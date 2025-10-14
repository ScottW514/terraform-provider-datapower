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

type SSHClientProfile struct {
	Id                          types.String                    `tfsdk:"id"`
	AppDomain                   types.String                    `tfsdk:"app_domain"`
	UserSummary                 types.String                    `tfsdk:"user_summary"`
	UserName                    types.String                    `tfsdk:"user_name"`
	ProfileUsage                types.String                    `tfsdk:"profile_usage"`
	SshUserAuthentication       *DmSSHUserAuthenticationMethods `tfsdk:"ssh_user_authentication"`
	UserPrivateKey              types.String                    `tfsdk:"user_private_key"`
	PasswordAlias               types.String                    `tfsdk:"password_alias"`
	PersistentConnections       types.Bool                      `tfsdk:"persistent_connections"`
	PersistentConnectionTimeout types.Int64                     `tfsdk:"persistent_connection_timeout"`
	StrictHostKeyChecking       types.Bool                      `tfsdk:"strict_host_key_checking"`
	Ciphers                     types.List                      `tfsdk:"ciphers"`
	KexAlg                      types.List                      `tfsdk:"kex_alg"`
	MacAlg                      types.List                      `tfsdk:"mac_alg"`
	DependencyActions           []*actions.DependencyAction     `tfsdk:"dependency_actions"`
}

var SSHClientProfileSSHUserAuthenticationCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfileSSHUserAuthenticationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfileUserPrivateKeyCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "profile_usage",
					AttrType:    "String",
					AttrDefault: "sftp",
					Value:       []string{"sftp"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "ssh_user_authentication",
					AttrType:    "DmSSHUserAuthenticationMethods",
					AttrDefault: "publickey+password",
					Value:       []string{"publickey"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"scc"},
		},
	},
}

var SSHClientProfileUserPrivateKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "profile_usage",
					AttrType:    "String",
					AttrDefault: "sftp",
					Value:       []string{"sftp"},
				},

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "ssh_user_authentication",
					AttrType:    "DmSSHUserAuthenticationMethods",
					AttrDefault: "publickey+password",
					Value:       []string{"publickey"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"sftp"},
		},
	},
}

var SSHClientProfilePasswordAliasCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"sftp"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssh_user_authentication",
			AttrType:    "DmSSHUserAuthenticationMethods",
			AttrDefault: "publickey+password",
			Value:       []string{"password"},
		},
	},
}

var SSHClientProfilePasswordAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "profile_usage",
					AttrType:    "String",
					AttrDefault: "sftp",
					Value:       []string{"sftp"},
				},

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "ssh_user_authentication",
					AttrType:    "DmSSHUserAuthenticationMethods",
					AttrDefault: "publickey+password",
					Value:       []string{"password"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"sftp"},
		},
	},
}

var SSHClientProfilePersistentConnectionsCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfilePersistentConnectionsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfilePersistentConnectionTimeoutCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"sftp"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "persistent_connections",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var SSHClientProfilePersistentConnectionTimeoutIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "profile_usage",
			AttrType:    "String",
			AttrDefault: "sftp",
			Value:       []string{"sftp"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "persistent_connections",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var SSHClientProfileStrictHostKeyCheckingCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfileStrictHostKeyCheckingIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"sftp"},
}

var SSHClientProfileCiphersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"scc"},
}

var SSHClientProfileKEXAlgIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"scc"},
}

var SSHClientProfileMACAlgIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "profile_usage",
	AttrType:    "String",
	AttrDefault: "sftp",
	Value:       []string{"scc"},
}

var SSHClientProfileObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"user_name":                     types.StringType,
	"profile_usage":                 types.StringType,
	"ssh_user_authentication":       types.ObjectType{AttrTypes: DmSSHUserAuthenticationMethodsObjectType},
	"user_private_key":              types.StringType,
	"password_alias":                types.StringType,
	"persistent_connections":        types.BoolType,
	"persistent_connection_timeout": types.Int64Type,
	"strict_host_key_checking":      types.BoolType,
	"ciphers":                       types.ListType{ElemType: types.StringType},
	"kex_alg":                       types.ListType{ElemType: types.StringType},
	"mac_alg":                       types.ListType{ElemType: types.StringType},
	"dependency_actions":            actions.ActionsListType,
}

