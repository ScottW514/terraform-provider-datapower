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

type Domain struct {
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	ConfigDir                  types.String                `tfsdk:"config_dir"`
	NeighborDomain             types.List                  `tfsdk:"neighbor_domain"`
	DomainUser                 types.List                  `tfsdk:"domain_user"`
	FileMap                    *DmDomainFileMap            `tfsdk:"file_map"`
	MonitoringMap              *DmDomainMonitoringMap      `tfsdk:"monitoring_map"`
	ConfigMode                 types.String                `tfsdk:"config_mode"`
	ImportUrl                  types.String                `tfsdk:"import_url"`
	ImportFormat               types.String                `tfsdk:"import_format"`
	DeploymentPolicy           types.String                `tfsdk:"deployment_policy"`
	DeploymentPolicyParameters types.String                `tfsdk:"deployment_policy_parameters"`
	LocalIpRewrite             types.Bool                  `tfsdk:"local_ip_rewrite"`
	MaxChkpoints               types.Int64                 `tfsdk:"max_chkpoints"`
	ConfigPermissionsMode      types.String                `tfsdk:"config_permissions_mode"`
	ConfigPermissionsProfile   types.String                `tfsdk:"config_permissions_profile"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var DomainObjectType = map[string]attr.Type{
	"app_domain":                   types.StringType,
	"user_summary":                 types.StringType,
	"config_dir":                   types.StringType,
	"neighbor_domain":              types.ListType{ElemType: types.StringType},
	"domain_user":                  types.ListType{ElemType: types.StringType},
	"file_map":                     types.ObjectType{AttrTypes: DmDomainFileMapObjectType},
	"monitoring_map":               types.ObjectType{AttrTypes: DmDomainMonitoringMapObjectType},
	"config_mode":                  types.StringType,
	"import_url":                   types.StringType,
	"import_format":                types.StringType,
	"deployment_policy":            types.StringType,
	"deployment_policy_parameters": types.StringType,
	"local_ip_rewrite":             types.BoolType,
	"max_chkpoints":                types.Int64Type,
	"config_permissions_mode":      types.StringType,
	"config_permissions_profile":   types.StringType,
	"dependency_actions":           actions.ActionsListType,
}

func (data Domain) GetPath() string {
	rest_path := "/mgmt/config/default/Domain/{domain}"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data Domain) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ConfigDir.IsNull() {
		return false
	}
	if !data.NeighborDomain.IsNull() {
		return false
	}
	if !data.DomainUser.IsNull() {
		return false
	}
	if data.FileMap != nil {
		if !data.FileMap.IsNull() {
			return false
		}
	}
	if data.MonitoringMap != nil {
		if !data.MonitoringMap.IsNull() {
			return false
		}
	}
	if !data.ConfigMode.IsNull() {
		return false
	}
	if !data.ImportUrl.IsNull() {
		return false
	}
	if !data.ImportFormat.IsNull() {
		return false
	}
	if !data.DeploymentPolicy.IsNull() {
		return false
	}
	if !data.DeploymentPolicyParameters.IsNull() {
		return false
	}
	if !data.LocalIpRewrite.IsNull() {
		return false
	}
	if !data.MaxChkpoints.IsNull() {
		return false
	}
	if !data.ConfigPermissionsMode.IsNull() {
		return false
	}
	if !data.ConfigPermissionsProfile.IsNull() {
		return false
	}
	return true
}

func (data Domain) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "Domain.name", data.AppDomain.ValueString())
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ConfigDir.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigDir`, data.ConfigDir.ValueString())
	}
	if !data.NeighborDomain.IsNull() {
		var values []string
		data.NeighborDomain.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`NeighborDomain`+".-1", map[string]string{"value": val})
		}
	}
	if !data.DomainUser.IsNull() {
		var values []string
		data.DomainUser.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`DomainUser`+".-1", map[string]string{"value": val})
		}
	}
	if data.FileMap != nil {
		if !data.FileMap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`FileMap`, data.FileMap.ToBody(ctx, ""))
		}
	}
	if data.MonitoringMap != nil {
		if !data.MonitoringMap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`MonitoringMap`, data.MonitoringMap.ToBody(ctx, ""))
		}
	}
	if !data.ConfigMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigMode`, data.ConfigMode.ValueString())
	}
	if !data.ImportUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ImportURL`, data.ImportUrl.ValueString())
	}
	if !data.ImportFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ImportFormat`, data.ImportFormat.ValueString())
	}
	if !data.DeploymentPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeploymentPolicy`, data.DeploymentPolicy.ValueString())
	}
	if !data.DeploymentPolicyParameters.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeploymentPolicyParameters`, data.DeploymentPolicyParameters.ValueString())
	}
	if !data.LocalIpRewrite.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalIPRewrite`, tfutils.StringFromBool(data.LocalIpRewrite, ""))
	}
	if !data.MaxChkpoints.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxChkpoints`, data.MaxChkpoints.ValueInt64())
	}
	if !data.ConfigPermissionsMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigPermissionsMode`, data.ConfigPermissionsMode.ValueString())
	}
	if !data.ConfigPermissionsProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConfigPermissionsProfile`, data.ConfigPermissionsProfile.ValueString())
	}
	return body
}

func (data *Domain) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigDir`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigDir = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigDir = types.StringNull()
	}
	if value := res.Get(pathRoot + `NeighborDomain`); value.Exists() {
		data.NeighborDomain = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.NeighborDomain = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DomainUser`); value.Exists() {
		data.DomainUser = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DomainUser = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FileMap`); value.Exists() {
		data.FileMap = &DmDomainFileMap{}
		data.FileMap.FromBody(ctx, "", value)
	} else {
		data.FileMap = nil
	}
	if value := res.Get(pathRoot + `MonitoringMap`); value.Exists() {
		data.MonitoringMap = &DmDomainMonitoringMap{}
		data.MonitoringMap.FromBody(ctx, "", value)
	} else {
		data.MonitoringMap = nil
	}
	if value := res.Get(pathRoot + `ConfigMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigMode = types.StringValue("local")
	}
	if value := res.Get(pathRoot + `ImportURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ImportUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ImportUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ImportFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ImportFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ImportFormat = types.StringValue("ZIP")
	}
	if value := res.Get(pathRoot + `DeploymentPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeploymentPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicyParameters`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeploymentPolicyParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicyParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalIPRewrite`); value.Exists() {
		data.LocalIpRewrite = tfutils.BoolFromString(value.String())
	} else {
		data.LocalIpRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxChkpoints`); value.Exists() {
		data.MaxChkpoints = types.Int64Value(value.Int())
	} else {
		data.MaxChkpoints = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `ConfigPermissionsMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigPermissionsMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigPermissionsMode = types.StringValue("scope-domain")
	}
	if value := res.Get(pathRoot + `ConfigPermissionsProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ConfigPermissionsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigPermissionsProfile = types.StringNull()
	}
}

