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

type DmNameServer struct {
	IpAddress  types.String `tfsdk:"ip_address"`
	UdpPort    types.Int64  `tfsdk:"udp_port"`
	TcpPort    types.Int64  `tfsdk:"tcp_port"`
	MaxRetries types.Int64  `tfsdk:"max_retries"`
}

var DmNameServerObjectType = map[string]attr.Type{
	"ip_address":  types.StringType,
	"udp_port":    types.Int64Type,
	"tcp_port":    types.Int64Type,
	"max_retries": types.Int64Type,
}
var DmNameServerObjectDefault = map[string]attr.Value{
	"ip_address":  types.StringNull(),
	"udp_port":    types.Int64Value(53),
	"tcp_port":    types.Int64Value(53),
	"max_retries": types.Int64Value(3),
}

func GetDmNameServerDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmNameServerDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"ip_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the IP address of DNS server.",
				Computed:            true,
			},
			"udp_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the UDP port that the DNS server monitors. The default value is 53. This setting is ignored with the first alive load balancing algorithm.",
				Computed:            true,
			},
			"tcp_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the TCP port that the DNS server monitors. The default value is 53. This setting is ignored with the first alive load balancing algorithm.",
				Computed:            true,
			},
			"max_retries": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the maximum number of times to send a query to the DNS server. By default, an unacknowledged resolution request is attempted 3 times. This setting is ignored with the first alive load balancing algorithm. For the first alive algorithm, define this behavior at the DNS settings level rather than the individual server level.",
				Computed:            true,
			},
		},
	}
	return DmNameServerDataSourceSchema
}
func GetDmNameServerResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmNameServerResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"ip_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address of DNS server.", "", "").String,
				Required:            true,
			},
			"udp_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the UDP port that the DNS server monitors. The default value is 53. This setting is ignored with the first alive load balancing algorithm.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("53").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(53),
			},
			"tcp_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TCP port that the DNS server monitors. The default value is 53. This setting is ignored with the first alive load balancing algorithm.", "", "").AddIntegerRange(1, 65535).AddDefaultValue("53").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(53),
			},
			"max_retries": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of times to send a query to the DNS server. By default, an unacknowledged resolution request is attempted 3 times. This setting is ignored with the first alive load balancing algorithm. For the first alive algorithm, define this behavior at the DNS settings level rather than the individual server level.", "", "").AddDefaultValue("3").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(3),
			},
		},
	}
	return DmNameServerResourceSchema
}

func (data DmNameServer) IsNull() bool {
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.UdpPort.IsNull() {
		return false
	}
	if !data.TcpPort.IsNull() {
		return false
	}
	if !data.MaxRetries.IsNull() {
		return false
	}
	return true
}

func (data DmNameServer) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPAddress`, data.IpAddress.ValueString())
	}
	if !data.UdpPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UDPPort`, data.UdpPort.ValueInt64())
	}
	if !data.TcpPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TCPPort`, data.TcpPort.ValueInt64())
	}
	if !data.MaxRetries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRetries`, data.MaxRetries.ValueInt64())
	}
	return body
}

func (data *DmNameServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `UDPPort`); value.Exists() {
		data.UdpPort = types.Int64Value(value.Int())
	} else {
		data.UdpPort = types.Int64Value(53)
	}
	if value := res.Get(pathRoot + `TCPPort`); value.Exists() {
		data.TcpPort = types.Int64Value(value.Int())
	} else {
		data.TcpPort = types.Int64Value(53)
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else {
		data.MaxRetries = types.Int64Value(3)
	}
}

func (data *DmNameServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `UDPPort`); value.Exists() && !data.UdpPort.IsNull() {
		data.UdpPort = types.Int64Value(value.Int())
	} else if data.UdpPort.ValueInt64() != 53 {
		data.UdpPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TCPPort`); value.Exists() && !data.TcpPort.IsNull() {
		data.TcpPort = types.Int64Value(value.Int())
	} else if data.TcpPort.ValueInt64() != 53 {
		data.TcpPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() && !data.MaxRetries.IsNull() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else if data.MaxRetries.ValueInt64() != 3 {
		data.MaxRetries = types.Int64Null()
	}
}
