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

type DmWSDLUserPolicyToggles struct {
	Enable                               types.Bool `tfsdk:"enable"`
	Publish                              types.Bool `tfsdk:"publish"`
	VerifyFaults                         types.Bool `tfsdk:"verify_faults"`
	VerifyHeaders                        types.Bool `tfsdk:"verify_headers"`
	NoRequestValidation                  types.Bool `tfsdk:"no_request_validation"`
	NoResponseValidation                 types.Bool `tfsdk:"no_response_validation"`
	SuppressFaultsElementsForRpcWrappers types.Bool `tfsdk:"suppress_faults_elements_for_rpc_wrappers"`
	NoWsA                                types.Bool `tfsdk:"no_ws_a"`
	NoWsRm                               types.Bool `tfsdk:"no_ws_rm"`
	AllowXopInclude                      types.Bool `tfsdk:"allow_xop_include"`
}

var DmWSDLUserPolicyTogglesObjectType = map[string]attr.Type{
	"enable":                 types.BoolType,
	"publish":                types.BoolType,
	"verify_faults":          types.BoolType,
	"verify_headers":         types.BoolType,
	"no_request_validation":  types.BoolType,
	"no_response_validation": types.BoolType,
	"suppress_faults_elements_for_rpc_wrappers": types.BoolType,
	"no_ws_a":           types.BoolType,
	"no_ws_rm":          types.BoolType,
	"allow_xop_include": types.BoolType,
}
var DmWSDLUserPolicyTogglesObjectDefault = map[string]attr.Value{
	"enable":                 types.BoolValue(true),
	"publish":                types.BoolValue(true),
	"verify_faults":          types.BoolValue(true),
	"verify_headers":         types.BoolValue(false),
	"no_request_validation":  types.BoolValue(false),
	"no_response_validation": types.BoolValue(false),
	"suppress_faults_elements_for_rpc_wrappers": types.BoolValue(false),
	"no_ws_a":           types.BoolValue(false),
	"no_ws_rm":          types.BoolValue(false),
	"allow_xop_include": types.BoolValue(true),
}

func GetDmWSDLUserPolicyTogglesDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmWSDLUserPolicyTogglesDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"enable": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Enable this component",
				Computed:            true,
			},
			"publish": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Publish in WSDL",
				Computed:            true,
			},
			"verify_faults": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Schema validate faults messages",
				Computed:            true,
			},
			"verify_headers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Schema validate SOAP headers",
				Computed:            true,
			},
			"no_request_validation": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "No Request Validation",
				Computed:            true,
			},
			"no_response_validation": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "No Response Validation",
				Computed:            true,
			},
			"suppress_faults_elements_for_rpc_wrappers": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Strict Fault Document Style",
				Computed:            true,
			},
			"no_ws_a": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Opt out of WS-Addressing",
				Computed:            true,
			},
			"no_ws_rm": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Opt out of WS-ReliableMessaging",
				Computed:            true,
			},
			"allow_xop_include": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Accept MTOM/XOP Optimized Messages",
				Computed:            true,
			},
		},
	}
	DmWSDLUserPolicyTogglesDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmWSDLUserPolicyTogglesDataSourceSchema
}
func GetDmWSDLUserPolicyTogglesResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmWSDLUserPolicyTogglesResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmWSDLUserPolicyTogglesObjectType,
				DmWSDLUserPolicyTogglesObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"enable": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable this component", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"publish": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Publish in WSDL", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"verify_faults": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schema validate faults messages", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"verify_headers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schema validate SOAP headers", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_request_validation": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("No Request Validation", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_response_validation": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("No Response Validation", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"suppress_faults_elements_for_rpc_wrappers": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict Fault Document Style", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_ws_a": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Opt out of WS-Addressing", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_ws_rm": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Opt out of WS-ReliableMessaging", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_xop_include": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Accept MTOM/XOP Optimized Messages", "", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
		},
	}
	DmWSDLUserPolicyTogglesResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmWSDLUserPolicyTogglesResourceSchema.Required = true
	} else {
		DmWSDLUserPolicyTogglesResourceSchema.Optional = true
		DmWSDLUserPolicyTogglesResourceSchema.Computed = true
	}
	return DmWSDLUserPolicyTogglesResourceSchema
}

func (data DmWSDLUserPolicyToggles) IsNull() bool {
	if !data.Enable.IsNull() {
		return false
	}
	if !data.Publish.IsNull() {
		return false
	}
	if !data.VerifyFaults.IsNull() {
		return false
	}
	if !data.VerifyHeaders.IsNull() {
		return false
	}
	if !data.NoRequestValidation.IsNull() {
		return false
	}
	if !data.NoResponseValidation.IsNull() {
		return false
	}
	if !data.SuppressFaultsElementsForRpcWrappers.IsNull() {
		return false
	}
	if !data.NoWsA.IsNull() {
		return false
	}
	if !data.NoWsRm.IsNull() {
		return false
	}
	if !data.AllowXopInclude.IsNull() {
		return false
	}
	return true
}

