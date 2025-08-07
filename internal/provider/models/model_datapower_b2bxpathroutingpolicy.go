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
	SenderXPath       types.String                `tfsdk:"sender_x_path"`
	ReceiverXPath     types.String                `tfsdk:"receiver_x_path"`
	DocumentIdxPath   types.String                `tfsdk:"document_idx_path"`
	DateTimeXPath     types.String                `tfsdk:"date_time_x_path"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BXPathRoutingPolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"sender_x_path":      types.StringType,
	"receiver_x_path":    types.StringType,
	"document_idx_path":  types.StringType,
	"date_time_x_path":   types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data B2BXPathRoutingPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BXPathRoutingPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
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
	if !data.SenderXPath.IsNull() {
		return false
	}
	if !data.ReceiverXPath.IsNull() {
		return false
	}
	if !data.DocumentIdxPath.IsNull() {
		return false
	}
	if !data.DateTimeXPath.IsNull() {
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
	if !data.SenderXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SenderXPath`, data.SenderXPath.ValueString())
	}
	if !data.ReceiverXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReceiverXPath`, data.ReceiverXPath.ValueString())
	}
	if !data.DocumentIdxPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentIDXPath`, data.DocumentIdxPath.ValueString())
	}
	if !data.DateTimeXPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DateTimeXPath`, data.DateTimeXPath.ValueString())
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
		data.SenderXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReceiverXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentIDXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocumentIdxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentIdxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DateTimeXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DateTimeXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DateTimeXPath = types.StringNull()
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
	if value := res.Get(pathRoot + `SenderXPath`); value.Exists() && !data.SenderXPath.IsNull() {
		data.SenderXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverXPath`); value.Exists() && !data.ReceiverXPath.IsNull() {
		data.ReceiverXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverXPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentIDXPath`); value.Exists() && !data.DocumentIdxPath.IsNull() {
		data.DocumentIdxPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentIdxPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `DateTimeXPath`); value.Exists() && !data.DateTimeXPath.IsNull() {
		data.DateTimeXPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DateTimeXPath = types.StringNull()
	}
}
