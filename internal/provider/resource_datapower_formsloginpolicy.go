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

package provider

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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &FormsLoginPolicyResource{}

func NewFormsLoginPolicyResource() resource.Resource {
	return &FormsLoginPolicyResource{}
}

type FormsLoginPolicyResource struct {
	client *client.DatapowerClient
}

func (r *FormsLoginPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_formsloginpolicy"
}

func (r *FormsLoginPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("HTML forms login policy", "forms-login-policy", "").String,

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
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"login_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Login", "login-url", "").AddDefaultValue("/LoginPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/LoginPage.htm"),
			},
			"use_cookie_attributes": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attach cookie attribute policy", "use-cookie-attribute", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"cookie_attributes": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cookie attributes", "cookie-attributes", "cookieattributepolicy").String,
				Optional:            true,
			},
			"use_ssl_for_login": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use TLS for Login", "use-ssl", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"enable_migration": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable session migration", "enable-migration", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ssl_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS port", "ssl-port", "").AddIntegerRange(1, 65535).AddDefaultValue("8080").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(8080),
			},
			"shared_secret": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Shared secret", "shared-secret", "cryptosskey").String,
				Optional:            true,
			},
			"error_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Error", "error-url", "").AddDefaultValue("/ErrorPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/ErrorPage.htm"),
			},
			"logout_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Logout", "logout-url", "").AddDefaultValue("/LogoutPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/LogoutPage.htm"),
			},
			"default_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default", "default-url", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"forms_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Forms storage location", "forms-location", "").AddStringEnum("backend", "appliance").AddDefaultValue("backend").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("backend", "appliance"),
				},
				Default: stringdefault.StaticString("backend"),
			},
			"local_login_form": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local login form", "local-login-page", "").AddDefaultValue("store:///LoginPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///LoginPage.htm"),
			},
			"local_error_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local error page", "local-error-page", "").AddDefaultValue("store:///ErrorPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///ErrorPage.htm"),
			},
			"local_logout_page": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local logout page", "local-logout-page", "").AddDefaultValue("store:///LogoutPage.htm").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///LogoutPage.htm"),
			},
			"username_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username field name", "username-field", "").AddDefaultValue("j_username").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("j_username"),
			},
			"password_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password field name", "password-field", "").AddDefaultValue("j_password").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("j_password"),
			},
			"redirect_field": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target URL field name", "redirect-field", "").AddDefaultValue("originalUrl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("originalUrl"),
			},
			"form_processing_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Form POST action URL", "post-action-url", "").AddDefaultValue("/j_security_check").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/j_security_check"),
			},
			"inactivity_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Inactivity timeout", "inactivity-timeout", "").AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(600),
			},
			"session_lifetime": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Session lifetime", "session-lifetime", "").AddDefaultValue("10800").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10800),
			},
			"redirect_url_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Redirect URL Type", "redirect-url-type", "").AddStringEnum("urlin", "host").AddDefaultValue("urlin").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("urlin", "host"),
				},
				Default: stringdefault.StaticString("urlin"),
			},
			"form_support_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Source for form processing", "support-type", "").AddStringEnum("static", "custom").AddDefaultValue("static").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("static", "custom"),
				},
				Default: stringdefault.StaticString("static"),
			},
			"form_support_script": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom processing for form", "script-location", "").AddDefaultValue("store:///Form-Generate-HTML.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///Form-Generate-HTML.xsl"),
			},
		},
	}
}

func (r *FormsLoginPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *FormsLoginPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.FormsLoginPolicy

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `FormsLoginPolicy`)
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

func (r *FormsLoginPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.FormsLoginPolicy

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
		data.FromBody(ctx, `FormsLoginPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `FormsLoginPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FormsLoginPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.FormsLoginPolicy

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `FormsLoginPolicy`))
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

func (r *FormsLoginPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.FormsLoginPolicy

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