func (data DmWSDLUserPolicyToggles) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Enable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Enable`, tfutils.StringFromBool(data.Enable, ""))
	}
	if !data.Publish.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Publish`, tfutils.StringFromBool(data.Publish, ""))
	}
	if !data.VerifyFaults.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyFaults`, tfutils.StringFromBool(data.VerifyFaults, ""))
	}
	if !data.VerifyHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`VerifyHeaders`, tfutils.StringFromBool(data.VerifyHeaders, ""))
	}
	if !data.NoRequestValidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoRequestValidation`, tfutils.StringFromBool(data.NoRequestValidation, ""))
	}
	if !data.NoResponseValidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoResponseValidation`, tfutils.StringFromBool(data.NoResponseValidation, ""))
	}
	if !data.SuppressFaultsElementsForRpcWrappers.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SuppressFaultsElementsForRPCWrappers`, tfutils.StringFromBool(data.SuppressFaultsElementsForRpcWrappers, ""))
	}
	if !data.NoWsA.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoWSA`, tfutils.StringFromBool(data.NoWsA, ""))
	}
	if !data.NoWsRm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoWSRM`, tfutils.StringFromBool(data.NoWsRm, ""))
	}
	if !data.AllowXopInclude.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowXOPInclude`, tfutils.StringFromBool(data.AllowXopInclude, ""))
	}
	return body
}

func (data *DmWSDLUserPolicyToggles) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Enable`); value.Exists() {
		data.Enable = tfutils.BoolFromString(value.String())
	} else {
		data.Enable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Publish`); value.Exists() {
		data.Publish = tfutils.BoolFromString(value.String())
	} else {
		data.Publish = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyFaults`); value.Exists() {
		data.VerifyFaults = tfutils.BoolFromString(value.String())
	} else {
		data.VerifyFaults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyHeaders`); value.Exists() {
		data.VerifyHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.VerifyHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoRequestValidation`); value.Exists() {
		data.NoRequestValidation = tfutils.BoolFromString(value.String())
	} else {
		data.NoRequestValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoResponseValidation`); value.Exists() {
		data.NoResponseValidation = tfutils.BoolFromString(value.String())
	} else {
		data.NoResponseValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuppressFaultsElementsForRPCWrappers`); value.Exists() {
		data.SuppressFaultsElementsForRpcWrappers = tfutils.BoolFromString(value.String())
	} else {
		data.SuppressFaultsElementsForRpcWrappers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoWSA`); value.Exists() {
		data.NoWsA = tfutils.BoolFromString(value.String())
	} else {
		data.NoWsA = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoWSRM`); value.Exists() {
		data.NoWsRm = tfutils.BoolFromString(value.String())
	} else {
		data.NoWsRm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowXOPInclude`); value.Exists() {
		data.AllowXopInclude = tfutils.BoolFromString(value.String())
	} else {
		data.AllowXopInclude = types.BoolNull()
	}
}

func (data *DmWSDLUserPolicyToggles) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Enable`); value.Exists() && !data.Enable.IsNull() {
		data.Enable = tfutils.BoolFromString(value.String())
	} else if !data.Enable.ValueBool() {
		data.Enable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Publish`); value.Exists() && !data.Publish.IsNull() {
		data.Publish = tfutils.BoolFromString(value.String())
	} else if !data.Publish.ValueBool() {
		data.Publish = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyFaults`); value.Exists() && !data.VerifyFaults.IsNull() {
		data.VerifyFaults = tfutils.BoolFromString(value.String())
	} else if !data.VerifyFaults.ValueBool() {
		data.VerifyFaults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VerifyHeaders`); value.Exists() && !data.VerifyHeaders.IsNull() {
		data.VerifyHeaders = tfutils.BoolFromString(value.String())
	} else if data.VerifyHeaders.ValueBool() {
		data.VerifyHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoRequestValidation`); value.Exists() && !data.NoRequestValidation.IsNull() {
		data.NoRequestValidation = tfutils.BoolFromString(value.String())
	} else if data.NoRequestValidation.ValueBool() {
		data.NoRequestValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoResponseValidation`); value.Exists() && !data.NoResponseValidation.IsNull() {
		data.NoResponseValidation = tfutils.BoolFromString(value.String())
	} else if data.NoResponseValidation.ValueBool() {
		data.NoResponseValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SuppressFaultsElementsForRPCWrappers`); value.Exists() && !data.SuppressFaultsElementsForRpcWrappers.IsNull() {
		data.SuppressFaultsElementsForRpcWrappers = tfutils.BoolFromString(value.String())
	} else if data.SuppressFaultsElementsForRpcWrappers.ValueBool() {
		data.SuppressFaultsElementsForRpcWrappers = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoWSA`); value.Exists() && !data.NoWsA.IsNull() {
		data.NoWsA = tfutils.BoolFromString(value.String())
	} else if data.NoWsA.ValueBool() {
		data.NoWsA = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NoWSRM`); value.Exists() && !data.NoWsRm.IsNull() {
		data.NoWsRm = tfutils.BoolFromString(value.String())
	} else if data.NoWsRm.ValueBool() {
		data.NoWsRm = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowXOPInclude`); value.Exists() && !data.AllowXopInclude.IsNull() {
		data.AllowXopInclude = tfutils.BoolFromString(value.String())
	} else if !data.AllowXopInclude.ValueBool() {
		data.AllowXopInclude = types.BoolNull()
	}
}
