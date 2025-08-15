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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &XSLCoprocServiceResource{}
var _ resource.ResourceWithValidateConfig = &XSLCoprocServiceResource{}

func NewXSLCoprocServiceResource() resource.Resource {
	return &XSLCoprocServiceResource{}
}

type XSLCoprocServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *XSLCoprocServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xsl_coproc_service"
}

func (r *XSLCoprocServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>The XSL Coprocessor is obsolete.</p>", "xslcoproc", "").AddActions("quiesce").String,
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
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "port", "").AddIntegerRange(1, 65535).String,
				Required:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"url_rewrite_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "urlrewrite-policy", "url_rewrite_policy").String,
				Optional:            true,
			},
			"style_policy_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "stylesheet-rule", "style_policy_rule").String,
				Optional:            true,
			},
			"connection_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "connection-timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(3, 7200),
				},
				Default: int64default.StaticInt64(60),
			},
			"intermediate_result_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "intermediate-result-timeout", "").AddIntegerRange(1, 600).AddDefaultValue("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(20),
			},
			"cache_relative_url": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "cache-relative-url", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_client_uri_resolver": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "use-client-resolver", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"crypto_extensions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "crypto-extensions", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"default_param_namespace": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "default-param-namespace", "").AddDefaultValue("http://www.datapower.com/param/config").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("http://www.datapower.com/param/config"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 250),
				},
				Default: int64default.StaticInt64(25),
			},
			"debug_trigger": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "debug-trigger", "").String,
				NestedObject:        models.DmMSDebugTriggerTypeResourceSchema,
				Optional:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-server", "ssl_server_profile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "ssl-sni-server", "ssl_sni_server_profile").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter a host alias or the IP address that the service listens on. Host aliases can ease migration tasks among appliances.</p><ul><li>0 or 0.0.0.0 indicates all configured IPv4 addresses.</li><li>:: indicates all configured IPv4 and IPv6 addresses.</li></ul><p><b>Attention:</b> For management services, the value of 0.0.0.0 or :: is a security risk. Use an explicit IP address to isolate management traffic from application data traffic.</p>", "ip-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *XSLCoprocServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *XSLCoprocServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLCoprocService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XSLCoprocService`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XSLCoprocServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLCoprocService
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

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `XSLCoprocService`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `XSLCoprocService`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XSLCoprocServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLCoprocService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XSLCoprocService`))
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

func (r *XSLCoprocServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XSLCoprocService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *XSLCoprocServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.XSLCoprocService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
