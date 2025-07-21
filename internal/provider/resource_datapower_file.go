// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"strings"

	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type File struct {
	AppDomain  types.String `tfsdk:"app_domain"`
	RemotePath types.String `tfsdk:"remote_path"`
	LocalPath  types.String `tfsdk:"local_path"`
	Hash       types.String `tfsdk:"hash"`
}

var _ resource.ResourceWithModifyPlan = &FileResource{}

func NewFileResource() resource.Resource {
	return &FileResource{}
}

type FileResource struct {
	client *client.DatapowerClient
}

func (r *FileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Files", "", "").String,

		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"remote_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File remote path", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"local_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Path to local file, which will be uploaded.", "", "").String,
				Required:            true,
			},
			"hash": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Provider calculated hash to track file changes.", "", "").String,
				Optional:            true,
				Computed:            true,
				DeprecationMessage:  "This attribute is for INTERNAL PROVIDER USE. Set values are ignored.",
			},
		},
	}
}

func (r *FileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data File

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	fileData, err := r.loadLocalFile(data.LocalPath.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s", err))
		return
	}
	fPath := "/" + strings.ReplaceAll(data.RemotePath.ValueString(), "://", "")
	bPath := fmt.Sprintf("/mgmt/filestore/%s", data.AppDomain.ValueString())
	dir := filepath.ToSlash(filepath.Dir(fPath))

	// Create directory, if needed
	if len(strings.Split(dir, "/")) > 2 {
		res, err := r.client.Post(bPath+dir, fmt.Sprintf(`{"directory": {"name": "%s"}}`, dir))
		if err != nil && res.RawResponse.StatusCode != 409 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to create remote directory (%s), got error: %s", dir, err))
			return
		}
	}

	body := fmt.Sprintf(`{"file": {"name": "%s", "content": "%s"}}`, filepath.Base(fPath), fileData)
	_, err = r.client.Post(bPath+dir, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to upload remote file (%s), got error: %s", bPath+dir, err))
		return
	}

	data.Hash = types.StringValue(r.generateHash(fileData))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data File

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	remoteFileData, err := r.loadRemoteFile(data.AppDomain.ValueString(), data.RemotePath.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 403") && !strings.Contains(err.Error(), "status 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	} else if err != nil && strings.Contains(err.Error(), "status 403") {
		// Ignore 403 forbidden, as we cannot read files in some protected directorys, and must assume they are unchanged on the remote end
	} else if err != nil && strings.Contains(err.Error(), "status 404") {
		resp.State.RemoveResource(ctx)
		return
	} else {
		data.Hash = types.StringValue(r.generateHash(remoteFileData))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data File

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	fileData, err := r.loadLocalFile(data.LocalPath.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s", err))
		return
	}
	fPath := "/" + strings.ReplaceAll(data.RemotePath.ValueString(), "://", "")
	bPath := fmt.Sprintf("/mgmt/filestore/%s", data.AppDomain.ValueString())

	body := fmt.Sprintf(`{"file": {"name": "%s", "content": "%s"}}`, filepath.Base(fPath), fileData)
	_, err = r.client.Put(bPath+fPath, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to update remote file (%s), got error: %s\n%s", bPath+fPath, err, body))
		return
	}

	data.Hash = types.StringValue(r.generateHash(fileData))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data File

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := fmt.Sprintf("/mgmt/filestore/%s/%s", data.AppDomain.ValueString(), strings.ReplaceAll(data.RemotePath.ValueString(), "://", ""))
	res, err := r.client.Delete(path)
	if err != nil && res.RawResponse.StatusCode != 404 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to delete remote file (%s), got error: %s", path, err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *FileResource) generateHash(data string) string {
	encodedBytes := []byte(data)
	table := crc32.MakeTable(crc32.Castagnoli)
	hasher := crc32.New(table)
	hasher.Write(encodedBytes)
	checksum := hasher.Sum32()
	return fmt.Sprintf("%08x", checksum)
}

func (r *FileResource) loadLocalFile(localPath string) (string, error) {
	data, err := os.ReadFile(localPath)
	if err != nil {
		return "", fmt.Errorf("failed to load local file (%s), got error: %s", localPath, err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func (r *FileResource) loadRemoteFile(domain string, remotePath string) (string, error) {
	path := fmt.Sprintf("/mgmt/filestore/%s/%s", domain, strings.ReplaceAll(remotePath, "://", ""))
	res, err := r.client.Get(path)
	if err != nil {
		return "", fmt.Errorf("failed to download remote file (%s), got error: %s", path, err)
	}
	if data := res.Get("file"); data.Exists() {
		return data.String(), nil
	} else {
		return "", fmt.Errorf("failed to download remote file (%s), invalid response", path)
	}
}

func (r *FileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var state, plan File

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Hash.IsUnknown() {
		localFileData, err := r.loadLocalFile(plan.LocalPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
		if plan.Hash.ValueString() != state.Hash.ValueString() {
			plan.Hash = types.StringUnknown()
		} else {
			plan.Hash = types.StringValue(r.generateHash(localFileData))
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)

}
