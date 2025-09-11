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

type User struct {
	Id                types.String                `tfsdk:"id"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	PasswordWo        types.String                `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64                 `tfsdk:"password_wo_version"`
	AccessLevel       types.String                `tfsdk:"access_level"`
	GroupName         types.String                `tfsdk:"group_name"`
	SnmpCreds         types.List                  `tfsdk:"snmp_creds"`
	HashedSnmpCreds   types.List                  `tfsdk:"hashed_snmp_creds"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}
type UserWO struct {
	Id                types.String                `tfsdk:"id"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	AccessLevel       types.String                `tfsdk:"access_level"`
	GroupName         types.String                `tfsdk:"group_name"`
	SnmpCreds         types.List                  `tfsdk:"snmp_creds"`
	HashedSnmpCreds   types.List                  `tfsdk:"hashed_snmp_creds"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var UserGroupNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "access_level",
	AttrType:    "String",
	AttrDefault: "group-defined",
	Value:       []string{"group-defined"},
}
var UserGroupNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var UserHashedSnmpCredsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var UserObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"user_summary":        types.StringType,
	"password_wo":         types.StringType,
	"password_wo_version": types.Int64Type,
	"access_level":        types.StringType,
	"group_name":          types.StringType,
	"snmp_creds":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpCredObjectType}},
	"hashed_snmp_creds":   types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}},
	"dependency_actions":  actions.ActionsListType,
}
var UserObjectTypeWO = map[string]attr.Type{
	"id":                 types.StringType,
	"user_summary":       types.StringType,
	"access_level":       types.StringType,
	"group_name":         types.StringType,
	"snmp_creds":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO}},
	"hashed_snmp_creds":  types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data User) GetPath() string {
	rest_path := "/mgmt/config/default/User"
	return rest_path
}

func (data UserWO) GetPath() string {
	rest_path := "/mgmt/config/default/User"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	return rest_path
}

func (data User) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PasswordWo.IsNull() {
		return false
	}
	if !data.AccessLevel.IsNull() {
		return false
	}
	if !data.GroupName.IsNull() {
		return false
	}
	if !data.SnmpCreds.IsNull() {
		return false
	}
	if !data.HashedSnmpCreds.IsNull() {
		return false
	}
	return true
}
func (data UserWO) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AccessLevel.IsNull() {
		return false
	}
	if !data.GroupName.IsNull() {
		return false
	}
	if !data.SnmpCreds.IsNull() {
		return false
	}
	if !data.HashedSnmpCreds.IsNull() {
		return false
	}
	return true
}

func (data User) ToBody(ctx context.Context, pathRoot string, config *User) string {
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
	if !data.PasswordWo.IsNull() || !data.PasswordWoVersion.IsNull() {
		if data.PasswordWo.IsNull() && config != nil {
			data.PasswordWo = config.PasswordWo
		}
		body, _ = sjson.Set(body, pathRoot+`Password`, data.PasswordWo.ValueString())
	}
	if !data.AccessLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccessLevel`, data.AccessLevel.ValueString())
	}
	if !data.GroupName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GroupName`, data.GroupName.ValueString())
	}
	if !data.SnmpCreds.IsNull() {
		var dataValues []DmSnmpCred
		data.SnmpCreds.ElementsAs(ctx, &dataValues, false)
		var configValues []DmSnmpCred
		config.SnmpCreds.ElementsAs(ctx, &configValues, false)
		for idx, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`SnmpCreds`+".-1", val.ToBody(ctx, "", &configValues[idx]))
		}
	}
	if !data.HashedSnmpCreds.IsNull() {
		var dataValues []DmSnmpCredMasked
		data.HashedSnmpCreds.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`HashedSnmpCreds`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *User) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Password`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccessLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccessLevel = types.StringValue("group-defined")
	}
	if value := res.Get(pathRoot + `GroupName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GroupName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GroupName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SnmpCreds`); value.Exists() {
		l := []DmSnmpCredWO{}
		if value := res.Get(`SnmpCreds`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpCredWO{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO}, l)
		} else {
			data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO})
		}
	} else {
		data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO})
	}
	if value := res.Get(pathRoot + `HashedSnmpCreds`); value.Exists() {
		l := []DmSnmpCredMasked{}
		if value := res.Get(`HashedSnmpCreds`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpCredMasked{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HashedSnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}, l)
		} else {
			data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
		}
	} else {
		data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
	}
}
func (data *UserWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccessLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccessLevel = types.StringValue("group-defined")
	}
	if value := res.Get(pathRoot + `GroupName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GroupName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GroupName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SnmpCreds`); value.Exists() {
		l := []DmSnmpCredWO{}
		if value := res.Get(`SnmpCreds`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpCredWO{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO}, l)
		} else {
			data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO})
		}
	} else {
		data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectTypeWO})
	}
	if value := res.Get(pathRoot + `HashedSnmpCreds`); value.Exists() {
		l := []DmSnmpCredMasked{}
		if value := res.Get(`HashedSnmpCreds`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpCredMasked{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HashedSnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}, l)
		} else {
			data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
		}
	} else {
		data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
	}
}

