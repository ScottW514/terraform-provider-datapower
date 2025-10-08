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

type DmStaticHost struct {
	Hostname    types.String `tfsdk:"hostname"`
	IpAddress   types.String `tfsdk:"ip_address"`
	UserSummary types.String `tfsdk:"user_summary"`
}

var DmStaticHostObjectType = map[string]attr.Type{
	"hostname":     types.StringType,
	"ip_address":   types.StringType,
	"user_summary": types.StringType,
}
var DmStaticHostObjectDefault = map[string]attr.Value{
	"hostname":     types.StringNull(),
	"ip_address":   types.StringNull(),
	"user_summary": types.StringNull(),
}

func GetDmStaticHostDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmStaticHostDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"hostname": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the hostname of the target host.",
				Computed:            true,
			},
			"ip_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the IP address for the target host.",
				Computed:            true,
			},
			"user_summary": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a brief, descriptive summary for the configuration.",
				Computed:            true,
			},
		},
	}
	return DmStaticHostDataSourceSchema
}
func GetDmStaticHostResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmStaticHostResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"hostname": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname of the target host.", "", "").String,
				Required:            true,
			},
			"ip_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address for the target host.", "", "").String,
				Required:            true,
			},
			"user_summary": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a brief, descriptive summary for the configuration.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmStaticHostResourceSchema
}

func (data DmStaticHost) IsNull() bool {
	if !data.Hostname.IsNull() {
		return false
	}
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data DmStaticHost) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Hostname`, data.Hostname.ValueString())
	}
	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPAddress`, data.IpAddress.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *DmStaticHost) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *DmStaticHost) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
