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

type DmRadiusServer struct {
	Number types.Int64  `tfsdk:"number"`
	Host   types.String `tfsdk:"host"`
	Port   types.Int64  `tfsdk:"port"`
	Secret types.String `tfsdk:"secret"`
}

var DmRadiusServerObjectType = map[string]attr.Type{
	"number": types.Int64Type,
	"host":   types.StringType,
	"port":   types.Int64Type,
	"secret": types.StringType,
}
var DmRadiusServerObjectDefault = map[string]attr.Value{
	"number": types.Int64Null(),
	"host":   types.StringNull(),
	"port":   types.Int64Value(1812),
	"secret": types.StringNull(),
}
var DmRadiusServerDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"number": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Number", "", "").AddIntegerRange(0, 2147483647).String,
			Computed:            true,
		},
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server address", "", "").String,
			Computed:            true,
		},
		"port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server port", "", "").AddDefaultValue("1812").String,
			Computed:            true,
		},
		"secret": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Secret", "", "").String,
			Computed:            true,
		},
	},
}
var DmRadiusServerResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"number": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Number", "", "").AddIntegerRange(0, 2147483647).String,
			Required:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 2147483647),
			},
		},
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server address", "", "").String,
			Required:            true,
		},
		"port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Server port", "", "").AddDefaultValue("1812").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(1812),
		},
		"secret": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Secret", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmRadiusServer) IsNull() bool {
	if !data.Number.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Secret.IsNull() {
		return false
	}
	return true
}

func (data DmRadiusServer) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Number.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Number`, data.Number.ValueInt64())
	}
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.Secret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Secret`, data.Secret.ValueString())
	}
	return body
}

func (data *DmRadiusServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Number`); value.Exists() {
		data.Number = types.Int64Value(value.Int())
	} else {
		data.Number = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(1812)
	}
	if value := res.Get(pathRoot + `Secret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Secret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Secret = types.StringNull()
	}
}

func (data *DmRadiusServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Number`); value.Exists() && !data.Number.IsNull() {
		data.Number = types.Int64Value(value.Int())
	} else {
		data.Number = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 1812 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Secret`); value.Exists() && !data.Secret.IsNull() {
		data.Secret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Secret = types.StringNull()
	}
}
