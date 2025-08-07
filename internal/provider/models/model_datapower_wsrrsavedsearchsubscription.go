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

type WSRRSavedSearchSubscription struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	Server                 types.String                `tfsdk:"server"`
	SavedSearchName        types.String                `tfsdk:"saved_search_name"`
	SavedSearchParameters  types.List                  `tfsdk:"saved_search_parameters"`
	Method                 types.String                `tfsdk:"method"`
	RefreshInterval        types.Int64                 `tfsdk:"refresh_interval"`
	FetchPolicyAttachments types.Bool                  `tfsdk:"fetch_policy_attachments"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WSRRSavedSearchSubscriptionObjectType = map[string]attr.Type{
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"server":                   types.StringType,
	"saved_search_name":        types.StringType,
	"saved_search_parameters":  types.ListType{ElemType: types.StringType},
	"method":                   types.StringType,
	"refresh_interval":         types.Int64Type,
	"fetch_policy_attachments": types.BoolType,
	"user_summary":             types.StringType,
	"dependency_actions":       actions.ActionsListType,
}

func (data WSRRSavedSearchSubscription) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSRRSavedSearchSubscription"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSRRSavedSearchSubscription) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Server.IsNull() {
		return false
	}
	if !data.SavedSearchName.IsNull() {
		return false
	}
	if !data.SavedSearchParameters.IsNull() {
		return false
	}
	if !data.Method.IsNull() {
		return false
	}
	if !data.RefreshInterval.IsNull() {
		return false
	}
	if !data.FetchPolicyAttachments.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data WSRRSavedSearchSubscription) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Server.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Server`, data.Server.ValueString())
	}
	if !data.SavedSearchName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SavedSearchName`, data.SavedSearchName.ValueString())
	}
	if !data.SavedSearchParameters.IsNull() {
		var values []string
		data.SavedSearchParameters.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`SavedSearchParameters`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	if !data.RefreshInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshInterval`, data.RefreshInterval.ValueInt64())
	}
	if !data.FetchPolicyAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FetchPolicyAttachments`, tfutils.StringFromBool(data.FetchPolicyAttachments, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *WSRRSavedSearchSubscription) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `SavedSearchName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SavedSearchName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SavedSearchName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SavedSearchParameters`); value.Exists() {
		data.SavedSearchParameters = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SavedSearchParameters = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Method = types.StringValue("poll")
	}
	if value := res.Get(pathRoot + `RefreshInterval`); value.Exists() {
		data.RefreshInterval = types.Int64Value(value.Int())
	} else {
		data.RefreshInterval = types.Int64Value(86400)
	}
	if value := res.Get(pathRoot + `FetchPolicyAttachments`); value.Exists() {
		data.FetchPolicyAttachments = tfutils.BoolFromString(value.String())
	} else {
		data.FetchPolicyAttachments = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *WSRRSavedSearchSubscription) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Server`); value.Exists() && !data.Server.IsNull() {
		data.Server = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Server = types.StringNull()
	}
	if value := res.Get(pathRoot + `SavedSearchName`); value.Exists() && !data.SavedSearchName.IsNull() {
		data.SavedSearchName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SavedSearchName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SavedSearchParameters`); value.Exists() && !data.SavedSearchParameters.IsNull() {
		data.SavedSearchParameters = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SavedSearchParameters = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Method`); value.Exists() && !data.Method.IsNull() {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else if data.Method.ValueString() != "poll" {
		data.Method = types.StringNull()
	}
	if value := res.Get(pathRoot + `RefreshInterval`); value.Exists() && !data.RefreshInterval.IsNull() {
		data.RefreshInterval = types.Int64Value(value.Int())
	} else if data.RefreshInterval.ValueInt64() != 86400 {
		data.RefreshInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FetchPolicyAttachments`); value.Exists() && !data.FetchPolicyAttachments.IsNull() {
		data.FetchPolicyAttachments = tfutils.BoolFromString(value.String())
	} else if data.FetchPolicyAttachments.ValueBool() {
		data.FetchPolicyAttachments = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
