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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmTAMRASTrace struct {
	TamTraceEnable     types.Bool   `tfsdk:"tam_trace_enable"`
	TamTraceFile       types.String `tfsdk:"tam_trace_file"`
	TamTraceSize       types.Int64  `tfsdk:"tam_trace_size"`
	TamTraceType       types.String `tfsdk:"tam_trace_type"`
	TamTraceComponents types.String `tfsdk:"tam_trace_components"`
	LdapTraceEnable    types.Bool   `tfsdk:"ldap_trace_enable"`
	LdapTraceFile      types.String `tfsdk:"ldap_trace_file"`
	LdapTraceSize      types.Int64  `tfsdk:"ldap_trace_size"`
	LdapTraceLevel     types.Int64  `tfsdk:"ldap_trace_level"`
	GsKitTraceEnable   types.Bool   `tfsdk:"gs_kit_trace_enable"`
	GsKitTraceFile     types.String `tfsdk:"gs_kit_trace_file"`
	GsKitTraceFlush    types.Bool   `tfsdk:"gs_kit_trace_flush"`
}

var DmTAMRASTraceTAMTraceFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceTAMTraceSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceTAMTraceTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceTAMTraceComponentsCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceLDAPTraceFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceLDAPTraceSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceLDAPTraceLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceGSKitTraceFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gs_kit_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceGSKitTraceFlushCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gs_kit_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmTAMRASTraceTAMTraceFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceTAMTraceSizeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceTAMTraceTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceTAMTraceComponentsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "tam_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceLDAPTraceFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceLDAPTraceSizeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceLDAPTraceLevelIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ldap_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceGSKitTraceFileIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gs_kit_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmTAMRASTraceGSKitTraceFlushIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "gs_kit_trace_enable",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmTAMRASTraceObjectType = map[string]attr.Type{
	"tam_trace_enable":     types.BoolType,
	"tam_trace_file":       types.StringType,
	"tam_trace_size":       types.Int64Type,
	"tam_trace_type":       types.StringType,
	"tam_trace_components": types.StringType,
	"ldap_trace_enable":    types.BoolType,
	"ldap_trace_file":      types.StringType,
	"ldap_trace_size":      types.Int64Type,
	"ldap_trace_level":     types.Int64Type,
	"gs_kit_trace_enable":  types.BoolType,
	"gs_kit_trace_file":    types.StringType,
	"gs_kit_trace_flush":   types.BoolType,
}
var DmTAMRASTraceObjectDefault = map[string]attr.Value{
	"tam_trace_enable":     types.BoolValue(false),
	"tam_trace_file":       types.StringNull(),
	"tam_trace_size":       types.Int64Value(100),
	"tam_trace_type":       types.StringNull(),
	"tam_trace_components": types.StringValue("*:*.9"),
	"ldap_trace_enable":    types.BoolValue(false),
	"ldap_trace_file":      types.StringNull(),
	"ldap_trace_size":      types.Int64Value(10000),
	"ldap_trace_level":     types.Int64Value(1),
	"gs_kit_trace_enable":  types.BoolValue(false),
	"gs_kit_trace_file":    types.StringNull(),
	"gs_kit_trace_flush":   types.BoolValue(false),
}

func GetDmTAMRASTraceDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmTAMRASTraceDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"tam_trace_enable": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable tracing in the client libraries for the Access Manager client.", "enable-tam-trace", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"tam_trace_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name of the trace file for Access Manager trace messages. The DataPower Gateway writes events to this trace file in the temporary: <i>client-name</i> directory.", "tam-trace-file", "").AddRequiredWhen(DmTAMRASTraceTAMTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceFileIgnoreVal.String()).String,
				Computed:            true,
			},
			"tam_trace_size": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of log entries for the trace file. Enter a value in the range 100 - 1000000. The default value is 100.", "tam-trace-size", "").AddIntegerRange(100, 1000000).AddDefaultValue("100").AddRequiredWhen(DmTAMRASTraceTAMTraceSizeCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceSizeIgnoreVal.String()).String,
				Computed:            true,
			},
			"tam_trace_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the format of trace messages in the file.", "tam-trace-type", "").AddStringEnum("textfile", "utf8file", "xmlfile").AddRequiredWhen(DmTAMRASTraceTAMTraceTypeCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"tam_trace_components": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the components and debug level. The default value is <tt>*:*.9</tt> , which is to trace all components and subcomponents at the highest debug level. Contact IBM Support to set the value needed to help resolve your problem.", "tam-trace-components", "").AddDefaultValue("*:*.9").AddRequiredWhen(DmTAMRASTraceTAMTraceComponentsCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceComponentsIgnoreVal.String()).String,
				Computed:            true,
			},
			"ldap_trace_enable": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable LDAP tracing in the client libraries for the Access Manager client.", "enable-ldap-trace", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"ldap_trace_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name of the trace file for LDAP trace messages. The DataPower Gateway writes events to this trace file in the temporary: <i>client-name</i> directory.", "ldap-trace-file", "").AddRequiredWhen(DmTAMRASTraceLDAPTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceFileIgnoreVal.String()).String,
				Computed:            true,
			},
			"ldap_trace_size": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of characters for the trace file. Enter a value in the range 10000 - 10000000. The default value is 10000.", "ldap-trace-size", "").AddIntegerRange(10000, 10000000).AddDefaultValue("10000").AddRequiredWhen(DmTAMRASTraceLDAPTraceSizeCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceSizeIgnoreVal.String()).String,
				Computed:            true,
			},
			"ldap_trace_level": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the levels for the LDAP trace. Enter a value in the range 1 - 65535. The default value is 1. A value of 65535 enables all levels. Contact IBM Support to set the value needed to help resolve your problem.", "ldap-trace-level", "").AddIntegerRange(1, 65535).AddDefaultValue("1").AddRequiredWhen(DmTAMRASTraceLDAPTraceLevelCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceLevelIgnoreVal.String()).String,
				Computed:            true,
			},
			"gs_kit_trace_enable": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable GSKit tracing in the client libraries for the Access Manager client.", "enable-gskit-trace", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"gs_kit_trace_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name for the trace file for GSKit messages. The DataPower Gateway writes event to this trace file in the temporary: <i>client-name</i> directory.", "gskit-trace-file", "").AddRequiredWhen(DmTAMRASTraceGSKitTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceGSKitTraceFileIgnoreVal.String()).String,
				Computed:            true,
			},
			"gs_kit_trace_flush": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to sets the NOBUFFERING flag to force the trace to be written to the file without buffering.", "gskit-trace-flush", "").AddDefaultValue("false").AddRequiredWhen(DmTAMRASTraceGSKitTraceFlushCondVal.String()).AddNotValidWhen(DmTAMRASTraceGSKitTraceFlushIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	DmTAMRASTraceDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmTAMRASTraceDataSourceSchema
}
func GetDmTAMRASTraceResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmTAMRASTraceResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmTAMRASTraceObjectType,
				DmTAMRASTraceObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"tam_trace_enable": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable tracing in the client libraries for the Access Manager client.", "enable-tam-trace", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_trace_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name of the trace file for Access Manager trace messages. The DataPower Gateway writes events to this trace file in the temporary: <i>client-name</i> directory.", "tam-trace-file", "").AddRequiredWhen(DmTAMRASTraceTAMTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmTAMRASTraceTAMTraceFileCondVal, DmTAMRASTraceTAMTraceFileIgnoreVal, false),
				},
			},
			"tam_trace_size": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of log entries for the trace file. Enter a value in the range 100 - 1000000. The default value is 100.", "tam-trace-size", "").AddIntegerRange(100, 1000000).AddDefaultValue("100").AddRequiredWhen(DmTAMRASTraceTAMTraceSizeCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceSizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(100, 1000000),
					validators.ConditionalRequiredInt64(DmTAMRASTraceTAMTraceSizeCondVal, DmTAMRASTraceTAMTraceSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(100),
			},
			"tam_trace_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the format of trace messages in the file.", "tam-trace-type", "").AddStringEnum("textfile", "utf8file", "xmlfile").AddRequiredWhen(DmTAMRASTraceTAMTraceTypeCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("textfile", "utf8file", "xmlfile"),
					validators.ConditionalRequiredString(DmTAMRASTraceTAMTraceTypeCondVal, DmTAMRASTraceTAMTraceTypeIgnoreVal, false),
				},
			},
			"tam_trace_components": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the components and debug level. The default value is <tt>*:*.9</tt> , which is to trace all components and subcomponents at the highest debug level. Contact IBM Support to set the value needed to help resolve your problem.", "tam-trace-components", "").AddDefaultValue("*:*.9").AddRequiredWhen(DmTAMRASTraceTAMTraceComponentsCondVal.String()).AddNotValidWhen(DmTAMRASTraceTAMTraceComponentsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmTAMRASTraceTAMTraceComponentsCondVal, DmTAMRASTraceTAMTraceComponentsIgnoreVal, true),
				},
				Default: stringdefault.StaticString("*:*.9"),
			},
			"ldap_trace_enable": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable LDAP tracing in the client libraries for the Access Manager client.", "enable-ldap-trace", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_trace_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name of the trace file for LDAP trace messages. The DataPower Gateway writes events to this trace file in the temporary: <i>client-name</i> directory.", "ldap-trace-file", "").AddRequiredWhen(DmTAMRASTraceLDAPTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmTAMRASTraceLDAPTraceFileCondVal, DmTAMRASTraceLDAPTraceFileIgnoreVal, false),
				},
			},
			"ldap_trace_size": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the maximum number of characters for the trace file. Enter a value in the range 10000 - 10000000. The default value is 10000.", "ldap-trace-size", "").AddIntegerRange(10000, 10000000).AddDefaultValue("10000").AddRequiredWhen(DmTAMRASTraceLDAPTraceSizeCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceSizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10000, 10000000),
					validators.ConditionalRequiredInt64(DmTAMRASTraceLDAPTraceSizeCondVal, DmTAMRASTraceLDAPTraceSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(10000),
			},
			"ldap_trace_level": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the levels for the LDAP trace. Enter a value in the range 1 - 65535. The default value is 1. A value of 65535 enables all levels. Contact IBM Support to set the value needed to help resolve your problem.", "ldap-trace-level", "").AddIntegerRange(1, 65535).AddDefaultValue("1").AddRequiredWhen(DmTAMRASTraceLDAPTraceLevelCondVal.String()).AddNotValidWhen(DmTAMRASTraceLDAPTraceLevelIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(DmTAMRASTraceLDAPTraceLevelCondVal, DmTAMRASTraceLDAPTraceLevelIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"gs_kit_trace_enable": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to enable GSKit tracing in the client libraries for the Access Manager client.", "enable-gskit-trace", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"gs_kit_trace_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the base name for the trace file for GSKit messages. The DataPower Gateway writes event to this trace file in the temporary: <i>client-name</i> directory.", "gskit-trace-file", "").AddRequiredWhen(DmTAMRASTraceGSKitTraceFileCondVal.String()).AddNotValidWhen(DmTAMRASTraceGSKitTraceFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmTAMRASTraceGSKitTraceFileCondVal, DmTAMRASTraceGSKitTraceFileIgnoreVal, false),
				},
			},
			"gs_kit_trace_flush": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether to sets the NOBUFFERING flag to force the trace to be written to the file without buffering.", "gskit-trace-flush", "").AddDefaultValue("false").AddRequiredWhen(DmTAMRASTraceGSKitTraceFlushCondVal.String()).AddNotValidWhen(DmTAMRASTraceGSKitTraceFlushIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmTAMRASTraceResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmTAMRASTraceResourceSchema.Required = true
	} else {
		DmTAMRASTraceResourceSchema.Optional = true
		DmTAMRASTraceResourceSchema.Computed = true
	}
	return DmTAMRASTraceResourceSchema
}

