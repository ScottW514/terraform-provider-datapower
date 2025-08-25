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

type DmFTPQuotedCommand struct {
	QuotedCommand types.String `tfsdk:"quoted_command"`
}

var DmFTPQuotedCommandObjectType = map[string]attr.Type{
	"quoted_command": types.StringType,
}
var DmFTPQuotedCommandObjectDefault = map[string]attr.Value{
	"quoted_command": types.StringNull(),
}

func GetDmFTPQuotedCommandDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmFTPQuotedCommandDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"quoted_command": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The command to send to the remote FTP server. The command must not be one that would create an FTP data connection. Normally this will be a SITE command. This command must return a result code in the 200 series.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmFTPQuotedCommandDataSourceSchema
}
func GetDmFTPQuotedCommandResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmFTPQuotedCommandResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"quoted_command": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The command to send to the remote FTP server. The command must not be one that would create an FTP data connection. Normally this will be a SITE command. This command must return a result code in the 200 series.", "", "").String,
				Required:            true,
			},
		},
	}
	return DmFTPQuotedCommandResourceSchema
}

func (data DmFTPQuotedCommand) IsNull() bool {
	if !data.QuotedCommand.IsNull() {
		return false
	}
	return true
}

func (data DmFTPQuotedCommand) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.QuotedCommand.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QuotedCommand`, data.QuotedCommand.ValueString())
	}
	return body
}

func (data *DmFTPQuotedCommand) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `QuotedCommand`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QuotedCommand = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommand = types.StringNull()
	}
}

func (data *DmFTPQuotedCommand) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `QuotedCommand`); value.Exists() && !data.QuotedCommand.IsNull() {
		data.QuotedCommand = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommand = types.StringNull()
	}
}
