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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmFileSystemUsage struct {
	Name              types.String `tfsdk:"name"`
	WarningThreshold  types.Int64  `tfsdk:"warning_threshold"`
	CriticalThreshold types.Int64  `tfsdk:"critical_threshold"`
}

var DmFileSystemUsageObjectType = map[string]attr.Type{
	"name":               types.StringType,
	"warning_threshold":  types.Int64Type,
	"critical_threshold": types.Int64Type,
}
var DmFileSystemUsageObjectDefault = map[string]attr.Value{
	"name":               types.StringNull(),
	"warning_threshold":  types.Int64Value(75),
	"critical_threshold": types.Int64Value(90),
}
var DmFileSystemUsageDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("File system", "name", "").AddStringEnum("system", "raid", "temporary", "mqroot", "mqbackup", "mqdiag", "mqerr", "mqtrace").String,
			Computed:            true,
		},
		"warning_threshold": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Warning threshold", "warning", "").AddIntegerRange(0, 100).AddDefaultValue("75").String,
			Computed:            true,
		},
		"critical_threshold": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Critical threshold", "critical", "").AddIntegerRange(0, 100).AddDefaultValue("90").String,
			Computed:            true,
		},
	},
}
var DmFileSystemUsageResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("File system", "name", "").AddStringEnum("system", "raid", "temporary", "mqroot", "mqbackup", "mqdiag", "mqerr", "mqtrace").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("system", "raid", "temporary", "mqroot", "mqbackup", "mqdiag", "mqerr", "mqtrace"),
			},
		},
		"warning_threshold": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Warning threshold", "warning", "").AddIntegerRange(0, 100).AddDefaultValue("75").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 100),
			},
			Default: int64default.StaticInt64(75),
		},
		"critical_threshold": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Critical threshold", "critical", "").AddIntegerRange(0, 100).AddDefaultValue("90").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 100),
			},
			Default: int64default.StaticInt64(90),
		},
	},
}

func (data DmFileSystemUsage) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.WarningThreshold.IsNull() {
		return false
	}
	if !data.CriticalThreshold.IsNull() {
		return false
	}
	return true
}

func (data DmFileSystemUsage) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.WarningThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WarningThreshold`, data.WarningThreshold.ValueInt64())
	}
	if !data.CriticalThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CriticalThreshold`, data.CriticalThreshold.ValueInt64())
	}
	return body
}

func (data *DmFileSystemUsage) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `WarningThreshold`); value.Exists() {
		data.WarningThreshold = types.Int64Value(value.Int())
	} else {
		data.WarningThreshold = types.Int64Value(75)
	}
	if value := res.Get(pathRoot + `CriticalThreshold`); value.Exists() {
		data.CriticalThreshold = types.Int64Value(value.Int())
	} else {
		data.CriticalThreshold = types.Int64Value(90)
	}
}

func (data *DmFileSystemUsage) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `WarningThreshold`); value.Exists() && !data.WarningThreshold.IsNull() {
		data.WarningThreshold = types.Int64Value(value.Int())
	} else if data.WarningThreshold.ValueInt64() != 75 {
		data.WarningThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CriticalThreshold`); value.Exists() && !data.CriticalThreshold.IsNull() {
		data.CriticalThreshold = types.Int64Value(value.Int())
	} else if data.CriticalThreshold.ValueInt64() != 90 {
		data.CriticalThreshold = types.Int64Null()
	}
}
