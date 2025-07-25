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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSUserToggles struct {
	WsdlName        types.String             `tfsdk:"wsdl_name"`
	ServiceName     types.String             `tfsdk:"service_name"`
	ServicePortName types.String             `tfsdk:"service_port_name"`
	PortTypeName    types.String             `tfsdk:"port_type_name"`
	BindingName     types.String             `tfsdk:"binding_name"`
	OperationName   types.String             `tfsdk:"operation_name"`
	Toggles         *DmWSDLUserPolicyToggles `tfsdk:"toggles"`
	Subscription    types.String             `tfsdk:"subscription"`
	UseFragmentId   types.Bool               `tfsdk:"use_fragment_id"`
	FragmentId      types.String             `tfsdk:"fragment_id"`
}

var DmWSUserTogglesObjectType = map[string]attr.Type{
	"wsdl_name":         types.StringType,
	"service_name":      types.StringType,
	"service_port_name": types.StringType,
	"port_type_name":    types.StringType,
	"binding_name":      types.StringType,
	"operation_name":    types.StringType,
	"toggles":           types.ObjectType{AttrTypes: DmWSDLUserPolicyTogglesObjectType},
	"subscription":      types.StringType,
	"use_fragment_id":   types.BoolType,
	"fragment_id":       types.StringType,
}
var DmWSUserTogglesObjectDefault = map[string]attr.Value{
	"wsdl_name":         types.StringNull(),
	"service_name":      types.StringNull(),
	"service_port_name": types.StringNull(),
	"port_type_name":    types.StringNull(),
	"binding_name":      types.StringNull(),
	"operation_name":    types.StringNull(),
	"toggles":           types.ObjectValueMust(DmWSDLUserPolicyTogglesObjectType, DmWSDLUserPolicyTogglesObjectDefault),
	"subscription":      types.StringNull(),
	"use_fragment_id":   types.BoolValue(false),
	"fragment_id":       types.StringNull(),
}
var DmWSUserTogglesDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"wsdl_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL file", "", "").String,
			Computed:            true,
		},
		"service_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "", "").String,
			Computed:            true,
		},
		"service_port_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "", "").String,
			Computed:            true,
		},
		"port_type_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PortType", "", "").String,
			Computed:            true,
		},
		"binding_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Binding", "", "").String,
			Computed:            true,
		},
		"operation_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Operation", "", "").String,
			Computed:            true,
		},
		"toggles": GetDmWSDLUserPolicyTogglesDataSourceSchema("Policy Toggles", "", ""),
		"subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "", "").String,
			Computed:            true,
		},
		"use_fragment_id": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Fragment ID", "use-fragid", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSUserTogglesResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"wsdl_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL file", "", "").String,
			Required:            true,
		},
		"service_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Service", "", "").String,
			Required:            true,
		},
		"service_port_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Port", "", "").String,
			Required:            true,
		},
		"port_type_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("PortType", "", "").String,
			Required:            true,
		},
		"binding_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Binding", "", "").String,
			Required:            true,
		},
		"operation_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Operation", "", "").String,
			Required:            true,
		},
		"toggles": GetDmWSDLUserPolicyTogglesResourceSchema("Policy Toggles", "", "", false),
		"subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "", "").String,
			Optional:            true,
		},
		"use_fragment_id": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use Fragment ID", "use-fragid", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSUserToggles) IsNull() bool {
	if !data.WsdlName.IsNull() {
		return false
	}
	if !data.ServiceName.IsNull() {
		return false
	}
	if !data.ServicePortName.IsNull() {
		return false
	}
	if !data.PortTypeName.IsNull() {
		return false
	}
	if !data.BindingName.IsNull() {
		return false
	}
	if !data.OperationName.IsNull() {
		return false
	}
	if data.Toggles != nil {
		if !data.Toggles.IsNull() {
			return false
		}
	}
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.UseFragmentId.IsNull() {
		return false
	}
	if !data.FragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSUserToggles) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.WsdlName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLName`, data.WsdlName.ValueString())
	}
	if !data.ServiceName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServiceName`, data.ServiceName.ValueString())
	}
	if !data.ServicePortName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServicePortName`, data.ServicePortName.ValueString())
	}
	if !data.PortTypeName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PortTypeName`, data.PortTypeName.ValueString())
	}
	if !data.BindingName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BindingName`, data.BindingName.ValueString())
	}
	if !data.OperationName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OperationName`, data.OperationName.ValueString())
	}
	if data.Toggles != nil {
		if !data.Toggles.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Toggles`, data.Toggles.ToBody(ctx, ""))
		}
	}
	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.UseFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseFragmentID`, tfutils.StringFromBool(data.UseFragmentId, ""))
	}
	if !data.FragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FragmentID`, data.FragmentId.ValueString())
	}
	return body
}

func (data *DmWSUserToggles) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServiceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServiceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServicePortName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServicePortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServicePortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PortTypeName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PortTypeName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PortTypeName = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindingName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BindingName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindingName = types.StringNull()
	}
	if value := res.Get(pathRoot + `OperationName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Toggles`); value.Exists() {
		data.Toggles = &DmWSDLUserPolicyToggles{}
		data.Toggles.FromBody(ctx, "", value)
	} else {
		data.Toggles = nil
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseFragmentID`); value.Exists() {
		data.UseFragmentId = tfutils.BoolFromString(value.String())
	} else {
		data.UseFragmentId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FragmentId = types.StringNull()
	}
}

func (data *DmWSUserToggles) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLName`); value.Exists() && !data.WsdlName.IsNull() {
		data.WsdlName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceName`); value.Exists() && !data.ServiceName.IsNull() {
		data.ServiceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServiceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServicePortName`); value.Exists() && !data.ServicePortName.IsNull() {
		data.ServicePortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServicePortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PortTypeName`); value.Exists() && !data.PortTypeName.IsNull() {
		data.PortTypeName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PortTypeName = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindingName`); value.Exists() && !data.BindingName.IsNull() {
		data.BindingName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindingName = types.StringNull()
	}
	if value := res.Get(pathRoot + `OperationName`); value.Exists() && !data.OperationName.IsNull() {
		data.OperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Toggles`); value.Exists() {
		data.Toggles.UpdateFromBody(ctx, "", value)
	} else {
		data.Toggles = nil
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseFragmentID`); value.Exists() && !data.UseFragmentId.IsNull() {
		data.UseFragmentId = tfutils.BoolFromString(value.String())
	} else if data.UseFragmentId.ValueBool() {
		data.UseFragmentId = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FragmentID`); value.Exists() && !data.FragmentId.IsNull() {
		data.FragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FragmentId = types.StringNull()
	}
}
