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

type ErrorReportSettings struct {
	Enabled               types.Bool                  `tfsdk:"enabled"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	UploadReport          types.Bool                  `tfsdk:"upload_report"`
	UseSmtp               types.Bool                  `tfsdk:"use_smtp"`
	InternalState         types.Bool                  `tfsdk:"internal_state"`
	FfdcPacketCapture     types.Bool                  `tfsdk:"ffdc_packet_capture"`
	FfdcEventLogCapture   types.Bool                  `tfsdk:"ffdc_event_log_capture"`
	FfdcMemoryLeakCapture types.Bool                  `tfsdk:"ffdc_memory_leak_capture"`
	AlwaysOnStartup       types.Bool                  `tfsdk:"always_on_startup"`
	AlwaysOnShutdown      types.Bool                  `tfsdk:"always_on_shutdown"`
	Protocol              types.String                `tfsdk:"protocol"`
	LocationIdentifier    types.String                `tfsdk:"location_identifier"`
	SmtpServer            types.String                `tfsdk:"smtp_server"`
	EmailAddress          types.String                `tfsdk:"email_address"`
	EmailSenderAddress    types.String                `tfsdk:"email_sender_address"`
	FtpServer             types.String                `tfsdk:"ftp_server"`
	FtpPath               types.String                `tfsdk:"ftp_path"`
	FtpUserAgent          types.String                `tfsdk:"ftp_user_agent"`
	NfsMount              types.String                `tfsdk:"nfs_mount"`
	NfsPath               types.String                `tfsdk:"nfs_path"`
	RaidVolume            types.String                `tfsdk:"raid_volume"`
	RaidPath              types.String                `tfsdk:"raid_path"`
	ReportHistoryKept     types.Int64                 `tfsdk:"report_history_kept"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var ErrorReportSettingsProtocolCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "upload_report",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var ErrorReportSettingsLocationIdentifierCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_smtp",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var ErrorReportSettingsSmtpServerCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "use_smtp",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_report",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "protocol",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var ErrorReportSettingsEmailAddressCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "use_smtp",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_report",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "protocol",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var ErrorReportSettingsFTPServerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_report",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "protocol",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"ftp"},
		},
	},
}
var ErrorReportSettingsFTPUserAgentCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_report",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "protocol",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"ftp"},
		},
	},
}
var ErrorReportSettingsNFSMountCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_report",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "protocol",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"nfs"},
		},
	},
}
var ErrorReportSettingsRaidVolumeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_report",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "protocol",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"raid"},
		},
	},
}

var ErrorReportSettingsObjectType = map[string]attr.Type{
	"enabled":                  types.BoolType,
	"user_summary":             types.StringType,
	"upload_report":            types.BoolType,
	"use_smtp":                 types.BoolType,
	"internal_state":           types.BoolType,
	"ffdc_packet_capture":      types.BoolType,
	"ffdc_event_log_capture":   types.BoolType,
	"ffdc_memory_leak_capture": types.BoolType,
	"always_on_startup":        types.BoolType,
	"always_on_shutdown":       types.BoolType,
	"protocol":                 types.StringType,
	"location_identifier":      types.StringType,
	"smtp_server":              types.StringType,
	"email_address":            types.StringType,
	"email_sender_address":     types.StringType,
	"ftp_server":               types.StringType,
	"ftp_path":                 types.StringType,
	"ftp_user_agent":           types.StringType,
	"nfs_mount":                types.StringType,
	"nfs_path":                 types.StringType,
	"raid_volume":              types.StringType,
	"raid_path":                types.StringType,
	"report_history_kept":      types.Int64Type,
	"dependency_actions":       actions.ActionsListType,
}

func (data ErrorReportSettings) GetPath() string {
	rest_path := "/mgmt/config/default/ErrorReportSettings/Error-Report"
	return rest_path
}

