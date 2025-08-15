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

type CryptoIdentCred struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	Key               types.String                `tfsdk:"key"`
	Certificate       types.String                `tfsdk:"certificate"`
	Ca                types.List                  `tfsdk:"ca"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CryptoIdentCredObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"key":                types.StringType,
	"certificate":        types.StringType,
	"ca":                 types.ListType{ElemType: types.StringType},
	"dependency_actions": actions.ActionsListType,
}

func (data CryptoIdentCred) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoIdentCred"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoIdentCred) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Key.IsNull() {
		return false
	}
	if !data.Certificate.IsNull() {
		return false
	}
	if !data.Ca.IsNull() {
		return false
	}
	return true
}

func (data CryptoIdentCred) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Key.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Key`, data.Key.ValueString())
	}
	if !data.Certificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Certificate`, data.Certificate.ValueString())
	}
	if !data.Ca.IsNull() {
		var values []string
		data.Ca.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`CA`+".-1", map[string]string{"value": val})
		}
	}
	return body
}

func (data *CryptoIdentCred) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `CA`); value.Exists() {
		data.Ca = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ca = types.ListNull(types.StringType)
	}
}

func (data *CryptoIdentCred) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && !data.Key.IsNull() {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `CA`); value.Exists() && !data.Ca.IsNull() {
		data.Ca = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ca = types.ListNull(types.StringType)
	}
}
