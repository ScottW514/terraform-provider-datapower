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

type CryptoKey struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Filename          types.String                `tfsdk:"filename"`
	Alias             types.String                `tfsdk:"alias"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var CryptoKeyObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"filename":           types.StringType,
	"alias":              types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data CryptoKey) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoKey"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoKey) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Filename.IsNull() {
		return false
	}
	if !data.Alias.IsNull() {
		return false
	}
	return true
}

func (data CryptoKey) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Filename.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Filename`, data.Filename.ValueString())
	}
	if !data.Alias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Alias`, data.Alias.ValueString())
	}
	return body
}

func (data *CryptoKey) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Filename`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Filename = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Filename = types.StringNull()
	}
	if value := res.Get(pathRoot + `Alias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Alias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Alias = types.StringNull()
	}
}

func (data *CryptoKey) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Filename`); value.Exists() && !data.Filename.IsNull() {
		data.Filename = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Filename = types.StringNull()
	}
	if value := res.Get(pathRoot + `Alias`); value.Exists() && !data.Alias.IsNull() {
		data.Alias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Alias = types.StringNull()
	}
}
