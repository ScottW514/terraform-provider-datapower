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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FTPServerSourceProtocolHandler struct {
	Id                          types.String `tfsdk:"id"`
	AppDomain                   types.String `tfsdk:"app_domain"`
	UserSummary                 types.String `tfsdk:"user_summary"`
	LocalAddress                types.String `tfsdk:"local_address"`
	LocalPort                   types.Int64  `tfsdk:"local_port"`
	FilesystemType              types.String `tfsdk:"filesystem_type"`
	PersistentFilesystemTimeout types.Int64  `tfsdk:"persistent_filesystem_timeout"`
	VirtualDirectories          types.List   `tfsdk:"virtual_directories"`
	DefaultDirectory            types.String `tfsdk:"default_directory"`
	MaxFilenameLength           types.Int64  `tfsdk:"max_filename_length"`
	Acl                         types.String `tfsdk:"acl"`
	RequireTls                  types.String `tfsdk:"require_tls"`
	PasswordAaaPolicy           types.String `tfsdk:"password_aaa_policy"`
	CertificateAaaPolicy        types.String `tfsdk:"certificate_aaa_policy"`
	AllowCcc                    types.Bool   `tfsdk:"allow_ccc"`
	Passive                     types.String `tfsdk:"passive"`
	UsePasvPortRange            types.Bool   `tfsdk:"use_pasv_port_range"`
	PasvMinPort                 types.Int64  `tfsdk:"pasv_min_port"`
	PasvMaxPort                 types.Int64  `tfsdk:"pasv_max_port"`
	PasvIdleTimeOut             types.Int64  `tfsdk:"pasv_idle_time_out"`
	DisablePasvipCheck          types.Bool   `tfsdk:"disable_pasvip_check"`
	DisablePortipCheck          types.Bool   `tfsdk:"disable_portip_check"`
	UseAlternatePasvAddr        types.Bool   `tfsdk:"use_alternate_pasv_addr"`
	AlternatePasvAddr           types.String `tfsdk:"alternate_pasv_addr"`
	AllowListCmd                types.Bool   `tfsdk:"allow_list_cmd"`
	AllowDeleCmd                types.Bool   `tfsdk:"allow_dele_cmd"`
	DataEncryption              types.String `tfsdk:"data_encryption"`
	AllowCompression            types.Bool   `tfsdk:"allow_compression"`
	AllowStou                   types.Bool   `tfsdk:"allow_stou"`
	UniqueFilenamePrefix        types.String `tfsdk:"unique_filename_prefix"`
	AllowRest                   types.Bool   `tfsdk:"allow_rest"`
	RestartTimeout              types.Int64  `tfsdk:"restart_timeout"`
	IdleTimeout                 types.Int64  `tfsdk:"idle_timeout"`
	ResponseType                types.String `tfsdk:"response_type"`
	ResponseStorage             types.String `tfsdk:"response_storage"`
	TemporaryStorageSize        types.Int64  `tfsdk:"temporary_storage_size"`
	ResponseNfsMount            types.String `tfsdk:"response_nfs_mount"`
	ResponseSuffix              types.String `tfsdk:"response_suffix"`
	SslServerConfigType         types.String `tfsdk:"ssl_server_config_type"`
	SslServer                   types.String `tfsdk:"ssl_server"`
	SslsniServer                types.String `tfsdk:"sslsni_server"`
}

var FTPServerSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"local_address":                 types.StringType,
	"local_port":                    types.Int64Type,
	"filesystem_type":               types.StringType,
	"persistent_filesystem_timeout": types.Int64Type,
	"virtual_directories":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType}},
	"default_directory":             types.StringType,
	"max_filename_length":           types.Int64Type,
	"acl":                           types.StringType,
	"require_tls":                   types.StringType,
	"password_aaa_policy":           types.StringType,
	"certificate_aaa_policy":        types.StringType,
	"allow_ccc":                     types.BoolType,
	"passive":                       types.StringType,
	"use_pasv_port_range":           types.BoolType,
	"pasv_min_port":                 types.Int64Type,
	"pasv_max_port":                 types.Int64Type,
	"pasv_idle_time_out":            types.Int64Type,
	"disable_pasvip_check":          types.BoolType,
	"disable_portip_check":          types.BoolType,
	"use_alternate_pasv_addr":       types.BoolType,
	"alternate_pasv_addr":           types.StringType,
	"allow_list_cmd":                types.BoolType,
	"allow_dele_cmd":                types.BoolType,
	"data_encryption":               types.StringType,
	"allow_compression":             types.BoolType,
	"allow_stou":                    types.BoolType,
	"unique_filename_prefix":        types.StringType,
	"allow_rest":                    types.BoolType,
	"restart_timeout":               types.Int64Type,
	"idle_timeout":                  types.Int64Type,
	"response_type":                 types.StringType,
	"response_storage":              types.StringType,
	"temporary_storage_size":        types.Int64Type,
	"response_nfs_mount":            types.StringType,
	"response_suffix":               types.StringType,
	"ssl_server_config_type":        types.StringType,
	"ssl_server":                    types.StringType,
	"sslsni_server":                 types.StringType,
}

