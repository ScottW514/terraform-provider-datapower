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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type FormsLoginPolicyList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &FormsLoginPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &FormsLoginPolicyDataSource{}
)

func NewFormsLoginPolicyDataSource() datasource.DataSource {
	return &FormsLoginPolicyDataSource{}
}

type FormsLoginPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *FormsLoginPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_forms_login_policy"
}

func (d *FormsLoginPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Define the policy for HTML form behavior. This policy identifies the location of the HTML login, logout, and error pages.",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"login_form": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment of the login page on the application server. This page collects username and password information. The default value is <tt>/LoginPage.htm</tt> .",
							Computed:            true,
						},
						"use_cookie_attributes": schema.BoolAttribute{
							MarkdownDescription: "Whether to attach a cookie attribute policy to allow predefined or custom attributes to be included in the forms login cookie. If you do not use a cookie attribute policy, the configuration use the properties of the forms login policy. By default, does not attach a cookie attribute policy.",
							Computed:            true,
						},
						"cookie_attributes": schema.StringAttribute{
							MarkdownDescription: "Specify the cookie attribute policy that allows predefined or custom attributes to be included in forms login cookie. <ul><li><tt>Max-Age</tt> overrides inactivity timeout and session lifetime.</li><li><tt>Path</tt> replaces the form POST action URL.</li><li><tt>Secure</tt> replaces the use TLS for login.</li></ul>",
							Computed:            true,
						},
						"use_ssl_for_login": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use TLS for login transactions. By default, uses TLS.",
							Computed:            true,
						},
						"enable_migration": schema.BoolAttribute{
							MarkdownDescription: "Specify whether users need to reauthenticate if their session moves to another member in the standby group if the current member becomes unavailable. Whether another member is available depends on the standby control configuration. In a standby configuration, each member must use the same shared secret. By default, session migration is enabled.",
							Computed:            true,
						},
						"ssl_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the listening port to secure TLS redirection during the login transaction. The default value is 8080.",
							Computed:            true,
						},
						"shared_secret": schema.StringAttribute{
							MarkdownDescription: "Specify the shared secret that protects cookies. The shared secret key protects cookies that the service generates and consumes. If allowed, all members in the standby group must use the same shared secret.",
							Computed:            true,
						},
						"error_page": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment of the error page on the application server. This fragment is relative to the root directory. With the address of the service, the URL fragment forms the complete URL of the error page on the application server. The default value is <tt>/ErrorPage.htm</tt> .",
							Computed:            true,
						},
						"logout_page": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment of the logout page on the application server. This fragment is relative to the root directory. With the address of the service, the URL fragment forms the complete URL of the logout page on the application server. An empty string disables logout. The default value is <tt>/LogoutPage.htm</tt> .",
							Computed:            true,
						},
						"default_url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL fragment on the application server to go to after a successful login. The field is only used when the user was not redirected to the login page. This fragment is relative to the root directory. The URL fragment with the address of the service forms the complete URL of the default page for the application. The default value is <tt>/</tt> .",
							Computed:            true,
						},
						"forms_location": schema.StringAttribute{
							MarkdownDescription: "Specify the location to retrieve HTML pages. These pages can be on the DataPower Gateway or on an application server. The default value is from an application server.",
							Computed:            true,
						},
						"local_login_form": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the login form on the DataPower Gateway. The default value is <tt>store:///LoginPage.htm</tt> .",
							Computed:            true,
						},
						"local_error_page": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the error page on the DataPower Gateway. The default value is <tt>store:///ErrorPage.htm</tt> .",
							Computed:            true,
						},
						"local_logout_page": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the logout page on the DataPower Gateway. The default value is <tt>store:///LogoutPage.htm</tt> .",
							Computed:            true,
						},
						"username_field": schema.StringAttribute{
							MarkdownDescription: "Specify the attribute value of the username field. The default value is <tt>j_username</tt> .",
							Computed:            true,
						},
						"password_field": schema.StringAttribute{
							MarkdownDescription: "Specify the attribute value of the password field. The default value is <tt>j_password</tt> .",
							Computed:            true,
						},
						"redirect_field": schema.StringAttribute{
							MarkdownDescription: "Specify the attribute value for the hidden field that contains the target URL. If an unauthenticated user attempts to access a secured page, the application server displays the login page. The login form contains a hidden field to collect the URL of the secured page that the user attempted to access. After authentication, the application server displays the secured page. The default value is <tt>originalUrl</tt> .",
							Computed:            true,
						},
						"form_processing_url": schema.StringAttribute{
							MarkdownDescription: "Specify the processing URL fragment where the login form posts information. The default value is <tt>/j_security_check</tt> .",
							Computed:            true,
						},
						"inactivity_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds for an inactive session. The default value is 600. A value of 0 disables the timer.",
							Computed:            true,
						},
						"session_lifetime": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum duration in seconds before the user must reauthentication. The default value is 10800. The minimum value is 1. The special value of 0 sets the value to the default value of 10800.",
							Computed:            true,
						},
						"redirect_url_type": schema.StringAttribute{
							MarkdownDescription: "Specify how to retrieve host information for URL redirects. By default, uses the variable value.",
							Computed:            true,
						},
						"form_support_type": schema.StringAttribute{
							MarkdownDescription: "Specify whether to use static HTML pages or a custom file to generate the HTML pages. By default, uses static pages from an explicit location.",
							Computed:            true,
						},
						"form_support_script": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the custom file that generates the HTML pages. This stylesheet or GatewayScript file must generate the following pages. <ul><li>A login form request to generate a login form with the <tt>login</tt> operation.</li><li>A logout page request to generate a logout page with the <tt>logout</tt> operation.</li><li>An error page request to generate an error page with the <tt>error</tt> operation.</li></ul>",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *FormsLoginPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *FormsLoginPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data FormsLoginPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.FormsLoginPolicy{
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
	l := []models.FormsLoginPolicy{}
	if resFound {
		if value := res.Get(`FormsLoginPolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.FormsLoginPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.FormsLoginPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.FormsLoginPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
