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

type DmDefinitionLink struct {
	ShortName  types.String `tfsdk:"short_name"`
	Definition types.String `tfsdk:"definition"`
}

var DmDefinitionLinkObjectType = map[string]attr.Type{
	"short_name": types.StringType,
	"definition": types.StringType,
}
var DmDefinitionLinkObjectDefault = map[string]attr.Value{
	"short_name": types.StringNull(),
	"definition": types.StringNull(),
}
var DmDefinitionLinkDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"short_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Short name", "name", "").String,
			Computed:            true,
		},
		"definition": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Definition", "definition", "ratelimitdefinition").String,
			Computed:            true,
		},
	},
}
var DmDefinitionLinkResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"short_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Short name", "name", "").String,
			Required:            true,
		},
		"definition": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Definition", "definition", "ratelimitdefinition").String,
			Required:            true,
		},
	},
}

func (data DmDefinitionLink) IsNull() bool {
	if !data.ShortName.IsNull() {
		return false
	}
	if !data.Definition.IsNull() {
		return false
	}
	return true
}

func (data DmDefinitionLink) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.ShortName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ShortName`, data.ShortName.ValueString())
	}
	if !data.Definition.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Definition`, data.Definition.ValueString())
	}
	return body
}

func (data *DmDefinitionLink) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ShortName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ShortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Definition`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Definition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Definition = types.StringNull()
	}
}

func (data *DmDefinitionLink) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ShortName`); value.Exists() && !data.ShortName.IsNull() {
		data.ShortName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Definition`); value.Exists() && !data.Definition.IsNull() {
		data.Definition = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Definition = types.StringNull()
	}
}
