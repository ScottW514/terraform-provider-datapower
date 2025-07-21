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

type DmFTPPolicy struct {
	RegExp         types.String `tfsdk:"reg_exp"`
	Passive        types.String `tfsdk:"passive"`
	AuthTls        types.String `tfsdk:"auth_tls"`
	UseCcc         types.String `tfsdk:"use_ccc"`
	EncryptData    types.String `tfsdk:"encrypt_data"`
	DataType       types.String `tfsdk:"data_type"`
	SlashStou      types.String `tfsdk:"slash_stou"`
	QuotedCommands types.String `tfsdk:"quoted_commands"`
	SizeCheck      types.String `tfsdk:"size_check"`
}

var DmFTPPolicyObjectType = map[string]attr.Type{
	"reg_exp":         types.StringType,
	"passive":         types.StringType,
	"auth_tls":        types.StringType,
	"use_ccc":         types.StringType,
	"encrypt_data":    types.StringType,
	"data_type":       types.StringType,
	"slash_stou":      types.StringType,
	"quoted_commands": types.StringType,
	"size_check":      types.StringType,
}
var DmFTPPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":         types.StringNull(),
	"passive":         types.StringValue("pasv-req"),
	"auth_tls":        types.StringValue("auth-off"),
	"use_ccc":         types.StringValue("ccc-off"),
	"encrypt_data":    types.StringValue("enc-data-off"),
	"data_type":       types.StringValue("binary"),
	"slash_stou":      types.StringValue("slash-stou-on"),
	"quoted_commands": types.StringNull(),
	"size_check":      types.StringValue("size-check-optional"),
}
var DmFTPPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Computed:            true,
		},
		"passive": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Passive mode", "", "").AddStringEnum("pasv-off", "pasv-opt", "pasv-req").AddDefaultValue("pasv-req").String,
			Computed:            true,
		},
		"auth_tls": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Encrypt command connection", "", "").AddStringEnum("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp").AddDefaultValue("auth-off").String,
			Computed:            true,
		},
		"use_ccc": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stop command encryption after authentication", "", "").AddStringEnum("ccc-off", "ccc-opt", "ccc-req").AddDefaultValue("ccc-off").String,
			Computed:            true,
		},
		"encrypt_data": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Encrypt file transfers", "", "").AddStringEnum("enc-data-off", "enc-data-opt", "enc-data-req").AddDefaultValue("enc-data-off").String,
			Computed:            true,
		},
		"data_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data type", "", "").AddStringEnum("ascii", "binary").AddDefaultValue("binary").String,
			Computed:            true,
		},
		"slash_stou": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Write unique file name if trailing slash", "", "").AddStringEnum("slash-stou-off", "slash-stou-on").AddDefaultValue("slash-stou-on").String,
			Computed:            true,
		},
		"quoted_commands": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Quoted commands", "", "ftpquotecommands").String,
			Computed:            true,
		},
		"size_check": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Size check", "", "").AddStringEnum("size-check-optional", "size-check-disabled").AddDefaultValue("size-check-optional").String,
			Computed:            true,
		},
	},
}
var DmFTPPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL matching expression", "", "").String,
			Required:            true,
		},
		"passive": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Passive mode", "", "").AddStringEnum("pasv-off", "pasv-opt", "pasv-req").AddDefaultValue("pasv-req").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("pasv-off", "pasv-opt", "pasv-req"),
			},
			Default: stringdefault.StaticString("pasv-req"),
		},
		"auth_tls": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Encrypt command connection", "", "").AddStringEnum("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp").AddDefaultValue("auth-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp"),
			},
			Default: stringdefault.StaticString("auth-off"),
		},
		"use_ccc": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Stop command encryption after authentication", "", "").AddStringEnum("ccc-off", "ccc-opt", "ccc-req").AddDefaultValue("ccc-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("ccc-off", "ccc-opt", "ccc-req"),
			},
			Default: stringdefault.StaticString("ccc-off"),
		},
		"encrypt_data": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Encrypt file transfers", "", "").AddStringEnum("enc-data-off", "enc-data-opt", "enc-data-req").AddDefaultValue("enc-data-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("enc-data-off", "enc-data-opt", "enc-data-req"),
			},
			Default: stringdefault.StaticString("enc-data-off"),
		},
		"data_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Data type", "", "").AddStringEnum("ascii", "binary").AddDefaultValue("binary").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("ascii", "binary"),
			},
			Default: stringdefault.StaticString("binary"),
		},
		"slash_stou": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Write unique file name if trailing slash", "", "").AddStringEnum("slash-stou-off", "slash-stou-on").AddDefaultValue("slash-stou-on").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("slash-stou-off", "slash-stou-on"),
			},
			Default: stringdefault.StaticString("slash-stou-on"),
		},
		"quoted_commands": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Quoted commands", "", "ftpquotecommands").String,
			Optional:            true,
		},
		"size_check": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Size check", "", "").AddStringEnum("size-check-optional", "size-check-disabled").AddDefaultValue("size-check-optional").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("size-check-optional", "size-check-disabled"),
			},
			Default: stringdefault.StaticString("size-check-optional"),
		},
	},
}

