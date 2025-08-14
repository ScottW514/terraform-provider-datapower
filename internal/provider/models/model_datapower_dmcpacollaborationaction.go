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

type DmCPACollaborationAction struct {
	Name            types.String `tfsdk:"name"`
	Value           types.String `tfsdk:"value"`
	Capability      types.String `tfsdk:"capability"`
	SenderSetting   types.String `tfsdk:"sender_setting"`
	ReceiverSetting types.String `tfsdk:"receiver_setting"`
}

var DmCPACollaborationActionObjectType = map[string]attr.Type{
	"name":             types.StringType,
	"value":            types.StringType,
	"capability":       types.StringType,
	"sender_setting":   types.StringType,
	"receiver_setting": types.StringType,
}
var DmCPACollaborationActionObjectDefault = map[string]attr.Value{
	"name":             types.StringNull(),
	"value":            types.StringNull(),
	"capability":       types.StringValue("cansend"),
	"sender_setting":   types.StringNull(),
	"receiver_setting": types.StringNull(),
}
var DmCPACollaborationActionDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies an ID for naming the action", "name", "").String,
			Computed:            true,
		},
		"value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of Action. For outbound message, the value is used in the Action element of the ebXML Message Header. For inbound transaction, the Action value is used to identify the action binding for processing the incoming message within the Service.", "value", "").String,
			Computed:            true,
		},
		"capability": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the type of this action binding.", "capability", "").AddStringEnum("cansend", "canreceive").AddDefaultValue("cansend").String,
			Computed:            true,
		},
		"sender_setting": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of sender setting to bind. A sender setting defines the Message-sending characteristics in Delivery Channels. It consists of document-exchange configurations and transport configurations.", "sender-setting", "b2bcpasendersetting").String,
			Computed:            true,
		},
		"receiver_setting": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of receiver setting to bind. A receiver setting defines the Message-receiving characteristics in Delivery Channels. It consists of document-exchange configurations and transport configurations.", "receiver-setting", "b2bcpareceiversetting").String,
			Computed:            true,
		},
	},
}
var DmCPACollaborationActionResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies an ID for naming the action", "name", "").String,
			Required:            true,
		},
		"value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of Action. For outbound message, the value is used in the Action element of the ebXML Message Header. For inbound transaction, the Action value is used to identify the action binding for processing the incoming message within the Service.", "value", "").String,
			Required:            true,
		},
		"capability": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the type of this action binding.", "capability", "").AddStringEnum("cansend", "canreceive").AddDefaultValue("cansend").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("cansend", "canreceive"),
			},
			Default: stringdefault.StaticString("cansend"),
		},
		"sender_setting": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of sender setting to bind. A sender setting defines the Message-sending characteristics in Delivery Channels. It consists of document-exchange configurations and transport configurations.", "sender-setting", "b2bcpasendersetting").String,
			Optional:            true,
		},
		"receiver_setting": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of receiver setting to bind. A receiver setting defines the Message-receiving characteristics in Delivery Channels. It consists of document-exchange configurations and transport configurations.", "receiver-setting", "b2bcpareceiversetting").String,
			Optional:            true,
		},
	},
}

func (data DmCPACollaborationAction) IsNull() bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.Capability.IsNull() {
		return false
	}
	if !data.SenderSetting.IsNull() {
		return false
	}
	if !data.ReceiverSetting.IsNull() {
		return false
	}
	return true
}

func (data DmCPACollaborationAction) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.Capability.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Capability`, data.Capability.ValueString())
	}
	if !data.SenderSetting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SenderSetting`, data.SenderSetting.ValueString())
	}
	if !data.ReceiverSetting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReceiverSetting`, data.ReceiverSetting.ValueString())
	}
	return body
}

func (data *DmCPACollaborationAction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Capability`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Capability = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Capability = types.StringValue("cansend")
	}
	if value := res.Get(pathRoot + `SenderSetting`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SenderSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverSetting`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReceiverSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverSetting = types.StringNull()
	}
}

func (data *DmCPACollaborationAction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `Capability`); value.Exists() && !data.Capability.IsNull() {
		data.Capability = tfutils.ParseStringFromGJSON(value)
	} else if data.Capability.ValueString() != "cansend" {
		data.Capability = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderSetting`); value.Exists() && !data.SenderSetting.IsNull() {
		data.SenderSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverSetting`); value.Exists() && !data.ReceiverSetting.IsNull() {
		data.ReceiverSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverSetting = types.StringNull()
	}
}
