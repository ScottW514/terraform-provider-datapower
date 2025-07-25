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

type DmWSOperationPolicySubjectOptOut struct {
	IgnoredSubjects                       *DmPolicySubjectBitmap `tfsdk:"ignored_subjects"`
	PolicySubjectOptOutWsdlComponentType  types.String           `tfsdk:"policy_subject_opt_out_wsdl_component_type"`
	PolicySubjectOptOutWsdlComponentValue types.String           `tfsdk:"policy_subject_opt_out_wsdl_component_value"`
	PolicySubjectOptOutSubscription       types.String           `tfsdk:"policy_subject_opt_out_subscription"`
	PolicySubjectOptOutFragmentId         types.String           `tfsdk:"policy_subject_opt_out_fragment_id"`
}

var DmWSOperationPolicySubjectOptOutObjectType = map[string]attr.Type{
	"ignored_subjects":                            types.ObjectType{AttrTypes: DmPolicySubjectBitmapObjectType},
	"policy_subject_opt_out_wsdl_component_type":  types.StringType,
	"policy_subject_opt_out_wsdl_component_value": types.StringType,
	"policy_subject_opt_out_subscription":         types.StringType,
	"policy_subject_opt_out_fragment_id":          types.StringType,
}
var DmWSOperationPolicySubjectOptOutObjectDefault = map[string]attr.Value{
	"ignored_subjects":                            types.ObjectValueMust(DmPolicySubjectBitmapObjectType, DmPolicySubjectBitmapObjectDefault),
	"policy_subject_opt_out_wsdl_component_type":  types.StringValue("all"),
	"policy_subject_opt_out_wsdl_component_value": types.StringNull(),
	"policy_subject_opt_out_subscription":         types.StringNull(),
	"policy_subject_opt_out_fragment_id":          types.StringNull(),
}
var DmWSOperationPolicySubjectOptOutDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"ignored_subjects": GetDmPolicySubjectBitmapDataSourceSchema("Ignored Subjects", "", ""),
		"policy_subject_opt_out_wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Type", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
		},
		"policy_subject_opt_out_wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Value", "", "").String,
			Computed:            true,
		},
		"policy_subject_opt_out_subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "", "").String,
			Computed:            true,
		},
		"policy_subject_opt_out_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSOperationPolicySubjectOptOutResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"ignored_subjects": GetDmPolicySubjectBitmapResourceSchema("Ignored Subjects", "", "", false),
		"policy_subject_opt_out_wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Type", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"policy_subject_opt_out_wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Value", "", "").String,
			Optional:            true,
		},
		"policy_subject_opt_out_subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "", "").String,
			Optional:            true,
		},
		"policy_subject_opt_out_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSOperationPolicySubjectOptOut) IsNull() bool {
	if data.IgnoredSubjects != nil {
		if !data.IgnoredSubjects.IsNull() {
			return false
		}
	}
	if !data.PolicySubjectOptOutWsdlComponentType.IsNull() {
		return false
	}
	if !data.PolicySubjectOptOutWsdlComponentValue.IsNull() {
		return false
	}
	if !data.PolicySubjectOptOutSubscription.IsNull() {
		return false
	}
	if !data.PolicySubjectOptOutFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSOperationPolicySubjectOptOut) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if data.IgnoredSubjects != nil {
		if !data.IgnoredSubjects.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`IgnoredSubjects`, data.IgnoredSubjects.ToBody(ctx, ""))
		}
	}
	if !data.PolicySubjectOptOutWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicySubjectOptOutWSDLComponentType`, data.PolicySubjectOptOutWsdlComponentType.ValueString())
	}
	if !data.PolicySubjectOptOutWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicySubjectOptOutWSDLComponentValue`, data.PolicySubjectOptOutWsdlComponentValue.ValueString())
	}
	if !data.PolicySubjectOptOutSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicySubjectOptOutSubscription`, data.PolicySubjectOptOutSubscription.ValueString())
	}
	if !data.PolicySubjectOptOutFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicySubjectOptOutFragmentID`, data.PolicySubjectOptOutFragmentId.ValueString())
	}
	return body
}

func (data *DmWSOperationPolicySubjectOptOut) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IgnoredSubjects`); value.Exists() {
		data.IgnoredSubjects = &DmPolicySubjectBitmap{}
		data.IgnoredSubjects.FromBody(ctx, "", value)
	} else {
		data.IgnoredSubjects = nil
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicySubjectOptOutWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutWsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicySubjectOptOutWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicySubjectOptOutSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicySubjectOptOutFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutFragmentId = types.StringNull()
	}
}

func (data *DmWSOperationPolicySubjectOptOut) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IgnoredSubjects`); value.Exists() {
		data.IgnoredSubjects.UpdateFromBody(ctx, "", value)
	} else {
		data.IgnoredSubjects = nil
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutWSDLComponentType`); value.Exists() && !data.PolicySubjectOptOutWsdlComponentType.IsNull() {
		data.PolicySubjectOptOutWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.PolicySubjectOptOutWsdlComponentType.ValueString() != "all" {
		data.PolicySubjectOptOutWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutWSDLComponentValue`); value.Exists() && !data.PolicySubjectOptOutWsdlComponentValue.IsNull() {
		data.PolicySubjectOptOutWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutSubscription`); value.Exists() && !data.PolicySubjectOptOutSubscription.IsNull() {
		data.PolicySubjectOptOutSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicySubjectOptOutFragmentID`); value.Exists() && !data.PolicySubjectOptOutFragmentId.IsNull() {
		data.PolicySubjectOptOutFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicySubjectOptOutFragmentId = types.StringNull()
	}
}
