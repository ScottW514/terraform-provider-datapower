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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSFTPServerVirtualDirectory struct {
	VirtualPath types.String `tfsdk:"virtual_path"`
}

var DmSFTPServerVirtualDirectoryObjectType = map[string]attr.Type{
	"virtual_path": types.StringType,
}
var DmSFTPServerVirtualDirectoryObjectDefault = map[string]attr.Value{
	"virtual_path": types.StringNull(),
}
var DmSFTPServerVirtualDirectoryDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"virtual_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory in the virtual file system of the SFTP server.", "", "").String,
			Computed:            true,
		},
	},
}
var DmSFTPServerVirtualDirectoryResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"virtual_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory in the virtual file system of the SFTP server.", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmSFTPServerVirtualDirectory) IsNull() bool {
	if !data.VirtualPath.IsNull() {
		return false
	}
	return true
}

func (data DmSFTPServerVirtualDirectory) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.VirtualPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VirtualPath`, data.VirtualPath.ValueString())
	}
	return body
}

func (data *DmSFTPServerVirtualDirectory) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `VirtualPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.VirtualPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VirtualPath = types.StringNull()
	}
}

func (data *DmSFTPServerVirtualDirectory) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `VirtualPath`); value.Exists() && !data.VirtualPath.IsNull() {
		data.VirtualPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.VirtualPath = types.StringNull()
	}
}
