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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmBasicAuthPolicy struct {
	RegExp        types.String `tfsdk:"reg_exp"`
	UserName      types.String `tfsdk:"user_name"`
	PasswordAlias types.String `tfsdk:"password_alias"`
}

var DmBasicAuthPolicyObjectType = map[string]attr.Type{
	"reg_exp":        types.StringType,
	"user_name":      types.StringType,
	"password_alias": types.StringType,
}
var DmBasicAuthPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":        types.StringNull(),
	"user_name":      types.StringNull(),
	"password_alias": types.StringNull(),
}
var DmBasicAuthPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Computed:            true,
		},
		"user_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the username for basic authentication.", "", "").String,
			Computed:            true,
		},
		"password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias that maps to the password for the username.", "", "password_alias").String,
			Computed:            true,
		},
	},
}
var DmBasicAuthPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Required:            true,
		},
		"user_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the username for basic authentication.", "", "").String,
			Required:            true,
		},
		"password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias that maps to the password for the username.", "", "password_alias").String,
			Optional:            true,
		},
	},
}

func (data DmBasicAuthPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	return true
}

func (data DmBasicAuthPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	return body
}

func (data *DmBasicAuthPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
}

func (data *DmBasicAuthPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
}
