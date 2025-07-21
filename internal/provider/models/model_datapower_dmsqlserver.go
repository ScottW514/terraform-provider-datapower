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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSQLServer struct {
	Host         types.String `tfsdk:"host"`
	Port         types.Int64  `tfsdk:"port"`
	Type         types.String `tfsdk:"type"`
	DataSourceId types.String `tfsdk:"data_source_id"`
}

var DmSQLServerObjectType = map[string]attr.Type{
	"host":           types.StringType,
	"port":           types.Int64Type,
	"type":           types.StringType,
	"data_source_id": types.StringType,
}
var DmSQLServerObjectDefault = map[string]attr.Value{
	"host":           types.StringNull(),
	"port":           types.Int64Value(1521),
	"type":           types.StringValue("OracleListener"),
	"data_source_id": types.StringNull(),
}
var DmSQLServerDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data source host", "host", "").String,
			Computed:            true,
		},
		"port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data source port", "port", "").AddDefaultValue("1521").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "type", "").AddStringEnum("Unspecified", "OracleListener", "OracleONS").AddDefaultValue("OracleListener").String,
			Computed:            true,
		},
		"data_source_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data Source ID", "id", "").String,
			Computed:            true,
		},
	},
}
var DmSQLServerResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data source host", "host", "").String,
			Required:            true,
		},
		"port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data source port", "port", "").AddDefaultValue("1521").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(1521),
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Type", "type", "").AddStringEnum("Unspecified", "OracleListener", "OracleONS").AddDefaultValue("OracleListener").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("Unspecified", "OracleListener", "OracleONS"),
			},
			Default: stringdefault.StaticString("OracleListener"),
		},
		"data_source_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data Source ID", "id", "").String,
			Optional:            true,
		},
	},
}

func (data DmSQLServer) IsNull() bool {
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.DataSourceId.IsNull() {
		return false
	}
	return true
}

func (data DmSQLServer) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.DataSourceId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ID`, data.DataSourceId.ValueString())
	}
	return body
}

func (data *DmSQLServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(1521)
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("OracleListener")
	}
	if value := res.Get(pathRoot + `ID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataSourceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceId = types.StringNull()
	}
}

func (data *DmSQLServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 1521 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "OracleListener" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `ID`); value.Exists() && !data.DataSourceId.IsNull() {
		data.DataSourceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceId = types.StringNull()
	}
}
