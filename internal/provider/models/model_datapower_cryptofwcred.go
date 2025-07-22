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

type CryptoFWCred struct {
	Id              types.String `tfsdk:"id"`
	AppDomain       types.String `tfsdk:"app_domain"`
	PrivateKey      types.List   `tfsdk:"private_key"`
	SharedSecretKey types.List   `tfsdk:"shared_secret_key"`
	Certificate     types.List   `tfsdk:"certificate"`
}

var CryptoFWCredObjectType = map[string]attr.Type{
	"id":                types.StringType,
	"app_domain":        types.StringType,
	"private_key":       types.ListType{ElemType: types.StringType},
	"shared_secret_key": types.ListType{ElemType: types.StringType},
	"certificate":       types.ListType{ElemType: types.StringType},
}

func (data CryptoFWCred) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoFWCred"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoFWCred) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.PrivateKey.IsNull() {
		return false
	}
	if !data.SharedSecretKey.IsNull() {
		return false
	}
	if !data.Certificate.IsNull() {
		return false
	}
	return true
}

func (data CryptoFWCred) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.PrivateKey.IsNull() {
		var values []string
		data.PrivateKey.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`PrivateKey`+".-1", map[string]string{"value": val})
		}
	}
	if !data.SharedSecretKey.IsNull() {
		var values []string
		data.SharedSecretKey.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`SharedSecretKey`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Certificate.IsNull() {
		var values []string
		data.Certificate.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Certificate`+".-1", map[string]string{"value": val})
		}
	}
	return body
}

func (data *CryptoFWCred) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivateKey`); value.Exists() {
		data.PrivateKey = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PrivateKey = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SharedSecretKey`); value.Exists() {
		data.SharedSecretKey = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SharedSecretKey = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() {
		data.Certificate = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Certificate = types.ListNull(types.StringType)
	}
}

func (data *CryptoFWCred) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivateKey`); value.Exists() && !data.PrivateKey.IsNull() {
		data.PrivateKey = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PrivateKey = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SharedSecretKey`); value.Exists() && !data.SharedSecretKey.IsNull() {
		data.SharedSecretKey = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SharedSecretKey = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Certificate = types.ListNull(types.StringType)
	}
}
