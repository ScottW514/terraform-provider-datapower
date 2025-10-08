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

type DmB2BBackupMsgType struct {
	InMessage       types.Bool `tfsdk:"in_message"`
	OutMessage      types.Bool `tfsdk:"out_message"`
	InMdn           types.Bool `tfsdk:"in_mdn"`
	OutMdn          types.Bool `tfsdk:"out_mdn"`
	EbmsInMessage   types.Bool `tfsdk:"ebms_in_message"`
	EbmsOutMessage  types.Bool `tfsdk:"ebms_out_message"`
	EbmsInAck       types.Bool `tfsdk:"ebms_in_ack"`
	EbmsOutAck      types.Bool `tfsdk:"ebms_out_ack"`
	Ebms3InMessage  types.Bool `tfsdk:"ebms3_in_message"`
	Ebms3OutMessage types.Bool `tfsdk:"ebms3_out_message"`
	Ebms3InRcpt     types.Bool `tfsdk:"ebms3_in_rcpt"`
	Ebms3OutRcpt    types.Bool `tfsdk:"ebms3_out_rcpt"`
}

var DmB2BBackupMsgTypeObjectType = map[string]attr.Type{
	"in_message":        types.BoolType,
	"out_message":       types.BoolType,
	"in_mdn":            types.BoolType,
	"out_mdn":           types.BoolType,
	"ebms_in_message":   types.BoolType,
	"ebms_out_message":  types.BoolType,
	"ebms_in_ack":       types.BoolType,
	"ebms_out_ack":      types.BoolType,
	"ebms3_in_message":  types.BoolType,
	"ebms3_out_message": types.BoolType,
	"ebms3_in_rcpt":     types.BoolType,
	"ebms3_out_rcpt":    types.BoolType,
}
var DmB2BBackupMsgTypeObjectDefault = map[string]attr.Value{
	"in_message":        types.BoolValue(true),
	"out_message":       types.BoolValue(true),
	"in_mdn":            types.BoolValue(true),
	"out_mdn":           types.BoolValue(true),
	"ebms_in_message":   types.BoolValue(true),
	"ebms_out_message":  types.BoolValue(true),
	"ebms_in_ack":       types.BoolValue(true),
	"ebms_out_ack":      types.BoolValue(true),
	"ebms3_in_message":  types.BoolValue(true),
	"ebms3_out_message": types.BoolValue(true),
	"ebms3_in_rcpt":     types.BoolValue(true),
	"ebms3_out_rcpt":    types.BoolValue(true),
}

func GetDmB2BBackupMsgTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmB2BBackupMsgTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"in_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "AS inbound message",
				Computed:            true,
			},
			"out_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "AS outbound message",
				Computed:            true,
			},
			"in_mdn": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "AS inbound MDN",
				Computed:            true,
			},
			"out_mdn": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "AS outbound MDN",
				Computed:            true,
			},
			"ebms_in_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS2 inbound message",
				Computed:            true,
			},
			"ebms_out_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS2 outbound message",
				Computed:            true,
			},
			"ebms_in_ack": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS2 inbound ACK",
				Computed:            true,
			},
			"ebms_out_ack": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS2 outbound ACK",
				Computed:            true,
			},
			"ebms3_in_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS3 inbound message",
				Computed:            true,
			},
			"ebms3_out_message": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS3 outbound message",
				Computed:            true,
			},
			"ebms3_in_rcpt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS3 inbound receipt",
				Computed:            true,
			},
			"ebms3_out_rcpt": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "ebMS3 outbound receipt",
				Computed:            true,
			},
		},
	}
	DmB2BBackupMsgTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmB2BBackupMsgTypeDataSourceSchema
}
func GetDmB2BBackupMsgTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmB2BBackupMsgTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmB2BBackupMsgTypeObjectType,
				DmB2BBackupMsgTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"in_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AS inbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"out_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AS outbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"in_mdn": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AS inbound MDN", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"out_mdn": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AS outbound MDN", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms_in_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS2 inbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms_out_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS2 outbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms_in_ack": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS2 inbound ACK", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms_out_ack": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS2 outbound ACK", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms3_in_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS3 inbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms3_out_message": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS3 outbound message", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms3_in_rcpt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS3 inbound receipt", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ebms3_out_rcpt": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("ebMS3 outbound receipt", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmB2BBackupMsgTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmB2BBackupMsgTypeResourceSchema.Required = true
	} else {
		DmB2BBackupMsgTypeResourceSchema.Optional = true
		DmB2BBackupMsgTypeResourceSchema.Computed = true
	}
	return DmB2BBackupMsgTypeResourceSchema
}

func (data DmB2BBackupMsgType) IsNull() bool {
	if !data.InMessage.IsNull() {
		return false
	}
	if !data.OutMessage.IsNull() {
		return false
	}
	if !data.InMdn.IsNull() {
		return false
	}
	if !data.OutMdn.IsNull() {
		return false
	}
	if !data.EbmsInMessage.IsNull() {
		return false
	}
	if !data.EbmsOutMessage.IsNull() {
		return false
	}
	if !data.EbmsInAck.IsNull() {
		return false
	}
	if !data.EbmsOutAck.IsNull() {
		return false
	}
	if !data.Ebms3InMessage.IsNull() {
		return false
	}
	if !data.Ebms3OutMessage.IsNull() {
		return false
	}
	if !data.Ebms3InRcpt.IsNull() {
		return false
	}
	if !data.Ebms3OutRcpt.IsNull() {
		return false
	}
	return true
}

