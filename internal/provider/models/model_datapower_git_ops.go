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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type GitOps struct {
	AppDomain             types.String                `tfsdk:"app_domain"`
	Enabled               types.Bool                  `tfsdk:"enabled"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	ConnectionType        types.String                `tfsdk:"connection_type"`
	Mode                  types.String                `tfsdk:"mode"`
	CommitIdentifierType  types.String                `tfsdk:"commit_identifier_type"`
	CommitIdentifier      types.String                `tfsdk:"commit_identifier"`
	RemoteLocation        types.String                `tfsdk:"remote_location"`
	Interval              types.Int64                 `tfsdk:"interval"`
	SshClientProfile      types.String                `tfsdk:"ssh_client_profile"`
	Username              types.String                `tfsdk:"username"`
	Password              types.String                `tfsdk:"password"`
	SshAuthorizedKeysFile types.String                `tfsdk:"ssh_authorized_keys_file"`
	TlsValcred            types.String                `tfsdk:"tls_valcred"`
	GitUser               types.String                `tfsdk:"git_user"`
	GitEmail              types.String                `tfsdk:"git_email"`
	JsonParseSettings     types.String                `tfsdk:"json_parse_settings"`
	TemplatePolicies      types.List                  `tfsdk:"template_policies"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var GitOpsIntervalCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "read-write",
	Value:       []string{"read-only"},
}
var GitOpsSSHClientProfileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "connection_type",
	AttrType:    "String",
	AttrDefault: "https",
	Value:       []string{"ssh"},
}
var GitOpsUsernameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "connection_type",
	AttrType:    "String",
	AttrDefault: "https",
	Value:       []string{"https"},
}
var GitOpsPasswordCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "connection_type",
	AttrType:    "String",
	AttrDefault: "https",
	Value:       []string{"https"},
}
var GitOpsSSHAuthorizedKeysFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "connection_type",
	AttrType:    "String",
	AttrDefault: "https",
	Value:       []string{"ssh"},
}
var GitOpsTLSValcredCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "connection_type",
	AttrType:    "String",
	AttrDefault: "https",
	Value:       []string{"https"},
}
var GitOpsGitUserCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "read-write",
	Value:       []string{"read-write"},
}
var GitOpsGitEmailCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "read-write",
	Value:       []string{"read-write"},
}

var GitOpsObjectType = map[string]attr.Type{
	"app_domain":               types.StringType,
	"enabled":                  types.BoolType,
	"user_summary":             types.StringType,
	"connection_type":          types.StringType,
	"mode":                     types.StringType,
	"commit_identifier_type":   types.StringType,
	"commit_identifier":        types.StringType,
	"remote_location":          types.StringType,
	"interval":                 types.Int64Type,
	"ssh_client_profile":       types.StringType,
	"username":                 types.StringType,
	"password":                 types.StringType,
	"ssh_authorized_keys_file": types.StringType,
	"tls_valcred":              types.StringType,
	"git_user":                 types.StringType,
	"git_email":                types.StringType,
	"json_parse_settings":      types.StringType,
	"template_policies":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType}},
	"dependency_actions":       actions.ActionsListType,
}

func (data GitOps) GetPath() string {
	rest_path := "/mgmt/config/{domain}/GitOps/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data GitOps) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ConnectionType.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.CommitIdentifierType.IsNull() {
		return false
	}
	if !data.CommitIdentifier.IsNull() {
		return false
	}
	if !data.RemoteLocation.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	if !data.SshClientProfile.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	if !data.SshAuthorizedKeysFile.IsNull() {
		return false
	}
	if !data.TlsValcred.IsNull() {
		return false
	}
	if !data.GitUser.IsNull() {
		return false
	}
	if !data.GitEmail.IsNull() {
		return false
	}
	if !data.JsonParseSettings.IsNull() {
		return false
	}
	if !data.TemplatePolicies.IsNull() {
		return false
	}
	return true
}

