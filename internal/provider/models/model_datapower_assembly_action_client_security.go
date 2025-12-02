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

type AssemblyActionClientSecurity struct {
	Id                       types.String                `tfsdk:"id"`
	AppDomain                types.String                `tfsdk:"app_domain"`
	StopOnError              types.Bool                  `tfsdk:"stop_on_error"`
	SecretRequired           types.Bool                  `tfsdk:"secret_required"`
	ExtractCredentialMethod  types.String                `tfsdk:"extract_credential_method"`
	IdName                   types.String                `tfsdk:"id_name"`
	SecretName               types.String                `tfsdk:"secret_name"`
	HttpType                 types.String                `tfsdk:"http_type"`
	AuthenticateClientMethod types.String                `tfsdk:"authenticate_client_method"`
	UserRegistry             types.String                `tfsdk:"user_registry"`
	UserSummary              types.String                `tfsdk:"user_summary"`
	Title                    types.String                `tfsdk:"title"`
	CorrelationPath          types.String                `tfsdk:"correlation_path"`
	ActionDebug              types.Bool                  `tfsdk:"action_debug"`
	DependencyActions        []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget           types.String                `tfsdk:"provider_target"`
}

var AssemblyActionClientSecurityIdNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_credential_method",
	AttrType:    "String",
	AttrDefault: "header",
	Value:       []string{"header", "cookie", "query", "form", "context-var"},
}

var AssemblyActionClientSecurityIdNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AssemblyActionClientSecuritySecretNameCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "secret_required",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "extract_credential_method",
			AttrType:    "String",
			AttrDefault: "header",
			Value:       []string{"header", "cookie", "query", "form", "context-var"},
		},
	},
}

var AssemblyActionClientSecuritySecretNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "extract_credential_method",
			AttrType:    "String",
			AttrDefault: "header",
			Value:       []string{"http"},
		},
	},
}

var AssemblyActionClientSecurityHTTPTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_credential_method",
	AttrType:    "String",
	AttrDefault: "header",
	Value:       []string{"http"},
}

var AssemblyActionClientSecurityHTTPTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AssemblyActionClientSecurityUserRegistryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "authenticate_client_method",
	AttrType:    "String",
	AttrDefault: "native",
	Value:       []string{"third-party"},
}

var AssemblyActionClientSecurityUserRegistryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var AssemblyActionClientSecurityObjectType = map[string]attr.Type{
	"provider_target":            types.StringType,
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"stop_on_error":              types.BoolType,
	"secret_required":            types.BoolType,
	"extract_credential_method":  types.StringType,
	"id_name":                    types.StringType,
	"secret_name":                types.StringType,
	"http_type":                  types.StringType,
	"authenticate_client_method": types.StringType,
	"user_registry":              types.StringType,
	"user_summary":               types.StringType,
	"title":                      types.StringType,
	"correlation_path":           types.StringType,
	"action_debug":               types.BoolType,
	"dependency_actions":         actions.ActionsListType,
}

func (data AssemblyActionClientSecurity) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionClientSecurity"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionClientSecurity) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.StopOnError.IsNull() {
		return false
	}
	if !data.SecretRequired.IsNull() {
		return false
	}
	if !data.ExtractCredentialMethod.IsNull() {
		return false
	}
	if !data.IdName.IsNull() {
		return false
	}
	if !data.SecretName.IsNull() {
		return false
	}
	if !data.HttpType.IsNull() {
		return false
	}
	if !data.AuthenticateClientMethod.IsNull() {
		return false
	}
	if !data.UserRegistry.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.CorrelationPath.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	return true
}

func (data AssemblyActionClientSecurity) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.StopOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StopOnError`, tfutils.StringFromBool(data.StopOnError, ""))
	}
	if !data.SecretRequired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecretRequired`, tfutils.StringFromBool(data.SecretRequired, ""))
	}
	if !data.ExtractCredentialMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExtractCredentialMethod`, data.ExtractCredentialMethod.ValueString())
	}
	if !data.IdName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdName`, data.IdName.ValueString())
	}
	if !data.SecretName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecretName`, data.SecretName.ValueString())
	}
	if !data.HttpType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPType`, data.HttpType.ValueString())
	}
	if !data.AuthenticateClientMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthenticateClientMethod`, data.AuthenticateClientMethod.ValueString())
	}
	if !data.UserRegistry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserRegistry`, data.UserRegistry.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.CorrelationPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CorrelationPath`, data.CorrelationPath.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	return body
}

func (data *AssemblyActionClientSecurity) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `StopOnError`); value.Exists() {
		data.StopOnError = tfutils.BoolFromString(value.String())
	} else {
		data.StopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SecretRequired`); value.Exists() {
		data.SecretRequired = tfutils.BoolFromString(value.String())
	} else {
		data.SecretRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExtractCredentialMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExtractCredentialMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtractCredentialMethod = types.StringValue("header")
	}
	if value := res.Get(pathRoot + `IdName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IdName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IdName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecretName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecretName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecretName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpType = types.StringValue("basic")
	}
	if value := res.Get(pathRoot + `AuthenticateClientMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthenticateClientMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthenticateClientMethod = types.StringValue("native")
	}
	if value := res.Get(pathRoot + `UserRegistry`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserRegistry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserRegistry = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
}

func (data *AssemblyActionClientSecurity) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `StopOnError`); value.Exists() && !data.StopOnError.IsNull() {
		data.StopOnError = tfutils.BoolFromString(value.String())
	} else if !data.StopOnError.ValueBool() {
		data.StopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SecretRequired`); value.Exists() && !data.SecretRequired.IsNull() {
		data.SecretRequired = tfutils.BoolFromString(value.String())
	} else if !data.SecretRequired.ValueBool() {
		data.SecretRequired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ExtractCredentialMethod`); value.Exists() && !data.ExtractCredentialMethod.IsNull() {
		data.ExtractCredentialMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.ExtractCredentialMethod.ValueString() != "header" {
		data.ExtractCredentialMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `IdName`); value.Exists() && !data.IdName.IsNull() {
		data.IdName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IdName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecretName`); value.Exists() && !data.SecretName.IsNull() {
		data.SecretName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecretName = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPType`); value.Exists() && !data.HttpType.IsNull() {
		data.HttpType = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpType.ValueString() != "basic" {
		data.HttpType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthenticateClientMethod`); value.Exists() && !data.AuthenticateClientMethod.IsNull() {
		data.AuthenticateClientMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthenticateClientMethod.ValueString() != "native" {
		data.AuthenticateClientMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserRegistry`); value.Exists() && !data.UserRegistry.IsNull() {
		data.UserRegistry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserRegistry = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `CorrelationPath`); value.Exists() && !data.CorrelationPath.IsNull() {
		data.CorrelationPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorrelationPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
}
