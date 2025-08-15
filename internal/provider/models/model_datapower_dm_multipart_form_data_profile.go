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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmMultipartFormDataProfile struct {
	MaxParts              types.Int64 `tfsdk:"max_parts"`
	MaxSizePerPart        types.Int64 `tfsdk:"max_size_per_part"`
	MaxSize               types.Int64 `tfsdk:"max_size"`
	ContentTypeRestricted types.Bool  `tfsdk:"content_type_restricted"`
}

var DmMultipartFormDataProfileObjectType = map[string]attr.Type{
	"max_parts":               types.Int64Type,
	"max_size_per_part":       types.Int64Type,
	"max_size":                types.Int64Type,
	"content_type_restricted": types.BoolType,
}
var DmMultipartFormDataProfileObjectDefault = map[string]attr.Value{
	"max_parts":               types.Int64Value(4),
	"max_size_per_part":       types.Int64Value(5000000),
	"max_size":                types.Int64Value(5000000),
	"content_type_restricted": types.BoolValue(false),
}
var DmMultipartFormDataProfileDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"max_parts": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum number of parts allowed. Defaults to 4.", "", "").AddDefaultValue("4").String,
			Computed:            true,
		},
		"max_size_per_part": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum size allowed for any one part. Defaults to 5000000.", "", "").AddDefaultValue("5000000").String,
			Computed:            true,
		},
		"max_size": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum size allowed for all parts combined. Defaults to 5000000.", "", "").AddDefaultValue("5000000").String,
			Computed:            true,
		},
		"content_type_restricted": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("If set, this forces the individual form-data content types to be matched against the general list of request acceptable content-type expressions.", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmMultipartFormDataProfileResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmMultipartFormDataProfileObjectType,
			DmMultipartFormDataProfileObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"max_parts": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum number of parts allowed. Defaults to 4.", "", "").AddDefaultValue("4").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(4),
		},
		"max_size_per_part": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum size allowed for any one part. Defaults to 5000000.", "", "").AddDefaultValue("5000000").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(5000000),
		},
		"max_size": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Maximum size allowed for all parts combined. Defaults to 5000000.", "", "").AddDefaultValue("5000000").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(5000000),
		},
		"content_type_restricted": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("If set, this forces the individual form-data content types to be matched against the general list of request acceptable content-type expressions.", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmMultipartFormDataProfile) IsNull() bool {
	if !data.MaxParts.IsNull() {
		return false
	}
	if !data.MaxSizePerPart.IsNull() {
		return false
	}
	if !data.MaxSize.IsNull() {
		return false
	}
	if !data.ContentTypeRestricted.IsNull() {
		return false
	}
	return true
}
func GetDmMultipartFormDataProfileDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmMultipartFormDataProfileDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmMultipartFormDataProfileDataSourceSchema
}

func GetDmMultipartFormDataProfileResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmMultipartFormDataProfileResourceSchema.Required = true
	} else {
		DmMultipartFormDataProfileResourceSchema.Optional = true
		DmMultipartFormDataProfileResourceSchema.Computed = true
	}
	DmMultipartFormDataProfileResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmMultipartFormDataProfileResourceSchema
}

func (data DmMultipartFormDataProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.MaxParts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxParts`, data.MaxParts.ValueInt64())
	}
	if !data.MaxSizePerPart.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxSizePerPart`, data.MaxSizePerPart.ValueInt64())
	}
	if !data.MaxSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxSize`, data.MaxSize.ValueInt64())
	}
	if !data.ContentTypeRestricted.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContentTypeRestricted`, tfutils.StringFromBool(data.ContentTypeRestricted, ""))
	}
	return body
}

func (data *DmMultipartFormDataProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MaxParts`); value.Exists() {
		data.MaxParts = types.Int64Value(value.Int())
	} else {
		data.MaxParts = types.Int64Value(4)
	}
	if value := res.Get(pathRoot + `MaxSizePerPart`); value.Exists() {
		data.MaxSizePerPart = types.Int64Value(value.Int())
	} else {
		data.MaxSizePerPart = types.Int64Value(5000000)
	}
	if value := res.Get(pathRoot + `MaxSize`); value.Exists() {
		data.MaxSize = types.Int64Value(value.Int())
	} else {
		data.MaxSize = types.Int64Value(5000000)
	}
	if value := res.Get(pathRoot + `ContentTypeRestricted`); value.Exists() {
		data.ContentTypeRestricted = tfutils.BoolFromString(value.String())
	} else {
		data.ContentTypeRestricted = types.BoolNull()
	}
}

func (data *DmMultipartFormDataProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MaxParts`); value.Exists() && !data.MaxParts.IsNull() {
		data.MaxParts = types.Int64Value(value.Int())
	} else if data.MaxParts.ValueInt64() != 4 {
		data.MaxParts = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxSizePerPart`); value.Exists() && !data.MaxSizePerPart.IsNull() {
		data.MaxSizePerPart = types.Int64Value(value.Int())
	} else if data.MaxSizePerPart.ValueInt64() != 5000000 {
		data.MaxSizePerPart = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxSize`); value.Exists() && !data.MaxSize.IsNull() {
		data.MaxSize = types.Int64Value(value.Int())
	} else if data.MaxSize.ValueInt64() != 5000000 {
		data.MaxSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ContentTypeRestricted`); value.Exists() && !data.ContentTypeRestricted.IsNull() {
		data.ContentTypeRestricted = tfutils.BoolFromString(value.String())
	} else if data.ContentTypeRestricted.ValueBool() {
		data.ContentTypeRestricted = types.BoolNull()
	}
}
