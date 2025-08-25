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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmDomainFileMap struct {
	CopyFrom types.Bool `tfsdk:"copy_from"`
	CopyTo   types.Bool `tfsdk:"copy_to"`
	Delete   types.Bool `tfsdk:"delete"`
	Display  types.Bool `tfsdk:"display"`
	Exec     types.Bool `tfsdk:"exec"`
	Subdir   types.Bool `tfsdk:"subdir"`
}

var DmDomainFileMapObjectType = map[string]attr.Type{
	"copy_from": types.BoolType,
	"copy_to":   types.BoolType,
	"delete":    types.BoolType,
	"display":   types.BoolType,
	"exec":      types.BoolType,
	"subdir":    types.BoolType,
}
var DmDomainFileMapObjectDefault = map[string]attr.Value{
	"copy_from": types.BoolValue(true),
	"copy_to":   types.BoolValue(true),
	"delete":    types.BoolValue(true),
	"display":   types.BoolValue(true),
	"exec":      types.BoolValue(true),
	"subdir":    types.BoolValue(true),
}

func GetDmDomainFileMapDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmDomainFileMapDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"copy_from": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be copied from", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"copy_to": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be copied to", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"delete": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be deleted", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"display": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow file content to be displayed", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"exec": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be run as scripts", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
			"subdir": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow subdirectories to be created", "", "").AddDefaultValue("true").String,
				Computed:            true,
			},
		},
	}
	DmDomainFileMapDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmDomainFileMapDataSourceSchema
}
func GetDmDomainFileMapResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmDomainFileMapResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmDomainFileMapObjectType,
				DmDomainFileMapObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"copy_from": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be copied from", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"copy_to": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be copied to", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"delete": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be deleted", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"display": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow file content to be displayed", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"exec": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow files to be run as scripts", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"subdir": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow subdirectories to be created", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmDomainFileMapResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmDomainFileMapResourceSchema.Required = true
	} else {
		DmDomainFileMapResourceSchema.Optional = true
		DmDomainFileMapResourceSchema.Computed = true
	}
	return DmDomainFileMapResourceSchema
}

func (data DmDomainFileMap) IsNull() bool {
	if !data.CopyFrom.IsNull() {
		return false
	}
	if !data.CopyTo.IsNull() {
		return false
	}
	if !data.Delete.IsNull() {
		return false
	}
	if !data.Display.IsNull() {
		return false
	}
	if !data.Exec.IsNull() {
		return false
	}
	if !data.Subdir.IsNull() {
		return false
	}
	return true
}

func (data DmDomainFileMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.CopyFrom.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CopyFrom`, tfutils.StringFromBool(data.CopyFrom, ""))
	}
	if !data.CopyTo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CopyTo`, tfutils.StringFromBool(data.CopyTo, ""))
	}
	if !data.Delete.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Delete`, tfutils.StringFromBool(data.Delete, ""))
	}
	if !data.Display.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Display`, tfutils.StringFromBool(data.Display, ""))
	}
	if !data.Exec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Exec`, tfutils.StringFromBool(data.Exec, ""))
	}
	if !data.Subdir.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subdir`, tfutils.StringFromBool(data.Subdir, ""))
	}
	return body
}

func (data *DmDomainFileMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CopyFrom`); value.Exists() {
		data.CopyFrom = tfutils.BoolFromString(value.String())
	} else {
		data.CopyFrom = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CopyTo`); value.Exists() {
		data.CopyTo = tfutils.BoolFromString(value.String())
	} else {
		data.CopyTo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Delete`); value.Exists() {
		data.Delete = tfutils.BoolFromString(value.String())
	} else {
		data.Delete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Display`); value.Exists() {
		data.Display = tfutils.BoolFromString(value.String())
	} else {
		data.Display = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Exec`); value.Exists() {
		data.Exec = tfutils.BoolFromString(value.String())
	} else {
		data.Exec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Subdir`); value.Exists() {
		data.Subdir = tfutils.BoolFromString(value.String())
	} else {
		data.Subdir = types.BoolNull()
	}
}

func (data *DmDomainFileMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `CopyFrom`); value.Exists() && !data.CopyFrom.IsNull() {
		data.CopyFrom = tfutils.BoolFromString(value.String())
	} else if !data.CopyFrom.ValueBool() {
		data.CopyFrom = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CopyTo`); value.Exists() && !data.CopyTo.IsNull() {
		data.CopyTo = tfutils.BoolFromString(value.String())
	} else if !data.CopyTo.ValueBool() {
		data.CopyTo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Delete`); value.Exists() && !data.Delete.IsNull() {
		data.Delete = tfutils.BoolFromString(value.String())
	} else if !data.Delete.ValueBool() {
		data.Delete = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Display`); value.Exists() && !data.Display.IsNull() {
		data.Display = tfutils.BoolFromString(value.String())
	} else if !data.Display.ValueBool() {
		data.Display = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Exec`); value.Exists() && !data.Exec.IsNull() {
		data.Exec = tfutils.BoolFromString(value.String())
	} else if !data.Exec.ValueBool() {
		data.Exec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Subdir`); value.Exists() && !data.Subdir.IsNull() {
		data.Subdir = tfutils.BoolFromString(value.String())
	} else if !data.Subdir.ValueBool() {
		data.Subdir = types.BoolNull()
	}
}
