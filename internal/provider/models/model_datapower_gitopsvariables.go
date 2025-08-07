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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type GitOpsVariables struct {
	Enabled           types.Bool                  `tfsdk:"enabled"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Variables         types.List                  `tfsdk:"variables"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var GitOpsVariablesObjectType = map[string]attr.Type{
	"enabled":            types.BoolType,
	"user_summary":       types.StringType,
	"variables":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data GitOpsVariables) GetPath() string {
	rest_path := "/mgmt/config/default/GitOpsVariables/gitops-variables"
	return rest_path
}

func (data GitOpsVariables) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Variables.IsNull() {
		return false
	}
	return true
}

func (data GitOpsVariables) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "GitOpsVariables.name", path.Base("/mgmt/config/default/GitOpsVariables/gitops-variables"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Variables.IsNull() {
		var values []DmGitOpsVariableEntry
		data.Variables.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Variables`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *GitOpsVariables) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Variables`); value.Exists() {
		l := []DmGitOpsVariableEntry{}
		if value := res.Get(`Variables`); value.Exists() {
			for _, v := range value.Array() {
				item := DmGitOpsVariableEntry{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Variables, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType}, l)
		} else {
			data.Variables = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType})
		}
	} else {
		data.Variables = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType})
	}
}

func (data *GitOpsVariables) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Variables`); value.Exists() && !data.Variables.IsNull() {
		l := []DmGitOpsVariableEntry{}
		for _, v := range value.Array() {
			item := DmGitOpsVariableEntry{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Variables, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType}, l)
		} else {
			data.Variables = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType})
		}
	} else {
		data.Variables = types.ListNull(types.ObjectType{AttrTypes: DmGitOpsVariableEntryObjectType})
	}
}