func (data FTPServerSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/FTPServerSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data FTPServerSourceProtocolHandler) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
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
	if !data.FilesystemType.IsNull() {
		return false
	}
	if !data.PersistentFilesystemTimeout.IsNull() {
		return false
	}
	if !data.VirtualDirectories.IsNull() {
		return false
	}
	if !data.DefaultDirectory.IsNull() {
		return false
	}
	if !data.MaxFilenameLength.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if !data.RequireTls.IsNull() {
		return false
	}
	if !data.PasswordAaaPolicy.IsNull() {
		return false
	}
	if !data.CertificateAaaPolicy.IsNull() {
		return false
	}
	if !data.AllowCcc.IsNull() {
		return false
	}
	if !data.Passive.IsNull() {
		return false
	}
	if !data.UsePasvPortRange.IsNull() {
		return false
	}
	if !data.PasvMinPort.IsNull() {
		return false
	}
	if !data.PasvMaxPort.IsNull() {
		return false
	}
	if !data.PasvIdleTimeOut.IsNull() {
		return false
	}
	if !data.DisablePasvipCheck.IsNull() {
		return false
	}
	if !data.DisablePortipCheck.IsNull() {
		return false
	}
	if !data.UseAlternatePasvAddr.IsNull() {
		return false
	}
	if !data.AlternatePasvAddr.IsNull() {
		return false
	}
	if !data.AllowListCmd.IsNull() {
		return false
	}
	if !data.AllowDeleCmd.IsNull() {
		return false
	}
	if !data.DataEncryption.IsNull() {
		return false
	}
	if !data.AllowCompression.IsNull() {
		return false
	}
	if !data.AllowStou.IsNull() {
		return false
	}
	if !data.UniqueFilenamePrefix.IsNull() {
		return false
	}
	if !data.AllowRest.IsNull() {
		return false
	}
	if !data.RestartTimeout.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.ResponseType.IsNull() {
		return false
	}
	if !data.ResponseStorage.IsNull() {
		return false
	}
	if !data.TemporaryStorageSize.IsNull() {
		return false
	}
	if !data.ResponseNfsMount.IsNull() {
		return false
	}
	if !data.ResponseSuffix.IsNull() {
		return false
	}
	if !data.SslServerConfigType.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.SslsniServer.IsNull() {
		return false
	}
	return true
}

