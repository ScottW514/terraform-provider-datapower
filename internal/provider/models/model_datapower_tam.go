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

type TAM struct {
	Id                          types.String   `tfsdk:"id"`
	AppDomain                   types.String   `tfsdk:"app_domain"`
	UserSummary                 types.String   `tfsdk:"user_summary"`
	AdUseAd                     types.Bool     `tfsdk:"ad_use_ad"`
	TamVersion                  types.String   `tfsdk:"tam_version"`
	ConfigurationFile           types.String   `tfsdk:"configuration_file"`
	AdConfigurationFile         types.String   `tfsdk:"ad_configuration_file"`
	SslKeyFile                  types.String   `tfsdk:"ssl_key_file"`
	SslKeyStashFile             types.String   `tfsdk:"ssl_key_stash_file"`
	UseLocalMode                types.Bool     `tfsdk:"use_local_mode"`
	PollInterval                types.String   `tfsdk:"poll_interval"`
	ListenMode                  types.Bool     `tfsdk:"listen_mode"`
	ListenPort                  types.Int64    `tfsdk:"listen_port"`
	ReturningUserAttributes     types.Bool     `tfsdk:"returning_user_attributes"`
	LdapUseSsl                  types.Bool     `tfsdk:"ldap_use_ssl"`
	LdapsslPort                 types.Int64    `tfsdk:"ldapssl_port"`
	LdapsslKeyFile              types.String   `tfsdk:"ldapssl_key_file"`
	LdapsslKeyFilePassword      types.String   `tfsdk:"ldapssl_key_file_password"`
	LdapsslKeyFilePasswordAlias types.String   `tfsdk:"ldapssl_key_file_password_alias"`
	LdapsslKeyFileLabel         types.String   `tfsdk:"ldapssl_key_file_label"`
	TamUseFips                  types.Bool     `tfsdk:"tam_use_fips"`
	TamChooseNist               types.String   `tfsdk:"tam_choose_nist"`
	TamUseBasicUser             types.Bool     `tfsdk:"tam_use_basic_user"`
	UserPrincipalAttribute      types.String   `tfsdk:"user_principal_attribute"`
	UserNoDuplicates            types.Bool     `tfsdk:"user_no_duplicates"`
	UserSearchSuffixes          types.List     `tfsdk:"user_search_suffixes"`
	UserSuffixOptimiser         types.Bool     `tfsdk:"user_suffix_optimiser"`
	TamFedDirs                  types.List     `tfsdk:"tam_fed_dirs"`
	TamazReplicas               types.List     `tfsdk:"tamaz_replicas"`
	TamrasTrace                 *DmTAMRASTrace `tfsdk:"tamras_trace"`
	AutoRetry                   types.Bool     `tfsdk:"auto_retry"`
	RetryInterval               types.Int64    `tfsdk:"retry_interval"`
	RetryAttempts               types.Int64    `tfsdk:"retry_attempts"`
	LongRetryInterval           types.Int64    `tfsdk:"long_retry_interval"`
}

var TAMObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"ad_use_ad":                       types.BoolType,
	"tam_version":                     types.StringType,
	"configuration_file":              types.StringType,
	"ad_configuration_file":           types.StringType,
	"ssl_key_file":                    types.StringType,
	"ssl_key_stash_file":              types.StringType,
	"use_local_mode":                  types.BoolType,
	"poll_interval":                   types.StringType,
	"listen_mode":                     types.BoolType,
	"listen_port":                     types.Int64Type,
	"returning_user_attributes":       types.BoolType,
	"ldap_use_ssl":                    types.BoolType,
	"ldapssl_port":                    types.Int64Type,
	"ldapssl_key_file":                types.StringType,
	"ldapssl_key_file_password":       types.StringType,
	"ldapssl_key_file_password_alias": types.StringType,
	"ldapssl_key_file_label":          types.StringType,
	"tam_use_fips":                    types.BoolType,
	"tam_choose_nist":                 types.StringType,
	"tam_use_basic_user":              types.BoolType,
	"user_principal_attribute":        types.StringType,
	"user_no_duplicates":              types.BoolType,
	"user_search_suffixes":            types.ListType{ElemType: types.StringType},
	"user_suffix_optimiser":           types.BoolType,
	"tam_fed_dirs":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmTAMFedDirObjectType}},
	"tamaz_replicas":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType}},
	"tamras_trace":                    types.ObjectType{AttrTypes: DmTAMRASTraceObjectType},
	"auto_retry":                      types.BoolType,
	"retry_interval":                  types.Int64Type,
	"retry_attempts":                  types.Int64Type,
	"long_retry_interval":             types.Int64Type,
}

