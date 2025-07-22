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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmTAMFedDir struct {
	FedName                 types.String `tfsdk:"fed_name"`
	Suffix                  types.String `tfsdk:"suffix"`
	Host                    types.String `tfsdk:"host"`
	Port                    types.Int64  `tfsdk:"port"`
	BindDn                  types.String `tfsdk:"bind_dn"`
	BindPw                  types.String `tfsdk:"bind_pw"`
	UseSsl                  types.Bool   `tfsdk:"use_ssl"`
	KeyFileLabel            types.String `tfsdk:"key_file_label"`
	BasicPrincipalAttribute types.String `tfsdk:"basic_principal_attribute"`
}

var DmTAMFedDirObjectType = map[string]attr.Type{
	"fed_name":                  types.StringType,
	"suffix":                    types.StringType,
	"host":                      types.StringType,
	"port":                      types.Int64Type,
	"bind_dn":                   types.StringType,
	"bind_pw":                   types.StringType,
	"use_ssl":                   types.BoolType,
	"key_file_label":            types.StringType,
	"basic_principal_attribute": types.StringType,
}
var DmTAMFedDirObjectDefault = map[string]attr.Value{
	"fed_name":                  types.StringNull(),
	"suffix":                    types.StringNull(),
	"host":                      types.StringNull(),
	"port":                      types.Int64Value(389),
	"bind_dn":                   types.StringNull(),
	"bind_pw":                   types.StringNull(),
	"use_ssl":                   types.BoolValue(false),
	"key_file_label":            types.StringNull(),
	"basic_principal_attribute": types.StringValue("uid"),
}
var DmTAMFedDirDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"fed_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "federate-name", "").String,
			Computed:            true,
		},
		"suffix": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Suffix", "suffix", "").String,
			Computed:            true,
		},
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP host", "hostname", "").String,
			Computed:            true,
		},
		"port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP port", "port", "").AddDefaultValue("389").String,
			Computed:            true,
		},
		"bind_dn": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP administrator DN", "bind-dn", "").String,
			Computed:            true,
		},
		"bind_pw": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP administrator password alias", "bind-pw", "passwordalias").String,
			Computed:            true,
		},
		"use_ssl": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use TLS", "UseSSL", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"key_file_label": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP server TLS key file label", "key-file-label", "").String,
			Computed:            true,
		},
		"basic_principal_attribute": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Basic user principal attribute", "basic-principal-attribute", "").AddDefaultValue("uid").String,
			Computed:            true,
		},
	},
}
var DmTAMFedDirResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"fed_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Name", "federate-name", "").String,
			Required:            true,
		},
		"suffix": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Suffix", "suffix", "").String,
			Required:            true,
		},
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP host", "hostname", "").String,
			Required:            true,
		},
		"port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP port", "port", "").AddDefaultValue("389").String,
			Computed:            true,
			Optional:            true,
			Default:             int64default.StaticInt64(389),
		},
		"bind_dn": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP administrator DN", "bind-dn", "").String,
			Required:            true,
		},
		"bind_pw": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP administrator password alias", "bind-pw", "passwordalias").String,
			Required:            true,
		},
		"use_ssl": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Use TLS", "UseSSL", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"key_file_label": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("LDAP server TLS key file label", "key-file-label", "").String,
			Optional:            true,
		},
		"basic_principal_attribute": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Basic user principal attribute", "basic-principal-attribute", "").AddDefaultValue("uid").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("uid"),
		},
	},
}

func (data DmTAMFedDir) IsNull() bool {
	if !data.FedName.IsNull() {
		return false
	}
	if !data.Suffix.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.BindDn.IsNull() {
		return false
	}
	if !data.BindPw.IsNull() {
		return false
	}
	if !data.UseSsl.IsNull() {
		return false
	}
	if !data.KeyFileLabel.IsNull() {
		return false
	}
	if !data.BasicPrincipalAttribute.IsNull() {
		return false
	}
	return true
}

func (data DmTAMFedDir) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.FedName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FedName`, data.FedName.ValueString())
	}
	if !data.Suffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Suffix`, data.Suffix.ValueString())
	}
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.BindDn.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BindDN`, data.BindDn.ValueString())
	}
	if !data.BindPw.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BindPw`, data.BindPw.ValueString())
	}
	if !data.UseSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseSSL`, tfutils.StringFromBool(data.UseSsl, false))
	}
	if !data.KeyFileLabel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeyFileLabel`, data.KeyFileLabel.ValueString())
	}
	if !data.BasicPrincipalAttribute.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BasicPrincipalAttribute`, data.BasicPrincipalAttribute.ValueString())
	}
	return body
}

func (data *DmTAMFedDir) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FedName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FedName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FedName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Suffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Suffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Suffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(389)
	}
	if value := res.Get(pathRoot + `BindDN`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindPw`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BindPw = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindPw = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSSL`); value.Exists() {
		data.UseSsl = tfutils.BoolFromString(value.String())
	} else {
		data.UseSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeyFileLabel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KeyFileLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyFileLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `BasicPrincipalAttribute`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BasicPrincipalAttribute = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BasicPrincipalAttribute = types.StringValue("uid")
	}
}

func (data *DmTAMFedDir) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FedName`); value.Exists() && !data.FedName.IsNull() {
		data.FedName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FedName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Suffix`); value.Exists() && !data.Suffix.IsNull() {
		data.Suffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Suffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 389 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `BindDN`); value.Exists() && !data.BindDn.IsNull() {
		data.BindDn = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindDn = types.StringNull()
	}
	if value := res.Get(pathRoot + `BindPw`); value.Exists() && !data.BindPw.IsNull() {
		data.BindPw = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BindPw = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseSSL`); value.Exists() && !data.UseSsl.IsNull() {
		data.UseSsl = tfutils.BoolFromString(value.String())
	} else if data.UseSsl.ValueBool() {
		data.UseSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeyFileLabel`); value.Exists() && !data.KeyFileLabel.IsNull() {
		data.KeyFileLabel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeyFileLabel = types.StringNull()
	}
	if value := res.Get(pathRoot + `BasicPrincipalAttribute`); value.Exists() && !data.BasicPrincipalAttribute.IsNull() {
		data.BasicPrincipalAttribute = tfutils.ParseStringFromGJSON(value)
	} else if data.BasicPrincipalAttribute.ValueString() != "uid" {
		data.BasicPrincipalAttribute = types.StringNull()
	}
}
