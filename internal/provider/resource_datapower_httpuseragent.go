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

// This file is generated "gen/generator.go"
// !!CHANGES TO THIS FILE WILL BE OVERWRITTEN!!

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &HTTPUserAgentResource{}

func NewHTTPUserAgentResource() resource.Resource {
	return &HTTPUserAgentResource{}
}

type HTTPUserAgentResource struct {
	client *client.DatapowerClient
}

func (r *HTTPUserAgentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_httpuseragent"
}

func (r *HTTPUserAgentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("User agent", "user-agent", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP request-header", "identifier", "").String,
				Optional:            true,
			},
			"max_redirects": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max redirects", "max-redirects", "").AddIntegerRange(0, 128).AddDefaultValue("8").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 128),
				},
				Default: int64default.StaticInt64(8),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Timeout", "timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(300),
			},
			"proxy_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Proxy policy", "proxy", "").String,
				NestedObject:        models.DmProxyPolicyResourceSchema,
				Optional:            true,
			},
			"ssl_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS profile policy", "ssl", "").String,
				NestedObject:        models.DmSSLPolicyResourceSchema,
				Optional:            true,
			},
			"basic_auth_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Basic authentication policy", "basicauth", "").String,
				NestedObject:        models.DmBasicAuthPolicyResourceSchema,
				Optional:            true,
			},
			"soap_action_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAPAction policy", "soapaction", "").String,
				NestedObject:        models.DmSoapActionPolicyResourceSchema,
				Optional:            true,
			},
			"pubkey_auth_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Public key authentication policy", "pubkeyauth", "").String,
				NestedObject:        models.DmPubkeyAuthPolicyResourceSchema,
				Optional:            true,
			},
			"allow_compression_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow compression policy", "compression-policy", "").String,
				NestedObject:        models.DmAllowCompressionPolicyResourceSchema,
				Optional:            true,
			},
			"header_retention_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header retention policy", "headerretention-policy", "").String,
				NestedObject:        models.DmHeaderRetentionPolicyResourceSchema,
				Optional:            true,
			},
			"restrict10_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Restrict to HTTP/1.0 policy (deprecated)", "restrict-http-policy", "").String,
				NestedObject:        models.DmRestrict10PolicyResourceSchema,
				Optional:            true,
			},
			"http_version_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP version policy", "http-version-policy", "").String,
				NestedObject:        models.DmHTTPVersionPolicyResourceSchema,
				Optional:            true,
			},
			"add_header_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header injection policy", "add-header-policy", "").String,
				NestedObject:        models.DmAddHeaderPolicyResourceSchema,
				Optional:            true,
			},
			"upload_chunked_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Chunked upload policy", "chunked-uploads-policy", "").String,
				NestedObject:        models.DmUploadChunkedPolicyResourceSchema,
				Optional:            true,
			},
			"ftp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("FTP client policy", "ftp-policy", "").String,
				NestedObject:        models.DmFTPPolicyResourceSchema,
				Optional:            true,
			},
			"smtp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SMTP client policy", "smtp-policy", "").String,
				NestedObject:        models.DmSMTPPolicyResourceSchema,
				Optional:            true,
			},
			"sftp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SFTP client policy", "sftp-policy", "").String,
				NestedObject:        models.DmSFTPPolicyResourceSchema,
				Optional:            true,
			},
		},
	}
}

func (r *HTTPUserAgentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *HTTPUserAgentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.HTTPUserAgent

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `HTTPUserAgent`)
	res, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.HTTPUserAgent

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `HTTPUserAgent`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `HTTPUserAgent`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.HTTPUserAgent

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `HTTPUserAgent`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.HTTPUserAgent

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	resp.State.RemoveResource(ctx)
}