func (data TAM) GetPath() string {
	rest_path := "/mgmt/config/{domain}/TAM"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data TAM) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AdUseAd.IsNull() {
		return false
	}
	if !data.TamVersion.IsNull() {
		return false
	}
	if !data.ConfigurationFile.IsNull() {
		return false
	}
	if !data.AdConfigurationFile.IsNull() {
		return false
	}
	if !data.SslKeyFile.IsNull() {
		return false
	}
	if !data.SslKeyStashFile.IsNull() {
		return false
	}
	if !data.UseLocalMode.IsNull() {
		return false
	}
	if !data.PollInterval.IsNull() {
		return false
	}
	if !data.ListenMode.IsNull() {
		return false
	}
	if !data.ListenPort.IsNull() {
		return false
	}
	if !data.ReturningUserAttributes.IsNull() {
		return false
	}
	if !data.LdapUseSsl.IsNull() {
		return false
	}
	if !data.LdapsslPort.IsNull() {
		return false
	}
	if !data.LdapsslKeyFile.IsNull() {
		return false
	}
	if !data.LdapsslKeyFilePassword.IsNull() {
		return false
	}
	if !data.LdapsslKeyFilePasswordAlias.IsNull() {
		return false
	}
	if !data.LdapsslKeyFileLabel.IsNull() {
		return false
	}
	if !data.TamUseFips.IsNull() {
		return false
	}
	if !data.TamChooseNist.IsNull() {
		return false
	}
	if !data.TamUseBasicUser.IsNull() {
		return false
	}
	if !data.UserPrincipalAttribute.IsNull() {
		return false
	}
	if !data.UserNoDuplicates.IsNull() {
		return false
	}
	if !data.UserSearchSuffixes.IsNull() {
		return false
	}
	if !data.UserSuffixOptimiser.IsNull() {
		return false
	}
	if !data.TamFedDirs.IsNull() {
		return false
	}
	if !data.TamazReplicas.IsNull() {
		return false
	}
	if data.TamrasTrace != nil {
		if !data.TamrasTrace.IsNull() {
			return false
		}
	}
	if !data.AutoRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.RetryAttempts.IsNull() {
		return false
	}
	if !data.LongRetryInterval.IsNull() {
		return false
	}
	return true
}

