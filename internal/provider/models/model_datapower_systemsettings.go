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

type SystemSettings struct {
	Enabled                 types.Bool        `tfsdk:"enabled"`
	UserSummary             types.String      `tfsdk:"user_summary"`
	ProductOid              types.String      `tfsdk:"product_oid"`
	Description             types.String      `tfsdk:"description"`
	SerialNumber            types.String      `tfsdk:"serial_number"`
	EntitlementNumber       types.String      `tfsdk:"entitlement_number"`
	ProductId               types.String      `tfsdk:"product_id"`
	CapacityMode            types.String      `tfsdk:"capacity_mode"`
	Contact                 types.String      `tfsdk:"contact"`
	SystemName              types.String      `tfsdk:"system_name"`
	Location                types.String      `tfsdk:"location"`
	Services                types.Int64       `tfsdk:"services"`
	BackupMode              types.String      `tfsdk:"backup_mode"`
	ProductMode             types.String      `tfsdk:"product_mode"`
	CustomUiFile            types.String      `tfsdk:"custom_ui_file"`
	AuditReserve            types.Int64       `tfsdk:"audit_reserve"`
	DetectIntrusion         types.String      `tfsdk:"detect_intrusion"`
	HardwareXmlAcceleration types.Bool        `tfsdk:"hardware_xml_acceleration"`
	Locale                  types.String      `tfsdk:"locale"`
	SystemLogFixedFormat    types.Bool        `tfsdk:"system_log_fixed_format"`
	Uuid                    types.String      `tfsdk:"uuid"`
	ObjectActions           []*actions.Action `tfsdk:"object_actions"`
}

var SystemSettingsObjectType = map[string]attr.Type{
	"enabled":                   types.BoolType,
	"user_summary":              types.StringType,
	"product_oid":               types.StringType,
	"description":               types.StringType,
	"serial_number":             types.StringType,
	"entitlement_number":        types.StringType,
	"product_id":                types.StringType,
	"capacity_mode":             types.StringType,
	"contact":                   types.StringType,
	"system_name":               types.StringType,
	"location":                  types.StringType,
	"services":                  types.Int64Type,
	"backup_mode":               types.StringType,
	"product_mode":              types.StringType,
	"custom_ui_file":            types.StringType,
	"audit_reserve":             types.Int64Type,
	"detect_intrusion":          types.StringType,
	"hardware_xml_acceleration": types.BoolType,
	"locale":                    types.StringType,
	"system_log_fixed_format":   types.BoolType,
	"uuid":                      types.StringType,
	"object_actions":            actions.ActionsListType,
}

func (data SystemSettings) GetPath() string {
	rest_path := "/mgmt/config/default/SystemSettings/System-Settings"
	return rest_path
}

func (data SystemSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ProductOid.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.SerialNumber.IsNull() {
		return false
	}
	if !data.EntitlementNumber.IsNull() {
		return false
	}
	if !data.ProductId.IsNull() {
		return false
	}
	if !data.CapacityMode.IsNull() {
		return false
	}
	if !data.Contact.IsNull() {
		return false
	}
	if !data.SystemName.IsNull() {
		return false
	}
	if !data.Location.IsNull() {
		return false
	}
	if !data.Services.IsNull() {
		return false
	}
	if !data.BackupMode.IsNull() {
		return false
	}
	if !data.ProductMode.IsNull() {
		return false
	}
	if !data.CustomUiFile.IsNull() {
		return false
	}
	if !data.AuditReserve.IsNull() {
		return false
	}
	if !data.DetectIntrusion.IsNull() {
		return false
	}
	if !data.HardwareXmlAcceleration.IsNull() {
		return false
	}
	if !data.Locale.IsNull() {
		return false
	}
	if !data.SystemLogFixedFormat.IsNull() {
		return false
	}
	if !data.Uuid.IsNull() {
		return false
	}
	return true
}

