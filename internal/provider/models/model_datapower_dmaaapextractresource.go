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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPExtractResource struct {
	ErBitmap   *DmAAAPERBitmap `tfsdk:"er_bitmap"`
	ErxPath    types.String    `tfsdk:"erx_path"`
	ErMetadata types.String    `tfsdk:"er_metadata"`
}

var DmAAAPExtractResourceObjectType = map[string]attr.Type{
	"er_bitmap":   types.ObjectType{AttrTypes: DmAAAPERBitmapObjectType},
	"erx_path":    types.StringType,
	"er_metadata": types.StringType,
}
var DmAAAPExtractResourceObjectDefault = map[string]attr.Value{
	"er_bitmap":   types.ObjectValueMust(DmAAAPERBitmapObjectType, DmAAAPERBitmapObjectDefault),
	"erx_path":    types.StringNull(),
	"er_metadata": types.StringNull(),
}
var DmAAAPExtractResourceDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"er_bitmap": GetDmAAAPERBitmapDataSourceSchema("Specify the methods to extract resource.", "method", ""),
		"erx_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the incoming message.", "xpath", "").String,
			Computed:            true,
		},
		"er_metadata": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the configuration for processing metadata.", "metadata", "processingmetadata").String,
			Computed:            true,
		},
	},
}
var DmAAAPExtractResourceResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPExtractResourceObjectType,
			DmAAAPExtractResourceObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"er_bitmap": GetDmAAAPERBitmapResourceSchema("Specify the methods to extract resource.", "method", "", false),
		"erx_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the incoming message.", "xpath", "").String,
			Optional:            true,
		},
		"er_metadata": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the configuration for processing metadata.", "metadata", "processingmetadata").String,
			Optional:            true,
		},
	},
}

func (data DmAAAPExtractResource) IsNull() bool {
	if data.ErBitmap != nil {
		if !data.ErBitmap.IsNull() {
			return false
		}
	}
	if !data.ErxPath.IsNull() {
		return false
	}
	if !data.ErMetadata.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPExtractResourceDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPExtractResourceDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPExtractResourceDataSourceSchema
}

func GetDmAAAPExtractResourceResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPExtractResourceResourceSchema.Required = true
	} else {
		DmAAAPExtractResourceResourceSchema.Optional = true
		DmAAAPExtractResourceResourceSchema.Computed = true
	}
	DmAAAPExtractResourceResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPExtractResourceResourceSchema
}

func (data DmAAAPExtractResource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if data.ErBitmap != nil {
		if !data.ErBitmap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ERBitmap`, data.ErBitmap.ToBody(ctx, ""))
		}
	}
	if !data.ErxPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ERXPath`, data.ErxPath.ValueString())
	}
	if !data.ErMetadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ERMetadata`, data.ErMetadata.ValueString())
	}
	return body
}

func (data *DmAAAPExtractResource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ERBitmap`); value.Exists() {
		data.ErBitmap = &DmAAAPERBitmap{}
		data.ErBitmap.FromBody(ctx, "", value)
	} else {
		data.ErBitmap = nil
	}
	if value := res.Get(pathRoot + `ERXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ERMetadata`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErMetadata = types.StringNull()
	}
}

func (data *DmAAAPExtractResource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ERBitmap`); value.Exists() {
		data.ErBitmap.UpdateFromBody(ctx, "", value)
	} else {
		data.ErBitmap = nil
	}
	if value := res.Get(pathRoot + `ERXPath`); value.Exists() && !data.ErxPath.IsNull() {
		data.ErxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ERMetadata`); value.Exists() && !data.ErMetadata.IsNull() {
		data.ErMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErMetadata = types.StringNull()
	}
}
