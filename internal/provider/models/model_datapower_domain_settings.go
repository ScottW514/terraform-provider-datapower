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
	"path"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DomainSettings struct {
	AppDomain         types.String                `tfsdk:"app_domain"`
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	PasswordTreatment types.String                `tfsdk:"password_treatment"`
	Passphrase        types.String                `tfsdk:"passphrase"`
	PassphraseUpdate  types.Bool                  `tfsdk:"passphrase_update"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}
type DomainSettingsWO struct {
	AppDomain         types.String                `tfsdk:"app_domain"`
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	PasswordTreatment types.String                `tfsdk:"password_treatment"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var DomainSettingsObjectType = map[string]attr.Type{
	"app_domain":         types.StringType,
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"password_treatment": types.StringType,
	"passphrase":         types.StringType,
	"passphrase_update":  types.BoolType,
	"dependency_actions": actions.ActionsListType,
}
var DomainSettingsObjectTypeWO = map[string]attr.Type{
	"app_domain":         types.StringType,
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"password_treatment": types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data DomainSettings) GetPath() string {
	rest_path := "/mgmt/config/{domain}/DomainSettings/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data DomainSettingsWO) GetPath() string {
	rest_path := "/mgmt/config/{domain}/DomainSettings/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data DomainSettings) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PasswordTreatment.IsNull() {
		return false
	}
	if !data.Passphrase.IsNull() {
		return false
	}
	return true
}
func (data DomainSettingsWO) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PasswordTreatment.IsNull() {
		return false
	}
	return true
}

func (data DomainSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "DomainSettings.name", path.Base("/mgmt/config/{domain}/DomainSettings/default"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.PasswordTreatment.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordTreatment`, data.PasswordTreatment.ValueString())
	}
	if !data.Passphrase.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Passphrase`, data.Passphrase.ValueString())
	}
	return body
}

func (data *DomainSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordTreatment`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordTreatment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordTreatment = types.StringValue("masked")
	}
	if value := res.Get(pathRoot + `Passphrase`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Passphrase = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Passphrase = types.StringNull()
	}
}
func (data *DomainSettingsWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordTreatment`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordTreatment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordTreatment = types.StringValue("masked")
	}
}

func (data *DomainSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordTreatment`); value.Exists() && !data.PasswordTreatment.IsNull() {
		data.PasswordTreatment = tfutils.ParseStringFromGJSON(value)
	} else if data.PasswordTreatment.ValueString() != "masked" {
		data.PasswordTreatment = types.StringNull()
	}
	if value := res.Get(pathRoot + `Passphrase`); value.Exists() && !data.Passphrase.IsNull() {
		data.Passphrase = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Passphrase = types.StringNull()
	}
}
func (data *DomainSettings) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Enabled.IsUnknown() {
		if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = tfutils.BoolFromString(value.String())
		} else {
			data.Enabled = types.BoolNull()
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.PasswordTreatment.IsUnknown() {
		if value := res.Get(pathRoot + `PasswordTreatment`); value.Exists() && !data.PasswordTreatment.IsNull() {
			data.PasswordTreatment = tfutils.ParseStringFromGJSON(value)
		} else if data.PasswordTreatment.ValueString() != "masked" {
			data.PasswordTreatment = types.StringNull()
		}
	}
	if data.Passphrase.IsUnknown() {
		if value := res.Get(pathRoot + `Passphrase`); value.Exists() && !data.Passphrase.IsNull() {
			data.Passphrase = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Passphrase = types.StringNull()
		}
	}
}
