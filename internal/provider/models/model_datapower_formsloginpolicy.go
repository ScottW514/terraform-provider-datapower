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

type FormsLoginPolicy struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	LoginForm           types.String                `tfsdk:"login_form"`
	UseCookieAttributes types.Bool                  `tfsdk:"use_cookie_attributes"`
	CookieAttributes    types.String                `tfsdk:"cookie_attributes"`
	UseSslForLogin      types.Bool                  `tfsdk:"use_ssl_for_login"`
	EnableMigration     types.Bool                  `tfsdk:"enable_migration"`
	SslPort             types.Int64                 `tfsdk:"ssl_port"`
	SharedSecret        types.String                `tfsdk:"shared_secret"`
	ErrorPage           types.String                `tfsdk:"error_page"`
	LogoutPage          types.String                `tfsdk:"logout_page"`
	DefaultUrl          types.String                `tfsdk:"default_url"`
	FormsLocation       types.String                `tfsdk:"forms_location"`
	LocalLoginForm      types.String                `tfsdk:"local_login_form"`
	LocalErrorPage      types.String                `tfsdk:"local_error_page"`
	LocalLogoutPage     types.String                `tfsdk:"local_logout_page"`
	UsernameField       types.String                `tfsdk:"username_field"`
	PasswordField       types.String                `tfsdk:"password_field"`
	RedirectField       types.String                `tfsdk:"redirect_field"`
	FormProcessingUrl   types.String                `tfsdk:"form_processing_url"`
	InactivityTimeout   types.Int64                 `tfsdk:"inactivity_timeout"`
	SessionLifetime     types.Int64                 `tfsdk:"session_lifetime"`
	RedirectUrlType     types.String                `tfsdk:"redirect_url_type"`
	FormSupportType     types.String                `tfsdk:"form_support_type"`
	FormSupportScript   types.String                `tfsdk:"form_support_script"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var FormsLoginPolicyObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"login_form":            types.StringType,
	"use_cookie_attributes": types.BoolType,
	"cookie_attributes":     types.StringType,
	"use_ssl_for_login":     types.BoolType,
	"enable_migration":      types.BoolType,
	"ssl_port":              types.Int64Type,
	"shared_secret":         types.StringType,
	"error_page":            types.StringType,
	"logout_page":           types.StringType,
	"default_url":           types.StringType,
	"forms_location":        types.StringType,
	"local_login_form":      types.StringType,
	"local_error_page":      types.StringType,
	"local_logout_page":     types.StringType,
	"username_field":        types.StringType,
	"password_field":        types.StringType,
	"redirect_field":        types.StringType,
	"form_processing_url":   types.StringType,
	"inactivity_timeout":    types.Int64Type,
	"session_lifetime":      types.Int64Type,
	"redirect_url_type":     types.StringType,
	"form_support_type":     types.StringType,
	"form_support_script":   types.StringType,
	"dependency_actions":    actions.ActionsListType,
}

func (data FormsLoginPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/FormsLoginPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data FormsLoginPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LoginForm.IsNull() {
		return false
	}
	if !data.UseCookieAttributes.IsNull() {
		return false
	}
	if !data.CookieAttributes.IsNull() {
		return false
	}
	if !data.UseSslForLogin.IsNull() {
		return false
	}
	if !data.EnableMigration.IsNull() {
		return false
	}
	if !data.SslPort.IsNull() {
		return false
	}
	if !data.SharedSecret.IsNull() {
		return false
	}
	if !data.ErrorPage.IsNull() {
		return false
	}
	if !data.LogoutPage.IsNull() {
		return false
	}
	if !data.DefaultUrl.IsNull() {
		return false
	}
	if !data.FormsLocation.IsNull() {
		return false
	}
	if !data.LocalLoginForm.IsNull() {
		return false
	}
	if !data.LocalErrorPage.IsNull() {
		return false
	}
	if !data.LocalLogoutPage.IsNull() {
		return false
	}
	if !data.UsernameField.IsNull() {
		return false
	}
	if !data.PasswordField.IsNull() {
		return false
	}
	if !data.RedirectField.IsNull() {
		return false
	}
	if !data.FormProcessingUrl.IsNull() {
		return false
	}
	if !data.InactivityTimeout.IsNull() {
		return false
	}
	if !data.SessionLifetime.IsNull() {
		return false
	}
	if !data.RedirectUrlType.IsNull() {
		return false
	}
	if !data.FormSupportType.IsNull() {
		return false
	}
	if !data.FormSupportScript.IsNull() {
		return false
	}
	return true
}

func (data FormsLoginPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LoginForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoginForm`, data.LoginForm.ValueString())
	}
	if !data.UseCookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCookieAttributes`, tfutils.StringFromBool(data.UseCookieAttributes, "flag"))
	}
	if !data.CookieAttributes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CookieAttributes`, data.CookieAttributes.ValueString())
	}
	if !data.UseSslForLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseSSLForLogin`, tfutils.StringFromBool(data.UseSslForLogin, ""))
	}
	if !data.EnableMigration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableMigration`, tfutils.StringFromBool(data.EnableMigration, ""))
	}
	if !data.SslPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLPort`, data.SslPort.ValueInt64())
	}
	if !data.SharedSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SharedSecret`, data.SharedSecret.ValueString())
	}
	if !data.ErrorPage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorPage`, data.ErrorPage.ValueString())
	}
	if !data.LogoutPage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogoutPage`, data.LogoutPage.ValueString())
	}
	if !data.DefaultUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultURL`, data.DefaultUrl.ValueString())
	}
	if !data.FormsLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FormsLocation`, data.FormsLocation.ValueString())
	}
	if !data.LocalLoginForm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalLoginForm`, data.LocalLoginForm.ValueString())
	}
	if !data.LocalErrorPage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalErrorPage`, data.LocalErrorPage.ValueString())
	}
	if !data.LocalLogoutPage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalLogoutPage`, data.LocalLogoutPage.ValueString())
	}
	if !data.UsernameField.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UsernameField`, data.UsernameField.ValueString())
	}
	if !data.PasswordField.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordField`, data.PasswordField.ValueString())
	}
	if !data.RedirectField.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RedirectField`, data.RedirectField.ValueString())
	}
	if !data.FormProcessingUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FormProcessingURL`, data.FormProcessingUrl.ValueString())
	}
	if !data.InactivityTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InactivityTimeout`, data.InactivityTimeout.ValueInt64())
	}
	if !data.SessionLifetime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SessionLifetime`, data.SessionLifetime.ValueInt64())
	}
	if !data.RedirectUrlType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RedirectURLType`, data.RedirectUrlType.ValueString())
	}
	if !data.FormSupportType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FormSupportType`, data.FormSupportType.ValueString())
	}
	if !data.FormSupportScript.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FormSupportScript`, data.FormSupportScript.ValueString())
	}
	return body
}

