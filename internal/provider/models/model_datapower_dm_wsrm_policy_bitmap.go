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

type DmWSRMPolicyBitmap struct {
	Optional                  types.Bool `tfsdk:"optional"`
	SequenceTransportSecurity types.Bool `tfsdk:"sequence_transport_security"`
	InOrder                   types.Bool `tfsdk:"in_order"`
	TwoWay                    types.Bool `tfsdk:"two_way"`
}

var DmWSRMPolicyBitmapObjectType = map[string]attr.Type{
	"optional":                    types.BoolType,
	"sequence_transport_security": types.BoolType,
	"in_order":                    types.BoolType,
	"two_way":                     types.BoolType,
}
var DmWSRMPolicyBitmapObjectDefault = map[string]attr.Value{
	"optional":                    types.BoolValue(false),
	"sequence_transport_security": types.BoolValue(false),
	"in_order":                    types.BoolValue(false),
	"two_way":                     types.BoolValue(false),
}

func GetDmWSRMPolicyBitmapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmWSRMPolicyBitmapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"optional": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reliable Messaging is optional", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"sequence_transport_security": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RM Sequence must be bound to underlying transport-level security protocol", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"in_order": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RM messages must be delivered in the same order as they have been sent by the source", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"two_way": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response messages require RM Sequence headers", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmWSRMPolicyBitmapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmWSRMPolicyBitmapDataSourceSchema
}
func GetDmWSRMPolicyBitmapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmWSRMPolicyBitmapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmWSRMPolicyBitmapObjectType,
				DmWSRMPolicyBitmapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"optional": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reliable Messaging is optional", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sequence_transport_security": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RM Sequence must be bound to underlying transport-level security protocol", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"in_order": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("RM messages must be delivered in the same order as they have been sent by the source", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"two_way": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response messages require RM Sequence headers", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmWSRMPolicyBitmapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmWSRMPolicyBitmapResourceSchema.Required = true
	} else {
		DmWSRMPolicyBitmapResourceSchema.Optional = true
		DmWSRMPolicyBitmapResourceSchema.Computed = true
	}
	return DmWSRMPolicyBitmapResourceSchema
}

func (data DmWSRMPolicyBitmap) IsNull() bool {
	if !data.Optional.IsNull() {
		return false
	}
	if !data.SequenceTransportSecurity.IsNull() {
		return false
	}
	if !data.InOrder.IsNull() {
		return false
	}
	if !data.TwoWay.IsNull() {
		return false
	}
	return true
}

func (data DmWSRMPolicyBitmap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Optional.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Optional`, tfutils.StringFromBool(data.Optional, ""))
	}
	if !data.SequenceTransportSecurity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SequenceTransportSecurity`, tfutils.StringFromBool(data.SequenceTransportSecurity, ""))
	}
	if !data.InOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InOrder`, tfutils.StringFromBool(data.InOrder, ""))
	}
	if !data.TwoWay.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TwoWay`, tfutils.StringFromBool(data.TwoWay, ""))
	}
	return body
}

func (data *DmWSRMPolicyBitmap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Optional`); value.Exists() {
		data.Optional = tfutils.BoolFromString(value.String())
	} else {
		data.Optional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SequenceTransportSecurity`); value.Exists() {
		data.SequenceTransportSecurity = tfutils.BoolFromString(value.String())
	} else {
		data.SequenceTransportSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InOrder`); value.Exists() {
		data.InOrder = tfutils.BoolFromString(value.String())
	} else {
		data.InOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TwoWay`); value.Exists() {
		data.TwoWay = tfutils.BoolFromString(value.String())
	} else {
		data.TwoWay = types.BoolNull()
	}
}

func (data *DmWSRMPolicyBitmap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Optional`); value.Exists() && !data.Optional.IsNull() {
		data.Optional = tfutils.BoolFromString(value.String())
	} else if data.Optional.ValueBool() {
		data.Optional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SequenceTransportSecurity`); value.Exists() && !data.SequenceTransportSecurity.IsNull() {
		data.SequenceTransportSecurity = tfutils.BoolFromString(value.String())
	} else if data.SequenceTransportSecurity.ValueBool() {
		data.SequenceTransportSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InOrder`); value.Exists() && !data.InOrder.IsNull() {
		data.InOrder = tfutils.BoolFromString(value.String())
	} else if data.InOrder.ValueBool() {
		data.InOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TwoWay`); value.Exists() && !data.TwoWay.IsNull() {
		data.TwoWay = tfutils.BoolFromString(value.String())
	} else if data.TwoWay.ValueBool() {
		data.TwoWay = types.BoolNull()
	}
}
