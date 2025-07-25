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

type SSHServerSourceProtocolHandler struct {
	Id                          types.String                    `tfsdk:"id"`
	AppDomain                   types.String                    `tfsdk:"app_domain"`
	UserSummary                 types.String                    `tfsdk:"user_summary"`
	LocalAddress                types.String                    `tfsdk:"local_address"`
	LocalPort                   types.Int64                     `tfsdk:"local_port"`
	Acl                         types.String                    `tfsdk:"acl"`
	HostPrivateKeys             types.List                      `tfsdk:"host_private_keys"`
	SshUserAuthentication       *DmSSHUserAuthenticationMethods `tfsdk:"ssh_user_authentication"`
	AllowBackendListings        types.Bool                      `tfsdk:"allow_backend_listings"`
	AllowBackendDelete          types.Bool                      `tfsdk:"allow_backend_delete"`
	AllowBackendStat            types.Bool                      `tfsdk:"allow_backend_stat"`
	AllowBackendMkdir           types.Bool                      `tfsdk:"allow_backend_mkdir"`
	AllowBackendRmdir           types.Bool                      `tfsdk:"allow_backend_rmdir"`
	AllowBackendRename          types.Bool                      `tfsdk:"allow_backend_rename"`
	AaaPolicy                   types.String                    `tfsdk:"aaa_policy"`
	FilesystemType              types.String                    `tfsdk:"filesystem_type"`
	DefaultDirectory            types.String                    `tfsdk:"default_directory"`
	IdleTimeout                 types.Int64                     `tfsdk:"idle_timeout"`
	PersistentFilesystemTimeout types.Int64                     `tfsdk:"persistent_filesystem_timeout"`
	VirtualDirectories          types.List                      `tfsdk:"virtual_directories"`
}

var SSHServerSourceProtocolHandlerObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"local_address":                 types.StringType,
	"local_port":                    types.Int64Type,
	"acl":                           types.StringType,
	"host_private_keys":             types.ListType{ElemType: types.StringType},
	"ssh_user_authentication":       types.ObjectType{AttrTypes: DmSSHUserAuthenticationMethodsObjectType},
	"allow_backend_listings":        types.BoolType,
	"allow_backend_delete":          types.BoolType,
	"allow_backend_stat":            types.BoolType,
	"allow_backend_mkdir":           types.BoolType,
	"allow_backend_rmdir":           types.BoolType,
	"allow_backend_rename":          types.BoolType,
	"aaa_policy":                    types.StringType,
	"filesystem_type":               types.StringType,
	"default_directory":             types.StringType,
	"idle_timeout":                  types.Int64Type,
	"persistent_filesystem_timeout": types.Int64Type,
	"virtual_directories":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType}},
}

func (data SSHServerSourceProtocolHandler) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SSHServerSourceProtocolHandler"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SSHServerSourceProtocolHandler) IsNull() bool {
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
	if !data.Acl.IsNull() {
		return false
	}
	if !data.HostPrivateKeys.IsNull() {
		return false
	}
	if data.SshUserAuthentication != nil {
		if !data.SshUserAuthentication.IsNull() {
			return false
		}
	}
	if !data.AllowBackendListings.IsNull() {
		return false
	}
	if !data.AllowBackendDelete.IsNull() {
		return false
	}
	if !data.AllowBackendStat.IsNull() {
		return false
	}
	if !data.AllowBackendMkdir.IsNull() {
		return false
	}
	if !data.AllowBackendRmdir.IsNull() {
		return false
	}
	if !data.AllowBackendRename.IsNull() {
		return false
	}
	if !data.AaaPolicy.IsNull() {
		return false
	}
	if !data.FilesystemType.IsNull() {
		return false
	}
	if !data.DefaultDirectory.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.PersistentFilesystemTimeout.IsNull() {
		return false
	}
	if !data.VirtualDirectories.IsNull() {
		return false
	}
	return true
}

