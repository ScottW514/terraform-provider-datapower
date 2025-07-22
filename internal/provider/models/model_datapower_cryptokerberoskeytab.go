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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CryptoKerberosKeytab struct {
	Id                  types.String        `tfsdk:"id"`
	AppDomain           types.String        `tfsdk:"app_domain"`
	UserSummary         types.String        `tfsdk:"user_summary"`
	Filename            types.String        `tfsdk:"filename"`
	UseReplayCache      types.Bool          `tfsdk:"use_replay_cache"`
	GenerateGssChecksum types.Bool          `tfsdk:"generate_gss_checksum"`
	GssChecksumFlags    *DmGssChecksumFlags `tfsdk:"gss_checksum_flags"`
}

var CryptoKerberosKeytabObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"filename":              types.StringType,
	"use_replay_cache":      types.BoolType,
	"generate_gss_checksum": types.BoolType,
	"gss_checksum_flags":    types.ObjectType{AttrTypes: DmGssChecksumFlagsObjectType},
}

func (data CryptoKerberosKeytab) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CryptoKerberosKeytab"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CryptoKerberosKeytab) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Filename.IsNull() {
		return false
	}
	if !data.UseReplayCache.IsNull() {
		return false
	}
	if !data.GenerateGssChecksum.IsNull() {
		return false
	}
	if data.GssChecksumFlags != nil {
		if !data.GssChecksumFlags.IsNull() {
			return false
		}
	}
	return true
}

func (data CryptoKerberosKeytab) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Filename.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Filename`, data.Filename.ValueString())
	}
	if !data.UseReplayCache.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseReplayCache`, tfutils.StringFromBool(data.UseReplayCache))
	}
	if !data.GenerateGssChecksum.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GenerateGssChecksum`, tfutils.StringFromBool(data.GenerateGssChecksum))
	}
	if data.GssChecksumFlags != nil {
		if !data.GssChecksumFlags.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`GssChecksumFlags`, data.GssChecksumFlags.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *CryptoKerberosKeytab) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Filename`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Filename = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Filename = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseReplayCache`); value.Exists() {
		data.UseReplayCache = tfutils.BoolFromString(value.String())
	} else {
		data.UseReplayCache = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GenerateGssChecksum`); value.Exists() {
		data.GenerateGssChecksum = tfutils.BoolFromString(value.String())
	} else {
		data.GenerateGssChecksum = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GssChecksumFlags`); value.Exists() {
		data.GssChecksumFlags = &DmGssChecksumFlags{}
		data.GssChecksumFlags.FromBody(ctx, "", value)
	} else {
		data.GssChecksumFlags = nil
	}
}

func (data *CryptoKerberosKeytab) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Filename`); value.Exists() && !data.Filename.IsNull() {
		data.Filename = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Filename = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseReplayCache`); value.Exists() && !data.UseReplayCache.IsNull() {
		data.UseReplayCache = tfutils.BoolFromString(value.String())
	} else if !data.UseReplayCache.ValueBool() {
		data.UseReplayCache = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GenerateGssChecksum`); value.Exists() && !data.GenerateGssChecksum.IsNull() {
		data.GenerateGssChecksum = tfutils.BoolFromString(value.String())
	} else if data.GenerateGssChecksum.ValueBool() {
		data.GenerateGssChecksum = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GssChecksumFlags`); value.Exists() {
		data.GssChecksumFlags.UpdateFromBody(ctx, "", value)
	} else {
		data.GssChecksumFlags = nil
	}
}
