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

type SSHService struct {
	Enabled           types.Bool                  `tfsdk:"enabled"`
	LocalPort         types.Int64                 `tfsdk:"local_port"`
	Acl               types.String                `tfsdk:"acl"`
	ConnectionLimit   types.Int64                 `tfsdk:"connection_limit"`
	LocalAddress      types.String                `tfsdk:"local_address"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SSHServiceObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"local_port":         types.Int64Type,
	"acl":                types.StringType,
	"connection_limit":   types.Int64Type,
	"local_address":      types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data SSHService) GetPath() string {
	rest_path := "/mgmt/config/default/SSHService/SSH Service"
	return rest_path
}

func (data SSHService) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.ConnectionLimit.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	return true
}

func (data SSHService) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "SSHService.name", path.Base("/mgmt/config/default/SSHService/SSH Service"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.ConnectionLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectionLimit`, data.ConnectionLimit.ValueInt64())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	return body
}

func (data *SSHService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(22)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringValue("ssh")
	}
	if value := res.Get(pathRoot + `ConnectionLimit`); value.Exists() {
		data.ConnectionLimit = types.Int64Value(value.Int())
	} else {
		data.ConnectionLimit = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
}

func (data *SSHService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 22 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else if data.Acl.ValueString() != "ssh" {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionLimit`); value.Exists() && !data.ConnectionLimit.IsNull() {
		data.ConnectionLimit = types.Int64Value(value.Int())
	} else if data.ConnectionLimit.ValueInt64() != 0 {
		data.ConnectionLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
}
