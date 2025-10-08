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

type DmInvokeErrorType struct {
	ConnectionError types.Bool `tfsdk:"connection_error"`
	SoapError       types.Bool `tfsdk:"soap_error"`
	OperationError  types.Bool `tfsdk:"operation_error"`
}

var DmInvokeErrorTypeObjectType = map[string]attr.Type{
	"connection_error": types.BoolType,
	"soap_error":       types.BoolType,
	"operation_error":  types.BoolType,
}
var DmInvokeErrorTypeObjectDefault = map[string]attr.Value{
	"connection_error": types.BoolValue(false),
	"soap_error":       types.BoolValue(false),
	"operation_error":  types.BoolValue(false),
}

func GetDmInvokeErrorTypeDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmInvokeErrorTypeDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"connection_error": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Connection error",
				Computed:            true,
			},
			"soap_error": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "SOAP error",
				Computed:            true,
			},
			"operation_error": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Operation error",
				Computed:            true,
			},
		},
	}
	DmInvokeErrorTypeDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmInvokeErrorTypeDataSourceSchema
}
func GetDmInvokeErrorTypeResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmInvokeErrorTypeResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmInvokeErrorTypeObjectType,
				DmInvokeErrorTypeObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"connection_error": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Connection error", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"soap_error": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP error", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"operation_error": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Operation error", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmInvokeErrorTypeResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmInvokeErrorTypeResourceSchema.Required = true
	} else {
		DmInvokeErrorTypeResourceSchema.Optional = true
		DmInvokeErrorTypeResourceSchema.Computed = true
	}
	return DmInvokeErrorTypeResourceSchema
}

func (data DmInvokeErrorType) IsNull() bool {
	if !data.ConnectionError.IsNull() {
		return false
	}
	if !data.SoapError.IsNull() {
		return false
	}
	if !data.OperationError.IsNull() {
		return false
	}
	return true
}

func (data DmInvokeErrorType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ConnectionError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectionError`, tfutils.StringFromBool(data.ConnectionError, ""))
	}
	if !data.SoapError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPError`, tfutils.StringFromBool(data.SoapError, ""))
	}
	if !data.OperationError.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OperationError`, tfutils.StringFromBool(data.OperationError, ""))
	}
	return body
}

func (data *DmInvokeErrorType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ConnectionError`); value.Exists() {
		data.ConnectionError = tfutils.BoolFromString(value.String())
	} else {
		data.ConnectionError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPError`); value.Exists() {
		data.SoapError = tfutils.BoolFromString(value.String())
	} else {
		data.SoapError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OperationError`); value.Exists() {
		data.OperationError = tfutils.BoolFromString(value.String())
	} else {
		data.OperationError = types.BoolNull()
	}
}

func (data *DmInvokeErrorType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ConnectionError`); value.Exists() && !data.ConnectionError.IsNull() {
		data.ConnectionError = tfutils.BoolFromString(value.String())
	} else if data.ConnectionError.ValueBool() {
		data.ConnectionError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SOAPError`); value.Exists() && !data.SoapError.IsNull() {
		data.SoapError = tfutils.BoolFromString(value.String())
	} else if data.SoapError.ValueBool() {
		data.SoapError = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OperationError`); value.Exists() && !data.OperationError.IsNull() {
		data.OperationError = tfutils.BoolFromString(value.String())
	} else if data.OperationError.ValueBool() {
		data.OperationError = types.BoolNull()
	}
}
