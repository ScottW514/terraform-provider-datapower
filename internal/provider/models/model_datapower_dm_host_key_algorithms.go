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

type DmHostKeyAlgorithms struct {
	SshEd25519        types.Bool `tfsdk:"ssh_ed25519"`
	EcdsaSha2Nistp256 types.Bool `tfsdk:"ecdsa_sha2_nistp256"`
	EcdsaSha2Nistp384 types.Bool `tfsdk:"ecdsa_sha2_nistp384"`
	EcdsaSha2Nistp521 types.Bool `tfsdk:"ecdsa_sha2_nistp521"`
	RsaSha2512        types.Bool `tfsdk:"rsa_sha2_512"`
	RsaSha2256        types.Bool `tfsdk:"rsa_sha2_256"`
}

var DmHostKeyAlgorithmsObjectType = map[string]attr.Type{
	"ssh_ed25519":         types.BoolType,
	"ecdsa_sha2_nistp256": types.BoolType,
	"ecdsa_sha2_nistp384": types.BoolType,
	"ecdsa_sha2_nistp521": types.BoolType,
	"rsa_sha2_512":        types.BoolType,
	"rsa_sha2_256":        types.BoolType,
}
var DmHostKeyAlgorithmsObjectDefault = map[string]attr.Value{
	"ssh_ed25519":         types.BoolValue(false),
	"ecdsa_sha2_nistp256": types.BoolValue(false),
	"ecdsa_sha2_nistp384": types.BoolValue(false),
	"ecdsa_sha2_nistp521": types.BoolValue(false),
	"rsa_sha2_512":        types.BoolValue(false),
	"rsa_sha2_256":        types.BoolValue(false),
}
var DmHostKeyAlgorithmsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"ssh_ed25519": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ssh-ed25519", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ecdsa_sha2_nistp256": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp256", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ecdsa_sha2_nistp384": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp384", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ecdsa_sha2_nistp521": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp521", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"rsa_sha2_512": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("rsa-sha2-512", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"rsa_sha2_256": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("rsa-sha2-256", "", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmHostKeyAlgorithmsResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmHostKeyAlgorithmsObjectType,
			DmHostKeyAlgorithmsObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"ssh_ed25519": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ssh-ed25519", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ecdsa_sha2_nistp256": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp256", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ecdsa_sha2_nistp384": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp384", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ecdsa_sha2_nistp521": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("ecdsa-sha2-nistp521", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"rsa_sha2_512": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("rsa-sha2-512", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"rsa_sha2_256": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("rsa-sha2-256", "", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmHostKeyAlgorithms) IsNull() bool {
	if !data.SshEd25519.IsNull() {
		return false
	}
	if !data.EcdsaSha2Nistp256.IsNull() {
		return false
	}
	if !data.EcdsaSha2Nistp384.IsNull() {
		return false
	}
	if !data.EcdsaSha2Nistp521.IsNull() {
		return false
	}
	if !data.RsaSha2512.IsNull() {
		return false
	}
	if !data.RsaSha2256.IsNull() {
		return false
	}
	return true
}
func GetDmHostKeyAlgorithmsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmHostKeyAlgorithmsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmHostKeyAlgorithmsDataSourceSchema
}

func GetDmHostKeyAlgorithmsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmHostKeyAlgorithmsResourceSchema.Required = true
	} else {
		DmHostKeyAlgorithmsResourceSchema.Optional = true
		DmHostKeyAlgorithmsResourceSchema.Computed = true
	}
	DmHostKeyAlgorithmsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmHostKeyAlgorithmsResourceSchema
}

func (data DmHostKeyAlgorithms) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SshEd25519.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSH-ED25519`, tfutils.StringFromBool(data.SshEd25519, ""))
	}
	if !data.EcdsaSha2Nistp256.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ECDSA-SHA2-NISTP256`, tfutils.StringFromBool(data.EcdsaSha2Nistp256, ""))
	}
	if !data.EcdsaSha2Nistp384.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ECDSA-SHA2-NISTP384`, tfutils.StringFromBool(data.EcdsaSha2Nistp384, ""))
	}
	if !data.EcdsaSha2Nistp521.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ECDSA-SHA2-NISTP521`, tfutils.StringFromBool(data.EcdsaSha2Nistp521, ""))
	}
	if !data.RsaSha2512.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RSA-SHA2-512`, tfutils.StringFromBool(data.RsaSha2512, ""))
	}
	if !data.RsaSha2256.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RSA-SHA2-256`, tfutils.StringFromBool(data.RsaSha2256, ""))
	}
	return body
}

func (data *DmHostKeyAlgorithms) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SSH-ED25519`); value.Exists() {
		data.SshEd25519 = tfutils.BoolFromString(value.String())
	} else {
		data.SshEd25519 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP256`); value.Exists() {
		data.EcdsaSha2Nistp256 = tfutils.BoolFromString(value.String())
	} else {
		data.EcdsaSha2Nistp256 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP384`); value.Exists() {
		data.EcdsaSha2Nistp384 = tfutils.BoolFromString(value.String())
	} else {
		data.EcdsaSha2Nistp384 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP521`); value.Exists() {
		data.EcdsaSha2Nistp521 = tfutils.BoolFromString(value.String())
	} else {
		data.EcdsaSha2Nistp521 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RSA-SHA2-512`); value.Exists() {
		data.RsaSha2512 = tfutils.BoolFromString(value.String())
	} else {
		data.RsaSha2512 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RSA-SHA2-256`); value.Exists() {
		data.RsaSha2256 = tfutils.BoolFromString(value.String())
	} else {
		data.RsaSha2256 = types.BoolNull()
	}
}

func (data *DmHostKeyAlgorithms) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SSH-ED25519`); value.Exists() && !data.SshEd25519.IsNull() {
		data.SshEd25519 = tfutils.BoolFromString(value.String())
	} else if data.SshEd25519.ValueBool() {
		data.SshEd25519 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP256`); value.Exists() && !data.EcdsaSha2Nistp256.IsNull() {
		data.EcdsaSha2Nistp256 = tfutils.BoolFromString(value.String())
	} else if data.EcdsaSha2Nistp256.ValueBool() {
		data.EcdsaSha2Nistp256 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP384`); value.Exists() && !data.EcdsaSha2Nistp384.IsNull() {
		data.EcdsaSha2Nistp384 = tfutils.BoolFromString(value.String())
	} else if data.EcdsaSha2Nistp384.ValueBool() {
		data.EcdsaSha2Nistp384 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ECDSA-SHA2-NISTP521`); value.Exists() && !data.EcdsaSha2Nistp521.IsNull() {
		data.EcdsaSha2Nistp521 = tfutils.BoolFromString(value.String())
	} else if data.EcdsaSha2Nistp521.ValueBool() {
		data.EcdsaSha2Nistp521 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RSA-SHA2-512`); value.Exists() && !data.RsaSha2512.IsNull() {
		data.RsaSha2512 = tfutils.BoolFromString(value.String())
	} else if data.RsaSha2512.ValueBool() {
		data.RsaSha2512 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RSA-SHA2-256`); value.Exists() && !data.RsaSha2256.IsNull() {
		data.RsaSha2256 = tfutils.BoolFromString(value.String())
	} else if data.RsaSha2256.ValueBool() {
		data.RsaSha2256 = types.BoolNull()
	}
}
