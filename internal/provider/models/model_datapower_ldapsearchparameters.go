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

type LDAPSearchParameters struct {
	Id                    types.String `tfsdk:"id"`
	AppDomain             types.String `tfsdk:"app_domain"`
	UserSummary           types.String `tfsdk:"user_summary"`
	LdapBaseDn            types.String `tfsdk:"ldap_base_dn"`
	LdapReturnedAttribute types.String `tfsdk:"ldap_returned_attribute"`
	LdapFilterPrefix      types.String `tfsdk:"ldap_filter_prefix"`
	LdapFilterSuffix      types.String `tfsdk:"ldap_filter_suffix"`
	LdapScope             types.String `tfsdk:"ldap_scope"`
}

var LDAPSearchParametersObjectType = map[string]attr.Type{
	"id":                      types.StringType,
	"app_domain":              types.StringType,
	"user_summary":            types.StringType,
	"ldap_base_dn":            types.StringType,
	"ldap_returned_attribute": types.StringType,
	"ldap_filter_prefix":      types.StringType,
	"ldap_filter_suffix":      types.StringType,
	"ldap_scope":              types.StringType,
}

func (data LDAPSearchParameters) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LDAPSearchParameters"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LDAPSearchParameters) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LdapBaseDn.IsNull() {
		return false
	}
	if !data.LdapReturnedAttribute.IsNull() {
		return false
	}
	if !data.LdapFilterPrefix.IsNull() {
		return false
	}
	if !data.LdapFilterSuffix.IsNull() {
		return false
	}
	if !data.LdapScope.IsNull() {
		return false
	}
	return true
}

func (data LDAPSearchParameters) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LdapBaseDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPBaseDN`, data.LdapBaseDn.ValueString())
	}
	if !data.LdapReturnedAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPReturnedAttribute`, data.LdapReturnedAttribute.ValueString())
	}
	if !data.LdapFilterPrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPFilterPrefix`, data.LdapFilterPrefix.ValueString())
	}
	if !data.LdapFilterSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPFilterSuffix`, data.LdapFilterSuffix.ValueString())
	}
	if !data.LdapScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPScope`, data.LdapScope.ValueString())
	}
	return body
}

func (data *LDAPSearchParameters) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LDAPBaseDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapBaseDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBaseDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPReturnedAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapReturnedAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapReturnedAttribute = types.StringValue("dn")
	}
	if value := res.Get(pathRoot + `LDAPFilterPrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapFilterPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapFilterPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPFilterSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapFilterSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapFilterSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapScope = types.StringValue("subtree")
	}
}

func (data *LDAPSearchParameters) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LDAPBaseDN`); value.Exists() && !data.LdapBaseDn.IsNull() {
		data.LdapBaseDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapBaseDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPReturnedAttribute`); value.Exists() && !data.LdapReturnedAttribute.IsNull() {
		data.LdapReturnedAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapReturnedAttribute.ValueString() != "dn" {
		data.LdapReturnedAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPFilterPrefix`); value.Exists() && !data.LdapFilterPrefix.IsNull() {
		data.LdapFilterPrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapFilterPrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPFilterSuffix`); value.Exists() && !data.LdapFilterSuffix.IsNull() {
		data.LdapFilterSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapFilterSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPScope`); value.Exists() && !data.LdapScope.IsNull() {
		data.LdapScope = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapScope.ValueString() != "subtree" {
		data.LdapScope = types.StringNull()
	}
}