func (data DmFTPPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.Passive.IsNull() {
		return false
	}
	if !data.AuthTls.IsNull() {
		return false
	}
	if !data.UseCcc.IsNull() {
		return false
	}
	if !data.EncryptData.IsNull() {
		return false
	}
	if !data.DataType.IsNull() {
		return false
	}
	if !data.SlashStou.IsNull() {
		return false
	}
	if !data.QuotedCommands.IsNull() {
		return false
	}
	if !data.SizeCheck.IsNull() {
		return false
	}
	return true
}

func (data DmFTPPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.Passive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Passive`, data.Passive.ValueString())
	}
	if !data.AuthTls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthTLS`, data.AuthTls.ValueString())
	}
	if !data.UseCcc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCCC`, data.UseCcc.ValueString())
	}
	if !data.EncryptData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptData`, data.EncryptData.ValueString())
	}
	if !data.DataType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataType`, data.DataType.ValueString())
	}
	if !data.SlashStou.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SlashSTOU`, data.SlashStou.ValueString())
	}
	if !data.QuotedCommands.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QuotedCommands`, data.QuotedCommands.ValueString())
	}
	if !data.SizeCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SizeCheck`, data.SizeCheck.ValueString())
	}
	return body
}

func (data *DmFTPPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Passive = types.StringValue("pasv-req")
	}
	if value := res.Get(pathRoot + `AuthTLS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthTls = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthTls = types.StringValue("auth-off")
	}
	if value := res.Get(pathRoot + `UseCCC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UseCcc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UseCcc = types.StringValue("ccc-off")
	}
	if value := res.Get(pathRoot + `EncryptData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptData = types.StringValue("enc-data-off")
	}
	if value := res.Get(pathRoot + `DataType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataType = types.StringValue("binary")
	}
	if value := res.Get(pathRoot + `SlashSTOU`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SlashStou = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SlashStou = types.StringValue("slash-stou-on")
	}
	if value := res.Get(pathRoot + `QuotedCommands`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QuotedCommands = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommands = types.StringNull()
	}
	if value := res.Get(pathRoot + `SizeCheck`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SizeCheck = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SizeCheck = types.StringValue("size-check-optional")
	}
}

func (data *DmFTPPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && !data.Passive.IsNull() {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else if data.Passive.ValueString() != "pasv-req" {
		data.Passive = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthTLS`); value.Exists() && !data.AuthTls.IsNull() {
		data.AuthTls = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthTls.ValueString() != "auth-off" {
		data.AuthTls = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseCCC`); value.Exists() && !data.UseCcc.IsNull() {
		data.UseCcc = tfutils.ParseStringFromGJSON(value)
	} else if data.UseCcc.ValueString() != "ccc-off" {
		data.UseCcc = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptData`); value.Exists() && !data.EncryptData.IsNull() {
		data.EncryptData = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptData.ValueString() != "enc-data-off" {
		data.EncryptData = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataType`); value.Exists() && !data.DataType.IsNull() {
		data.DataType = tfutils.ParseStringFromGJSON(value)
	} else if data.DataType.ValueString() != "binary" {
		data.DataType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SlashSTOU`); value.Exists() && !data.SlashStou.IsNull() {
		data.SlashStou = tfutils.ParseStringFromGJSON(value)
	} else if data.SlashStou.ValueString() != "slash-stou-on" {
		data.SlashStou = types.StringNull()
	}
	if value := res.Get(pathRoot + `QuotedCommands`); value.Exists() && !data.QuotedCommands.IsNull() {
		data.QuotedCommands = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommands = types.StringNull()
	}
	if value := res.Get(pathRoot + `SizeCheck`); value.Exists() && !data.SizeCheck.IsNull() {
		data.SizeCheck = tfutils.ParseStringFromGJSON(value)
	} else if data.SizeCheck.ValueString() != "size-check-optional" {
		data.SizeCheck = types.StringNull()
	}
}