func (data FTPServerSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.FilesystemType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FilesystemType`, data.FilesystemType.ValueString())
	}
	if !data.PersistentFilesystemTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentFilesystemTimeout`, data.PersistentFilesystemTimeout.ValueInt64())
	}
	if !data.VirtualDirectories.IsNull() {
		var values []DmFTPServerVirtualDirectory
		data.VirtualDirectories.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`VirtualDirectories`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.DefaultDirectory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultDirectory`, data.DefaultDirectory.ValueString())
	}
	if !data.MaxFilenameLength.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxFilenameLength`, data.MaxFilenameLength.ValueInt64())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.RequireTls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireTLS`, data.RequireTls.ValueString())
	}
	if !data.PasswordAaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAAAPolicy`, data.PasswordAaaPolicy.ValueString())
	}
	if !data.CertificateAaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CertificateAAAPolicy`, data.CertificateAaaPolicy.ValueString())
	}
	if !data.AllowCcc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCCC`, tfutils.StringFromBool(data.AllowCcc, ""))
	}
	if !data.Passive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Passive`, data.Passive.ValueString())
	}
	if !data.UsePasvPortRange.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UsePasvPortRange`, tfutils.StringFromBool(data.UsePasvPortRange, ""))
	}
	if !data.PasvMinPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasvMinPort`, data.PasvMinPort.ValueInt64())
	}
	if !data.PasvMaxPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasvMaxPort`, data.PasvMaxPort.ValueInt64())
	}
	if !data.PasvIdleTimeOut.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasvIdleTimeOut`, data.PasvIdleTimeOut.ValueInt64())
	}
	if !data.DisablePasvipCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisablePASVIPCheck`, tfutils.StringFromBool(data.DisablePasvipCheck, ""))
	}
	if !data.DisablePortipCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DisablePORTIPCheck`, tfutils.StringFromBool(data.DisablePortipCheck, ""))
	}
	if !data.UseAlternatePasvAddr.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseAlternatePASVAddr`, tfutils.StringFromBool(data.UseAlternatePasvAddr, ""))
	}
	if !data.AlternatePasvAddr.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AlternatePASVAddr`, data.AlternatePasvAddr.ValueString())
	}
	if !data.AllowListCmd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowLISTCmd`, tfutils.StringFromBool(data.AllowListCmd, ""))
	}
	if !data.AllowDeleCmd.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowDELECmd`, tfutils.StringFromBool(data.AllowDeleCmd, ""))
	}
	if !data.DataEncryption.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataEncryption`, data.DataEncryption.ValueString())
	}
	if !data.AllowCompression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowCompression`, tfutils.StringFromBool(data.AllowCompression, ""))
	}
	if !data.AllowStou.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowSTOU`, tfutils.StringFromBool(data.AllowStou, ""))
	}
	if !data.UniqueFilenamePrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UniqueFilenamePrefix`, data.UniqueFilenamePrefix.ValueString())
	}
	if !data.AllowRest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowREST`, tfutils.StringFromBool(data.AllowRest, ""))
	}
	if !data.RestartTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RestartTimeout`, data.RestartTimeout.ValueInt64())
	}
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.ResponseType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseType`, data.ResponseType.ValueString())
	}
	if !data.ResponseStorage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseStorage`, data.ResponseStorage.ValueString())
	}
	if !data.TemporaryStorageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TemporaryStorageSize`, data.TemporaryStorageSize.ValueInt64())
	}
	if !data.ResponseNfsMount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseNFSMount`, data.ResponseNfsMount.ValueString())
	}
	if !data.ResponseSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResponseSuffix`, data.ResponseSuffix.ValueString())
	}
	if !data.SslServerConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServerConfigType`, data.SslServerConfigType.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.SslsniServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLSNIServer`, data.SslsniServer.ValueString())
	}
	return body
}

func (data *FTPServerSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(21)
	}
	if value := res.Get(pathRoot + `FilesystemType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FilesystemType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FilesystemType = types.StringValue("virtual-ephemeral")
	}
	if value := res.Get(pathRoot + `PersistentFilesystemTimeout`); value.Exists() {
		data.PersistentFilesystemTimeout = types.Int64Value(value.Int())
	} else {
		data.PersistentFilesystemTimeout = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `VirtualDirectories`); value.Exists() {
		l := []DmFTPServerVirtualDirectory{}
		if value := res.Get(`VirtualDirectories`); value.Exists() {
			for _, v := range value.Array() {
				item := DmFTPServerVirtualDirectory{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.VirtualDirectories, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType}, l)
		} else {
			data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType})
		}
	} else {
		data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType})
	}
	if value := res.Get(pathRoot + `DefaultDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultDirectory = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `MaxFilenameLength`); value.Exists() {
		data.MaxFilenameLength = types.Int64Value(value.Int())
	} else {
		data.MaxFilenameLength = types.Int64Value(256)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequireTLS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequireTls = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequireTls = types.StringValue("off")
	}
	if value := res.Get(pathRoot + `PasswordAAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CertificateAAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CertificateAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CertificateAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowCCC`); value.Exists() {
		data.AllowCcc = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCcc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Passive = types.StringValue("allow")
	}
	if value := res.Get(pathRoot + `UsePasvPortRange`); value.Exists() {
		data.UsePasvPortRange = tfutils.BoolFromString(value.String())
	} else {
		data.UsePasvPortRange = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasvMinPort`); value.Exists() {
		data.PasvMinPort = types.Int64Value(value.Int())
	} else {
		data.PasvMinPort = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `PasvMaxPort`); value.Exists() {
		data.PasvMaxPort = types.Int64Value(value.Int())
	} else {
		data.PasvMaxPort = types.Int64Value(1050)
	}
	if value := res.Get(pathRoot + `PasvIdleTimeOut`); value.Exists() {
		data.PasvIdleTimeOut = types.Int64Value(value.Int())
	} else {
		data.PasvIdleTimeOut = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `DisablePASVIPCheck`); value.Exists() {
		data.DisablePasvipCheck = tfutils.BoolFromString(value.String())
	} else {
		data.DisablePasvipCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisablePORTIPCheck`); value.Exists() {
		data.DisablePortipCheck = tfutils.BoolFromString(value.String())
	} else {
		data.DisablePortipCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseAlternatePASVAddr`); value.Exists() {
		data.UseAlternatePasvAddr = tfutils.BoolFromString(value.String())
	} else {
		data.UseAlternatePasvAddr = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlternatePASVAddr`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AlternatePasvAddr = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AlternatePasvAddr = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowLISTCmd`); value.Exists() {
		data.AllowListCmd = tfutils.BoolFromString(value.String())
	} else {
		data.AllowListCmd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowDELECmd`); value.Exists() {
		data.AllowDeleCmd = tfutils.BoolFromString(value.String())
	} else {
		data.AllowDeleCmd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DataEncryption`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataEncryption = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataEncryption = types.StringValue("allow")
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowSTOU`); value.Exists() {
		data.AllowStou = tfutils.BoolFromString(value.String())
	} else {
		data.AllowStou = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UniqueFilenamePrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UniqueFilenamePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UniqueFilenamePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowREST`); value.Exists() {
		data.AllowRest = tfutils.BoolFromString(value.String())
	} else {
		data.AllowRest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RestartTimeout`); value.Exists() {
		data.RestartTimeout = types.Int64Value(value.Int())
	} else {
		data.RestartTimeout = types.Int64Value(240)
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseType = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `ResponseStorage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseStorage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseStorage = types.StringValue("temporary")
	}
	if value := res.Get(pathRoot + `TemporaryStorageSize`); value.Exists() {
		data.TemporaryStorageSize = types.Int64Value(value.Int())
	} else {
		data.TemporaryStorageSize = types.Int64Value(32)
	}
	if value := res.Get(pathRoot + `ResponseNFSMount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseNfsMount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseNfsMount = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResponseSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServerConfigType = types.StringValue("server")
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
}

