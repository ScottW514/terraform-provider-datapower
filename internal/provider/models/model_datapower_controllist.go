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

type ControlList struct {
	Id              types.String `tfsdk:"id"`
	AppDomain       types.String `tfsdk:"app_domain"`
	UserSummary     types.String `tfsdk:"user_summary"`
	Type            types.String `tfsdk:"type"`
	Value           types.List   `tfsdk:"value"`
	CaseInsensitive types.Bool   `tfsdk:"case_insensitive"`
}

var ControlListObjectType = map[string]attr.Type{
	"id":               types.StringType,
	"app_domain":       types.StringType,
	"user_summary":     types.StringType,
	"type":             types.StringType,
	"value":            types.ListType{ElemType: types.StringType},
	"case_insensitive": types.BoolType,
}

func (data ControlList) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ControlList"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ControlList) IsNull() bool {
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
	if !data.Value.IsNull() {
		return false
	}
	if !data.CaseInsensitive.IsNull() {
		return false
	}
	return true
}

func (data ControlList) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Value.IsNull() {
		var values []string
		data.Value.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Value`+".-1", map[string]string{"value": val})
		}
	}
	if !data.CaseInsensitive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CaseInsensitive`, tfutils.StringFromBool(data.CaseInsensitive, ""))
	}
	return body
}

func (data *ControlList) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() {
		data.Value = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Value = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CaseInsensitive`); value.Exists() {
		data.CaseInsensitive = tfutils.BoolFromString(value.String())
	} else {
		data.CaseInsensitive = types.BoolNull()
	}
}

func (data *ControlList) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Value = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CaseInsensitive`); value.Exists() && !data.CaseInsensitive.IsNull() {
		data.CaseInsensitive = tfutils.BoolFromString(value.String())
	} else if data.CaseInsensitive.ValueBool() {
		data.CaseInsensitive = types.BoolNull()
	}
}
