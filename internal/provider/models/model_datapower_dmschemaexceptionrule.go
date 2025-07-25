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

type DmSchemaExceptionRule struct {
	XPath         types.String `tfsdk:"x_path"`
	ExceptionType types.String `tfsdk:"exception_type"`
}

var DmSchemaExceptionRuleObjectType = map[string]attr.Type{
	"x_path":         types.StringType,
	"exception_type": types.StringType,
}
var DmSchemaExceptionRuleObjectDefault = map[string]attr.Value{
	"x_path":         types.StringNull(),
	"exception_type": types.StringValue("AllowEncrypted"),
}
var DmSchemaExceptionRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"x_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath Expression", "", "").String,
			Computed:            true,
		},
		"exception_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").AddStringEnum("AllowEncrypted", "RequireEncrypted").AddDefaultValue("AllowEncrypted").String,
			Computed:            true,
		},
	},
}
var DmSchemaExceptionRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"x_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("XPath Expression", "", "").String,
			Required:            true,
		},
		"exception_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "", "").AddStringEnum("AllowEncrypted", "RequireEncrypted").AddDefaultValue("AllowEncrypted").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("AllowEncrypted", "RequireEncrypted"),
			},
			Default: stringdefault.StaticString("AllowEncrypted"),
		},
	},
}

func (data DmSchemaExceptionRule) IsNull() bool {
	if !data.XPath.IsNull() {
		return false
	}
	if !data.ExceptionType.IsNull() {
		return false
	}
	return true
}

func (data DmSchemaExceptionRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.XPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.XPath.ValueString())
	}
	if !data.ExceptionType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExceptionType`, data.ExceptionType.ValueString())
	}
	return body
}

func (data *DmSchemaExceptionRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExceptionType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExceptionType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExceptionType = types.StringValue("AllowEncrypted")
	}
}

func (data *DmSchemaExceptionRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.XPath.IsNull() {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExceptionType`); value.Exists() && !data.ExceptionType.IsNull() {
		data.ExceptionType = tfutils.ParseStringFromGJSON(value)
	} else if data.ExceptionType.ValueString() != "AllowEncrypted" {
		data.ExceptionType = types.StringNull()
	}
}
