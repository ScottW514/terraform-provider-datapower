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

type FilterAction struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Type              types.String                `tfsdk:"type"`
	LogLevel          types.String                `tfsdk:"log_level"`
	BlockInterval     types.Int64                 `tfsdk:"block_interval"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var FilterActionBlockIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "notify",
	Value:       []string{"notify"},
}

var FilterActionObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"type":               types.StringType,
	"log_level":          types.StringType,
	"block_interval":     types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data FilterAction) GetPath() string {
	rest_path := "/mgmt/config/{domain}/FilterAction"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data FilterAction) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.LogLevel.IsNull() {
		return false
	}
	if !data.BlockInterval.IsNull() {
		return false
	}
	return true
}

func (data FilterAction) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.LogLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogLevel`, data.LogLevel.ValueString())
	}
	if !data.BlockInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BlockInterval`, data.BlockInterval.ValueInt64())
	}
	return body
}

func (data *FilterAction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("notify")
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogLevel = types.StringValue("debug")
	}
	if value := res.Get(pathRoot + `BlockInterval`); value.Exists() {
		data.BlockInterval = types.Int64Value(value.Int())
	} else {
		data.BlockInterval = types.Int64Null()
	}
}

func (data *FilterAction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "notify" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && !data.LogLevel.IsNull() {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.LogLevel.ValueString() != "debug" {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `BlockInterval`); value.Exists() && !data.BlockInterval.IsNull() {
		data.BlockInterval = types.Int64Value(value.Int())
	} else {
		data.BlockInterval = types.Int64Null()
	}
}
