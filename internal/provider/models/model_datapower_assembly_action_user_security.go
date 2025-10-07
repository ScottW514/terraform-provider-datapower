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

type AssemblyActionUserSecurity struct {
	Id                                types.String                `tfsdk:"id"`
	AppDomain                         types.String                `tfsdk:"app_domain"`
	FactorId                          types.String                `tfsdk:"factor_id"`
	ExtractIdentityMethod             types.String                `tfsdk:"extract_identity_method"`
	EiStopOnError                     types.Bool                  `tfsdk:"ei_stop_on_error"`
	UserContextVariable               types.String                `tfsdk:"user_context_variable"`
	PassContextVariable               types.String                `tfsdk:"pass_context_variable"`
	RedirectUrl                       types.String                `tfsdk:"redirect_url"`
	RedirectTimeLimit                 types.Int64                 `tfsdk:"redirect_time_limit"`
	QueryParameters                   types.String                `tfsdk:"query_parameters"`
	EiDefaultForm                     types.Bool                  `tfsdk:"ei_default_form"`
	EiCustomForm                      types.String                `tfsdk:"ei_custom_form"`
	EiCustomFormClientProfile         types.String                `tfsdk:"ei_custom_form_client_profile"`
	EiCustomFormContentSecurityPolicy types.String                `tfsdk:"ei_custom_form_content_security_policy"`
	EiFormTimeLimit                   types.Int64                 `tfsdk:"ei_form_time_limit"`
	UserAuthMethod                    types.String                `tfsdk:"user_auth_method"`
	AuStopOnError                     types.Bool                  `tfsdk:"au_stop_on_error"`
	UserRegistry                      types.String                `tfsdk:"user_registry"`
	AuthResponseHeadersPattern        types.String                `tfsdk:"auth_response_headers_pattern"`
	AuthResponseCredentialHeader      types.String                `tfsdk:"auth_response_credential_header"`
	UserAzMethod                      types.String                `tfsdk:"user_az_method"`
	AzStopOnError                     types.Bool                  `tfsdk:"az_stop_on_error"`
	AzDefaultForm                     types.Bool                  `tfsdk:"az_default_form"`
	AzCustomForm                      types.String                `tfsdk:"az_custom_form"`
	AzCustomFormClientProfile         types.String                `tfsdk:"az_custom_form_client_profile"`
	AzCustomFormContentSecurityPolicy types.String                `tfsdk:"az_custom_form_content_security_policy"`
	AzFormTimeLimit                   types.Int64                 `tfsdk:"az_form_time_limit"`
	AzTableDisplayCheckboxes          types.Bool                  `tfsdk:"az_table_display_checkboxes"`
	AzTableDynamicEntries             types.String                `tfsdk:"az_table_dynamic_entries"`
	AzTableDefaultEntry               types.List                  `tfsdk:"az_table_default_entry"`
	Hostname                          types.String                `tfsdk:"hostname"`
	UserSummary                       types.String                `tfsdk:"user_summary"`
	Title                             types.String                `tfsdk:"title"`
	CorrelationPath                   types.String                `tfsdk:"correlation_path"`
	ActionDebug                       types.Bool                  `tfsdk:"action_debug"`
	DependencyActions                 []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AssemblyActionUserSecurityEIStopOnErrorCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"disabled"},
}
var AssemblyActionUserSecurityUserContextVariableCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"context-var"},
}
var AssemblyActionUserSecurityPassContextVariableCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"context-var"},
}
var AssemblyActionUserSecurityRedirectURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"redirect"},
}
var AssemblyActionUserSecurityRedirectTimeLimitCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"redirect"},
}
var AssemblyActionUserSecurityEIDefaultFormCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityEICustomFormCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "extract_identity_method",
			AttrType:    "String",
			AttrDefault: "basic",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ei_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}
var AssemblyActionUserSecurityEIFormTimeLimitCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityAUStopOnErrorCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_auth_method",
	AttrType:    "String",
	AttrDefault: "user-registry",
	Value:       []string{"disabled"},
}
var AssemblyActionUserSecurityUserRegistryCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_auth_method",
	AttrType:    "String",
	AttrDefault: "user-registry",
	Value:       []string{"disabled"},
}
var AssemblyActionUserSecurityAZStopOnErrorCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"disabled"},
}
var AssemblyActionUserSecurityAZDefaultFormCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityAZCustomFormCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "user_az_method",
			AttrType:    "String",
			AttrDefault: "authenticated",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}
