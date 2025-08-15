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

type DmSoapActionPolicy struct {
	RegExp     types.String `tfsdk:"reg_exp"`
	SoapAction types.String `tfsdk:"soap_action"`
}

var DmSoapActionPolicyObjectType = map[string]attr.Type{
	"reg_exp":     types.StringType,
	"soap_action": types.StringType,
}
var DmSoapActionPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":     types.StringNull(),
	"soap_action": types.StringNull(),
}
var DmSoapActionPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Computed:            true,
		},
		"soap_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the HTTP SOAPAction header.", "", "").String,
			Computed:            true,
		},
	},
}
var DmSoapActionPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Required:            true,
		},
		"soap_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value for the HTTP SOAPAction header.", "", "").String,
			Required:            true,
		},
	},
}

func (data DmSoapActionPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.SoapAction.IsNull() {
		return false
	}
	return true
}

func (data DmSoapActionPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.SoapAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SoapAction`, data.SoapAction.ValueString())
	}
	return body
}

func (data *DmSoapActionPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SoapAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapAction = types.StringNull()
	}
}

func (data *DmSoapActionPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SoapAction`); value.Exists() && !data.SoapAction.IsNull() {
		data.SoapAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapAction = types.StringNull()
	}
}
