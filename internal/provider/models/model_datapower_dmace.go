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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmACE struct {
	Access  types.String `tfsdk:"access"`
	Address types.String `tfsdk:"address"`
}

var DmACEObjectType = map[string]attr.Type{
	"access":  types.StringType,
	"address": types.StringType,
}
var DmACEObjectDefault = map[string]attr.Value{
	"access":  types.StringNull(),
	"address": types.StringNull(),
}
var DmACEDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"access": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Access", "", "").AddStringEnum("allow", "deny").String,
			Computed:            true,
		},
		"address": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Address range", "", "").String,
			Computed:            true,
		},
	},
}
var DmACEResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"access": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Access", "", "").AddStringEnum("allow", "deny").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("allow", "deny"),
			},
		},
		"address": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Address range", "", "").String,
			Required:            true,
		},
	},
}

func (data DmACE) IsNull() bool {
	if !data.Access.IsNull() {
		return false
	}
	if !data.Address.IsNull() {
		return false
	}
	return true
}

func (data DmACE) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Access.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Access`, data.Access.ValueString())
	}
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Address`, data.Address.ValueString())
	}
	return body
}

func (data *DmACE) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Access`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Access = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get(pathRoot + `Address`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Address = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Address = types.StringNull()
	}
}

func (data *DmACE) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Access`); value.Exists() && !data.Access.IsNull() {
		data.Access = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get(pathRoot + `Address`); value.Exists() && !data.Address.IsNull() {
		data.Address = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Address = types.StringNull()
	}
}