var AssemblyActionUserSecurityAZFormTimeLimitCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityEIStopOnErrorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityUserContextVariableIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityPassContextVariableIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityRedirectURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityRedirectTimeLimitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityQueryParametersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "extract_identity_method",
	AttrType:    "String",
	AttrDefault: "basic",
	Value:       []string{"redirect"},
}
var AssemblyActionUserSecurityEIDefaultFormIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityEICustomFormIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityEICustomFormClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "extract_identity_method",
			AttrType:    "String",
			AttrDefault: "basic",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ei_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var AssemblyActionUserSecurityEICustomFormContentSecurityPolicyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "extract_identity_method",
			AttrType:    "String",
			AttrDefault: "basic",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ei_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var AssemblyActionUserSecurityEIFormTimeLimitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAUStopOnErrorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityUserRegistryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAuthResponseHeadersPatternIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_auth_method",
	AttrType:    "String",
	AttrDefault: "user-registry",
	Value:       []string{"user-registry"},
}
var AssemblyActionUserSecurityAuthResponseCredentialHeaderIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_auth_method",
	AttrType:    "String",
	AttrDefault: "user-registry",
	Value:       []string{"user-registry"},
}
var AssemblyActionUserSecurityAZStopOnErrorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAZDefaultFormIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAZCustomFormIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAZCustomFormClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "user_az_method",
			AttrType:    "String",
			AttrDefault: "authenticated",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var AssemblyActionUserSecurityAZCustomFormContentSecurityPolicyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "user_az_method",
			AttrType:    "String",
			AttrDefault: "authenticated",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_default_form",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}
var AssemblyActionUserSecurityAZFormTimeLimitIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var AssemblyActionUserSecurityAZTableDisplayCheckboxesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityAZTableDynamicEntriesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityAZTableDefaultEntryIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "user_az_method",
	AttrType:    "String",
	AttrDefault: "authenticated",
	Value:       []string{"html-form"},
}
var AssemblyActionUserSecurityHostnameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "extract_identity_method",
			AttrType:    "String",
			AttrDefault: "basic",
			Value:       []string{"redirect"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "extract_identity_method",
			AttrType:    "String",
			AttrDefault: "basic",
			Value:       []string{"html-form"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "user_az_method",
			AttrType:    "String",
			AttrDefault: "authenticated",
			Value:       []string{"html-form"},
		},
	},
}

var AssemblyActionUserSecurityObjectType = map[string]attr.Type{
	"id":                                     types.StringType,
	"app_domain":                             types.StringType,
	"factor_id":                              types.StringType,
	"extract_identity_method":                types.StringType,
	"ei_stop_on_error":                       types.BoolType,
	"user_context_variable":                  types.StringType,
	"pass_context_variable":                  types.StringType,
	"redirect_url":                           types.StringType,
	"redirect_time_limit":                    types.Int64Type,
	"query_parameters":                       types.StringType,
	"ei_default_form":                        types.BoolType,
	"ei_custom_form":                         types.StringType,
	"ei_custom_form_client_profile":          types.StringType,
	"ei_custom_form_content_security_policy": types.StringType,
	"ei_form_time_limit":                     types.Int64Type,
	"user_auth_method":                       types.StringType,
	"au_stop_on_error":                       types.BoolType,
	"user_registry":                          types.StringType,
	"auth_response_headers_pattern":          types.StringType,
	"auth_response_credential_header":        types.StringType,
	"user_az_method":                         types.StringType,
	"az_stop_on_error":                       types.BoolType,
	"az_default_form":                        types.BoolType,
	"az_custom_form":                         types.StringType,
	"az_custom_form_client_profile":          types.StringType,
	"az_custom_form_content_security_policy": types.StringType,
	"az_form_time_limit":                     types.Int64Type,
	"az_table_display_checkboxes":            types.BoolType,
	"az_table_dynamic_entries":               types.StringType,
	"az_table_default_entry":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmTableEntryObjectType}},
	"hostname":                               types.StringType,
	"user_summary":                           types.StringType,
	"title":                                  types.StringType,
	"correlation_path":                       types.StringType,
	"action_debug":                           types.BoolType,
	"dependency_actions":                     actions.ActionsListType,
}

