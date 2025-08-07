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

type CompileOptionsPolicy struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	XsltVersion            types.String                `tfsdk:"xslt_version"`
	Strict                 types.Bool                  `tfsdk:"strict"`
	Profile                types.String                `tfsdk:"profile"`
	Debug                  types.String                `tfsdk:"debug"`
	Stream                 types.String                `tfsdk:"stream"`
	TryStream              types.String                `tfsdk:"try_stream"`
	MinimumEscaping        types.String                `tfsdk:"minimum_escaping"`
	StackSize              types.Int64                 `tfsdk:"stack_size"`
	PreferXg4              types.String                `tfsdk:"prefer_xg4"`
	DisallowXg4            types.String                `tfsdk:"disallow_xg4"`
	WsiValidation          types.String                `tfsdk:"wsi_validation"`
	WsdlValidateBody       types.String                `tfsdk:"wsdl_validate_body"`
	WsdlValidateHeaders    types.String                `tfsdk:"wsdl_validate_headers"`
	WsdlValidateFaults     types.String                `tfsdk:"wsdl_validate_faults"`
	WsdlWrappedFaults      types.Bool                  `tfsdk:"wsdl_wrapped_faults"`
	AllowSoapEncArray      types.String                `tfsdk:"allow_soap_enc_array"`
	ValidateSoapEncArray   types.String                `tfsdk:"validate_soap_enc_array"`
	WildcardsIgnoreXsiType types.String                `tfsdk:"wildcards_ignore_xsi_type"`
	WsdlStrictSoapVersion  types.Bool                  `tfsdk:"wsdl_strict_soap_version"`
	XacmlDebug             types.Bool                  `tfsdk:"xacml_debug"`
	AllowXopInclude        types.String                `tfsdk:"allow_xop_include"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var CompileOptionsPolicyObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"xslt_version":              types.StringType,
	"strict":                    types.BoolType,
	"profile":                   types.StringType,
	"debug":                     types.StringType,
	"stream":                    types.StringType,
	"try_stream":                types.StringType,
	"minimum_escaping":          types.StringType,
	"stack_size":                types.Int64Type,
	"prefer_xg4":                types.StringType,
	"disallow_xg4":              types.StringType,
	"wsi_validation":            types.StringType,
	"wsdl_validate_body":        types.StringType,
	"wsdl_validate_headers":     types.StringType,
	"wsdl_validate_faults":      types.StringType,
	"wsdl_wrapped_faults":       types.BoolType,
	"allow_soap_enc_array":      types.StringType,
	"validate_soap_enc_array":   types.StringType,
	"wildcards_ignore_xsi_type": types.StringType,
	"wsdl_strict_soap_version":  types.BoolType,
	"xacml_debug":               types.BoolType,
	"allow_xop_include":         types.StringType,
	"dependency_actions":        actions.ActionsListType,
}

func (data CompileOptionsPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CompileOptionsPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CompileOptionsPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.XsltVersion.IsNull() {
		return false
	}
	if !data.Strict.IsNull() {
		return false
	}
	if !data.Profile.IsNull() {
		return false
	}
	if !data.Debug.IsNull() {
		return false
	}
	if !data.Stream.IsNull() {
		return false
	}
	if !data.TryStream.IsNull() {
		return false
	}
	if !data.MinimumEscaping.IsNull() {
		return false
	}
	if !data.StackSize.IsNull() {
		return false
	}
	if !data.PreferXg4.IsNull() {
		return false
	}
	if !data.DisallowXg4.IsNull() {
		return false
	}
	if !data.WsiValidation.IsNull() {
		return false
	}
	if !data.WsdlValidateBody.IsNull() {
		return false
	}
	if !data.WsdlValidateHeaders.IsNull() {
		return false
	}
	if !data.WsdlValidateFaults.IsNull() {
		return false
	}
	if !data.WsdlWrappedFaults.IsNull() {
		return false
	}
	if !data.AllowSoapEncArray.IsNull() {
		return false
	}
	if !data.ValidateSoapEncArray.IsNull() {
		return false
	}
	if !data.WildcardsIgnoreXsiType.IsNull() {
		return false
	}
	if !data.WsdlStrictSoapVersion.IsNull() {
		return false
	}
	if !data.XacmlDebug.IsNull() {
		return false
	}
	if !data.AllowXopInclude.IsNull() {
		return false
	}
	return true
}

func (data CompileOptionsPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.XsltVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XSLTVersion`, data.XsltVersion.ValueString())
	}
	if !data.Strict.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Strict`, tfutils.StringFromBool(data.Strict, ""))
	}
	if !data.Profile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Profile`, data.Profile.ValueString())
	}
	if !data.Debug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Debug`, data.Debug.ValueString())
	}
	if !data.Stream.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Stream`, data.Stream.ValueString())
	}
	if !data.TryStream.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TryStream`, data.TryStream.ValueString())
	}
	if !data.MinimumEscaping.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MinimumEscaping`, data.MinimumEscaping.ValueString())
	}
	if !data.StackSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StackSize`, data.StackSize.ValueInt64())
	}
	if !data.PreferXg4.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PreferXG4`, data.PreferXg4.ValueString())
	}
	if !data.DisallowXg4.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisallowXG4`, data.DisallowXg4.ValueString())
	}
	if !data.WsiValidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSIValidation`, data.WsiValidation.ValueString())
	}
	if !data.WsdlValidateBody.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLValidateBody`, data.WsdlValidateBody.ValueString())
	}
	if !data.WsdlValidateHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLValidateHeaders`, data.WsdlValidateHeaders.ValueString())
	}
	if !data.WsdlValidateFaults.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLValidateFaults`, data.WsdlValidateFaults.ValueString())
	}
	if !data.WsdlWrappedFaults.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLWrappedFaults`, tfutils.StringFromBool(data.WsdlWrappedFaults, ""))
	}
	if !data.AllowSoapEncArray.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowSoapEncArray`, data.AllowSoapEncArray.ValueString())
	}
	if !data.ValidateSoapEncArray.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateSoapEncArray`, data.ValidateSoapEncArray.ValueString())
	}
	if !data.WildcardsIgnoreXsiType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WildcardsIgnoreXsiType`, data.WildcardsIgnoreXsiType.ValueString())
	}
	if !data.WsdlStrictSoapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLStrictSOAPVersion`, tfutils.StringFromBool(data.WsdlStrictSoapVersion, ""))
	}
	if !data.XacmlDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XACMLDebug`, tfutils.StringFromBool(data.XacmlDebug, ""))
	}
	if !data.AllowXopInclude.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowXOPInclude`, data.AllowXopInclude.ValueString())
	}
	return body
}

