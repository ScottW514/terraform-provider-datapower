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

type B2BProfileGroup struct {
	Id                types.String      `tfsdk:"id"`
	AppDomain         types.String      `tfsdk:"app_domain"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	B2bProfiles       types.List        `tfsdk:"b2b_profiles"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var B2BProfileGroupObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"b2b_profiles":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data B2BProfileGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BProfileGroup"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BProfileGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.B2bProfiles.IsNull() {
		return false
	}
	return true
}

func (data B2BProfileGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.B2bProfiles.IsNull() {
		var values []DmB2BGroupedProfile
		data.B2bProfiles.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`B2BProfiles`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *B2BProfileGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `B2BProfiles`); value.Exists() {
		l := []DmB2BGroupedProfile{}
		if value := res.Get(`B2BProfiles`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BGroupedProfile{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.B2bProfiles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType}, l)
		} else {
			data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType})
		}
	} else {
		data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType})
	}
}

func (data *B2BProfileGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `B2BProfiles`); value.Exists() && !data.B2bProfiles.IsNull() {
		l := []DmB2BGroupedProfile{}
		for _, v := range value.Array() {
			item := DmB2BGroupedProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.B2bProfiles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType}, l)
		} else {
			data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType})
		}
	} else {
		data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BGroupedProfileObjectType})
	}
}
