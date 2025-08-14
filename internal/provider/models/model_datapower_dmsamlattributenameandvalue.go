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

type DmSAMLAttributeNameAndValue struct {
	Uri       types.String `tfsdk:"uri"`
	LocalName types.String `tfsdk:"local_name"`
	Value     types.String `tfsdk:"value"`
}

var DmSAMLAttributeNameAndValueObjectType = map[string]attr.Type{
	"uri":        types.StringType,
	"local_name": types.StringType,
	"value":      types.StringType,
}
var DmSAMLAttributeNameAndValueObjectDefault = map[string]attr.Value{
	"uri":        types.StringNull(),
	"local_name": types.StringNull(),
	"value":      types.StringNull(),
}
var DmSAMLAttributeNameAndValueDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"uri": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the namespace URI for the attribute. The namespace URI must match to a name. If blank, the null namespace is used. For example, <tt>http://www.examples.com</tt> matches a message with the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Computed:            true,
		},
		"local_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the local name of the attribute. For example, <tt>cats</tt> matches a message with the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value for the attribute with the corresponding name. For example, <tt>Winchester</tt> matches the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Computed:            true,
		},
	},
}
var DmSAMLAttributeNameAndValueResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"uri": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the namespace URI for the attribute. The namespace URI must match to a name. If blank, the null namespace is used. For example, <tt>http://www.examples.com</tt> matches a message with the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Optional:            true,
		},
		"local_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the local name of the attribute. For example, <tt>cats</tt> matches a message with the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Optional:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the value for the attribute with the corresponding name. For example, <tt>Winchester</tt> matches the following attribute.</p><pre><tt>&lt;Attribute AttributeName=\"cats\" AttributeNamespace=\"http://www.example.com\"></tt><tt> &lt;AttributeValue>Winchester&lt;/AttributeValue></tt><tt>&lt;Attribute></tt></pre>", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmSAMLAttributeNameAndValue) IsNull() bool {
	if !data.Uri.IsNull() {
		return false
	}
	if !data.LocalName.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmSAMLAttributeNameAndValue) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Uri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URI`, data.Uri.ValueString())
	}
	if !data.LocalName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalName`, data.LocalName.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmSAMLAttributeNameAndValue) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uri = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmSAMLAttributeNameAndValue) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && !data.Uri.IsNull() {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uri = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalName`); value.Exists() && !data.LocalName.IsNull() {
		data.LocalName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