func (data *FormsLoginPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LoginForm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LoginForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoginForm = types.StringValue("/LoginPage.htm")
	}
	if value := res.Get(pathRoot + `UseCookieAttributes`); value.Exists() {
		data.UseCookieAttributes = tfutils.BoolFromString(value.String())
	} else {
		data.UseCookieAttributes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieAttributes`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSSLForLogin`); value.Exists() {
		data.UseSslForLogin = tfutils.BoolFromString(value.String())
	} else {
		data.UseSslForLogin = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableMigration`); value.Exists() {
		data.EnableMigration = tfutils.BoolFromString(value.String())
	} else {
		data.EnableMigration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLPort`); value.Exists() {
		data.SslPort = types.Int64Value(value.Int())
	} else {
		data.SslPort = types.Int64Value(8080)
	}
	if value := res.Get(pathRoot + `SharedSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SharedSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SharedSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPage = types.StringValue("/ErrorPage.htm")
	}
	if value := res.Get(pathRoot + `LogoutPage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogoutPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogoutPage = types.StringValue("/LogoutPage.htm")
	}
	if value := res.Get(pathRoot + `DefaultURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultUrl = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `FormsLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FormsLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FormsLocation = types.StringValue("backend")
	}
	if value := res.Get(pathRoot + `LocalLoginForm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalLoginForm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalLoginForm = types.StringValue("store:///LoginPage.htm")
	}
	if value := res.Get(pathRoot + `LocalErrorPage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalErrorPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalErrorPage = types.StringValue("store:///ErrorPage.htm")
	}
	if value := res.Get(pathRoot + `LocalLogoutPage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalLogoutPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalLogoutPage = types.StringValue("store:///LogoutPage.htm")
	}
	if value := res.Get(pathRoot + `UsernameField`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UsernameField = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UsernameField = types.StringValue("j_username")
	}
	if value := res.Get(pathRoot + `PasswordField`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordField = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordField = types.StringValue("j_password")
	}
	if value := res.Get(pathRoot + `RedirectField`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RedirectField = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RedirectField = types.StringValue("originalUrl")
	}
	if value := res.Get(pathRoot + `FormProcessingURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FormProcessingUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FormProcessingUrl = types.StringValue("/j_security_check")
	}
	if value := res.Get(pathRoot + `InactivityTimeout`); value.Exists() {
		data.InactivityTimeout = types.Int64Value(value.Int())
	} else {
		data.InactivityTimeout = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `SessionLifetime`); value.Exists() {
		data.SessionLifetime = types.Int64Value(value.Int())
	} else {
		data.SessionLifetime = types.Int64Value(10800)
	}
	if value := res.Get(pathRoot + `RedirectURLType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RedirectUrlType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RedirectUrlType = types.StringValue("urlin")
	}
	if value := res.Get(pathRoot + `FormSupportType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FormSupportType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FormSupportType = types.StringValue("static")
	}
	if value := res.Get(pathRoot + `FormSupportScript`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FormSupportScript = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FormSupportScript = types.StringValue("store:///Form-Generate-HTML.xsl")
	}
}

