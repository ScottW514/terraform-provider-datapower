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

type B2BXPathRoutingPolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	SenderXpath       types.String                `tfsdk:"sender_xpath"`
	ReceiverXpath     types.String                `tfsdk:"receiver_xpath"`
	DocumentIdXpath   types.String                `tfsdk:"document_id_xpath"`
	DateTimeXpath     types.String                `tfsdk:"date_time_xpath"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BXPathRoutingPolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"sender_xpath":       types.StringType,
	"receiver_xpath":     types.StringType,
	"document_id_xpath":  types.StringType,
	"date_time_xpath":    types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data B2BXPathRoutingPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BXPathRoutingPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BXPathRoutingPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.SenderXpath.IsNull() {
		return false
	}
	if !data.ReceiverXpath.IsNull() {
		return false
	}
	if !data.DocumentIdXpath.IsNull() {
		return false
	}
	if !data.DateTimeXpath.IsNull() {
		return false
	}
	return true
}

func (data B2BXPathRoutingPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.SenderXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SenderXPath`, data.SenderXpath.ValueString())
	}
	if !data.ReceiverXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReceiverXPath`, data.ReceiverXpath.ValueString())
	}
	if !data.DocumentIdXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentIDXPath`, data.DocumentIdXpath.ValueString())
	}
	if !data.DateTimeXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DateTimeXPath`, data.DateTimeXpath.ValueString())
	}
	return body
}

func (data *B2BXPathRoutingPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SenderXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SenderXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReceiverXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentIDXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocumentIdXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentIdXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DateTimeXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DateTimeXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DateTimeXpath = types.StringNull()
	}
}

func (data *B2BXPathRoutingPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SenderXPath`); value.Exists() && !data.SenderXpath.IsNull() {
		data.SenderXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverXPath`); value.Exists() && !data.ReceiverXpath.IsNull() {
		data.ReceiverXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentIDXPath`); value.Exists() && !data.DocumentIdXpath.IsNull() {
		data.DocumentIdXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentIdXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DateTimeXPath`); value.Exists() && !data.DateTimeXpath.IsNull() {
		data.DateTimeXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DateTimeXpath = types.StringNull()
	}
}
