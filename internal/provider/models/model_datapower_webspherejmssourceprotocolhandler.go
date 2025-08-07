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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebSphereJMSSourceProtocolHandler struct {
	Id                     types.String      `tfsdk:"id"`
	AppDomain              types.String      `tfsdk:"app_domain"`
	Server                 types.String      `tfsdk:"server"`
	RequestTopicSpace      types.String      `tfsdk:"request_topic_space"`
	ReplyTopicSpace        types.String      `tfsdk:"reply_topic_space"`
	StrictMessageOrder     types.Bool        `tfsdk:"strict_message_order"`
	UserSummary            types.String      `tfsdk:"user_summary"`
	GetQueue               types.String      `tfsdk:"get_queue"`
	PutQueue               types.String      `tfsdk:"put_queue"`
	Selector               types.String      `tfsdk:"selector"`
	AsyncMessageProcessing types.Bool        `tfsdk:"async_message_processing"`
	DependencyActions      []*actions.Action `tfsdk:"dependency_actions"`
}

var WebSphereJMSSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"server":                   types.StringType,
	"request_topic_space":      types.StringType,
	"reply_topic_space":        types.StringType,
	"strict_message_order":     types.BoolType,
	"user_summary":             types.StringType,
	"get_queue":                types.StringType,
	"put_queue":                types.StringType,
	"selector":                 types.StringType,
	"async_message_processing": types.BoolType,
	"dependency_actions":       actions.ActionsListType,
}

func (data WebSphereJMSSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebSphereJMSSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebSphereJMSSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Server.IsNull() {
		return false
	}
	if !data.RequestTopicSpace.IsNull() {
		return false
	}
	if !data.ReplyTopicSpace.IsNull() {
		return false
	}
	if !data.StrictMessageOrder.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.GetQueue.IsNull() {
		return false
	}
	if !data.PutQueue.IsNull() {
		return false
	}
	if !data.Selector.IsNull() {
		return false
	}
	if !data.AsyncMessageProcessing.IsNull() {
		return false
	}
	return true
}

func (data WebSphereJMSSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Server.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Server`, data.Server.ValueString())
	}
	if !data.RequestTopicSpace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestTopicSpace`, data.RequestTopicSpace.ValueString())
	}
	if !data.ReplyTopicSpace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReplyTopicSpace`, data.ReplyTopicSpace.ValueString())
	}
	if !data.StrictMessageOrder.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StrictMessageOrder`, tfutils.StringFromBool(data.StrictMessageOrder, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.GetQueue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetQueue`, data.GetQueue.ValueString())
	}
	if !data.PutQueue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PutQueue`, data.PutQueue.ValueString())
	}
	if !data.Selector.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Selector`, data.Selector.ValueString())
	}
	if !data.AsyncMessageProcessing.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AsyncMessageProcessing`, tfutils.StringFromBool(data.AsyncMessageProcessing, ""))
	}
	return body
}

func (data *WebSphereJMSSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopicSpace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestTopicSpace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopicSpace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReplyTopicSpace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReplyTopicSpace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReplyTopicSpace = types.StringNull()
	}
	if value := res.Get(pathRoot + `StrictMessageOrder`); value.Exists() {
		data.StrictMessageOrder = tfutils.BoolFromString(value.String())
	} else {
		data.StrictMessageOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PutQueue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PutQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PutQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Selector`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Selector = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Selector = types.StringNull()
	}
	if value := res.Get(pathRoot + `AsyncMessageProcessing`); value.Exists() {
		data.AsyncMessageProcessing = tfutils.BoolFromString(value.String())
	} else {
		data.AsyncMessageProcessing = types.BoolNull()
	}
}

func (data *WebSphereJMSSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && !data.Server.IsNull() {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopicSpace`); value.Exists() && !data.RequestTopicSpace.IsNull() {
		data.RequestTopicSpace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopicSpace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReplyTopicSpace`); value.Exists() && !data.ReplyTopicSpace.IsNull() {
		data.ReplyTopicSpace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReplyTopicSpace = types.StringNull()
	}
	if value := res.Get(pathRoot + `StrictMessageOrder`); value.Exists() && !data.StrictMessageOrder.IsNull() {
		data.StrictMessageOrder = tfutils.BoolFromString(value.String())
	} else if data.StrictMessageOrder.ValueBool() {
		data.StrictMessageOrder = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && !data.GetQueue.IsNull() {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PutQueue`); value.Exists() && !data.PutQueue.IsNull() {
		data.PutQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PutQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `Selector`); value.Exists() && !data.Selector.IsNull() {
		data.Selector = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Selector = types.StringNull()
	}
	if value := res.Get(pathRoot + `AsyncMessageProcessing`); value.Exists() && !data.AsyncMessageProcessing.IsNull() {
		data.AsyncMessageProcessing = tfutils.BoolFromString(value.String())
	} else if data.AsyncMessageProcessing.ValueBool() {
		data.AsyncMessageProcessing = types.BoolNull()
	}
}
