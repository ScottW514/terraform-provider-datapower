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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PeerGroup struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Type              types.String                `tfsdk:"type"`
	Url               types.List                  `tfsdk:"url"`
	IpMulticast       types.String                `tfsdk:"ip_multicast"`
	UpdateInterval    types.Int64                 `tfsdk:"update_interval"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var PeerGroupIPMulticastCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "slm",
	Value:       []string{"slm-multicast"},
}

var PeerGroupIPMulticastIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var PeerGroupUpdateIntervalCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "slm",
	Value:       []string{"slm-multicast"},
}

var PeerGroupUpdateIntervalIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var PeerGroupObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"type":               types.StringType,
	"url":                types.ListType{ElemType: types.StringType},
	"ip_multicast":       types.StringType,
	"update_interval":    types.Int64Type,
	"dependency_actions": actions.ActionsListType,
}

func (data PeerGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/PeerGroup"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data PeerGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.IpMulticast.IsNull() {
		return false
	}
	if !data.UpdateInterval.IsNull() {
		return false
	}
	return true
}

func (data PeerGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Url.IsNull() {
		var dataValues []string
		data.Url.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`URL`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`URL`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`URL`, "[]")
	}
	if !data.IpMulticast.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPMulticast`, data.IpMulticast.ValueString())
	}
	if !data.UpdateInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UpdateInterval`, data.UpdateInterval.ValueInt64())
	}
	return body
}

func (data *PeerGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("slm")
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() {
		data.Url = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Url = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `IPMulticast`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpMulticast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpMulticast = types.StringNull()
	}
	if value := res.Get(pathRoot + `UpdateInterval`); value.Exists() {
		data.UpdateInterval = types.Int64Value(value.Int())
	} else {
		data.UpdateInterval = types.Int64Value(10)
	}
}

func (data *PeerGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "slm" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Url = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `IPMulticast`); value.Exists() && !data.IpMulticast.IsNull() {
		data.IpMulticast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpMulticast = types.StringNull()
	}
	if value := res.Get(pathRoot + `UpdateInterval`); value.Exists() && !data.UpdateInterval.IsNull() {
		data.UpdateInterval = types.Int64Value(value.Int())
	} else if data.UpdateInterval.ValueInt64() != 10 {
		data.UpdateInterval = types.Int64Null()
	}
}
