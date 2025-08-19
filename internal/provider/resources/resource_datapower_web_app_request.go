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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &WebAppRequestResource{}

func NewWebAppRequestResource() resource.Resource {
	return &WebAppRequestResource{}
}

type WebAppRequestResource struct {
	pData *tfutils.ProviderData
}

func (r *WebAppRequestResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_app_request"
}

func (r *WebAppRequestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The web request profile specifies various properties about the request side of the transaction that must be satisfied.", "webapp-request-profile", "").String,
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
			"policy_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the satisfaction policy for the profile. The default value is Admission.", "policy-type", "").AddStringEnum("pre-requisite", "admission").AddDefaultValue("admission").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pre-requisite", "admission"),
				},
				Default: stringdefault.StaticString("admission"),
			},
			"ssl_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how the client side of the inspected transaction handles SSL. The default value is Allow.", "request-ssl-policy", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "require", "deny"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"aaa": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this property references a AAA policy, then that AAA policy will be run as a filter on this transaction and the success of that AAA run will be required to continue processing. Any input to this transaction as XML, application/www-url-encoded, or multipart/form-data MIME types will be automatically provided to the AAA processing policy.", "aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the Shared Secret Key for both signing or encrypting.", "ss-key", "crypto_sskey").String,
				Optional:            true,
			},
			"rate_limiter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A rate limiting policy restricts identities (as determined by AAA or the client IP address if AAA is not available) to a specific number of transactions per second.", "ratelimiter-policy", "simple_count_monitor").String,
				Optional:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This Access Control List will be used to allow or deny access to this service based on the IP address of the client. When attached to a service, an Access Control List (ACL) denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.", "acl", "access_control_list").String,
				Optional:            true,
			},
			"ok_methods":  models.GetDmHTTPRequestMethodsResourceSchema("Specify which HTTP methods are acceptable from the client acceptable.", "request-methods", "", false),
			"ok_versions": models.GetDmHTTPVersionMaskResourceSchema("Specify which HTTP versions are acceptable from the client.", "request-versions", "", false),
			"min_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum size of the request body.", "request-body-min", "").String,
				Optional:            true,
			},
			"max_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size of the request body.", "request-body-max", "").AddDefaultValue("128000000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128000000),
			},
			"xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the device handles requests that contain an XML MIME type.", "request-xml-policy", "").AddStringEnum("nothing", "xml", "soap").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "xml", "soap"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This is the transformation rule that is run when the request contains an XML MIME type and the XML processing policy is set to XML or SOAP.", "request-xml-rule", "style_policy_rule").String,
				Optional:            true,
			},
			"non_xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the device handles requests that do not contain an XML MIME type.", "request-nonxml-policy", "").AddStringEnum("nothing", "side", "binary").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "side", "binary"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"non_xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This is the transformation rule that is run when the request does not contain an XML MIME type and the Non-XML processing policy is set to binary or side-effect.", "request-nonxml-rule", "style_policy_rule").String,
				Optional:            true,
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this request policy is violated the firewall error policy will be invoked unless this more specific error policy is provided, in which case this policy takes precedence.", "error-policy-override", "web_app_error_handling_policy").String,
				Optional:            true,
			},
			"session_management_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The session management policy enforces the start pages acceptable for requests that match this security profile. If no policy is referenced, any page is an acceptable start page.", "session-policy", "web_app_session_policy").String,
				Optional:            true,
			},
			"header_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name-value profile allows you to specify what headers are expected, what headers should be stripped, and what headers should be mapped to known values. If no profile is specified, any header is allowed.", "request-header-profile", "name_value_profile").String,
				Optional:            true,
			},
			"url_encoded_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name-value profile allows you to specify what form elements are expected, what form elements should be stripped, and what form elements should be mapped to known values. If no profile is specified, any set of pairs is allowed.", "request-body-profile", "name_value_profile").String,
				Optional:            true,
			},
			"query_string_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how the client URL handles query strings. The default is Allow.", "request-qs-policy", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "require", "deny"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"query_string_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name-value profile for query-string. If not present, no profile is enforced. The profile allows you to validate data members of the query string, filter out unknown ones, or map certain names to known values.", "request-qs-profile", "name_value_profile").String,
				Optional:            true,
			},
			"sql_injection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Data parameters from the query string, application/www-urlencoded requests, and multipart/form-data requests will be passed through the standard SQL Injection filter if this property is enabled.", "request-sql-policy", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_uri_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URI may be no longer than the value specified here.", "request-uri-max", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"uri_filter_unicode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If Unicode is detected in the URI and this property is enabled then the transaction will be rejected", "request-uri-filter-unicode", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_dot_dot": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Filter Requests with .. in the URI after URI normalization", "request-uri-filter-dotdot", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_exe": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Filter Requests with .exe in the URI after URI normalization", "request-uri-filter-exe", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_fragment": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select how to handle URI fragments in requests. A URI fragment is the portion of a URI after the # symbol. The default is Truncate", "request-uri-filter-fragment", "").AddStringEnum("allow", "reject", "truncate").AddDefaultValue("truncate").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "reject", "truncate"),
				},
				Default: stringdefault.StaticString("truncate"),
			},
			"content_types": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A list of PCRE regular expressions that indicate acceptable content-type MIME headers on the request. If this list is empty, any content-type is acceptable. If the request does not have a content type that will be represented as an empty string for matching purposes. Requests without a body (GET, HEAD, and so forth) are not subject to this constraint. An empty list will match all content types.", "request-content-type", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("application/www-url-encoded"),
				})),
			},
			"multipart_form_data": models.GetDmMultipartFormDataProfileResourceSchema("Multipart/Form-Data Profile", "multipart-form-data", "", false),
			"cookie_profile":      models.GetDmCookieProfileResourceSchema("The cookie management profile allows you to specify validation profiles for incoming cookies, whether cookies should be allowed at all, and the signature and encryption policies for cookies.", "cookie-policy", "", false),
			"process_all_cookie": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The process will sign or encrypt all cookies when enabled. The default is enabled.", "process-all-cookie", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cookie_name_vector": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The list of cookies, by name, that the process signs and encrypts.", "cookie-namelist", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"sql_injection_patterns_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The patterns file that the SQL injection filter uses.", "sql-injection-patterns-file", "").AddDefaultValue("store:///SQL-Injection-Patterns.xml").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///SQL-Injection-Patterns.xml"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebAppRequestResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebAppRequestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppRequest
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebAppRequest`)
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

func (r *WebAppRequestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppRequest
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
		data.FromBody(ctx, `WebAppRequest`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WebAppRequest`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppRequestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppRequest
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebAppRequest`))
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

func (r *WebAppRequestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppRequest
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

func (r *WebAppRequestResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebAppRequest

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
