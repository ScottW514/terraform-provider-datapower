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

type DmSnmpContext struct {
	Context types.String `tfsdk:"context"`
	Domain  types.String `tfsdk:"domain"`
}

var DmSnmpContextObjectType = map[string]attr.Type{
	"context": types.StringType,
	"domain":  types.StringType,
}
var DmSnmpContextObjectDefault = map[string]attr.Value{
	"context": types.StringNull(),
	"domain":  types.StringNull(),
}
var DmSnmpContextDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"context": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Context Name", "", "").String,
			Computed:            true,
		},
		"domain": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Application Domain", "", "domain").String,
			Computed:            true,
		},
	},
}
var DmSnmpContextResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"context": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Context Name", "", "").String,
			Required:            true,
		},
		"domain": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Application Domain", "", "domain").String,
			Required:            true,
		},
	},
}

func (data DmSnmpContext) IsNull() bool {
	if !data.Context.IsNull() {
		return false
	}
	if !data.Domain.IsNull() {
		return false
	}
	return true
}

func (data DmSnmpContext) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Context.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Context`, data.Context.ValueString())
	}
	if !data.Domain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Domain`, data.Domain.ValueString())
	}
	return body
}

func (data *DmSnmpContext) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Context`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Context = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Context = types.StringNull()
	}
	if value := res.Get(pathRoot + `Domain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Domain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Domain = types.StringNull()
	}
}

func (data *DmSnmpContext) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Context`); value.Exists() && !data.Context.IsNull() {
		data.Context = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Context = types.StringNull()
	}
	if value := res.Get(pathRoot + `Domain`); value.Exists() && !data.Domain.IsNull() {
		data.Domain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Domain = types.StringNull()
	}
}
