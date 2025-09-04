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

var _ resource.Resource = &ConformancePolicyResource{}

func NewConformancePolicyResource() resource.Resource {
	return &ConformancePolicyResource{}
}

type ConformancePolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *ConformancePolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_conformance_policy"
}

func (r *ConformancePolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Conformance Policy", "conformancepolicy", "").String,
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
			"profiles": models.GetDmConformanceProfilesResourceSchema("Profiles against which to check conformance", "profiles", "", false),
			"ignored_requirements": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Requirements that should not be validated. A requirement is specified by a string of the form \"&lt;profile>:&lt;reqid>\", where &lt;profile> names the profile and is one of BP1.0, BP1.1, BSP1.0 or AP1.0, and &lt;reqid> names the requirement within that profile, and follows the naming convention used by the profile itself. For example, requirement R4221 in the Basic Security Profile 1.0 would be named as \"BSP1.0:R4221\".", "ignored-requirements", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"fixup_stylesheets": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheets to invoke after conformance analysis. These stylesheets can manipulate the analysis results or repair instances of nonconformance.", "fixup-stylesheets", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"assert_bp10_conformance": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attach a Basic Profile 1.0 conformance assertion to messages that conform to BP 1.0, or remove a Basic Profile 1.0 conformance assertion to the messages that don't conform to BP 1.0.", "assert-bp10-conformance", "").AddDefaultValue("false").AddNotValidWhen(models.ConformancePolicyAssertBP10ConformanceIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"report_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the degree of nonconformance to cause a conformance report to be recorded.", "report-level", "").AddStringEnum("never", "failure", "warning", "always").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "failure", "warning", "always"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"log_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target URL to which conformance reports will be sent", "report-target", "").AddRequiredWhen(models.ConformancePolicyLogTargetCondVal.String()).AddNotValidWhen(models.ConformancePolicyLogTargetIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.ConformancePolicyLogTargetCondVal, models.ConformancePolicyLogTargetIgnoreVal, false),
				},
			},
			"reject_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the degree of nonconformance to cause the message to be rejected.", "reject-level", "").AddStringEnum("never", "failure", "warning").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "failure", "warning"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"reject_include_summary": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Usually, a rejection response contains little information about the reason that the message was rejected. Setting this property causes the conformance action to include summary information about the conformance errors found.", "reject-include-summary", "").AddDefaultValue("false").AddRequiredWhen(models.ConformancePolicyRejectIncludeSummaryCondVal.String()).AddNotValidWhen(models.ConformancePolicyRejectIncludeSummaryIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"result_is_conformance_report": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The normal behavior of the conformance action is to deliver the original message, possibly modified by one or more stylesheets, to the next multistep stage. Setting this property will instead cause the analysis result to be used as the output. This is primarily intended for use within a loopback firewall, which will return the analysis results to the client.", "result-is-conformance-report", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"response_properties_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When placed inside a single conformance check action (as is typical in an XML gateway), a single set of logging and behavior parameters is sufficent. However, sometimes (as in the case of auto-generated WS-Proxy conformance checking), the same policy is used in checks in both the request and response directions. In this case, the conformance reports should likely be sent to different targets. This toggle allows for an alternate set of logging and rejection parameters to be specified for messages in the response direction.", "response-properties-enabled", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"response_report_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the degree of nonconformance in a response message to cause a conformance report to be recorded.", "response-report-level", "").AddStringEnum("never", "failure", "warning", "always").AddDefaultValue("never").AddRequiredWhen(models.ConformancePolicyResponseReportLevelCondVal.String()).AddNotValidWhen(models.ConformancePolicyResponseReportLevelIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "failure", "warning", "always"),
					validators.ConditionalRequiredString(models.ConformancePolicyResponseReportLevelCondVal, models.ConformancePolicyResponseReportLevelIgnoreVal, true),
				},
				Default: stringdefault.StaticString("never"),
			},
			"response_log_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target URL to which response conformance reports will be sent", "response-report-target", "").AddRequiredWhen(models.ConformancePolicyResponseLogTargetCondVal.String()).AddNotValidWhen(models.ConformancePolicyResponseLogTargetIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.ConformancePolicyResponseLogTargetCondVal, models.ConformancePolicyResponseLogTargetIgnoreVal, false),
				},
			},
			"response_reject_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the degree of nonconformance to cause a response message to be rejected.", "response-reject-level", "").AddStringEnum("never", "failure", "warning").AddDefaultValue("never").AddRequiredWhen(models.ConformancePolicyResponseRejectLevelCondVal.String()).AddNotValidWhen(models.ConformancePolicyResponseRejectLevelIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "failure", "warning"),
					validators.ConditionalRequiredString(models.ConformancePolicyResponseRejectLevelCondVal, models.ConformancePolicyResponseRejectLevelIgnoreVal, true),
				},
				Default: stringdefault.StaticString("never"),
			},
			"response_reject_include_summary": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Usually, a rejection response contains little information about the reason that the message was rejected. Setting this property causes the conformance action to include summary information about the conformance errors found in response messages.", "response-reject-include-summary", "").AddDefaultValue("false").AddRequiredWhen(models.ConformancePolicyResponseRejectIncludeSummaryCondVal.String()).AddNotValidWhen(models.ConformancePolicyResponseRejectIncludeSummaryIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ConformancePolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ConformancePolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConformancePolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `ConformancePolicy`)
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

func (r *ConformancePolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConformancePolicy
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
		data.FromBody(ctx, `ConformancePolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `ConformancePolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConformancePolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConformancePolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `ConformancePolicy`))
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

func (r *ConformancePolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConformancePolicy
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

func (r *ConformancePolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.ConformancePolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
