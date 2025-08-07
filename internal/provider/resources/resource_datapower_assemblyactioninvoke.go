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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AssemblyActionInvokeResource{}

func NewAssemblyActionInvokeResource() resource.Resource {
	return &AssemblyActionInvokeResource{}
}

type AssemblyActionInvokeResource struct {
	client *client.DatapowerClient
}

func (r *AssemblyActionInvokeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assemblyactioninvoke"
}

func (r *AssemblyActionInvokeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Invoke assembly action", "assembly-invoke", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
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
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL", "url", "").String,
				Required:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "").String,
				Optional:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Timeout", "timeout", "").AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "password", "passwordalias").String,
				Optional:            true,
			},
			"method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP method", "method", "").AddStringEnum("Keep", "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE").AddDefaultValue("Keep").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Keep", "GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE"),
				},
				Default: stringdefault.StaticString("Keep"),
			},
			"backend_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backend type", "backend-type", "").AddStringEnum("detect", "xml", "json", "binary", "graphql").AddDefaultValue("detect").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("detect", "xml", "json", "binary", "graphql"),
				},
				Default: stringdefault.StaticString("detect"),
			},
			"graph_ql_send_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("GraphQL send type", "graphql-send-type", "").AddStringEnum("detect", "graphql", "json").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("detect", "graphql", "json"),
				},
			},
			"compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable compression", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache type", "cache-type", "").AddStringEnum("Protocol", "NoCache", "TimeToLive").AddDefaultValue("Protocol").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Protocol", "NoCache", "TimeToLive"),
				},
				Default: stringdefault.StaticString("Protocol"),
			},
			"time_to_live": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Time to live", "ttl", "").AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(900),
			},
			"cache_unsafe_response": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache response to POST and PUT requests", "cache-unsafe-response", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cache_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache key", "cache-key", "").String,
				Optional:            true,
			},
			"follow_redirects": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Follow redirects", "follow-redirects", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP version to server", "http-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1", "HTTP/2").AddDefaultValue("HTTP/1.1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1", "HTTP/2"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"http2_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP/2 required", "http2-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"do_chunked_upload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow chunked uploads", "chunked-uploads", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"persistent_connection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistent connection", "persistent-connection", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stop on error", "stop-on-error", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"error_types": models.GetDmInvokeErrorTypeResourceSchema("Error types", "error-types", "", false),
			"output": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Output", "output", "").AddDefaultValue("message").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("message"),
			},
			"decode_request_params": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Decode request parameters", "decode-request-params", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encode_plus_char": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encode + characters in query", "encode-plus-char", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"keep_payload": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Keep payload", "keep-payload", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"inject_user_agent_header": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inject User-Agent header", "inject-user-agent", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"inject_proxy_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inject proxy headers", "inject-proxy-headers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"header_control_list": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header control list", "header-control-list", "controllist").AddDefaultValue("default-accept-all").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-accept-all"),
			},
			"parameter_control_list": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parameter control list", "parameter-control-list", "controllist").AddDefaultValue("default-reject-all").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default-reject-all"),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Title", "title", "").String,
				Optional:            true,
			},
			"correlation_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Correlation path", "correlation-path", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable debugging", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AssemblyActionInvokeResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *AssemblyActionInvokeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)

	body := data.ToBody(ctx, `AssemblyActionInvoke`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `AssemblyActionInvoke`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AssemblyActionInvoke`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AssemblyActionInvoke`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionInvokeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *AssemblyActionInvokeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AssemblyActionInvoke

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
