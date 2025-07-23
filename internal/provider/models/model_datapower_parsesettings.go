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

type ParseSettings struct {
	Id                 types.String `tfsdk:"id"`
	AppDomain          types.String `tfsdk:"app_domain"`
	UserSummary        types.String `tfsdk:"user_summary"`
	DocumentType       types.String `tfsdk:"document_type"`
	DocumentSize       types.Int64  `tfsdk:"document_size"`
	NestingDepth       types.Int64  `tfsdk:"nesting_depth"`
	Width              types.Int64  `tfsdk:"width"`
	NameLength         types.Int64  `tfsdk:"name_length"`
	ValueLength        types.Int64  `tfsdk:"value_length"`
	UniquePrefixes     types.Int64  `tfsdk:"unique_prefixes"`
	UniqueNamespaces   types.Int64  `tfsdk:"unique_namespaces"`
	UniqueNames        types.Int64  `tfsdk:"unique_names"`
	NumberLength       types.Int64  `tfsdk:"number_length"`
	StrictUtf8Encoding types.Bool   `tfsdk:"strict_utf8_encoding"`
}

var ParseSettingsObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"document_type":        types.StringType,
	"document_size":        types.Int64Type,
	"nesting_depth":        types.Int64Type,
	"width":                types.Int64Type,
	"name_length":          types.Int64Type,
	"value_length":         types.Int64Type,
	"unique_prefixes":      types.Int64Type,
	"unique_namespaces":    types.Int64Type,
	"unique_names":         types.Int64Type,
	"number_length":        types.Int64Type,
	"strict_utf8_encoding": types.BoolType,
}

func (data ParseSettings) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ParseSettings"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ParseSettings) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.DocumentType.IsNull() {
		return false
	}
	if !data.DocumentSize.IsNull() {
		return false
	}
	if !data.NestingDepth.IsNull() {
		return false
	}
	if !data.Width.IsNull() {
		return false
	}
	if !data.NameLength.IsNull() {
		return false
	}
	if !data.ValueLength.IsNull() {
		return false
	}
	if !data.UniquePrefixes.IsNull() {
		return false
	}
	if !data.UniqueNamespaces.IsNull() {
		return false
	}
	if !data.UniqueNames.IsNull() {
		return false
	}
	if !data.NumberLength.IsNull() {
		return false
	}
	if !data.StrictUtf8Encoding.IsNull() {
		return false
	}
	return true
}

