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

package datasources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type WebAppRequestList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppRequestDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppRequestDataSource{}
)

func NewWebAppRequestDataSource() datasource.DataSource {
	return &WebAppRequestDataSource{}
}

type WebAppRequestDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebAppRequestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_app_request"
}

func (d *WebAppRequestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The web request profile specifies various properties about the request side of the transaction that must be satisfied.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"policy_type": schema.StringAttribute{
							MarkdownDescription: "Select the satisfaction policy for the profile. The default value is Admission.",
							Computed:            true,
						},
						"ssl_policy": schema.StringAttribute{
							MarkdownDescription: "Select how the client side of the inspected transaction handles SSL. The default value is Allow.",
							Computed:            true,
						},
						"aaa": schema.StringAttribute{
							MarkdownDescription: "If this property references a AAA policy, then that AAA policy will be run as a filter on this transaction and the success of that AAA run will be required to continue processing. Any input to this transaction as XML, application/www-url-encoded, or multipart/form-data MIME types will be automatically provided to the AAA processing policy.",
							Computed:            true,
						},
						"sskey": schema.StringAttribute{
							MarkdownDescription: "Select the Shared Secret Key for both signing or encrypting.",
							Computed:            true,
						},
						"rate_limiter": schema.StringAttribute{
							MarkdownDescription: "A rate limiting policy restricts identities (as determined by AAA or the client IP address if AAA is not available) to a specific number of transactions per second.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "This Access Control List will be used to allow or deny access to this service based on the IP address of the client. When attached to a service, an Access Control List (ACL) denies all access by default. To deny access to only selected addresses, first grant access to all addresses (allow 0.0.0.0) and then create deny entries for the desired hosts.",
							Computed:            true,
						},
						"ok_methods":  models.GetDmHTTPRequestMethodsDataSourceSchema("Specify which HTTP methods are acceptable from the client acceptable.", "request-methods", ""),
						"ok_versions": models.GetDmHTTPVersionMaskDataSourceSchema("Specify which HTTP versions are acceptable from the client.", "request-versions", ""),
						"min_body_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the minimum size of the request body.",
							Computed:            true,
						},
						"max_body_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size of the request body.",
							Computed:            true,
						},
						"xml_policy": schema.StringAttribute{
							MarkdownDescription: "Specify how the device handles requests that contain an XML MIME type.",
							Computed:            true,
						},
						"xml_rule": schema.StringAttribute{
							MarkdownDescription: "This is the transformation rule that is run when the request contains an XML MIME type and the XML processing policy is set to XML or SOAP.",
							Computed:            true,
						},
						"non_xml_policy": schema.StringAttribute{
							MarkdownDescription: "Specify how the device handles requests that do not contain an XML MIME type.",
							Computed:            true,
						},
						"non_xml_rule": schema.StringAttribute{
							MarkdownDescription: "This is the transformation rule that is run when the request does not contain an XML MIME type and the Non-XML processing policy is set to binary or side-effect.",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "If this request policy is violated the firewall error policy will be invoked unless this more specific error policy is provided, in which case this policy takes precedence.",
							Computed:            true,
						},
						"session_management_profile": schema.StringAttribute{
							MarkdownDescription: "The session management policy enforces the start pages acceptable for requests that match this security profile. If no policy is referenced, any page is an acceptable start page.",
							Computed:            true,
						},
						"header_gnvc": schema.StringAttribute{
							MarkdownDescription: "The name-value profile allows you to specify what headers are expected, what headers should be stripped, and what headers should be mapped to known values. If no profile is specified, any header is allowed.",
							Computed:            true,
						},
						"url_encoded_gnvc": schema.StringAttribute{
							MarkdownDescription: "The name-value profile allows you to specify what form elements are expected, what form elements should be stripped, and what form elements should be mapped to known values. If no profile is specified, any set of pairs is allowed.",
							Computed:            true,
						},
						"query_string_policy": schema.StringAttribute{
							MarkdownDescription: "Select how the client URL handles query strings. The default is Allow.",
							Computed:            true,
						},
						"query_string_gnvc": schema.StringAttribute{
							MarkdownDescription: "The name-value profile for query-string. If not present, no profile is enforced. The profile allows you to validate data members of the query string, filter out unknown ones, or map certain names to known values.",
							Computed:            true,
						},
						"sql_injection": schema.BoolAttribute{
							MarkdownDescription: "Data parameters from the query string, application/www-urlencoded requests, and multipart/form-data requests will be passed through the standard SQL Injection filter if this property is enabled.",
							Computed:            true,
						},
						"max_uri_size": schema.Int64Attribute{
							MarkdownDescription: "The URI may be no longer than the value specified here.",
							Computed:            true,
						},
						"uri_filter_unicode": schema.BoolAttribute{
							MarkdownDescription: "If Unicode is detected in the URI and this property is enabled then the transaction will be rejected",
							Computed:            true,
						},
						"uri_filter_dot_dot": schema.BoolAttribute{
							MarkdownDescription: "Filter Requests with .. in the URI after URI normalization",
							Computed:            true,
						},
						"uri_filter_exe": schema.BoolAttribute{
							MarkdownDescription: "Filter Requests with .exe in the URI after URI normalization",
							Computed:            true,
						},
						"uri_filter_fragment": schema.StringAttribute{
							MarkdownDescription: "Select how to handle URI fragments in requests. A URI fragment is the portion of a URI after the # symbol. The default is Truncate",
							Computed:            true,
						},
						"content_types": schema.ListAttribute{
							MarkdownDescription: "A list of PCRE regular expressions that indicate acceptable content-type MIME headers on the request. If this list is empty, any content-type is acceptable. If the request does not have a content type that will be represented as an empty string for matching purposes. Requests without a body (GET, HEAD, and so forth) are not subject to this constraint. An empty list will match all content types.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"multipart_form_data": models.GetDmMultipartFormDataProfileDataSourceSchema("Multipart/Form-Data Profile", "multipart-form-data", ""),
						"cookie_profile":      models.GetDmCookieProfileDataSourceSchema("The cookie management profile allows you to specify validation profiles for incoming cookies, whether cookies should be allowed at all, and the signature and encryption policies for cookies.", "cookie-policy", ""),
						"process_all_cookie": schema.BoolAttribute{
							MarkdownDescription: "The process will sign or encrypt all cookies when enabled. The default is enabled.",
							Computed:            true,
						},
						"cookie_name_vector": schema.ListAttribute{
							MarkdownDescription: "The list of cookies, by name, that the process signs and encrypts.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sql_injection_patterns_file": schema.StringAttribute{
							MarkdownDescription: "The patterns file that the SQL injection filter uses.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WebAppRequestDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebAppRequestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebAppRequestList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.WebAppRequest{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.WebAppRequest{}
	if resFound {
		if value := res.Get(`WebAppRequest`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WebAppRequest{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppRequestObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppRequestObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
