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

type DmLogEvent struct {
	Class    types.String `tfsdk:"class"`
	Priority types.String `tfsdk:"priority"`
}

var DmLogEventObjectType = map[string]attr.Type{
	"class":    types.StringType,
	"priority": types.StringType,
}
var DmLogEventObjectDefault = map[string]attr.Value{
	"class":    types.StringNull(),
	"priority": types.StringValue("notice"),
}
var DmLogEventDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"class": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Event category", "", "loglabel").String,
			Computed:            true,
		},
		"priority": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Minimum event priority", "", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("notice").String,
			Computed:            true,
		},
	},
}
var DmLogEventResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"class": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Event category", "", "loglabel").String,
			Optional:            true,
		},
		"priority": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Minimum event priority", "", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("notice").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
			},
			Default: stringdefault.StaticString("notice"),
		},
	},
}

func (data DmLogEvent) IsNull() bool {
	if !data.Class.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	return true
}

func (data DmLogEvent) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Class.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Class`, data.Class.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	return body
}

func (data *DmLogEvent) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Class`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Class = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Class = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("notice")
	}
}

func (data *DmLogEvent) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Class`); value.Exists() && !data.Class.IsNull() {
		data.Class = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Class = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "notice" {
		data.Priority = types.StringNull()
	}
}