func (data ParseSettings) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.DocumentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentType`, data.DocumentType.ValueString())
	}
	if !data.DocumentSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentSize`, data.DocumentSize.ValueInt64())
	}
	if !data.NestingDepth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NestingDepth`, data.NestingDepth.ValueInt64())
	}
	if !data.Width.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Width`, data.Width.ValueInt64())
	}
	if !data.NameLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NameLength`, data.NameLength.ValueInt64())
	}
	if !data.ValueLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueLength`, data.ValueLength.ValueInt64())
	}
	if !data.UniquePrefixes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UniquePrefixes`, data.UniquePrefixes.ValueInt64())
	}
	if !data.UniqueNamespaces.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UniqueNamespaces`, data.UniqueNamespaces.ValueInt64())
	}
	if !data.UniqueNames.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UniqueNames`, data.UniqueNames.ValueInt64())
	}
	if !data.NumberLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NumberLength`, data.NumberLength.ValueInt64())
	}
	if !data.StrictUtf8Encoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StrictUTF8Encoding`, tfutils.StringFromBool(data.StrictUtf8Encoding, ""))
	}
	return body
}

func (data *ParseSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DocumentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocumentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentType = types.StringValue("detect")
	}
	if value := res.Get(pathRoot + `DocumentSize`); value.Exists() {
		data.DocumentSize = types.Int64Value(value.Int())
	} else {
		data.DocumentSize = types.Int64Value(4194304)
	}
	if value := res.Get(pathRoot + `NestingDepth`); value.Exists() {
		data.NestingDepth = types.Int64Value(value.Int())
	} else {
		data.NestingDepth = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `Width`); value.Exists() {
		data.Width = types.Int64Value(value.Int())
	} else {
		data.Width = types.Int64Value(4096)
	}
	if value := res.Get(pathRoot + `NameLength`); value.Exists() {
		data.NameLength = types.Int64Value(value.Int())
	} else {
		data.NameLength = types.Int64Value(256)
	}
	if value := res.Get(pathRoot + `ValueLength`); value.Exists() {
		data.ValueLength = types.Int64Value(value.Int())
	} else {
		data.ValueLength = types.Int64Value(8192)
	}
	if value := res.Get(pathRoot + `UniquePrefixes`); value.Exists() {
		data.UniquePrefixes = types.Int64Value(value.Int())
	} else {
		data.UniquePrefixes = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `UniqueNamespaces`); value.Exists() {
		data.UniqueNamespaces = types.Int64Value(value.Int())
	} else {
		data.UniqueNamespaces = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `UniqueNames`); value.Exists() {
		data.UniqueNames = types.Int64Value(value.Int())
	} else {
		data.UniqueNames = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `NumberLength`); value.Exists() {
		data.NumberLength = types.Int64Value(value.Int())
	} else {
		data.NumberLength = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `StrictUTF8Encoding`); value.Exists() {
		data.StrictUtf8Encoding = tfutils.BoolFromString(value.String())
	} else {
		data.StrictUtf8Encoding = types.BoolNull()
	}
}

func (data *ParseSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DocumentType`); value.Exists() && !data.DocumentType.IsNull() {
		data.DocumentType = tfutils.ParseStringFromGJSON(value)
	} else if data.DocumentType.ValueString() != "detect" {
		data.DocumentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentSize`); value.Exists() && !data.DocumentSize.IsNull() {
		data.DocumentSize = types.Int64Value(value.Int())
	} else if data.DocumentSize.ValueInt64() != 4194304 {
		data.DocumentSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `NestingDepth`); value.Exists() && !data.NestingDepth.IsNull() {
		data.NestingDepth = types.Int64Value(value.Int())
	} else if data.NestingDepth.ValueInt64() != 512 {
		data.NestingDepth = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Width`); value.Exists() && !data.Width.IsNull() {
		data.Width = types.Int64Value(value.Int())
	} else if data.Width.ValueInt64() != 4096 {
		data.Width = types.Int64Null()
	}
	if value := res.Get(pathRoot + `NameLength`); value.Exists() && !data.NameLength.IsNull() {
		data.NameLength = types.Int64Value(value.Int())
	} else if data.NameLength.ValueInt64() != 256 {
		data.NameLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ValueLength`); value.Exists() && !data.ValueLength.IsNull() {
		data.ValueLength = types.Int64Value(value.Int())
	} else if data.ValueLength.ValueInt64() != 8192 {
		data.ValueLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UniquePrefixes`); value.Exists() && !data.UniquePrefixes.IsNull() {
		data.UniquePrefixes = types.Int64Value(value.Int())
	} else if data.UniquePrefixes.ValueInt64() != 1024 {
		data.UniquePrefixes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UniqueNamespaces`); value.Exists() && !data.UniqueNamespaces.IsNull() {
		data.UniqueNamespaces = types.Int64Value(value.Int())
	} else if data.UniqueNamespaces.ValueInt64() != 1024 {
		data.UniqueNamespaces = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UniqueNames`); value.Exists() && !data.UniqueNames.IsNull() {
		data.UniqueNames = types.Int64Value(value.Int())
	} else if data.UniqueNames.ValueInt64() != 1024 {
		data.UniqueNames = types.Int64Null()
	}
	if value := res.Get(pathRoot + `NumberLength`); value.Exists() && !data.NumberLength.IsNull() {
		data.NumberLength = types.Int64Value(value.Int())
	} else if data.NumberLength.ValueInt64() != 128 {
		data.NumberLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `StrictUTF8Encoding`); value.Exists() && !data.StrictUtf8Encoding.IsNull() {
		data.StrictUtf8Encoding = tfutils.BoolFromString(value.String())
	} else if data.StrictUtf8Encoding.ValueBool() {
		data.StrictUtf8Encoding = types.BoolNull()
	}
}
