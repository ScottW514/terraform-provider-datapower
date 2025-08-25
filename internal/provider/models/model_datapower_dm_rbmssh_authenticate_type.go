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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmRBMSSHAuthenticateType struct {
	Certificate types.Bool `tfsdk:"certificate"`
	Password    types.Bool `tfsdk:"password"`
}

var DmRBMSSHAuthenticateTypeObjectType = map[string]attr.Type{
	"certificate": types.BoolType,
	"password":    types.BoolType,
}
var DmRBMSSHAuthenticateTypeObjectDefault = map[string]attr.Value{
	"certificate": types.BoolValue(false),
	"password":    types.BoolValue(false),
}

func GetDmRBMSSHAuthenticateTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmRBMSSHAuthenticateTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"certificate": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CA-signed user certificate", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"password": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmRBMSSHAuthenticateTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmRBMSSHAuthenticateTypeDataSourceSchema
}
func GetDmRBMSSHAuthenticateTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmRBMSSHAuthenticateTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmRBMSSHAuthenticateTypeObjectType,
				DmRBMSSHAuthenticateTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"certificate": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CA-signed user certificate", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"password": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmRBMSSHAuthenticateTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmRBMSSHAuthenticateTypeResourceSchema.Required = true
	} else {
		DmRBMSSHAuthenticateTypeResourceSchema.Optional = true
		DmRBMSSHAuthenticateTypeResourceSchema.Computed = true
	}
	return DmRBMSSHAuthenticateTypeResourceSchema
}

func (data DmRBMSSHAuthenticateType) IsNull() bool {
	if !data.Certificate.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	return true
}

func (data DmRBMSSHAuthenticateType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Certificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`certificate`, tfutils.StringFromBool(data.Certificate, ""))
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`password`, tfutils.StringFromBool(data.Password, ""))
	}
	return body
}

func (data *DmRBMSSHAuthenticateType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `certificate`); value.Exists() {
		data.Certificate = tfutils.BoolFromString(value.String())
	} else {
		data.Certificate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() {
		data.Password = tfutils.BoolFromString(value.String())
	} else {
		data.Password = types.BoolNull()
	}
}

func (data *DmRBMSSHAuthenticateType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.BoolFromString(value.String())
	} else if data.Certificate.ValueBool() {
		data.Certificate = types.BoolNull()
	}
	if value := res.Get(pathRoot + `password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.BoolFromString(value.String())
	} else if data.Password.ValueBool() {
		data.Password = types.BoolNull()
	}
}