func (data SystemSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "SystemSettings.name", path.Base("/mgmt/config/default/SystemSettings/System-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.ProductOid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductOID`, data.ProductOid.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Description`, data.Description.ValueString())
	}
	if !data.SerialNumber.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SerialNumber`, data.SerialNumber.ValueString())
	}
	if !data.EntitlementNumber.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EntitlementNumber`, data.EntitlementNumber.ValueString())
	}
	if !data.ProductId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductID`, data.ProductId.ValueString())
	}
	if !data.CapacityMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CapacityMode`, data.CapacityMode.ValueString())
	}
	if !data.Contact.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Contact`, data.Contact.ValueString())
	}
	if !data.SystemName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SystemName`, data.SystemName.ValueString())
	}
	if !data.Location.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Location`, data.Location.ValueString())
	}
	if !data.Services.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Services`, data.Services.ValueInt64())
	}
	if !data.BackupMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BackupMode`, data.BackupMode.ValueString())
	}
	if !data.ProductMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductMode`, data.ProductMode.ValueString())
	}
	if !data.CustomUiFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomUIFile`, data.CustomUiFile.ValueString())
	}
	if !data.AuditReserve.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuditReserve`, data.AuditReserve.ValueInt64())
	}
	if !data.DetectIntrusion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DetectIntrusion`, data.DetectIntrusion.ValueString())
	}
	if !data.HardwareXmlAcceleration.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HardwareXMLAcceleration`, tfutils.StringFromBool(data.HardwareXmlAcceleration, ""))
	}
	if !data.Locale.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Locale`, data.Locale.ValueString())
	}
	if !data.SystemLogFixedFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SystemLogFixedFormat`, tfutils.StringFromBool(data.SystemLogFixedFormat, ""))
	}
	if !data.Uuid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UUID`, data.Uuid.ValueString())
	}
	return body
}

func (data *SystemSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ProductOID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductOid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductOid = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `SerialNumber`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SerialNumber = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get(pathRoot + `EntitlementNumber`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EntitlementNumber = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EntitlementNumber = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductId = types.StringNull()
	}
	if value := res.Get(pathRoot + `CapacityMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CapacityMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CapacityMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `Contact`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Contact = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Contact = types.StringNull()
	}
	if value := res.Get(pathRoot + `SystemName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SystemName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SystemName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Location`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Location = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get(pathRoot + `Services`); value.Exists() {
		data.Services = types.Int64Value(value.Int())
	} else {
		data.Services = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackupMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BackupMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackupMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomUIFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomUiFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomUiFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuditReserve`); value.Exists() {
		data.AuditReserve = types.Int64Value(value.Int())
	} else {
		data.AuditReserve = types.Int64Value(40)
	}
	if value := res.Get(pathRoot + `DetectIntrusion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DetectIntrusion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DetectIntrusion = types.StringNull()
	}
	if value := res.Get(pathRoot + `HardwareXMLAcceleration`); value.Exists() {
		data.HardwareXmlAcceleration = tfutils.BoolFromString(value.String())
	} else {
		data.HardwareXmlAcceleration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Locale`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Locale = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Locale = types.StringNull()
	}
	if value := res.Get(pathRoot + `SystemLogFixedFormat`); value.Exists() {
		data.SystemLogFixedFormat = tfutils.BoolFromString(value.String())
	} else {
		data.SystemLogFixedFormat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UUID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Uuid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uuid = types.StringNull()
	}
}

func (data *SystemSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ProductOID`); value.Exists() && !data.ProductOid.IsNull() {
		data.ProductOid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductOid = types.StringNull()
	}
	if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
		data.Description = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get(pathRoot + `SerialNumber`); value.Exists() && !data.SerialNumber.IsNull() {
		data.SerialNumber = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SerialNumber = types.StringNull()
	}
	if value := res.Get(pathRoot + `EntitlementNumber`); value.Exists() && !data.EntitlementNumber.IsNull() {
		data.EntitlementNumber = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EntitlementNumber = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductID`); value.Exists() && !data.ProductId.IsNull() {
		data.ProductId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductId = types.StringNull()
	}
	if value := res.Get(pathRoot + `CapacityMode`); value.Exists() && !data.CapacityMode.IsNull() {
		data.CapacityMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CapacityMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `Contact`); value.Exists() && !data.Contact.IsNull() {
		data.Contact = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Contact = types.StringNull()
	}
	if value := res.Get(pathRoot + `SystemName`); value.Exists() && !data.SystemName.IsNull() {
		data.SystemName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SystemName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Location`); value.Exists() && !data.Location.IsNull() {
		data.Location = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Location = types.StringNull()
	}
	if value := res.Get(pathRoot + `Services`); value.Exists() && !data.Services.IsNull() {
		data.Services = types.Int64Value(value.Int())
	} else {
		data.Services = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BackupMode`); value.Exists() && !data.BackupMode.IsNull() {
		data.BackupMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BackupMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductMode`); value.Exists() && !data.ProductMode.IsNull() {
		data.ProductMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `CustomUIFile`); value.Exists() && !data.CustomUiFile.IsNull() {
		data.CustomUiFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomUiFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuditReserve`); value.Exists() && !data.AuditReserve.IsNull() {
		data.AuditReserve = types.Int64Value(value.Int())
	} else if data.AuditReserve.ValueInt64() != 40 {
		data.AuditReserve = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DetectIntrusion`); value.Exists() && !data.DetectIntrusion.IsNull() {
		data.DetectIntrusion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DetectIntrusion = types.StringNull()
	}
	if value := res.Get(pathRoot + `HardwareXMLAcceleration`); value.Exists() && !data.HardwareXmlAcceleration.IsNull() {
		data.HardwareXmlAcceleration = tfutils.BoolFromString(value.String())
	} else {
		data.HardwareXmlAcceleration = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Locale`); value.Exists() && !data.Locale.IsNull() {
		data.Locale = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Locale = types.StringNull()
	}
	if value := res.Get(pathRoot + `SystemLogFixedFormat`); value.Exists() && !data.SystemLogFixedFormat.IsNull() {
		data.SystemLogFixedFormat = tfutils.BoolFromString(value.String())
	} else if data.SystemLogFixedFormat.ValueBool() {
		data.SystemLogFixedFormat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UUID`); value.Exists() && !data.Uuid.IsNull() {
		data.Uuid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uuid = types.StringNull()
	}
}
func (data *SystemSettings) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if data.ProductOid.IsUnknown() {
		if value := res.Get(pathRoot + `ProductOID`); value.Exists() && !data.ProductOid.IsNull() {
			data.ProductOid = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ProductOid = types.StringNull()
		}
	}
	if data.Description.IsUnknown() {
		if value := res.Get(pathRoot + `Description`); value.Exists() && !data.Description.IsNull() {
			data.Description = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Description = types.StringNull()
		}
	}
	if data.SerialNumber.IsUnknown() {
		if value := res.Get(pathRoot + `SerialNumber`); value.Exists() && !data.SerialNumber.IsNull() {
			data.SerialNumber = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SerialNumber = types.StringNull()
		}
	}
	if data.EntitlementNumber.IsUnknown() {
		if value := res.Get(pathRoot + `EntitlementNumber`); value.Exists() && !data.EntitlementNumber.IsNull() {
			data.EntitlementNumber = tfutils.ParseStringFromGJSON(value)
		} else {
			data.EntitlementNumber = types.StringNull()
		}
	}
	if data.ProductId.IsUnknown() {
		if value := res.Get(pathRoot + `ProductID`); value.Exists() && !data.ProductId.IsNull() {
			data.ProductId = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ProductId = types.StringNull()
		}
	}
	if data.CapacityMode.IsUnknown() {
		if value := res.Get(pathRoot + `CapacityMode`); value.Exists() && !data.CapacityMode.IsNull() {
			data.CapacityMode = tfutils.ParseStringFromGJSON(value)
		} else {
			data.CapacityMode = types.StringNull()
		}
	}
	if data.Contact.IsUnknown() {
		if value := res.Get(pathRoot + `Contact`); value.Exists() && !data.Contact.IsNull() {
			data.Contact = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Contact = types.StringNull()
		}
	}
	if data.SystemName.IsUnknown() {
		if value := res.Get(pathRoot + `SystemName`); value.Exists() && !data.SystemName.IsNull() {
			data.SystemName = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SystemName = types.StringNull()
		}
	}
	if data.Location.IsUnknown() {
		if value := res.Get(pathRoot + `Location`); value.Exists() && !data.Location.IsNull() {
			data.Location = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Location = types.StringNull()
		}
	}
	if data.Services.IsUnknown() {
		if value := res.Get(pathRoot + `Services`); value.Exists() && !data.Services.IsNull() {
			data.Services = types.Int64Value(value.Int())
		} else {
			data.Services = types.Int64Null()
		}
	}
	if data.BackupMode.IsUnknown() {
		if value := res.Get(pathRoot + `BackupMode`); value.Exists() && !data.BackupMode.IsNull() {
			data.BackupMode = tfutils.ParseStringFromGJSON(value)
		} else {
			data.BackupMode = types.StringNull()
		}
	}
	if data.ProductMode.IsUnknown() {
		if value := res.Get(pathRoot + `ProductMode`); value.Exists() && !data.ProductMode.IsNull() {
			data.ProductMode = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ProductMode = types.StringNull()
		}
	}
	if data.CustomUiFile.IsUnknown() {
		if value := res.Get(pathRoot + `CustomUIFile`); value.Exists() && !data.CustomUiFile.IsNull() {
			data.CustomUiFile = tfutils.ParseStringFromGJSON(value)
		} else {
			data.CustomUiFile = types.StringNull()
		}
	}
	if data.AuditReserve.IsUnknown() {
		if value := res.Get(pathRoot + `AuditReserve`); value.Exists() && !data.AuditReserve.IsNull() {
			data.AuditReserve = types.Int64Value(value.Int())
		} else if data.AuditReserve.ValueInt64() != 40 {
			data.AuditReserve = types.Int64Null()
		}
	}
	if data.DetectIntrusion.IsUnknown() {
		if value := res.Get(pathRoot + `DetectIntrusion`); value.Exists() && !data.DetectIntrusion.IsNull() {
			data.DetectIntrusion = tfutils.ParseStringFromGJSON(value)
		} else {
			data.DetectIntrusion = types.StringNull()
		}
	}
	if data.HardwareXmlAcceleration.IsUnknown() {
		if value := res.Get(pathRoot + `HardwareXMLAcceleration`); value.Exists() && !data.HardwareXmlAcceleration.IsNull() {
			data.HardwareXmlAcceleration = tfutils.BoolFromString(value.String())
		} else {
			data.HardwareXmlAcceleration = types.BoolNull()
		}
	}
	if data.Locale.IsUnknown() {
		if value := res.Get(pathRoot + `Locale`); value.Exists() && !data.Locale.IsNull() {
			data.Locale = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Locale = types.StringNull()
		}
	}
	if data.SystemLogFixedFormat.IsUnknown() {
		if value := res.Get(pathRoot + `SystemLogFixedFormat`); value.Exists() && !data.SystemLogFixedFormat.IsNull() {
			data.SystemLogFixedFormat = tfutils.BoolFromString(value.String())
		} else {
			data.SystemLogFixedFormat = types.BoolNull()
		}
	}
	if data.Uuid.IsUnknown() {
		if value := res.Get(pathRoot + `UUID`); value.Exists() && !data.Uuid.IsNull() {
			data.Uuid = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Uuid = types.StringNull()
		}
	}
}