func (data TAM) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AdUseAd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ADUseAD`, tfutils.StringFromBool(data.AdUseAd))
	}
	if !data.TamVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMVersion`, data.TamVersion.ValueString())
	}
	if !data.ConfigurationFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigurationFile`, data.ConfigurationFile.ValueString())
	}
	if !data.AdConfigurationFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ADConfigurationFile`, data.AdConfigurationFile.ValueString())
	}
	if !data.SslKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLKeyFile`, data.SslKeyFile.ValueString())
	}
	if !data.SslKeyStashFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLKeyStashFile`, data.SslKeyStashFile.ValueString())
	}
	if !data.UseLocalMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseLocalMode`, tfutils.StringFromBool(data.UseLocalMode))
	}
	if !data.PollInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PollInterval`, data.PollInterval.ValueString())
	}
	if !data.ListenMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ListenMode`, tfutils.StringFromBool(data.ListenMode))
	}
	if !data.ListenPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ListenPort`, data.ListenPort.ValueInt64())
	}
	if !data.ReturningUserAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReturningUserAttributes`, tfutils.StringFromBool(data.ReturningUserAttributes))
	}
	if !data.LdapUseSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPUseSSL`, tfutils.StringFromBool(data.LdapUseSsl))
	}
	if !data.LdapsslPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLPort`, data.LdapsslPort.ValueInt64())
	}
	if !data.LdapsslKeyFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLKeyFile`, data.LdapsslKeyFile.ValueString())
	}
	if !data.LdapsslKeyFilePassword.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLKeyFilePassword`, data.LdapsslKeyFilePassword.ValueString())
	}
	if !data.LdapsslKeyFilePasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLKeyFilePasswordAlias`, data.LdapsslKeyFilePasswordAlias.ValueString())
	}
	if !data.LdapsslKeyFileLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPSSLKeyFileLabel`, data.LdapsslKeyFileLabel.ValueString())
	}
	if !data.TamUseFips.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMUseFIPS`, tfutils.StringFromBool(data.TamUseFips))
	}
	if !data.TamChooseNist.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMChooseNIST`, data.TamChooseNist.ValueString())
	}
	if !data.TamUseBasicUser.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMUseBasicUser`, tfutils.StringFromBool(data.TamUseBasicUser))
	}
	if !data.UserPrincipalAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserPrincipalAttribute`, data.UserPrincipalAttribute.ValueString())
	}
	if !data.UserNoDuplicates.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserNoDuplicates`, tfutils.StringFromBool(data.UserNoDuplicates))
	}
	if !data.UserSearchSuffixes.IsNull() {
		var values []string
		data.UserSearchSuffixes.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`UserSearchSuffixes`+".-1", map[string]string{"value": val})
		}
	}
	if !data.UserSuffixOptimiser.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSuffixOptimiser`, tfutils.StringFromBool(data.UserSuffixOptimiser))
	}
	if !data.TamFedDirs.IsNull() {
		var values []DmTAMFedDir
		data.TamFedDirs.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`TAMFedDirs`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.TamazReplicas.IsNull() {
		var values []DmTAMAZReplica
		data.TamazReplicas.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`TAMAZReplicas`+".-1", val.ToBody(ctx, ""))
		}
	}
	if data.TamrasTrace != nil {
		if !data.TamrasTrace.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`TAMRASTrace`, data.TamrasTrace.ToBody(ctx, ""))
		}
	}
	if !data.AutoRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRetry`, tfutils.StringFromBool(data.AutoRetry))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.RetryAttempts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryAttempts`, data.RetryAttempts.ValueInt64())
	}
	if !data.LongRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LongRetryInterval`, data.LongRetryInterval.ValueInt64())
	}
	return body
}