func (data *FTPServerSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 21 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FilesystemType`); value.Exists() && !data.FilesystemType.IsNull() {
		data.FilesystemType = tfutils.ParseStringFromGJSON(value)
	} else if data.FilesystemType.ValueString() != "virtual-ephemeral" {
		data.FilesystemType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentFilesystemTimeout`); value.Exists() && !data.PersistentFilesystemTimeout.IsNull() {
		data.PersistentFilesystemTimeout = types.Int64Value(value.Int())
	} else if data.PersistentFilesystemTimeout.ValueInt64() != 600 {
		data.PersistentFilesystemTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `VirtualDirectories`); value.Exists() && !data.VirtualDirectories.IsNull() {
		l := []DmFTPServerVirtualDirectory{}
		for _, v := range value.Array() {
			item := DmFTPServerVirtualDirectory{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.VirtualDirectories, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType}, l)
		} else {
			data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType})
		}
	} else {
		data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmFTPServerVirtualDirectoryObjectType})
	}
	if value := res.Get(pathRoot + `DefaultDirectory`); value.Exists() && !data.DefaultDirectory.IsNull() {
		data.DefaultDirectory = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultDirectory.ValueString() != "/" {
		data.DefaultDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxFilenameLength`); value.Exists() && !data.MaxFilenameLength.IsNull() {
		data.MaxFilenameLength = types.Int64Value(value.Int())
	} else if data.MaxFilenameLength.ValueInt64() != 256 {
		data.MaxFilenameLength = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequireTLS`); value.Exists() && !data.RequireTls.IsNull() {
		data.RequireTls = tfutils.ParseStringFromGJSON(value)
	} else if data.RequireTls.ValueString() != "off" {
		data.RequireTls = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAAAPolicy`); value.Exists() && !data.PasswordAaaPolicy.IsNull() {
		data.PasswordAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CertificateAAAPolicy`); value.Exists() && !data.CertificateAaaPolicy.IsNull() {
		data.CertificateAaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CertificateAaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowCCC`); value.Exists() && !data.AllowCcc.IsNull() {
		data.AllowCcc = tfutils.BoolFromString(value.String())
	} else if !data.AllowCcc.ValueBool() {
		data.AllowCcc = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && !data.Passive.IsNull() {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else if data.Passive.ValueString() != "allow" {
		data.Passive = types.StringNull()
	}
	if value := res.Get(pathRoot + `UsePasvPortRange`); value.Exists() && !data.UsePasvPortRange.IsNull() {
		data.UsePasvPortRange = tfutils.BoolFromString(value.String())
	} else if data.UsePasvPortRange.ValueBool() {
		data.UsePasvPortRange = types.BoolNull()
	}
	if value := res.Get(pathRoot + `PasvMinPort`); value.Exists() && !data.PasvMinPort.IsNull() {
		data.PasvMinPort = types.Int64Value(value.Int())
	} else if data.PasvMinPort.ValueInt64() != 1024 {
		data.PasvMinPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PasvMaxPort`); value.Exists() && !data.PasvMaxPort.IsNull() {
		data.PasvMaxPort = types.Int64Value(value.Int())
	} else if data.PasvMaxPort.ValueInt64() != 1050 {
		data.PasvMaxPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PasvIdleTimeOut`); value.Exists() && !data.PasvIdleTimeOut.IsNull() {
		data.PasvIdleTimeOut = types.Int64Value(value.Int())
	} else if data.PasvIdleTimeOut.ValueInt64() != 60 {
		data.PasvIdleTimeOut = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DisablePASVIPCheck`); value.Exists() && !data.DisablePasvipCheck.IsNull() {
		data.DisablePasvipCheck = tfutils.BoolFromString(value.String())
	} else if data.DisablePasvipCheck.ValueBool() {
		data.DisablePasvipCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DisablePORTIPCheck`); value.Exists() && !data.DisablePortipCheck.IsNull() {
		data.DisablePortipCheck = tfutils.BoolFromString(value.String())
	} else if data.DisablePortipCheck.ValueBool() {
		data.DisablePortipCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseAlternatePASVAddr`); value.Exists() && !data.UseAlternatePasvAddr.IsNull() {
		data.UseAlternatePasvAddr = tfutils.BoolFromString(value.String())
	} else if data.UseAlternatePasvAddr.ValueBool() {
		data.UseAlternatePasvAddr = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AlternatePASVAddr`); value.Exists() && !data.AlternatePasvAddr.IsNull() {
		data.AlternatePasvAddr = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AlternatePasvAddr = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowLISTCmd`); value.Exists() && !data.AllowListCmd.IsNull() {
		data.AllowListCmd = tfutils.BoolFromString(value.String())
	} else if data.AllowListCmd.ValueBool() {
		data.AllowListCmd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowDELECmd`); value.Exists() && !data.AllowDeleCmd.IsNull() {
		data.AllowDeleCmd = tfutils.BoolFromString(value.String())
	} else if data.AllowDeleCmd.ValueBool() {
		data.AllowDeleCmd = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DataEncryption`); value.Exists() && !data.DataEncryption.IsNull() {
		data.DataEncryption = tfutils.ParseStringFromGJSON(value)
	} else if data.DataEncryption.ValueString() != "allow" {
		data.DataEncryption = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowCompression`); value.Exists() && !data.AllowCompression.IsNull() {
		data.AllowCompression = tfutils.BoolFromString(value.String())
	} else if !data.AllowCompression.ValueBool() {
		data.AllowCompression = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowSTOU`); value.Exists() && !data.AllowStou.IsNull() {
		data.AllowStou = tfutils.BoolFromString(value.String())
	} else if data.AllowStou.ValueBool() {
		data.AllowStou = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UniqueFilenamePrefix`); value.Exists() && !data.UniqueFilenamePrefix.IsNull() {
		data.UniqueFilenamePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UniqueFilenamePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `AllowREST`); value.Exists() && !data.AllowRest.IsNull() {
		data.AllowRest = tfutils.BoolFromString(value.String())
	} else if data.AllowRest.ValueBool() {
		data.AllowRest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RestartTimeout`); value.Exists() && !data.RestartTimeout.IsNull() {
		data.RestartTimeout = types.Int64Value(value.Int())
	} else if data.RestartTimeout.ValueInt64() != 240 {
		data.RestartTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 0 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ResponseType`); value.Exists() && !data.ResponseType.IsNull() {
		data.ResponseType = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseType.ValueString() != "none" {
		data.ResponseType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseStorage`); value.Exists() && !data.ResponseStorage.IsNull() {
		data.ResponseStorage = tfutils.ParseStringFromGJSON(value)
	} else if data.ResponseStorage.ValueString() != "temporary" {
		data.ResponseStorage = types.StringNull()
	}
	if value := res.Get(pathRoot + `TemporaryStorageSize`); value.Exists() && !data.TemporaryStorageSize.IsNull() {
		data.TemporaryStorageSize = types.Int64Value(value.Int())
	} else if data.TemporaryStorageSize.ValueInt64() != 32 {
		data.TemporaryStorageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ResponseNFSMount`); value.Exists() && !data.ResponseNfsMount.IsNull() {
		data.ResponseNfsMount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseNfsMount = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSuffix`); value.Exists() && !data.ResponseSuffix.IsNull() {
		data.ResponseSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResponseSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServerConfigType`); value.Exists() && !data.SslServerConfigType.IsNull() {
		data.SslServerConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslServerConfigType.ValueString() != "server" {
		data.SslServerConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLSNIServer`); value.Exists() && !data.SslsniServer.IsNull() {
		data.SslsniServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslsniServer = types.StringNull()
	}
}
