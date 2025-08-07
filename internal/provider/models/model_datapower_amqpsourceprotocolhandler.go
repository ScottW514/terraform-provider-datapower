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

type AMQPSourceProtocolHandler struct {
	Id                types.String      `tfsdk:"id"`
	AppDomain         types.String      `tfsdk:"app_domain"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	Broker            types.String      `tfsdk:"broker"`
	From              types.String      `tfsdk:"from"`
	To                types.String      `tfsdk:"to"`
	Credit            types.Int64       `tfsdk:"credit"`
	IgnoreReplyTo     types.Bool        `tfsdk:"ignore_reply_to"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var AMQPSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"broker":             types.StringType,
	"from":               types.StringType,
	"to":                 types.StringType,
	"credit":             types.Int64Type,
	"ignore_reply_to":    types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data AMQPSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AMQPSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AMQPSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Broker.IsNull() {
		return false
	}
	if !data.From.IsNull() {
		return false
	}
	if !data.To.IsNull() {
		return false
	}
	if !data.Credit.IsNull() {
		return false
	}
	if !data.IgnoreReplyTo.IsNull() {
		return false
	}
	return true
}

func (data AMQPSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Broker.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Broker`, data.Broker.ValueString())
	}
	if !data.From.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`From`, data.From.ValueString())
	}
	if !data.To.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`To`, data.To.ValueString())
	}
	if !data.Credit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Credit`, data.Credit.ValueInt64())
	}
	if !data.IgnoreReplyTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IgnoreReplyTo`, tfutils.StringFromBool(data.IgnoreReplyTo, ""))
	}
	return body
}

func (data *AMQPSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Broker`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Broker = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Broker = types.StringNull()
	}
	if value := res.Get(pathRoot + `From`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.From = tfutils.ParseStringFromGJSON(value)
	} else {
		data.From = types.StringNull()
	}
	if value := res.Get(pathRoot + `To`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.To = tfutils.ParseStringFromGJSON(value)
	} else {
		data.To = types.StringNull()
	}
	if value := res.Get(pathRoot + `Credit`); value.Exists() {
		data.Credit = types.Int64Value(value.Int())
	} else {
		data.Credit = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `IgnoreReplyTo`); value.Exists() {
		data.IgnoreReplyTo = tfutils.BoolFromString(value.String())
	} else {
		data.IgnoreReplyTo = types.BoolNull()
	}
}

func (data *AMQPSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Broker`); value.Exists() && !data.Broker.IsNull() {
		data.Broker = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Broker = types.StringNull()
	}
	if value := res.Get(pathRoot + `From`); value.Exists() && !data.From.IsNull() {
		data.From = tfutils.ParseStringFromGJSON(value)
	} else {
		data.From = types.StringNull()
	}
	if value := res.Get(pathRoot + `To`); value.Exists() && !data.To.IsNull() {
		data.To = tfutils.ParseStringFromGJSON(value)
	} else {
		data.To = types.StringNull()
	}
	if value := res.Get(pathRoot + `Credit`); value.Exists() && !data.Credit.IsNull() {
		data.Credit = types.Int64Value(value.Int())
	} else if data.Credit.ValueInt64() != 100 {
		data.Credit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IgnoreReplyTo`); value.Exists() && !data.IgnoreReplyTo.IsNull() {
		data.IgnoreReplyTo = tfutils.BoolFromString(value.String())
	} else if !data.IgnoreReplyTo.ValueBool() {
		data.IgnoreReplyTo = types.BoolNull()
	}
}
