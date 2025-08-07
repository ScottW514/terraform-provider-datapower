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

type ODR struct {
	Enabled             types.Bool        `tfsdk:"enabled"`
	UserSummary         types.String      `tfsdk:"user_summary"`
	OdrServerName       types.String      `tfsdk:"odr_server_name"`
	OdrConnectorGroups  types.List        `tfsdk:"odr_connector_groups"`
	OdrCustomProperties types.List        `tfsdk:"odr_custom_properties"`
	DependencyActions   []*actions.Action `tfsdk:"dependency_actions"`
}

var ODRObjectType = map[string]attr.Type{
	"enabled":               types.BoolType,
	"user_summary":          types.StringType,
	"odr_server_name":       types.StringType,
	"odr_connector_groups":  types.ListType{ElemType: types.StringType},
	"odr_custom_properties": types.ListType{ElemType: types.ObjectType{AttrTypes: DmODRPropertyObjectType}},
	"dependency_actions":    actions.ActionsListType,
}

func (data ODR) GetPath() string {
	rest_path := "/mgmt/config/default/ODR/ODRInstance"
	return rest_path
}

func (data ODR) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.OdrServerName.IsNull() {
		return false
	}
	if !data.OdrConnectorGroups.IsNull() {
		return false
	}
	if !data.OdrCustomProperties.IsNull() {
		return false
	}
	return true
}

func (data ODR) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "ODR.name", path.Base("/mgmt/config/default/ODR/ODRInstance"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.OdrServerName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OdrServerName`, data.OdrServerName.ValueString())
	}
	if !data.OdrConnectorGroups.IsNull() {
		var values []string
		data.OdrConnectorGroups.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`ODRConnectorGroups`+".-1", map[string]string{"value": val})
		}
	}
	if !data.OdrCustomProperties.IsNull() {
		var values []DmODRProperty
		data.OdrCustomProperties.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ODRCustomProperties`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *ODR) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OdrServerName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OdrServerName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OdrServerName = types.StringValue("dp_set")
	}
	if value := res.Get(pathRoot + `ODRConnectorGroups`); value.Exists() {
		data.OdrConnectorGroups = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.OdrConnectorGroups = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ODRCustomProperties`); value.Exists() {
		l := []DmODRProperty{}
		if value := res.Get(`ODRCustomProperties`); value.Exists() {
			for _, v := range value.Array() {
				item := DmODRProperty{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OdrCustomProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRPropertyObjectType}, l)
		} else {
			data.OdrCustomProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRPropertyObjectType})
		}
	} else {
		data.OdrCustomProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRPropertyObjectType})
	}
}

func (data *ODR) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OdrServerName`); value.Exists() && !data.OdrServerName.IsNull() {
		data.OdrServerName = tfutils.ParseStringFromGJSON(value)
	} else if data.OdrServerName.ValueString() != "dp_set" {
		data.OdrServerName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ODRConnectorGroups`); value.Exists() && !data.OdrConnectorGroups.IsNull() {
		data.OdrConnectorGroups = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.OdrConnectorGroups = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ODRCustomProperties`); value.Exists() && !data.OdrCustomProperties.IsNull() {
		l := []DmODRProperty{}
		for _, v := range value.Array() {
			item := DmODRProperty{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OdrCustomProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRPropertyObjectType}, l)
		} else {
			data.OdrCustomProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRPropertyObjectType})
		}
	} else {
		data.OdrCustomProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRPropertyObjectType})
	}
}
