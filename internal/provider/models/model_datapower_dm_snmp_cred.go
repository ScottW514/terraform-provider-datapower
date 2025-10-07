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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSnmpCred struct {
	EngineId            types.String `tfsdk:"engine_id"`
	AuthProtocol        types.String `tfsdk:"auth_protocol"`
	AuthSecretType      types.String `tfsdk:"auth_secret_type"`
	AuthSecretWo        types.String `tfsdk:"auth_secret_wo"`
	AuthSecretWoVersion types.Int64  `tfsdk:"auth_secret_wo_version"`
	PrivProtocol        types.String `tfsdk:"priv_protocol"`
	PrivSecretType      types.String `tfsdk:"priv_secret_type"`
	PrivSecretWo        types.String `tfsdk:"priv_secret_wo"`
	PrivSecretWoVersion types.Int64  `tfsdk:"priv_secret_wo_version"`
}
type DmSnmpCredWO struct {
	EngineId       types.String `tfsdk:"engine_id"`
	AuthProtocol   types.String `tfsdk:"auth_protocol"`
	AuthSecretType types.String `tfsdk:"auth_secret_type"`
	PrivProtocol   types.String `tfsdk:"priv_protocol"`
	PrivSecretType types.String `tfsdk:"priv_secret_type"`
}

var DmSnmpCredAuthSecretTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "auth_protocol",
	AttrType:    "String",
	AttrDefault: "sha",
	Value:       []string{"none"},
}
var DmSnmpCredAuthSecretCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "auth_protocol",
	AttrType:    "String",
	AttrDefault: "sha",
	Value:       []string{"none"},
}
var DmSnmpCredPrivSecretTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "priv_protocol",
	AttrType:    "String",
	AttrDefault: "des",
	Value:       []string{"none"},
}
var DmSnmpCredPrivSecretCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "priv_protocol",
	AttrType:    "String",
	AttrDefault: "des",
	Value:       []string{"none"},
}
var DmSnmpCredAuthSecretTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmSnmpCredAuthSecretIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "auth_protocol",
	AttrType:    "String",
	AttrDefault: "sha",
	Value:       []string{"none"},
}
var DmSnmpCredPrivSecretTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmSnmpCredPrivSecretIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "priv_protocol",
	AttrType:    "String",
	AttrDefault: "des",
	Value:       []string{"none"},
}

var DmSnmpCredObjectType = map[string]attr.Type{
	"engine_id":              types.StringType,
	"auth_protocol":          types.StringType,
	"auth_secret_type":       types.StringType,
	"auth_secret_wo":         types.StringType,
	"auth_secret_wo_version": types.Int64Type,
	"priv_protocol":          types.StringType,
	"priv_secret_type":       types.StringType,
	"priv_secret_wo":         types.StringType,
	"priv_secret_wo_version": types.Int64Type,
}
var DmSnmpCredObjectTypeWO = map[string]attr.Type{
	"engine_id":        types.StringType,
	"auth_protocol":    types.StringType,
	"auth_secret_type": types.StringType,
	"priv_protocol":    types.StringType,
	"priv_secret_type": types.StringType,
}
var DmSnmpCredObjectDefault = map[string]attr.Value{
	"engine_id":        types.StringNull(),
	"auth_protocol":    types.StringValue("sha"),
	"auth_secret_type": types.StringValue("password"),
	"auth_secret_wo":   types.StringNull(),
	"priv_protocol":    types.StringValue("des"),
	"priv_secret_type": types.StringValue("password"),
	"priv_secret_wo":   types.StringNull(),
}

func GetDmSnmpCredDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSnmpCredDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"engine_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SNMPv3 engine ID. The value of 0 represents the local engine ID. For any other engine ID, the value is a hex string that represents the 5 - 32 byte value.", "", "").String,
				Computed:            true,
			},
			"auth_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the authentication protocol.", "", "").AddStringEnum("none", "md5", "sha").AddDefaultValue("sha").String,
				Computed:            true,
			},
			"auth_secret_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the format of the authentication key.", "", "").AddStringEnum("password", "key").AddDefaultValue("password").AddRequiredWhen(DmSnmpCredAuthSecretTypeCondVal.String()).AddNotValidWhen(DmSnmpCredAuthSecretTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"priv_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the privacy protocol.", "", "").AddStringEnum("none", "des", "aes").AddDefaultValue("des").String,
				Computed:            true,
			},
			"priv_secret_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the format of the privacy key. When the privacy protocol is AED or DES, whether the <tt>PrivKey</tt> is generated from a plaintext string or is an explicit key value.", "", "").AddStringEnum("password", "key").AddDefaultValue("password").AddRequiredWhen(DmSnmpCredPrivSecretTypeCondVal.String()).AddNotValidWhen(DmSnmpCredPrivSecretTypeIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	return DmSnmpCredDataSourceSchema
}
func GetDmSnmpCredResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSnmpCredResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"engine_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SNMPv3 engine ID. The value of 0 represents the local engine ID. For any other engine ID, the value is a hex string that represents the 5 - 32 byte value.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
					stringvalidator.RegexMatches(regexp.MustCompile("^(0|[0-9a-fA-F]{10,64})$"), "Must match :"+"^(0|[0-9a-fA-F]{10,64})$"),
				},
			},
			"auth_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the authentication protocol.", "", "").AddStringEnum("none", "md5", "sha").AddDefaultValue("sha").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "md5", "sha"),
				},
				Default: stringdefault.StaticString("sha"),
			},
			"auth_secret_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the format of the authentication key.", "", "").AddStringEnum("password", "key").AddDefaultValue("password").AddRequiredWhen(DmSnmpCredAuthSecretTypeCondVal.String()).AddNotValidWhen(DmSnmpCredAuthSecretTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("password", "key"),
					validators.ConditionalRequiredString(DmSnmpCredAuthSecretTypeCondVal, DmSnmpCredAuthSecretTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("password"),
			},
			"auth_secret_wo": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the authentication value. The <tt>AuthKey</tt> can be an explicit value or generated by the DataPower Gateway. An explicit value is useful when the key is created on another system. Enter the plaintext string to generate the localized key, or enter the hex representation of the localized key. <ul><li>When the DataPower Gateway generates and stores the appropriate length key, enter a plaintext string that is at least 8 characters long as the hash.</li><li>When the value type is a localized key, enter the explicit key value. You can use colons (:) between every 2 hex characters.</li><ul><li>For MD5 authentication, enter the hex representation of the 16-byte key.</li><li>For SHA authentication, enter the hex representation of a 20-byte key.</li></ul></ul>", "", "").AddRequiredWhen(DmSnmpCredAuthSecretCondVal.String()).AddNotValidWhen(DmSnmpCredAuthSecretIgnoreVal.String()).String,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSnmpCredAuthSecretCondVal, DmSnmpCredAuthSecretIgnoreVal, false),
				},
			},
			"auth_secret_wo_version": ResourceSchema.Int64Attribute{
				MarkdownDescription: "Changes to this value trigger an update to `write_only` value.",
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(
						validators.Evaluation{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "auth_secret_wo",
							AttrType:    "String",
							AttrDefault: "",
							Value:       []string{""},
						}, validators.Evaluation{}, false),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"priv_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the privacy protocol.", "", "").AddStringEnum("none", "des", "aes").AddDefaultValue("des").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "des", "aes"),
				},
				Default: stringdefault.StaticString("des"),
			},
			"priv_secret_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the format of the privacy key. When the privacy protocol is AED or DES, whether the <tt>PrivKey</tt> is generated from a plaintext string or is an explicit key value.", "", "").AddStringEnum("password", "key").AddDefaultValue("password").AddRequiredWhen(DmSnmpCredPrivSecretTypeCondVal.String()).AddNotValidWhen(DmSnmpCredPrivSecretTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("password", "key"),
					validators.ConditionalRequiredString(DmSnmpCredPrivSecretTypeCondVal, DmSnmpCredPrivSecretTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("password"),
			},
			"priv_secret_wo": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the privacy value. The <tt>PrivKey</tt> can be an explicit value or generated by the DataPower Gateway. An explicit value is useful when the key is created on another system. Enter the plaintext string to generate the localized key, or enter the hex representation of the localized privacy key. <ul><li>When the DataPower Gateway generates and stores the appropriate length key, enter a plaintext string that is at least 8 characters long as the hash.</li><li>When the value type is a localized key, enter the explicit key value. You can use colons (:) between every 2 hex characters.</li><ul><li>For MD5 authentication, enter the hex representation of the 16-byte key.</li><li>For SHA authentication, enter the hex representation of a 20-byte key.</li></ul></ul>", "", "").AddRequiredWhen(DmSnmpCredPrivSecretCondVal.String()).AddNotValidWhen(DmSnmpCredPrivSecretIgnoreVal.String()).String,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSnmpCredPrivSecretCondVal, DmSnmpCredPrivSecretIgnoreVal, false),
				},
			},
			"priv_secret_wo_version": ResourceSchema.Int64Attribute{
				MarkdownDescription: "Changes to this value trigger an update to `write_only` value.",
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(
						validators.Evaluation{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "priv_secret_wo",
							AttrType:    "String",
							AttrDefault: "",
							Value:       []string{""},
						}, validators.Evaluation{}, false),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
	return DmSnmpCredResourceSchema
}