func (data DmB2BBackupMsgType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.InMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InMessage`, tfutils.StringFromBool(data.InMessage, ""))
	}
	if !data.OutMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutMessage`, tfutils.StringFromBool(data.OutMessage, ""))
	}
	if !data.InMdn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InMDN`, tfutils.StringFromBool(data.InMdn, ""))
	}
	if !data.OutMdn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutMDN`, tfutils.StringFromBool(data.OutMdn, ""))
	}
	if !data.EbmsInMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInMessage`, tfutils.StringFromBool(data.EbmsInMessage, ""))
	}
	if !data.EbmsOutMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutMessage`, tfutils.StringFromBool(data.EbmsOutMessage, ""))
	}
	if !data.EbmsInAck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSInACK`, tfutils.StringFromBool(data.EbmsInAck, ""))
	}
	if !data.EbmsOutAck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutACK`, tfutils.StringFromBool(data.EbmsOutAck, ""))
	}
	if !data.Ebms3InMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InMessage`, tfutils.StringFromBool(data.Ebms3InMessage, ""))
	}
	if !data.Ebms3OutMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutMessage`, tfutils.StringFromBool(data.Ebms3OutMessage, ""))
	}
	if !data.Ebms3InRcpt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3InRCPT`, tfutils.StringFromBool(data.Ebms3InRcpt, ""))
	}
	if !data.Ebms3OutRcpt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMS3OutRCPT`, tfutils.StringFromBool(data.Ebms3OutRcpt, ""))
	}
	return body
}

func (data *DmB2BBackupMsgType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InMessage`); value.Exists() {
		data.InMessage = tfutils.BoolFromString(value.String())
	} else {
		data.InMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutMessage`); value.Exists() {
		data.OutMessage = tfutils.BoolFromString(value.String())
	} else {
		data.OutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InMDN`); value.Exists() {
		data.InMdn = tfutils.BoolFromString(value.String())
	} else {
		data.InMdn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutMDN`); value.Exists() {
		data.OutMdn = tfutils.BoolFromString(value.String())
	} else {
		data.OutMdn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInMessage`); value.Exists() {
		data.EbmsInMessage = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutMessage`); value.Exists() {
		data.EbmsOutMessage = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInACK`); value.Exists() {
		data.EbmsInAck = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsInAck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutACK`); value.Exists() {
		data.EbmsOutAck = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutAck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InMessage`); value.Exists() {
		data.Ebms3InMessage = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3InMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutMessage`); value.Exists() {
		data.Ebms3OutMessage = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3OutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InRCPT`); value.Exists() {
		data.Ebms3InRcpt = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3InRcpt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutRCPT`); value.Exists() {
		data.Ebms3OutRcpt = tfutils.BoolFromString(value.String())
	} else {
		data.Ebms3OutRcpt = types.BoolNull()
	}
}

func (data *DmB2BBackupMsgType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `InMessage`); value.Exists() && !data.InMessage.IsNull() {
		data.InMessage = tfutils.BoolFromString(value.String())
	} else if !data.InMessage.ValueBool() {
		data.InMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutMessage`); value.Exists() && !data.OutMessage.IsNull() {
		data.OutMessage = tfutils.BoolFromString(value.String())
	} else if !data.OutMessage.ValueBool() {
		data.OutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InMDN`); value.Exists() && !data.InMdn.IsNull() {
		data.InMdn = tfutils.BoolFromString(value.String())
	} else if !data.InMdn.ValueBool() {
		data.InMdn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutMDN`); value.Exists() && !data.OutMdn.IsNull() {
		data.OutMdn = tfutils.BoolFromString(value.String())
	} else if !data.OutMdn.ValueBool() {
		data.OutMdn = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInMessage`); value.Exists() && !data.EbmsInMessage.IsNull() {
		data.EbmsInMessage = tfutils.BoolFromString(value.String())
	} else if !data.EbmsInMessage.ValueBool() {
		data.EbmsInMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutMessage`); value.Exists() && !data.EbmsOutMessage.IsNull() {
		data.EbmsOutMessage = tfutils.BoolFromString(value.String())
	} else if !data.EbmsOutMessage.ValueBool() {
		data.EbmsOutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSInACK`); value.Exists() && !data.EbmsInAck.IsNull() {
		data.EbmsInAck = tfutils.BoolFromString(value.String())
	} else if !data.EbmsInAck.ValueBool() {
		data.EbmsInAck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutACK`); value.Exists() && !data.EbmsOutAck.IsNull() {
		data.EbmsOutAck = tfutils.BoolFromString(value.String())
	} else if !data.EbmsOutAck.ValueBool() {
		data.EbmsOutAck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InMessage`); value.Exists() && !data.Ebms3InMessage.IsNull() {
		data.Ebms3InMessage = tfutils.BoolFromString(value.String())
	} else if !data.Ebms3InMessage.ValueBool() {
		data.Ebms3InMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutMessage`); value.Exists() && !data.Ebms3OutMessage.IsNull() {
		data.Ebms3OutMessage = tfutils.BoolFromString(value.String())
	} else if !data.Ebms3OutMessage.ValueBool() {
		data.Ebms3OutMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3InRCPT`); value.Exists() && !data.Ebms3InRcpt.IsNull() {
		data.Ebms3InRcpt = tfutils.BoolFromString(value.String())
	} else if !data.Ebms3InRcpt.ValueBool() {
		data.Ebms3InRcpt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMS3OutRCPT`); value.Exists() && !data.Ebms3OutRcpt.IsNull() {
		data.Ebms3OutRcpt = tfutils.BoolFromString(value.String())
	} else if !data.Ebms3OutRcpt.ValueBool() {
		data.Ebms3OutRcpt = types.BoolNull()
	}
}