func (data *User) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Password`); value.Exists() && !data.PasswordWo.IsNull() {
		data.PasswordWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && !data.AccessLevel.IsNull() {
		data.AccessLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.AccessLevel.ValueString() != "group-defined" {
		data.AccessLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `GroupName`); value.Exists() && !data.GroupName.IsNull() {
		data.GroupName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GroupName = types.StringNull()
	}
	if value := res.Get(pathRoot + `SnmpCreds`); value.Exists() && !data.SnmpCreds.IsNull() {
		l := []DmSnmpCred{}
		for _, v := range value.Array() {
			item := DmSnmpCred{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredObjectType}, l)
		} else {
			data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectType})
		}
	} else {
		data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectType})
	}
	if value := res.Get(pathRoot + `HashedSnmpCreds`); value.Exists() && !data.HashedSnmpCreds.IsNull() {
		l := []DmSnmpCredMasked{}
		for _, v := range value.Array() {
			item := DmSnmpCredMasked{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HashedSnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}, l)
		} else {
			data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
		}
	} else {
		data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
	}
}
func (data *User) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Id.IsUnknown() {
		if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
			data.Id = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Id = types.StringNull()
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.PasswordWo.IsUnknown() {
		if value := res.Get(pathRoot + `Password`); value.Exists() && !data.PasswordWo.IsNull() {
			data.PasswordWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.PasswordWo = types.StringNull()
		}
	}
	if data.AccessLevel.IsUnknown() {
		if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && !data.AccessLevel.IsNull() {
			data.AccessLevel = tfutils.ParseStringFromGJSON(value)
		} else if data.AccessLevel.ValueString() != "group-defined" {
			data.AccessLevel = types.StringNull()
		}
	}
	if data.GroupName.IsUnknown() {
		if value := res.Get(pathRoot + `GroupName`); value.Exists() && !data.GroupName.IsNull() {
			data.GroupName = tfutils.ParseStringFromGJSON(value)
		} else {
			data.GroupName = types.StringNull()
		}
	}
	if data.SnmpCreds.IsUnknown() {
		if value := res.Get(pathRoot + `SnmpCreds`); value.Exists() && !data.SnmpCreds.IsNull() {
			l := []DmSnmpCred{}
			if value := res.Get(`SnmpCreds`); value.Exists() {
				for _, v := range value.Array() {
					item := DmSnmpCred{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.SnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredObjectType}, l)
			} else {
				data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectType})
			}
		} else {
			data.SnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredObjectType})
		}
	}
	if data.HashedSnmpCreds.IsUnknown() {
		if value := res.Get(pathRoot + `HashedSnmpCreds`); value.Exists() && !data.HashedSnmpCreds.IsNull() {
			l := []DmSnmpCredMasked{}
			if value := res.Get(`HashedSnmpCreds`); value.Exists() {
				for _, v := range value.Array() {
					item := DmSnmpCredMasked{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.HashedSnmpCreds, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType}, l)
			} else {
				data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
			}
		} else {
			data.HashedSnmpCreds = types.ListNull(types.ObjectType{AttrTypes: DmSnmpCredMaskedObjectType})
		}
	}
}
