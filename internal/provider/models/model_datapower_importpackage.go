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

type ImportPackage struct {
	Id                         types.String      `tfsdk:"id"`
	AppDomain                  types.String      `tfsdk:"app_domain"`
	UserSummary                types.String      `tfsdk:"user_summary"`
	Url                        types.String      `tfsdk:"url"`
	ImportFormat               types.String      `tfsdk:"import_format"`
	OverwriteFiles             types.Bool        `tfsdk:"overwrite_files"`
	OverwriteObjects           types.Bool        `tfsdk:"overwrite_objects"`
	DestinationDomain          types.String      `tfsdk:"destination_domain"`
	DeploymentPolicy           types.String      `tfsdk:"deployment_policy"`
	DeploymentPolicyParameters types.String      `tfsdk:"deployment_policy_parameters"`
	LocalIpRewrite             types.Bool        `tfsdk:"local_ip_rewrite"`
	OnStartup                  types.Bool        `tfsdk:"on_startup"`
	ObjectActions              []*actions.Action `tfsdk:"object_actions"`
}

var ImportPackageObjectType = map[string]attr.Type{
	"id":                           types.StringType,
	"app_domain":                   types.StringType,
	"user_summary":                 types.StringType,
	"url":                          types.StringType,
	"import_format":                types.StringType,
	"overwrite_files":              types.BoolType,
	"overwrite_objects":            types.BoolType,
	"destination_domain":           types.StringType,
	"deployment_policy":            types.StringType,
	"deployment_policy_parameters": types.StringType,
	"local_ip_rewrite":             types.BoolType,
	"on_startup":                   types.BoolType,
	"object_actions":               actions.ActionsListType,
}

func (data ImportPackage) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ImportPackage"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ImportPackage) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.ImportFormat.IsNull() {
		return false
	}
	if !data.OverwriteFiles.IsNull() {
		return false
	}
	if !data.OverwriteObjects.IsNull() {
		return false
	}
	if !data.DestinationDomain.IsNull() {
		return false
	}
	if !data.DeploymentPolicy.IsNull() {
		return false
	}
	if !data.DeploymentPolicyParameters.IsNull() {
		return false
	}
	if !data.LocalIpRewrite.IsNull() {
		return false
	}
	if !data.OnStartup.IsNull() {
		return false
	}
	return true
}

func (data ImportPackage) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.ImportFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ImportFormat`, data.ImportFormat.ValueString())
	}
	if !data.OverwriteFiles.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OverwriteFiles`, tfutils.StringFromBool(data.OverwriteFiles, ""))
	}
	if !data.OverwriteObjects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OverwriteObjects`, tfutils.StringFromBool(data.OverwriteObjects, ""))
	}
	if !data.DestinationDomain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DestinationDomain`, data.DestinationDomain.ValueString())
	}
	if !data.DeploymentPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeploymentPolicy`, data.DeploymentPolicy.ValueString())
	}
	if !data.DeploymentPolicyParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeploymentPolicyParameters`, data.DeploymentPolicyParameters.ValueString())
	}
	if !data.LocalIpRewrite.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalIPRewrite`, tfutils.StringFromBool(data.LocalIpRewrite, ""))
	}
	if !data.OnStartup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OnStartup`, tfutils.StringFromBool(data.OnStartup, ""))
	}
	return body
}

func (data *ImportPackage) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ImportFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ImportFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ImportFormat = types.StringValue("ZIP")
	}
	if value := res.Get(pathRoot + `OverwriteFiles`); value.Exists() {
		data.OverwriteFiles = tfutils.BoolFromString(value.String())
	} else {
		data.OverwriteFiles = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OverwriteObjects`); value.Exists() {
		data.OverwriteObjects = tfutils.BoolFromString(value.String())
	} else {
		data.OverwriteObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DestinationDomain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DestinationDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestinationDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeploymentPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicyParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeploymentPolicyParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicyParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalIPRewrite`); value.Exists() {
		data.LocalIpRewrite = tfutils.BoolFromString(value.String())
	} else {
		data.LocalIpRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OnStartup`); value.Exists() {
		data.OnStartup = tfutils.BoolFromString(value.String())
	} else {
		data.OnStartup = types.BoolNull()
	}
}

func (data *ImportPackage) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `ImportFormat`); value.Exists() && !data.ImportFormat.IsNull() {
		data.ImportFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.ImportFormat.ValueString() != "ZIP" {
		data.ImportFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `OverwriteFiles`); value.Exists() && !data.OverwriteFiles.IsNull() {
		data.OverwriteFiles = tfutils.BoolFromString(value.String())
	} else if !data.OverwriteFiles.ValueBool() {
		data.OverwriteFiles = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OverwriteObjects`); value.Exists() && !data.OverwriteObjects.IsNull() {
		data.OverwriteObjects = tfutils.BoolFromString(value.String())
	} else if !data.OverwriteObjects.ValueBool() {
		data.OverwriteObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DestinationDomain`); value.Exists() && !data.DestinationDomain.IsNull() {
		data.DestinationDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestinationDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicy`); value.Exists() && !data.DeploymentPolicy.IsNull() {
		data.DeploymentPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicyParameters`); value.Exists() && !data.DeploymentPolicyParameters.IsNull() {
		data.DeploymentPolicyParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicyParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalIPRewrite`); value.Exists() && !data.LocalIpRewrite.IsNull() {
		data.LocalIpRewrite = tfutils.BoolFromString(value.String())
	} else if !data.LocalIpRewrite.ValueBool() {
		data.LocalIpRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OnStartup`); value.Exists() && !data.OnStartup.IsNull() {
		data.OnStartup = tfutils.BoolFromString(value.String())
	} else if !data.OnStartup.ValueBool() {
		data.OnStartup = types.BoolNull()
	}
}
