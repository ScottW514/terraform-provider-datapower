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

type DmAssemblyActionRedact struct {
	Path   types.String `tfsdk:"path"`
	Action types.String `tfsdk:"action"`
}

var DmAssemblyActionRedactObjectType = map[string]attr.Type{
	"path":   types.StringType,
	"action": types.StringType,
}
var DmAssemblyActionRedactObjectDefault = map[string]attr.Value{
	"path":   types.StringNull(),
	"action": types.StringValue("redact"),
}

func GetDmAssemblyActionRedactDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAssemblyActionRedactDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"path": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JSONata path expression to the content.", "path", "").String,
				Computed:            true,
			},
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to redact or remove the content.", "action", "").AddStringEnum("redact", "remove").AddDefaultValue("redact").String,
				Computed:            true,
			},
		},
	}
	return DmAssemblyActionRedactDataSourceSchema
}
func GetDmAssemblyActionRedactResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAssemblyActionRedactResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the JSONata path expression to the content.", "path", "").String,
				Required:            true,
			},
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to redact or remove the content.", "action", "").AddStringEnum("redact", "remove").AddDefaultValue("redact").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("redact", "remove"),
				},
				Default: stringdefault.StaticString("redact"),
			},
		},
	}
	return DmAssemblyActionRedactResourceSchema
}

func (data DmAssemblyActionRedact) IsNull() bool {
	if !data.Path.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyActionRedact) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Path.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Path`, data.Path.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmAssemblyActionRedact) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringValue("redact")
	}
}

func (data *DmAssemblyActionRedact) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Path = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else if data.Action.ValueString() != "redact" {
		data.Action = types.StringNull()
	}
}
