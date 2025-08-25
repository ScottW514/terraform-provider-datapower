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

type DmGitOpsTemplatePolicy struct {
	RemotePolicy types.String `tfsdk:"remote_policy"`
}

var DmGitOpsTemplatePolicyObjectType = map[string]attr.Type{
	"remote_policy": types.StringType,
}
var DmGitOpsTemplatePolicyObjectDefault = map[string]attr.Value{
	"remote_policy": types.StringNull(),
}

func GetDmGitOpsTemplatePolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmGitOpsTemplatePolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"remote_policy": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmGitOpsTemplatePolicyDataSourceSchema
}
func GetDmGitOpsTemplatePolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmGitOpsTemplatePolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"remote_policy": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmGitOpsTemplatePolicyResourceSchema
}

func (data DmGitOpsTemplatePolicy) IsNull() bool {
	if !data.RemotePolicy.IsNull() {
		return false
	}
	return true
}

func (data DmGitOpsTemplatePolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.RemotePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePolicy`, data.RemotePolicy.ValueString())
	}
	return body
}

func (data *DmGitOpsTemplatePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RemotePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemotePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemotePolicy = types.StringNull()
	}
}

func (data *DmGitOpsTemplatePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RemotePolicy`); value.Exists() && !data.RemotePolicy.IsNull() {
		data.RemotePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemotePolicy = types.StringNull()
	}
}
