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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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

var _ resource.Resource = &OAuthSupportedClientGroupResource{}
var _ resource.ResourceWithImportState = &OAuthSupportedClientGroupResource{}

func NewOAuthSupportedClientGroupResource() resource.Resource {
	return &OAuthSupportedClientGroupResource{}
}

type OAuthSupportedClientGroupResource struct {
	pData *tfutils.ProviderData
}

func (r *OAuthSupportedClientGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_supported_client_group"
}

func (r *OAuthSupportedClientGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>To support the OAuth 2.0 protocol, an AAA policy requires the configuration of an OAuth client group. An OAuth client group contains the configured OAuth clients that the DataPower Gateway accepts requests from.</p><p>When creating an OAuth client group for an AAA policy, the OAuth client group consists of one or more OAuth clients with the same OAuth roles.</p>", "oauth-supported-client-group", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies a brief comment that describes the configuration.", "summary", "").String,
				Optional:            true,
			},
			"customized": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the configuration is for a customized OAuth client group.", "customized", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"customized_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sets the method to customize an OAuth client.", "customized-type", "").AddStringEnum("custom", "template").AddDefaultValue("custom").AddNotValidWhen(models.OAuthSupportedClientGroupCustomizedTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("custom", "template"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.OAuthSupportedClientGroupCustomizedTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("custom"),
			},
			"oauth_role": models.GetDmOAuthRoleResourceSchema("Identifies the roles of clients in the group. This property is mutually exclusive to the <b>Customized OAuth</b> property.", "oauth-role", "", false),
			"client": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Manages the group of OAuth clients. Use the controls to add or remove clients from the group.", "client", "oauth_supported_client").AddRequiredWhen(models.OAuthSupportedClientGroupClientCondVal.String()).AddNotValidWhen(models.OAuthSupportedClientGroupClientIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(),
					validators.ConditionalRequiredList(models.OAuthSupportedClientGroupClientCondVal, models.OAuthSupportedClientGroupClientIgnoreVal, false),
				},
			},
			"template_process_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the location of the stylesheet or GatewayScript file that defines the OAuth client. You can define parts of the configuration parameters in the stylesheet or GatewayScript file and then specify an OAuth client template to derive the remaining information. Note that the stylesheet or GatewayScript file must at least provide the \"client-id\" node. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>The stylesheet or GatewayScript file must follow the guidelines when it returns the information: <ul><li>The stylesheet or GatewayScript file must return the &lt;client-id> element.</li><li>If the &lt;customized> element value is set to \"on\", the &lt;customized-process-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;customized> element value is set to \"on\" in the client template, the &lt;customized-process-url> element cannot unset this value.</li><li>If the &lt;use-validation-url> element value is set to \"on\", the &lt;validation-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;custom-scope-check> element value is set to \"on\", the &lt;scope-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;custom-resource-owner> element value is set to \"on\", the &lt;resource-owner-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;role> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;client-type> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;az-grant> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;az-grant> element value is set to \"+code+\" or \"+token+\", the &lt;local-az-page-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;caching> element value is set to \"custom\", the &lt;additional-oauth-processing-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;refresh-token-allowed> is set to a non-zero value, the &lt;refresh-token-lifetime> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;check-client-credential> element value is set to \"on\", the &lt;client-authen-method> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;client-authen-method> element value is set to \"secret\", the &lt;client-secret> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;client-authen-method> element value is set to \"ssl\", the &lt;client-valcred> element value must be provided by either the stylesheet or GatewayScript file or the template.</li></ul></p>", "template-process-url", "").AddRequiredWhen(models.OAuthSupportedClientGroupTemplateProcessUrlCondVal.String()).AddNotValidWhen(models.OAuthSupportedClientGroupTemplateProcessUrlIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.OAuthSupportedClientGroupTemplateProcessUrlCondVal, models.OAuthSupportedClientGroupTemplateProcessUrlIgnoreVal, false),
				},
			},
			"client_template": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the OAuth client template that is used to derive the configuration parameters that are not specified in the stylesheet or GatewayScript file of the OAuth client.</p>", "client-template", "oauth_supported_client").AddRequiredWhen(models.OAuthSupportedClientGroupClientTemplateCondVal.String()).AddNotValidWhen(models.OAuthSupportedClientGroupClientTemplateIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.OAuthSupportedClientGroupClientTemplateCondVal, models.OAuthSupportedClientGroupClientTemplateIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *OAuthSupportedClientGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *OAuthSupportedClientGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClientGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `OAuthSupportedClientGroup`)
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

func (r *OAuthSupportedClientGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClientGroup
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

	data.UpdateFromBody(ctx, `OAuthSupportedClientGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClientGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `OAuthSupportedClientGroup`))
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

func (r *OAuthSupportedClientGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.OAuthSupportedClientGroup
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

func (r *OAuthSupportedClientGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.OAuthSupportedClientGroup
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

	data.FromBody(ctx, `OAuthSupportedClientGroup`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OAuthSupportedClientGroupResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.OAuthSupportedClientGroup

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
