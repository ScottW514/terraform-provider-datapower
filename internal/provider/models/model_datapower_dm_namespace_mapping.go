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

type DmNamespaceMapping struct {
	Prefix types.String `tfsdk:"prefix"`
	Uri    types.String `tfsdk:"uri"`
}

var DmNamespaceMappingObjectType = map[string]attr.Type{
	"prefix": types.StringType,
	"uri":    types.StringType,
}
var DmNamespaceMappingObjectDefault = map[string]attr.Value{
	"prefix": types.StringNull(),
	"uri":    types.StringNull(),
}

func GetDmNamespaceMappingDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmNamespaceMappingDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"prefix": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The prefix (Prefix:) used to map namespaces that might be encountered in client requests.", "", "").String,
				Computed:            true,
			},
			"uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URI (URI:) used to map namespaces that might be encountered in client requests.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmNamespaceMappingDataSourceSchema
}
func GetDmNamespaceMappingResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmNamespaceMappingResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"prefix": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The prefix (Prefix:) used to map namespaces that might be encountered in client requests.", "", "").String,
				Optional:            true,
			},
			"uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The URI (URI:) used to map namespaces that might be encountered in client requests.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmNamespaceMappingResourceSchema
}

func (data DmNamespaceMapping) IsNull() bool {
	if !data.Prefix.IsNull() {
		return false
	}
	if !data.Uri.IsNull() {
		return false
	}
	return true
}

func (data DmNamespaceMapping) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Prefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Prefix`, data.Prefix.ValueString())
	}
	if !data.Uri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URI`, data.Uri.ValueString())
	}
	return body
}

func (data *DmNamespaceMapping) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Prefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Prefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Prefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uri = types.StringNull()
	}
}

func (data *DmNamespaceMapping) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Prefix`); value.Exists() && !data.Prefix.IsNull() {
		data.Prefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Prefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && !data.Uri.IsNull() {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uri = types.StringNull()
	}
}
