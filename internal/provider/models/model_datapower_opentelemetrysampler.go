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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type OpenTelemetrySampler struct {
	Id          types.String `tfsdk:"id"`
	AppDomain   types.String `tfsdk:"app_domain"`
	UserSummary types.String `tfsdk:"user_summary"`
	ParentBased types.Bool   `tfsdk:"parent_based"`
	Type        types.String `tfsdk:"type"`
	Ratio       types.Int64  `tfsdk:"ratio"`
}

var OpenTelemetrySamplerObjectType = map[string]attr.Type{
	"id":           types.StringType,
	"app_domain":   types.StringType,
	"user_summary": types.StringType,
	"parent_based": types.BoolType,
	"type":         types.StringType,
	"ratio":        types.Int64Type,
}

func (data OpenTelemetrySampler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OpenTelemetrySampler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OpenTelemetrySampler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ParentBased.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Ratio.IsNull() {
		return false
	}
	return true
}

func (data OpenTelemetrySampler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ParentBased.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParentBased`, tfutils.StringFromBool(data.ParentBased, ""))
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Ratio.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Ratio`, data.Ratio.ValueInt64())
	}
	return body
}

func (data *OpenTelemetrySampler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ParentBased`); value.Exists() {
		data.ParentBased = tfutils.BoolFromString(value.String())
	} else {
		data.ParentBased = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("always-on")
	}
	if value := res.Get(pathRoot + `Ratio`); value.Exists() {
		data.Ratio = types.Int64Value(value.Int())
	} else {
		data.Ratio = types.Int64Null()
	}
}

func (data *OpenTelemetrySampler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ParentBased`); value.Exists() && !data.ParentBased.IsNull() {
		data.ParentBased = tfutils.BoolFromString(value.String())
	} else if !data.ParentBased.ValueBool() {
		data.ParentBased = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "always-on" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Ratio`); value.Exists() && !data.Ratio.IsNull() {
		data.Ratio = types.Int64Value(value.Int())
	} else {
		data.Ratio = types.Int64Null()
	}
}
