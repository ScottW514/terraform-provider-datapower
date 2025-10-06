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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &HTTPUserAgentResource{}
var _ resource.ResourceWithImportState = &HTTPUserAgentResource{}

func NewHTTPUserAgentResource() resource.Resource {
	return &HTTPUserAgentResource{}
}

type HTTPUserAgentResource struct {
	pData *tfutils.ProviderData
}

func (r *HTTPUserAgentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http_user_agent"
}

func (r *HTTPUserAgentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A user agent define how to retrieve resources from remote servers.", "user-agent", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the string that the user agent includes as the <tt>request-header</tt> field. This field contains information about the user agent that initiates the request. By default, the system does not include a <tt>request-header</tt> field.", "identifier", "").String,
				Optional:            true,
			},
			"max_redirects": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of HTTP redirect messages received before the target URL is declared unreachable. Enter a value in the range 0 - 128. The default value is 8.", "max-redirects", "").AddIntegerRange(0, 128).AddDefaultValue("8").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 128),
				},
				Default: int64default.StaticInt64(8),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("the maximum idle time in seconds before an established connection to a remote server is torn down. Enter a avlue in the range 1 - 86400. The default value is 300.", "timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(300),
			},
			"proxy_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the proxy policy that associates a set of URLs with a specific HTTP proxy.", "proxy", "").String,
				NestedObject:        models.GetDmProxyPolicyResourceSchema(),
				Optional:            true,
			},
			"ssl_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLs with a specific TLS profile. When a URL matches the expression, the agent uses the corresponding TLS profile to secure connections with the resource.", "ssl", "").String,
				NestedObject:        models.GetDmSSLPolicyResourceSchema(),
				Optional:            true,
			},
			"basic_auth_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLs with a specific username and password for basic authentication.", "basicauth", "").String,
				NestedObject:        models.GetDmBasicAuthPolicyResourceSchema(),
				Optional:            true,
			},
			"soap_action_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLs with a specific HTTP SOAPAction header.", "soapaction", "").String,
				NestedObject:        models.GetDmSoapActionPolicyResourceSchema(),
				Optional:            true,
			},
			"pubkey_auth_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLs with a specific private key for public key authentication. The remote host must possess and reference the corresponding public key (certificate) to connect successfully.", "pubkeyauth", "").String,
				NestedObject:        models.GetDmPubkeyAuthPolicyResourceSchema(),
				Optional:            true,
			},
			"allow_compression_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLS that allow compression.", "compression-policy", "").String,
				NestedObject:        models.GetDmAllowCompressionPolicyResourceSchema(),
				Optional:            true,
			},
			"header_retention_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLS to retain specific heads in messages.", "headerretention-policy", "").String,
				NestedObject:        models.GetDmHeaderRetentionPolicyResourceSchema(),
				Optional:            true,
			},
			"http_version_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLs to specific HTTP versions. This policy is cumulative. If any transaction, URL match, or gateway have an HTTP version policy, that transaction is processed at the requested HTTP version.", "http-version-policy", "").String,
				NestedObject:        models.GetDmHTTPVersionPolicyResourceSchema(),
				Optional:            true,
			},
			"add_header_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLS to inject HTTP headers into the message.", "add-header-policy", "").String,
				NestedObject:        models.GetDmAddHeaderPolicyResourceSchema(),
				Optional:            true,
			},
			"upload_chunked_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URL to control whether to send chunked-encoded documents. With HTTP/1.1, the body of the document can be delimited by <tt>Content-Length</tt> or chunked encoding. All servers understand <tt>Content-Length</tt> but many applications fail to understand chunked encoding. Therefore, <tt>Content-Length</tt> is used. However, the use of <tt>Content-Length</tt> interferes with the ability of the service to fully stream. <p>Unlike all other HTTP/1.1 features that can be negotiated down at run time, you must know beforehand that the target server is RFC 2616 compatible.</p>", "chunked-uploads-policy", "").String,
				NestedObject:        models.GetDmUploadChunkedPolicyResourceSchema(),
				Optional:            true,
			},
			"ftp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associate a set of URLs to control FTP client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that initiates the file transfer.", "ftp-policy", "").String,
				NestedObject:        models.GetDmFTPPolicyResourceSchema(),
				Optional:            true,
			},
			"smtp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associates a set of URLS to control SMTP client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that sends the e-mail message.", "smtp-policy", "").String,
				NestedObject:        models.GetDmSMTPPolicyResourceSchema(),
				Optional:            true,
			},
			"sftp_policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the policy that associate a set of URLs to control SSH client options for outgoing connections. These settings override the compiled-in defaults and can be further overridden by query parameters that initiates the file transfer.", "sftp-policy", "").String,
				NestedObject:        models.GetDmSFTPPolicyResourceSchema(),
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *HTTPUserAgentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *HTTPUserAgentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPUserAgent
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `HTTPUserAgent`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPUserAgent
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	data.UpdateFromBody(ctx, `HTTPUserAgent`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPUserAgent
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `HTTPUserAgent`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.HTTPUserAgent
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *HTTPUserAgentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.HTTPUserAgent
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `HTTPUserAgent`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HTTPUserAgentResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.HTTPUserAgent

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
