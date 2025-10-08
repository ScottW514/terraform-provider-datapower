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

type APILDAPRegistry struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	LdapHost               types.String                `tfsdk:"ldap_host"`
	LdapPort               types.Int64                 `tfsdk:"ldap_port"`
	SslClientProfile       types.String                `tfsdk:"ssl_client_profile"`
	LdapVersion            types.String                `tfsdk:"ldap_version"`
	LdapAuthMethod         types.String                `tfsdk:"ldap_auth_method"`
	LdapBindDn             types.String                `tfsdk:"ldap_bind_dn"`
	LdapBindPasswordAlias  types.String                `tfsdk:"ldap_bind_password_alias"`
	LdapSearchParameters   types.String                `tfsdk:"ldap_search_parameters"`
	LdapReadTimeout        types.Int64                 `tfsdk:"ldap_read_timeout"`
	LdapGroupAuthEnabled   types.Bool                  `tfsdk:"ldap_group_auth_enabled"`
	LdapGroupAuthType      types.String                `tfsdk:"ldap_group_auth_type"`
	LdapGroupScope         types.String                `tfsdk:"ldap_group_scope"`
	LdapGroupBaseDn        types.String                `tfsdk:"ldap_group_base_dn"`
	LdapGroupFilterPrefix  types.String                `tfsdk:"ldap_group_filter_prefix"`
	LdapGroupFilterSuffix  types.String                `tfsdk:"ldap_group_filter_suffix"`
	LdapGroupDynamicFilter types.String                `tfsdk:"ldap_group_dynamic_filter"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APILDAPRegistryLDAPGroupAuthTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_group_auth_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APILDAPRegistryLDAPGroupAuthTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ldap_group_auth_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APILDAPRegistryLDAPGroupScopeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_group_auth_enabled",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var APILDAPRegistryLDAPGroupBaseDNCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"static"},
		},
	},
}

var APILDAPRegistryLDAPGroupBaseDNIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_enabled",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"dynamic"},
				},
			},
		},
	},
}

var APILDAPRegistryLDAPGroupFilterPrefixCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"static"},
		},
	},
}

var APILDAPRegistryLDAPGroupFilterPrefixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_enabled",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"dynamic"},
				},
			},
		},
	},
}

var APILDAPRegistryLDAPGroupFilterSuffixCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"static"},
		},
	},
}

var APILDAPRegistryLDAPGroupFilterSuffixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_enabled",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"dynamic"},
				},
			},
		},
	},
}

var APILDAPRegistryLDAPGroupDynamicFilterCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"dynamic"},
		},
	},
}

var APILDAPRegistryLDAPGroupDynamicFilterIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ldap_group_auth_enabled",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_enabled",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ldap_group_auth_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"static"},
				},
			},
		},
	},
}

var APILDAPRegistryObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"ldap_host":                 types.StringType,
	"ldap_port":                 types.Int64Type,
	"ssl_client_profile":        types.StringType,
	"ldap_version":              types.StringType,
	"ldap_auth_method":          types.StringType,
	"ldap_bind_dn":              types.StringType,
	"ldap_bind_password_alias":  types.StringType,
	"ldap_search_parameters":    types.StringType,
	"ldap_read_timeout":         types.Int64Type,
	"ldap_group_auth_enabled":   types.BoolType,
	"ldap_group_auth_type":      types.StringType,
	"ldap_group_scope":          types.StringType,
	"ldap_group_base_dn":        types.StringType,
	"ldap_group_filter_prefix":  types.StringType,
	"ldap_group_filter_suffix":  types.StringType,
	"ldap_group_dynamic_filter": types.StringType,
	"dependency_actions":        actions.ActionsListType,
}

func (data APILDAPRegistry) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APILDAPRegistry"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APILDAPRegistry) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LdapHost.IsNull() {
		return false
	}
	if !data.LdapPort.IsNull() {
		return false
	}
	if !data.SslClientProfile.IsNull() {
		return false
	}
	if !data.LdapVersion.IsNull() {
		return false
	}
	if !data.LdapAuthMethod.IsNull() {
		return false
	}
	if !data.LdapBindDn.IsNull() {
		return false
	}
	if !data.LdapBindPasswordAlias.IsNull() {
		return false
	}
	if !data.LdapSearchParameters.IsNull() {
		return false
	}
	if !data.LdapReadTimeout.IsNull() {
		return false
	}
	if !data.LdapGroupAuthEnabled.IsNull() {
		return false
	}
	if !data.LdapGroupAuthType.IsNull() {
		return false
	}
	if !data.LdapGroupScope.IsNull() {
		return false
	}
	if !data.LdapGroupBaseDn.IsNull() {
		return false
	}
	if !data.LdapGroupFilterPrefix.IsNull() {
		return false
	}
	if !data.LdapGroupFilterSuffix.IsNull() {
		return false
	}
	if !data.LdapGroupDynamicFilter.IsNull() {
		return false
	}
	return true
}

func (data APILDAPRegistry) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LdapHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPHost`, data.LdapHost.ValueString())
	}
	if !data.LdapPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPPort`, data.LdapPort.ValueInt64())
	}
	if !data.SslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientProfile`, data.SslClientProfile.ValueString())
	}
	if !data.LdapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPVersion`, data.LdapVersion.ValueString())
	}
	if !data.LdapAuthMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPAuthMethod`, data.LdapAuthMethod.ValueString())
	}
	if !data.LdapBindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPBindDN`, data.LdapBindDn.ValueString())
	}
	if !data.LdapBindPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPBindPasswordAlias`, data.LdapBindPasswordAlias.ValueString())
	}
	if !data.LdapSearchParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSearchParameters`, data.LdapSearchParameters.ValueString())
	}
	if !data.LdapReadTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPReadTimeout`, data.LdapReadTimeout.ValueInt64())
	}
	if !data.LdapGroupAuthEnabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupAuthEnabled`, tfutils.StringFromBool(data.LdapGroupAuthEnabled, ""))
	}
	if !data.LdapGroupAuthType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupAuthType`, data.LdapGroupAuthType.ValueString())
	}
	if !data.LdapGroupScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupScope`, data.LdapGroupScope.ValueString())
	}
	if !data.LdapGroupBaseDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupBaseDN`, data.LdapGroupBaseDn.ValueString())
	}
	if !data.LdapGroupFilterPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupFilterPrefix`, data.LdapGroupFilterPrefix.ValueString())
	}
	if !data.LdapGroupFilterSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupFilterSuffix`, data.LdapGroupFilterSuffix.ValueString())
	}
	if !data.LdapGroupDynamicFilter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPGroupDynamicFilter`, data.LdapGroupDynamicFilter.ValueString())
	}
	return body
}

func (data *APILDAPRegistry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LDAPHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPPort`); value.Exists() {
		data.LdapPort = types.Int64Value(value.Int())
	} else {
		data.LdapPort = types.Int64Value(636)
	}
	if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapVersion = types.StringValue("v3")
	}
	if value := res.Get(pathRoot + `LDAPAuthMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapAuthMethod = types.StringValue("searchDN")
	}
	if value := res.Get(pathRoot + `LDAPBindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPBindPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSearchParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPReadTimeout`); value.Exists() {
		data.LdapReadTimeout = types.Int64Value(value.Int())
	} else {
		data.LdapReadTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `LDAPGroupAuthEnabled`); value.Exists() {
		data.LdapGroupAuthEnabled = tfutils.BoolFromString(value.String())
	} else {
		data.LdapGroupAuthEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupAuthType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupAuthType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupAuthType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupScope = types.StringValue("subtree")
	}
	if value := res.Get(pathRoot + `LDAPGroupBaseDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupBaseDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupBaseDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupFilterPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupFilterPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupFilterPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupFilterSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupFilterSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupFilterSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupDynamicFilter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapGroupDynamicFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupDynamicFilter = types.StringNull()
	}
}

