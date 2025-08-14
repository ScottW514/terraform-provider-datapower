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

type DmDynamicStylePolicyActionBaseReference struct {
	Url     types.String `tfsdk:"url"`
	Literal types.String `tfsdk:"literal"`
	Default types.String `tfsdk:"default"`
}

var DmDynamicStylePolicyActionBaseReferenceObjectType = map[string]attr.Type{
	"url":     types.StringType,
	"literal": types.StringType,
	"default": types.StringType,
}
var DmDynamicStylePolicyActionBaseReferenceObjectDefault = map[string]attr.Value{
	"url":     types.StringNull(),
	"literal": types.StringNull(),
	"default": types.StringNull(),
}
var DmDynamicStylePolicyActionBaseReferenceDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies a URL to a file that contains serialized XML or JSON properties to be merged into the dynamic object. These properties override any existing literal or default properties. The URL can contain variable references, and fields within the associated file can also contain variable references.", "url", "").String,
			Computed:            true,
		},
		"literal": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies a literal string that defines serialized XML or JSON properties for merging into the dynamic object. These properties override the existing default properties. The literal string can contain variable references.", "literal", "").String,
			Computed:            true,
		},
		"default": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies an existing object from which to retrieve default property values for the dynamic object. If an object is not specified, then the URL reference, the literal configuration, or the combination of URL reference and literal configuration fully define the action.", "default", "").String,
			Computed:            true,
		},
	},
}
var DmDynamicStylePolicyActionBaseReferenceResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies a URL to a file that contains serialized XML or JSON properties to be merged into the dynamic object. These properties override any existing literal or default properties. The URL can contain variable references, and fields within the associated file can also contain variable references.", "url", "").String,
			Optional:            true,
		},
		"literal": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies a literal string that defines serialized XML or JSON properties for merging into the dynamic object. These properties override the existing default properties. The literal string can contain variable references.", "literal", "").String,
			Optional:            true,
		},
		"default": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies an existing object from which to retrieve default property values for the dynamic object. If an object is not specified, then the URL reference, the literal configuration, or the combination of URL reference and literal configuration fully define the action.", "default", "").String,
			Optional:            true,
		},
	},
}

func (data DmDynamicStylePolicyActionBaseReference) IsNull() bool {
	if !data.Url.IsNull() {
		return false
	}
	if !data.Literal.IsNull() {
		return false
	}
	if !data.Default.IsNull() {
		return false
	}
	return true
}

func (data DmDynamicStylePolicyActionBaseReference) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.Literal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Literal`, data.Literal.ValueString())
	}
	if !data.Default.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Default`, data.Default.ValueString())
	}
	return body
}

func (data *DmDynamicStylePolicyActionBaseReference) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `Literal`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Literal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Literal = types.StringNull()
	}
	if value := res.Get(pathRoot + `Default`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Default = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Default = types.StringNull()
	}
}

func (data *DmDynamicStylePolicyActionBaseReference) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `Literal`); value.Exists() && !data.Literal.IsNull() {
		data.Literal = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Literal = types.StringNull()
	}
	if value := res.Get(pathRoot + `Default`); value.Exists() && !data.Default.IsNull() {
		data.Default = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Default = types.StringNull()
	}
}
