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

type DmGatewayPeeringGroupClusterNode struct {
	Host       types.String `tfsdk:"host"`
	LocalNodes types.String `tfsdk:"local_nodes"`
}

var DmGatewayPeeringGroupClusterNodeObjectType = map[string]attr.Type{
	"host":        types.StringType,
	"local_nodes": types.StringType,
}
var DmGatewayPeeringGroupClusterNodeObjectDefault = map[string]attr.Value{
	"host":        types.StringNull(),
	"local_nodes": types.StringNull(),
}
var DmGatewayPeeringGroupClusterNodeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the node.", "", "").String,
			Computed:            true,
		},
		"local_nodes": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a comma-separated list of the local IP addresses or host aliases of the other nodes that are in the same data center.", "", "").String,
			Computed:            true,
		},
	},
}
var DmGatewayPeeringGroupClusterNodeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the node.", "", "").String,
			Required:            true,
		},
		"local_nodes": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify a comma-separated list of the local IP addresses or host aliases of the other nodes that are in the same data center.", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmGatewayPeeringGroupClusterNode) IsNull() bool {
	if !data.Host.IsNull() {
		return false
	}
	if !data.LocalNodes.IsNull() {
		return false
	}
	return true
}

func (data DmGatewayPeeringGroupClusterNode) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.LocalNodes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalNodes`, data.LocalNodes.ValueString())
	}
	return body
}

func (data *DmGatewayPeeringGroupClusterNode) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalNodes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalNodes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalNodes = types.StringNull()
	}
}

func (data *DmGatewayPeeringGroupClusterNode) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalNodes`); value.Exists() && !data.LocalNodes.IsNull() {
		data.LocalNodes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalNodes = types.StringNull()
	}
}
