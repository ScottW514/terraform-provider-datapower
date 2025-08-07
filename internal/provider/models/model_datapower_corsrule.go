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

type CORSRule struct {
	Id                types.String             `tfsdk:"id"`
	AppDomain         types.String             `tfsdk:"app_domain"`
	UserSummary       types.String             `tfsdk:"user_summary"`
	AllowOrigin       types.List               `tfsdk:"allow_origin"`
	AllowCredentials  types.Bool               `tfsdk:"allow_credentials"`
	ExposeHeaders     *DmCORSRuleExposeHeaders `tfsdk:"expose_headers"`
	DependencyActions []*actions.Action        `tfsdk:"dependency_actions"`
}

var CORSRuleObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"allow_origin":       types.ListType{ElemType: types.StringType},
	"allow_credentials":  types.BoolType,
	"expose_headers":     types.ObjectType{AttrTypes: DmCORSRuleExposeHeadersObjectType},
	"dependency_actions": actions.ActionsListType,
}

func (data CORSRule) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CORSRule"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CORSRule) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AllowOrigin.IsNull() {
		return false
	}
	if !data.AllowCredentials.IsNull() {
		return false
	}
	if data.ExposeHeaders != nil {
		if !data.ExposeHeaders.IsNull() {
			return false
		}
	}
	return true
}

func (data CORSRule) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AllowOrigin.IsNull() {
		var values []string
		data.AllowOrigin.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`AllowOrigin`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AllowCredentials.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCredentials`, tfutils.StringFromBool(data.AllowCredentials, ""))
	}
	if data.ExposeHeaders != nil {
		if !data.ExposeHeaders.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ExposeHeaders`, data.ExposeHeaders.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *CORSRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AllowOrigin`); value.Exists() {
		data.AllowOrigin = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AllowOrigin = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AllowCredentials`); value.Exists() {
		data.AllowCredentials = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCredentials = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExposeHeaders`); value.Exists() {
		data.ExposeHeaders = &DmCORSRuleExposeHeaders{}
		data.ExposeHeaders.FromBody(ctx, "", value)
	} else {
		data.ExposeHeaders = nil
	}
}

func (data *CORSRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AllowOrigin`); value.Exists() && !data.AllowOrigin.IsNull() {
		data.AllowOrigin = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AllowOrigin = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AllowCredentials`); value.Exists() && !data.AllowCredentials.IsNull() {
		data.AllowCredentials = tfutils.BoolFromString(value.String())
	} else if data.AllowCredentials.ValueBool() {
		data.AllowCredentials = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExposeHeaders`); value.Exists() {
		data.ExposeHeaders.UpdateFromBody(ctx, "", value)
	} else {
		data.ExposeHeaders = nil
	}
}
