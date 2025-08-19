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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmGatewayPeeringGroupPeerNode struct {
	Host     types.String `tfsdk:"host"`
	Priority types.Int64  `tfsdk:"priority"`
}

var DmGatewayPeeringGroupPeerNodeObjectType = map[string]attr.Type{
	"host":     types.StringType,
	"priority": types.Int64Type,
}
var DmGatewayPeeringGroupPeerNodeObjectDefault = map[string]attr.Value{
	"host":     types.StringNull(),
	"priority": types.Int64Value(100),
}
var DmGatewayPeeringGroupPeerNodeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the peer.", "", "").String,
			Computed:            true,
		},
		"priority": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the priority to elect the new primary. When failover occurs, the secondary member with the lowest priority is promoted as the new primary. Enter a value in the range 0 - 255. The default value is 100. A secondary member with a priority of 0 can never be promoted.", "", "").AddIntegerRange(0, 255).AddDefaultValue("100").String,
			Computed:            true,
		},
	},
}
var DmGatewayPeeringGroupPeerNodeResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address or host alias of the peer.", "", "").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.NoneOf([]string{"127.0.0.1", "0.0.0.0", "::", "::1"}...),
			},
		},
		"priority": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the priority to elect the new primary. When failover occurs, the secondary member with the lowest priority is promoted as the new primary. Enter a value in the range 0 - 255. The default value is 100. A secondary member with a priority of 0 can never be promoted.", "", "").AddIntegerRange(0, 255).AddDefaultValue("100").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(0, 255),
			},
			Default: int64default.StaticInt64(100),
		},
	},
}

func (data DmGatewayPeeringGroupPeerNode) IsNull() bool {
	if !data.Host.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	return true
}

func (data DmGatewayPeeringGroupPeerNode) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueInt64())
	}
	return body
}

func (data *DmGatewayPeeringGroupPeerNode) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Value(100)
	}
}

func (data *DmGatewayPeeringGroupPeerNode) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else if data.Priority.ValueInt64() != 100 {
		data.Priority = types.Int64Null()
	}
}
