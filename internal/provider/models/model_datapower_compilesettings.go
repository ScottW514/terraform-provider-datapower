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

type CompileSettings struct {
	Id                     types.String `tfsdk:"id"`
	AppDomain              types.String `tfsdk:"app_domain"`
	UserSummary            types.String `tfsdk:"user_summary"`
	XsltVersion            types.String `tfsdk:"xslt_version"`
	Strict                 types.Bool   `tfsdk:"strict"`
	Profile                types.Bool   `tfsdk:"profile"`
	Debug                  types.Bool   `tfsdk:"debug"`
	Stream                 types.Bool   `tfsdk:"stream"`
	TryStream              types.Bool   `tfsdk:"try_stream"`
	MinimumEscaping        types.Bool   `tfsdk:"minimum_escaping"`
	StackSize              types.Int64  `tfsdk:"stack_size"`
	WsiValidation          types.String `tfsdk:"wsi_validation"`
	WsdlValidateBody       types.String `tfsdk:"wsdl_validate_body"`
	WsdlValidateHeaders    types.String `tfsdk:"wsdl_validate_headers"`
	WsdlValidateFaults     types.String `tfsdk:"wsdl_validate_faults"`
	WsdlWrappedFaults      types.Bool   `tfsdk:"wsdl_wrapped_faults"`
	AllowSoapEncArray      types.Bool   `tfsdk:"allow_soap_enc_array"`
	ValidateSoapEncArray   types.Bool   `tfsdk:"validate_soap_enc_array"`
	WildcardsIgnoreXsiType types.Bool   `tfsdk:"wildcards_ignore_xsi_type"`
	WsdlStrictSoapVersion  types.Bool   `tfsdk:"wsdl_strict_soap_version"`
	XacmlDebug             types.Bool   `tfsdk:"xacml_debug"`
	AllowXopInclude        types.Bool   `tfsdk:"allow_xop_include"`
}

var CompileSettingsObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"xslt_version":              types.StringType,
	"strict":                    types.BoolType,
	"profile":                   types.BoolType,
	"debug":                     types.BoolType,
	"stream":                    types.BoolType,
	"try_stream":                types.BoolType,
	"minimum_escaping":          types.BoolType,
	"stack_size":                types.Int64Type,
	"wsi_validation":            types.StringType,
	"wsdl_validate_body":        types.StringType,
	"wsdl_validate_headers":     types.StringType,
	"wsdl_validate_faults":      types.StringType,
	"wsdl_wrapped_faults":       types.BoolType,
	"allow_soap_enc_array":      types.BoolType,
	"validate_soap_enc_array":   types.BoolType,
	"wildcards_ignore_xsi_type": types.BoolType,
	"wsdl_strict_soap_version":  types.BoolType,
	"xacml_debug":               types.BoolType,
	"allow_xop_include":         types.BoolType,
}

func (data CompileSettings) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CompileSettings"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CompileSettings) IsNull() bool {
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

func (data CompileSettings) ToBody(ctx context.Context, pathRoot string) string {
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
		body, _ = sjson.Set(body, pathRoot+`Profile`, tfutils.StringFromBool(data.Profile, ""))
	}
	if !data.Debug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Debug`, tfutils.StringFromBool(data.Debug, ""))
	}
	if !data.Stream.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Stream`, tfutils.StringFromBool(data.Stream, ""))
	}
	if !data.TryStream.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TryStream`, tfutils.StringFromBool(data.TryStream, ""))
	}
	if !data.MinimumEscaping.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MinimumEscaping`, tfutils.StringFromBool(data.MinimumEscaping, ""))
	}
	if !data.StackSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StackSize`, data.StackSize.ValueInt64())
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
		body, _ = sjson.Set(body, pathRoot+`AllowSoapEncArray`, tfutils.StringFromBool(data.AllowSoapEncArray, ""))
	}
	if !data.ValidateSoapEncArray.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateSoapEncArray`, tfutils.StringFromBool(data.ValidateSoapEncArray, ""))
	}
	if !data.WildcardsIgnoreXsiType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WildcardsIgnoreXsiType`, tfutils.StringFromBool(data.WildcardsIgnoreXsiType, ""))
	}
	if !data.WsdlStrictSoapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLStrictSOAPVersion`, tfutils.StringFromBool(data.WsdlStrictSoapVersion, ""))
	}
	if !data.XacmlDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XACMLDebug`, tfutils.StringFromBool(data.XacmlDebug, ""))
	}
	if !data.AllowXopInclude.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowXOPInclude`, tfutils.StringFromBool(data.AllowXopInclude, ""))
	}
	return body
}