func (data *APILDAPRegistry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LDAPHost`); value.Exists() && !data.LdapHost.IsNull() {
		data.LdapHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPPort`); value.Exists() && !data.LdapPort.IsNull() {
		data.LdapPort = types.Int64Value(value.Int())
	} else if data.LdapPort.ValueInt64() != 636 {
		data.LdapPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && !data.SslClientProfile.IsNull() {
		data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && !data.LdapVersion.IsNull() {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapVersion.ValueString() != "v3" {
		data.LdapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPAuthMethod`); value.Exists() && !data.LdapAuthMethod.IsNull() {
		data.LdapAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapAuthMethod.ValueString() != "searchDN" {
		data.LdapAuthMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPBindDN`); value.Exists() && !data.LdapBindDn.IsNull() {
		data.LdapBindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPBindPasswordAlias`); value.Exists() && !data.LdapBindPasswordAlias.IsNull() {
		data.LdapBindPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBindPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSearchParameters`); value.Exists() && !data.LdapSearchParameters.IsNull() {
		data.LdapSearchParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSearchParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPReadTimeout`); value.Exists() && !data.LdapReadTimeout.IsNull() {
		data.LdapReadTimeout = types.Int64Value(value.Int())
	} else if data.LdapReadTimeout.ValueInt64() != 60 {
		data.LdapReadTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LDAPGroupAuthEnabled`); value.Exists() && !data.LdapGroupAuthEnabled.IsNull() {
		data.LdapGroupAuthEnabled = tfutils.BoolFromString(value.String())
	} else if data.LdapGroupAuthEnabled.ValueBool() {
		data.LdapGroupAuthEnabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupAuthType`); value.Exists() && !data.LdapGroupAuthType.IsNull() {
		data.LdapGroupAuthType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupAuthType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupScope`); value.Exists() && !data.LdapGroupScope.IsNull() {
		data.LdapGroupScope = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapGroupScope.ValueString() != "subtree" {
		data.LdapGroupScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupBaseDN`); value.Exists() && !data.LdapGroupBaseDn.IsNull() {
		data.LdapGroupBaseDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupBaseDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupFilterPrefix`); value.Exists() && !data.LdapGroupFilterPrefix.IsNull() {
		data.LdapGroupFilterPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupFilterPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupFilterSuffix`); value.Exists() && !data.LdapGroupFilterSuffix.IsNull() {
		data.LdapGroupFilterSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupFilterSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPGroupDynamicFilter`); value.Exists() && !data.LdapGroupDynamicFilter.IsNull() {
		data.LdapGroupDynamicFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapGroupDynamicFilter = types.StringNull()
	}
}
