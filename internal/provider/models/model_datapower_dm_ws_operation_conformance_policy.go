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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSOperationConformancePolicy struct {
	ConformancePolicy                   types.String `tfsdk:"conformance_policy"`
	ConformancePolicyWsdlComponentType  types.String `tfsdk:"conformance_policy_wsdl_component_type"`
	ConformancePolicyWsdlComponentValue types.String `tfsdk:"conformance_policy_wsdl_component_value"`
	ConformancePolicySubscription       types.String `tfsdk:"conformance_policy_subscription"`
	ConformancePolicyFragmentId         types.String `tfsdk:"conformance_policy_fragment_id"`
}

var DmWSOperationConformancePolicyConformancePolicyWSDLComponentValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "conformance_policy_wsdl_component_type",
	AttrType:    "String",
	AttrDefault: "all",
	Value:       []string{"subscription"},
}

var DmWSOperationConformancePolicyConformancePolicySubscriptionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "conformance_policy_wsdl_component_type",
	AttrType:    "String",
	AttrDefault: "all",
	Value:       []string{"subscription"},
}

var DmWSOperationConformancePolicyConformancePolicySubscriptionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmWSOperationConformancePolicyObjectType = map[string]attr.Type{
	"conformance_policy":                      types.StringType,
	"conformance_policy_wsdl_component_type":  types.StringType,
	"conformance_policy_wsdl_component_value": types.StringType,
	"conformance_policy_subscription":         types.StringType,
	"conformance_policy_fragment_id":          types.StringType,
}
var DmWSOperationConformancePolicyObjectDefault = map[string]attr.Value{
	"conformance_policy":                      types.StringNull(),
	"conformance_policy_wsdl_component_type":  types.StringValue("all"),
	"conformance_policy_wsdl_component_value": types.StringNull(),
	"conformance_policy_subscription":         types.StringNull(),
	"conformance_policy_fragment_id":          types.StringNull(),
}

func GetDmWSOperationConformancePolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSOperationConformancePolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"conformance_policy": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Conformance Policy", "", "conformance_policy").String,
				Computed:            true,
			},
			"conformance_policy_wsdl_component_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
				Computed:            true,
			},
			"conformance_policy_wsdl_component_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").AddNotValidWhen(DmWSOperationConformancePolicyConformancePolicyWSDLComponentValueIgnoreVal.String()).String,
				Computed:            true,
			},
			"conformance_policy_subscription": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").AddRequiredWhen(DmWSOperationConformancePolicyConformancePolicySubscriptionCondVal.String()).AddNotValidWhen(DmWSOperationConformancePolicyConformancePolicySubscriptionIgnoreVal.String()).String,
				Computed:            true,
			},
			"conformance_policy_fragment_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmWSOperationConformancePolicyDataSourceSchema
}
func GetDmWSOperationConformancePolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSOperationConformancePolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"conformance_policy": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Conformance Policy", "", "conformance_policy").String,
				Required:            true,
			},
			"conformance_policy_wsdl_component_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
				},
				Default: stringdefault.StaticString("all"),
			},
			"conformance_policy_wsdl_component_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").AddNotValidWhen(DmWSOperationConformancePolicyConformancePolicyWSDLComponentValueIgnoreVal.String()).String,
				Optional:            true,
			},
			"conformance_policy_subscription": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").AddRequiredWhen(DmWSOperationConformancePolicyConformancePolicySubscriptionCondVal.String()).AddNotValidWhen(DmWSOperationConformancePolicyConformancePolicySubscriptionIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmWSOperationConformancePolicyConformancePolicySubscriptionCondVal, DmWSOperationConformancePolicyConformancePolicySubscriptionIgnoreVal, false),
				},
			},
			"conformance_policy_fragment_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmWSOperationConformancePolicyResourceSchema
}

func (data DmWSOperationConformancePolicy) IsNull() bool {
	if !data.ConformancePolicy.IsNull() {
		return false
	}
	if !data.ConformancePolicyWsdlComponentType.IsNull() {
		return false
	}
	if !data.ConformancePolicyWsdlComponentValue.IsNull() {
		return false
	}
	if !data.ConformancePolicySubscription.IsNull() {
		return false
	}
	if !data.ConformancePolicyFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSOperationConformancePolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ConformancePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConformancePolicy`, data.ConformancePolicy.ValueString())
	}
	if !data.ConformancePolicyWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConformancePolicyWSDLComponentType`, data.ConformancePolicyWsdlComponentType.ValueString())
	}
	if !data.ConformancePolicyWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConformancePolicyWSDLComponentValue`, data.ConformancePolicyWsdlComponentValue.ValueString())
	}
	if !data.ConformancePolicySubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConformancePolicySubscription`, data.ConformancePolicySubscription.ValueString())
	}
	if !data.ConformancePolicyFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConformancePolicyFragmentID`, data.ConformancePolicyFragmentId.ValueString())
	}
	return body
}

func (data *DmWSOperationConformancePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ConformancePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConformancePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicyWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConformancePolicyWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicyWsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `ConformancePolicyWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConformancePolicyWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicyWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicySubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConformancePolicySubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicySubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicyFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConformancePolicyFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicyFragmentId = types.StringNull()
	}
}

func (data *DmWSOperationConformancePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ConformancePolicy`); value.Exists() && !data.ConformancePolicy.IsNull() {
		data.ConformancePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicyWSDLComponentType`); value.Exists() && !data.ConformancePolicyWsdlComponentType.IsNull() {
		data.ConformancePolicyWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.ConformancePolicyWsdlComponentType.ValueString() != "all" {
		data.ConformancePolicyWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicyWSDLComponentValue`); value.Exists() && !data.ConformancePolicyWsdlComponentValue.IsNull() {
		data.ConformancePolicyWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicyWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicySubscription`); value.Exists() && !data.ConformancePolicySubscription.IsNull() {
		data.ConformancePolicySubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicySubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConformancePolicyFragmentID`); value.Exists() && !data.ConformancePolicyFragmentId.IsNull() {
		data.ConformancePolicyFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConformancePolicyFragmentId = types.StringNull()
	}
}
