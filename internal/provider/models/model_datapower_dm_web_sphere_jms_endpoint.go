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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWebSphereJMSEndpoint struct {
	Host      types.String `tfsdk:"host"`
	Port      types.Int64  `tfsdk:"port"`
	Transport types.String `tfsdk:"transport"`
}

var DmWebSphereJMSEndpointObjectType = map[string]attr.Type{
	"host":      types.StringType,
	"port":      types.Int64Type,
	"transport": types.StringType,
}
var DmWebSphereJMSEndpointObjectDefault = map[string]attr.Value{
	"host":      types.StringNull(),
	"port":      types.Int64Null(),
	"transport": types.StringValue("TCP"),
}

func GetDmWebSphereJMSEndpointDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWebSphereJMSEndpointDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"host": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the host name or IP address of a bootstrap server.",
				Computed:            true,
			},
			"port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the listening port number of a bootstrap server.",
				Computed:            true,
			},
			"transport": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the predefined transport chain for the bootstrap server. This connection is used to exchange information between the DataPower Gateway and the bootstrap server. By default, uses the predefined <tt>BootstrapBasicMessaging</tt> transport chain.",
				Computed:            true,
			},
		},
	}
	return DmWebSphereJMSEndpointDataSourceSchema
}
func GetDmWebSphereJMSEndpointResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWebSphereJMSEndpointResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the host name or IP address of a bootstrap server.", "", "").String,
				Required:            true,
			},
			"port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port number of a bootstrap server.", "", "").String,
				Required:            true,
			},
			"transport": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the predefined transport chain for the bootstrap server. This connection is used to exchange information between the DataPower Gateway and the bootstrap server. By default, uses the predefined <tt>BootstrapBasicMessaging</tt> transport chain.", "", "").AddStringEnum("TCP", "SSL", "HTTP", "HTTPS").AddDefaultValue("TCP").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TCP", "SSL", "HTTP", "HTTPS"),
				},
				Default: stringdefault.StaticString("TCP"),
			},
		},
	}
	return DmWebSphereJMSEndpointResourceSchema
}

func (data DmWebSphereJMSEndpoint) IsNull() bool {
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Transport.IsNull() {
		return false
	}
	return true
}

func (data DmWebSphereJMSEndpoint) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.Transport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transport`, data.Transport.ValueString())
	}
	return body
}

func (data *DmWebSphereJMSEndpoint) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
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
	if value := res.Get(pathRoot + `Transport`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transport = types.StringValue("TCP")
	}
}

func (data *DmWebSphereJMSEndpoint) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
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
	if value := res.Get(pathRoot + `Transport`); value.Exists() && !data.Transport.IsNull() {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else if data.Transport.ValueString() != "TCP" {
		data.Transport = types.StringNull()
	}
}