func (data *TAM) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ADUseAD`); value.Exists() {
		data.AdUseAd = tfutils.BoolFromString(value.String())
	} else {
		data.AdUseAd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigurationFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigurationFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigurationFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ADConfigurationFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdConfigurationFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdConfigurationFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLKeyStashFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslKeyStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKeyStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseLocalMode`); value.Exists() {
		data.UseLocalMode = tfutils.BoolFromString(value.String())
	} else {
		data.UseLocalMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PollInterval`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PollInterval = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PollInterval = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `ListenMode`); value.Exists() {
		data.ListenMode = tfutils.BoolFromString(value.String())
	} else {
		data.ListenMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ListenPort`); value.Exists() {
		data.ListenPort = types.Int64Value(value.Int())
	} else {
		data.ListenPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReturningUserAttributes`); value.Exists() {
		data.ReturningUserAttributes = tfutils.BoolFromString(value.String())
	} else {
		data.ReturningUserAttributes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPUseSSL`); value.Exists() {
		data.LdapUseSsl = tfutils.BoolFromString(value.String())
	} else {
		data.LdapUseSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLPort`); value.Exists() {
		data.LdapsslPort = types.Int64Value(value.Int())
	} else {
		data.LdapsslPort = types.Int64Value(636)
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFilePassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFilePasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFileLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapsslKeyFileLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFileLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMUseFIPS`); value.Exists() {
		data.TamUseFips = tfutils.BoolFromString(value.String())
	} else {
		data.TamUseFips = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMChooseNIST`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamChooseNist = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamChooseNist = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMUseBasicUser`); value.Exists() {
		data.TamUseBasicUser = tfutils.BoolFromString(value.String())
	} else {
		data.TamUseBasicUser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserPrincipalAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserPrincipalAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserPrincipalAttribute = types.StringValue("uid")
	}
	if value := res.Get(pathRoot + `UserNoDuplicates`); value.Exists() {
		data.UserNoDuplicates = tfutils.BoolFromString(value.String())
	} else {
		data.UserNoDuplicates = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSearchSuffixes`); value.Exists() {
		data.UserSearchSuffixes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.UserSearchSuffixes = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `UserSuffixOptimiser`); value.Exists() {
		data.UserSuffixOptimiser = tfutils.BoolFromString(value.String())
	} else {
		data.UserSuffixOptimiser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMFedDirs`); value.Exists() {
		l := []DmTAMFedDir{}
		if value := res.Get(`TAMFedDirs`); value.Exists() {
			for _, v := range value.Array() {
				item := DmTAMFedDir{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TamFedDirs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTAMFedDirObjectType}, l)
		} else {
			data.TamFedDirs = types.ListNull(types.ObjectType{AttrTypes: DmTAMFedDirObjectType})
		}
	} else {
		data.TamFedDirs = types.ListNull(types.ObjectType{AttrTypes: DmTAMFedDirObjectType})
	}
	if value := res.Get(pathRoot + `TAMAZReplicas`); value.Exists() {
		l := []DmTAMAZReplica{}
		if value := res.Get(`TAMAZReplicas`); value.Exists() {
			for _, v := range value.Array() {
				item := DmTAMAZReplica{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TamazReplicas, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType}, l)
		} else {
			data.TamazReplicas = types.ListNull(types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType})
		}
	} else {
		data.TamazReplicas = types.ListNull(types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType})
	}
	if value := res.Get(pathRoot + `TAMRASTrace`); value.Exists() {
		data.TamrasTrace = &DmTAMRASTrace{}
		data.TamrasTrace.FromBody(ctx, "", value)
	} else {
		data.TamrasTrace = nil
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else {
		data.RetryAttempts = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else {
		data.LongRetryInterval = types.Int64Value(900)
	}
}

func (data *TAM) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ADUseAD`); value.Exists() && !data.AdUseAd.IsNull() {
		data.AdUseAd = tfutils.BoolFromString(value.String())
	} else if data.AdUseAd.ValueBool() {
		data.AdUseAd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMVersion`); value.Exists() && !data.TamVersion.IsNull() {
		data.TamVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigurationFile`); value.Exists() && !data.ConfigurationFile.IsNull() {
		data.ConfigurationFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigurationFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ADConfigurationFile`); value.Exists() && !data.AdConfigurationFile.IsNull() {
		data.AdConfigurationFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdConfigurationFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLKeyFile`); value.Exists() && !data.SslKeyFile.IsNull() {
		data.SslKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLKeyStashFile`); value.Exists() && !data.SslKeyStashFile.IsNull() {
		data.SslKeyStashFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslKeyStashFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseLocalMode`); value.Exists() && !data.UseLocalMode.IsNull() {
		data.UseLocalMode = tfutils.BoolFromString(value.String())
	} else if data.UseLocalMode.ValueBool() {
		data.UseLocalMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PollInterval`); value.Exists() && !data.PollInterval.IsNull() {
		data.PollInterval = tfutils.ParseStringFromGJSON(value)
	} else if data.PollInterval.ValueString() != "default" {
		data.PollInterval = types.StringNull()
	}
	if value := res.Get(pathRoot + `ListenMode`); value.Exists() && !data.ListenMode.IsNull() {
		data.ListenMode = tfutils.BoolFromString(value.String())
	} else if data.ListenMode.ValueBool() {
		data.ListenMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ListenPort`); value.Exists() && !data.ListenPort.IsNull() {
		data.ListenPort = types.Int64Value(value.Int())
	} else {
		data.ListenPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ReturningUserAttributes`); value.Exists() && !data.ReturningUserAttributes.IsNull() {
		data.ReturningUserAttributes = tfutils.BoolFromString(value.String())
	} else if data.ReturningUserAttributes.ValueBool() {
		data.ReturningUserAttributes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPUseSSL`); value.Exists() && !data.LdapUseSsl.IsNull() {
		data.LdapUseSsl = tfutils.BoolFromString(value.String())
	} else if data.LdapUseSsl.ValueBool() {
		data.LdapUseSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLPort`); value.Exists() && !data.LdapsslPort.IsNull() {
		data.LdapsslPort = types.Int64Value(value.Int())
	} else if data.LdapsslPort.ValueInt64() != 636 {
		data.LdapsslPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFile`); value.Exists() && !data.LdapsslKeyFile.IsNull() {
		data.LdapsslKeyFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFilePassword`); value.Exists() && !data.LdapsslKeyFilePassword.IsNull() {
		data.LdapsslKeyFilePassword = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFilePassword = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFilePasswordAlias`); value.Exists() && !data.LdapsslKeyFilePasswordAlias.IsNull() {
		data.LdapsslKeyFilePasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFilePasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPSSLKeyFileLabel`); value.Exists() && !data.LdapsslKeyFileLabel.IsNull() {
		data.LdapsslKeyFileLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapsslKeyFileLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMUseFIPS`); value.Exists() && !data.TamUseFips.IsNull() {
		data.TamUseFips = tfutils.BoolFromString(value.String())
	} else if data.TamUseFips.ValueBool() {
		data.TamUseFips = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMChooseNIST`); value.Exists() && !data.TamChooseNist.IsNull() {
		data.TamChooseNist = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamChooseNist = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMUseBasicUser`); value.Exists() && !data.TamUseBasicUser.IsNull() {
		data.TamUseBasicUser = tfutils.BoolFromString(value.String())
	} else if data.TamUseBasicUser.ValueBool() {
		data.TamUseBasicUser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserPrincipalAttribute`); value.Exists() && !data.UserPrincipalAttribute.IsNull() {
		data.UserPrincipalAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.UserPrincipalAttribute.ValueString() != "uid" {
		data.UserPrincipalAttribute = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserNoDuplicates`); value.Exists() && !data.UserNoDuplicates.IsNull() {
		data.UserNoDuplicates = tfutils.BoolFromString(value.String())
	} else if !data.UserNoDuplicates.ValueBool() {
		data.UserNoDuplicates = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSearchSuffixes`); value.Exists() && !data.UserSearchSuffixes.IsNull() {
		data.UserSearchSuffixes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.UserSearchSuffixes = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `UserSuffixOptimiser`); value.Exists() && !data.UserSuffixOptimiser.IsNull() {
		data.UserSuffixOptimiser = tfutils.BoolFromString(value.String())
	} else if data.UserSuffixOptimiser.ValueBool() {
		data.UserSuffixOptimiser = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMFedDirs`); value.Exists() && !data.TamFedDirs.IsNull() {
		l := []DmTAMFedDir{}
		for _, v := range value.Array() {
			item := DmTAMFedDir{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.TamFedDirs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTAMFedDirObjectType}, l)
		} else {
			data.TamFedDirs = types.ListNull(types.ObjectType{AttrTypes: DmTAMFedDirObjectType})
		}
	} else {
		data.TamFedDirs = types.ListNull(types.ObjectType{AttrTypes: DmTAMFedDirObjectType})
	}
	if value := res.Get(pathRoot + `TAMAZReplicas`); value.Exists() && !data.TamazReplicas.IsNull() {
		l := []DmTAMAZReplica{}
		for _, v := range value.Array() {
			item := DmTAMAZReplica{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.TamazReplicas, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType}, l)
		} else {
			data.TamazReplicas = types.ListNull(types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType})
		}
	} else {
		data.TamazReplicas = types.ListNull(types.ObjectType{AttrTypes: DmTAMAZReplicaObjectType})
	}
	if value := res.Get(pathRoot + `TAMRASTrace`); value.Exists() {
		data.TamrasTrace.UpdateFromBody(ctx, "", value)
	} else {
		data.TamrasTrace = nil
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() && !data.AutoRetry.IsNull() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else if data.AutoRetry.ValueBool() {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 180 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() && !data.RetryAttempts.IsNull() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else if data.RetryAttempts.ValueInt64() != 3 {
		data.RetryAttempts = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() && !data.LongRetryInterval.IsNull() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else if data.LongRetryInterval.ValueInt64() != 900 {
		data.LongRetryInterval = types.Int64Null()
	}
}
