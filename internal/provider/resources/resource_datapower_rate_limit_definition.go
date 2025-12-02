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
	_ resource.Resource                = &RateLimitDefinitionResource{}
	_ resource.ResourceWithImportState = &RateLimitDefinitionResource{}
)

func NewRateLimitDefinitionResource() resource.Resource {
	return &RateLimitDefinitionResource{}
}

type RateLimitDefinitionResource struct {
	pData *tfutils.ProviderData
}

func (r *RateLimitDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rate_limit_definition"
}

func (r *RateLimitDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A rate limit definition contains the parameters to enforce a rate limit.", "rate-limit-definition", "").String,
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
			"short_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name shown in rate limit response headers. If not specified, the object name is used in response headers.", "short-name", "").String,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Limit type", "type", "").AddStringEnum("rate", "burst", "count").AddDefaultValue("rate").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rate", "burst", "count"),
				},
				Default: stringdefault.StaticString("rate"),
			},
			"rate": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of requests that the API gateway can handle in an interval. The value of 0 indicates no limit.", "rate", "").AddIntegerRange(0, 4294967295).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval for the rate limit. Enter a value that is greater than or equal to 1. The default value is 1.", "interval", "").AddIntegerRange(1, 65535).AddDefaultValue("1").AddNotValidWhen(models.RateLimitDefinitionIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.RateLimitDefinitionIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"unit": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time unit for the rate limit. The default value is minute. When type is burst, the unit can be second or minute.", "unit", "").AddStringEnum("second", "minute", "hour", "day", "week").AddDefaultValue("minute").AddNotValidWhen(models.RateLimitDefinitionUnitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("second", "minute", "hour", "day", "week"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.RateLimitDefinitionUnitIgnoreVal, true),
				},
				Default: stringdefault.StaticString("minute"),
			},
			"hard_limit": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to reject requests when the specified rate limit is exceeded. By default, requests are rejected when the limit is exceeded. When disabled, requests are accepted but a warning is logged.", "hard-limit", "").AddDefaultValue("true").AddNotValidWhen(models.RateLimitDefinitionHardLimitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"is_client": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to apply the rate limit to the client or to an internal component. By default, the rate limit is applied to the client. Client rate limits return a 429 error when exceeded. When disabled, rate limit information is not applied to the client. Non-client rate limits return a 503 error when exceeded.", "is-client", "").AddDefaultValue("true").AddNotValidWhen(models.RateLimitDefinitionIsClientIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"use_api_name": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the API name in the rate limit key. By default, the API name is not included. When enabled, the API name is included.", "use-api-name", "").AddDefaultValue("false").AddNotValidWhen(models.RateLimitDefinitionUseApiNameIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_app_id": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the application ID in the rate limit key. By default, the application ID is not included. When enabled, the application ID is included.", "use-app-id", "").AddDefaultValue("false").AddNotValidWhen(models.RateLimitDefinitionUseAppIdIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_client_id": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the client ID in the rate limit key. By default, the client ID is not included. When enabled, the client ID is included.", "use-client-id", "").AddDefaultValue("false").AddNotValidWhen(models.RateLimitDefinitionUseClientIdIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"auto_replenish": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the count limit is automatically replenished at the end of the transaction. By default, the count limit is automatically replenished. When disabled, the count limit is replenished only by applying a rate limit assembly action that contains the count limit with a replenish operation.", "auto-replenish", "").AddDefaultValue("true").AddNotValidWhen(models.RateLimitDefinitionAutoReplenishIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dynamic_value": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the dynamic value string for the rate limit, which should contain one or more context variables. The default value is an empty string. <p>The dynamic value makes it possible to use a context variable to enforce the rate limit based on parameters other than those defined in the rate limit scheme, such as a username, incoming IP address, or server name. The context variable can be set in a GatewayScript action and then included in the dynamic value.</p><p>The following example uses the context object in a GatewayScript action to add the <tt>my.server</tt> variable to the API context. The dynamic value can then include the variable <tt>my.server</tt> , which resolves to the server name <tt>server34</tt> .</p><p><tt>context.set(\"my.server\", \"server34\")</tt></p>", "dynamic-value", "").AddNotValidWhen(models.RateLimitDefinitionDynamicValueIgnoreVal.String()).String,
				Optional:            true,
			},
			"weight": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a JSONata expression to assign a weight value to the rate limit. For each API call, the value computed by the weight expression is applied to the rate limit. The default value is 1. If the weight expression evaluates to a value that is less than or equal to 0, it is set to 1. An empty string results in an error.", "weight", "").AddDefaultValue("1").AddNotValidWhen(models.RateLimitDefinitionWeightIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("1"),
			},
			"response_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether response headers include rate limit information. By default, headers include rate limit information. When disabled, headers exclude rate limit information.", "response-headers", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"emulate_burst_headers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to return information about the rate limit in burst limit response headers instead of in rate limit response headers. By default, the information is in rate limit headers. When enabled, information is in burst limit headers.", "emulate-burst-headers", "").AddDefaultValue("false").AddNotValidWhen(models.RateLimitDefinitionEmulateBurstHeadersIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_interval_offset": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to allow limit intervals to start at different offsets. By default, intervals can start at different offsets. When disabled, intervals cannot start at different offsets.", "use-interval-offset", "").AddDefaultValue("true").AddNotValidWhen(models.RateLimitDefinitionUseIntervalOffsetIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_cache_fallback": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the cache as a fallback when gateway-peering instances cannot be contacted. By default, the cache can enforce rate limits when the cache is disabled. When disabled, the cache cannot enforce rate limits.", "allow-cache-fallback", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"use_cache": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use a cache to store rate limit information. A cache might be faster when the number of API calls is low. A cache can cause degraded performance when the number of API calls is exceptionally high. By default, a cache cannot store information. When enabled, the cache can store information.", "use-cache", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"parameters": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parameters", "parameter", "").String,
				NestedObject:        models.GetDmRateLimitDefinitionNameValuePairResourceSchema(),
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *RateLimitDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *RateLimitDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RateLimitDefinition
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

	body := data.ToBody(ctx, `RateLimitDefinition`)
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

func (r *RateLimitDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RateLimitDefinition
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

	data.UpdateFromBody(ctx, `RateLimitDefinition`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RateLimitDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RateLimitDefinition
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
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `RateLimitDefinition`), data.ProviderTarget)
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

func (r *RateLimitDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.RateLimitDefinition
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

func (r *RateLimitDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.RateLimitDefinition
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

	data.FromBody(ctx, `RateLimitDefinition`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
