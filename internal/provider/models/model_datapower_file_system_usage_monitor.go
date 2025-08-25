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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FileSystemUsageMonitor struct {
	Enabled                    types.Bool                  `tfsdk:"enabled"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	PollingInterval            types.Int64                 `tfsdk:"polling_interval"`
	AllSystem                  types.Bool                  `tfsdk:"all_system"`
	AllSystemWarningThreshold  types.Int64                 `tfsdk:"all_system_warning_threshold"`
	AllSystemCriticalThreshold types.Int64                 `tfsdk:"all_system_critical_threshold"`
	System                     types.List                  `tfsdk:"system"`
	AllQueueManagers           types.Bool                  `tfsdk:"all_queue_managers"`
	AllQmWarningThreshold      types.Int64                 `tfsdk:"all_qm_warning_threshold"`
	AllQmCriticalThreshold     types.Int64                 `tfsdk:"all_qm_critical_threshold"`
	QueueManager               types.List                  `tfsdk:"queue_manager"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var FileSystemUsageMonitorAllSystemWarningThresholdCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "all_system",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var FileSystemUsageMonitorAllSystemCriticalThresholdCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "all_system",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var FileSystemUsageMonitorAllQMWarningThresholdCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "all_queue_managers",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var FileSystemUsageMonitorAllQMCriticalThresholdCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "all_queue_managers",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var FileSystemUsageMonitorObjectType = map[string]attr.Type{
	"enabled":                       types.BoolType,
	"user_summary":                  types.StringType,
	"polling_interval":              types.Int64Type,
	"all_system":                    types.BoolType,
	"all_system_warning_threshold":  types.Int64Type,
	"all_system_critical_threshold": types.Int64Type,
	"system":                        types.ListType{ElemType: types.ObjectType{AttrTypes: DmFileSystemUsageObjectType}},
	"all_queue_managers":            types.BoolType,
	"all_qm_warning_threshold":      types.Int64Type,
	"all_qm_critical_threshold":     types.Int64Type,
	"queue_manager":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType}},
	"dependency_actions":            actions.ActionsListType,
}

func (data FileSystemUsageMonitor) GetPath() string {
	rest_path := "/mgmt/config/default/FileSystemUsageMonitor/FileSystemUsageMonitor"
	return rest_path
}

func (data FileSystemUsageMonitor) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PollingInterval.IsNull() {
		return false
	}
	if !data.AllSystem.IsNull() {
		return false
	}
	if !data.AllSystemWarningThreshold.IsNull() {
		return false
	}
	if !data.AllSystemCriticalThreshold.IsNull() {
		return false
	}
	if !data.System.IsNull() {
		return false
	}
	if !data.AllQueueManagers.IsNull() {
		return false
	}
	if !data.AllQmWarningThreshold.IsNull() {
		return false
	}
	if !data.AllQmCriticalThreshold.IsNull() {
		return false
	}
	if !data.QueueManager.IsNull() {
		return false
	}
	return true
}

