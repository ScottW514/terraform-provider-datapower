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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmPolicyAttachmentPoint struct {
	PolicyAttachWsdlComponentType  types.String `tfsdk:"policy_attach_wsdl_component_type"`
	PolicyAttachWsdlComponentValue types.String `tfsdk:"policy_attach_wsdl_component_value"`
	PolicyAttachFragmentId         types.String `tfsdk:"policy_attach_fragment_id"`
}

var DmPolicyAttachmentPointObjectType = map[string]attr.Type{
	"policy_attach_wsdl_component_type":  types.StringType,
	"policy_attach_wsdl_component_value": types.StringType,
	"policy_attach_fragment_id":          types.StringType,
}
var DmPolicyAttachmentPointObjectDefault = map[string]attr.Value{
	"policy_attach_wsdl_component_type":  types.StringNull(),
	"policy_attach_wsdl_component_value": types.StringNull(),
	"policy_attach_fragment_id":          types.StringNull(),
}
var DmPolicyAttachmentPointDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"policy_attach_wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Type", "", "").AddStringEnum("service", "port", "fragmentid", "rest").String,
			Computed:            true,
		},
		"policy_attach_wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Value", "", "").String,
			Computed:            true,
		},
		"policy_attach_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmPolicyAttachmentPointResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"policy_attach_wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Type", "", "").AddStringEnum("service", "port", "fragmentid", "rest").String,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("service", "port", "fragmentid", "rest"),
			},
		},
		"policy_attach_wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Component Value", "", "").String,
			Optional:            true,
		},
		"policy_attach_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmPolicyAttachmentPoint) IsNull() bool {
	if !data.PolicyAttachWsdlComponentType.IsNull() {
		return false
	}
	if !data.PolicyAttachWsdlComponentValue.IsNull() {
		return false
	}
	if !data.PolicyAttachFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmPolicyAttachmentPoint) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.PolicyAttachWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyAttachWSDLComponentType`, data.PolicyAttachWsdlComponentType.ValueString())
	}
	if !data.PolicyAttachWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyAttachWSDLComponentValue`, data.PolicyAttachWsdlComponentValue.ValueString())
	}
	if !data.PolicyAttachFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyAttachFragmentID`, data.PolicyAttachFragmentId.ValueString())
	}
	return body
}

func (data *DmPolicyAttachmentPoint) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PolicyAttachWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyAttachWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyAttachWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyAttachFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachFragmentId = types.StringNull()
	}
}

func (data *DmPolicyAttachmentPoint) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `PolicyAttachWSDLComponentType`); value.Exists() && !data.PolicyAttachWsdlComponentType.IsNull() {
		data.PolicyAttachWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachWSDLComponentValue`); value.Exists() && !data.PolicyAttachWsdlComponentValue.IsNull() {
		data.PolicyAttachWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachFragmentID`); value.Exists() && !data.PolicyAttachFragmentId.IsNull() {
		data.PolicyAttachFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachFragmentId = types.StringNull()
	}
}
