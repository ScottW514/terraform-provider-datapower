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

var _ resource.Resource = &AssemblyActionUserSecurityResource{}
var _ resource.ResourceWithImportState = &AssemblyActionUserSecurityResource{}

func NewAssemblyActionUserSecurityResource() resource.Resource {
	return &AssemblyActionUserSecurityResource{}
}

type AssemblyActionUserSecurityResource struct {
	pData *tfutils.ProviderData
}

func (r *AssemblyActionUserSecurityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_user_security"
}

func (r *AssemblyActionUserSecurityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The user security assembly actions extracts identity, authenticates, and authorizes end users.", "assembly-user-security", "").String,
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
			"factor_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Factor identifier", "factor-id", "").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"extract_identity_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identity extraction method", "extract-identity-method", "").AddStringEnum("disabled", "basic", "context-var", "html-form", "redirect").AddDefaultValue("basic").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "basic", "context-var", "html-form", "redirect"),
				},
				Default: stringdefault.StaticString("basic"),
			},
			"ei_stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to stop processing if identity extraction fails. If failed, stops the assembly and return an error.", "ei-stop-on-error", "").AddDefaultValue("true").AddRequiredWhen(models.AssemblyActionUserSecurityEIStopOnErrorCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityEIStopOnErrorIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_context_variable": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username context variable", "user-context-var", "").AddRequiredWhen(models.AssemblyActionUserSecurityUserContextVariableCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityUserContextVariableIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityUserContextVariableCondVal, models.AssemblyActionUserSecurityUserContextVariableIgnoreVal, false),
				},
			},
			"pass_context_variable": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password context variable", "pass-context-var", "").AddRequiredWhen(models.AssemblyActionUserSecurityPassContextVariableCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityPassContextVariableIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityPassContextVariableCondVal, models.AssemblyActionUserSecurityPassContextVariableIgnoreVal, false),
				},
			},
			"redirect_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment to redirect the request to obtain user credentials. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "redirect-url", "").AddRequiredWhen(models.AssemblyActionUserSecurityRedirectURLCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityRedirectURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityRedirectURLCondVal, models.AssemblyActionUserSecurityRedirectURLIgnoreVal, false),
				},
			},
			"redirect_time_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds for a transaction to complete before the redirect fails. Enter a value in the range 10 - 6000. The default value is 300.", "redirect-time-limit", "").AddIntegerRange(10, 6000).AddDefaultValue("300").AddRequiredWhen(models.AssemblyActionUserSecurityRedirectTimeLimitCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityRedirectTimeLimitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 6000),
					validators.ConditionalRequiredInt64(models.AssemblyActionUserSecurityRedirectTimeLimitCondVal, models.AssemblyActionUserSecurityRedirectTimeLimitIgnoreVal, true),
				},
				Default: int64default.StaticInt64(300),
			},
			"query_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Query parameters", "query-parameters", "").AddNotValidWhen(models.AssemblyActionUserSecurityQueryParametersIgnoreVal.String()).String,
				Optional:            true,
			},
			"ei_default_form": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the default form or a custom form. When enabled, returns the default login page to obtain credentials. When disabled, define the configuration to return the custom login page.", "ei-default-form", "").AddDefaultValue("true").AddRequiredWhen(models.AssemblyActionUserSecurityEIDefaultFormCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityEIDefaultFormIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ei_custom_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment of the custom login page. This page collects user name and password information. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "ei-custom-form", "").AddRequiredWhen(models.AssemblyActionUserSecurityEICustomFormCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityEICustomFormIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityEICustomFormCondVal, models.AssemblyActionUserSecurityEICustomFormIgnoreVal, false),
				},
			},
			"ei_custom_form_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom form TLS client profile", "ei-custom-form-tls-client-profile", "ssl_client_profile").AddNotValidWhen(models.AssemblyActionUserSecurityEICustomFormClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"ei_custom_form_content_security_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value to use for the HTTP <tt>Content-Security-Policy</tt> response header for the custom login page. This response header allows you to control which resources the user agent can load. Generally, you set server origins and script endpoints to detect and mitigate cross-site scripting (XSS), clickjacking, and other injection attacks.", "ei-custom-form-csp", "").AddDefaultValue("default-src 'self'").AddNotValidWhen(models.AssemblyActionUserSecurityEICustomFormContentSecurityPolicyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default-src 'self'"),
			},
			"ei_form_time_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds for a transaction to complete before the identity extraction request fails. Enter a value in the range 10 - 600. The default value is 300.", "ei-form-time-limit", "").AddIntegerRange(10, 600).AddDefaultValue("300").AddRequiredWhen(models.AssemblyActionUserSecurityEIFormTimeLimitCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityEIFormTimeLimitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 600),
					validators.ConditionalRequiredInt64(models.AssemblyActionUserSecurityEIFormTimeLimitCondVal, models.AssemblyActionUserSecurityEIFormTimeLimitIgnoreVal, true),
				},
				Default: int64default.StaticInt64(300),
			},
			"user_auth_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authentication method", "user-auth-method", "").AddStringEnum("disabled", "user-registry").AddDefaultValue("user-registry").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "user-registry"),
				},
				Default: stringdefault.StaticString("user-registry"),
			},
			"au_stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to stop processing if authentication fails. If failed, stops the assembly and return an error.", "au-stop-on-error", "").AddDefaultValue("true").AddRequiredWhen(models.AssemblyActionUserSecurityAUStopOnErrorCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityAUStopOnErrorIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_registry": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API user registry to authenticate incoming API requests. The supported registries are API authentication URL and API LDAP.", "user-registry", "").AddRequiredWhen(models.AssemblyActionUserSecurityUserRegistryCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityUserRegistryIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityUserRegistryCondVal, models.AssemblyActionUserSecurityUserRegistryIgnoreVal, false),
				},
			},
			"auth_response_headers_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the regular expression to select which response headers to add to the API context for access by subsequent actions. The default value is a case-insensitive search on the <tt>x-api</tt> prefix. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "auth-response-headers-pattern", "").AddDefaultValue("(?i)x-api*").AddNotValidWhen(models.AssemblyActionUserSecurityAuthResponseHeadersPatternIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("(?i)x-api*"),
			},
			"auth_response_credential_header": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the response header that contains the authenticated credentials. The default value is <tt>X-API-Authenticated-Credential</tt> . The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "auth-response-header-credential", "").AddDefaultValue("X-API-Authenticated-Credential").AddNotValidWhen(models.AssemblyActionUserSecurityAuthResponseCredentialHeaderIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("X-API-Authenticated-Credential"),
			},
			"user_az_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization method", "user-az-method", "").AddStringEnum("disabled", "authenticated", "html-form").AddDefaultValue("authenticated").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "authenticated", "html-form"),
				},
				Default: stringdefault.StaticString("authenticated"),
			},
			"az_stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to stop processing if authorization fails. If failed, stops the assembly and return an error.", "az-stop-on-error", "").AddDefaultValue("true").AddRequiredWhen(models.AssemblyActionUserSecurityAZStopOnErrorCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityAZStopOnErrorIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"az_default_form": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use the default form or a custom form. When enabled, returns the default authorization page to obtain authorization. When disabled, define the configuration to return the custom authorization page.", "az-default-form", "").AddDefaultValue("true").AddRequiredWhen(models.AssemblyActionUserSecurityAZDefaultFormCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityAZDefaultFormIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"az_custom_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment of the custom authorization page. This page obtains permission from the end user. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.", "az-custom-form", "").AddRequiredWhen(models.AssemblyActionUserSecurityAZCustomFormCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityAZCustomFormIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AssemblyActionUserSecurityAZCustomFormCondVal, models.AssemblyActionUserSecurityAZCustomFormIgnoreVal, false),
				},
			},
			"az_custom_form_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom form TLS client profile", "az-custom-form-tls-client-profile", "ssl_client_profile").AddNotValidWhen(models.AssemblyActionUserSecurityAZCustomFormClientProfileIgnoreVal.String()).String,
				Optional:            true,
			},
			"az_custom_form_content_security_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the HTTP <tt>Content-Security-Policy</tt> response header for the custom authorization page. This response header allows you to control which resources the user agent can load. Generally, you set server origins and script endpoints to detect and mitigate cross-site scripting (XSS), clickjacking, and other injection attacks.", "az-custom-form-csp", "").AddDefaultValue("default-src 'self'").AddNotValidWhen(models.AssemblyActionUserSecurityAZCustomFormContentSecurityPolicyIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("default-src 'self'"),
			},
			"az_form_time_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds for a transaction to complete before the authorization request fails. Enter a value in the range 10 - 600. The default value is 300.", "az-form-time-limit", "").AddIntegerRange(10, 600).AddDefaultValue("300").AddRequiredWhen(models.AssemblyActionUserSecurityAZFormTimeLimitCondVal.String()).AddNotValidWhen(models.AssemblyActionUserSecurityAZFormTimeLimitIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 600),
					validators.ConditionalRequiredInt64(models.AssemblyActionUserSecurityAZFormTimeLimitCondVal, models.AssemblyActionUserSecurityAZFormTimeLimitIgnoreVal, true),
				},
				Default: int64default.StaticInt64(300),
			},
			"az_table_display_checkboxes": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Display table check boxes", "az-table-display-checkboxes", "").AddDefaultValue("false").AddNotValidWhen(models.AssemblyActionUserSecurityAZTableDisplayCheckboxesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"az_table_dynamic_entries": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the period-delimited context variable that adds dynamic entries to display. This context variable supports space delimited names, a JSON array of names, or a JSON array of objects with name and description.", "az-table-dynamic-entries", "").AddDefaultValue("user.default.az.dynamic_entries").AddNotValidWhen(models.AssemblyActionUserSecurityAZTableDynamicEntriesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("user.default.az.dynamic_entries"),
			},
			"az_table_default_entry": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default table entry", "az-table-default-entry", "").AddNotValidWhen(models.AssemblyActionUserSecurityAZTableDefaultEntryIgnoreVal.String()).String,
				NestedObject:        models.GetDmTableEntryResourceSchema(),
				Optional:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Hostname", "hostname", "").AddNotValidWhen(models.AssemblyActionUserSecurityHostnameIgnoreVal.String()).String,
				Optional:            true,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.", "correlation-path", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>", "debug", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AssemblyActionUserSecurityResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AssemblyActionUserSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionUserSecurity
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AssemblyActionUserSecurity`)
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

func (r *AssemblyActionUserSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionUserSecurity
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

	data.UpdateFromBody(ctx, `AssemblyActionUserSecurity`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionUserSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionUserSecurity
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AssemblyActionUserSecurity`))
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

func (r *AssemblyActionUserSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionUserSecurity
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

func (r *AssemblyActionUserSecurityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.AssemblyActionUserSecurity
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

	data.FromBody(ctx, `AssemblyActionUserSecurity`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionUserSecurityResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AssemblyActionUserSecurity

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