func (data DmTAMRASTrace) IsNull() bool {
	if !data.TamTraceEnable.IsNull() {
		return false
	}
	if !data.TamTraceFile.IsNull() {
		return false
	}
	if !data.TamTraceSize.IsNull() {
		return false
	}
	if !data.TamTraceType.IsNull() {
		return false
	}
	if !data.TamTraceComponents.IsNull() {
		return false
	}
	if !data.LdapTraceEnable.IsNull() {
		return false
	}
	if !data.LdapTraceFile.IsNull() {
		return false
	}
	if !data.LdapTraceSize.IsNull() {
		return false
	}
	if !data.LdapTraceLevel.IsNull() {
		return false
	}
	if !data.GsKitTraceEnable.IsNull() {
		return false
	}
	if !data.GsKitTraceFile.IsNull() {
		return false
	}
	if !data.GsKitTraceFlush.IsNull() {
		return false
	}
	return true
}

func (data DmTAMRASTrace) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.TamTraceEnable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMTraceEnable`, tfutils.StringFromBool(data.TamTraceEnable, ""))
	}
	if !data.TamTraceFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMTraceFile`, data.TamTraceFile.ValueString())
	}
	if !data.TamTraceSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMTraceSize`, data.TamTraceSize.ValueInt64())
	}
	if !data.TamTraceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMTraceType`, data.TamTraceType.ValueString())
	}
	if !data.TamTraceComponents.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMTraceComponents`, data.TamTraceComponents.ValueString())
	}
	if !data.LdapTraceEnable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPTraceEnable`, tfutils.StringFromBool(data.LdapTraceEnable, ""))
	}
	if !data.LdapTraceFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPTraceFile`, data.LdapTraceFile.ValueString())
	}
	if !data.LdapTraceSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPTraceSize`, data.LdapTraceSize.ValueInt64())
	}
	if !data.LdapTraceLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPTraceLevel`, data.LdapTraceLevel.ValueInt64())
	}
	if !data.GsKitTraceEnable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GSKitTraceEnable`, tfutils.StringFromBool(data.GsKitTraceEnable, ""))
	}
	if !data.GsKitTraceFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GSKitTraceFile`, data.GsKitTraceFile.ValueString())
	}
	if !data.GsKitTraceFlush.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GSKitTraceFlush`, tfutils.StringFromBool(data.GsKitTraceFlush, ""))
	}
	return body
}

