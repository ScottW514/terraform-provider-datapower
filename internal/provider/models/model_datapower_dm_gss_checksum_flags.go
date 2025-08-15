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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmGssChecksumFlags struct {
	Replay   types.Bool `tfsdk:"replay"`
	Sequence types.Bool `tfsdk:"sequence"`
	Conf     types.Bool `tfsdk:"conf"`
	Integ    types.Bool `tfsdk:"integ"`
}

var DmGssChecksumFlagsObjectType = map[string]attr.Type{
	"replay":   types.BoolType,
	"sequence": types.BoolType,
	"conf":     types.BoolType,
	"integ":    types.BoolType,
}
var DmGssChecksumFlagsObjectDefault = map[string]attr.Value{
	"replay":   types.BoolValue(false),
	"sequence": types.BoolValue(false),
	"conf":     types.BoolValue(false),
	"integ":    types.BoolValue(false),
}
var DmGssChecksumFlagsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"replay": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("REPLAY", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"sequence": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SEQUENCE", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"conf": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CONF", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"integ": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("INTEG", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmGssChecksumFlagsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmGssChecksumFlagsObjectType,
			DmGssChecksumFlagsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"replay": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("REPLAY", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"sequence": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("SEQUENCE", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"conf": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("CONF", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"integ": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("INTEG", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmGssChecksumFlags) IsNull() bool {
	if !data.Replay.IsNull() {
		return false
	}
	if !data.Sequence.IsNull() {
		return false
	}
	if !data.Conf.IsNull() {
		return false
	}
	if !data.Integ.IsNull() {
		return false
	}
	return true
}
func GetDmGssChecksumFlagsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmGssChecksumFlagsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmGssChecksumFlagsDataSourceSchema
}

func GetDmGssChecksumFlagsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmGssChecksumFlagsResourceSchema.Required = true
	} else {
		DmGssChecksumFlagsResourceSchema.Optional = true
		DmGssChecksumFlagsResourceSchema.Computed = true
	}
	DmGssChecksumFlagsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmGssChecksumFlagsResourceSchema
}

func (data DmGssChecksumFlags) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Replay.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`REPLAY`, tfutils.StringFromBool(data.Replay, ""))
	}
	if !data.Sequence.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SEQUENCE`, tfutils.StringFromBool(data.Sequence, ""))
	}
	if !data.Conf.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CONF`, tfutils.StringFromBool(data.Conf, ""))
	}
	if !data.Integ.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`INTEG`, tfutils.StringFromBool(data.Integ, ""))
	}
	return body
}

func (data *DmGssChecksumFlags) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `REPLAY`); value.Exists() {
		data.Replay = tfutils.BoolFromString(value.String())
	} else {
		data.Replay = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SEQUENCE`); value.Exists() {
		data.Sequence = tfutils.BoolFromString(value.String())
	} else {
		data.Sequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CONF`); value.Exists() {
		data.Conf = tfutils.BoolFromString(value.String())
	} else {
		data.Conf = types.BoolNull()
	}
	if value := res.Get(pathRoot + `INTEG`); value.Exists() {
		data.Integ = tfutils.BoolFromString(value.String())
	} else {
		data.Integ = types.BoolNull()
	}
}

func (data *DmGssChecksumFlags) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `REPLAY`); value.Exists() && !data.Replay.IsNull() {
		data.Replay = tfutils.BoolFromString(value.String())
	} else if data.Replay.ValueBool() {
		data.Replay = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SEQUENCE`); value.Exists() && !data.Sequence.IsNull() {
		data.Sequence = tfutils.BoolFromString(value.String())
	} else if data.Sequence.ValueBool() {
		data.Sequence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CONF`); value.Exists() && !data.Conf.IsNull() {
		data.Conf = tfutils.BoolFromString(value.String())
	} else if data.Conf.ValueBool() {
		data.Conf = types.BoolNull()
	}
	if value := res.Get(pathRoot + `INTEG`); value.Exists() && !data.Integ.IsNull() {
		data.Integ = tfutils.BoolFromString(value.String())
	} else if data.Integ.ValueBool() {
		data.Integ = types.BoolNull()
	}
}
