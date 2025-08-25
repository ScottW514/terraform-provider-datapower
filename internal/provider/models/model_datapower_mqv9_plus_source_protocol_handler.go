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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MQv9PlusSourceProtocolHandler struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	QueueManager            types.String                `tfsdk:"queue_manager"`
	GetQueue                types.String                `tfsdk:"get_queue"`
	SubscribeTopicString    types.String                `tfsdk:"subscribe_topic_string"`
	SubscriptionName        types.String                `tfsdk:"subscription_name"`
	PutQueue                types.String                `tfsdk:"put_queue"`
	PublishTopicString      types.String                `tfsdk:"publish_topic_string"`
	CodePage                types.Int64                 `tfsdk:"code_page"`
	GetMessageOptions       types.Int64                 `tfsdk:"get_message_options"`
	MessageSelector         types.String                `tfsdk:"message_selector"`
	ParseProperties         types.Bool                  `tfsdk:"parse_properties"`
	AsyncPut                types.Bool                  `tfsdk:"async_put"`
	ExcludeHeaders          *DmMQHeaders                `tfsdk:"exclude_headers"`
	ConcurrentConnections   types.Int64                 `tfsdk:"concurrent_connections"`
	PollingInterval         types.Int64                 `tfsdk:"polling_interval"`
	BatchSize               types.Int64                 `tfsdk:"batch_size"`
	ContentTypeHeader       types.String                `tfsdk:"content_type_header"`
	ContentTypeXPath        types.String                `tfsdk:"content_type_x_path"`
	RetrieveBackoutSettings types.Bool                  `tfsdk:"retrieve_backout_settings"`
	UseQmNameInUrl          types.Bool                  `tfsdk:"use_qm_name_in_url"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var MQv9PlusSourceProtocolHandlerGetQueueCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "subscribe_topic_string",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{""},
}
var MQv9PlusSourceProtocolHandlerSubscribeTopicStringCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "get_queue",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{""},
}
var MQv9PlusSourceProtocolHandlerContentTypeXPathCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "content_type_header",
	AttrType:    "String",
	AttrDefault: "None",
	Value:       []string{"MQRFH", "MQRFH2"},
}

var MQv9PlusSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"queue_manager":             types.StringType,
	"get_queue":                 types.StringType,
	"subscribe_topic_string":    types.StringType,
	"subscription_name":         types.StringType,
	"put_queue":                 types.StringType,
	"publish_topic_string":      types.StringType,
	"code_page":                 types.Int64Type,
	"get_message_options":       types.Int64Type,
	"message_selector":          types.StringType,
	"parse_properties":          types.BoolType,
	"async_put":                 types.BoolType,
	"exclude_headers":           types.ObjectType{AttrTypes: DmMQHeadersObjectType},
	"concurrent_connections":    types.Int64Type,
	"polling_interval":          types.Int64Type,
	"batch_size":                types.Int64Type,
	"content_type_header":       types.StringType,
	"content_type_x_path":       types.StringType,
	"retrieve_backout_settings": types.BoolType,
	"use_qm_name_in_url":        types.BoolType,
	"dependency_actions":        actions.ActionsListType,
}

func (data MQv9PlusSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MQv9PlusSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MQv9PlusSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.QueueManager.IsNull() {
		return false
	}
	if !data.GetQueue.IsNull() {
		return false
	}
	if !data.SubscribeTopicString.IsNull() {
		return false
	}
	if !data.SubscriptionName.IsNull() {
		return false
	}
	if !data.PutQueue.IsNull() {
		return false
	}
	if !data.PublishTopicString.IsNull() {
		return false
	}
	if !data.CodePage.IsNull() {
		return false
	}
	if !data.GetMessageOptions.IsNull() {
		return false
	}
	if !data.MessageSelector.IsNull() {
		return false
	}
	if !data.ParseProperties.IsNull() {
		return false
	}
	if !data.AsyncPut.IsNull() {
		return false
	}
	if data.ExcludeHeaders != nil {
		if !data.ExcludeHeaders.IsNull() {
			return false
		}
	}
	if !data.ConcurrentConnections.IsNull() {
		return false
	}
	if !data.PollingInterval.IsNull() {
		return false
	}
	if !data.BatchSize.IsNull() {
		return false
	}
	if !data.ContentTypeHeader.IsNull() {
		return false
	}
	if !data.ContentTypeXPath.IsNull() {
		return false
	}
	if !data.RetrieveBackoutSettings.IsNull() {
		return false
	}
	if !data.UseQmNameInUrl.IsNull() {
		return false
	}
	return true
}

func (data MQv9PlusSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.QueueManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueueManager`, data.QueueManager.ValueString())
	}
	if !data.GetQueue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetQueue`, data.GetQueue.ValueString())
	}
	if !data.SubscribeTopicString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SubscribeTopicString`, data.SubscribeTopicString.ValueString())
	}
	if !data.SubscriptionName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SubscriptionName`, data.SubscriptionName.ValueString())
	}
	if !data.PutQueue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PutQueue`, data.PutQueue.ValueString())
	}
	if !data.PublishTopicString.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PublishTopicString`, data.PublishTopicString.ValueString())
	}
	if !data.CodePage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CodePage`, data.CodePage.ValueInt64())
	}
	if !data.GetMessageOptions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetMessageOptions`, data.GetMessageOptions.ValueInt64())
	}
	if !data.MessageSelector.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessageSelector`, data.MessageSelector.ValueString())
	}
	if !data.ParseProperties.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParseProperties`, tfutils.StringFromBool(data.ParseProperties, ""))
	}
	if !data.AsyncPut.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AsyncPut`, tfutils.StringFromBool(data.AsyncPut, ""))
	}
	if data.ExcludeHeaders != nil {
		if !data.ExcludeHeaders.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ExcludeHeaders`, data.ExcludeHeaders.ToBody(ctx, ""))
		}
	}
	if !data.ConcurrentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConcurrentConnections`, data.ConcurrentConnections.ValueInt64())
	}
	if !data.PollingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PollingInterval`, data.PollingInterval.ValueInt64())
	}
	if !data.BatchSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BatchSize`, data.BatchSize.ValueInt64())
	}
	if !data.ContentTypeHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContentTypeHeader`, data.ContentTypeHeader.ValueString())
	}
	if !data.ContentTypeXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ContentTypeXPath`, data.ContentTypeXPath.ValueString())
	}
	if !data.RetrieveBackoutSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetrieveBackoutSettings`, tfutils.StringFromBool(data.RetrieveBackoutSettings, ""))
	}
	if !data.UseQmNameInUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseQMNameInURL`, tfutils.StringFromBool(data.UseQmNameInUrl, ""))
	}
	return body
}

