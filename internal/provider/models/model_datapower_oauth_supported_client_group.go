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

type OAuthSupportedClientGroup struct {
	Id                 types.String                `tfsdk:"id"`
	AppDomain          types.String                `tfsdk:"app_domain"`
	UserSummary        types.String                `tfsdk:"user_summary"`
	Customized         types.Bool                  `tfsdk:"customized"`
	CustomizedType     types.String                `tfsdk:"customized_type"`
	OauthRole          *DmOAuthRole                `tfsdk:"oauth_role"`
	Client             types.List                  `tfsdk:"client"`
	TemplateProcessUrl types.String                `tfsdk:"template_process_url"`
	ClientTemplate     types.String                `tfsdk:"client_template"`
	DependencyActions  []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var OAuthSupportedClientGroupOAuthRoleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "customized",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var OAuthSupportedClientGroupClientCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized_type",
					AttrType:    "String",
					AttrDefault: "custom",
					Value:       []string{"custom"},
				},
			},
		},
	},
}
var OAuthSupportedClientGroupTemplateProcessUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized_type",
			AttrType:    "String",
			AttrDefault: "custom",
			Value:       []string{"template"},
		},
	},
}
var OAuthSupportedClientGroupClientTemplateCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized_type",
			AttrType:    "String",
			AttrDefault: "custom",
			Value:       []string{"template"},
		},
	},
}
var OAuthSupportedClientGroupCustomizedTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "customized",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var OAuthSupportedClientGroupOAuthRoleIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthSupportedClientGroupClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthSupportedClientGroupTemplateProcessUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var OAuthSupportedClientGroupClientTemplateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientGroupObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"customized":           types.BoolType,
	"customized_type":      types.StringType,
	"oauth_role":           types.ObjectType{AttrTypes: DmOAuthRoleObjectType},
	"client":               types.ListType{ElemType: types.StringType},
	"template_process_url": types.StringType,
	"client_template":      types.StringType,
	"dependency_actions":   actions.ActionsListType,
}

func (data OAuthSupportedClientGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OAuthSupportedClientGroup"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OAuthSupportedClientGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Customized.IsNull() {
		return false
	}
	if !data.CustomizedType.IsNull() {
		return false
	}
	if data.OauthRole != nil {
		if !data.OauthRole.IsNull() {
			return false
		}
	}
	if !data.Client.IsNull() {
		return false
	}
	if !data.TemplateProcessUrl.IsNull() {
		return false
	}
	if !data.ClientTemplate.IsNull() {
		return false
	}
	return true
}

func (data OAuthSupportedClientGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Customized.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Customized`, tfutils.StringFromBool(data.Customized, ""))
	}
	if !data.CustomizedType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomizedType`, data.CustomizedType.ValueString())
	}
	if data.OauthRole != nil {
		if !data.OauthRole.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OAuthRole`, data.OauthRole.ToBody(ctx, ""))
		}
	}
	if !data.Client.IsNull() {
		var dataValues []string
		data.Client.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Client`+".-1", map[string]string{"value": val})
		}
	}
	if !data.TemplateProcessUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TemplateProcessUrl`, data.TemplateProcessUrl.ValueString())
	}
	if !data.ClientTemplate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientTemplate`, data.ClientTemplate.ValueString())
	}
	return body
}

func (data *OAuthSupportedClientGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomizedType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedType = types.StringValue("custom")
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OauthRole = &DmOAuthRole{}
		data.OauthRole.FromBody(ctx, "", value)
	} else {
		data.OauthRole = nil
	}
	if value := res.Get(pathRoot + `Client`); value.Exists() {
		data.Client = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Client = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `TemplateProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TemplateProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TemplateProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientTemplate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientTemplate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientTemplate = types.StringNull()
	}
}

func (data *OAuthSupportedClientGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() && !data.Customized.IsNull() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else if data.Customized.ValueBool() {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedType`); value.Exists() && !data.CustomizedType.IsNull() {
		data.CustomizedType = tfutils.ParseStringFromGJSON(value)
	} else if data.CustomizedType.ValueString() != "custom" {
		data.CustomizedType = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OauthRole.UpdateFromBody(ctx, "", value)
	} else {
		data.OauthRole = nil
	}
	if value := res.Get(pathRoot + `Client`); value.Exists() && !data.Client.IsNull() {
		data.Client = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Client = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `TemplateProcessUrl`); value.Exists() && !data.TemplateProcessUrl.IsNull() {
		data.TemplateProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TemplateProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientTemplate`); value.Exists() && !data.ClientTemplate.IsNull() {
		data.ClientTemplate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientTemplate = types.StringNull()
	}
}
