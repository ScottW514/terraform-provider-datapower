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

type SSHDomainClientProfile struct {
	AppDomain         types.String                `tfsdk:"app_domain"`
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Ciphers           types.List                  `tfsdk:"ciphers"`
	KexAlg            types.List                  `tfsdk:"kex_alg"`
	MacAlg            types.List                  `tfsdk:"mac_alg"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var SSHDomainClientProfileObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"app_domain":         types.StringType,
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"ciphers":            types.ListType{ElemType: types.StringType},
	"kex_alg":            types.ListType{ElemType: types.StringType},
	"mac_alg":            types.ListType{ElemType: types.StringType},
	"dependency_actions": actions.ActionsListType,
}

func (data SSHDomainClientProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSHDomainClientProfile/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSHDomainClientProfile) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Ciphers.IsNull() {
		return false
	}
	if !data.KexAlg.IsNull() {
		return false
	}
	if !data.MacAlg.IsNull() {
		return false
	}
	return true
}
func (data *SSHDomainClientProfile) ToDefault() {
	data.Enabled = types.BoolValue(true)
	data.UserSummary = types.StringNull()
	data.Ciphers = types.ListNull(types.StringType)
	data.KexAlg = types.ListNull(types.StringType)
	data.MacAlg = types.ListNull(types.StringType)
}

func (data SSHDomainClientProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "SSHDomainClientProfile.name", path.Base("/mgmt/config/{domain}/SSHDomainClientProfile/default"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Ciphers.IsNull() {
		var dataValues []string
		data.Ciphers.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Ciphers`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
	}
	if !data.KexAlg.IsNull() {
		var dataValues []string
		data.KexAlg.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`KEXAlg`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`KEXAlg`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`KEXAlg`, "[]")
	}
	if !data.MacAlg.IsNull() {
		var dataValues []string
		data.MacAlg.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`MACAlg`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`MACAlg`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`MACAlg`, "[]")
	}
	return body
}

func (data *SSHDomainClientProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `KEXAlg`); value.Exists() {
		data.KexAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.KexAlg = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MACAlg`); value.Exists() {
		data.MacAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MacAlg = types.ListNull(types.StringType)
	}
}

func (data *SSHDomainClientProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() && !data.Ciphers.IsNull() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `KEXAlg`); value.Exists() && !data.KexAlg.IsNull() {
		data.KexAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.KexAlg = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MACAlg`); value.Exists() && !data.MacAlg.IsNull() {
		data.MacAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MacAlg = types.ListNull(types.StringType)
	}
}