func (data ErrorReportSettings) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UploadReport.IsNull() {
		return false
	}
	if !data.UseSmtp.IsNull() {
		return false
	}
	if !data.InternalState.IsNull() {
		return false
	}
	if !data.FfdcPacketCapture.IsNull() {
		return false
	}
	if !data.FfdcEventLogCapture.IsNull() {
		return false
	}
	if !data.FfdcMemoryLeakCapture.IsNull() {
		return false
	}
	if !data.AlwaysOnStartup.IsNull() {
		return false
	}
	if !data.AlwaysOnShutdown.IsNull() {
		return false
	}
	if !data.Protocol.IsNull() {
		return false
	}
	if !data.LocationIdentifier.IsNull() {
		return false
	}
	if !data.SmtpServer.IsNull() {
		return false
	}
	if !data.EmailAddress.IsNull() {
		return false
	}
	if !data.EmailSenderAddress.IsNull() {
		return false
	}
	if !data.FtpServer.IsNull() {
		return false
	}
	if !data.FtpPath.IsNull() {
		return false
	}
	if !data.FtpUserAgent.IsNull() {
		return false
	}
	if !data.NfsMount.IsNull() {
		return false
	}
	if !data.NfsPath.IsNull() {
		return false
	}
	if !data.RaidVolume.IsNull() {
		return false
	}
	if !data.RaidPath.IsNull() {
		return false
	}
	if !data.ReportHistoryKept.IsNull() {
		return false
	}
	return true
}

