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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmConfigSequenceCapabilities struct {
	ApiConnect         types.Bool `tfsdk:"api_connect"`
	MonitorPersistence types.Bool `tfsdk:"monitor_persistence"`
	ApplyAllObjects    types.Bool `tfsdk:"apply_all_objects"`
	MarkExternal       types.Bool `tfsdk:"mark_external"`
	DeleteConfig       types.Bool `tfsdk:"delete_config"`
	Batch              types.Bool `tfsdk:"batch"`
	GitOpsMode         types.Bool `tfsdk:"git_ops_mode"`
}

var DmConfigSequenceCapabilitiesObjectType = map[string]attr.Type{
	"api_connect":         types.BoolType,
	"monitor_persistence": types.BoolType,
	"apply_all_objects":   types.BoolType,
	"mark_external":       types.BoolType,
	"delete_config":       types.BoolType,
	"batch":               types.BoolType,
	"git_ops_mode":        types.BoolType,
}
var DmConfigSequenceCapabilitiesObjectDefault = map[string]attr.Value{
	"api_connect":         types.BoolNull(),
	"monitor_persistence": types.BoolNull(),
	"apply_all_objects":   types.BoolNull(),
	"mark_external":       types.BoolNull(),
	"delete_config":       types.BoolNull(),
	"batch":               types.BoolNull(),
	"git_ops_mode":        types.BoolNull(),
}
var DmConfigSequenceCapabilitiesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"api_connect": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Optimize for API Connect processing", "", "").String,
			Computed:            true,
		},
		"monitor_persistence": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Monitor changes for persistence", "", "").String,
			Computed:            true,
		},
		"apply_all_objects": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Apply to all objects", "", "").String,
			Computed:            true,
		},
		"mark_external": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Mark objects as external", "", "").String,
			Computed:            true,
		},
		"delete_config": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Delete files after processing", "", "").String,
			Computed:            true,
		},
		"batch": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Batch changes from configuration files", "", "").String,
			Computed:            true,
		},
		"git_ops_mode": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Process in GitOps mode", "", "").String,
			Computed:            true,
		},
	},
}
var DmConfigSequenceCapabilitiesResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmConfigSequenceCapabilitiesObjectType,
			DmConfigSequenceCapabilitiesObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"api_connect": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Optimize for API Connect processing", "", "").String,
			Optional:            true,
		},
		"monitor_persistence": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Monitor changes for persistence", "", "").String,
			Optional:            true,
		},
		"apply_all_objects": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Apply to all objects", "", "").String,
			Optional:            true,
		},
		"mark_external": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Mark objects as external", "", "").String,
			Optional:            true,
		},
		"delete_config": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Delete files after processing", "", "").String,
			Optional:            true,
		},
		"batch": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Batch changes from configuration files", "", "").String,
			Optional:            true,
		},
		"git_ops_mode": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Process in GitOps mode", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmConfigSequenceCapabilities) IsNull() bool {
	if !data.ApiConnect.IsNull() {
		return false
	}
	if !data.MonitorPersistence.IsNull() {
		return false
	}
	if !data.ApplyAllObjects.IsNull() {
		return false
	}
	if !data.MarkExternal.IsNull() {
		return false
	}
	if !data.DeleteConfig.IsNull() {
		return false
	}
	if !data.Batch.IsNull() {
		return false
	}
	if !data.GitOpsMode.IsNull() {
		return false
	}
	return true
}
func GetDmConfigSequenceCapabilitiesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmConfigSequenceCapabilitiesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmConfigSequenceCapabilitiesDataSourceSchema
}

func GetDmConfigSequenceCapabilitiesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmConfigSequenceCapabilitiesResourceSchema.Required = true
	} else {
		DmConfigSequenceCapabilitiesResourceSchema.Optional = true
		DmConfigSequenceCapabilitiesResourceSchema.Computed = true
	}
	DmConfigSequenceCapabilitiesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmConfigSequenceCapabilitiesResourceSchema
}

func (data DmConfigSequenceCapabilities) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.ApiConnect.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIConnect`, tfutils.StringFromBool(data.ApiConnect, ""))
	}
	if !data.MonitorPersistence.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MonitorPersistence`, tfutils.StringFromBool(data.MonitorPersistence, ""))
	}
	if !data.ApplyAllObjects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ApplyAllObjects`, tfutils.StringFromBool(data.ApplyAllObjects, ""))
	}
	if !data.MarkExternal.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MarkExternal`, tfutils.StringFromBool(data.MarkExternal, ""))
	}
	if !data.DeleteConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeleteConfig`, tfutils.StringFromBool(data.DeleteConfig, ""))
	}
	if !data.Batch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Batch`, tfutils.StringFromBool(data.Batch, ""))
	}
	if !data.GitOpsMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GitOpsMode`, tfutils.StringFromBool(data.GitOpsMode, ""))
	}
	return body
}

func (data *DmConfigSequenceCapabilities) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `APIConnect`); value.Exists() {
		data.ApiConnect = tfutils.BoolFromString(value.String())
	} else {
		data.ApiConnect = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MonitorPersistence`); value.Exists() {
		data.MonitorPersistence = tfutils.BoolFromString(value.String())
	} else {
		data.MonitorPersistence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplyAllObjects`); value.Exists() {
		data.ApplyAllObjects = tfutils.BoolFromString(value.String())
	} else {
		data.ApplyAllObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MarkExternal`); value.Exists() {
		data.MarkExternal = tfutils.BoolFromString(value.String())
	} else {
		data.MarkExternal = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DeleteConfig`); value.Exists() {
		data.DeleteConfig = tfutils.BoolFromString(value.String())
	} else {
		data.DeleteConfig = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Batch`); value.Exists() {
		data.Batch = tfutils.BoolFromString(value.String())
	} else {
		data.Batch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GitOpsMode`); value.Exists() {
		data.GitOpsMode = tfutils.BoolFromString(value.String())
	} else {
		data.GitOpsMode = types.BoolNull()
	}
}

func (data *DmConfigSequenceCapabilities) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `APIConnect`); value.Exists() && !data.ApiConnect.IsNull() {
		data.ApiConnect = tfutils.BoolFromString(value.String())
	} else {
		data.ApiConnect = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MonitorPersistence`); value.Exists() && !data.MonitorPersistence.IsNull() {
		data.MonitorPersistence = tfutils.BoolFromString(value.String())
	} else {
		data.MonitorPersistence = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplyAllObjects`); value.Exists() && !data.ApplyAllObjects.IsNull() {
		data.ApplyAllObjects = tfutils.BoolFromString(value.String())
	} else {
		data.ApplyAllObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MarkExternal`); value.Exists() && !data.MarkExternal.IsNull() {
		data.MarkExternal = tfutils.BoolFromString(value.String())
	} else {
		data.MarkExternal = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DeleteConfig`); value.Exists() && !data.DeleteConfig.IsNull() {
		data.DeleteConfig = tfutils.BoolFromString(value.String())
	} else {
		data.DeleteConfig = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Batch`); value.Exists() && !data.Batch.IsNull() {
		data.Batch = tfutils.BoolFromString(value.String())
	} else {
		data.Batch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GitOpsMode`); value.Exists() && !data.GitOpsMode.IsNull() {
		data.GitOpsMode = tfutils.BoolFromString(value.String())
	} else {
		data.GitOpsMode = types.BoolNull()
	}
}
