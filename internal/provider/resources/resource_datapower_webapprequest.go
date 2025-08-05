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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &WebAppRequestResource{}

func NewWebAppRequestResource() resource.Resource {
	return &WebAppRequestResource{}
}

type WebAppRequestResource struct {
	client *client.DatapowerClient
}

func (r *WebAppRequestResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webapprequest"
}

func (r *WebAppRequestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Web Request Profile", "webapp-request-profile", "").String,

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
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"policy_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Style", "policy-type", "").AddStringEnum("pre-requisite", "admission").AddDefaultValue("admission").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pre-requisite", "admission"),
				},
				Default: stringdefault.StaticString("admission"),
			},
			"ssl_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow SSL", "request-ssl-policy", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "require", "deny"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"aaa": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AAA Policy", "aaa-policy", "aaapolicy").String,
				Optional:            true,
			},
			"ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Shared Secret Key", "ss-key", "cryptosskey").String,
				Optional:            true,
			},
			"rate_limiter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate Limiting", "ratelimiter-policy", "simplecountmonitor").String,
				Optional:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access Control List", "acl", "accesscontrollist").String,
				Optional:            true,
			},
			"ok_methods":  models.GetDmHTTPRequestMethodsResourceSchema("Methods", "request-methods", "", false),
			"ok_versions": models.GetDmHTTPVersionMaskResourceSchema("Versions", "request-versions", "", false),
			"min_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Minimum Size", "request-body-min", "").String,
				Optional:            true,
			},
			"max_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum Size", "request-body-max", "").AddDefaultValue("128000000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128000000),
			},
			"xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Processing", "request-xml-policy", "").AddStringEnum("nothing", "xml", "soap").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "xml", "soap"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Transformation Rule", "request-xml-rule", "stylepolicyrule").String,
				Optional:            true,
			},
			"non_xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Non-XML Processing", "request-nonxml-policy", "").AddStringEnum("nothing", "side", "binary").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "side", "binary"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"non_xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Non-XML Processing Rule", "request-nonxml-rule", "stylepolicyrule").String,
				Optional:            true,
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error Policy", "error-policy-override", "webapperrorhandlingpolicy").String,
				Optional:            true,
			},
			"session_management_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Session Policy", "session-policy", "webappsessionpolicy").String,
				Optional:            true,
			},
			"header_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header Name-Value Profile", "request-header-profile", "namevalueprofile").String,
				Optional:            true,
			},
			"url_encoded_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL-Encoded Body Name-Value Profile", "request-body-profile", "namevalueprofile").String,
				Optional:            true,
			},
			"query_string_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow Query String", "request-qs-policy", "").AddStringEnum("allow", "require", "deny").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "require", "deny"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"query_string_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("QueryString Name-Value Profile", "request-qs-profile", "namevalueprofile").String,
				Optional:            true,
			},
			"sql_injection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL Injection Filter", "request-sql-policy", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_uri_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum URI Length", "request-uri-max", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"uri_filter_unicode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Filter Unicode", "request-uri-filter-unicode", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_dot_dot": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Filter Dot Dot", "request-uri-filter-dotdot", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_exe": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Filter .exe", "request-uri-filter-exe", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"uri_filter_fragment": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Fragmented URI Policy", "request-uri-filter-fragment", "").AddStringEnum("allow", "reject", "truncate").AddDefaultValue("truncate").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "reject", "truncate"),
				},
				Default: stringdefault.StaticString("truncate"),
			},
			"content_types": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Content-Type List", "request-content-type", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("application/www-url-encoded"),
				})),
			},
			"multipart_form_data": models.GetDmMultipartFormDataProfileResourceSchema("Multipart/Form-Data Profile", "multipart-form-data", "", false),
			"cookie_profile":      models.GetDmCookieProfileResourceSchema("Cookie Profile", "cookie-policy", "", false),
			"process_all_cookie": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sign or Encrypt All Cookies", "process-all-cookie", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cookie_name_vector": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cookie Names", "cookie-namelist", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"sql_injection_patterns_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL Injection Patterns File", "sql-injection-patterns-file", "").AddDefaultValue("store:///SQL-Injection-Patterns.xml").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///SQL-Injection-Patterns.xml"),
			},
		},
	}
}

func (r *WebAppRequestResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *WebAppRequestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.WebAppRequest

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebAppRequest`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppRequestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.WebAppRequest

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
		data.FromBody(ctx, `WebAppRequest`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WebAppRequest`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppRequestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.WebAppRequest

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebAppRequest`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppRequestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.WebAppRequest

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}
