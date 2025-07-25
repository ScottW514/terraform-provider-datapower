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

type OperationRateLimit struct {
	Id                types.String `tfsdk:"id"`
	AppDomain         types.String `tfsdk:"app_domain"`
	UserSummary       types.String `tfsdk:"user_summary"`
	Operation         types.String `tfsdk:"operation"`
	UseRateLimitGroup types.Bool   `tfsdk:"use_rate_limit_group"`
	RateLimit         types.List   `tfsdk:"rate_limit"`
	RateLimitGroup    types.String `tfsdk:"rate_limit_group"`
}

var OperationRateLimitObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"operation":            types.StringType,
	"use_rate_limit_group": types.BoolType,
	"rate_limit":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"rate_limit_group":     types.StringType,
}

func (data OperationRateLimit) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OperationRateLimit"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OperationRateLimit) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Operation.IsNull() {
		return false
	}
	if !data.UseRateLimitGroup.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.RateLimitGroup.IsNull() {
		return false
	}
	return true
}

func (data OperationRateLimit) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Operation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Operation`, data.Operation.ValueString())
	}
	if !data.UseRateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseRateLimitGroup`, tfutils.StringFromBool(data.UseRateLimitGroup, ""))
	}
	if !data.RateLimit.IsNull() {
		var values []DmAPIRateLimit
		data.RateLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`RateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimitGroup`, data.RateLimitGroup.ValueString())
	}
	return body
}

func (data *OperationRateLimit) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Operation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Operation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Operation = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		l := []DmAPIRateLimit{}
		if value := res.Get(`RateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
}

func (data *OperationRateLimit) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Operation`); value.Exists() && !data.Operation.IsNull() {
		data.Operation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Operation = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() && !data.UseRateLimitGroup.IsNull() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else if data.UseRateLimitGroup.ValueBool() {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		l := []DmAPIRateLimit{}
		for _, v := range value.Array() {
			item := DmAPIRateLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && !data.RateLimitGroup.IsNull() {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
}
