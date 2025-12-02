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

type KafkaSourceProtocolHandler struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Cluster           types.String                `tfsdk:"cluster"`
	RequestTopic      types.String                `tfsdk:"request_topic"`
	ResponseTopic     types.String                `tfsdk:"response_topic"`
	ConsumerGroup     types.String                `tfsdk:"consumer_group"`
	BatchSize         types.Int64                 `tfsdk:"batch_size"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var KafkaSourceProtocolHandlerObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"cluster":            types.StringType,
	"request_topic":      types.StringType,
	"response_topic":     types.StringType,
	"consumer_group":     types.StringType,
	"batch_size":         types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data KafkaSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/KafkaSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data KafkaSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Cluster.IsNull() {
		return false
	}
	if !data.RequestTopic.IsNull() {
		return false
	}
	if !data.ResponseTopic.IsNull() {
		return false
	}
	if !data.ConsumerGroup.IsNull() {
		return false
	}
	if !data.BatchSize.IsNull() {
		return false
	}
	return true
}

func (data KafkaSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Cluster.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Cluster`, data.Cluster.ValueString())
	}
	if !data.RequestTopic.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestTopic`, data.RequestTopic.ValueString())
	}
	if !data.ResponseTopic.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseTopic`, data.ResponseTopic.ValueString())
	}
	if !data.ConsumerGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConsumerGroup`, data.ConsumerGroup.ValueString())
	}
	if !data.BatchSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BatchSize`, data.BatchSize.ValueInt64())
	}
	return body
}

func (data *KafkaSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Cluster`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Cluster = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Cluster = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopic`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseTopic`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConsumerGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConsumerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConsumerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `BatchSize`); value.Exists() {
		data.BatchSize = types.Int64Value(value.Int())
	} else {
		data.BatchSize = types.Int64Value(1)
	}
}

func (data *KafkaSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Cluster`); value.Exists() && !data.Cluster.IsNull() {
		data.Cluster = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Cluster = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopic`); value.Exists() && !data.RequestTopic.IsNull() {
		data.RequestTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseTopic`); value.Exists() && !data.ResponseTopic.IsNull() {
		data.ResponseTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConsumerGroup`); value.Exists() && !data.ConsumerGroup.IsNull() {
		data.ConsumerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConsumerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `BatchSize`); value.Exists() && !data.BatchSize.IsNull() {
		data.BatchSize = types.Int64Value(value.Int())
	} else if data.BatchSize.ValueInt64() != 1 {
		data.BatchSize = types.Int64Null()
	}
}