func (data *CompileOptionsPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `XSLTVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XsltVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XsltVersion = types.StringValue("XSLT10")
	}
	if value := res.Get(pathRoot + `Strict`); value.Exists() {
		data.Strict = tfutils.BoolFromString(value.String())
	} else {
		data.Strict = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Profile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Profile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Profile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Debug`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Debug = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Debug = types.StringNull()
	}
	if value := res.Get(pathRoot + `Stream`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Stream = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stream = types.StringNull()
	}
	if value := res.Get(pathRoot + `TryStream`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TryStream = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TryStream = types.StringNull()
	}
	if value := res.Get(pathRoot + `MinimumEscaping`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MinimumEscaping = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MinimumEscaping = types.StringNull()
	}
	if value := res.Get(pathRoot + `StackSize`); value.Exists() {
		data.StackSize = types.Int64Value(value.Int())
	} else {
		data.StackSize = types.Int64Value(524288)
	}
	if value := res.Get(pathRoot + `PreferXG4`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PreferXg4 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PreferXg4 = types.StringNull()
	}
	if value := res.Get(pathRoot + `DisallowXG4`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DisallowXg4 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DisallowXg4 = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSIValidation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsiValidation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsiValidation = types.StringValue("ignore")
	}
	if value := res.Get(pathRoot + `WSDLValidateBody`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlValidateBody = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlValidateBody = types.StringValue("strict")
	}
	if value := res.Get(pathRoot + `WSDLValidateHeaders`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlValidateHeaders = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlValidateHeaders = types.StringValue("lax")
	}
	if value := res.Get(pathRoot + `WSDLValidateFaults`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlValidateFaults = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlValidateFaults = types.StringValue("strict")
	}
	if value := res.Get(pathRoot + `WSDLWrappedFaults`); value.Exists() {
		data.WsdlWrappedFaults = tfutils.BoolFromString(value.String())
	} else {
		data.WsdlWrappedFaults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowSoapEncArray`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AllowSoapEncArray = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AllowSoapEncArray = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateSoapEncArray`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidateSoapEncArray = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateSoapEncArray = types.StringNull()
	}
	if value := res.Get(pathRoot + `WildcardsIgnoreXsiType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WildcardsIgnoreXsiType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WildcardsIgnoreXsiType = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLStrictSOAPVersion`); value.Exists() {
		data.WsdlStrictSoapVersion = tfutils.BoolFromString(value.String())
	} else {
		data.WsdlStrictSoapVersion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XACMLDebug`); value.Exists() {
		data.XacmlDebug = tfutils.BoolFromString(value.String())
	} else {
		data.XacmlDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowXOPInclude`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AllowXopInclude = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AllowXopInclude = types.StringNull()
	}
}