func (data *CompileSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Profile`); value.Exists() {
		data.Profile = tfutils.BoolFromString(value.String())
	} else {
		data.Profile = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Debug`); value.Exists() {
		data.Debug = tfutils.BoolFromString(value.String())
	} else {
		data.Debug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Stream`); value.Exists() {
		data.Stream = tfutils.BoolFromString(value.String())
	} else {
		data.Stream = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TryStream`); value.Exists() {
		data.TryStream = tfutils.BoolFromString(value.String())
	} else {
		data.TryStream = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MinimumEscaping`); value.Exists() {
		data.MinimumEscaping = tfutils.BoolFromString(value.String())
	} else {
		data.MinimumEscaping = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StackSize`); value.Exists() {
		data.StackSize = types.Int64Value(value.Int())
	} else {
		data.StackSize = types.Int64Value(1048576)
	}
	if value := res.Get(pathRoot + `WSIValidation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsiValidation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsiValidation = types.StringValue("warn")
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
	if value := res.Get(pathRoot + `AllowSoapEncArray`); value.Exists() {
		data.AllowSoapEncArray = tfutils.BoolFromString(value.String())
	} else {
		data.AllowSoapEncArray = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ValidateSoapEncArray`); value.Exists() {
		data.ValidateSoapEncArray = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateSoapEncArray = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WildcardsIgnoreXsiType`); value.Exists() {
		data.WildcardsIgnoreXsiType = tfutils.BoolFromString(value.String())
	} else {
		data.WildcardsIgnoreXsiType = types.BoolNull()
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
	if value := res.Get(pathRoot + `AllowXOPInclude`); value.Exists() {
		data.AllowXopInclude = tfutils.BoolFromString(value.String())
	} else {
		data.AllowXopInclude = types.BoolNull()
	}
}

func (data *CompileSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if !data.Strict.ValueBool() {
		data.Strict = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Profile`); value.Exists() && !data.Profile.IsNull() {
		data.Profile = tfutils.BoolFromString(value.String())
	} else if data.Profile.ValueBool() {
		data.Profile = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Debug`); value.Exists() && !data.Debug.IsNull() {
		data.Debug = tfutils.BoolFromString(value.String())
	} else if data.Debug.ValueBool() {
		data.Debug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Stream`); value.Exists() && !data.Stream.IsNull() {
		data.Stream = tfutils.BoolFromString(value.String())
	} else if data.Stream.ValueBool() {
		data.Stream = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TryStream`); value.Exists() && !data.TryStream.IsNull() {
		data.TryStream = tfutils.BoolFromString(value.String())
	} else if data.TryStream.ValueBool() {
		data.TryStream = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MinimumEscaping`); value.Exists() && !data.MinimumEscaping.IsNull() {
		data.MinimumEscaping = tfutils.BoolFromString(value.String())
	} else if data.MinimumEscaping.ValueBool() {
		data.MinimumEscaping = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StackSize`); value.Exists() && !data.StackSize.IsNull() {
		data.StackSize = types.Int64Value(value.Int())
	} else if data.StackSize.ValueInt64() != 1048576 {
		data.StackSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSIValidation`); value.Exists() && !data.WsiValidation.IsNull() {
		data.WsiValidation = tfutils.ParseStringFromGJSON(value)
	} else if data.WsiValidation.ValueString() != "warn" {
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
		data.AllowSoapEncArray = tfutils.BoolFromString(value.String())
	} else if data.AllowSoapEncArray.ValueBool() {
		data.AllowSoapEncArray = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ValidateSoapEncArray`); value.Exists() && !data.ValidateSoapEncArray.IsNull() {
		data.ValidateSoapEncArray = tfutils.BoolFromString(value.String())
	} else if data.ValidateSoapEncArray.ValueBool() {
		data.ValidateSoapEncArray = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WildcardsIgnoreXsiType`); value.Exists() && !data.WildcardsIgnoreXsiType.IsNull() {
		data.WildcardsIgnoreXsiType = tfutils.BoolFromString(value.String())
	} else if data.WildcardsIgnoreXsiType.ValueBool() {
		data.WildcardsIgnoreXsiType = types.BoolNull()
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
		data.AllowXopInclude = tfutils.BoolFromString(value.String())
	} else if data.AllowXopInclude.ValueBool() {
		data.AllowXopInclude = types.BoolNull()
	}
}