func (data *Domain) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigDir`); value.Exists() && !data.ConfigDir.IsNull() {
		data.ConfigDir = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigDir = types.StringNull()
	}
	if value := res.Get(pathRoot + `NeighborDomain`); value.Exists() && !data.NeighborDomain.IsNull() {
		data.NeighborDomain = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.NeighborDomain = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DomainUser`); value.Exists() && !data.DomainUser.IsNull() {
		data.DomainUser = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.DomainUser = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `FileMap`); value.Exists() {
		data.FileMap.UpdateFromBody(ctx, "", value)
	} else {
		data.FileMap = nil
	}
	if value := res.Get(pathRoot + `MonitoringMap`); value.Exists() {
		data.MonitoringMap.UpdateFromBody(ctx, "", value)
	} else {
		data.MonitoringMap = nil
	}
	if value := res.Get(pathRoot + `ConfigMode`); value.Exists() && !data.ConfigMode.IsNull() {
		data.ConfigMode = tfutils.ParseStringFromGJSON(value)
	} else if data.ConfigMode.ValueString() != "local" {
		data.ConfigMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ImportURL`); value.Exists() && !data.ImportUrl.IsNull() {
		data.ImportUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ImportUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ImportFormat`); value.Exists() && !data.ImportFormat.IsNull() {
		data.ImportFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.ImportFormat.ValueString() != "ZIP" {
		data.ImportFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicy`); value.Exists() && !data.DeploymentPolicy.IsNull() {
		data.DeploymentPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeploymentPolicyParameters`); value.Exists() && !data.DeploymentPolicyParameters.IsNull() {
		data.DeploymentPolicyParameters = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentPolicyParameters = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalIPRewrite`); value.Exists() && !data.LocalIpRewrite.IsNull() {
		data.LocalIpRewrite = tfutils.BoolFromString(value.String())
	} else if !data.LocalIpRewrite.ValueBool() {
		data.LocalIpRewrite = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxChkpoints`); value.Exists() && !data.MaxChkpoints.IsNull() {
		data.MaxChkpoints = types.Int64Value(value.Int())
	} else if data.MaxChkpoints.ValueInt64() != 3 {
		data.MaxChkpoints = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConfigPermissionsMode`); value.Exists() && !data.ConfigPermissionsMode.IsNull() {
		data.ConfigPermissionsMode = tfutils.ParseStringFromGJSON(value)
	} else if data.ConfigPermissionsMode.ValueString() != "scope-domain" {
		data.ConfigPermissionsMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ConfigPermissionsProfile`); value.Exists() && !data.ConfigPermissionsProfile.IsNull() {
		data.ConfigPermissionsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ConfigPermissionsProfile = types.StringNull()
	}
}
func (data *Domain) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.ConfigDir.IsUnknown() {
		if value := res.Get(pathRoot + `ConfigDir`); value.Exists() && !data.ConfigDir.IsNull() {
			data.ConfigDir = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ConfigDir = types.StringNull()
		}
	}
	if data.NeighborDomain.IsUnknown() {
		if value := res.Get(pathRoot + `NeighborDomain`); value.Exists() && !data.NeighborDomain.IsNull() {
			data.NeighborDomain = tfutils.ParseStringListFromGJSON(value)
		} else {
			data.NeighborDomain = types.ListNull(types.StringType)
		}
	}
	if data.DomainUser.IsUnknown() {
		if value := res.Get(pathRoot + `DomainUser`); value.Exists() && !data.DomainUser.IsNull() {
			data.DomainUser = tfutils.ParseStringListFromGJSON(value)
		} else {
			data.DomainUser = types.ListNull(types.StringType)
		}
	}
	if data.FileMap == nil {
		if value := res.Get(pathRoot + `FileMap`); value.Exists() {
			d := DmDomainFileMap{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.FileMap = &d
			}
		}
	}
	if data.MonitoringMap == nil {
		if value := res.Get(pathRoot + `MonitoringMap`); value.Exists() {
			d := DmDomainMonitoringMap{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.MonitoringMap = &d
			}
		}
	}
	if data.ConfigMode.IsUnknown() {
		if value := res.Get(pathRoot + `ConfigMode`); value.Exists() && !data.ConfigMode.IsNull() {
			data.ConfigMode = tfutils.ParseStringFromGJSON(value)
		} else if data.ConfigMode.ValueString() != "local" {
			data.ConfigMode = types.StringNull()
		}
	}
	if data.ImportUrl.IsUnknown() {
		if value := res.Get(pathRoot + `ImportURL`); value.Exists() && !data.ImportUrl.IsNull() {
			data.ImportUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ImportUrl = types.StringNull()
		}
	}
	if data.ImportFormat.IsUnknown() {
		if value := res.Get(pathRoot + `ImportFormat`); value.Exists() && !data.ImportFormat.IsNull() {
			data.ImportFormat = tfutils.ParseStringFromGJSON(value)
		} else if data.ImportFormat.ValueString() != "ZIP" {
			data.ImportFormat = types.StringNull()
		}
	}
	if data.DeploymentPolicy.IsUnknown() {
		if value := res.Get(pathRoot + `DeploymentPolicy`); value.Exists() && !data.DeploymentPolicy.IsNull() {
			data.DeploymentPolicy = tfutils.ParseStringFromGJSON(value)
		} else {
			data.DeploymentPolicy = types.StringNull()
		}
	}
	if data.DeploymentPolicyParameters.IsUnknown() {
		if value := res.Get(pathRoot + `DeploymentPolicyParameters`); value.Exists() && !data.DeploymentPolicyParameters.IsNull() {
			data.DeploymentPolicyParameters = tfutils.ParseStringFromGJSON(value)
		} else {
			data.DeploymentPolicyParameters = types.StringNull()
		}
	}
	if data.LocalIpRewrite.IsUnknown() {
		if value := res.Get(pathRoot + `LocalIPRewrite`); value.Exists() && !data.LocalIpRewrite.IsNull() {
			data.LocalIpRewrite = tfutils.BoolFromString(value.String())
		} else {
			data.LocalIpRewrite = types.BoolNull()
		}
	}
	if data.MaxChkpoints.IsUnknown() {
		if value := res.Get(pathRoot + `MaxChkpoints`); value.Exists() && !data.MaxChkpoints.IsNull() {
			data.MaxChkpoints = types.Int64Value(value.Int())
		} else if data.MaxChkpoints.ValueInt64() != 3 {
			data.MaxChkpoints = types.Int64Null()
		}
	}
	if data.ConfigPermissionsMode.IsUnknown() {
		if value := res.Get(pathRoot + `ConfigPermissionsMode`); value.Exists() && !data.ConfigPermissionsMode.IsNull() {
			data.ConfigPermissionsMode = tfutils.ParseStringFromGJSON(value)
		} else if data.ConfigPermissionsMode.ValueString() != "scope-domain" {
			data.ConfigPermissionsMode = types.StringNull()
		}
	}
	if data.ConfigPermissionsProfile.IsUnknown() {
		if value := res.Get(pathRoot + `ConfigPermissionsProfile`); value.Exists() && !data.ConfigPermissionsProfile.IsNull() {
			data.ConfigPermissionsProfile = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ConfigPermissionsProfile = types.StringNull()
		}
	}
}
