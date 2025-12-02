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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var (
	_ resource.Resource                   = &WebTokenServiceResource{}
	_ resource.ResourceWithValidateConfig = &WebTokenServiceResource{}
	_ resource.ResourceWithImportState    = &WebTokenServiceResource{}
)

func NewWebTokenServiceResource() resource.Resource {
	return &WebTokenServiceResource{}
}

type WebTokenServiceResource struct {
	pData *tfutils.ProviderData
}

func (r *WebTokenServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_token_service"
}

func (r *WebTokenServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The web token service provides on-box HTTP services.", "web-token-service", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a brief, descriptive comment.", "summary", "").String,
				Optional:            true,
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the service-level priority. When resources are in high demand, the appliance handles transactions for higher priority services over lower priority services.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assign the XML manager that obtains and manages documents on behalf of the service. The user agent, which is referenced by this XML manager, controls access to remote destinations. Unless you have service-specific needs, use the <tt>default</tt> XML manager.", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"request_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Characterizes the traffic that originates from the client. The default is Non-XML.", "request-type", "").AddStringEnum("soap", "xml", "unprocessed", "preprocessed", "json").AddDefaultValue("preprocessed").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap", "xml", "unprocessed", "preprocessed", "json"),
				},
				Default: stringdefault.StaticString("preprocessed"),
			},
			"front_side": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Define the addresses and ports that the service listens on. You can define listening sources for HTTP and HTTPS requests. For HTTPS request, set the listening source to use TLS parameter and assign the TLS profile.", "listen-on", "").String,
				NestedObject:        models.GetDmSSLFrontSideResourceSchema(),
				Required:            true,
			},
			"style_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Assign the processing policy to the service. The processing policy defines the actions to perform on security token requests.", "stylepolicy", "style_policy").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"rewrite_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error messages after a decryption action can provide an attacker who is using the padding oracle attack method with enough information to determine the contents of the plain-text data. When enabled, the default, the client receives error messages without the internal information that could lead to a discovery. When disabled, the client receives the original message with this information.", "rewrite-errors", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The timing difference of the error messages returned after a decryption action can provide an attacker with enough information to determine the contents of the plain-text data. When enabled, the default, the appliance delays error messages for the defined duration. When disabled, the appliance does not delay error messages.", "delay-errors", "").AddDefaultValue("true").AddNotValidWhen(models.WebTokenServiceDelayErrorsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delay_errors_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When enabling the delay of error messages, specify the delay duration in milliseconds. If delaying messages for 3000 ms, the appliance will not send error messages to the client until 3 seconds have elapsed since the appliance performed decryption on the requests. Enter any value in the range 250 - 300000. The default value is 1000.", "delay-errors-duration", "").AddIntegerRange(250, 300000).AddDefaultValue("1000").AddRequiredWhen(models.WebTokenServiceDelayErrorsDurationCondVal.String()).AddNotValidWhen(models.WebTokenServiceDelayErrorsDurationIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(250, 300000),
					validators.ConditionalRequiredInt64(models.WebTokenServiceDelayErrorsDurationCondVal, models.WebTokenServiceDelayErrorsDurationIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1000),
			},
			"front_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the intra-transaction timeout for Web Token Service to client connections. This value is the maximum idle time to allow in a transaction on the Web Token Service to client connection. This timer monitors idle time in the data transfer process. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 120.", "front-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"front_persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the inter-transaction timeout for Web Token Service to client connections. This value is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction on the Web Token Service to client connection. If the specified idle time is exceeded, the connection is torn down. Enter a value in the range 0 - 86400. The default value is 180. A value of 0 disables persistent connections.", "front-persistent-timeout", "").AddIntegerRange(0, 86400).AddDefaultValue("180").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"front_http_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the HTTP version for client-to-service connections. If the client submits an HTTP/1.0 request, the service on the DataPower appliance always replies with HTTP/1.0 compatible responses regardless of this setting. The default version is HTTP 1.1.", "http-front-version", "").AddStringEnum("HTTP/1.0", "HTTP/1.1").AddDefaultValue("HTTP/1.1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("HTTP/1.0", "HTTP/1.1"),
				},
				Default: stringdefault.StaticString("HTTP/1.1"),
			},
			"http_client_ip_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the HTTP header that contains the client IP address. When defined, the client IP address is read from this HTTP header. This IP address will then be used for monitoring and logging. Retain the default value <tt>X-Client-IP</tt> or specify another value; for example, <tt>X-Forwarded-For</tt> .", "http-client-ip-label", "").AddDefaultValue("X-Client-IP").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Client-IP"),
			},
			"http_log_cor_id_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of an HTTP Header to read to determine the global transaction ID for chained services. This value defaults to X-Global-Transaction-ID.", "http-global-tranID-label", "").AddDefaultValue("X-Global-Transaction-ID").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-Global-Transaction-ID"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the diagnostic mode for processing policies. When enabled, you can view details about the state of variables and contexts for a captured transaction in the probe. The default value is <tt>off</tt> .</p><p>Transaction diagnostic mode is not intended for use in a production environment. Transaction diagnostic mode consumes significant resources that can slow down transaction processing.</p>", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Set the number of records for transaction diagnostic mode in the probe. Enter a value in the range 10 - 250. The default value is 25.", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").AddRequiredWhen(models.WebTokenServiceDebugHistoryCondVal.String()).AddNotValidWhen(models.WebTokenServiceDebugHistoryIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 250),
					validators.ConditionalRequiredInt64(models.WebTokenServiceDebugHistoryCondVal, models.WebTokenServiceDebugHistoryIgnoreVal, true),
				},
				Default: int64default.StaticInt64(25),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebTokenServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebTokenServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebTokenService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebTokenService`)
	_, err := r.pData.Client.Post(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebTokenServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebTokenService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `WebTokenService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebTokenServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebTokenService
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebTokenService`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebTokenServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebTokenService
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *WebTokenServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.WebTokenService
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `WebTokenService`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *WebTokenServiceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebTokenService

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
