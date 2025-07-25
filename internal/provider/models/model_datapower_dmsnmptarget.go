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

type DmSnmpTarget struct {
	Host          types.String `tfsdk:"host"`
	Port          types.Int64  `tfsdk:"port"`
	Community     types.String `tfsdk:"community"`
	TrapVersion   types.String `tfsdk:"trap_version"`
	SecurityName  types.String `tfsdk:"security_name"`
	SecurityLevel types.String `tfsdk:"security_level"`
}

var DmSnmpTargetObjectType = map[string]attr.Type{
	"host":           types.StringType,
	"port":           types.Int64Type,
	"community":      types.StringType,
	"trap_version":   types.StringType,
	"security_name":  types.StringType,
	"security_level": types.StringType,
}
var DmSnmpTargetObjectDefault = map[string]attr.Value{
	"host":           types.StringNull(),
	"port":           types.Int64Value(162),
	"community":      types.StringValue("public"),
	"trap_version":   types.StringValue("1"),
	"security_name":  types.StringNull(),
	"security_level": types.StringValue("authPriv"),
}
var DmSnmpTargetDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Host Address", "", "").String,
			Computed:            true,
		},
		"port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Port", "", "").AddDefaultValue("162").String,
			Computed:            true,
		},
		"community": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Community", "", "").AddDefaultValue("public").String,
			Computed:            true,
		},
		"trap_version": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Version", "", "").AddStringEnum("1", "2c", "3").AddDefaultValue("1").String,
			Computed:            true,
		},
		"security_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security Name", "", "").String,
			Computed:            true,
		},
		"security_level": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security Level", "", "").AddStringEnum("noAuthNoPriv", "authNoPriv", "authPriv").AddDefaultValue("authPriv").String,
			Computed:            true,
		},
	},
}
var DmSnmpTargetResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Host Address", "", "").String,
			Required:            true,
		},
		"port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Port", "", "").AddDefaultValue("162").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(162),
		},
		"community": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Community", "", "").AddDefaultValue("public").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("public"),
		},
		"trap_version": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Version", "", "").AddStringEnum("1", "2c", "3").AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("1", "2c", "3"),
			},
			Default: stringdefault.StaticString("1"),
		},
		"security_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security Name", "", "").String,
			Optional:            true,
		},
		"security_level": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Security Level", "", "").AddStringEnum("noAuthNoPriv", "authNoPriv", "authPriv").AddDefaultValue("authPriv").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("noAuthNoPriv", "authNoPriv", "authPriv"),
			},
			Default: stringdefault.StaticString("authPriv"),
		},
	},
}

func (data DmSnmpTarget) IsNull() bool {
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Community.IsNull() {
		return false
	}
	if !data.TrapVersion.IsNull() {
		return false
	}
	if !data.SecurityName.IsNull() {
		return false
	}
	if !data.SecurityLevel.IsNull() {
		return false
	}
	return true
}

func (data DmSnmpTarget) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Community.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Community`, data.Community.ValueString())
	}
	if !data.TrapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TrapVersion`, data.TrapVersion.ValueString())
	}
	if !data.SecurityName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecurityName`, data.SecurityName.ValueString())
	}
	if !data.SecurityLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecurityLevel`, data.SecurityLevel.ValueString())
	}
	return body
}

func (data *DmSnmpTarget) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Port = types.Int64Value(162)
	}
	if value := res.Get(pathRoot + `Community`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Community = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Community = types.StringValue("public")
	}
	if value := res.Get(pathRoot + `TrapVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TrapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TrapVersion = types.StringValue("1")
	}
	if value := res.Get(pathRoot + `SecurityName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecurityName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecurityName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecurityLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecurityLevel = types.StringValue("authPriv")
	}
}

func (data *DmSnmpTarget) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.Port.ValueInt64() != 162 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Community`); value.Exists() && !data.Community.IsNull() {
		data.Community = tfutils.ParseStringFromGJSON(value)
	} else if data.Community.ValueString() != "public" {
		data.Community = types.StringNull()
	}
	if value := res.Get(pathRoot + `TrapVersion`); value.Exists() && !data.TrapVersion.IsNull() {
		data.TrapVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.TrapVersion.ValueString() != "1" {
		data.TrapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityName`); value.Exists() && !data.SecurityName.IsNull() {
		data.SecurityName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecurityName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityLevel`); value.Exists() && !data.SecurityLevel.IsNull() {
		data.SecurityLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.SecurityLevel.ValueString() != "authPriv" {
		data.SecurityLevel = types.StringNull()
	}
}
