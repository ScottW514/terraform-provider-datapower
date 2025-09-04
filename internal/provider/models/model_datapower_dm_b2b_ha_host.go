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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmB2BHAHost struct {
	Hostname types.String `tfsdk:"hostname"`
	Port     types.Int64  `tfsdk:"port"`
}

var DmB2BHAHostHostnameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ha_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	AttrPath:    "../",
	Value:       []string{"true"},
}
var DmB2BHAHostPortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ha_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	AttrPath:    "../",
	Value:       []string{"true"},
}
var DmB2BHAHostHostnameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmB2BHAHostPortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmB2BHAHostObjectType = map[string]attr.Type{
	"hostname": types.StringType,
	"port":     types.Int64Type,
}
var DmB2BHAHostObjectDefault = map[string]attr.Value{
	"hostname": types.StringNull(),
	"port":     types.Int64Value(1320),
}

func GetDmB2BHAHostDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmB2BHAHostDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"hostname": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name of the other system in the high availability cluster.", "", "").AddRequiredWhen(DmB2BHAHostHostnameCondVal.String()).AddNotValidWhen(DmB2BHAHostHostnameIgnoreVal.String()).String,
				Computed:            true,
			},
			"port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the other system in the high availability cluster.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("1320").AddRequiredWhen(DmB2BHAHostPortCondVal.String()).AddNotValidWhen(DmB2BHAHostPortIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	DmB2BHAHostDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmB2BHAHostDataSourceSchema
}
func GetDmB2BHAHostResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmB2BHAHostResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmB2BHAHostObjectType,
				DmB2BHAHostObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"hostname": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name of the other system in the high availability cluster.", "", "").AddRequiredWhen(DmB2BHAHostHostnameCondVal.String()).AddNotValidWhen(DmB2BHAHostHostnameIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmB2BHAHostHostnameCondVal, DmB2BHAHostHostnameIgnoreVal, false),
				},
			},
			"port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the other system in the high availability cluster.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("1320").AddRequiredWhen(DmB2BHAHostPortCondVal.String()).AddNotValidWhen(DmB2BHAHostPortIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(DmB2BHAHostPortCondVal, DmB2BHAHostPortIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1320),
			},
		},
	}
	DmB2BHAHostResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmB2BHAHostResourceSchema.Required = true
	} else {
		DmB2BHAHostResourceSchema.Optional = true
		DmB2BHAHostResourceSchema.Computed = true
	}
	return DmB2BHAHostResourceSchema
}

func (data DmB2BHAHost) IsNull() bool {
	if !data.Hostname.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	return true
}

func (data DmB2BHAHost) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Hostname`, data.Hostname.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	return body
}

func (data *DmB2BHAHost) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(1320)
	}
}

func (data *DmB2BHAHost) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 1320 {
		data.Port = types.Int64Null()
	}
}
