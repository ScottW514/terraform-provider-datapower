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

type DmOAuthType struct {
	ClientType types.String `tfsdk:"client_type"`
	GrantType  types.String `tfsdk:"grant_type"`
}

var DmOAuthTypeObjectType = map[string]attr.Type{
	"client_type": types.StringType,
	"grant_type":  types.StringType,
}
var DmOAuthTypeObjectDefault = map[string]attr.Value{
	"client_type": types.StringValue("confidential"),
	"grant_type":  types.StringValue("implicit"),
}

func GetDmOAuthTypeDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmOAuthTypeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"client_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Client type",
				Computed:            true,
			},
			"grant_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Flow",
				Computed:            true,
			},
		},
	}
	return DmOAuthTypeDataSourceSchema
}
func GetDmOAuthTypeResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmOAuthTypeResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"client_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client type", "", "").AddStringEnum("confidential", "public").AddDefaultValue("confidential").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("confidential", "public"),
				},
				Default: stringdefault.StaticString("confidential"),
			},
			"grant_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Flow", "", "").AddStringEnum("implicit", "password", "application", "accessCode").AddDefaultValue("implicit").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("implicit", "password", "application", "accessCode"),
				},
				Default: stringdefault.StaticString("implicit"),
			},
		},
	}
	return DmOAuthTypeResourceSchema
}

func (data DmOAuthType) IsNull() bool {
	if !data.ClientType.IsNull() {
		return false
	}
	if !data.GrantType.IsNull() {
		return false
	}
	return true
}

func (data DmOAuthType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ClientType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientType`, data.ClientType.ValueString())
	}
	if !data.GrantType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GrantType`, data.GrantType.ValueString())
	}
	return body
}

func (data *DmOAuthType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientType = types.StringValue("confidential")
	}
	if value := res.Get(pathRoot + `GrantType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GrantType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GrantType = types.StringValue("implicit")
	}
}

func (data *DmOAuthType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && !data.ClientType.IsNull() {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientType.ValueString() != "confidential" {
		data.ClientType = types.StringNull()
	}
	if value := res.Get(pathRoot + `GrantType`); value.Exists() && !data.GrantType.IsNull() {
		data.GrantType = tfutils.ParseStringFromGJSON(value)
	} else if data.GrantType.ValueString() != "implicit" {
		data.GrantType = types.StringNull()
	}
}
