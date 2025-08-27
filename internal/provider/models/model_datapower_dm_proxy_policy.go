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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmProxyPolicy struct {
	RegExp        types.String `tfsdk:"reg_exp"`
	Skip          types.Bool   `tfsdk:"skip"`
	RemoteAddress types.String `tfsdk:"remote_address"`
	RemotePort    types.Int64  `tfsdk:"remote_port"`
}

var DmProxyPolicyRemoteAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmProxyPolicyRemotePortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmProxyPolicyRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmProxyPolicyRemotePortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmProxyPolicyObjectType = map[string]attr.Type{
	"reg_exp":        types.StringType,
	"skip":           types.BoolType,
	"remote_address": types.StringType,
	"remote_port":    types.Int64Type,
}
var DmProxyPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":        types.StringNull(),
	"skip":           types.BoolValue(false),
	"remote_address": types.StringNull(),
	"remote_port":    types.Int64Null(),
}

func GetDmProxyPolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmProxyPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"reg_exp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Computed:            true,
			},
			"skip": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to forward requests to the remote HTTP server. When not enabled, specify the remote host and port of the HTTP server.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"remote_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname or IP address of the remote HTTP server.", "", "").String,
				Computed:            true,
			},
			"remote_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port on the remote HTTP server.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmProxyPolicyDataSourceSchema
}
func GetDmProxyPolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmProxyPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"reg_exp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
				Required:            true,
			},
			"skip": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to forward requests to the remote HTTP server. When not enabled, specify the remote host and port of the HTTP server.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"remote_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname or IP address of the remote HTTP server.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmProxyPolicyRemoteAddressCondVal, validators.Evaluation{}, false),
				},
			},
			"remote_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port on the remote HTTP server.", "", "").String,
				Required:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmProxyPolicyRemotePortCondVal, validators.Evaluation{}, false),
				},
			},
		},
	}
	return DmProxyPolicyResourceSchema
}

func (data DmProxyPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.Skip.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	return true
}

func (data DmProxyPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.Skip.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Skip`, tfutils.StringFromBool(data.Skip, ""))
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	return body
}

func (data *DmProxyPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Skip`); value.Exists() {
		data.Skip = tfutils.BoolFromString(value.String())
	} else {
		data.Skip = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
}

func (data *DmProxyPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Skip`); value.Exists() && !data.Skip.IsNull() {
		data.Skip = tfutils.BoolFromString(value.String())
	} else if data.Skip.ValueBool() {
		data.Skip = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
}
