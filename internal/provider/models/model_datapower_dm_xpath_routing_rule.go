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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmXPathRoutingRule struct {
	XPath types.String `tfsdk:"x_path"`
	Host  types.String `tfsdk:"host"`
	Port  types.Int64  `tfsdk:"port"`
	Ssl   types.Bool   `tfsdk:"ssl"`
}

var DmXPathRoutingRuleObjectType = map[string]attr.Type{
	"x_path": types.StringType,
	"host":   types.StringType,
	"port":   types.Int64Type,
	"ssl":    types.BoolType,
}
var DmXPathRoutingRuleObjectDefault = map[string]attr.Value{
	"x_path": types.StringNull(),
	"host":   types.StringNull(),
	"port":   types.Int64Null(),
	"ssl":    types.BoolValue(false),
}

func GetDmXPathRoutingRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmXPathRoutingRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"x_path": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The XPath expression applied to submitted documents. This expression evaluates to true or false. If the expression points to a particular node and that node is present in the submitted document, the expression evaluates to true.</p><p>This expression cannot exceed 330 characters. Use the Namespace Mapping tab to establish mapping that then allow the use of qualified names in the XPath expression, shortening the expression.</p>", "", "").String,
				Computed:            true,
			},
			"host": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address to which matching documents should be routed.", "", "").String,
				Computed:            true,
			},
			"port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port to which matching documents should be routed.", "", "").String,
				Computed:            true,
			},
			"ssl": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the connection to the target destination uses TLS communications. The default is off. When set to on, the DataPower Gateway uses the TLS profile that is specified at the service level to establish TLS communications to the destination host.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	return DmXPathRoutingRuleDataSourceSchema
}
func GetDmXPathRoutingRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmXPathRoutingRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"x_path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The XPath expression applied to submitted documents. This expression evaluates to true or false. If the expression points to a particular node and that node is present in the submitted document, the expression evaluates to true.</p><p>This expression cannot exceed 330 characters. Use the Namespace Mapping tab to establish mapping that then allow the use of qualified names in the XPath expression, shortening the expression.</p>", "", "").String,
				Required:            true,
			},
			"host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address to which matching documents should be routed.", "", "").String,
				Required:            true,
			},
			"port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port to which matching documents should be routed.", "", "").String,
				Required:            true,
			},
			"ssl": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the connection to the target destination uses TLS communications. The default is off. When set to on, the DataPower Gateway uses the TLS profile that is specified at the service level to establish TLS communications to the destination host.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	return DmXPathRoutingRuleResourceSchema
}

func (data DmXPathRoutingRule) IsNull() bool {
	if !data.XPath.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Ssl.IsNull() {
		return false
	}
	return true
}

func (data DmXPathRoutingRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.XPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.XPath.ValueString())
	}
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.Ssl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSL`, tfutils.StringFromBool(data.Ssl, ""))
	}
	return body
}

func (data *DmXPathRoutingRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSL`); value.Exists() {
		data.Ssl = tfutils.BoolFromString(value.String())
	} else {
		data.Ssl = types.BoolNull()
	}
}

func (data *DmXPathRoutingRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.XPath.IsNull() {
		data.XPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSL`); value.Exists() && !data.Ssl.IsNull() {
		data.Ssl = tfutils.BoolFromString(value.String())
	} else if data.Ssl.ValueBool() {
		data.Ssl = types.BoolNull()
	}
}