func (data SSHClientProfile) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSHClientProfile"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSHClientProfile) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.ProfileUsage.IsNull() {
		return false
	}
	if data.SshUserAuthentication != nil {
		if !data.SshUserAuthentication.IsNull() {
			return false
		}
	}
	if !data.UserPrivateKey.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.PersistentConnections.IsNull() {
		return false
	}
	if !data.PersistentConnectionTimeout.IsNull() {
		return false
	}
	if !data.StrictHostKeyChecking.IsNull() {
		return false
	}
	if !data.Ciphers.IsNull() {
		return false
	}
	if !data.KexAlg.IsNull() {
		return false
	}
	if !data.MacAlg.IsNull() {
		return false
	}
	return true
}

func (data SSHClientProfile) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.ProfileUsage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProfileUsage`, data.ProfileUsage.ValueString())
	}
	if data.SshUserAuthentication != nil {
		if !data.SshUserAuthentication.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSHUserAuthentication`, data.SshUserAuthentication.ToBody(ctx, ""))
		}
	}
	if !data.UserPrivateKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserPrivateKey`, data.UserPrivateKey.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.PersistentConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnections`, tfutils.StringFromBool(data.PersistentConnections, ""))
	}
	if !data.PersistentConnectionTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnectionTimeout`, data.PersistentConnectionTimeout.ValueInt64())
	}
	if !data.StrictHostKeyChecking.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StrictHostKeyChecking`, tfutils.StringFromBool(data.StrictHostKeyChecking, ""))
	}
	if !data.Ciphers.IsNull() {
		var dataValues []string
		data.Ciphers.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`Ciphers`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Ciphers`, "[]")
	}
	if !data.KexAlg.IsNull() {
		var dataValues []string
		data.KexAlg.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`KEXAlg`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`KEXAlg`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`KEXAlg`, "[]")
	}
	if !data.MacAlg.IsNull() {
		var dataValues []string
		data.MacAlg.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`MACAlg`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`MACAlg`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`MACAlg`, "[]")
	}
	return body
}

func (data *SSHClientProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileUsage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProfileUsage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProfileUsage = types.StringValue("sftp")
	}
	if value := res.Get(pathRoot + `SSHUserAuthentication`); value.Exists() {
		data.SshUserAuthentication = &DmSSHUserAuthenticationMethods{}
		data.SshUserAuthentication.FromBody(ctx, "", value)
	} else {
		data.SshUserAuthentication = nil
	}
	if value := res.Get(pathRoot + `UserPrivateKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserPrivateKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserPrivateKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnectionTimeout`); value.Exists() {
		data.PersistentConnectionTimeout = types.Int64Value(value.Int())
	} else {
		data.PersistentConnectionTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `StrictHostKeyChecking`); value.Exists() {
		data.StrictHostKeyChecking = tfutils.BoolFromString(value.String())
	} else {
		data.StrictHostKeyChecking = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `KEXAlg`); value.Exists() {
		data.KexAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.KexAlg = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MACAlg`); value.Exists() {
		data.MacAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MacAlg = types.ListNull(types.StringType)
	}
}

func (data *SSHClientProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProfileUsage`); value.Exists() && !data.ProfileUsage.IsNull() {
		data.ProfileUsage = tfutils.ParseStringFromGJSON(value)
	} else if data.ProfileUsage.ValueString() != "sftp" {
		data.ProfileUsage = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHUserAuthentication`); value.Exists() {
		data.SshUserAuthentication.UpdateFromBody(ctx, "", value)
	} else {
		data.SshUserAuthentication = nil
	}
	if value := res.Get(pathRoot + `UserPrivateKey`); value.Exists() && !data.UserPrivateKey.IsNull() {
		data.UserPrivateKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserPrivateKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentConnections`); value.Exists() && !data.PersistentConnections.IsNull() {
		data.PersistentConnections = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnections.ValueBool() {
		data.PersistentConnections = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PersistentConnectionTimeout`); value.Exists() && !data.PersistentConnectionTimeout.IsNull() {
		data.PersistentConnectionTimeout = types.Int64Value(value.Int())
	} else if data.PersistentConnectionTimeout.ValueInt64() != 120 {
		data.PersistentConnectionTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `StrictHostKeyChecking`); value.Exists() && !data.StrictHostKeyChecking.IsNull() {
		data.StrictHostKeyChecking = tfutils.BoolFromString(value.String())
	} else if data.StrictHostKeyChecking.ValueBool() {
		data.StrictHostKeyChecking = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Ciphers`); value.Exists() && !data.Ciphers.IsNull() {
		data.Ciphers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Ciphers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `KEXAlg`); value.Exists() && !data.KexAlg.IsNull() {
		data.KexAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.KexAlg = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MACAlg`); value.Exists() && !data.MacAlg.IsNull() {
		data.MacAlg = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MacAlg = types.ListNull(types.StringType)
	}
}
