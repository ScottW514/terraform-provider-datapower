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

var _ resource.Resource = &FormsLoginPolicyResource{}

func NewFormsLoginPolicyResource() resource.Resource {
	return &FormsLoginPolicyResource{}
}

type FormsLoginPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *FormsLoginPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_forms_login_policy"
}

func (r *FormsLoginPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Define the policy for HTML form behavior. This policy identifies the location of the HTML login, logout, and error pages.", "forms-login-policy", "").String,
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
			"login_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment of the login page on the application server. This page collects username and password information. The default value is <tt>/LoginPage.htm</tt> .", "login-url", "").AddDefaultValue("/LoginPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/LoginPage.htm"),
			},
			"use_cookie_attributes": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Whether to attach a cookie attribute policy to allow predefined or custom attributes to be included in the forms login cookie. If you do not use a cookie attribute policy, the configuration use the properties of the forms login policy. By default, does not attach a cookie attribute policy.", "use-cookie-attribute", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cookie_attributes": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie attribute policy that allows predefined or custom attributes to be included in forms login cookie. <ul><li><tt>Max-Age</tt> overrides inactivity timeout and session lifetime.</li><li><tt>Path</tt> replaces the form POST action URL.</li><li><tt>Secure</tt> replaces the use TLS for login.</li></ul>", "cookie-attributes", "cookie_attribute_policy").String,
				Optional:            true,
			},
			"use_ssl_for_login": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use TLS for login transactions. By default, uses TLS.", "use-ssl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_migration": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether users need to reauthenticate if their session moves to another member in the standby group if the current member becomes unavailable. Whether another member is available depends on the standby control configuration. In a standby configuration, each member must use the same shared secret. By default, session migration is enabled.", "enable-migration", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port to secure TLS redirection during the login transaction. The default value is 8080.", "ssl-port", "").AddIntegerRange(1, 65535).AddDefaultValue("8080").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(8080),
			},
			"shared_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret that protects cookies. The shared secret key protects cookies that the service generates and consumes. If allowed, all members in the standby group must use the same shared secret.", "shared-secret", "crypto_sskey").String,
				Optional:            true,
			},
			"error_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment of the error page on the application server. This fragment is relative to the root directory. With the address of the service, the URL fragment forms the complete URL of the error page on the application server. The default value is <tt>/ErrorPage.htm</tt> .", "error-url", "").AddDefaultValue("/ErrorPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/ErrorPage.htm"),
			},
			"logout_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment of the logout page on the application server. This fragment is relative to the root directory. With the address of the service, the URL fragment forms the complete URL of the logout page on the application server. An empty string disables logout. The default value is <tt>/LogoutPage.htm</tt> .", "logout-url", "").AddDefaultValue("/LogoutPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/LogoutPage.htm"),
			},
			"default_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL fragment on the application server to go to after a successful login. The field is only used when the user was not redirected to the login page. This fragment is relative to the root directory. The URL fragment with the address of the service forms the complete URL of the default page for the application. The default value is <tt>/</tt> .", "default-url", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"forms_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location to retrieve HTML pages. These pages can be on the DataPower Gateway or on an application server. The default value is from an application server.", "forms-location", "").AddStringEnum("backend", "appliance").AddDefaultValue("backend").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("backend", "appliance"),
				},
				Default: stringdefault.StaticString("backend"),
			},
			"local_login_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the login form on the DataPower Gateway. The default value is <tt>store:///LoginPage.htm</tt> .", "local-login-page", "").AddDefaultValue("store:///LoginPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///LoginPage.htm"),
			},
			"local_error_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the error page on the DataPower Gateway. The default value is <tt>store:///ErrorPage.htm</tt> .", "local-error-page", "").AddDefaultValue("store:///ErrorPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///ErrorPage.htm"),
			},
			"local_logout_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the logout page on the DataPower Gateway. The default value is <tt>store:///LogoutPage.htm</tt> .", "local-logout-page", "").AddDefaultValue("store:///LogoutPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///LogoutPage.htm"),
			},
			"username_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute value of the username field. The default value is <tt>j_username</tt> .", "username-field", "").AddDefaultValue("j_username").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("j_username"),
			},
			"password_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute value of the password field. The default value is <tt>j_password</tt> .", "password-field", "").AddDefaultValue("j_password").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("j_password"),
			},
			"redirect_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute value for the hidden field that contains the target URL. If an unauthenticated user attempts to access a secured page, the application server displays the login page. The login form contains a hidden field to collect the URL of the secured page that the user attempted to access. After authentication, the application server displays the secured page. The default value is <tt>originalUrl</tt> .", "redirect-field", "").AddDefaultValue("originalUrl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("originalUrl"),
			},
			"form_processing_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the processing URL fragment where the login form posts information. The default value is <tt>/j_security_check</tt> .", "post-action-url", "").AddDefaultValue("/j_security_check").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/j_security_check"),
			},
			"inactivity_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds for an inactive session. The default value is 600. A value of 0 disables the timer.", "inactivity-timeout", "").AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(600),
			},
			"session_lifetime": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum duration in seconds before the user must reauthentication. The default value is 10800. The minimum value is 1. The special value of 0 sets the value to the default value of 10800.", "session-lifetime", "").AddDefaultValue("10800").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10800),
			},
			"redirect_url_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to retrieve host information for URL redirects. By default, uses the variable value.", "redirect-url-type", "").AddStringEnum("urlin", "host").AddDefaultValue("urlin").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("urlin", "host"),
				},
				Default: stringdefault.StaticString("urlin"),
			},
			"form_support_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use static HTML pages or a custom file to generate the HTML pages. By default, uses static pages from an explicit location.", "support-type", "").AddStringEnum("static", "custom").AddDefaultValue("static").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static", "custom"),
				},
				Default: stringdefault.StaticString("static"),
			},
			"form_support_script": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the custom file that generates the HTML pages. This stylesheet or GatewayScript file must generate the following pages. <ul><li>A login form request to generate a login form with the <tt>login</tt> operation.</li><li>A logout page request to generate a logout page with the <tt>logout</tt> operation.</li><li>An error page request to generate an error page with the <tt>error</tt> operation.</li></ul>", "script-location", "").AddDefaultValue("store:///Form-Generate-HTML.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///Form-Generate-HTML.xsl"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *FormsLoginPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *FormsLoginPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.FormsLoginPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `FormsLoginPolicy`)
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

func (r *FormsLoginPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.FormsLoginPolicy
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
		data.FromBody(ctx, `FormsLoginPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `FormsLoginPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FormsLoginPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.FormsLoginPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `FormsLoginPolicy`))
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

func (r *FormsLoginPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.FormsLoginPolicy
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

func (r *FormsLoginPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.FormsLoginPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
