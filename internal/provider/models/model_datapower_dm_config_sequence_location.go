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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmConfigSequenceLocation struct {
	Directory         types.String `tfsdk:"directory"`
	AccessProfileName types.String `tfsdk:"access_profile_name"`
}

var DmConfigSequenceLocationObjectType = map[string]attr.Type{
	"directory":           types.StringType,
	"access_profile_name": types.StringType,
}
var DmConfigSequenceLocationObjectDefault = map[string]attr.Value{
	"directory":           types.StringValue("local:///"),
	"access_profile_name": types.StringNull(),
}

func GetDmConfigSequenceLocationDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmConfigSequenceLocationDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"directory": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory to watch for changes to configuration files. Processing scans the directory for configuration files that match a given PCRE pattern and runs only the matching files. By default, this directory is also where the generated output files are stored.", "directory", "").AddDefaultValue("local:///").String,
				Computed:            true,
			},
			"access_profile_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the access policies to enforce. Each policy defines permissions to a specific resource group.", "access-profile", "access_profile").String,
				Computed:            true,
			},
		},
	}
	return DmConfigSequenceLocationDataSourceSchema
}
func GetDmConfigSequenceLocationResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmConfigSequenceLocationResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"directory": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory to watch for changes to configuration files. Processing scans the directory for configuration files that match a given PCRE pattern and runs only the matching files. By default, this directory is also where the generated output files are stored.", "directory", "").AddDefaultValue("local:///").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("local:///"),
			},
			"access_profile_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the access policies to enforce. Each policy defines permissions to a specific resource group.", "access-profile", "access_profile").String,
				Optional:            true,
			},
		},
	}
	return DmConfigSequenceLocationResourceSchema
}

func (data DmConfigSequenceLocation) IsNull() bool {
	if !data.Directory.IsNull() {
		return false
	}
	if !data.AccessProfileName.IsNull() {
		return false
	}
	return true
}

func (data DmConfigSequenceLocation) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Directory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Directory`, data.Directory.ValueString())
	}
	if !data.AccessProfileName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccessProfileName`, data.AccessProfileName.ValueString())
	}
	return body
}

func (data *DmConfigSequenceLocation) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Directory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Directory = types.StringValue("local:///")
	}
	if value := res.Get(pathRoot + `AccessProfileName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AccessProfileName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccessProfileName = types.StringNull()
	}
}

func (data *DmConfigSequenceLocation) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Directory`); value.Exists() && !data.Directory.IsNull() {
		data.Directory = tfutils.ParseStringFromGJSON(value)
	} else if data.Directory.ValueString() != "local:///" {
		data.Directory = types.StringNull()
	}
	if value := res.Get(pathRoot + `AccessProfileName`); value.Exists() && !data.AccessProfileName.IsNull() {
		data.AccessProfileName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AccessProfileName = types.StringNull()
	}
}
