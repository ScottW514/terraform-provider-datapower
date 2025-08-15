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

package models

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSPolicyParameters struct {
	PolicyParamParameters         types.String `tfsdk:"policy_param_parameters"`
	PolicyParamWsdlComponentType  types.String `tfsdk:"policy_param_wsdl_component_type"`
	PolicyParamWsdlComponentValue types.String `tfsdk:"policy_param_wsdl_component_value"`
	PolicyParamSubscription       types.String `tfsdk:"policy_param_subscription"`
	PolicyParamFragmentId         types.String `tfsdk:"policy_param_fragment_id"`
}

var DmWSPolicyParametersObjectType = map[string]attr.Type{
	"policy_param_parameters":           types.StringType,
	"policy_param_wsdl_component_type":  types.StringType,
	"policy_param_wsdl_component_value": types.StringType,
	"policy_param_subscription":         types.StringType,
	"policy_param_fragment_id":          types.StringType,
}
var DmWSPolicyParametersObjectDefault = map[string]attr.Value{
	"policy_param_parameters":           types.StringNull(),
	"policy_param_wsdl_component_type":  types.StringValue("all"),
	"policy_param_wsdl_component_value": types.StringNull(),
	"policy_param_subscription":         types.StringNull(),
	"policy_param_fragment_id":          types.StringNull(),
}
var DmWSPolicyParametersDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"policy_param_parameters": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reference to policy parameter object.", "", "policy_parameters").String,
			Computed:            true,
		},
		"policy_param_wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
		},
		"policy_param_wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Computed:            true,
		},
		"policy_param_subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Computed:            true,
		},
		"policy_param_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSPolicyParametersResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"policy_param_parameters": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reference to policy parameter object.", "", "policy_parameters").String,
			Required:            true,
		},
		"policy_param_wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"policy_param_wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Optional:            true,
		},
		"policy_param_subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Optional:            true,
		},
		"policy_param_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSPolicyParameters) IsNull() bool {
	if !data.PolicyParamParameters.IsNull() {
		return false
	}
	if !data.PolicyParamWsdlComponentType.IsNull() {
		return false
	}
	if !data.PolicyParamWsdlComponentValue.IsNull() {
		return false
	}
	if !data.PolicyParamSubscription.IsNull() {
		return false
	}
	if !data.PolicyParamFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSPolicyParameters) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.PolicyParamParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyParamParameters`, data.PolicyParamParameters.ValueString())
	}
	if !data.PolicyParamWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyParamWSDLComponentType`, data.PolicyParamWsdlComponentType.ValueString())
	}
	if !data.PolicyParamWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyParamWSDLComponentValue`, data.PolicyParamWsdlComponentValue.ValueString())
	}
	if !data.PolicyParamSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyParamSubscription`, data.PolicyParamSubscription.ValueString())
	}
	if !data.PolicyParamFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyParamFragmentID`, data.PolicyParamFragmentId.ValueString())
	}
	return body
}

func (data *DmWSPolicyParameters) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PolicyParamParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyParamParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyParamWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamWsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `PolicyParamWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyParamWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyParamSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyParamFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamFragmentId = types.StringNull()
	}
}

func (data *DmWSPolicyParameters) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PolicyParamParameters`); value.Exists() && !data.PolicyParamParameters.IsNull() {
		data.PolicyParamParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamWSDLComponentType`); value.Exists() && !data.PolicyParamWsdlComponentType.IsNull() {
		data.PolicyParamWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.PolicyParamWsdlComponentType.ValueString() != "all" {
		data.PolicyParamWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamWSDLComponentValue`); value.Exists() && !data.PolicyParamWsdlComponentValue.IsNull() {
		data.PolicyParamWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamSubscription`); value.Exists() && !data.PolicyParamSubscription.IsNull() {
		data.PolicyParamSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyParamFragmentID`); value.Exists() && !data.PolicyParamFragmentId.IsNull() {
		data.PolicyParamFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyParamFragmentId = types.StringNull()
	}
}
