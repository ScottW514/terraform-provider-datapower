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

type ConfigDeploymentPolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	AcceptedConfig    types.List                  `tfsdk:"accepted_config"`
	FilteredConfig    types.List                  `tfsdk:"filtered_config"`
	ModifiedConfig    types.List                  `tfsdk:"modified_config"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var ConfigDeploymentPolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"accepted_config":    types.ListType{ElemType: types.StringType},
	"filtered_config":    types.ListType{ElemType: types.StringType},
	"modified_config":    types.ListType{ElemType: types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data ConfigDeploymentPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ConfigDeploymentPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ConfigDeploymentPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AcceptedConfig.IsNull() {
		return false
	}
	if !data.FilteredConfig.IsNull() {
		return false
	}
	if !data.ModifiedConfig.IsNull() {
		return false
	}
	return true
}

func (data ConfigDeploymentPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AcceptedConfig.IsNull() {
		var dataValues []string
		data.AcceptedConfig.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`AcceptedConfig`+".-1", map[string]string{"value": val})
		}
	}
	if !data.FilteredConfig.IsNull() {
		var dataValues []string
		data.FilteredConfig.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`FilteredConfig`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ModifiedConfig.IsNull() {
		var dataValues []DmConfigModifyType
		data.ModifiedConfig.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ModifiedConfig`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *ConfigDeploymentPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AcceptedConfig`); value.Exists() {
		data.AcceptedConfig = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AcceptedConfig = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FilteredConfig`); value.Exists() {
		data.FilteredConfig = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FilteredConfig = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ModifiedConfig`); value.Exists() {
		l := []DmConfigModifyType{}
		if value := res.Get(`ModifiedConfig`); value.Exists() {
			for _, v := range value.Array() {
				item := DmConfigModifyType{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ModifiedConfig, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType}, l)
		} else {
			data.ModifiedConfig = types.ListNull(types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType})
		}
	} else {
		data.ModifiedConfig = types.ListNull(types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType})
	}
}

func (data *ConfigDeploymentPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AcceptedConfig`); value.Exists() && !data.AcceptedConfig.IsNull() {
		data.AcceptedConfig = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AcceptedConfig = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FilteredConfig`); value.Exists() && !data.FilteredConfig.IsNull() {
		data.FilteredConfig = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FilteredConfig = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ModifiedConfig`); value.Exists() && !data.ModifiedConfig.IsNull() {
		l := []DmConfigModifyType{}
		for _, v := range value.Array() {
			item := DmConfigModifyType{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ModifiedConfig, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType}, l)
		} else {
			data.ModifiedConfig = types.ListNull(types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType})
		}
	} else {
		data.ModifiedConfig = types.ListNull(types.ObjectType{AttrTypes: DmConfigModifyTypeObjectType})
	}
}
