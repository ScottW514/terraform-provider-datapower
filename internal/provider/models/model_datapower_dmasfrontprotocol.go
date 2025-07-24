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

type DmASFrontProtocol struct {
	FrontProtocol types.String `tfsdk:"front_protocol"`
	MdnReceiver   types.Bool   `tfsdk:"mdn_receiver"`
}

var DmASFrontProtocolObjectType = map[string]attr.Type{
	"front_protocol": types.StringType,
	"mdn_receiver":   types.BoolType,
}
var DmASFrontProtocolObjectDefault = map[string]attr.Value{
	"front_protocol": types.StringNull(),
	"mdn_receiver":   types.BoolValue(false),
}
var DmASFrontProtocolDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"front_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Front protocol (reference to front protocol handler)", "front-protocol", "").String,
			Computed:            true,
		},
		"mdn_receiver": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set as MDN receiver", "mdn-receiver", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmASFrontProtocolResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"front_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Front protocol (reference to front protocol handler)", "front-protocol", "").String,
			Required:            true,
		},
		"mdn_receiver": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Set as MDN receiver", "mdn-receiver", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmASFrontProtocol) IsNull() bool {
	if !data.FrontProtocol.IsNull() {
		return false
	}
	if !data.MdnReceiver.IsNull() {
		return false
	}
	return true
}

func (data DmASFrontProtocol) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.FrontProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontProtocol`, data.FrontProtocol.ValueString())
	}
	if !data.MdnReceiver.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MDNReceiver`, tfutils.StringFromBool(data.MdnReceiver, ""))
	}
	return body
}

func (data *DmASFrontProtocol) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `MDNReceiver`); value.Exists() {
		data.MdnReceiver = tfutils.BoolFromString(value.String())
	} else {
		data.MdnReceiver = types.BoolNull()
	}
}

func (data *DmASFrontProtocol) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && !data.FrontProtocol.IsNull() {
		data.FrontProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `MDNReceiver`); value.Exists() && !data.MdnReceiver.IsNull() {
		data.MdnReceiver = tfutils.BoolFromString(value.String())
	} else if data.MdnReceiver.ValueBool() {
		data.MdnReceiver = types.BoolNull()
	}
}
