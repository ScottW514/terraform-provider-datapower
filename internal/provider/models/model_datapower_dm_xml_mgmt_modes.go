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

type DmXMLMgmtModes struct {
	Any              types.Bool `tfsdk:"any"`
	Soma             types.Bool `tfsdk:"soma"`
	V2004            types.Bool `tfsdk:"v2004"`
	Amp              types.Bool `tfsdk:"amp"`
	Slm              types.Bool `tfsdk:"slm"`
	Wsm              types.Bool `tfsdk:"wsm"`
	Wsdm             types.Bool `tfsdk:"wsdm"`
	WsrrSubscription types.Bool `tfsdk:"wsrr_subscription"`
}

var DmXMLMgmtModesObjectType = map[string]attr.Type{
	"any":               types.BoolType,
	"soma":              types.BoolType,
	"v2004":             types.BoolType,
	"amp":               types.BoolType,
	"slm":               types.BoolType,
	"wsm":               types.BoolType,
	"wsdm":              types.BoolType,
	"wsrr_subscription": types.BoolType,
}
var DmXMLMgmtModesObjectDefault = map[string]attr.Value{
	"any":               types.BoolValue(true),
	"soma":              types.BoolValue(true),
	"v2004":             types.BoolValue(true),
	"amp":               types.BoolValue(true),
	"slm":               types.BoolValue(true),
	"wsm":               types.BoolValue(false),
	"wsdm":              types.BoolValue(false),
	"wsrr_subscription": types.BoolValue(true),
}

func GetDmXMLMgmtModesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmXMLMgmtModesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"any": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP management URI", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"soma": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP configuration management", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"v2004": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP configuration management (v2004)", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"amp": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AMP endpoint", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"slm": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SLM endpoint", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"wsm": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Management endpoint", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"wsdm": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDM endpoint", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"wsrr_subscription": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR subscription", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmXMLMgmtModesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmXMLMgmtModesDataSourceSchema
}
func GetDmXMLMgmtModesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmXMLMgmtModesResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmXMLMgmtModesObjectType,
				DmXMLMgmtModesObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"any": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP management URI", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"soma": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP configuration management", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"v2004": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP configuration management (v2004)", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"amp": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AMP endpoint", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"slm": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SLM endpoint", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"wsm": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Management endpoint", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsdm": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSDM endpoint", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsrr_subscription": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR subscription", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmXMLMgmtModesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmXMLMgmtModesResourceSchema.Required = true
	} else {
		DmXMLMgmtModesResourceSchema.Optional = true
		DmXMLMgmtModesResourceSchema.Computed = true
	}
	return DmXMLMgmtModesResourceSchema
}

func (data DmXMLMgmtModes) IsNull() bool {
	if !data.Any.IsNull() {
		return false
	}
	if !data.Soma.IsNull() {
		return false
	}
	if !data.V2004.IsNull() {
		return false
	}
	if !data.Amp.IsNull() {
		return false
	}
	if !data.Slm.IsNull() {
		return false
	}
	if !data.Wsm.IsNull() {
		return false
	}
	if !data.Wsdm.IsNull() {
		return false
	}
	if !data.WsrrSubscription.IsNull() {
		return false
	}
	return true
}

func (data DmXMLMgmtModes) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Any.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`any`, tfutils.StringFromBool(data.Any, ""))
	}
	if !data.Soma.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`soma`, tfutils.StringFromBool(data.Soma, ""))
	}
	if !data.V2004.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`v2004`, tfutils.StringFromBool(data.V2004, ""))
	}
	if !data.Amp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`amp`, tfutils.StringFromBool(data.Amp, ""))
	}
	if !data.Slm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`slm`, tfutils.StringFromBool(data.Slm, ""))
	}
	if !data.Wsm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wsm`, tfutils.StringFromBool(data.Wsm, ""))
	}
	if !data.Wsdm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wsdm`, tfutils.StringFromBool(data.Wsdm, ""))
	}
	if !data.WsrrSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`wsrr-subscription`, tfutils.StringFromBool(data.WsrrSubscription, ""))
	}
	return body
}

func (data *DmXMLMgmtModes) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `any`); value.Exists() {
		data.Any = tfutils.BoolFromString(value.String())
	} else {
		data.Any = types.BoolNull()
	}
	if value := res.Get(pathRoot + `soma`); value.Exists() {
		data.Soma = tfutils.BoolFromString(value.String())
	} else {
		data.Soma = types.BoolNull()
	}
	if value := res.Get(pathRoot + `v2004`); value.Exists() {
		data.V2004 = tfutils.BoolFromString(value.String())
	} else {
		data.V2004 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `amp`); value.Exists() {
		data.Amp = tfutils.BoolFromString(value.String())
	} else {
		data.Amp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `slm`); value.Exists() {
		data.Slm = tfutils.BoolFromString(value.String())
	} else {
		data.Slm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsm`); value.Exists() {
		data.Wsm = tfutils.BoolFromString(value.String())
	} else {
		data.Wsm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsdm`); value.Exists() {
		data.Wsdm = tfutils.BoolFromString(value.String())
	} else {
		data.Wsdm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsrr-subscription`); value.Exists() {
		data.WsrrSubscription = tfutils.BoolFromString(value.String())
	} else {
		data.WsrrSubscription = types.BoolNull()
	}
}

func (data *DmXMLMgmtModes) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `any`); value.Exists() && !data.Any.IsNull() {
		data.Any = tfutils.BoolFromString(value.String())
	} else if !data.Any.ValueBool() {
		data.Any = types.BoolNull()
	}
	if value := res.Get(pathRoot + `soma`); value.Exists() && !data.Soma.IsNull() {
		data.Soma = tfutils.BoolFromString(value.String())
	} else if !data.Soma.ValueBool() {
		data.Soma = types.BoolNull()
	}
	if value := res.Get(pathRoot + `v2004`); value.Exists() && !data.V2004.IsNull() {
		data.V2004 = tfutils.BoolFromString(value.String())
	} else if !data.V2004.ValueBool() {
		data.V2004 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `amp`); value.Exists() && !data.Amp.IsNull() {
		data.Amp = tfutils.BoolFromString(value.String())
	} else if !data.Amp.ValueBool() {
		data.Amp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `slm`); value.Exists() && !data.Slm.IsNull() {
		data.Slm = tfutils.BoolFromString(value.String())
	} else if !data.Slm.ValueBool() {
		data.Slm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsm`); value.Exists() && !data.Wsm.IsNull() {
		data.Wsm = tfutils.BoolFromString(value.String())
	} else if data.Wsm.ValueBool() {
		data.Wsm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsdm`); value.Exists() && !data.Wsdm.IsNull() {
		data.Wsdm = tfutils.BoolFromString(value.String())
	} else if data.Wsdm.ValueBool() {
		data.Wsdm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `wsrr-subscription`); value.Exists() && !data.WsrrSubscription.IsNull() {
		data.WsrrSubscription = tfutils.BoolFromString(value.String())
	} else if !data.WsrrSubscription.ValueBool() {
		data.WsrrSubscription = types.BoolNull()
	}
}
