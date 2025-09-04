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

type DmAssemblyActionSetVar struct {
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

var DmAssemblyActionSetVarValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "action",
	AttrType:    "String",
	AttrDefault: "set",
	Value:       []string{"clear"},
}
var DmAssemblyActionSetVarTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "action",
	AttrType:    "String",
	AttrDefault: "set",
	Value:       []string{"clear"},
}
var DmAssemblyActionSetVarValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "action",
	AttrType:    "String",
	AttrDefault: "set",
	Value:       []string{"clear"},
}

var DmAssemblyActionSetVarObjectType = map[string]attr.Type{
	"action": types.StringType,
	"name":   types.StringType,
	"type":   types.StringType,
	"value":  types.StringType,
}
var DmAssemblyActionSetVarObjectDefault = map[string]attr.Value{
	"action": types.StringValue("set"),
	"name":   types.StringNull(),
	"type":   types.StringValue("any"),
	"value":  types.StringNull(),
}

func GetDmAssemblyActionSetVarDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAssemblyActionSetVarDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to manage the variable.", "action", "").AddStringEnum("set", "add", "clear").AddDefaultValue("set").String,
				Computed:            true,
			},
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the variable. <p>You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short form <tt>$(property_name)</tt> when the assembly action does not have a property with the same name.</p>", "name", "").String,
				Computed:            true,
			},
			"type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the data type of the variable. You must define this property to set or add a variable.", "type", "").AddStringEnum("any", "string", "number", "boolean").AddDefaultValue("any").AddNotValidWhen(DmAssemblyActionSetVarTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the variable. You must define this property to set or add a variable. <p>You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short form <tt>$(property_name)</tt> when the assembly action does not have a property with the same name.</p><p>When you assign value, comply with the following rules. Otherwise, error occurs and the action fails.</p><ul><li>The value must match the specified data type: number, string, or Boolean.</li><li>The value for the <tt>message.status.code</tt> variable must be a valid HTTP status code.</li></ul>", "value", "").AddRequiredWhen(DmAssemblyActionSetVarValueCondVal.String()).AddNotValidWhen(DmAssemblyActionSetVarValueIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	return DmAssemblyActionSetVarDataSourceSchema
}
func GetDmAssemblyActionSetVarResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAssemblyActionSetVarResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to manage the variable.", "action", "").AddStringEnum("set", "add", "clear").AddDefaultValue("set").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("set", "add", "clear"),
				},
				Default: stringdefault.StaticString("set"),
			},
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the variable. <p>You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short form <tt>$(property_name)</tt> when the assembly action does not have a property with the same name.</p>", "name", "").String,
				Required:            true,
			},
			"type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the data type of the variable. You must define this property to set or add a variable.", "type", "").AddStringEnum("any", "string", "number", "boolean").AddDefaultValue("any").AddNotValidWhen(DmAssemblyActionSetVarTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("any", "string", "number", "boolean"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAssemblyActionSetVarTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("any"),
			},
			"value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the variable. You must define this property to set or add a variable. <p>You can reference a custom API property that resolves as the value. To reference an API property, use the <tt>$(api.properties. <i>property_name</i> )</tt> format, where <tt><i>property_name</i></tt> is the name of the property to reference. You can use the short form <tt>$(property_name)</tt> when the assembly action does not have a property with the same name.</p><p>When you assign value, comply with the following rules. Otherwise, error occurs and the action fails.</p><ul><li>The value must match the specified data type: number, string, or Boolean.</li><li>The value for the <tt>message.status.code</tt> variable must be a valid HTTP status code.</li></ul>", "value", "").AddRequiredWhen(DmAssemblyActionSetVarValueCondVal.String()).AddNotValidWhen(DmAssemblyActionSetVarValueIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAssemblyActionSetVarValueCondVal, DmAssemblyActionSetVarValueIgnoreVal, false),
				},
			},
		},
	}
	return DmAssemblyActionSetVarResourceSchema
}

func (data DmAssemblyActionSetVar) IsNull() bool {
	if !data.Action.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyActionSetVar) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmAssemblyActionSetVar) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringValue("set")
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("any")
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmAssemblyActionSetVar) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else if data.Action.ValueString() != "set" {
		data.Action = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "any" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