func (data *DmTAMRASTrace) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TAMTraceEnable`); value.Exists() {
		data.TamTraceEnable = tfutils.BoolFromString(value.String())
	} else {
		data.TamTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMTraceFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMTraceSize`); value.Exists() {
		data.TamTraceSize = types.Int64Value(value.Int())
	} else {
		data.TamTraceSize = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `TAMTraceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamTraceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamTraceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMTraceComponents`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamTraceComponents = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamTraceComponents = types.StringValue("*:*.9")
	}
	if value := res.Get(pathRoot + `LDAPTraceEnable`); value.Exists() {
		data.LdapTraceEnable = tfutils.BoolFromString(value.String())
	} else {
		data.LdapTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPTraceFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPTraceSize`); value.Exists() {
		data.LdapTraceSize = types.Int64Value(value.Int())
	} else {
		data.LdapTraceSize = types.Int64Value(10000)
	}
	if value := res.Get(pathRoot + `LDAPTraceLevel`); value.Exists() {
		data.LdapTraceLevel = types.Int64Value(value.Int())
	} else {
		data.LdapTraceLevel = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `GSKitTraceEnable`); value.Exists() {
		data.GsKitTraceEnable = tfutils.BoolFromString(value.String())
	} else {
		data.GsKitTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GSKitTraceFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GsKitTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GsKitTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `GSKitTraceFlush`); value.Exists() {
		data.GsKitTraceFlush = tfutils.BoolFromString(value.String())
	} else {
		data.GsKitTraceFlush = types.BoolNull()
	}
}

func (data *DmTAMRASTrace) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TAMTraceEnable`); value.Exists() && !data.TamTraceEnable.IsNull() {
		data.TamTraceEnable = tfutils.BoolFromString(value.String())
	} else if data.TamTraceEnable.ValueBool() {
		data.TamTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TAMTraceFile`); value.Exists() && !data.TamTraceFile.IsNull() {
		data.TamTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMTraceSize`); value.Exists() && !data.TamTraceSize.IsNull() {
		data.TamTraceSize = types.Int64Value(value.Int())
	} else if data.TamTraceSize.ValueInt64() != 100 {
		data.TamTraceSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TAMTraceType`); value.Exists() && !data.TamTraceType.IsNull() {
		data.TamTraceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamTraceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMTraceComponents`); value.Exists() && !data.TamTraceComponents.IsNull() {
		data.TamTraceComponents = tfutils.ParseStringFromGJSON(value)
	} else if data.TamTraceComponents.ValueString() != "*:*.9" {
		data.TamTraceComponents = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPTraceEnable`); value.Exists() && !data.LdapTraceEnable.IsNull() {
		data.LdapTraceEnable = tfutils.BoolFromString(value.String())
	} else if data.LdapTraceEnable.ValueBool() {
		data.LdapTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LDAPTraceFile`); value.Exists() && !data.LdapTraceFile.IsNull() {
		data.LdapTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPTraceSize`); value.Exists() && !data.LdapTraceSize.IsNull() {
		data.LdapTraceSize = types.Int64Value(value.Int())
	} else if data.LdapTraceSize.ValueInt64() != 10000 {
		data.LdapTraceSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LDAPTraceLevel`); value.Exists() && !data.LdapTraceLevel.IsNull() {
		data.LdapTraceLevel = types.Int64Value(value.Int())
	} else if data.LdapTraceLevel.ValueInt64() != 1 {
		data.LdapTraceLevel = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GSKitTraceEnable`); value.Exists() && !data.GsKitTraceEnable.IsNull() {
		data.GsKitTraceEnable = tfutils.BoolFromString(value.String())
	} else if data.GsKitTraceEnable.ValueBool() {
		data.GsKitTraceEnable = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GSKitTraceFile`); value.Exists() && !data.GsKitTraceFile.IsNull() {
		data.GsKitTraceFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GsKitTraceFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `GSKitTraceFlush`); value.Exists() && !data.GsKitTraceFlush.IsNull() {
		data.GsKitTraceFlush = tfutils.BoolFromString(value.String())
	} else if data.GsKitTraceFlush.ValueBool() {
		data.GsKitTraceFlush = types.BoolNull()
	}
}
