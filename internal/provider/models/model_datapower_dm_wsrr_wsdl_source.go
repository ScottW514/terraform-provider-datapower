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

type DmWSRRWSDLSource struct {
	WsrrSubscription types.String `tfsdk:"wsrr_subscription"`
	WsrrAttachment   types.String `tfsdk:"wsrr_attachment"`
}

var DmWSRRWSDLSourceObjectType = map[string]attr.Type{
	"wsrr_subscription": types.StringType,
	"wsrr_attachment":   types.StringType,
}
var DmWSRRWSDLSourceObjectDefault = map[string]attr.Value{
	"wsrr_subscription": types.StringNull(),
	"wsrr_attachment":   types.StringNull(),
}

func GetDmWSRRWSDLSourceDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSRRWSDLSourceDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"wsrr_subscription": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select WSRR subscription", "", "wsrr_subscription").String,
				Computed:            true,
			},
			"wsrr_attachment": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select WS-Policy attachment", "", "policy_attachments").String,
				Computed:            true,
			},
		},
	}
	return DmWSRRWSDLSourceDataSourceSchema
}
func GetDmWSRRWSDLSourceResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSRRWSDLSourceResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"wsrr_subscription": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select WSRR subscription", "", "wsrr_subscription").String,
				Optional:            true,
			},
			"wsrr_attachment": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select WS-Policy attachment", "", "policy_attachments").String,
				Optional:            true,
			},
		},
	}
	return DmWSRRWSDLSourceResourceSchema
}

func (data DmWSRRWSDLSource) IsNull() bool {
	if !data.WsrrSubscription.IsNull() {
		return false
	}
	if !data.WsrrAttachment.IsNull() {
		return false
	}
	return true
}

func (data DmWSRRWSDLSource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.WsrrSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRSubscription`, data.WsrrSubscription.ValueString())
	}
	if !data.WsrrAttachment.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRAttachment`, data.WsrrAttachment.ValueString())
	}
	return body
}

func (data *DmWSRRWSDLSource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSRRSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRAttachment`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrAttachment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrAttachment = types.StringNull()
	}
}

func (data *DmWSRRWSDLSource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSRRSubscription`); value.Exists() && !data.WsrrSubscription.IsNull() {
		data.WsrrSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRAttachment`); value.Exists() && !data.WsrrAttachment.IsNull() {
		data.WsrrAttachment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrAttachment = types.StringNull()
	}
}
