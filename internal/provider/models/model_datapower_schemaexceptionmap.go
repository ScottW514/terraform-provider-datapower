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

type SchemaExceptionMap struct {
	Id                   types.String      `tfsdk:"id"`
	AppDomain            types.String      `tfsdk:"app_domain"`
	OriginalSchemaUrl    types.String      `tfsdk:"original_schema_url"`
	SchemaExceptionRules types.List        `tfsdk:"schema_exception_rules"`
	UserSummary          types.String      `tfsdk:"user_summary"`
	ObjectActions        []*actions.Action `tfsdk:"object_actions"`
}

var SchemaExceptionMapObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"original_schema_url":    types.StringType,
	"schema_exception_rules": types.ListType{ElemType: types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType}},
	"user_summary":           types.StringType,
	"object_actions":         actions.ActionsListType,
}

func (data SchemaExceptionMap) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SchemaExceptionMap"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SchemaExceptionMap) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.OriginalSchemaUrl.IsNull() {
		return false
	}
	if !data.SchemaExceptionRules.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data SchemaExceptionMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.OriginalSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OriginalSchemaURL`, data.OriginalSchemaUrl.ValueString())
	}
	if !data.SchemaExceptionRules.IsNull() {
		var values []DmSchemaExceptionRule
		data.SchemaExceptionRules.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`SchemaExceptionRules`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *SchemaExceptionMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `OriginalSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OriginalSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OriginalSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaExceptionRules`); value.Exists() {
		l := []DmSchemaExceptionRule{}
		if value := res.Get(`SchemaExceptionRules`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSchemaExceptionRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SchemaExceptionRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType}, l)
		} else {
			data.SchemaExceptionRules = types.ListNull(types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType})
		}
	} else {
		data.SchemaExceptionRules = types.ListNull(types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *SchemaExceptionMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `OriginalSchemaURL`); value.Exists() && !data.OriginalSchemaUrl.IsNull() {
		data.OriginalSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OriginalSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaExceptionRules`); value.Exists() && !data.SchemaExceptionRules.IsNull() {
		l := []DmSchemaExceptionRule{}
		for _, v := range value.Array() {
			item := DmSchemaExceptionRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SchemaExceptionRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType}, l)
		} else {
			data.SchemaExceptionRules = types.ListNull(types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType})
		}
	} else {
		data.SchemaExceptionRules = types.ListNull(types.ObjectType{AttrTypes: DmSchemaExceptionRuleObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
