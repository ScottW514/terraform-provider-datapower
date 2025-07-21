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

type DmSnmpCredMasked struct {
	EngineId       types.String `tfsdk:"engine_id"`
	AuthProtocol   types.String `tfsdk:"auth_protocol"`
	AuthSecretType types.String `tfsdk:"auth_secret_type"`
	AuthSecret     types.String `tfsdk:"auth_secret"`
	PrivProtocol   types.String `tfsdk:"priv_protocol"`
	PrivSecretType types.String `tfsdk:"priv_secret_type"`
	PrivSecret     types.String `tfsdk:"priv_secret"`
}

var DmSnmpCredMaskedObjectType = map[string]attr.Type{
	"engine_id":        types.StringType,
	"auth_protocol":    types.StringType,
	"auth_secret_type": types.StringType,
	"auth_secret":      types.StringType,
	"priv_protocol":    types.StringType,
	"priv_secret_type": types.StringType,
	"priv_secret":      types.StringType,
}
var DmSnmpCredMaskedObjectDefault = map[string]attr.Value{
	"engine_id":        types.StringNull(),
	"auth_protocol":    types.StringValue("sha"),
	"auth_secret_type": types.StringValue("password"),
	"auth_secret":      types.StringNull(),
	"priv_protocol":    types.StringValue("des"),
	"priv_secret_type": types.StringValue("password"),
	"priv_secret":      types.StringNull(),
}
var DmSnmpCredMaskedDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"engine_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"auth_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("none", "md5", "sha").AddDefaultValue("sha").String,
			Computed:            true,
		},
		"auth_secret_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("password", "key").AddDefaultValue("password").String,
			Computed:            true,
		},
		"auth_secret": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
		"priv_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("none", "des", "aes").AddDefaultValue("des").String,
			Computed:            true,
		},
		"priv_secret_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("password", "key").AddDefaultValue("password").String,
			Computed:            true,
		},
		"priv_secret": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Computed:            true,
		},
	},
}
var DmSnmpCredMaskedResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"engine_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Required:            true,
		},
		"auth_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("none", "md5", "sha").AddDefaultValue("sha").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "md5", "sha"),
			},
			Default: stringdefault.StaticString("sha"),
		},
		"auth_secret_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("password", "key").AddDefaultValue("password").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("password", "key"),
			},
			Default: stringdefault.StaticString("password"),
		},
		"auth_secret": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
		"priv_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("none", "des", "aes").AddDefaultValue("des").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "des", "aes"),
			},
			Default: stringdefault.StaticString("des"),
		},
		"priv_secret_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddStringEnum("password", "key").AddDefaultValue("password").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("password", "key"),
			},
			Default: stringdefault.StaticString("password"),
		},
		"priv_secret": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmSnmpCredMasked) IsNull() bool {
	if !data.EngineId.IsNull() {
		return false
	}
	if !data.AuthProtocol.IsNull() {
		return false
	}
	if !data.AuthSecretType.IsNull() {
		return false
	}
	if !data.AuthSecret.IsNull() {
		return false
	}
	if !data.PrivProtocol.IsNull() {
		return false
	}
	if !data.PrivSecretType.IsNull() {
		return false
	}
	if !data.PrivSecret.IsNull() {
		return false
	}
	return true
}

func (data DmSnmpCredMasked) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.EngineId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EngineID`, data.EngineId.ValueString())
	}
	if !data.AuthProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthProtocol`, data.AuthProtocol.ValueString())
	}
	if !data.AuthSecretType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthSecretType`, data.AuthSecretType.ValueString())
	}
	if !data.AuthSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthSecret`, data.AuthSecret.ValueString())
	}
	if !data.PrivProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivProtocol`, data.PrivProtocol.ValueString())
	}
	if !data.PrivSecretType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivSecretType`, data.PrivSecretType.ValueString())
	}
	if !data.PrivSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivSecret`, data.PrivSecret.ValueString())
	}
	return body
}

func (data *DmSnmpCredMasked) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EngineID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EngineId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthProtocol = types.StringValue("sha")
	}
	if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecretType = types.StringValue("password")
	}
	if value := res.Get(pathRoot + `AuthSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivProtocol = types.StringValue("des")
	}
	if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecretType = types.StringValue("password")
	}
	if value := res.Get(pathRoot + `PrivSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecret = types.StringNull()
	}
}

func (data *DmSnmpCredMasked) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EngineID`); value.Exists() && !data.EngineId.IsNull() {
		data.EngineId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && !data.AuthProtocol.IsNull() {
		data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthProtocol.ValueString() != "sha" {
		data.AuthProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && !data.AuthSecretType.IsNull() {
		data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthSecretType.ValueString() != "password" {
		data.AuthSecretType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthSecret`); value.Exists() && !data.AuthSecret.IsNull() {
		data.AuthSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && !data.PrivProtocol.IsNull() {
		data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.PrivProtocol.ValueString() != "des" {
		data.PrivProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && !data.PrivSecretType.IsNull() {
		data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
	} else if data.PrivSecretType.ValueString() != "password" {
		data.PrivSecretType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivSecret`); value.Exists() && !data.PrivSecret.IsNull() {
		data.PrivSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecret = types.StringNull()
	}
}
