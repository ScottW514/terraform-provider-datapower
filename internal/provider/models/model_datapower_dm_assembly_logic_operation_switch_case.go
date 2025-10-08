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

type DmAssemblyLogicOperationSwitchCase struct {
	Execute     types.String `tfsdk:"execute"`
	OperationId types.String `tfsdk:"operation_id"`
	Path        types.String `tfsdk:"path"`
	Method      types.String `tfsdk:"method"`
}

var DmAssemblyLogicOperationSwitchCaseObjectType = map[string]attr.Type{
	"execute":      types.StringType,
	"operation_id": types.StringType,
	"path":         types.StringType,
	"method":       types.StringType,
}
var DmAssemblyLogicOperationSwitchCaseObjectDefault = map[string]attr.Value{
	"execute":      types.StringNull(),
	"operation_id": types.StringNull(),
	"path":         types.StringNull(),
	"method":       types.StringValue("GET"),
}

func GetDmAssemblyLogicOperationSwitchCaseDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAssemblyLogicOperationSwitchCaseDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"execute": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the API rule to run when the case is matched.",
				Computed:            true,
			},
			"operation_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the unique operation ID to match.",
				Computed:            true,
			},
			"path": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the relative API path to match.",
				Computed:            true,
			},
			"method": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the HTTP method to match.",
				Computed:            true,
			},
		},
	}
	return DmAssemblyLogicOperationSwitchCaseDataSourceSchema
}
func GetDmAssemblyLogicOperationSwitchCaseResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAssemblyLogicOperationSwitchCaseResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"execute": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API rule to run when the case is matched.", "", "").String,
				Required:            true,
			},
			"operation_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the unique operation ID to match.", "", "").String,
				Optional:            true,
			},
			"path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the relative API path to match.", "", "").String,
				Optional:            true,
			},
			"method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTTP method to match.", "", "").AddStringEnum("GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE").AddDefaultValue("GET").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("GET", "POST", "PUT", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE"),
				},
				Default: stringdefault.StaticString("GET"),
			},
		},
	}
	return DmAssemblyLogicOperationSwitchCaseResourceSchema
}

func (data DmAssemblyLogicOperationSwitchCase) IsNull() bool {
	if !data.Execute.IsNull() {
		return false
	}
	if !data.OperationId.IsNull() {
		return false
	}
	if !data.Path.IsNull() {
		return false
	}
	if !data.Method.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyLogicOperationSwitchCase) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Execute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Execute`, data.Execute.ValueString())
	}
	if !data.OperationId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OperationId`, data.OperationId.ValueString())
	}
	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Path`, data.Path.ValueString())
	}
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	return body
}

func (data *DmAssemblyLogicOperationSwitchCase) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
	if value := res.Get(pathRoot + `OperationId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OperationId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Method = types.StringValue("GET")
	}
}

func (data *DmAssemblyLogicOperationSwitchCase) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Execute`); value.Exists() && !data.Execute.IsNull() {
		data.Execute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Execute = types.StringNull()
	}
	if value := res.Get(pathRoot + `OperationId`); value.Exists() && !data.OperationId.IsNull() {
		data.OperationId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && !data.Method.IsNull() {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else if data.Method.ValueString() != "GET" {
		data.Method = types.StringNull()
	}
}
