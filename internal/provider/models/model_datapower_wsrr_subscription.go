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

type WSRRSubscription struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	Server                 types.String                `tfsdk:"server"`
	Namespace              types.String                `tfsdk:"namespace"`
	ObjectType             types.String                `tfsdk:"object_type"`
	ObjectName             types.String                `tfsdk:"object_name"`
	Method                 types.String                `tfsdk:"method"`
	RefreshInterval        types.Int64                 `tfsdk:"refresh_interval"`
	UseVersion             types.Bool                  `tfsdk:"use_version"`
	ObjectVersion          types.String                `tfsdk:"object_version"`
	FetchPolicyAttachments types.Bool                  `tfsdk:"fetch_policy_attachments"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget         types.String                `tfsdk:"provider_target"`
}

var WSRRSubscriptionNamespaceCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "object_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"service-version"},
}

var WSRRSubscriptionRefreshIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "method",
	AttrType:    "String",
	AttrDefault: "poll",
	Value:       []string{"poll"},
}

var WSRRSubscriptionObjectVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_version",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var WSRRSubscriptionObjectType = map[string]attr.Type{
	"provider_target":          types.StringType,
	"id":                       types.StringType,
	"app_domain":               types.StringType,
	"server":                   types.StringType,
	"namespace":                types.StringType,
	"object_type":              types.StringType,
	"object_name":              types.StringType,
	"method":                   types.StringType,
	"refresh_interval":         types.Int64Type,
	"use_version":              types.BoolType,
	"object_version":           types.StringType,
	"fetch_policy_attachments": types.BoolType,
	"user_summary":             types.StringType,
	"dependency_actions":       actions.ActionsListType,
}

func (data WSRRSubscription) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSRRSubscription"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSRRSubscription) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Server.IsNull() {
		return false
	}
	if !data.Namespace.IsNull() {
		return false
	}
	if !data.ObjectType.IsNull() {
		return false
	}
	if !data.ObjectName.IsNull() {
		return false
	}
	if !data.Method.IsNull() {
		return false
	}
	if !data.RefreshInterval.IsNull() {
		return false
	}
	if !data.UseVersion.IsNull() {
		return false
	}
	if !data.ObjectVersion.IsNull() {
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

func (data WSRRSubscription) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Namespace.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Namespace`, data.Namespace.ValueString())
	}
	if !data.ObjectType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ObjectType`, data.ObjectType.ValueString())
	}
	if !data.ObjectName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ObjectName`, data.ObjectName.ValueString())
	}
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	if !data.RefreshInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshInterval`, data.RefreshInterval.ValueInt64())
	}
	if !data.UseVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseVersion`, tfutils.StringFromBool(data.UseVersion, ""))
	}
	if !data.ObjectVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ObjectVersion`, data.ObjectVersion.ValueString())
	}
	if !data.FetchPolicyAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FetchPolicyAttachments`, tfutils.StringFromBool(data.FetchPolicyAttachments, ""))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *WSRRSubscription) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Namespace`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Namespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ObjectType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ObjectType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ObjectName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ObjectName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectName = types.StringNull()
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
	if value := res.Get(pathRoot + `UseVersion`); value.Exists() {
		data.UseVersion = tfutils.BoolFromString(value.String())
	} else {
		data.UseVersion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ObjectVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ObjectVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectVersion = types.StringNull()
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

func (data *WSRRSubscription) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Namespace`); value.Exists() && !data.Namespace.IsNull() {
		data.Namespace = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Namespace = types.StringNull()
	}
	if value := res.Get(pathRoot + `ObjectType`); value.Exists() && !data.ObjectType.IsNull() {
		data.ObjectType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ObjectName`); value.Exists() && !data.ObjectName.IsNull() {
		data.ObjectName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectName = types.StringNull()
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
	if value := res.Get(pathRoot + `UseVersion`); value.Exists() && !data.UseVersion.IsNull() {
		data.UseVersion = tfutils.BoolFromString(value.String())
	} else if data.UseVersion.ValueBool() {
		data.UseVersion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ObjectVersion`); value.Exists() && !data.ObjectVersion.IsNull() {
		data.ObjectVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ObjectVersion = types.StringNull()
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
