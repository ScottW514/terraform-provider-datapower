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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmGatewayPeeringClusterNode struct {
	Address   types.String `tfsdk:"address"`
	Port      types.Int64  `tfsdk:"port"`
	LocalNode types.Bool   `tfsdk:"local_node"`
}

var DmGatewayPeeringClusterNodeObjectType = map[string]attr.Type{
	"address":    types.StringType,
	"port":       types.Int64Type,
	"local_node": types.BoolType,
}
var DmGatewayPeeringClusterNodeObjectDefault = map[string]attr.Value{
	"address":    types.StringNull(),
	"port":       types.Int64Null(),
	"local_node": types.BoolValue(true),
}
var DmGatewayPeeringClusterNodeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"address": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the node.", "", "").String,
			Computed:            true,
		},
		"port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local port of the node.", "", "").String,
			Computed:            true,
		},
		"local_node": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the node is local to the data center.", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
	},
}
var DmGatewayPeeringClusterNodeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"address": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the node.", "", "").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.NoneOf([]string{"127.0.0.1", "0.0.0.0", "::", "::1"}...),
			},
		},
		"port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local port of the node.", "", "").String,
			Required:            true,
		},
		"local_node": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the node is local to the data center.", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
	},
}

func (data DmGatewayPeeringClusterNode) IsNull() bool {
	if !data.Address.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.LocalNode.IsNull() {
		return false
	}
	return true
}

func (data DmGatewayPeeringClusterNode) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Address`, data.Address.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.LocalNode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalNode`, tfutils.StringFromBool(data.LocalNode, ""))
	}
	return body
}

func (data *DmGatewayPeeringClusterNode) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Address`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Address = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LocalNode`); value.Exists() {
		data.LocalNode = tfutils.BoolFromString(value.String())
	} else {
		data.LocalNode = types.BoolNull()
	}
}

func (data *DmGatewayPeeringClusterNode) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Address`); value.Exists() && !data.Address.IsNull() {
		data.Address = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LocalNode`); value.Exists() && !data.LocalNode.IsNull() {
		data.LocalNode = tfutils.BoolFromString(value.String())
	} else if !data.LocalNode.ValueBool() {
		data.LocalNode = types.BoolNull()
	}
}
