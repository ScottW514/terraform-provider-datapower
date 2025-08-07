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

type RateLimitConfiguration struct {
	AppDomain         types.String      `tfsdk:"app_domain"`
	Enabled           types.Bool        `tfsdk:"enabled"`
	Parameters        types.List        `tfsdk:"parameters"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var RateLimitConfigurationObjectType = map[string]attr.Type{
	"app_domain":         types.StringType,
	"enabled":            types.BoolType,
	"parameters":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data RateLimitConfiguration) GetPath() string {
	rest_path := "/mgmt/config/{domain}/RateLimitConfiguration/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data RateLimitConfiguration) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.Parameters.IsNull() {
		return false
	}
	return true
}

func (data RateLimitConfiguration) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "RateLimitConfiguration.name", path.Base("/mgmt/config/{domain}/RateLimitConfiguration/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.Parameters.IsNull() {
		var values []DmRateLimitConfigurationNameValuePair
		data.Parameters.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Parameters`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *RateLimitConfiguration) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Parameters`); value.Exists() {
		l := []DmRateLimitConfigurationNameValuePair{}
		if value := res.Get(`Parameters`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRateLimitConfigurationNameValuePair{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType}, l)
		} else {
			data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType})
		}
	} else {
		data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType})
	}
}

func (data *RateLimitConfiguration) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Parameters`); value.Exists() && !data.Parameters.IsNull() {
		l := []DmRateLimitConfigurationNameValuePair{}
		for _, v := range value.Array() {
			item := DmRateLimitConfigurationNameValuePair{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Parameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType}, l)
		} else {
			data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType})
		}
	} else {
		data.Parameters = types.ListNull(types.ObjectType{AttrTypes: DmRateLimitConfigurationNameValuePairObjectType})
	}
}
