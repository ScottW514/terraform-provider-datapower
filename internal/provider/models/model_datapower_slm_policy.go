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

type SLMPolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	ExecutionPolicy   types.String                `tfsdk:"execution_policy"`
	Statement         types.List                  `tfsdk:"statement"`
	PeerGroup         types.String                `tfsdk:"peer_group"`
	ApiMgmt           types.Bool                  `tfsdk:"api_mgmt"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SLMPolicyObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"execution_policy":   types.StringType,
	"statement":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmSLMStatementObjectType}},
	"peer_group":         types.StringType,
	"api_mgmt":           types.BoolType,
	"dependency_actions": actions.ActionsListType,
}

func (data SLMPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SLMPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SLMPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ExecutionPolicy.IsNull() {
		return false
	}
	if !data.Statement.IsNull() {
		return false
	}
	if !data.PeerGroup.IsNull() {
		return false
	}
	if !data.ApiMgmt.IsNull() {
		return false
	}
	return true
}

func (data SLMPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ExecutionPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExecutionPolicy`, data.ExecutionPolicy.ValueString())
	}
	if !data.Statement.IsNull() {
		var dataValues []DmSLMStatement
		data.Statement.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Statement`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.PeerGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PeerGroup`, data.PeerGroup.ValueString())
	}
	if !data.ApiMgmt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIMgmt`, tfutils.StringFromBool(data.ApiMgmt, ""))
	}
	return body
}

func (data *SLMPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ExecutionPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExecutionPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExecutionPolicy = types.StringValue("execute-all-statements")
	}
	if value := res.Get(pathRoot + `Statement`); value.Exists() {
		l := []DmSLMStatement{}
		if value := res.Get(`Statement`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSLMStatement{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Statement, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSLMStatementObjectType}, l)
		} else {
			data.Statement = types.ListNull(types.ObjectType{AttrTypes: DmSLMStatementObjectType})
		}
	} else {
		data.Statement = types.ListNull(types.ObjectType{AttrTypes: DmSLMStatementObjectType})
	}
	if value := res.Get(pathRoot + `PeerGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PeerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PeerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIMgmt`); value.Exists() {
		data.ApiMgmt = tfutils.BoolFromString(value.String())
	} else {
		data.ApiMgmt = types.BoolNull()
	}
}

func (data *SLMPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ExecutionPolicy`); value.Exists() && !data.ExecutionPolicy.IsNull() {
		data.ExecutionPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.ExecutionPolicy.ValueString() != "execute-all-statements" {
		data.ExecutionPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `Statement`); value.Exists() && !data.Statement.IsNull() {
		l := []DmSLMStatement{}
		for _, v := range value.Array() {
			item := DmSLMStatement{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Statement, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSLMStatementObjectType}, l)
		} else {
			data.Statement = types.ListNull(types.ObjectType{AttrTypes: DmSLMStatementObjectType})
		}
	} else {
		data.Statement = types.ListNull(types.ObjectType{AttrTypes: DmSLMStatementObjectType})
	}
	if value := res.Get(pathRoot + `PeerGroup`); value.Exists() && !data.PeerGroup.IsNull() {
		data.PeerGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PeerGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIMgmt`); value.Exists() && !data.ApiMgmt.IsNull() {
		data.ApiMgmt = tfutils.BoolFromString(value.String())
	} else {
		data.ApiMgmt = types.BoolNull()
	}
}