func (data DmSnmpCred) IsNull() bool {
	if !data.EngineId.IsNull() {
		return false
	}
	if !data.AuthProtocol.IsNull() {
		return false
	}
	if !data.AuthSecretType.IsNull() {
		return false
	}
	if !data.AuthSecretWo.IsNull() {
		return false
	}
	if !data.PrivProtocol.IsNull() {
		return false
	}
	if !data.PrivSecretType.IsNull() {
		return false
	}
	if !data.PrivSecretWo.IsNull() {
		return false
	}
	return true
}
func (data DmSnmpCredWO) IsNull() bool {
	if !data.EngineId.IsNull() {
		return false
	}
	if !data.AuthProtocol.IsNull() {
		return false
	}
	if !data.AuthSecretType.IsNull() {
		return false
	}
	if !data.PrivProtocol.IsNull() {
		return false
	}
	if !data.PrivSecretType.IsNull() {
		return false
	}
	return true
}

func (data DmSnmpCred) ToBody(ctx context.Context, pathRoot string, config *DmSnmpCred) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.EngineId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EngineID`, data.EngineId.ValueString())
	}
	if !data.AuthProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthProtocol`, data.AuthProtocol.ValueString())
	}
	if !data.AuthSecretType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthSecretType`, data.AuthSecretType.ValueString())
	}
	if !data.AuthSecretWo.IsNull() || !data.AuthSecretWoVersion.IsNull() {
		if data.AuthSecretWo.IsNull() && config != nil {
			data.AuthSecretWo = config.AuthSecretWo
		}
		body, _ = sjson.Set(body, pathRoot+`AuthSecret`, data.AuthSecretWo.ValueString())
	}
	if !data.PrivProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivProtocol`, data.PrivProtocol.ValueString())
	}
	if !data.PrivSecretType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PrivSecretType`, data.PrivSecretType.ValueString())
	}
	if !data.PrivSecretWo.IsNull() || !data.PrivSecretWoVersion.IsNull() {
		if data.PrivSecretWo.IsNull() && config != nil {
			data.PrivSecretWo = config.PrivSecretWo
		}
		body, _ = sjson.Set(body, pathRoot+`PrivSecret`, data.PrivSecretWo.ValueString())
	}
	return body
}

func (data *DmSnmpCred) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EngineID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EngineId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthProtocol = types.StringValue("sha")
	}
	if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecretType = types.StringValue("password")
	}
	if value := res.Get(pathRoot + `AuthSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecretWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivProtocol = types.StringValue("des")
	}
	if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecretType = types.StringValue("password")
	}
	if value := res.Get(pathRoot + `PrivSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecretWo = types.StringNull()
	}
}
func (data *DmSnmpCredWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EngineID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EngineId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthProtocol = types.StringValue("sha")
	}
	if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecretType = types.StringValue("password")
	}
	if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivProtocol = types.StringValue("des")
	}
	if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecretType = types.StringValue("password")
	}
}

func (data *DmSnmpCred) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EngineID`); value.Exists() && !data.EngineId.IsNull() {
		data.EngineId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EngineId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && !data.AuthProtocol.IsNull() {
		data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthProtocol.ValueString() != "sha" {
		data.AuthProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && !data.AuthSecretType.IsNull() {
		data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthSecretType.ValueString() != "password" {
		data.AuthSecretType = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthSecret`); value.Exists() && !data.AuthSecretWo.IsNull() {
		data.AuthSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthSecretWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && !data.PrivProtocol.IsNull() {
		data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.PrivProtocol.ValueString() != "des" {
		data.PrivProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && !data.PrivSecretType.IsNull() {
		data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
	} else if data.PrivSecretType.ValueString() != "password" {
		data.PrivSecretType = types.StringNull()
	}
	if value := res.Get(pathRoot + `PrivSecret`); value.Exists() && !data.PrivSecretWo.IsNull() {
		data.PrivSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PrivSecretWo = types.StringNull()
	}
}
func (data *DmSnmpCred) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.EngineId.IsUnknown() {
		if value := res.Get(pathRoot + `EngineID`); value.Exists() && !data.EngineId.IsNull() {
			data.EngineId = tfutils.ParseStringFromGJSON(value)
		} else {
			data.EngineId = types.StringNull()
		}
	}
	if data.AuthProtocol.IsUnknown() {
		if value := res.Get(pathRoot + `AuthProtocol`); value.Exists() && !data.AuthProtocol.IsNull() {
			data.AuthProtocol = tfutils.ParseStringFromGJSON(value)
		} else if data.AuthProtocol.ValueString() != "sha" {
			data.AuthProtocol = types.StringNull()
		}
	}
	if data.AuthSecretType.IsUnknown() {
		if value := res.Get(pathRoot + `AuthSecretType`); value.Exists() && !data.AuthSecretType.IsNull() {
			data.AuthSecretType = tfutils.ParseStringFromGJSON(value)
		} else if data.AuthSecretType.ValueString() != "password" {
			data.AuthSecretType = types.StringNull()
		}
	}
	if data.AuthSecretWo.IsUnknown() {
		if value := res.Get(pathRoot + `AuthSecret`); value.Exists() && !data.AuthSecretWo.IsNull() {
			data.AuthSecretWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.AuthSecretWo = types.StringNull()
		}
	}
	if data.PrivProtocol.IsUnknown() {
		if value := res.Get(pathRoot + `PrivProtocol`); value.Exists() && !data.PrivProtocol.IsNull() {
			data.PrivProtocol = tfutils.ParseStringFromGJSON(value)
		} else if data.PrivProtocol.ValueString() != "des" {
			data.PrivProtocol = types.StringNull()
		}
	}
	if data.PrivSecretType.IsUnknown() {
		if value := res.Get(pathRoot + `PrivSecretType`); value.Exists() && !data.PrivSecretType.IsNull() {
			data.PrivSecretType = tfutils.ParseStringFromGJSON(value)
		} else if data.PrivSecretType.ValueString() != "password" {
			data.PrivSecretType = types.StringNull()
		}
	}
	if data.PrivSecretWo.IsUnknown() {
		if value := res.Get(pathRoot + `PrivSecret`); value.Exists() && !data.PrivSecretWo.IsNull() {
			data.PrivSecretWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.PrivSecretWo = types.StringNull()
		}
	}
}
