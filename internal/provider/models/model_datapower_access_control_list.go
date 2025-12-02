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

type AccessControlList struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	AccessControlEntry types.List                  `tfsdk:"access_control_entry"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget     types.String                `tfsdk:"provider_target"`
}

var AccessControlListObjectType = map[string]attr.Type{
	"provider_target":      types.StringType,
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"access_control_entry": types.ListType{ElemType: types.ObjectType{AttrTypes: DmACEObjectType}},
	"dependency_actions":   actions.ActionsListType,
}

func (data AccessControlList) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AccessControlList"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AccessControlList) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.AccessControlEntry.IsNull() {
		return false
	}
	return true
}

func (data AccessControlList) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.AccessControlEntry.IsNull() {
		var dataValues []DmACE
		data.AccessControlEntry.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`AccessControlEntry`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`AccessControlEntry`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`AccessControlEntry`, "[]")
	}
	return body
}

func (data *AccessControlList) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessControlEntry`); value.Exists() {
		l := []DmACE{}
		if value := res.Get(`AccessControlEntry`); value.Exists() {
			for _, v := range value.Array() {
				item := DmACE{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AccessControlEntry, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmACEObjectType}, l)
		} else {
			data.AccessControlEntry = types.ListNull(types.ObjectType{AttrTypes: DmACEObjectType})
		}
	} else {
		data.AccessControlEntry = types.ListNull(types.ObjectType{AttrTypes: DmACEObjectType})
	}
}

func (data *AccessControlList) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessControlEntry`); value.Exists() && !data.AccessControlEntry.IsNull() {
		l := []DmACE{}
		e := []DmACE{}
		data.AccessControlEntry.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmACE{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AccessControlEntry, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmACEObjectType}, l)
		} else {
			data.AccessControlEntry = types.ListNull(types.ObjectType{AttrTypes: DmACEObjectType})
		}
	} else {
		data.AccessControlEntry = types.ListNull(types.ObjectType{AttrTypes: DmACEObjectType})
	}
}