func (data GitOps) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "GitOps.name", path.Base("/mgmt/config/{domain}/GitOps/default"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ConnectionType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectionType`, data.ConnectionType.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Mode`, data.Mode.ValueString())
	}
	if !data.CommitIdentifierType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CommitIdentifierType`, data.CommitIdentifierType.ValueString())
	}
	if !data.CommitIdentifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CommitIdentifier`, data.CommitIdentifier.ValueString())
	}
	if !data.RemoteLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteLocation`, data.RemoteLocation.ValueString())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	if !data.SshClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSHClientProfile`, data.SshClientProfile.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Username`, data.Username.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Password`, data.Password.ValueString())
	}
	if !data.SshAuthorizedKeysFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSHAuthorizedKeysFile`, data.SshAuthorizedKeysFile.ValueString())
	}
	if !data.TlsValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TLSValcred`, data.TlsValcred.ValueString())
	}
	if !data.GitUser.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GitUser`, data.GitUser.ValueString())
	}
	if !data.GitEmail.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GitEmail`, data.GitEmail.ValueString())
	}
	if !data.JsonParseSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONParseSettings`, data.JsonParseSettings.ValueString())
	}
	if !data.TemplatePolicies.IsNull() {
		var dataValues []DmGitOpsTemplatePolicy
		data.TemplatePolicies.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`TemplatePolicies`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *GitOps) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ConnectionType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConnectionType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConnectionType = types.StringValue("https")
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringValue("read-write")
	}
	if value := res.Get(pathRoot + `CommitIdentifierType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CommitIdentifierType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CommitIdentifierType = types.StringValue("branch")
	}
	if value := res.Get(pathRoot + `CommitIdentifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CommitIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CommitIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Value(5)
	}
	if value := res.Get(pathRoot + `SSHClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SshClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHAuthorizedKeysFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SshAuthorizedKeysFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshAuthorizedKeysFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `TLSValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TlsValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TlsValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GitUser`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GitUser = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GitUser = types.StringNull()
	}
	if value := res.Get(pathRoot + `GitEmail`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GitEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GitEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `JSONParseSettings`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JsonParseSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonParseSettings = types.StringNull()
	}
	if value := res.Get(pathRoot + `TemplatePolicies`); value.Exists() {
		l := []DmGitOpsTemplatePolicy{}
		if value := res.Get(`TemplatePolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmGitOpsTemplatePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TemplatePolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType}, l)
		} else {
			data.TemplatePolicies = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType})
		}
	} else {
		data.TemplatePolicies = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType})
	}
}

func (data *GitOps) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConnectionType`); value.Exists() && !data.ConnectionType.IsNull() {
		data.ConnectionType = tfutils.ParseStringFromGJSON(value)
	} else if data.ConnectionType.ValueString() != "https" {
		data.ConnectionType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && !data.Mode.IsNull() {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else if data.Mode.ValueString() != "read-write" {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `CommitIdentifierType`); value.Exists() && !data.CommitIdentifierType.IsNull() {
		data.CommitIdentifierType = tfutils.ParseStringFromGJSON(value)
	} else if data.CommitIdentifierType.ValueString() != "branch" {
		data.CommitIdentifierType = types.StringNull()
	}
	if value := res.Get(pathRoot + `CommitIdentifier`); value.Exists() && !data.CommitIdentifier.IsNull() {
		data.CommitIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CommitIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteLocation`); value.Exists() && !data.RemoteLocation.IsNull() {
		data.RemoteLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else if data.Interval.ValueInt64() != 5 {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSHClientProfile`); value.Exists() && !data.SshClientProfile.IsNull() {
		data.SshClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && !data.Username.IsNull() {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHAuthorizedKeysFile`); value.Exists() && !data.SshAuthorizedKeysFile.IsNull() {
		data.SshAuthorizedKeysFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshAuthorizedKeysFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `TLSValcred`); value.Exists() && !data.TlsValcred.IsNull() {
		data.TlsValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TlsValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GitUser`); value.Exists() && !data.GitUser.IsNull() {
		data.GitUser = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GitUser = types.StringNull()
	}
	if value := res.Get(pathRoot + `GitEmail`); value.Exists() && !data.GitEmail.IsNull() {
		data.GitEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GitEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `JSONParseSettings`); value.Exists() && !data.JsonParseSettings.IsNull() {
		data.JsonParseSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonParseSettings = types.StringNull()
	}
	if value := res.Get(pathRoot + `TemplatePolicies`); value.Exists() && !data.TemplatePolicies.IsNull() {
		l := []DmGitOpsTemplatePolicy{}
		for _, v := range value.Array() {
			item := DmGitOpsTemplatePolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.TemplatePolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType}, l)
		} else {
			data.TemplatePolicies = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType})
		}
	} else {
		data.TemplatePolicies = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsTemplatePolicyObjectType})
	}
}
