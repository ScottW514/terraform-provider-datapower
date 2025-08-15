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

type DmAssemblyCatch struct {
	Error   types.String `tfsdk:"error"`
	Handler types.String `tfsdk:"handler"`
}

var DmAssemblyCatchObjectType = map[string]attr.Type{
	"error":   types.StringType,
	"handler": types.StringType,
}
var DmAssemblyCatchObjectDefault = map[string]attr.Value{
	"error":   types.StringNull(),
	"handler": types.StringNull(),
}
var DmAssemblyCatchDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"error": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of a custom error.", "", "").String,
			Computed:            true,
		},
		"handler": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of the error handler for the custom error.", "", "api_rule").String,
			Computed:            true,
		},
	},
}
var DmAssemblyCatchResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"error": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of a custom error.", "", "").String,
			Required:            true,
		},
		"handler": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of the error handler for the custom error.", "", "api_rule").String,
			Required:            true,
		},
	},
}

func (data DmAssemblyCatch) IsNull() bool {
	if !data.Error.IsNull() {
		return false
	}
	if !data.Handler.IsNull() {
		return false
	}
	return true
}

func (data DmAssemblyCatch) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Error.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Error`, data.Error.ValueString())
	}
	if !data.Handler.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Handler`, data.Handler.ValueString())
	}
	return body
}

func (data *DmAssemblyCatch) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Error`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Error = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Error = types.StringNull()
	}
	if value := res.Get(pathRoot + `Handler`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Handler = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Handler = types.StringNull()
	}
}

func (data *DmAssemblyCatch) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Error`); value.Exists() && !data.Error.IsNull() {
		data.Error = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Error = types.StringNull()
	}
	if value := res.Get(pathRoot + `Handler`); value.Exists() && !data.Handler.IsNull() {
		data.Handler = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Handler = types.StringNull()
	}
}
