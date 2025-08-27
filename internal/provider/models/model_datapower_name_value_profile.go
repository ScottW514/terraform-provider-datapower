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

type NameValueProfile struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	MaxAttributes          types.Int64                 `tfsdk:"max_attributes"`
	MaxAggregateSize       types.Int64                 `tfsdk:"max_aggregate_size"`
	MaxNameSize            types.Int64                 `tfsdk:"max_name_size"`
	MaxValueSize           types.Int64                 `tfsdk:"max_value_size"`
	ValidationList         types.List                  `tfsdk:"validation_list"`
	DefaultFixup           types.String                `tfsdk:"default_fixup"`
	DefaultMapValue        types.String                `tfsdk:"default_map_value"`
	DefaultXss             types.Bool                  `tfsdk:"default_xss"`
	NoMatchXssPatternsFile types.String                `tfsdk:"no_match_xss_patterns_file"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var NameValueProfileNoMatchXSSPatternsFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "default_xss",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var NameValueProfileDefaultMapValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "default_fixup",
	AttrType:    "String",
	AttrDefault: "strip",
	Value:       []string{"set"},
}
var NameValueProfileNoMatchXSSPatternsFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "default_xss",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var NameValueProfileObjectType = map[string]attr.Type{
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"user_summary":               types.StringType,
	"max_attributes":             types.Int64Type,
	"max_aggregate_size":         types.Int64Type,
	"max_name_size":              types.Int64Type,
	"max_value_size":             types.Int64Type,
	"validation_list":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmValidationTypeObjectType}},
	"default_fixup":              types.StringType,
	"default_map_value":          types.StringType,
	"default_xss":                types.BoolType,
	"no_match_xss_patterns_file": types.StringType,
	"dependency_actions":         actions.ActionsListType,
}

func (data NameValueProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/NameValueProfile"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data NameValueProfile) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MaxAttributes.IsNull() {
		return false
	}
	if !data.MaxAggregateSize.IsNull() {
		return false
	}
	if !data.MaxNameSize.IsNull() {
		return false
	}
	if !data.MaxValueSize.IsNull() {
		return false
	}
	if !data.ValidationList.IsNull() {
		return false
	}
	if !data.DefaultFixup.IsNull() {
		return false
	}
	if !data.DefaultMapValue.IsNull() {
		return false
	}
	if !data.DefaultXss.IsNull() {
		return false
	}
	if !data.NoMatchXssPatternsFile.IsNull() {
		return false
	}
	return true
}

func (data NameValueProfile) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MaxAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxAttributes`, data.MaxAttributes.ValueInt64())
	}
	if !data.MaxAggregateSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxAggregateSize`, data.MaxAggregateSize.ValueInt64())
	}
	if !data.MaxNameSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxNameSize`, data.MaxNameSize.ValueInt64())
	}
	if !data.MaxValueSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxValueSize`, data.MaxValueSize.ValueInt64())
	}
	if !data.ValidationList.IsNull() {
		var dataValues []DmValidationType
		data.ValidationList.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ValidationList`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.DefaultFixup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultFixup`, data.DefaultFixup.ValueString())
	}
	if !data.DefaultMapValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultMapValue`, data.DefaultMapValue.ValueString())
	}
	if !data.DefaultXss.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultXSS`, tfutils.StringFromBool(data.DefaultXss, ""))
	}
	if !data.NoMatchXssPatternsFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoMatchXSSPatternsFile`, data.NoMatchXssPatternsFile.ValueString())
	}
	return body
}

func (data *NameValueProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MaxAttributes`); value.Exists() {
		data.MaxAttributes = types.Int64Value(value.Int())
	} else {
		data.MaxAttributes = types.Int64Value(256)
	}
	if value := res.Get(pathRoot + `MaxAggregateSize`); value.Exists() {
		data.MaxAggregateSize = types.Int64Value(value.Int())
	} else {
		data.MaxAggregateSize = types.Int64Value(128000)
	}
	if value := res.Get(pathRoot + `MaxNameSize`); value.Exists() {
		data.MaxNameSize = types.Int64Value(value.Int())
	} else {
		data.MaxNameSize = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `MaxValueSize`); value.Exists() {
		data.MaxValueSize = types.Int64Value(value.Int())
	} else {
		data.MaxValueSize = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ValidationList`); value.Exists() {
		l := []DmValidationType{}
		if value := res.Get(`ValidationList`); value.Exists() {
			for _, v := range value.Array() {
				item := DmValidationType{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ValidationList, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmValidationTypeObjectType}, l)
		} else {
			data.ValidationList = types.ListNull(types.ObjectType{AttrTypes: DmValidationTypeObjectType})
		}
	} else {
		data.ValidationList = types.ListNull(types.ObjectType{AttrTypes: DmValidationTypeObjectType})
	}
	if value := res.Get(pathRoot + `DefaultFixup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultFixup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultFixup = types.StringValue("strip")
	}
	if value := res.Get(pathRoot + `DefaultMapValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultMapValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultMapValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultXSS`); value.Exists() {
		data.DefaultXss = tfutils.BoolFromString(value.String())
	} else {
		data.DefaultXss = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoMatchXSSPatternsFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NoMatchXssPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NoMatchXssPatternsFile = types.StringValue("store:///XSS-Patterns.xml")
	}
}

func (data *NameValueProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MaxAttributes`); value.Exists() && !data.MaxAttributes.IsNull() {
		data.MaxAttributes = types.Int64Value(value.Int())
	} else if data.MaxAttributes.ValueInt64() != 256 {
		data.MaxAttributes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxAggregateSize`); value.Exists() && !data.MaxAggregateSize.IsNull() {
		data.MaxAggregateSize = types.Int64Value(value.Int())
	} else if data.MaxAggregateSize.ValueInt64() != 128000 {
		data.MaxAggregateSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxNameSize`); value.Exists() && !data.MaxNameSize.IsNull() {
		data.MaxNameSize = types.Int64Value(value.Int())
	} else if data.MaxNameSize.ValueInt64() != 512 {
		data.MaxNameSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxValueSize`); value.Exists() && !data.MaxValueSize.IsNull() {
		data.MaxValueSize = types.Int64Value(value.Int())
	} else if data.MaxValueSize.ValueInt64() != 1024 {
		data.MaxValueSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ValidationList`); value.Exists() && !data.ValidationList.IsNull() {
		l := []DmValidationType{}
		for _, v := range value.Array() {
			item := DmValidationType{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ValidationList, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmValidationTypeObjectType}, l)
		} else {
			data.ValidationList = types.ListNull(types.ObjectType{AttrTypes: DmValidationTypeObjectType})
		}
	} else {
		data.ValidationList = types.ListNull(types.ObjectType{AttrTypes: DmValidationTypeObjectType})
	}
	if value := res.Get(pathRoot + `DefaultFixup`); value.Exists() && !data.DefaultFixup.IsNull() {
		data.DefaultFixup = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultFixup.ValueString() != "strip" {
		data.DefaultFixup = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultMapValue`); value.Exists() && !data.DefaultMapValue.IsNull() {
		data.DefaultMapValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultMapValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultXSS`); value.Exists() && !data.DefaultXss.IsNull() {
		data.DefaultXss = tfutils.BoolFromString(value.String())
	} else if data.DefaultXss.ValueBool() {
		data.DefaultXss = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoMatchXSSPatternsFile`); value.Exists() && !data.NoMatchXssPatternsFile.IsNull() {
		data.NoMatchXssPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else if data.NoMatchXssPatternsFile.ValueString() != "store:///XSS-Patterns.xml" {
		data.NoMatchXssPatternsFile = types.StringNull()
	}
}
