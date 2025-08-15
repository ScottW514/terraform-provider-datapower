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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSSLPolicy struct {
	RegExp              types.String `tfsdk:"reg_exp"`
	SslClientConfigType types.String `tfsdk:"ssl_client_config_type"`
	SslClient           types.String `tfsdk:"ssl_client"`
}

var DmSSLPolicyObjectType = map[string]attr.Type{
	"reg_exp":                types.StringType,
	"ssl_client_config_type": types.StringType,
	"ssl_client":             types.StringType,
}
var DmSSLPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":                types.StringNull(),
	"ssl_client_config_type": types.StringValue("client"),
	"ssl_client":             types.StringNull(),
}
var DmSSLPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Computed:            true,
		},
		"ssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile to secure connections with targets", "", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
		},
		"ssl_client": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections with targets", "", "ssl_client_profile").String,
			Computed:            true,
		},
	},
}
var DmSSLPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").String,
			Required:            true,
		},
		"ssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile to secure connections with targets", "", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("client"),
			},
			Default: stringdefault.StaticString("client"),
		},
		"ssl_client": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections with targets", "", "ssl_client_profile").String,
			Optional:            true,
		},
	},
}

func (data DmSSLPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data DmSSLPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *DmSSLPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *DmSSLPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}
