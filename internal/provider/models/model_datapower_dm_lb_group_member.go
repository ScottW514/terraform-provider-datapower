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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmLBGroupMember struct {
	Server        types.String `tfsdk:"server"`
	Weight        types.Int64  `tfsdk:"weight"`
	MappedPort    types.Int64  `tfsdk:"mapped_port"`
	Activity      types.String `tfsdk:"activity"`
	HealthPort    types.Int64  `tfsdk:"health_port"`
	LbMemberState types.String `tfsdk:"lb_member_state"`
}

var DmLBGroupMemberObjectType = map[string]attr.Type{
	"server":          types.StringType,
	"weight":          types.Int64Type,
	"mapped_port":     types.Int64Type,
	"activity":        types.StringType,
	"health_port":     types.Int64Type,
	"lb_member_state": types.StringType,
}
var DmLBGroupMemberObjectDefault = map[string]attr.Value{
	"server":          types.StringNull(),
	"weight":          types.Int64Value(1),
	"mapped_port":     types.Int64Null(),
	"activity":        types.StringNull(),
	"health_port":     types.Int64Null(),
	"lb_member_state": types.StringValue("enabled"),
}
var DmLBGroupMemberDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"server": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the host name or IP address of the real server.", "server", "").String,
			Computed:            true,
		},
		"weight": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>When balancing with Weighted Round Robin, the weight property indicates the relative priority of this server versus the other members of the group. The greater the weight indicates the higher the preference. Use an integer between 1 and 65000.</p><p>For example, assume three server pool members, A, B, and C, with assigned weights of 5, 3, and 2 respectively. As a result of this weighting, server A can expect to receive approximately 50% of client requests, server B can expect to receive approximately 30% of client requests, with server C receiving the remaining 20%.</p>", "weight", "").AddIntegerRange(1, 65000).AddDefaultValue("1").String,
			Computed:            true,
		},
		"mapped_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the port on the real server. If non-zero, the associated real server is contacted on this port. Normally the real server is contacted on the same port number as that of the virtual server and this should not be set to a value other than zero. However, if you have services that run on different ports for different members of the group, you might need to define this value.</p><p>The DataPower Gateway checks this port when the health check type is IMS Connect.</p>", "", "").String,
			Computed:            true,
		},
		"activity": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("DEPRECATED.", "", "").String,
			Computed:            true,
		},
		"health_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the TCP Port number to test.", "", "").String,
			Computed:            true,
		},
		"lb_member_state": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable or Disable the Load Balancer Member. If the Administrative State is disabled, the member are not used in health or forwarding checks.", "admin-state", "").AddStringEnum("enabled", "disabled").AddDefaultValue("enabled").String,
			Computed:            true,
		},
	},
}
var DmLBGroupMemberResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"server": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the host name or IP address of the real server.", "server", "").String,
			Required:            true,
		},
		"weight": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>When balancing with Weighted Round Robin, the weight property indicates the relative priority of this server versus the other members of the group. The greater the weight indicates the higher the preference. Use an integer between 1 and 65000.</p><p>For example, assume three server pool members, A, B, and C, with assigned weights of 5, 3, and 2 respectively. As a result of this weighting, server A can expect to receive approximately 50% of client requests, server B can expect to receive approximately 30% of client requests, with server C receiving the remaining 20%.</p>", "weight", "").AddIntegerRange(1, 65000).AddDefaultValue("1").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 65000),
			},
			Default: int64default.StaticInt64(1),
		},
		"mapped_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies the port on the real server. If non-zero, the associated real server is contacted on this port. Normally the real server is contacted on the same port number as that of the virtual server and this should not be set to a value other than zero. However, if you have services that run on different ports for different members of the group, you might need to define this value.</p><p>The DataPower Gateway checks this port when the health check type is IMS Connect.</p>", "", "").String,
			Optional:            true,
		},
		"activity": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("DEPRECATED.", "", "").String,
			Optional:            true,
		},
		"health_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the TCP Port number to test.", "", "").String,
			Optional:            true,
		},
		"lb_member_state": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enable or Disable the Load Balancer Member. If the Administrative State is disabled, the member are not used in health or forwarding checks.", "admin-state", "").AddStringEnum("enabled", "disabled").AddDefaultValue("enabled").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("enabled", "disabled"),
			},
			Default: stringdefault.StaticString("enabled"),
		},
	},
}

func (data DmLBGroupMember) IsNull() bool {
	if !data.Server.IsNull() {
		return false
	}
	if !data.Weight.IsNull() {
		return false
	}
	if !data.MappedPort.IsNull() {
		return false
	}
	if !data.Activity.IsNull() {
		return false
	}
	if !data.HealthPort.IsNull() {
		return false
	}
	if !data.LbMemberState.IsNull() {
		return false
	}
	return true
}

func (data DmLBGroupMember) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Server.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Server`, data.Server.ValueString())
	}
	if !data.Weight.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Weight`, data.Weight.ValueInt64())
	}
	if !data.MappedPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MappedPort`, data.MappedPort.ValueInt64())
	}
	if !data.Activity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Activity`, data.Activity.ValueString())
	}
	if !data.HealthPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HealthPort`, data.HealthPort.ValueInt64())
	}
	if !data.LbMemberState.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LBMemberState`, data.LbMemberState.ValueString())
	}
	return body
}

func (data *DmLBGroupMember) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `Weight`); value.Exists() {
		data.Weight = types.Int64Value(value.Int())
	} else {
		data.Weight = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `MappedPort`); value.Exists() {
		data.MappedPort = types.Int64Value(value.Int())
	} else {
		data.MappedPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Activity`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Activity = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Activity = types.StringNull()
	}
	if value := res.Get(pathRoot + `HealthPort`); value.Exists() {
		data.HealthPort = types.Int64Value(value.Int())
	} else {
		data.HealthPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LBMemberState`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LbMemberState = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LbMemberState = types.StringValue("enabled")
	}
}

func (data *DmLBGroupMember) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && !data.Server.IsNull() {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `Weight`); value.Exists() && !data.Weight.IsNull() {
		data.Weight = types.Int64Value(value.Int())
	} else if data.Weight.ValueInt64() != 1 {
		data.Weight = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MappedPort`); value.Exists() && !data.MappedPort.IsNull() {
		data.MappedPort = types.Int64Value(value.Int())
	} else {
		data.MappedPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Activity`); value.Exists() && !data.Activity.IsNull() {
		data.Activity = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Activity = types.StringNull()
	}
	if value := res.Get(pathRoot + `HealthPort`); value.Exists() && !data.HealthPort.IsNull() {
		data.HealthPort = types.Int64Value(value.Int())
	} else {
		data.HealthPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LBMemberState`); value.Exists() && !data.LbMemberState.IsNull() {
		data.LbMemberState = tfutils.ParseStringFromGJSON(value)
	} else if data.LbMemberState.ValueString() != "enabled" {
		data.LbMemberState = types.StringNull()
	}
}
