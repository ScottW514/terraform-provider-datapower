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

type DmPolicySubjectBitmap struct {
	Service    types.Bool `tfsdk:"service"`
	Endpoint   types.Bool `tfsdk:"endpoint"`
	Operation  types.Bool `tfsdk:"operation"`
	MessageIn  types.Bool `tfsdk:"message_in"`
	MessageOut types.Bool `tfsdk:"message_out"`
}

var DmPolicySubjectBitmapObjectType = map[string]attr.Type{
	"service":     types.BoolType,
	"endpoint":    types.BoolType,
	"operation":   types.BoolType,
	"message_in":  types.BoolType,
	"message_out": types.BoolType,
}
var DmPolicySubjectBitmapObjectDefault = map[string]attr.Value{
	"service":     types.BoolValue(false),
	"endpoint":    types.BoolValue(false),
	"operation":   types.BoolValue(false),
	"message_in":  types.BoolValue(false),
	"message_out": types.BoolValue(false),
}

func GetDmPolicySubjectBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmPolicySubjectBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"service": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Service Subject",
				Computed:            true,
			},
			"endpoint": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Endpoint Subject",
				Computed:            true,
			},
			"operation": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Operation Subject",
				Computed:            true,
			},
			"message_in": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Message Input Subject",
				Computed:            true,
			},
			"message_out": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Message Output Subject",
				Computed:            true,
			},
		},
	}
	DmPolicySubjectBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmPolicySubjectBitmapDataSourceSchema
}
func GetDmPolicySubjectBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmPolicySubjectBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmPolicySubjectBitmapObjectType,
				DmPolicySubjectBitmapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"service": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Service Subject", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"endpoint": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Endpoint Subject", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"operation": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation Subject", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"message_in": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Message Input Subject", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"message_out": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Message Output Subject", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmPolicySubjectBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmPolicySubjectBitmapResourceSchema.Required = true
	} else {
		DmPolicySubjectBitmapResourceSchema.Optional = true
		DmPolicySubjectBitmapResourceSchema.Computed = true
	}
	return DmPolicySubjectBitmapResourceSchema
}

func (data DmPolicySubjectBitmap) IsNull() bool {
	if !data.Service.IsNull() {
		return false
	}
	if !data.Endpoint.IsNull() {
		return false
	}
	if !data.Operation.IsNull() {
		return false
	}
	if !data.MessageIn.IsNull() {
		return false
	}
	if !data.MessageOut.IsNull() {
		return false
	}
	return true
}

func (data DmPolicySubjectBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Service.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Service`, tfutils.StringFromBool(data.Service, ""))
	}
	if !data.Endpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Endpoint`, tfutils.StringFromBool(data.Endpoint, ""))
	}
	if !data.Operation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Operation`, tfutils.StringFromBool(data.Operation, ""))
	}
	if !data.MessageIn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessageIn`, tfutils.StringFromBool(data.MessageIn, ""))
	}
	if !data.MessageOut.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessageOut`, tfutils.StringFromBool(data.MessageOut, ""))
	}
	return body
}

func (data *DmPolicySubjectBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Service`); value.Exists() {
		data.Service = tfutils.BoolFromString(value.String())
	} else {
		data.Service = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() {
		data.Endpoint = tfutils.BoolFromString(value.String())
	} else {
		data.Endpoint = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() {
		data.Operation = tfutils.BoolFromString(value.String())
	} else {
		data.Operation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MessageIn`); value.Exists() {
		data.MessageIn = tfutils.BoolFromString(value.String())
	} else {
		data.MessageIn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MessageOut`); value.Exists() {
		data.MessageOut = tfutils.BoolFromString(value.String())
	} else {
		data.MessageOut = types.BoolNull()
	}
}

func (data *DmPolicySubjectBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Service`); value.Exists() && !data.Service.IsNull() {
		data.Service = tfutils.BoolFromString(value.String())
	} else if data.Service.ValueBool() {
		data.Service = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() && !data.Endpoint.IsNull() {
		data.Endpoint = tfutils.BoolFromString(value.String())
	} else if data.Endpoint.ValueBool() {
		data.Endpoint = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Operation`); value.Exists() && !data.Operation.IsNull() {
		data.Operation = tfutils.BoolFromString(value.String())
	} else if data.Operation.ValueBool() {
		data.Operation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MessageIn`); value.Exists() && !data.MessageIn.IsNull() {
		data.MessageIn = tfutils.BoolFromString(value.String())
	} else if data.MessageIn.ValueBool() {
		data.MessageIn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MessageOut`); value.Exists() && !data.MessageOut.IsNull() {
		data.MessageOut = tfutils.BoolFromString(value.String())
	} else if data.MessageOut.ValueBool() {
		data.MessageOut = types.BoolNull()
	}
}