func (data FileSystemUsageMonitor) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "FileSystemUsageMonitor.name", path.Base("/mgmt/config/default/FileSystemUsageMonitor/FileSystemUsageMonitor"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.PollingInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PollingInterval`, data.PollingInterval.ValueInt64())
	}
	if !data.AllSystem.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllSystem`, tfutils.StringFromBool(data.AllSystem, ""))
	}
	if !data.AllSystemWarningThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllSystemWarningThreshold`, data.AllSystemWarningThreshold.ValueInt64())
	}
	if !data.AllSystemCriticalThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllSystemCriticalThreshold`, data.AllSystemCriticalThreshold.ValueInt64())
	}
	if !data.System.IsNull() {
		var dataValues []DmFileSystemUsage
		data.System.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`System`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AllQueueManagers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllQueueManagers`, tfutils.StringFromBool(data.AllQueueManagers, ""))
	}
	if !data.AllQmWarningThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllQMWarningThreshold`, data.AllQmWarningThreshold.ValueInt64())
	}
	if !data.AllQmCriticalThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllQMCriticalThreshold`, data.AllQmCriticalThreshold.ValueInt64())
	}
	if !data.QueueManager.IsNull() {
		var dataValues []DmQMFileSystemUsage
		data.QueueManager.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`QueueManager`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *FileSystemUsageMonitor) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else {
		data.PollingInterval = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `AllSystem`); value.Exists() {
		data.AllSystem = tfutils.BoolFromString(value.String())
	} else {
		data.AllSystem = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllSystemWarningThreshold`); value.Exists() {
		data.AllSystemWarningThreshold = types.Int64Value(value.Int())
	} else {
		data.AllSystemWarningThreshold = types.Int64Value(75)
	}
	if value := res.Get(pathRoot + `AllSystemCriticalThreshold`); value.Exists() {
		data.AllSystemCriticalThreshold = types.Int64Value(value.Int())
	} else {
		data.AllSystemCriticalThreshold = types.Int64Value(90)
	}
	if value := res.Get(pathRoot + `System`); value.Exists() {
		l := []DmFileSystemUsage{}
		if value := res.Get(`System`); value.Exists() {
			for _, v := range value.Array() {
				item := DmFileSystemUsage{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.System, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFileSystemUsageObjectType}, l)
		} else {
			data.System = types.ListNull(types.ObjectType{AttrTypes: DmFileSystemUsageObjectType})
		}
	} else {
		data.System = types.ListNull(types.ObjectType{AttrTypes: DmFileSystemUsageObjectType})
	}
	if value := res.Get(pathRoot + `AllQueueManagers`); value.Exists() {
		data.AllQueueManagers = tfutils.BoolFromString(value.String())
	} else {
		data.AllQueueManagers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllQMWarningThreshold`); value.Exists() {
		data.AllQmWarningThreshold = types.Int64Value(value.Int())
	} else {
		data.AllQmWarningThreshold = types.Int64Value(75)
	}
	if value := res.Get(pathRoot + `AllQMCriticalThreshold`); value.Exists() {
		data.AllQmCriticalThreshold = types.Int64Value(value.Int())
	} else {
		data.AllQmCriticalThreshold = types.Int64Value(90)
	}
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() {
		l := []DmQMFileSystemUsage{}
		if value := res.Get(`QueueManager`); value.Exists() {
			for _, v := range value.Array() {
				item := DmQMFileSystemUsage{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.QueueManager, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType}, l)
		} else {
			data.QueueManager = types.ListNull(types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType})
		}
	} else {
		data.QueueManager = types.ListNull(types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType})
	}
}

func (data *FileSystemUsageMonitor) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PollingInterval`); value.Exists() && !data.PollingInterval.IsNull() {
		data.PollingInterval = types.Int64Value(value.Int())
	} else if data.PollingInterval.ValueInt64() != 60 {
		data.PollingInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AllSystem`); value.Exists() && !data.AllSystem.IsNull() {
		data.AllSystem = tfutils.BoolFromString(value.String())
	} else if !data.AllSystem.ValueBool() {
		data.AllSystem = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllSystemWarningThreshold`); value.Exists() && !data.AllSystemWarningThreshold.IsNull() {
		data.AllSystemWarningThreshold = types.Int64Value(value.Int())
	} else if data.AllSystemWarningThreshold.ValueInt64() != 75 {
		data.AllSystemWarningThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AllSystemCriticalThreshold`); value.Exists() && !data.AllSystemCriticalThreshold.IsNull() {
		data.AllSystemCriticalThreshold = types.Int64Value(value.Int())
	} else if data.AllSystemCriticalThreshold.ValueInt64() != 90 {
		data.AllSystemCriticalThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `System`); value.Exists() && !data.System.IsNull() {
		l := []DmFileSystemUsage{}
		for _, v := range value.Array() {
			item := DmFileSystemUsage{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.System, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFileSystemUsageObjectType}, l)
		} else {
			data.System = types.ListNull(types.ObjectType{AttrTypes: DmFileSystemUsageObjectType})
		}
	} else {
		data.System = types.ListNull(types.ObjectType{AttrTypes: DmFileSystemUsageObjectType})
	}
	if value := res.Get(pathRoot + `AllQueueManagers`); value.Exists() && !data.AllQueueManagers.IsNull() {
		data.AllQueueManagers = tfutils.BoolFromString(value.String())
	} else if !data.AllQueueManagers.ValueBool() {
		data.AllQueueManagers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllQMWarningThreshold`); value.Exists() && !data.AllQmWarningThreshold.IsNull() {
		data.AllQmWarningThreshold = types.Int64Value(value.Int())
	} else if data.AllQmWarningThreshold.ValueInt64() != 75 {
		data.AllQmWarningThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AllQMCriticalThreshold`); value.Exists() && !data.AllQmCriticalThreshold.IsNull() {
		data.AllQmCriticalThreshold = types.Int64Value(value.Int())
	} else if data.AllQmCriticalThreshold.ValueInt64() != 90 {
		data.AllQmCriticalThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `QueueManager`); value.Exists() && !data.QueueManager.IsNull() {
		l := []DmQMFileSystemUsage{}
		for _, v := range value.Array() {
			item := DmQMFileSystemUsage{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.QueueManager, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType}, l)
		} else {
			data.QueueManager = types.ListNull(types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType})
		}
	} else {
		data.QueueManager = types.ListNull(types.ObjectType{AttrTypes: DmQMFileSystemUsageObjectType})
	}
}
