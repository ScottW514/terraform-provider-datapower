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

type LunaPartition struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	PartitionName     types.String                `tfsdk:"partition_name"`
	PartitionSerial   types.String                `tfsdk:"partition_serial"`
	PasswordAlias     types.String                `tfsdk:"password_alias"`
	LoginRole         types.String                `tfsdk:"login_role"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var LunaPartitionObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"partition_name":     types.StringType,
	"partition_serial":   types.StringType,
	"password_alias":     types.StringType,
	"login_role":         types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data LunaPartition) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LunaPartition"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LunaPartition) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PartitionName.IsNull() {
		return false
	}
	if !data.PartitionSerial.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.LoginRole.IsNull() {
		return false
	}
	return true
}

func (data LunaPartition) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.PartitionName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PartitionName`, data.PartitionName.ValueString())
	}
	if !data.PartitionSerial.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PartitionSerial`, data.PartitionSerial.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.LoginRole.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoginRole`, data.LoginRole.ValueString())
	}
	return body
}

func (data *LunaPartition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PartitionName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PartitionName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartitionName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PartitionSerial`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PartitionSerial = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartitionSerial = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LoginRole`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LoginRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoginRole = types.StringValue("co")
	}
}

func (data *LunaPartition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PartitionName`); value.Exists() && !data.PartitionName.IsNull() {
		data.PartitionName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartitionName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PartitionSerial`); value.Exists() && !data.PartitionSerial.IsNull() {
		data.PartitionSerial = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PartitionSerial = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LoginRole`); value.Exists() && !data.LoginRole.IsNull() {
		data.LoginRole = tfutils.ParseStringFromGJSON(value)
	} else if data.LoginRole.ValueString() != "co" {
		data.LoginRole = types.StringNull()
	}
}
