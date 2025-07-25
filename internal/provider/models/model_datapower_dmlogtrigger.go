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

type DmLogTrigger struct {
	MessageId      types.String `tfsdk:"message_id"`
	Expression     types.String `tfsdk:"expression"`
	OnlyOnce       types.Bool   `tfsdk:"only_once"`
	StopProcessing types.Bool   `tfsdk:"stop_processing"`
	Command        types.String `tfsdk:"command"`
}

var DmLogTriggerObjectType = map[string]attr.Type{
	"message_id":      types.StringType,
	"expression":      types.StringType,
	"only_once":       types.BoolType,
	"stop_processing": types.BoolType,
	"command":         types.StringType,
}
var DmLogTriggerObjectDefault = map[string]attr.Value{
	"message_id":      types.StringNull(),
	"expression":      types.StringNull(),
	"only_once":       types.BoolValue(true),
	"stop_processing": types.BoolValue(true),
	"command":         types.StringNull(),
}
var DmLogTriggerDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"message_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Message ID", "", "").String,
			Computed:            true,
		},
		"expression": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Regular expression", "", "").String,
			Computed:            true,
		},
		"only_once": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Only one time", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"stop_processing": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Only this trigger", "", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"command": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Command", "", "").String,
			Computed:            true,
		},
	},
}
var DmLogTriggerResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"message_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Message ID", "", "").String,
			Optional:            true,
		},
		"expression": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Regular expression", "", "").String,
			Optional:            true,
		},
		"only_once": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Only one time", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"stop_processing": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Only this trigger", "", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"command": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Command", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmLogTrigger) IsNull() bool {
	if !data.MessageId.IsNull() {
		return false
	}
	if !data.Expression.IsNull() {
		return false
	}
	if !data.OnlyOnce.IsNull() {
		return false
	}
	if !data.StopProcessing.IsNull() {
		return false
	}
	if !data.Command.IsNull() {
		return false
	}
	return true
}

func (data DmLogTrigger) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.MessageId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ID`, data.MessageId.ValueString())
	}
	if !data.Expression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Expression`, data.Expression.ValueString())
	}
	if !data.OnlyOnce.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OnlyOnce`, tfutils.StringFromBool(data.OnlyOnce, ""))
	}
	if !data.StopProcessing.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StopProcessing`, tfutils.StringFromBool(data.StopProcessing, ""))
	}
	if !data.Command.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Command`, data.Command.ValueString())
	}
	return body
}

func (data *DmLogTrigger) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MessageId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Expression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Expression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Expression = types.StringNull()
	}
	if value := res.Get(pathRoot + `OnlyOnce`); value.Exists() {
		data.OnlyOnce = tfutils.BoolFromString(value.String())
	} else {
		data.OnlyOnce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StopProcessing`); value.Exists() {
		data.StopProcessing = tfutils.BoolFromString(value.String())
	} else {
		data.StopProcessing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Command`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Command = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Command = types.StringNull()
	}
}

func (data *DmLogTrigger) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ID`); value.Exists() && !data.MessageId.IsNull() {
		data.MessageId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Expression`); value.Exists() && !data.Expression.IsNull() {
		data.Expression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Expression = types.StringNull()
	}
	if value := res.Get(pathRoot + `OnlyOnce`); value.Exists() && !data.OnlyOnce.IsNull() {
		data.OnlyOnce = tfutils.BoolFromString(value.String())
	} else if !data.OnlyOnce.ValueBool() {
		data.OnlyOnce = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StopProcessing`); value.Exists() && !data.StopProcessing.IsNull() {
		data.StopProcessing = tfutils.BoolFromString(value.String())
	} else if !data.StopProcessing.ValueBool() {
		data.StopProcessing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Command`); value.Exists() && !data.Command.IsNull() {
		data.Command = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Command = types.StringNull()
	}
}