func (data *CompileOptionsPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `XSLTVersion`); value.Exists() && !data.XsltVersion.IsNull() {
		data.XsltVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.XsltVersion.ValueString() != "XSLT10" {
		data.XsltVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `Strict`); value.Exists() && !data.Strict.IsNull() {
		data.Strict = tfutils.BoolFromString(value.String())
	} else if data.Strict.ValueBool() {
		data.Strict = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Profile`); value.Exists() && !data.Profile.IsNull() {
		data.Profile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Profile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Debug`); value.Exists() && !data.Debug.IsNull() {
		data.Debug = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Debug = types.StringNull()
	}
	if value := res.Get(pathRoot + `Stream`); value.Exists() && !data.Stream.IsNull() {
		data.Stream = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Stream = types.StringNull()
	}
	if value := res.Get(pathRoot + `TryStream`); value.Exists() && !data.TryStream.IsNull() {
		data.TryStream = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TryStream = types.StringNull()
	}
	if value := res.Get(pathRoot + `MinimumEscaping`); value.Exists() && !data.MinimumEscaping.IsNull() {
		data.MinimumEscaping = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MinimumEscaping = types.StringNull()
	}
	if value := res.Get(pathRoot + `StackSize`); value.Exists() && !data.StackSize.IsNull() {
		data.StackSize = types.Int64Value(value.Int())
	} else if data.StackSize.ValueInt64() != 524288 {
		data.StackSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PreferXG4`); value.Exists() && !data.PreferXg4.IsNull() {
		data.PreferXg4 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PreferXg4 = types.StringNull()
	}
	if value := res.Get(pathRoot + `DisallowXG4`); value.Exists() && !data.DisallowXg4.IsNull() {
		data.DisallowXg4 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DisallowXg4 = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSIValidation`); value.Exists() && !data.WsiValidation.IsNull() {
		data.WsiValidation = tfutils.ParseStringFromGJSON(value)
	} else if data.WsiValidation.ValueString() != "ignore" {
		data.WsiValidation = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLValidateBody`); value.Exists() && !data.WsdlValidateBody.IsNull() {
		data.WsdlValidateBody = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlValidateBody.ValueString() != "strict" {
		data.WsdlValidateBody = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLValidateHeaders`); value.Exists() && !data.WsdlValidateHeaders.IsNull() {
		data.WsdlValidateHeaders = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlValidateHeaders.ValueString() != "lax" {
		data.WsdlValidateHeaders = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLValidateFaults`); value.Exists() && !data.WsdlValidateFaults.IsNull() {
		data.WsdlValidateFaults = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlValidateFaults.ValueString() != "strict" {
		data.WsdlValidateFaults = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLWrappedFaults`); value.Exists() && !data.WsdlWrappedFaults.IsNull() {
		data.WsdlWrappedFaults = tfutils.BoolFromString(value.String())
	} else if data.WsdlWrappedFaults.ValueBool() {
		data.WsdlWrappedFaults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowSoapEncArray`); value.Exists() && !data.AllowSoapEncArray.IsNull() {
		data.AllowSoapEncArray = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AllowSoapEncArray = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateSoapEncArray`); value.Exists() && !data.ValidateSoapEncArray.IsNull() {
		data.ValidateSoapEncArray = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateSoapEncArray = types.StringNull()
	}
	if value := res.Get(pathRoot + `WildcardsIgnoreXsiType`); value.Exists() && !data.WildcardsIgnoreXsiType.IsNull() {
		data.WildcardsIgnoreXsiType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WildcardsIgnoreXsiType = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLStrictSOAPVersion`); value.Exists() && !data.WsdlStrictSoapVersion.IsNull() {
		data.WsdlStrictSoapVersion = tfutils.BoolFromString(value.String())
	} else if data.WsdlStrictSoapVersion.ValueBool() {
		data.WsdlStrictSoapVersion = types.BoolNull()
	}
	if value := res.Get(pathRoot + `XACMLDebug`); value.Exists() && !data.XacmlDebug.IsNull() {
		data.XacmlDebug = tfutils.BoolFromString(value.String())
	} else if data.XacmlDebug.ValueBool() {
		data.XacmlDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowXOPInclude`); value.Exists() && !data.AllowXopInclude.IsNull() {
		data.AllowXopInclude = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AllowXopInclude = types.StringNull()
	}
}
