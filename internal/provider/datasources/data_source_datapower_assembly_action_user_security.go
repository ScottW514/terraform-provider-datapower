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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type AssemblyActionUserSecurityList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionUserSecurityDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionUserSecurityDataSource{}
)

func NewAssemblyActionUserSecurityDataSource() datasource.DataSource {
	return &AssemblyActionUserSecurityDataSource{}
}

type AssemblyActionUserSecurityDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionUserSecurityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_user_security"
}

func (d *AssemblyActionUserSecurityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The user security assembly actions extracts identity, authenticates, and authorizes end users.",
		Attributes: map[string]schema.Attribute{
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
						"factor_id": schema.StringAttribute{
							MarkdownDescription: "Factor identifier",
							Computed:            true,
						},
						"extract_identity_method": schema.StringAttribute{
							MarkdownDescription: "Identity extraction method",
							Computed:            true,
						},
						"ei_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to stop processing if identity extraction fails. If failed, stops the assembly and return an error.",
							Computed:            true,
						},
						"user_context_variable": schema.StringAttribute{
							MarkdownDescription: "Username context variable",
							Computed:            true,
						},
						"pass_context_variable": schema.StringAttribute{
							MarkdownDescription: "Password context variable",
							Computed:            true,
						},
						"redirect_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment to redirect the request to obtain user credentials. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"redirect_time_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds for a transaction to complete before the redirect fails. Enter a value in the range 10 - 6000. The default value is 300.",
							Computed:            true,
						},
						"query_parameters": schema.StringAttribute{
							MarkdownDescription: "Query parameters",
							Computed:            true,
						},
						"ei_default_form": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use the default form or a custom form. When enabled, returns the default login page to obtain credentials. When disabled, define the configuration to return the custom login page.",
							Computed:            true,
						},
						"ei_custom_form": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment of the custom login page. This page collects user name and password information. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"ei_custom_form_client_profile": schema.StringAttribute{
							MarkdownDescription: "Custom form TLS client profile",
							Computed:            true,
						},
						"ei_custom_form_content_security_policy": schema.StringAttribute{
							MarkdownDescription: "Specify the value to use for the HTTP <tt>Content-Security-Policy</tt> response header for the custom login page. This response header allows you to control which resources the user agent can load. Generally, you set server origins and script endpoints to detect and mitigate cross-site scripting (XSS), clickjacking, and other injection attacks.",
							Computed:            true,
						},
						"ei_form_time_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds for a transaction to complete before the identity extraction request fails. Enter a value in the range 10 - 600. The default value is 300.",
							Computed:            true,
						},
						"user_auth_method": schema.StringAttribute{
							MarkdownDescription: "Authentication method",
							Computed:            true,
						},
						"au_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to stop processing if authentication fails. If failed, stops the assembly and return an error.",
							Computed:            true,
						},
						"user_registry": schema.StringAttribute{
							MarkdownDescription: "Specify the API user registry to authenticate incoming API requests. The supported registries are API authentication URL and API LDAP.",
							Computed:            true,
						},
						"auth_response_headers_pattern": schema.StringAttribute{
							MarkdownDescription: "Specify the regular expression to select which response headers to add to the API context for access by subsequent actions. The default value is a case-insensitive search on the <tt>x-api</tt> prefix. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"auth_response_credential_header": schema.StringAttribute{
							MarkdownDescription: "Specify the response header that contains the authenticated credentials. The default value is <tt>X-API-Authenticated-Credential</tt> . The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"user_az_method": schema.StringAttribute{
							MarkdownDescription: "Authorization method",
							Computed:            true,
						},
						"az_stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to stop processing if authorization fails. If failed, stops the assembly and return an error.",
							Computed:            true,
						},
						"az_default_form": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use the default form or a custom form. When enabled, returns the default authorization page to obtain authorization. When disabled, define the configuration to return the custom authorization page.",
							Computed:            true,
						},
						"az_custom_form": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment of the custom authorization page. This page obtains permission from the end user. The value can include one or more runtime context variables in the <tt>$(variable)</tt> format.",
							Computed:            true,
						},
						"az_custom_form_client_profile": schema.StringAttribute{
							MarkdownDescription: "Custom form TLS client profile",
							Computed:            true,
						},
						"az_custom_form_content_security_policy": schema.StringAttribute{
							MarkdownDescription: "Specify the value for the HTTP <tt>Content-Security-Policy</tt> response header for the custom authorization page. This response header allows you to control which resources the user agent can load. Generally, you set server origins and script endpoints to detect and mitigate cross-site scripting (XSS), clickjacking, and other injection attacks.",
							Computed:            true,
						},
						"az_form_time_limit": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds for a transaction to complete before the authorization request fails. Enter a value in the range 10 - 600. The default value is 300.",
							Computed:            true,
						},
						"az_table_display_checkboxes": schema.BoolAttribute{
							MarkdownDescription: "Display table check boxes",
							Computed:            true,
						},
						"az_table_dynamic_entries": schema.StringAttribute{
							MarkdownDescription: "Specify the period-delimited context variable that adds dynamic entries to display. This context variable supports space delimited names, a JSON array of names, or a JSON array of objects with name and description.",
							Computed:            true,
						},
						"az_table_default_entry": schema.ListNestedAttribute{
							MarkdownDescription: "Default table entry",
							NestedObject:        models.GetDmTableEntryDataSourceSchema(),
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AssemblyActionUserSecurityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionUserSecurityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionUserSecurityList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionUserSecurity{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.AssemblyActionUserSecurity{}
	if resFound {
		if value := res.Get(`AssemblyActionUserSecurity`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionUserSecurity{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionUserSecurityObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionUserSecurityObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