func (data ErrorReportSettings) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "ErrorReportSettings.name", path.Base("/mgmt/config/default/ErrorReportSettings/Error-Report"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.UploadReport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UploadReport`, tfutils.StringFromBool(data.UploadReport, ""))
	}
	if !data.UseSmtp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseSmtp`, tfutils.StringFromBool(data.UseSmtp, ""))
	}
	if !data.InternalState.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InternalState`, tfutils.StringFromBool(data.InternalState, ""))
	}
	if !data.FfdcPacketCapture.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FFDCPacketCapture`, tfutils.StringFromBool(data.FfdcPacketCapture, ""))
	}
	if !data.FfdcEventLogCapture.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FFDCEventLogCapture`, tfutils.StringFromBool(data.FfdcEventLogCapture, ""))
	}
	if !data.FfdcMemoryLeakCapture.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FFDCMemoryLeakCapture`, tfutils.StringFromBool(data.FfdcMemoryLeakCapture, ""))
	}
	if !data.AlwaysOnStartup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AlwaysOnStartup`, tfutils.StringFromBool(data.AlwaysOnStartup, ""))
	}
	if !data.AlwaysOnShutdown.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AlwaysOnShutdown`, tfutils.StringFromBool(data.AlwaysOnShutdown, ""))
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Protocol`, data.Protocol.ValueString())
	}
	if !data.LocationIdentifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocationIdentifier`, data.LocationIdentifier.ValueString())
	}
	if !data.SmtpServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SmtpServer`, data.SmtpServer.ValueString())
	}
	if !data.EmailAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EmailAddress`, data.EmailAddress.ValueString())
	}
	if !data.EmailSenderAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EmailSenderAddress`, data.EmailSenderAddress.ValueString())
	}
	if !data.FtpServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FTPServer`, data.FtpServer.ValueString())
	}
	if !data.FtpPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FTPPath`, data.FtpPath.ValueString())
	}
	if !data.FtpUserAgent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FTPUserAgent`, data.FtpUserAgent.ValueString())
	}
	if !data.NfsMount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NFSMount`, data.NfsMount.ValueString())
	}
	if !data.NfsPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NFSPath`, data.NfsPath.ValueString())
	}
	if !data.RaidVolume.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RaidVolume`, data.RaidVolume.ValueString())
	}
	if !data.RaidPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RaidPath`, data.RaidPath.ValueString())
	}
	if !data.ReportHistoryKept.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReportHistoryKept`, data.ReportHistoryKept.ValueInt64())
	}
	return body
}

func (data *ErrorReportSettings) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UploadReport`); value.Exists() {
		data.UploadReport = tfutils.BoolFromString(value.String())
	} else {
		data.UploadReport = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseSmtp`); value.Exists() {
		data.UseSmtp = tfutils.BoolFromString(value.String())
	} else {
		data.UseSmtp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InternalState`); value.Exists() {
		data.InternalState = tfutils.BoolFromString(value.String())
	} else {
		data.InternalState = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCPacketCapture`); value.Exists() {
		data.FfdcPacketCapture = tfutils.BoolFromString(value.String())
	} else {
		data.FfdcPacketCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCEventLogCapture`); value.Exists() {
		data.FfdcEventLogCapture = tfutils.BoolFromString(value.String())
	} else {
		data.FfdcEventLogCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCMemoryLeakCapture`); value.Exists() {
		data.FfdcMemoryLeakCapture = tfutils.BoolFromString(value.String())
	} else {
		data.FfdcMemoryLeakCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysOnStartup`); value.Exists() {
		data.AlwaysOnStartup = tfutils.BoolFromString(value.String())
	} else {
		data.AlwaysOnStartup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysOnShutdown`); value.Exists() {
		data.AlwaysOnShutdown = tfutils.BoolFromString(value.String())
	} else {
		data.AlwaysOnShutdown = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Protocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Protocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocationIdentifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocationIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocationIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `SmtpServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SmtpServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailSenderAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EmailSenderAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailSenderAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FtpServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FtpPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPUserAgent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FtpUserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpUserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSMount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfsMount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsMount = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfsPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RaidPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReportHistoryKept`); value.Exists() {
		data.ReportHistoryKept = types.Int64Value(value.Int())
	} else {
		data.ReportHistoryKept = types.Int64Value(5)
	}
}

func (data *ErrorReportSettings) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UploadReport`); value.Exists() && !data.UploadReport.IsNull() {
		data.UploadReport = tfutils.BoolFromString(value.String())
	} else if data.UploadReport.ValueBool() {
		data.UploadReport = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseSmtp`); value.Exists() && !data.UseSmtp.IsNull() {
		data.UseSmtp = tfutils.BoolFromString(value.String())
	} else if data.UseSmtp.ValueBool() {
		data.UseSmtp = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InternalState`); value.Exists() && !data.InternalState.IsNull() {
		data.InternalState = tfutils.BoolFromString(value.String())
	} else if data.InternalState.ValueBool() {
		data.InternalState = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCPacketCapture`); value.Exists() && !data.FfdcPacketCapture.IsNull() {
		data.FfdcPacketCapture = tfutils.BoolFromString(value.String())
	} else if data.FfdcPacketCapture.ValueBool() {
		data.FfdcPacketCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCEventLogCapture`); value.Exists() && !data.FfdcEventLogCapture.IsNull() {
		data.FfdcEventLogCapture = tfutils.BoolFromString(value.String())
	} else if data.FfdcEventLogCapture.ValueBool() {
		data.FfdcEventLogCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `FFDCMemoryLeakCapture`); value.Exists() && !data.FfdcMemoryLeakCapture.IsNull() {
		data.FfdcMemoryLeakCapture = tfutils.BoolFromString(value.String())
	} else if data.FfdcMemoryLeakCapture.ValueBool() {
		data.FfdcMemoryLeakCapture = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysOnStartup`); value.Exists() && !data.AlwaysOnStartup.IsNull() {
		data.AlwaysOnStartup = tfutils.BoolFromString(value.String())
	} else if !data.AlwaysOnStartup.ValueBool() {
		data.AlwaysOnStartup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlwaysOnShutdown`); value.Exists() && !data.AlwaysOnShutdown.IsNull() {
		data.AlwaysOnShutdown = tfutils.BoolFromString(value.String())
	} else if !data.AlwaysOnShutdown.ValueBool() {
		data.AlwaysOnShutdown = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Protocol`); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocationIdentifier`); value.Exists() && !data.LocationIdentifier.IsNull() {
		data.LocationIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocationIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `SmtpServer`); value.Exists() && !data.SmtpServer.IsNull() {
		data.SmtpServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && !data.EmailAddress.IsNull() {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailSenderAddress`); value.Exists() && !data.EmailSenderAddress.IsNull() {
		data.EmailSenderAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailSenderAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPServer`); value.Exists() && !data.FtpServer.IsNull() {
		data.FtpServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPPath`); value.Exists() && !data.FtpPath.IsNull() {
		data.FtpPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `FTPUserAgent`); value.Exists() && !data.FtpUserAgent.IsNull() {
		data.FtpUserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FtpUserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSMount`); value.Exists() && !data.NfsMount.IsNull() {
		data.NfsMount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsMount = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSPath`); value.Exists() && !data.NfsPath.IsNull() {
		data.NfsPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidVolume`); value.Exists() && !data.RaidVolume.IsNull() {
		data.RaidVolume = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidVolume = types.StringNull()
	}
	if value := res.Get(pathRoot + `RaidPath`); value.Exists() && !data.RaidPath.IsNull() {
		data.RaidPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RaidPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReportHistoryKept`); value.Exists() && !data.ReportHistoryKept.IsNull() {
		data.ReportHistoryKept = types.Int64Value(value.Int())
	} else if data.ReportHistoryKept.ValueInt64() != 5 {
		data.ReportHistoryKept = types.Int64Null()
	}
}
