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

type DmRateLimitInfoDomainNamed struct {
	Name   types.String `tfsdk:"name"`
	Action types.String `tfsdk:"action"`
}

var DmRateLimitInfoDomainNamedObjectType = map[string]attr.Type{
	"name":   types.StringType,
	"action": types.StringType,
}
var DmRateLimitInfoDomainNamedObjectDefault = map[string]attr.Value{
	"name":   types.StringNull(),
	"action": types.StringValue("consume"),
}

func GetDmRateLimitInfoDomainNamedDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmRateLimitInfoDomainNamedDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the rate limit definition.", "name", "rate_limit_definition").String,
				Computed:            true,
			},
			"action": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to apply to the rate limit definition.", "action", "").AddStringEnum("consume", "replenish", "check", "update").AddDefaultValue("consume").String,
				Computed:            true,
			},
		},
	}
	return DmRateLimitInfoDomainNamedDataSourceSchema
}
func GetDmRateLimitInfoDomainNamedResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmRateLimitInfoDomainNamedResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the rate limit definition.", "name", "rate_limit_definition").String,
				Required:            true,
			},
			"action": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the action to apply to the rate limit definition.", "action", "").AddStringEnum("consume", "replenish", "check", "update").AddDefaultValue("consume").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("consume", "replenish", "check", "update"),
				},
				Default: stringdefault.StaticString("consume"),
			},
		},
	}
	return DmRateLimitInfoDomainNamedResourceSchema
}

func (data DmRateLimitInfoDomainNamed) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Action.IsNull() {
		return false
	}
	return true
}

func (data DmRateLimitInfoDomainNamed) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Action`, data.Action.ValueString())
	}
	return body
}

func (data *DmRateLimitInfoDomainNamed) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Action = types.StringValue("consume")
	}
}

func (data *DmRateLimitInfoDomainNamed) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Action`); value.Exists() && !data.Action.IsNull() {
		data.Action = tfutils.ParseStringFromGJSON(value)
	} else if data.Action.ValueString() != "consume" {
		data.Action = types.StringNull()
	}
}
