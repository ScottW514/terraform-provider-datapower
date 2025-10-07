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

type FTPQuoteCommands struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	FtpQuotedCommands types.List                  `tfsdk:"ftp_quoted_commands"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var FTPQuoteCommandsObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"ftp_quoted_commands": types.ListType{ElemType: types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType}},
	"dependency_actions":  actions.ActionsListType,
}

func (data FTPQuoteCommands) GetPath() string {
	rest_path := "/mgmt/config/{domain}/FTPQuoteCommands"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data FTPQuoteCommands) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.FtpQuotedCommands.IsNull() {
		return false
	}
	return true
}

func (data FTPQuoteCommands) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.FtpQuotedCommands.IsNull() {
		var dataValues []DmFTPQuotedCommand
		data.FtpQuotedCommands.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`FTPQuotedCommands`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *FTPQuoteCommands) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `FTPQuotedCommands`); value.Exists() {
		l := []DmFTPQuotedCommand{}
		if value := res.Get(`FTPQuotedCommands`); value.Exists() {
			for _, v := range value.Array() {
				item := DmFTPQuotedCommand{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FtpQuotedCommands, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType}, l)
		} else {
			data.FtpQuotedCommands = types.ListNull(types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType})
		}
	} else {
		data.FtpQuotedCommands = types.ListNull(types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType})
	}
}

func (data *FTPQuoteCommands) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `FTPQuotedCommands`); value.Exists() && !data.FtpQuotedCommands.IsNull() {
		l := []DmFTPQuotedCommand{}
		e := []DmFTPQuotedCommand{}
		data.FtpQuotedCommands.ElementsAs(ctx, &e, false)
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
				item := DmFTPQuotedCommand{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FtpQuotedCommands, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType}, l)
		} else {
			data.FtpQuotedCommands = types.ListNull(types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType})
		}
	} else {
		data.FtpQuotedCommands = types.ListNull(types.ObjectType{AttrTypes: DmFTPQuotedCommandObjectType})
	}
}
