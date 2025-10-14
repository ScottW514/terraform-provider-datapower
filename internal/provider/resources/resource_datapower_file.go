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

package resources

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

type File struct {
	AppDomain         types.String                `tfsdk:"app_domain"`
	RemotePath        types.String                `tfsdk:"remote_path"`
	LocalPath         types.String                `tfsdk:"local_path"`
	Content           types.String                `tfsdk:"content"`
	Hash              types.String                `tfsdk:"hash"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var _ resource.ResourceWithModifyPlan = &FileResource{}
var _ resource.ResourceWithValidateConfig = &FileResource{}

func NewFileResource() resource.Resource {
	return &FileResource{}
}

type FileResource struct {
	pData *tfutils.ProviderData
}

var localPathCondVal = validators.Evaluation{
	Evaluation: "property-null",
	Attribute:  "content",
	AttrType:   "String",
}

var localPathIgnoreVal = validators.Evaluation{
	Evaluation: "property-not-null",
	Attribute:  "content",
	AttrType:   "String",
}

var contentCondVal = validators.Evaluation{
	Evaluation: "property-null",
	Attribute:  "local_path",
	AttrType:   "String",
}

var contentIgnoreVal = validators.Evaluation{
	Evaluation: "property-not-null",
	Attribute:  "local_path",
	AttrType:   "String",
}

func (r *FileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The `datapower_file` resource manages files on an IBM DataPower Gateway, allowing practitioners to upload files to a specified remote path within an application domain. The resource supports uploading file content either by referencing a local file path or by providing the file content directly as a string. It tracks file changes using a computed hash, which is updated based on the local file or provided content. If the remote file is readable (e.g. not in a `cert` folder), the hash attribute is updated during the plan phase to reflect changes in the remote file. This resource is useful for managing stylesheets, gateway scripts, certificates, or other file based assets on the DataPower Gateway.", "", "").String,
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"remote_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File remote path", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^((shared|pub)?cert|chkpoints|config|image|local|policyframework|tasktemplates|temporary):///[a-zA-Z0-9_./-]+[a-zA-Z0-9_-]$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"local_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Path to local file, which will be uploaded. Not valid if `content` is set.", "", "").AddRequiredWhen(localPathCondVal.String()).AddNotValidWhen(localPathIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(localPathCondVal, localPathIgnoreVal, false),
				},
			},
			"content": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("String content of file (UTF-8 text only), which will be uploaded. Not valid if `local_file` is set. \n!> **Note:** This is string is stored in the state file and may not be appropriate for sensitive values.", "", "").AddRequiredWhen(contentCondVal.String()).AddNotValidWhen(contentIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(contentCondVal, contentIgnoreVal, false),
				},
			},
			"hash": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Provider calculated hash to track file changes.", "", "").String,
				Optional:            true,
				Computed:            true,
				DeprecationMessage:  "This attribute is for INTERNAL PROVIDER USE. Set values are ignored.",
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *FileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data File
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var fileData string
	if data.Content.IsUnknown() {
		data.Content = types.StringNull()
		var err error
		fileData, err = r.loadLocalFile(data.LocalPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s", err))
			return
		}
	} else {
		data.LocalPath = types.StringNull()
		fileData = base64.StdEncoding.EncodeToString([]byte(data.Content.ValueString()))
	}

	baseUrl, dpFileName, dpFullPath, dpStore, dpDirPath, dpDirs := data.getPathAttrs()

	// Create directory, if needed
	if len(dpDirs) > 0 {
		_, err := r.pData.Client.Post(baseUrl+"/"+dpStore, fmt.Sprintf(`{"directory": {"name": "%s"}}`, dpDirPath))
		if err != nil && !strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to create remote directory (%s), got error: %s", dpDirPath, err))
			return
		}
	}

	body := fmt.Sprintf(`{"file": {"name": "%s", "content": "%s"}}`, dpFileName, fileData)
	_, err := r.pData.Client.Post(fmt.Sprintf("%s/%s/%s", baseUrl, dpStore, dpDirPath), body)
	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to upload remote file (%s%s), got error: %s", baseUrl, dpFullPath, err))
		return
	} else if err != nil {
		_, err = r.pData.Client.Put(baseUrl+dpFullPath, body)
		if err != nil && !strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to upload (overwrite) remote file (%s%s), got error: %s", baseUrl, dpFullPath, err))
			return
		}
	}

	data.Hash = types.StringValue(r.generateHash(fileData))

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	var data File

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	baseUrl, _, dpFullPath, _, _, _ := data.getPathAttrs()

	remoteFileData, err := r.loadRemoteFile(baseUrl + dpFullPath)
	if err != nil && !strings.Contains(err.Error(), "status 403") && !strings.Contains(err.Error(), "status 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (%s), got error: %s", dpFullPath, err))
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
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	var data File

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}

	var fileData string
	if data.Content.IsUnknown() || data.Content.IsNull() {
		data.Content = types.StringNull()
		var err error
		fileData, err = r.loadLocalFile(data.LocalPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s", err))
			return
		}
	} else {
		data.LocalPath = types.StringNull()
		fileData = base64.StdEncoding.EncodeToString([]byte(data.Content.ValueString()))
	}
	baseUrl, dpFileName, dpFullPath, _, _, _ := data.getPathAttrs()

	body := fmt.Sprintf(`{"file": {"name": "%s", "content": "%s"}}`, dpFileName, fileData)
	_, err := r.pData.Client.Put(baseUrl+dpFullPath, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to update remote file (%s%s), got error: %s", baseUrl, dpFullPath, err))
		return
	}

	data.Hash = types.StringValue(r.generateHash(fileData))

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	var data File

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}

	baseUrl, _, dpFullPath, dpStore, _, dpDirs := data.getPathAttrs()

	_, err := r.pData.Client.Delete(baseUrl + dpFullPath)
	if err != nil && !strings.Contains(err.Error(), "status 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to delete remote file (%s%s), got error: %s", baseUrl, dpFullPath, err))
		return
	}

	dirCount := len(dpDirs)
	for {
		if dirCount <= 0 {
			break
		}
		dPath := fmt.Sprintf("%s/%s/%s", baseUrl, dpStore, strings.Join(dpDirs[0:dirCount], "/"))
		res, err := r.pData.Client.Get(dPath)
		if err != nil {
			if strings.Contains(err.Error(), "status 404") {
				break
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to read remote directory before deletion (%s), got error: %s", dPath, err))
			return
		}
		if files := res.Get("filestore.location.file"); files.Exists() {
			break
		} else if directories := res.Get("filestore.location.directory"); directories.Exists() {
			break
		}
		_, err = r.pData.Client.Delete(dPath)
		if err != nil {
			if strings.Contains(err.Error(), "status 404") {
				break
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to delete remote directory (%s), got error: %s", dPath, err))
			return
		}
		dirCount = dirCount - 1
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r *FileResource) generateHash(data string) string {
	encodedBytes := []byte(data)
	hasher := sha256.New()
	hasher.Write(encodedBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (r *FileResource) loadLocalFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to load local file (%s), got error: %s", path, err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func (r *FileResource) loadRemoteFile(path string) (string, error) {
	res, err := r.pData.Client.Get(path)
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
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
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
		var localFileData string
		if plan.Content.IsNull() {
			var err error
			localFileData, err = r.loadLocalFile(plan.LocalPath.ValueString())
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("%s", err))
				return
			}
		} else {
			localFileData = base64.StdEncoding.EncodeToString([]byte(plan.Content.ValueString()))
		}
		if plan.Hash.ValueString() != state.Hash.ValueString() {
			plan.Hash = types.StringUnknown()
		} else {
			plan.Hash = types.StringValue(r.generateHash(localFileData))
		}
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)

}

func (r *FileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data File

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

// Get Path attributes
// Returns:
//   - baseUrl: The base URL for the REST path
//   - dpFileName: Name of the file
//   - dpFullPath: Full path to the remote DataPower file, including the store.
//   - dpStore: The name of the DataPower file store, without slashes
//   - dpDirPath: The directory path within the DataPower file store, without leading or trailing slashes
//   - dpDirs: A slice containing each of the DataPower store directories, in order, without leading or trailing slashes
func (r File) getPathAttrs() (baseUrl string, dpFileName string, dpFullPath string, dpStore string, dpDirPath string, dpDirs []string) {
	baseUrl = fmt.Sprintf("/mgmt/filestore/%s", r.AppDomain.ValueString())
	dpFullPath = filepath.ToSlash("/" + strings.ReplaceAll(r.RemotePath.ValueString(), "://", ""))
	dpStore = strings.Split(dpFullPath, "/")[1]
	s := strings.Split(dpFullPath[1:], "/")
	dpFileName = s[len(s)-1]
	if len(s) > 2 {
		dpDirPath = strings.Join(s[1:len(s)-1], "/")
		dpDirs = strings.Split(dpDirPath, "/")
	}
	return
}
