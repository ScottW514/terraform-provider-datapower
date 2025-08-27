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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAPICGSProxyPolicy struct {
	ProxyPolicyEnable types.Bool   `tfsdk:"proxy_policy_enable"`
	RemoteAddress     types.String `tfsdk:"remote_address"`
	RemotePort        types.Int64  `tfsdk:"remote_port"`
}

var DmAPICGSProxyPolicyRemoteAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "proxy_policy_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmAPICGSProxyPolicyRemotePortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "proxy_policy_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmAPICGSProxyPolicyRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "proxy_policy_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmAPICGSProxyPolicyRemotePortIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "proxy_policy_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAPICGSProxyPolicyObjectType = map[string]attr.Type{
	"proxy_policy_enable": types.BoolType,
	"remote_address":      types.StringType,
	"remote_port":         types.Int64Type,
}
var DmAPICGSProxyPolicyObjectDefault = map[string]attr.Value{
	"proxy_policy_enable": types.BoolValue(false),
	"remote_address":      types.StringNull(),
	"remote_port":         types.Int64Null(),
}

func GetDmAPICGSProxyPolicyDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAPICGSProxyPolicyDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"proxy_policy_enable": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable a proxy to connect to API Manager.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"remote_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname or IP address of the proxy to connect to the API Manager endpoint.", "", "").String,
				Computed:            true,
			},
			"remote_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the proxy to connect to the API Manager endpoint.", "", "").String,
				Computed:            true,
			},
		},
	}
	DmAPICGSProxyPolicyDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAPICGSProxyPolicyDataSourceSchema
}
func GetDmAPICGSProxyPolicyResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAPICGSProxyPolicyResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAPICGSProxyPolicyObjectType,
				DmAPICGSProxyPolicyObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"proxy_policy_enable": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable a proxy to connect to API Manager.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"remote_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname or IP address of the proxy to connect to the API Manager endpoint.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAPICGSProxyPolicyRemoteAddressCondVal, validators.Evaluation{}, false),
				},
			},
			"remote_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the proxy to connect to the API Manager endpoint.", "", "").String,
				Required:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(DmAPICGSProxyPolicyRemotePortCondVal, validators.Evaluation{}, false),
				},
			},
		},
	}
	DmAPICGSProxyPolicyResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAPICGSProxyPolicyResourceSchema.Required = true
	} else {
		DmAPICGSProxyPolicyResourceSchema.Optional = true
		DmAPICGSProxyPolicyResourceSchema.Computed = true
	}
	return DmAPICGSProxyPolicyResourceSchema
}

func (data DmAPICGSProxyPolicy) IsNull() bool {
	if !data.ProxyPolicyEnable.IsNull() {
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

func (data DmAPICGSProxyPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ProxyPolicyEnable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProxyPolicyEnable`, tfutils.StringFromBool(data.ProxyPolicyEnable, ""))
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	return body
}

func (data *DmAPICGSProxyPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ProxyPolicyEnable`); value.Exists() {
		data.ProxyPolicyEnable = tfutils.BoolFromString(value.String())
	} else {
		data.ProxyPolicyEnable = types.BoolNull()
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

func (data *DmAPICGSProxyPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ProxyPolicyEnable`); value.Exists() && !data.ProxyPolicyEnable.IsNull() {
		data.ProxyPolicyEnable = tfutils.BoolFromString(value.String())
	} else if data.ProxyPolicyEnable.ValueBool() {
		data.ProxyPolicyEnable = types.BoolNull()
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
