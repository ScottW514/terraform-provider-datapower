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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmRadiusServer struct {
	Number          types.Int64  `tfsdk:"number"`
	Host            types.String `tfsdk:"host"`
	Port            types.Int64  `tfsdk:"port"`
	SecretWo        types.String `tfsdk:"secret_wo"`
	SecretWoVersion types.Int64  `tfsdk:"secret_wo_version"`
}
type DmRadiusServerWO struct {
	Number types.Int64  `tfsdk:"number"`
	Host   types.String `tfsdk:"host"`
	Port   types.Int64  `tfsdk:"port"`
}

var DmRadiusServerObjectType = map[string]attr.Type{
	"number":            types.Int64Type,
	"host":              types.StringType,
	"port":              types.Int64Type,
	"secret_wo":         types.StringType,
	"secret_wo_version": types.Int64Type,
}
var DmRadiusServerObjectTypeWO = map[string]attr.Type{
	"number": types.Int64Type,
	"host":   types.StringType,
	"port":   types.Int64Type,
}
var DmRadiusServerObjectDefault = map[string]attr.Value{
	"number":    types.Int64Null(),
	"host":      types.StringNull(),
	"port":      types.Int64Value(1812),
	"secret_wo": types.StringNull(),
}

func GetDmRadiusServerDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmRadiusServerDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"number": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list position of this RADIUS server within the list of all RADIUS servers known to the client implementation. The lower the number, the more preferred the server.", "", "").AddIntegerRange(0, 2147483647).String,
				Computed:            true,
			},
			"host": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address of the RADIUS server.", "", "").String,
				Computed:            true,
			},
			"port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the RADIUS server.", "", "").AddDefaultValue("1812").String,
				Computed:            true,
			},
		},
	}
	return DmRadiusServerDataSourceSchema
}
func GetDmRadiusServerResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmRadiusServerResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"number": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list position of this RADIUS server within the list of all RADIUS servers known to the client implementation. The lower the number, the more preferred the server.", "", "").AddIntegerRange(0, 2147483647).String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"host": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address of the RADIUS server.", "", "").String,
				Required:            true,
			},
			"port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the RADIUS server.", "", "").AddDefaultValue("1812").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(1812),
			},
			"secret_wo": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password login to the RADIUS server. You must confirm the password for accuracy.", "", "").String,
				WriteOnly:           true,
				Required:            true,
			},
			"secret_wo_version": ResourceSchema.Int64Attribute{
				MarkdownDescription: "Changes to this value trigger an update to `write_only` value.",
				Optional:            true,
				Validators: []validator.Int64{
					validators.ConditionalRequiredInt64(
						validators.Evaluation{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "secret_wo",
							AttrType:    "String",
							AttrDefault: "",
							Value:       []string{""},
						}, validators.Evaluation{}, false),
				},
			},
		},
	}
	return DmRadiusServerResourceSchema
}

func (data DmRadiusServer) IsNull() bool {
	if !data.Number.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.SecretWo.IsNull() {
		return false
	}
	return true
}
func (data DmRadiusServerWO) IsNull() bool {
	if !data.Number.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	return true
}

func (data DmRadiusServer) ToBody(ctx context.Context, pathRoot string, config *DmRadiusServer) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Number.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Number`, data.Number.ValueInt64())
	}
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.SecretWo.IsNull() || !data.SecretWoVersion.IsNull() {
		if data.SecretWo.IsNull() && config != nil {
			data.SecretWo = config.SecretWo
		}
		body, _ = sjson.Set(body, pathRoot+`Secret`, data.SecretWo.ValueString())
	}
	return body
}

func (data *DmRadiusServer) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Number`); value.Exists() {
		data.Number = types.Int64Value(value.Int())
	} else {
		data.Number = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(1812)
	}
	if value := res.Get(pathRoot + `Secret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecretWo = types.StringNull()
	}
}
func (data *DmRadiusServerWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Number`); value.Exists() {
		data.Number = types.Int64Value(value.Int())
	} else {
		data.Number = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(1812)
	}
}

func (data *DmRadiusServer) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Number`); value.Exists() && !data.Number.IsNull() {
		data.Number = types.Int64Value(value.Int())
	} else {
		data.Number = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 1812 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Secret`); value.Exists() && !data.SecretWo.IsNull() {
		data.SecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SecretWo = types.StringNull()
	}
}
func (data *DmRadiusServer) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Number.IsUnknown() {
		if value := res.Get(pathRoot + `Number`); value.Exists() && !data.Number.IsNull() {
			data.Number = types.Int64Value(value.Int())
		} else {
			data.Number = types.Int64Null()
		}
	}
	if data.Host.IsUnknown() {
		if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
			data.Host = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Host = types.StringNull()
		}
	}
	if data.Port.IsUnknown() {
		if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else if data.Port.ValueInt64() != 1812 {
			data.Port = types.Int64Null()
		}
	}
	if data.SecretWo.IsUnknown() {
		if value := res.Get(pathRoot + `Secret`); value.Exists() && !data.SecretWo.IsNull() {
			data.SecretWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SecretWo = types.StringNull()
		}
	}
}
