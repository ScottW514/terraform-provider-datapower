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

type DmODRConnector struct {
	DmgrHostname types.String `tfsdk:"dmgr_hostname"`
	DmgrPort     types.Int64  `tfsdk:"dmgr_port"`
}

var DmODRConnectorObjectType = map[string]attr.Type{
	"dmgr_hostname": types.StringType,
	"dmgr_port":     types.Int64Type,
}
var DmODRConnectorObjectDefault = map[string]attr.Value{
	"dmgr_hostname": types.StringNull(),
	"dmgr_port":     types.Int64Null(),
}
var DmODRConnectorDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"dmgr_hostname": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The host name of the server where the Intelligent management service is enabled.", "dmgr-host", "").String,
			Computed:            true,
		},
		"dmgr_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The port on the server where the Intelligent management service is enabled.", "dmgr-port", "").String,
			Computed:            true,
		},
	},
}
var DmODRConnectorResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"dmgr_hostname": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The host name of the server where the Intelligent management service is enabled.", "dmgr-host", "").String,
			Required:            true,
		},
		"dmgr_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The port on the server where the Intelligent management service is enabled.", "dmgr-port", "").String,
			Required:            true,
		},
	},
}

func (data DmODRConnector) IsNull() bool {
	if !data.DmgrHostname.IsNull() {
		return false
	}
	if !data.DmgrPort.IsNull() {
		return false
	}
	return true
}

func (data DmODRConnector) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.DmgrHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DmgrHostname`, data.DmgrHostname.ValueString())
	}
	if !data.DmgrPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DmgrPort`, data.DmgrPort.ValueInt64())
	}
	return body
}

func (data *DmODRConnector) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `DmgrHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DmgrHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DmgrHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `DmgrPort`); value.Exists() {
		data.DmgrPort = types.Int64Value(value.Int())
	} else {
		data.DmgrPort = types.Int64Null()
	}
}

func (data *DmODRConnector) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `DmgrHostname`); value.Exists() && !data.DmgrHostname.IsNull() {
		data.DmgrHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DmgrHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `DmgrPort`); value.Exists() && !data.DmgrPort.IsNull() {
		data.DmgrPort = types.Int64Value(value.Int())
	} else {
		data.DmgrPort = types.Int64Null()
	}
}
