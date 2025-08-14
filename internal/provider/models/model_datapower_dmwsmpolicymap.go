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

type DmWSMPolicyMap struct {
	WsdlComponentType  types.String `tfsdk:"wsdl_component_type"`
	WsdlComponentValue types.String `tfsdk:"wsdl_component_value"`
	Match              types.String `tfsdk:"match"`
	Rule               types.String `tfsdk:"rule"`
	Subscription       types.String `tfsdk:"subscription"`
	WsdlFragmentId     types.String `tfsdk:"wsdl_fragment_id"`
}

var DmWSMPolicyMapObjectType = map[string]attr.Type{
	"wsdl_component_type":  types.StringType,
	"wsdl_component_value": types.StringType,
	"match":                types.StringType,
	"rule":                 types.StringType,
	"subscription":         types.StringType,
	"wsdl_fragment_id":     types.StringType,
}
var DmWSMPolicyMapObjectDefault = map[string]attr.Value{
	"wsdl_component_type":  types.StringValue("all"),
	"wsdl_component_value": types.StringNull(),
	"match":                types.StringNull(),
	"rule":                 types.StringNull(),
	"subscription":         types.StringNull(),
	"wsdl_fragment_id":     types.StringNull(),
}
var DmWSMPolicyMapDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
		},
		"wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field. The selected rule will be run only if the component named here would be used in processing the client request.", "", "").String,
			Computed:            true,
		},
		"match": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select an existing Matching Rule.", "", "matching").String,
			Computed:            true,
		},
		"rule": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a WS-Proxy Processing Rule to run for matching transactions.", "", "wsstylepolicyrule").String,
			Computed:            true,
		},
		"subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription. The selected rule will be run for requests that correspond to services that belong to this subscription.", "", "").String,
			Computed:            true,
		},
		"wsdl_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSMPolicyMapResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field. The selected rule will be run only if the component named here would be used in processing the client request.", "", "").String,
			Optional:            true,
		},
		"match": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select an existing Matching Rule.", "", "matching").String,
			Required:            true,
		},
		"rule": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a WS-Proxy Processing Rule to run for matching transactions.", "", "wsstylepolicyrule").String,
			Required:            true,
		},
		"subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription. The selected rule will be run for requests that correspond to services that belong to this subscription.", "", "").String,
			Optional:            true,
		},
		"wsdl_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSMPolicyMap) IsNull() bool {
	if !data.WsdlComponentType.IsNull() {
		return false
	}
	if !data.WsdlComponentValue.IsNull() {
		return false
	}
	if !data.Match.IsNull() {
		return false
	}
	if !data.Rule.IsNull() {
		return false
	}
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.WsdlFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSMPolicyMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.WsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLComponentType`, data.WsdlComponentType.ValueString())
	}
	if !data.WsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLComponentValue`, data.WsdlComponentValue.ValueString())
	}
	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Rule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rule`, data.Rule.ValueString())
	}
	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.WsdlFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLFragmentID`, data.WsdlFragmentId.ValueString())
	}
	return body
}

func (data *DmWSMPolicyMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `WSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlFragmentId = types.StringNull()
	}
}

func (data *DmWSMPolicyMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLComponentType`); value.Exists() && !data.WsdlComponentType.IsNull() {
		data.WsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlComponentType.ValueString() != "all" {
		data.WsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLComponentValue`); value.Exists() && !data.WsdlComponentValue.IsNull() {
		data.WsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLFragmentID`); value.Exists() && !data.WsdlFragmentId.IsNull() {
		data.WsdlFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlFragmentId = types.StringNull()
	}
}