func (data *FormsLoginPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LoginForm`); value.Exists() && !data.LoginForm.IsNull() {
		data.LoginForm = tfutils.ParseStringFromGJSON(value)
	} else if data.LoginForm.ValueString() != "/LoginPage.htm" {
		data.LoginForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseCookieAttributes`); value.Exists() && !data.UseCookieAttributes.IsNull() {
		data.UseCookieAttributes = tfutils.BoolFromString(value.String())
	} else if data.UseCookieAttributes.ValueBool() {
		data.UseCookieAttributes = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieAttributes`); value.Exists() && !data.CookieAttributes.IsNull() {
		data.CookieAttributes = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CookieAttributes = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSSLForLogin`); value.Exists() && !data.UseSslForLogin.IsNull() {
		data.UseSslForLogin = tfutils.BoolFromString(value.String())
	} else if !data.UseSslForLogin.ValueBool() {
		data.UseSslForLogin = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableMigration`); value.Exists() && !data.EnableMigration.IsNull() {
		data.EnableMigration = tfutils.BoolFromString(value.String())
	} else if data.EnableMigration.ValueBool() {
		data.EnableMigration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLPort`); value.Exists() && !data.SslPort.IsNull() {
		data.SslPort = types.Int64Value(value.Int())
	} else if data.SslPort.ValueInt64() != 8080 {
		data.SslPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SharedSecret`); value.Exists() && !data.SharedSecret.IsNull() {
		data.SharedSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SharedSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPage`); value.Exists() && !data.ErrorPage.IsNull() {
		data.ErrorPage = tfutils.ParseStringFromGJSON(value)
	} else if data.ErrorPage.ValueString() != "/ErrorPage.htm" {
		data.ErrorPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogoutPage`); value.Exists() && !data.LogoutPage.IsNull() {
		data.LogoutPage = tfutils.ParseStringFromGJSON(value)
	} else if data.LogoutPage.ValueString() != "/LogoutPage.htm" {
		data.LogoutPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultURL`); value.Exists() && !data.DefaultUrl.IsNull() {
		data.DefaultUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultUrl.ValueString() != "/" {
		data.DefaultUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `FormsLocation`); value.Exists() && !data.FormsLocation.IsNull() {
		data.FormsLocation = tfutils.ParseStringFromGJSON(value)
	} else if data.FormsLocation.ValueString() != "backend" {
		data.FormsLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalLoginForm`); value.Exists() && !data.LocalLoginForm.IsNull() {
		data.LocalLoginForm = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalLoginForm.ValueString() != "store:///LoginPage.htm" {
		data.LocalLoginForm = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalErrorPage`); value.Exists() && !data.LocalErrorPage.IsNull() {
		data.LocalErrorPage = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalErrorPage.ValueString() != "store:///ErrorPage.htm" {
		data.LocalErrorPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalLogoutPage`); value.Exists() && !data.LocalLogoutPage.IsNull() {
		data.LocalLogoutPage = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalLogoutPage.ValueString() != "store:///LogoutPage.htm" {
		data.LocalLogoutPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `UsernameField`); value.Exists() && !data.UsernameField.IsNull() {
		data.UsernameField = tfutils.ParseStringFromGJSON(value)
	} else if data.UsernameField.ValueString() != "j_username" {
		data.UsernameField = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordField`); value.Exists() && !data.PasswordField.IsNull() {
		data.PasswordField = tfutils.ParseStringFromGJSON(value)
	} else if data.PasswordField.ValueString() != "j_password" {
		data.PasswordField = types.StringNull()
	}
	if value := res.Get(pathRoot + `RedirectField`); value.Exists() && !data.RedirectField.IsNull() {
		data.RedirectField = tfutils.ParseStringFromGJSON(value)
	} else if data.RedirectField.ValueString() != "originalUrl" {
		data.RedirectField = types.StringNull()
	}
	if value := res.Get(pathRoot + `FormProcessingURL`); value.Exists() && !data.FormProcessingUrl.IsNull() {
		data.FormProcessingUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.FormProcessingUrl.ValueString() != "/j_security_check" {
		data.FormProcessingUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `InactivityTimeout`); value.Exists() && !data.InactivityTimeout.IsNull() {
		data.InactivityTimeout = types.Int64Value(value.Int())
	} else if data.InactivityTimeout.ValueInt64() != 600 {
		data.InactivityTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SessionLifetime`); value.Exists() && !data.SessionLifetime.IsNull() {
		data.SessionLifetime = types.Int64Value(value.Int())
	} else if data.SessionLifetime.ValueInt64() != 10800 {
		data.SessionLifetime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RedirectURLType`); value.Exists() && !data.RedirectUrlType.IsNull() {
		data.RedirectUrlType = tfutils.ParseStringFromGJSON(value)
	} else if data.RedirectUrlType.ValueString() != "urlin" {
		data.RedirectUrlType = types.StringNull()
	}
	if value := res.Get(pathRoot + `FormSupportType`); value.Exists() && !data.FormSupportType.IsNull() {
		data.FormSupportType = tfutils.ParseStringFromGJSON(value)
	} else if data.FormSupportType.ValueString() != "static" {
		data.FormSupportType = types.StringNull()
	}
	if value := res.Get(pathRoot + `FormSupportScript`); value.Exists() && !data.FormSupportScript.IsNull() {
		data.FormSupportScript = tfutils.ParseStringFromGJSON(value)
	} else if data.FormSupportScript.ValueString() != "store:///Form-Generate-HTML.xsl" {
		data.FormSupportScript = types.StringNull()
	}
}
