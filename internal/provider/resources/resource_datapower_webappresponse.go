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

var _ resource.Resource = &WebAppResponseResource{}

func NewWebAppResponseResource() resource.Resource {
	return &WebAppResponseResource{}
}

type WebAppResponseResource struct {
	pData *tfutils.ProviderData
}

func (r *WebAppResponseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webappresponse"
}

func (r *WebAppResponseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The web response profile specifies various properties about the response side of the transaction that must be satisfied.", "webapp-response-profile", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Select the satisfaction policy for the profile. The default is Admission.", "policy-type", "").AddStringEnum("pre-requisite", "admission").AddDefaultValue("admission").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pre-requisite", "admission"),
				},
				Default: stringdefault.StaticString("admission"),
			},
			"ok_codes":    models.GetDmHTTPResponseCodesResourceSchema("Specify which response codes from the backend server are acceptable.", "response-codes", "", false),
			"ok_versions": models.GetDmHTTPVersionMaskResourceSchema("Specify which HTTP versions are acceptable from the backend server.", "response-versions", "", false),
			"min_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum size of the response body.", "response-body-min", "").String,
				Optional:            true,
			},
			"max_body_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size of the response body.", "response-body-max", "").AddDefaultValue("128000000").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128000000),
			},
			"header_gnvc": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The validation profile allows you to specify what headers are expected, what headers should be stripped, and what headers should be mapped to known values. If no profile is specified, any header is allowed.", "response-header-profile", "namevalueprofile").String,
				Optional:            true,
			},
			"content_types": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A list of PCRE regular expressions that indicate acceptable content-type MIME headers on the response. If this list is empty, any content-type is acceptable. If the response does not have a content type that will be represented as an empty string for matching purposes. An empty list will match all content types.", "response-content-type", "").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue(".*"),
				})),
			},
			"xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the device handles responses that contain an XML MIME type.", "response-xml-policy", "").AddStringEnum("nothing", "xml", "soap").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "xml", "soap"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This is the transformation rule that is run when the response contains an XML MIME type and the XML processing policy is set to XML or SOAP.", "response-xml-rule", "stylepolicyrule").String,
				Optional:            true,
			},
			"non_xml_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how the device handles responses that do not contain an XML MIME type.", "response-nonxml-policy", "").AddStringEnum("nothing", "side", "binary").AddDefaultValue("nothing").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("nothing", "side", "binary"),
				},
				Default: stringdefault.StaticString("nothing"),
			},
			"non_xml_rule": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This is the transformation rule that is run when the response does not contain an XML MIME type and the Non-XML processing policy is set to binary or side-effect.", "response-nonxml-rule", "stylepolicyrule").String,
				Optional:            true,
			},
			"error_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("If this response policy is violated the firewall error policy will be invoked unless this more specific error policy is provided, in which case this policy takes precedence.", "error-policy-override", "webapperrorhandlingpolicy").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebAppResponseResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebAppResponseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppResponse
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebAppResponse`)
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

func (r *WebAppResponseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppResponse
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
		data.FromBody(ctx, `WebAppResponse`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WebAppResponse`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebAppResponseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppResponse
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebAppResponse`))
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

func (r *WebAppResponseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebAppResponse
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

func (r *WebAppResponseResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebAppResponse

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