func (data SSHServerSourceProtocolHandler) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if !data.HostPrivateKeys.IsNull() {
		var values []string
		data.HostPrivateKeys.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`HostPrivateKeys`+".-1", map[string]string{"value": val})
		}
	}
	if data.SshUserAuthentication != nil {
		if !data.SshUserAuthentication.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`SSHUserAuthentication`, data.SshUserAuthentication.ToBody(ctx, ""))
		}
	}
	if !data.AllowBackendListings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendListings`, tfutils.StringFromBool(data.AllowBackendListings, ""))
	}
	if !data.AllowBackendDelete.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendDelete`, tfutils.StringFromBool(data.AllowBackendDelete, ""))
	}
	if !data.AllowBackendStat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendStat`, tfutils.StringFromBool(data.AllowBackendStat, ""))
	}
	if !data.AllowBackendMkdir.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendMkdir`, tfutils.StringFromBool(data.AllowBackendMkdir, ""))
	}
	if !data.AllowBackendRmdir.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendRmdir`, tfutils.StringFromBool(data.AllowBackendRmdir, ""))
	}
	if !data.AllowBackendRename.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowBackendRename`, tfutils.StringFromBool(data.AllowBackendRename, ""))
	}
	if !data.AaaPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAAPolicy`, data.AaaPolicy.ValueString())
	}
	if !data.FilesystemType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FilesystemType`, data.FilesystemType.ValueString())
	}
	if !data.DefaultDirectory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultDirectory`, data.DefaultDirectory.ValueString())
	}
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.PersistentFilesystemTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentFilesystemTimeout`, data.PersistentFilesystemTimeout.ValueInt64())
	}
	if !data.VirtualDirectories.IsNull() {
		var values []DmSFTPServerVirtualDirectory
		data.VirtualDirectories.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`VirtualDirectories`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *SSHServerSourceProtocolHandler) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.LocalPort = types.Int64Value(22)
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HostPrivateKeys`); value.Exists() {
		data.HostPrivateKeys = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.HostPrivateKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SSHUserAuthentication`); value.Exists() {
		data.SshUserAuthentication = &DmSSHUserAuthenticationMethods{}
		data.SshUserAuthentication.FromBody(ctx, "", value)
	} else {
		data.SshUserAuthentication = nil
	}
	if value := res.Get(pathRoot + `AllowBackendListings`); value.Exists() {
		data.AllowBackendListings = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendListings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendDelete`); value.Exists() {
		data.AllowBackendDelete = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendDelete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendStat`); value.Exists() {
		data.AllowBackendStat = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendStat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendMkdir`); value.Exists() {
		data.AllowBackendMkdir = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendMkdir = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendRmdir`); value.Exists() {
		data.AllowBackendRmdir = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendRmdir = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendRename`); value.Exists() {
		data.AllowBackendRename = tfutils.BoolFromString(value.String())
	} else {
		data.AllowBackendRename = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `FilesystemType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FilesystemType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FilesystemType = types.StringValue("transparent")
	}
	if value := res.Get(pathRoot + `DefaultDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultDirectory = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `PersistentFilesystemTimeout`); value.Exists() {
		data.PersistentFilesystemTimeout = types.Int64Value(value.Int())
	} else {
		data.PersistentFilesystemTimeout = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `VirtualDirectories`); value.Exists() {
		l := []DmSFTPServerVirtualDirectory{}
		if value := res.Get(`VirtualDirectories`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSFTPServerVirtualDirectory{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.VirtualDirectories, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType}, l)
		} else {
			data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType})
		}
	} else {
		data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType})
	}
}

func (data *SSHServerSourceProtocolHandler) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.LocalPort.ValueInt64() != 22 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `HostPrivateKeys`); value.Exists() && !data.HostPrivateKeys.IsNull() {
		data.HostPrivateKeys = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.HostPrivateKeys = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SSHUserAuthentication`); value.Exists() {
		data.SshUserAuthentication.UpdateFromBody(ctx, "", value)
	} else {
		data.SshUserAuthentication = nil
	}
	if value := res.Get(pathRoot + `AllowBackendListings`); value.Exists() && !data.AllowBackendListings.IsNull() {
		data.AllowBackendListings = tfutils.BoolFromString(value.String())
	} else if !data.AllowBackendListings.ValueBool() {
		data.AllowBackendListings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendDelete`); value.Exists() && !data.AllowBackendDelete.IsNull() {
		data.AllowBackendDelete = tfutils.BoolFromString(value.String())
	} else if data.AllowBackendDelete.ValueBool() {
		data.AllowBackendDelete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendStat`); value.Exists() && !data.AllowBackendStat.IsNull() {
		data.AllowBackendStat = tfutils.BoolFromString(value.String())
	} else if data.AllowBackendStat.ValueBool() {
		data.AllowBackendStat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendMkdir`); value.Exists() && !data.AllowBackendMkdir.IsNull() {
		data.AllowBackendMkdir = tfutils.BoolFromString(value.String())
	} else if data.AllowBackendMkdir.ValueBool() {
		data.AllowBackendMkdir = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendRmdir`); value.Exists() && !data.AllowBackendRmdir.IsNull() {
		data.AllowBackendRmdir = tfutils.BoolFromString(value.String())
	} else if data.AllowBackendRmdir.ValueBool() {
		data.AllowBackendRmdir = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowBackendRename`); value.Exists() && !data.AllowBackendRename.IsNull() {
		data.AllowBackendRename = tfutils.BoolFromString(value.String())
	} else if data.AllowBackendRename.ValueBool() {
		data.AllowBackendRename = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AAAPolicy`); value.Exists() && !data.AaaPolicy.IsNull() {
		data.AaaPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AaaPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `FilesystemType`); value.Exists() && !data.FilesystemType.IsNull() {
		data.FilesystemType = tfutils.ParseStringFromGJSON(value)
	} else if data.FilesystemType.ValueString() != "transparent" {
		data.FilesystemType = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultDirectory`); value.Exists() && !data.DefaultDirectory.IsNull() {
		data.DefaultDirectory = tfutils.ParseStringFromGJSON(value)
	} else if data.DefaultDirectory.ValueString() != "/" {
		data.DefaultDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 0 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PersistentFilesystemTimeout`); value.Exists() && !data.PersistentFilesystemTimeout.IsNull() {
		data.PersistentFilesystemTimeout = types.Int64Value(value.Int())
	} else if data.PersistentFilesystemTimeout.ValueInt64() != 600 {
		data.PersistentFilesystemTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `VirtualDirectories`); value.Exists() && !data.VirtualDirectories.IsNull() {
		l := []DmSFTPServerVirtualDirectory{}
		for _, v := range value.Array() {
			item := DmSFTPServerVirtualDirectory{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.VirtualDirectories, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType}, l)
		} else {
			data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType})
		}
	} else {
		data.VirtualDirectories = types.ListNull(types.ObjectType{AttrTypes: DmSFTPServerVirtualDirectoryObjectType})
	}
}
