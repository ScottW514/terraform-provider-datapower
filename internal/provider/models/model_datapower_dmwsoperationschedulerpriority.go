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

type DmWSOperationSchedulerPriority struct {
	SchedulerPriority                   types.String `tfsdk:"scheduler_priority"`
	SchedulerPriorityWsdlComponentType  types.String `tfsdk:"scheduler_priority_wsdl_component_type"`
	SchedulerPriorityWsdlComponentValue types.String `tfsdk:"scheduler_priority_wsdl_component_value"`
	SchedulerPrioritySubscription       types.String `tfsdk:"scheduler_priority_subscription"`
	SchedulerPriorityFragmentId         types.String `tfsdk:"scheduler_priority_fragment_id"`
}

var DmWSOperationSchedulerPriorityObjectType = map[string]attr.Type{
	"scheduler_priority":                      types.StringType,
	"scheduler_priority_wsdl_component_type":  types.StringType,
	"scheduler_priority_wsdl_component_value": types.StringType,
	"scheduler_priority_subscription":         types.StringType,
	"scheduler_priority_fragment_id":          types.StringType,
}
var DmWSOperationSchedulerPriorityObjectDefault = map[string]attr.Value{
	"scheduler_priority":                      types.StringValue("normal"),
	"scheduler_priority_wsdl_component_type":  types.StringValue("all"),
	"scheduler_priority_wsdl_component_value": types.StringNull(),
	"scheduler_priority_subscription":         types.StringNull(),
	"scheduler_priority_fragment_id":          types.StringNull(),
}
var DmWSOperationSchedulerPriorityDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"scheduler_priority": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service Priority", "", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
			Computed:            true,
		},
		"scheduler_priority_wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
		},
		"scheduler_priority_wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Computed:            true,
		},
		"scheduler_priority_subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Computed:            true,
		},
		"scheduler_priority_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSOperationSchedulerPriorityResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"scheduler_priority": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service Priority", "", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
			},
			Default: stringdefault.StaticString("normal"),
		},
		"scheduler_priority_wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"scheduler_priority_wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Optional:            true,
		},
		"scheduler_priority_subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Optional:            true,
		},
		"scheduler_priority_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSOperationSchedulerPriority) IsNull() bool {
	if !data.SchedulerPriority.IsNull() {
		return false
	}
	if !data.SchedulerPriorityWsdlComponentType.IsNull() {
		return false
	}
	if !data.SchedulerPriorityWsdlComponentValue.IsNull() {
		return false
	}
	if !data.SchedulerPrioritySubscription.IsNull() {
		return false
	}
	if !data.SchedulerPriorityFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSOperationSchedulerPriority) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SchedulerPriority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchedulerPriority`, data.SchedulerPriority.ValueString())
	}
	if !data.SchedulerPriorityWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchedulerPriorityWSDLComponentType`, data.SchedulerPriorityWsdlComponentType.ValueString())
	}
	if !data.SchedulerPriorityWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchedulerPriorityWSDLComponentValue`, data.SchedulerPriorityWsdlComponentValue.ValueString())
	}
	if !data.SchedulerPrioritySubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchedulerPrioritySubscription`, data.SchedulerPrioritySubscription.ValueString())
	}
	if !data.SchedulerPriorityFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchedulerPriorityFragmentID`, data.SchedulerPriorityFragmentId.ValueString())
	}
	return body
}

func (data *DmWSOperationSchedulerPriority) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SchedulerPriority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchedulerPriority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `SchedulerPriorityWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchedulerPriorityWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriorityWsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `SchedulerPriorityWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchedulerPriorityWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriorityWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPrioritySubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchedulerPrioritySubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPrioritySubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPriorityFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchedulerPriorityFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriorityFragmentId = types.StringNull()
	}
}

func (data *DmWSOperationSchedulerPriority) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SchedulerPriority`); value.Exists() && !data.SchedulerPriority.IsNull() {
		data.SchedulerPriority = tfutils.ParseStringFromGJSON(value)
	} else if data.SchedulerPriority.ValueString() != "normal" {
		data.SchedulerPriority = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPriorityWSDLComponentType`); value.Exists() && !data.SchedulerPriorityWsdlComponentType.IsNull() {
		data.SchedulerPriorityWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.SchedulerPriorityWsdlComponentType.ValueString() != "all" {
		data.SchedulerPriorityWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPriorityWSDLComponentValue`); value.Exists() && !data.SchedulerPriorityWsdlComponentValue.IsNull() {
		data.SchedulerPriorityWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriorityWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPrioritySubscription`); value.Exists() && !data.SchedulerPrioritySubscription.IsNull() {
		data.SchedulerPrioritySubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPrioritySubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchedulerPriorityFragmentID`); value.Exists() && !data.SchedulerPriorityFragmentId.IsNull() {
		data.SchedulerPriorityFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchedulerPriorityFragmentId = types.StringNull()
	}
}