func (data *MQv9PlusSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueueManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueueManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubscribeTopicString`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SubscribeTopicString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubscribeTopicString = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubscriptionName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SubscriptionName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubscriptionName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PutQueue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PutQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PutQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishTopicString`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PublishTopicString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishTopicString = types.StringNull()
	}
	if value := res.Get(pathRoot + `CodePage`); value.Exists() {
		data.CodePage = types.Int64Value(value.Int())
	} else {
		data.CodePage = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `GetMessageOptions`); value.Exists() {
		data.GetMessageOptions = types.Int64Value(value.Int())
	} else {
		data.GetMessageOptions = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `MessageSelector`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MessageSelector = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageSelector = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseProperties`); value.Exists() {
		data.ParseProperties = tfutils.BoolFromString(value.String())
	} else {
		data.ParseProperties = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AsyncPut`); value.Exists() {
		data.AsyncPut = tfutils.BoolFromString(value.String())
	} else {
		data.AsyncPut = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExcludeHeaders`); value.Exists() {
		data.ExcludeHeaders = &DmMQHeaders{}
		data.ExcludeHeaders.FromBody(ctx, "", value)
	} else {
		data.ExcludeHeaders = nil
	}
	if value := res.Get(pathRoot + `ConcurrentConnections`); value.Exists() {
		data.ConcurrentConnections = types.Int64Value(value.Int())
	} else {
		data.ConcurrentConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else {
		data.PollingInterval = types.Int64Value(30)
	}
	if value := res.Get(pathRoot + `BatchSize`); value.Exists() {
		data.BatchSize = types.Int64Value(value.Int())
	} else {
		data.BatchSize = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ContentTypeHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ContentTypeHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentTypeHeader = types.StringValue("None")
	}
	if value := res.Get(pathRoot + `ContentTypeXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ContentTypeXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentTypeXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `RetrieveBackoutSettings`); value.Exists() {
		data.RetrieveBackoutSettings = tfutils.BoolFromString(value.String())
	} else {
		data.RetrieveBackoutSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseQMNameInURL`); value.Exists() {
		data.UseQmNameInUrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseQmNameInUrl = types.BoolNull()
	}
}

func (data *MQv9PlusSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() && !data.QueueManager.IsNull() {
		data.QueueManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueueManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `GetQueue`); value.Exists() && !data.GetQueue.IsNull() {
		data.GetQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GetQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubscribeTopicString`); value.Exists() && !data.SubscribeTopicString.IsNull() {
		data.SubscribeTopicString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubscribeTopicString = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubscriptionName`); value.Exists() && !data.SubscriptionName.IsNull() {
		data.SubscriptionName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubscriptionName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PutQueue`); value.Exists() && !data.PutQueue.IsNull() {
		data.PutQueue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PutQueue = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishTopicString`); value.Exists() && !data.PublishTopicString.IsNull() {
		data.PublishTopicString = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishTopicString = types.StringNull()
	}
	if value := res.Get(pathRoot + `CodePage`); value.Exists() && !data.CodePage.IsNull() {
		data.CodePage = types.Int64Value(value.Int())
	} else if data.CodePage.ValueInt64() != 0 {
		data.CodePage = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GetMessageOptions`); value.Exists() && !data.GetMessageOptions.IsNull() {
		data.GetMessageOptions = types.Int64Value(value.Int())
	} else if data.GetMessageOptions.ValueInt64() != 1 {
		data.GetMessageOptions = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MessageSelector`); value.Exists() && !data.MessageSelector.IsNull() {
		data.MessageSelector = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageSelector = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseProperties`); value.Exists() && !data.ParseProperties.IsNull() {
		data.ParseProperties = tfutils.BoolFromString(value.String())
	} else if data.ParseProperties.ValueBool() {
		data.ParseProperties = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AsyncPut`); value.Exists() && !data.AsyncPut.IsNull() {
		data.AsyncPut = tfutils.BoolFromString(value.String())
	} else if data.AsyncPut.ValueBool() {
		data.AsyncPut = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExcludeHeaders`); value.Exists() {
		data.ExcludeHeaders.UpdateFromBody(ctx, "", value)
	} else {
		data.ExcludeHeaders = nil
	}
	if value := res.Get(pathRoot + `ConcurrentConnections`); value.Exists() && !data.ConcurrentConnections.IsNull() {
		data.ConcurrentConnections = types.Int64Value(value.Int())
	} else if data.ConcurrentConnections.ValueInt64() != 1 {
		data.ConcurrentConnections = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() && !data.PollingInterval.IsNull() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else if data.PollingInterval.ValueInt64() != 30 {
		data.PollingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BatchSize`); value.Exists() && !data.BatchSize.IsNull() {
		data.BatchSize = types.Int64Value(value.Int())
	} else if data.BatchSize.ValueInt64() != 0 {
		data.BatchSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ContentTypeHeader`); value.Exists() && !data.ContentTypeHeader.IsNull() {
		data.ContentTypeHeader = tfutils.ParseStringFromGJSON(value)
	} else if data.ContentTypeHeader.ValueString() != "None" {
		data.ContentTypeHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentTypeXPath`); value.Exists() && !data.ContentTypeXPath.IsNull() {
		data.ContentTypeXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ContentTypeXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `RetrieveBackoutSettings`); value.Exists() && !data.RetrieveBackoutSettings.IsNull() {
		data.RetrieveBackoutSettings = tfutils.BoolFromString(value.String())
	} else if data.RetrieveBackoutSettings.ValueBool() {
		data.RetrieveBackoutSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseQMNameInURL`); value.Exists() && !data.UseQmNameInUrl.IsNull() {
		data.UseQmNameInUrl = tfutils.BoolFromString(value.String())
	} else if data.UseQmNameInUrl.ValueBool() {
		data.UseQmNameInUrl = types.BoolNull()
	}
}
