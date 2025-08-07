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
	"path"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Statistics struct {
	AppDomain         types.String                `tfsdk:"app_domain"`
	Enabled           types.Bool                  `tfsdk:"enabled"`
	LoadInterval      types.Int64                 `tfsdk:"load_interval"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var StatisticsObjectType = map[string]attr.Type{
	"app_domain":         types.StringType,
	"enabled":            types.BoolType,
	"load_interval":      types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data Statistics) GetPath() string {
	rest_path := "/mgmt/config/{domain}/Statistics/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data Statistics) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.LoadInterval.IsNull() {
		return false
	}
	return true
}

func (data Statistics) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "Statistics.name", path.Base("/mgmt/config/{domain}/Statistics/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.LoadInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoadInterval`, data.LoadInterval.ValueInt64())
	}
	return body
}

func (data *Statistics) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoadInterval`); value.Exists() {
		data.LoadInterval = types.Int64Value(value.Int())
	} else {
		data.LoadInterval = types.Int64Value(1000)
	}
}

func (data *Statistics) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoadInterval`); value.Exists() && !data.LoadInterval.IsNull() {
		data.LoadInterval = types.Int64Value(value.Int())
	} else if data.LoadInterval.ValueInt64() != 1000 {
		data.LoadInterval = types.Int64Null()
	}
}
