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

type JSONSettings struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	JsonMaxNestingDepth types.Int64                 `tfsdk:"json_max_nesting_depth"`
	JsonMaxLabelLength  types.Int64                 `tfsdk:"json_max_label_length"`
	JsonMaxValueLength  types.Int64                 `tfsdk:"json_max_value_length"`
	JsonMaxNumberLength types.Int64                 `tfsdk:"json_max_number_length"`
	JsonDocumentSize    types.Int64                 `tfsdk:"json_document_size"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget      types.String                `tfsdk:"provider_target"`
}

var JSONSettingsObjectType = map[string]attr.Type{
	"provider_target":        types.StringType,
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"json_max_nesting_depth": types.Int64Type,
	"json_max_label_length":  types.Int64Type,
	"json_max_value_length":  types.Int64Type,
	"json_max_number_length": types.Int64Type,
	"json_document_size":     types.Int64Type,
	"dependency_actions":     actions.ActionsListType,
}

func (data JSONSettings) GetPath() string {
	rest_path := "/mgmt/config/{domain}/JSONSettings"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data JSONSettings) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.JsonMaxNestingDepth.IsNull() {
		return false
	}
	if !data.JsonMaxLabelLength.IsNull() {
		return false
	}
	if !data.JsonMaxValueLength.IsNull() {
		return false
	}
	if !data.JsonMaxNumberLength.IsNull() {
		return false
	}
	if !data.JsonDocumentSize.IsNull() {
		return false
	}
	return true
}

func (data JSONSettings) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.JsonMaxNestingDepth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONMaxNestingDepth`, data.JsonMaxNestingDepth.ValueInt64())
	}
	if !data.JsonMaxLabelLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONMaxLabelLength`, data.JsonMaxLabelLength.ValueInt64())
	}
	if !data.JsonMaxValueLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONMaxValueLength`, data.JsonMaxValueLength.ValueInt64())
	}
	if !data.JsonMaxNumberLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONMaxNumberLength`, data.JsonMaxNumberLength.ValueInt64())
	}
	if !data.JsonDocumentSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONDocumentSize`, data.JsonDocumentSize.ValueInt64())
	}
	return body
}

func (data *JSONSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JSONMaxNestingDepth`); value.Exists() {
		data.JsonMaxNestingDepth = types.Int64Value(value.Int())
	} else {
		data.JsonMaxNestingDepth = types.Int64Value(64)
	}
	if value := res.Get(pathRoot + `JSONMaxLabelLength`); value.Exists() {
		data.JsonMaxLabelLength = types.Int64Value(value.Int())
	} else {
		data.JsonMaxLabelLength = types.Int64Value(256)
	}
	if value := res.Get(pathRoot + `JSONMaxValueLength`); value.Exists() {
		data.JsonMaxValueLength = types.Int64Value(value.Int())
	} else {
		data.JsonMaxValueLength = types.Int64Value(8192)
	}
	if value := res.Get(pathRoot + `JSONMaxNumberLength`); value.Exists() {
		data.JsonMaxNumberLength = types.Int64Value(value.Int())
	} else {
		data.JsonMaxNumberLength = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `JSONDocumentSize`); value.Exists() {
		data.JsonDocumentSize = types.Int64Value(value.Int())
	} else {
		data.JsonDocumentSize = types.Int64Value(4194304)
	}
}

func (data *JSONSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JSONMaxNestingDepth`); value.Exists() && !data.JsonMaxNestingDepth.IsNull() {
		data.JsonMaxNestingDepth = types.Int64Value(value.Int())
	} else if data.JsonMaxNestingDepth.ValueInt64() != 64 {
		data.JsonMaxNestingDepth = types.Int64Null()
	}
	if value := res.Get(pathRoot + `JSONMaxLabelLength`); value.Exists() && !data.JsonMaxLabelLength.IsNull() {
		data.JsonMaxLabelLength = types.Int64Value(value.Int())
	} else if data.JsonMaxLabelLength.ValueInt64() != 256 {
		data.JsonMaxLabelLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `JSONMaxValueLength`); value.Exists() && !data.JsonMaxValueLength.IsNull() {
		data.JsonMaxValueLength = types.Int64Value(value.Int())
	} else if data.JsonMaxValueLength.ValueInt64() != 8192 {
		data.JsonMaxValueLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `JSONMaxNumberLength`); value.Exists() && !data.JsonMaxNumberLength.IsNull() {
		data.JsonMaxNumberLength = types.Int64Value(value.Int())
	} else if data.JsonMaxNumberLength.ValueInt64() != 128 {
		data.JsonMaxNumberLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `JSONDocumentSize`); value.Exists() && !data.JsonDocumentSize.IsNull() {
		data.JsonDocumentSize = types.Int64Value(value.Int())
	} else if data.JsonDocumentSize.ValueInt64() != 4194304 {
		data.JsonDocumentSize = types.Int64Null()
	}
}