func (data AssemblyActionUserSecurity) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AssemblyActionUserSecurity"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AssemblyActionUserSecurity) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.FactorId.IsNull() {
		return false
	}
	if !data.ExtractIdentityMethod.IsNull() {
		return false
	}
	if !data.EiStopOnError.IsNull() {
		return false
	}
	if !data.UserContextVariable.IsNull() {
		return false
	}
	if !data.PassContextVariable.IsNull() {
		return false
	}
	if !data.RedirectUrl.IsNull() {
		return false
	}
	if !data.RedirectTimeLimit.IsNull() {
		return false
	}
	if !data.QueryParameters.IsNull() {
		return false
	}
	if !data.EiDefaultForm.IsNull() {
		return false
	}
	if !data.EiCustomForm.IsNull() {
		return false
	}
	if !data.EiCustomFormClientProfile.IsNull() {
		return false
	}
	if !data.EiCustomFormContentSecurityPolicy.IsNull() {
		return false
	}
	if !data.EiFormTimeLimit.IsNull() {
		return false
	}
	if !data.UserAuthMethod.IsNull() {
		return false
	}
	if !data.AuStopOnError.IsNull() {
		return false
	}
	if !data.UserRegistry.IsNull() {
		return false
	}
	if !data.AuthResponseHeadersPattern.IsNull() {
		return false
	}
	if !data.AuthResponseCredentialHeader.IsNull() {
		return false
	}
	if !data.UserAzMethod.IsNull() {
		return false
	}
	if !data.AzStopOnError.IsNull() {
		return false
	}
	if !data.AzDefaultForm.IsNull() {
		return false
	}
	if !data.AzCustomForm.IsNull() {
		return false
	}
	if !data.AzCustomFormClientProfile.IsNull() {
		return false
	}
	if !data.AzCustomFormContentSecurityPolicy.IsNull() {
		return false
	}
	if !data.AzFormTimeLimit.IsNull() {
		return false
	}
	if !data.AzTableDisplayCheckboxes.IsNull() {
		return false
	}
	if !data.AzTableDynamicEntries.IsNull() {
		return false
	}
	if !data.AzTableDefaultEntry.IsNull() {
		return false
	}
	if !data.Hostname.IsNull() {
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

func (data AssemblyActionUserSecurity) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.FactorId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FactorID`, data.FactorId.ValueString())
	}
	if !data.ExtractIdentityMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExtractIdentityMethod`, data.ExtractIdentityMethod.ValueString())
	}
	if !data.EiStopOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIStopOnError`, tfutils.StringFromBool(data.EiStopOnError, ""))
	}
	if !data.UserContextVariable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserContextVariable`, data.UserContextVariable.ValueString())
	}
	if !data.PassContextVariable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PassContextVariable`, data.PassContextVariable.ValueString())
	}
	if !data.RedirectUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RedirectURL`, data.RedirectUrl.ValueString())
	}
	if !data.RedirectTimeLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RedirectTimeLimit`, data.RedirectTimeLimit.ValueInt64())
	}
	if !data.QueryParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryParameters`, data.QueryParameters.ValueString())
	}
	if !data.EiDefaultForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIDefaultForm`, tfutils.StringFromBool(data.EiDefaultForm, ""))
	}
	if !data.EiCustomForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EICustomForm`, data.EiCustomForm.ValueString())
	}
	if !data.EiCustomFormClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EICustomFormClientProfile`, data.EiCustomFormClientProfile.ValueString())
	}
	if !data.EiCustomFormContentSecurityPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EICustomFormContentSecurityPolicy`, data.EiCustomFormContentSecurityPolicy.ValueString())
	}
	if !data.EiFormTimeLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIFormTimeLimit`, data.EiFormTimeLimit.ValueInt64())
	}
	if !data.UserAuthMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserAuthMethod`, data.UserAuthMethod.ValueString())
	}
	if !data.AuStopOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUStopOnError`, tfutils.StringFromBool(data.AuStopOnError, ""))
	}
	if !data.UserRegistry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserRegistry`, data.UserRegistry.ValueString())
	}
	if !data.AuthResponseHeadersPattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthResponseHeadersPattern`, data.AuthResponseHeadersPattern.ValueString())
	}
	if !data.AuthResponseCredentialHeader.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthResponseCredentialHeader`, data.AuthResponseCredentialHeader.ValueString())
	}
	if !data.UserAzMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserAZMethod`, data.UserAzMethod.ValueString())
	}
	if !data.AzStopOnError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZStopOnError`, tfutils.StringFromBool(data.AzStopOnError, ""))
	}
	if !data.AzDefaultForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZDefaultForm`, tfutils.StringFromBool(data.AzDefaultForm, ""))
	}
	if !data.AzCustomForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCustomForm`, data.AzCustomForm.ValueString())
	}
	if !data.AzCustomFormClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCustomFormClientProfile`, data.AzCustomFormClientProfile.ValueString())
	}
	if !data.AzCustomFormContentSecurityPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZCustomFormContentSecurityPolicy`, data.AzCustomFormContentSecurityPolicy.ValueString())
	}
	if !data.AzFormTimeLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZFormTimeLimit`, data.AzFormTimeLimit.ValueInt64())
	}
	if !data.AzTableDisplayCheckboxes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTableDisplayCheckboxes`, tfutils.StringFromBool(data.AzTableDisplayCheckboxes, ""))
	}
	if !data.AzTableDynamicEntries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AZTableDynamicEntries`, data.AzTableDynamicEntries.ValueString())
	}
	if !data.AzTableDefaultEntry.IsNull() {
		var dataValues []DmTableEntry
		data.AzTableDefaultEntry.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AZTableDefaultEntry`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Hostname`, data.Hostname.ValueString())
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

func (data *AssemblyActionUserSecurity) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `FactorID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FactorId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FactorId = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `ExtractIdentityMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExtractIdentityMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExtractIdentityMethod = types.StringValue("basic")
	}
	if value := res.Get(pathRoot + `EIStopOnError`); value.Exists() {
		data.EiStopOnError = tfutils.BoolFromString(value.String())
	} else {
		data.EiStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserContextVariable`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserContextVariable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserContextVariable = types.StringNull()
	}
	if value := res.Get(pathRoot + `PassContextVariable`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PassContextVariable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PassContextVariable = types.StringNull()
	}
	if value := res.Get(pathRoot + `RedirectURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RedirectTimeLimit`); value.Exists() {
		data.RedirectTimeLimit = types.Int64Value(value.Int())
	} else {
		data.RedirectTimeLimit = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `QueryParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueryParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIDefaultForm`); value.Exists() {
		data.EiDefaultForm = tfutils.BoolFromString(value.String())
	} else {
		data.EiDefaultForm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EICustomForm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiCustomForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICustomFormClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiCustomFormClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomFormClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICustomFormContentSecurityPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiCustomFormContentSecurityPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomFormContentSecurityPolicy = types.StringValue("default-src 'self'")
	}
	if value := res.Get(pathRoot + `EIFormTimeLimit`); value.Exists() {
		data.EiFormTimeLimit = types.Int64Value(value.Int())
	} else {
		data.EiFormTimeLimit = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `UserAuthMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAuthMethod = types.StringValue("user-registry")
	}
	if value := res.Get(pathRoot + `AUStopOnError`); value.Exists() {
		data.AuStopOnError = tfutils.BoolFromString(value.String())
	} else {
		data.AuStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserRegistry`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserRegistry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserRegistry = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthResponseHeadersPattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthResponseHeadersPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthResponseHeadersPattern = types.StringValue("(?i)x-api*")
	}
	if value := res.Get(pathRoot + `AuthResponseCredentialHeader`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthResponseCredentialHeader = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthResponseCredentialHeader = types.StringValue("X-API-Authenticated-Credential")
	}
	if value := res.Get(pathRoot + `UserAZMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserAzMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAzMethod = types.StringValue("authenticated")
	}
	if value := res.Get(pathRoot + `AZStopOnError`); value.Exists() {
		data.AzStopOnError = tfutils.BoolFromString(value.String())
	} else {
		data.AzStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZDefaultForm`); value.Exists() {
		data.AzDefaultForm = tfutils.BoolFromString(value.String())
	} else {
		data.AzDefaultForm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZCustomForm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCustomForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCustomFormClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCustomFormClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomFormClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCustomFormContentSecurityPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzCustomFormContentSecurityPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomFormContentSecurityPolicy = types.StringValue("default-src 'self'")
	}
	if value := res.Get(pathRoot + `AZFormTimeLimit`); value.Exists() {
		data.AzFormTimeLimit = types.Int64Value(value.Int())
	} else {
		data.AzFormTimeLimit = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AZTableDisplayCheckboxes`); value.Exists() {
		data.AzTableDisplayCheckboxes = tfutils.BoolFromString(value.String())
	} else {
		data.AzTableDisplayCheckboxes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTableDynamicEntries`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AzTableDynamicEntries = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzTableDynamicEntries = types.StringValue("user.default.az.dynamic_entries")
	}
	if value := res.Get(pathRoot + `AZTableDefaultEntry`); value.Exists() {
		l := []DmTableEntry{}
		if value := res.Get(`AZTableDefaultEntry`); value.Exists() {
			for _, v := range value.Array() {
				item := DmTableEntry{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AzTableDefaultEntry, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTableEntryObjectType}, l)
		} else {
			data.AzTableDefaultEntry = types.ListNull(types.ObjectType{AttrTypes: DmTableEntryObjectType})
		}
	} else {
		data.AzTableDefaultEntry = types.ListNull(types.ObjectType{AttrTypes: DmTableEntryObjectType})
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
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

func (data *AssemblyActionUserSecurity) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `FactorID`); value.Exists() && !data.FactorId.IsNull() {
		data.FactorId = tfutils.ParseStringFromGJSON(value)
	} else if data.FactorId.ValueString() != "default" {
		data.FactorId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExtractIdentityMethod`); value.Exists() && !data.ExtractIdentityMethod.IsNull() {
		data.ExtractIdentityMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.ExtractIdentityMethod.ValueString() != "basic" {
		data.ExtractIdentityMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIStopOnError`); value.Exists() && !data.EiStopOnError.IsNull() {
		data.EiStopOnError = tfutils.BoolFromString(value.String())
	} else if !data.EiStopOnError.ValueBool() {
		data.EiStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserContextVariable`); value.Exists() && !data.UserContextVariable.IsNull() {
		data.UserContextVariable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserContextVariable = types.StringNull()
	}
	if value := res.Get(pathRoot + `PassContextVariable`); value.Exists() && !data.PassContextVariable.IsNull() {
		data.PassContextVariable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PassContextVariable = types.StringNull()
	}
	if value := res.Get(pathRoot + `RedirectURL`); value.Exists() && !data.RedirectUrl.IsNull() {
		data.RedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RedirectTimeLimit`); value.Exists() && !data.RedirectTimeLimit.IsNull() {
		data.RedirectTimeLimit = types.Int64Value(value.Int())
	} else if data.RedirectTimeLimit.ValueInt64() != 300 {
		data.RedirectTimeLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `QueryParameters`); value.Exists() && !data.QueryParameters.IsNull() {
		data.QueryParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIDefaultForm`); value.Exists() && !data.EiDefaultForm.IsNull() {
		data.EiDefaultForm = tfutils.BoolFromString(value.String())
	} else if !data.EiDefaultForm.ValueBool() {
		data.EiDefaultForm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EICustomForm`); value.Exists() && !data.EiCustomForm.IsNull() {
		data.EiCustomForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICustomFormClientProfile`); value.Exists() && !data.EiCustomFormClientProfile.IsNull() {
		data.EiCustomFormClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomFormClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICustomFormContentSecurityPolicy`); value.Exists() && !data.EiCustomFormContentSecurityPolicy.IsNull() {
		data.EiCustomFormContentSecurityPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.EiCustomFormContentSecurityPolicy.ValueString() != "default-src 'self'" {
		data.EiCustomFormContentSecurityPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIFormTimeLimit`); value.Exists() && !data.EiFormTimeLimit.IsNull() {
		data.EiFormTimeLimit = types.Int64Value(value.Int())
	} else if data.EiFormTimeLimit.ValueInt64() != 300 {
		data.EiFormTimeLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserAuthMethod`); value.Exists() && !data.UserAuthMethod.IsNull() {
		data.UserAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.UserAuthMethod.ValueString() != "user-registry" {
		data.UserAuthMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUStopOnError`); value.Exists() && !data.AuStopOnError.IsNull() {
		data.AuStopOnError = tfutils.BoolFromString(value.String())
	} else if !data.AuStopOnError.ValueBool() {
		data.AuStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserRegistry`); value.Exists() && !data.UserRegistry.IsNull() {
		data.UserRegistry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserRegistry = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthResponseHeadersPattern`); value.Exists() && !data.AuthResponseHeadersPattern.IsNull() {
		data.AuthResponseHeadersPattern = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthResponseHeadersPattern.ValueString() != "(?i)x-api*" {
		data.AuthResponseHeadersPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthResponseCredentialHeader`); value.Exists() && !data.AuthResponseCredentialHeader.IsNull() {
		data.AuthResponseCredentialHeader = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthResponseCredentialHeader.ValueString() != "X-API-Authenticated-Credential" {
		data.AuthResponseCredentialHeader = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserAZMethod`); value.Exists() && !data.UserAzMethod.IsNull() {
		data.UserAzMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.UserAzMethod.ValueString() != "authenticated" {
		data.UserAzMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZStopOnError`); value.Exists() && !data.AzStopOnError.IsNull() {
		data.AzStopOnError = tfutils.BoolFromString(value.String())
	} else if !data.AzStopOnError.ValueBool() {
		data.AzStopOnError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZDefaultForm`); value.Exists() && !data.AzDefaultForm.IsNull() {
		data.AzDefaultForm = tfutils.BoolFromString(value.String())
	} else if !data.AzDefaultForm.ValueBool() {
		data.AzDefaultForm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZCustomForm`); value.Exists() && !data.AzCustomForm.IsNull() {
		data.AzCustomForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCustomFormClientProfile`); value.Exists() && !data.AzCustomFormClientProfile.IsNull() {
		data.AzCustomFormClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AzCustomFormClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZCustomFormContentSecurityPolicy`); value.Exists() && !data.AzCustomFormContentSecurityPolicy.IsNull() {
		data.AzCustomFormContentSecurityPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.AzCustomFormContentSecurityPolicy.ValueString() != "default-src 'self'" {
		data.AzCustomFormContentSecurityPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZFormTimeLimit`); value.Exists() && !data.AzFormTimeLimit.IsNull() {
		data.AzFormTimeLimit = types.Int64Value(value.Int())
	} else if data.AzFormTimeLimit.ValueInt64() != 300 {
		data.AzFormTimeLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AZTableDisplayCheckboxes`); value.Exists() && !data.AzTableDisplayCheckboxes.IsNull() {
		data.AzTableDisplayCheckboxes = tfutils.BoolFromString(value.String())
	} else if data.AzTableDisplayCheckboxes.ValueBool() {
		data.AzTableDisplayCheckboxes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AZTableDynamicEntries`); value.Exists() && !data.AzTableDynamicEntries.IsNull() {
		data.AzTableDynamicEntries = tfutils.ParseStringFromGJSON(value)
	} else if data.AzTableDynamicEntries.ValueString() != "user.default.az.dynamic_entries" {
		data.AzTableDynamicEntries = types.StringNull()
	}
	if value := res.Get(pathRoot + `AZTableDefaultEntry`); value.Exists() && !data.AzTableDefaultEntry.IsNull() {
		l := []DmTableEntry{}
		e := []DmTableEntry{}
		data.AzTableDefaultEntry.ElementsAs(ctx, &e, false)
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
				item := DmTableEntry{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AzTableDefaultEntry, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTableEntryObjectType}, l)
		} else {
			data.AzTableDefaultEntry = types.ListNull(types.ObjectType{AttrTypes: DmTableEntryObjectType})
		}
	} else {
		data.AzTableDefaultEntry = types.ListNull(types.ObjectType{AttrTypes: DmTableEntryObjectType})
	}
	if value := res.Get(pathRoot + `Hostname`); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Hostname = types.StringNull()
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
