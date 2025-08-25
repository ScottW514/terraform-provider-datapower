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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSFTPPolicy struct {
	RegExp             types.String `tfsdk:"reg_exp"`
	SshClientProfile   types.String `tfsdk:"ssh_client_profile"`
	UseUniqueFilenames types.Bool   `tfsdk:"use_unique_filenames"`
}

var DmSFTPPolicyObjectType = map[string]attr.Type{
	"reg_exp":              types.StringType,
	"ssh_client_profile":   types.StringType,
	"use_unique_filenames": types.BoolType,
}
var DmSFTPPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":              types.StringNull(),
	"ssh_client_profile":   types.StringNull(),
	"use_unique_filenames": types.BoolValue(false),
}

func GetDmSFTPPolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSFTPPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"reg_exp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Computed:            true,
			},
			"ssh_client_profile": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH client profile.", "", "ssh_client_profile").String,
				Computed:            true,
			},
			"use_unique_filenames": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies whether to generate a unique file name for files written to an SSH FTP server. When the target URL represents a directory, a unique file name is generated. When the target URL represents a file name that already exists, any transfers result into a new unique file being created and its name is modified to include a unique prolog.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	return DmSFTPPolicyDataSourceSchema
}
func GetDmSFTPPolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSFTPPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"reg_exp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Required:            true,
			},
			"ssh_client_profile": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH client profile.", "", "ssh_client_profile").String,
				Required:            true,
			},
			"use_unique_filenames": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies whether to generate a unique file name for files written to an SSH FTP server. When the target URL represents a directory, a unique file name is generated. When the target URL represents a file name that already exists, any transfers result into a new unique file being created and its name is modified to include a unique prolog.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	return DmSFTPPolicyResourceSchema
}

func (data DmSFTPPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.SshClientProfile.IsNull() {
		return false
	}
	if !data.UseUniqueFilenames.IsNull() {
		return false
	}
	return true
}

func (data DmSFTPPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.SshClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSHClientProfile`, data.SshClientProfile.ValueString())
	}
	if !data.UseUniqueFilenames.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseUniqueFilenames`, tfutils.StringFromBool(data.UseUniqueFilenames, ""))
	}
	return body
}

func (data *DmSFTPPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SshClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseUniqueFilenames`); value.Exists() {
		data.UseUniqueFilenames = tfutils.BoolFromString(value.String())
	} else {
		data.UseUniqueFilenames = types.BoolNull()
	}
}

func (data *DmSFTPPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHClientProfile`); value.Exists() && !data.SshClientProfile.IsNull() {
		data.SshClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseUniqueFilenames`); value.Exists() && !data.UseUniqueFilenames.IsNull() {
		data.UseUniqueFilenames = tfutils.BoolFromString(value.String())
	} else if data.UseUniqueFilenames.ValueBool() {
		data.UseUniqueFilenames = types.BoolNull()
	}
}
