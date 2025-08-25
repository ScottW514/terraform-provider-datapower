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

type SNMPSettings struct {
	Enabled                        types.Bool                  `tfsdk:"enabled"`
	UserSummary                    types.String                `tfsdk:"user_summary"`
	LocalAddress                   types.String                `tfsdk:"local_address"`
	LocalPort                      types.Int64                 `tfsdk:"local_port"`
	Policies                       types.List                  `tfsdk:"policies"`
	PoliciesMq                     types.List                  `tfsdk:"policies_mq"`
	Targets                        types.List                  `tfsdk:"targets"`
	Users                          types.List                  `tfsdk:"users"`
	Contexts                       types.List                  `tfsdk:"contexts"`
	SecurityLevel                  types.String                `tfsdk:"security_level"`
	AccessLevel                    types.String                `tfsdk:"access_level"`
	EnableDefaultTrapSubscriptions types.Bool                  `tfsdk:"enable_default_trap_subscriptions"`
	TrapPriority                   types.String                `tfsdk:"trap_priority"`
	TrapEventCode                  types.List                  `tfsdk:"trap_event_code"`
	ConfigMib                      types.String                `tfsdk:"config_mib"`
	ConfigMibMq                    types.String                `tfsdk:"config_mib_mq"`
	StatusMib                      types.String                `tfsdk:"status_mib"`
	StatusMibMq                    types.String                `tfsdk:"status_mib_mq"`
	NotifMib                       types.String                `tfsdk:"notif_mib"`
	NotifMibMq                     types.String                `tfsdk:"notif_mib_mq"`
	DependencyActions              []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SNMPSettingsObjectType = map[string]attr.Type{
	"enabled":                           types.BoolType,
	"user_summary":                      types.StringType,
	"local_address":                     types.StringType,
	"local_port":                        types.Int64Type,
	"policies":                          types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpPolicyObjectType}},
	"policies_mq":                       types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType}},
	"targets":                           types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpTargetObjectType}},
	"users":                             types.ListType{ElemType: types.StringType},
	"contexts":                          types.ListType{ElemType: types.ObjectType{AttrTypes: DmSnmpContextObjectType}},
	"security_level":                    types.StringType,
	"access_level":                      types.StringType,
	"enable_default_trap_subscriptions": types.BoolType,
	"trap_priority":                     types.StringType,
	"trap_event_code":                   types.ListType{ElemType: types.StringType},
	"config_mib":                        types.StringType,
	"config_mib_mq":                     types.StringType,
	"status_mib":                        types.StringType,
	"status_mib_mq":                     types.StringType,
	"notif_mib":                         types.StringType,
	"notif_mib_mq":                      types.StringType,
	"dependency_actions":                actions.ActionsListType,
}

func (data SNMPSettings) GetPath() string {
	rest_path := "/mgmt/config/default/SNMPSettings/SNMP-Settings"
	return rest_path
}

func (data SNMPSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.Policies.IsNull() {
		return false
	}
	if !data.PoliciesMq.IsNull() {
		return false
	}
	if !data.Targets.IsNull() {
		return false
	}
	if !data.Users.IsNull() {
		return false
	}
	if !data.Contexts.IsNull() {
		return false
	}
	if !data.SecurityLevel.IsNull() {
		return false
	}
	if !data.AccessLevel.IsNull() {
		return false
	}
	if !data.EnableDefaultTrapSubscriptions.IsNull() {
		return false
	}
	if !data.TrapPriority.IsNull() {
		return false
	}
	if !data.TrapEventCode.IsNull() {
		return false
	}
	if !data.ConfigMib.IsNull() {
		return false
	}
	if !data.ConfigMibMq.IsNull() {
		return false
	}
	if !data.StatusMib.IsNull() {
		return false
	}
	if !data.StatusMibMq.IsNull() {
		return false
	}
	if !data.NotifMib.IsNull() {
		return false
	}
	if !data.NotifMibMq.IsNull() {
		return false
	}
	return true
}

