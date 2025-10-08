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

type DmWSRRSavedSearchWSDLSource struct {
	WsrrSavedSearchSubscription types.String `tfsdk:"wsrr_saved_search_subscription"`
	WsrrAttachment              types.String `tfsdk:"wsrr_attachment"`
}

var DmWSRRSavedSearchWSDLSourceObjectType = map[string]attr.Type{
	"wsrr_saved_search_subscription": types.StringType,
	"wsrr_attachment":                types.StringType,
}
var DmWSRRSavedSearchWSDLSourceObjectDefault = map[string]attr.Value{
	"wsrr_saved_search_subscription": types.StringNull(),
	"wsrr_attachment":                types.StringNull(),
}

func GetDmWSRRSavedSearchWSDLSourceDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSRRSavedSearchWSDLSourceDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"wsrr_saved_search_subscription": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a WSRR saved search subscription",
				Computed:            true,
			},
			"wsrr_attachment": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a WS-Policy attachment",
				Computed:            true,
			},
		},
	}
	return DmWSRRSavedSearchWSDLSourceDataSourceSchema
}
func GetDmWSRRSavedSearchWSDLSourceResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSRRSavedSearchWSDLSourceResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"wsrr_saved_search_subscription": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a WSRR saved search subscription", "", "wsrr_saved_search_subscription").String,
				Optional:            true,
			},
			"wsrr_attachment": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a WS-Policy attachment", "", "policy_attachments").String,
				Optional:            true,
			},
		},
	}
	return DmWSRRSavedSearchWSDLSourceResourceSchema
}

func (data DmWSRRSavedSearchWSDLSource) IsNull() bool {
	if !data.WsrrSavedSearchSubscription.IsNull() {
		return false
	}
	if !data.WsrrAttachment.IsNull() {
		return false
	}
	return true
}

func (data DmWSRRSavedSearchWSDLSource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.WsrrSavedSearchSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRSavedSearchSubscription`, data.WsrrSavedSearchSubscription.ValueString())
	}
	if !data.WsrrAttachment.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSRRAttachment`, data.WsrrAttachment.ValueString())
	}
	return body
}

func (data *DmWSRRSavedSearchWSDLSource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrSavedSearchSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSavedSearchSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRAttachment`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsrrAttachment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrAttachment = types.StringNull()
	}
}

func (data *DmWSRRSavedSearchWSDLSource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSRRSavedSearchSubscription`); value.Exists() && !data.WsrrSavedSearchSubscription.IsNull() {
		data.WsrrSavedSearchSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrSavedSearchSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSRRAttachment`); value.Exists() && !data.WsrrAttachment.IsNull() {
		data.WsrrAttachment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsrrAttachment = types.StringNull()
	}
}
