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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RADIUSSettings struct {
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Id                types.String                `tfsdk:"id"`
	Timeout           types.Int64                 `tfsdk:"timeout"`
	Retries           types.Int64                 `tfsdk:"retries"`
	AaaServers        types.List                  `tfsdk:"aaa_servers"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var RADIUSSettingsObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"id":                 types.StringType,
	"timeout":            types.Int64Type,
	"retries":            types.Int64Type,
	"aaa_servers":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmRadiusServerObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data RADIUSSettings) GetPath() string {
	rest_path := "/mgmt/config/default/RADIUSSettings/RADIUS-Settings"
	return rest_path
}

func (data RADIUSSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Id.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.Retries.IsNull() {
		return false
	}
	if !data.AaaServers.IsNull() {
		return false
	}
	return true
}

func (data RADIUSSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "RADIUSSettings.name", path.Base("/mgmt/config/default/RADIUSSettings/RADIUS-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdgId`, data.Id.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.Retries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Retries`, data.Retries.ValueInt64())
	}
	if !data.AaaServers.IsNull() {
		var values []DmRadiusServer
		data.AaaServers.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AAAServers`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *RADIUSSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `IdgId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `Retries`); value.Exists() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `AAAServers`); value.Exists() {
		l := []DmRadiusServer{}
		if value := res.Get(`AAAServers`); value.Exists() {
			for _, v := range value.Array() {
				item := DmRadiusServer{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AaaServers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRadiusServerObjectType}, l)
		} else {
			data.AaaServers = types.ListNull(types.ObjectType{AttrTypes: DmRadiusServerObjectType})
		}
	} else {
		data.AaaServers = types.ListNull(types.ObjectType{AttrTypes: DmRadiusServerObjectType})
	}
}

func (data *RADIUSSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `IdgId`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 1000 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Retries`); value.Exists() && !data.Retries.IsNull() {
		data.Retries = types.Int64Value(value.Int())
	} else if data.Retries.ValueInt64() != 3 {
		data.Retries = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AAAServers`); value.Exists() && !data.AaaServers.IsNull() {
		l := []DmRadiusServer{}
		for _, v := range value.Array() {
			item := DmRadiusServer{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AaaServers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmRadiusServerObjectType}, l)
		} else {
			data.AaaServers = types.ListNull(types.ObjectType{AttrTypes: DmRadiusServerObjectType})
		}
	} else {
		data.AaaServers = types.ListNull(types.ObjectType{AttrTypes: DmRadiusServerObjectType})
	}
}