func (data SNMPSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "SNMPSettings.name", path.Base("/mgmt/config/default/SNMPSettings/SNMP-Settings"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.Policies.IsNull() {
		var dataValues []DmSnmpPolicy
		data.Policies.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Policies`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.PoliciesMq.IsNull() {
		var dataValues []DmSnmpPolicyMQ
		data.PoliciesMq.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`PoliciesMQ`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Targets.IsNull() {
		var dataValues []DmSnmpTarget
		data.Targets.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Targets`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Users.IsNull() {
		var dataValues []string
		data.Users.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Users`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Contexts.IsNull() {
		var dataValues []DmSnmpContext
		data.Contexts.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Contexts`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SecurityLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SecurityLevel`, data.SecurityLevel.ValueString())
	}
	if !data.AccessLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccessLevel`, data.AccessLevel.ValueString())
	}
	if !data.EnableDefaultTrapSubscriptions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableDefaultTrapSubscriptions`, tfutils.StringFromBool(data.EnableDefaultTrapSubscriptions, ""))
	}
	if !data.TrapPriority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TrapPriority`, data.TrapPriority.ValueString())
	}
	if !data.TrapEventCode.IsNull() {
		var dataValues []string
		data.TrapEventCode.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`TrapEventCode`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ConfigMib.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigMib`, data.ConfigMib.ValueString())
	}
	if !data.ConfigMibMq.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigMibMQ`, data.ConfigMibMq.ValueString())
	}
	if !data.StatusMib.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StatusMib`, data.StatusMib.ValueString())
	}
	if !data.StatusMibMq.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StatusMibMQ`, data.StatusMibMq.ValueString())
	}
	if !data.NotifMib.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NotifMib`, data.NotifMib.ValueString())
	}
	if !data.NotifMibMq.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NotifMibMQ`, data.NotifMibMq.ValueString())
	}
	return body
}

func (data *SNMPSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(161)
	}
	if value := res.Get(pathRoot + `Policies`); value.Exists() {
		l := []DmSnmpPolicy{}
		if value := res.Get(`Policies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Policies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpPolicyObjectType}, l)
		} else {
			data.Policies = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyObjectType})
		}
	} else {
		data.Policies = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyObjectType})
	}
	if value := res.Get(pathRoot + `PoliciesMQ`); value.Exists() {
		l := []DmSnmpPolicyMQ{}
		if value := res.Get(`PoliciesMQ`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpPolicyMQ{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PoliciesMq, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType}, l)
		} else {
			data.PoliciesMq = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType})
		}
	} else {
		data.PoliciesMq = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType})
	}
	if value := res.Get(pathRoot + `Targets`); value.Exists() {
		l := []DmSnmpTarget{}
		if value := res.Get(`Targets`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpTarget{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Targets, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpTargetObjectType}, l)
		} else {
			data.Targets = types.ListNull(types.ObjectType{AttrTypes: DmSnmpTargetObjectType})
		}
	} else {
		data.Targets = types.ListNull(types.ObjectType{AttrTypes: DmSnmpTargetObjectType})
	}
	if value := res.Get(pathRoot + `Users`); value.Exists() {
		data.Users = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Users = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Contexts`); value.Exists() {
		l := []DmSnmpContext{}
		if value := res.Get(`Contexts`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSnmpContext{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Contexts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpContextObjectType}, l)
		} else {
			data.Contexts = types.ListNull(types.ObjectType{AttrTypes: DmSnmpContextObjectType})
		}
	} else {
		data.Contexts = types.ListNull(types.ObjectType{AttrTypes: DmSnmpContextObjectType})
	}
	if value := res.Get(pathRoot + `SecurityLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecurityLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecurityLevel = types.StringValue("authPriv")
	}
	if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccessLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccessLevel = types.StringValue("read-only")
	}
	if value := res.Get(pathRoot + `EnableDefaultTrapSubscriptions`); value.Exists() {
		data.EnableDefaultTrapSubscriptions = tfutils.BoolFromString(value.String())
	} else {
		data.EnableDefaultTrapSubscriptions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TrapPriority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TrapPriority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TrapPriority = types.StringValue("warn")
	}
	if value := res.Get(pathRoot + `TrapEventCode`); value.Exists() {
		data.TrapEventCode = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.TrapEventCode = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ConfigMib`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigMib = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigMib = types.StringValue("/drConfigMIB.txt")
	}
	if value := res.Get(pathRoot + `ConfigMibMQ`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigMibMq = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigMibMq = types.StringValue("/mqConfigMIB.txt")
	}
	if value := res.Get(pathRoot + `StatusMib`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StatusMib = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StatusMib = types.StringValue("/drStatusMIB.txt")
	}
	if value := res.Get(pathRoot + `StatusMibMQ`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StatusMibMq = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StatusMibMq = types.StringValue("/mqStatusMIB.txt")
	}
	if value := res.Get(pathRoot + `NotifMib`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NotifMib = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NotifMib = types.StringValue("/drNotificationMIB.txt")
	}
	if value := res.Get(pathRoot + `NotifMibMQ`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NotifMibMq = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NotifMibMq = types.StringValue("/mqNotificationMIB.txt")
	}
}

