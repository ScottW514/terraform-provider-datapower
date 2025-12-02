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

type ConformancePolicyList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ConformancePolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &ConformancePolicyDataSource{}
)

func NewConformancePolicyDataSource() datasource.DataSource {
	return &ConformancePolicyDataSource{}
}

type ConformancePolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ConformancePolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_conformance_policy"
}

func (d *ConformancePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Conformance Policy",
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
						"profiles": models.GetDmConformanceProfilesDataSourceSchema("Profiles against which to check conformance", "profiles", ""),
						"ignored_requirements": schema.ListAttribute{
							MarkdownDescription: "Requirements that should not be validated. A requirement is specified by a string of the form \"&lt;profile>:&lt;reqid>\", where &lt;profile> names the profile and is one of BP1.0, BP1.1, BSP1.0 or AP1.0, and &lt;reqid> names the requirement within that profile, and follows the naming convention used by the profile itself. For example, requirement R4221 in the Basic Security Profile 1.0 would be named as \"BSP1.0:R4221\".",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"fixup_stylesheets": schema.ListAttribute{
							MarkdownDescription: "Stylesheets to invoke after conformance analysis. These stylesheets can manipulate the analysis results or repair instances of nonconformance.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"assert_bp10_conformance": schema.BoolAttribute{
							MarkdownDescription: "Attach a Basic Profile 1.0 conformance assertion to messages that conform to BP 1.0, or remove a Basic Profile 1.0 conformance assertion to the messages that don't conform to BP 1.0.",
							Computed:            true,
						},
						"report_level": schema.StringAttribute{
							MarkdownDescription: "Select the degree of nonconformance to cause a conformance report to be recorded.",
							Computed:            true,
						},
						"log_target": schema.StringAttribute{
							MarkdownDescription: "Target URL to which conformance reports will be sent",
							Computed:            true,
						},
						"reject_level": schema.StringAttribute{
							MarkdownDescription: "Select the degree of nonconformance to cause the message to be rejected.",
							Computed:            true,
						},
						"reject_include_summary": schema.BoolAttribute{
							MarkdownDescription: "Usually, a rejection response contains little information about the reason that the message was rejected. Setting this property causes the conformance action to include summary information about the conformance errors found.",
							Computed:            true,
						},
						"result_is_conformance_report": schema.BoolAttribute{
							MarkdownDescription: "The normal behavior of the conformance action is to deliver the original message, possibly modified by one or more stylesheets, to the next multistep stage. Setting this property will instead cause the analysis result to be used as the output. This is primarily intended for use within a loopback firewall, which will return the analysis results to the client.",
							Computed:            true,
						},
						"response_properties_enabled": schema.BoolAttribute{
							MarkdownDescription: "When placed inside a single conformance check action (as is typical in an XML gateway), a single set of logging and behavior parameters is sufficent. However, sometimes (as in the case of auto-generated WS-Proxy conformance checking), the same policy is used in checks in both the request and response directions. In this case, the conformance reports should likely be sent to different targets. This toggle allows for an alternate set of logging and rejection parameters to be specified for messages in the response direction.",
							Computed:            true,
						},
						"response_report_level": schema.StringAttribute{
							MarkdownDescription: "Select the degree of nonconformance in a response message to cause a conformance report to be recorded.",
							Computed:            true,
						},
						"response_log_target": schema.StringAttribute{
							MarkdownDescription: "Target URL to which response conformance reports will be sent",
							Computed:            true,
						},
						"response_reject_level": schema.StringAttribute{
							MarkdownDescription: "Select the degree of nonconformance to cause a response message to be rejected.",
							Computed:            true,
						},
						"response_reject_include_summary": schema.BoolAttribute{
							MarkdownDescription: "Usually, a rejection response contains little information about the reason that the message was rejected. Setting this property causes the conformance action to include summary information about the conformance errors found in response messages.",
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

func (d *ConformancePolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ConformancePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ConformancePolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.ConformancePolicy{
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
	l := []models.ConformancePolicy{}
	if resFound {
		if value := res.Get(`ConformancePolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ConformancePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ConformancePolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ConformancePolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
