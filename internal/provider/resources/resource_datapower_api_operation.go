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
)

var _ resource.Resource = &APIOperationResource{}

func NewAPIOperationResource() resource.Resource {
	return &APIOperationResource{}
}

type APIOperationResource struct {
	pData *tfutils.ProviderData
}

func (r *APIOperationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_operation"
}

func (r *APIOperationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An API operation describes the actions to perform against the resource.", "api-operation", "").String,
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
			"method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Method", "method", "").AddStringEnum("GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE").AddDefaultValue("GET").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE"),
				},
				Default: stringdefault.StaticString("GET"),
			},
			"operation_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation ID", "operation-id", "").String,
				Optional:            true,
			},
			"remove_consume": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to remove the API-level consume declaration. By default, the API-level consume declaration is applied to the operation. When removed, the operation can always be performed regardless of the content type.", "remove-consume", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"consume": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify MIME types that the operation can consume. This setting overrides the API-level consume declaration that is defined in the API definition.", "consume", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_]{0,126}\\/[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_+.]{0,126}(;\\s*[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_=\"]{0,126})*$"), "Must match :"+"[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_]{0,126}\\/[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_+.]{0,126}(;\\s*[a-zA-Z0-9][a-zA-Z0-9!#$&\\-^_=\"]{0,126})*$"),
					),
				},
			},
			"produce": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify MIME types that the operation can produce. This setting overrides the API-level produce declaration that is defined in the API definition.", "produce", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(regexp.MustCompile("^(application|audio|image|message|multipart|text|video)\\/[a-zA-Z-]+$"), "Must match :"+"^(application|audio|image|message|multipart|text|video)\\/[a-zA-Z-]+$"),
					),
				},
			},
			"request_schema": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request schema", "request-schema", "api_schema").String,
				Optional:            true,
			},
			"response_schema": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response schema", "response-schema", "").String,
				NestedObject:        models.GetDmAPIResponseSchemaResourceSchema(),
				Optional:            true,
			},
			"parameter": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify applicable parameters for the API operation. This setting overrides the setting in the API path configuration for the same parameter.", "parameter", "").String,
				NestedObject:        models.GetDmAPIParameterResourceSchema(),
				Optional:            true,
			},
			"remove_security": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to remove the API-level security declaration that is defined for the API. By default, the API-level security declaration is applied to the operation. When removed, the operation can be performed without security check.", "remove-security", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"security": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the alternative security requirements to enforce for the operation (that is, there is a logical OR between the security requirements). This setting overrides any declared API-level security.", "security", "api_security_requirement").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"soap_action": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP action", "soap-action", "").String,
				Optional:            true,
			},
			"soap_element_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP element name", "soap-element-name", "").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APIOperationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APIOperationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIOperation
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APIOperation`)
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

func (r *APIOperationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIOperation
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
		data.FromBody(ctx, `APIOperation`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APIOperation`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APIOperationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIOperation
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APIOperation`))
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

func (r *APIOperationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APIOperation
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

func (r *APIOperationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APIOperation

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