func (data *SNMPSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 161 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Policies`); value.Exists() && !data.Policies.IsNull() {
		l := []DmSnmpPolicy{}
		for _, v := range value.Array() {
			item := DmSnmpPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Policies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpPolicyObjectType}, l)
		} else {
			data.Policies = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyObjectType})
		}
	} else {
		data.Policies = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyObjectType})
	}
	if value := res.Get(pathRoot + `PoliciesMQ`); value.Exists() && !data.PoliciesMq.IsNull() {
		l := []DmSnmpPolicyMQ{}
		for _, v := range value.Array() {
			item := DmSnmpPolicyMQ{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.PoliciesMq, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType}, l)
		} else {
			data.PoliciesMq = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType})
		}
	} else {
		data.PoliciesMq = types.ListNull(types.ObjectType{AttrTypes: DmSnmpPolicyMQObjectType})
	}
	if value := res.Get(pathRoot + `Targets`); value.Exists() && !data.Targets.IsNull() {
		l := []DmSnmpTarget{}
		for _, v := range value.Array() {
			item := DmSnmpTarget{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Targets, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpTargetObjectType}, l)
		} else {
			data.Targets = types.ListNull(types.ObjectType{AttrTypes: DmSnmpTargetObjectType})
		}
	} else {
		data.Targets = types.ListNull(types.ObjectType{AttrTypes: DmSnmpTargetObjectType})
	}
	if value := res.Get(pathRoot + `Users`); value.Exists() && !data.Users.IsNull() {
		data.Users = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Users = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Contexts`); value.Exists() && !data.Contexts.IsNull() {
		l := []DmSnmpContext{}
		for _, v := range value.Array() {
			item := DmSnmpContext{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Contexts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSnmpContextObjectType}, l)
		} else {
			data.Contexts = types.ListNull(types.ObjectType{AttrTypes: DmSnmpContextObjectType})
		}
	} else {
		data.Contexts = types.ListNull(types.ObjectType{AttrTypes: DmSnmpContextObjectType})
	}
	if value := res.Get(pathRoot + `SecurityLevel`); value.Exists() && !data.SecurityLevel.IsNull() {
		data.SecurityLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.SecurityLevel.ValueString() != "authPriv" {
		data.SecurityLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessLevel`); value.Exists() && !data.AccessLevel.IsNull() {
		data.AccessLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.AccessLevel.ValueString() != "read-only" {
		data.AccessLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnableDefaultTrapSubscriptions`); value.Exists() && !data.EnableDefaultTrapSubscriptions.IsNull() {
		data.EnableDefaultTrapSubscriptions = tfutils.BoolFromString(value.String())
	} else if !data.EnableDefaultTrapSubscriptions.ValueBool() {
		data.EnableDefaultTrapSubscriptions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TrapPriority`); value.Exists() && !data.TrapPriority.IsNull() {
		data.TrapPriority = tfutils.ParseStringFromGJSON(value)
	} else if data.TrapPriority.ValueString() != "warn" {
		data.TrapPriority = types.StringNull()
	}
	if value := res.Get(pathRoot + `TrapEventCode`); value.Exists() && !data.TrapEventCode.IsNull() {
		data.TrapEventCode = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.TrapEventCode = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ConfigMib`); value.Exists() && !data.ConfigMib.IsNull() {
		data.ConfigMib = tfutils.ParseStringFromGJSON(value)
	} else if data.ConfigMib.ValueString() != "/drConfigMIB.txt" {
		data.ConfigMib = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigMibMQ`); value.Exists() && !data.ConfigMibMq.IsNull() {
		data.ConfigMibMq = tfutils.ParseStringFromGJSON(value)
	} else if data.ConfigMibMq.ValueString() != "/mqConfigMIB.txt" {
		data.ConfigMibMq = types.StringNull()
	}
	if value := res.Get(pathRoot + `StatusMib`); value.Exists() && !data.StatusMib.IsNull() {
		data.StatusMib = tfutils.ParseStringFromGJSON(value)
	} else if data.StatusMib.ValueString() != "/drStatusMIB.txt" {
		data.StatusMib = types.StringNull()
	}
	if value := res.Get(pathRoot + `StatusMibMQ`); value.Exists() && !data.StatusMibMq.IsNull() {
		data.StatusMibMq = tfutils.ParseStringFromGJSON(value)
	} else if data.StatusMibMq.ValueString() != "/mqStatusMIB.txt" {
		data.StatusMibMq = types.StringNull()
	}
	if value := res.Get(pathRoot + `NotifMib`); value.Exists() && !data.NotifMib.IsNull() {
		data.NotifMib = tfutils.ParseStringFromGJSON(value)
	} else if data.NotifMib.ValueString() != "/drNotificationMIB.txt" {
		data.NotifMib = types.StringNull()
	}
	if value := res.Get(pathRoot + `NotifMibMQ`); value.Exists() && !data.NotifMibMq.IsNull() {
		data.NotifMibMq = tfutils.ParseStringFromGJSON(value)
	} else if data.NotifMibMq.ValueString() != "/mqNotificationMIB.txt" {
		data.NotifMibMq = types.StringNull()
	}
}
