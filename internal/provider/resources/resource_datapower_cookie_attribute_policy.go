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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &CookieAttributePolicyResource{}

func NewCookieAttributePolicyResource() resource.Resource {
	return &CookieAttributePolicyResource{}
}

type CookieAttributePolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *CookieAttributePolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cookie_attribute_policy"
}

func (r *CookieAttributePolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Cookie attribute policy manages pre-defined and custom attributes of a cookie.", "cookie-attribute-policy", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("A descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"cookie_attribute": models.GetDmCookieAttributeResourceSchema("Specifies the attributes to include in the cookie. Acceptable values include domain, path, secure, httponly, max-age, expires, and custom. When an attribute is enabled, you can enter its value in the corresponding property and the enabled attribute is included in the cookie.", "cookie-attribute", "", false),
			"domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies domain to which a cookie belongs. A browser accepts cookies only when the current domain matches the value you enter here. The maximum length of the domain is 256 characters.", "domain", "").AddRequiredWhen(models.CookieAttributePolicyDomainCondVal.String()).AddNotValidWhen(models.CookieAttributePolicyDomainIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 256),
					validators.ConditionalRequiredString(models.CookieAttributePolicyDomainCondVal, models.CookieAttributePolicyDomainIgnoreVal, false),
				},
			},
			"path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies path attribute of a cookie. A browser accepts cookies only when the current path matches the value you enter there. If this policy object is attached to HTML Forms Login Policy, this property overrides Form POST Action URL property. The maximum length of the path is 256 characters.", "path", "").AddDefaultValue("/").AddRequiredWhen(models.CookieAttributePolicyPathCondVal.String()).AddNotValidWhen(models.CookieAttributePolicyPathIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 256),
					validators.ConditionalRequiredString(models.CookieAttributePolicyPathCondVal, models.CookieAttributePolicyPathIgnoreVal, true),
				},
				Default: stringdefault.StaticString("/"),
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Sets the cookie's maximum age and/or the cookie's expiration date as an interval of seconds, relative to the time the transaction occurred on the object. For example, if this value is set to 3600 and the transaction on this object occurred on Feb 10, 2014 12:00:00 GMT, then the maximum age of the cookie is 3600 seconds and the expiration date is Feb 10, 2014 13:00:00 GMT, depending on whether the Max-Age and the Expires attribute are included.</p><p>When the maximum age or the expiration date is reached, the cookie is deleted. Enter a value in the range 1 - 2678400. The default value is 3600. Note that the Max-Age attribute in this policy overrides Inactivity Timeout and Session Lifetime attributes in HTML Forms Login policy.</p>", "interval", "").AddIntegerRange(1, 2678400).AddDefaultValue("3600").AddRequiredWhen(models.CookieAttributePolicyIntervalCondVal.String()).AddNotValidWhen(models.CookieAttributePolicyIntervalIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2678400),
					validators.ConditionalRequiredInt64(models.CookieAttributePolicyIntervalCondVal, models.CookieAttributePolicyIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(3600),
			},
			"custom_attribute": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The additional attributes to include in the cookie. Enter each attribute in name-value pair. When you enter multiple pairs, use a semicolon (;) to separate them. A name-value pair with an empty value (name-only portion) can also be specified here. You can use variables instead of name-value pair(s). Enter a context variable as var://variablename", "custom-attribute", "").AddRequiredWhen(models.CookieAttributePolicyCustomAttributeCondVal.String()).AddNotValidWhen(models.CookieAttributePolicyCustomAttributeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.CookieAttributePolicyCustomAttributeCondVal, models.CookieAttributePolicyCustomAttributeIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *CookieAttributePolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *CookieAttributePolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CookieAttributePolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `CookieAttributePolicy`)
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

func (r *CookieAttributePolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CookieAttributePolicy
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
		data.FromBody(ctx, `CookieAttributePolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `CookieAttributePolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CookieAttributePolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CookieAttributePolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `CookieAttributePolicy`))
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

func (r *CookieAttributePolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CookieAttributePolicy
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

func (r *CookieAttributePolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CookieAttributePolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
