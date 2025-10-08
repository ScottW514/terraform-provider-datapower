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
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

type DmAPIProxyPolicy struct {
	RegExp        types.String `tfsdk:"reg_exp"`
	Skip          types.Bool   `tfsdk:"skip"`
	RemoteAddress types.String `tfsdk:"remote_address"`
	RemotePort    types.Int64  `tfsdk:"remote_port"`
	UserName      types.String `tfsdk:"user_name"`
	Password      types.String `tfsdk:"password"`
}

var DmAPIProxyPolicyRemoteAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAPIProxyPolicyRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAPIProxyPolicyRemotePortCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAPIProxyPolicyRemotePortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAPIProxyPolicyUserNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAPIProxyPolicyPasswordIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "skip",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmAPIProxyPolicyObjectType = map[string]attr.Type{
	"reg_exp":        types.StringType,
	"skip":           types.BoolType,
	"remote_address": types.StringType,
	"remote_port":    types.Int64Type,
	"user_name":      types.StringType,
	"password":       types.StringType,
}
var DmAPIProxyPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":        types.StringNull(),
	"skip":           types.BoolValue(false),
	"remote_address": types.StringNull(),
	"remote_port":    types.Int64Null(),
	"user_name":      types.StringNull(),
	"password":       types.StringNull(),
}

func GetDmAPIProxyPolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmAPIProxyPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"reg_exp": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the shell-style match pattern that defines a URL set. The URL set is assigned to a specific HTTP proxy.",
				Computed:            true,
			},
			"skip": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Specify how to treat the URL set for the match pattern. When set to on, the URL set is not forwarded to an HTTP proxy, and the remote host and remote port of a proxy are not defined. When set to off, the URL set is forwarded to the HTTP proxy designated by the remote host and remote port.",
				Computed:            true,
			},
			"remote_address": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the hostname or IP address of an HTTP server. With the remote port, this setting designates the HTTP proxy that services the URL set for the match pattern. When Skip is on, the remote host is not used.",
				Computed:            true,
			},
			"remote_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Specify the port on the HTTP server. With the remote host, this setting designates the HTTP proxy that services the URL set for the match pattern. When Skip is on, the remote port is not used.",
				Computed:            true,
			},
			"user_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the username for authentication.",
				Computed:            true,
			},
			"password": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the password alias for authentication.",
				Computed:            true,
			},
		},
	}
	return DmAPIProxyPolicyDataSourceSchema
}
func GetDmAPIProxyPolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmAPIProxyPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"reg_exp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style match pattern that defines a URL set. The URL set is assigned to a specific HTTP proxy.", "", "").String,
				Required:            true,
			},
			"skip": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to treat the URL set for the match pattern. When set to on, the URL set is not forwarded to an HTTP proxy, and the remote host and remote port of a proxy are not defined. When set to off, the URL set is forwarded to the HTTP proxy designated by the remote host and remote port.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"remote_address": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hostname or IP address of an HTTP server. With the remote port, this setting designates the HTTP proxy that services the URL set for the match pattern. When Skip is on, the remote host is not used.", "", "").AddRequiredWhen(DmAPIProxyPolicyRemoteAddressCondVal.String()).AddNotValidWhen(DmAPIProxyPolicyRemoteAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAPIProxyPolicyRemoteAddressCondVal, DmAPIProxyPolicyRemoteAddressIgnoreVal, false),
				},
			},
			"remote_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the port on the HTTP server. With the remote host, this setting designates the HTTP proxy that services the URL set for the match pattern. When Skip is on, the remote port is not used.", "", "").AddIntegerRange(1, 65535).AddRequiredWhen(DmAPIProxyPolicyRemotePortCondVal.String()).AddNotValidWhen(DmAPIProxyPolicyRemotePortIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(DmAPIProxyPolicyRemotePortCondVal, DmAPIProxyPolicyRemotePortIgnoreVal, false),
				},
			},
			"user_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the username for authentication.", "", "").AddNotValidWhen(DmAPIProxyPolicyUserNameIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^ ]+$"), "Must match :"+"^[^ ]+$"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmAPIProxyPolicyUserNameIgnoreVal, false),
				},
			},
			"password": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias for authentication.", "", "password_alias").AddNotValidWhen(DmAPIProxyPolicyPasswordIgnoreVal.String()).String,
				Optional:            true,
			},
		},
	}
	return DmAPIProxyPolicyResourceSchema
}

func (data DmAPIProxyPolicy) IsNull() bool {
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
	if !data.UserName.IsNull() {
		return false
	}
	if !data.Password.IsNull() {
		return false
	}
	return true
}

func (data DmAPIProxyPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Password`, data.Password.ValueString())
	}
	return body
}

func (data *DmAPIProxyPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
}

func (data *DmAPIProxyPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Password`); value.Exists() && !data.Password.IsNull() {
		data.Password = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Password = types.StringNull()
	}
}
