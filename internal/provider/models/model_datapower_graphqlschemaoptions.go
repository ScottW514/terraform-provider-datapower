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

type GraphQLSchemaOptions struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	VisibilityList    types.String                `tfsdk:"visibility_list"`
	Api               types.String                `tfsdk:"api"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var GraphQLSchemaOptionsObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"visibility_list":    types.StringType,
	"api":                types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data GraphQLSchemaOptions) GetPath() string {
	rest_path := "/mgmt/config/{domain}/GraphQLSchemaOptions"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data GraphQLSchemaOptions) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.VisibilityList.IsNull() {
		return false
	}
	if !data.Api.IsNull() {
		return false
	}
	return true
}

func (data GraphQLSchemaOptions) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.VisibilityList.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VisibilityList`, data.VisibilityList.ValueString())
	}
	if !data.Api.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`API`, data.Api.ValueString())
	}
	return body
}

func (data *GraphQLSchemaOptions) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `VisibilityList`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VisibilityList = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VisibilityList = types.StringNull()
	}
	if value := res.Get(pathRoot + `API`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Api = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Api = types.StringNull()
	}
}

func (data *GraphQLSchemaOptions) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `VisibilityList`); value.Exists() && !data.VisibilityList.IsNull() {
		data.VisibilityList = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VisibilityList = types.StringNull()
	}
	if value := res.Get(pathRoot + `API`); value.Exists() && !data.Api.IsNull() {
		data.Api = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Api = types.StringNull()
	}
}
