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

type Luna struct {
	Id                types.String      `tfsdk:"id"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	RemoteAddress     types.String      `tfsdk:"remote_address"`
	ServerCert        types.String      `tfsdk:"server_cert"`
	SecurityOption    types.String      `tfsdk:"security_option"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var LunaObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"user_summary":       types.StringType,
	"remote_address":     types.StringType,
	"server_cert":        types.StringType,
	"security_option":    types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data Luna) GetPath() string {
	rest_path := "/mgmt/config/default/Luna"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	return rest_path
}

func (data Luna) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.ServerCert.IsNull() {
		return false
	}
	if !data.SecurityOption.IsNull() {
		return false
	}
	return true
}

func (data Luna) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.ServerCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServerCert`, data.ServerCert.ValueString())
	}
	if !data.SecurityOption.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecurityOption`, data.SecurityOption.ValueString())
	}
	return body
}

func (data *Luna) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityOption`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecurityOption = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecurityOption = types.StringValue("none")
	}
}

func (data *Luna) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServerCert`); value.Exists() && !data.ServerCert.IsNull() {
		data.ServerCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServerCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityOption`); value.Exists() && !data.SecurityOption.IsNull() {
		data.SecurityOption = tfutils.ParseStringFromGJSON(value)
	} else if data.SecurityOption.ValueString() != "none" {
		data.SecurityOption = types.StringNull()
	}
}
