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

type DmSSLHostnameValidationFlags struct {
	X509CheckFlagAlwaysCheckSubject    types.Bool `tfsdk:"x509_check_flag_always_check_subject"`
	X509CheckFlagNoWildcards           types.Bool `tfsdk:"x509_check_flag_no_wildcards"`
	X509CheckFlagNoPartialWildcards    types.Bool `tfsdk:"x509_check_flag_no_partial_wildcards"`
	X509CheckFlagMultiLabelWildcards   types.Bool `tfsdk:"x509_check_flag_multi_label_wildcards"`
	X509CheckFlagSingleLabelSubdomains types.Bool `tfsdk:"x509_check_flag_single_label_subdomains"`
}

var DmSSLHostnameValidationFlagsObjectType = map[string]attr.Type{
	"x509_check_flag_always_check_subject":    types.BoolType,
	"x509_check_flag_no_wildcards":            types.BoolType,
	"x509_check_flag_no_partial_wildcards":    types.BoolType,
	"x509_check_flag_multi_label_wildcards":   types.BoolType,
	"x509_check_flag_single_label_subdomains": types.BoolType,
}
var DmSSLHostnameValidationFlagsObjectDefault = map[string]attr.Value{
	"x509_check_flag_always_check_subject":    types.BoolValue(false),
	"x509_check_flag_no_wildcards":            types.BoolValue(false),
	"x509_check_flag_no_partial_wildcards":    types.BoolValue(false),
	"x509_check_flag_multi_label_wildcards":   types.BoolValue(false),
	"x509_check_flag_single_label_subdomains": types.BoolValue(false),
}

func GetDmSSLHostnameValidationFlagsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmSSLHostnameValidationFlagsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"x509_check_flag_always_check_subject": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"x509_check_flag_no_wildcards": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_NO_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"x509_check_flag_no_partial_wildcards": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"x509_check_flag_multi_label_wildcards": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"x509_check_flag_single_label_subdomains": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmSSLHostnameValidationFlagsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmSSLHostnameValidationFlagsDataSourceSchema
}
func GetDmSSLHostnameValidationFlagsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmSSLHostnameValidationFlagsResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmSSLHostnameValidationFlagsObjectType,
				DmSSLHostnameValidationFlagsObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"x509_check_flag_always_check_subject": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"x509_check_flag_no_wildcards": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_NO_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"x509_check_flag_no_partial_wildcards": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"x509_check_flag_multi_label_wildcards": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"x509_check_flag_single_label_subdomains": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmSSLHostnameValidationFlagsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmSSLHostnameValidationFlagsResourceSchema.Required = true
	} else {
		DmSSLHostnameValidationFlagsResourceSchema.Optional = true
		DmSSLHostnameValidationFlagsResourceSchema.Computed = true
	}
	return DmSSLHostnameValidationFlagsResourceSchema
}

func (data DmSSLHostnameValidationFlags) IsNull() bool {
	if !data.X509CheckFlagAlwaysCheckSubject.IsNull() {
		return false
	}
	if !data.X509CheckFlagNoWildcards.IsNull() {
		return false
	}
	if !data.X509CheckFlagNoPartialWildcards.IsNull() {
		return false
	}
	if !data.X509CheckFlagMultiLabelWildcards.IsNull() {
		return false
	}
	if !data.X509CheckFlagSingleLabelSubdomains.IsNull() {
		return false
	}
	return true
}

func (data DmSSLHostnameValidationFlags) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.X509CheckFlagAlwaysCheckSubject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT`, tfutils.StringFromBool(data.X509CheckFlagAlwaysCheckSubject, ""))
	}
	if !data.X509CheckFlagNoWildcards.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`X509_CHECK_FLAG_NO_WILDCARDS`, tfutils.StringFromBool(data.X509CheckFlagNoWildcards, ""))
	}
	if !data.X509CheckFlagNoPartialWildcards.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS`, tfutils.StringFromBool(data.X509CheckFlagNoPartialWildcards, ""))
	}
	if !data.X509CheckFlagMultiLabelWildcards.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS`, tfutils.StringFromBool(data.X509CheckFlagMultiLabelWildcards, ""))
	}
	if !data.X509CheckFlagSingleLabelSubdomains.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS`, tfutils.StringFromBool(data.X509CheckFlagSingleLabelSubdomains, ""))
	}
	return body
}

func (data *DmSSLHostnameValidationFlags) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT`); value.Exists() {
		data.X509CheckFlagAlwaysCheckSubject = tfutils.BoolFromString(value.String())
	} else {
		data.X509CheckFlagAlwaysCheckSubject = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_NO_WILDCARDS`); value.Exists() {
		data.X509CheckFlagNoWildcards = tfutils.BoolFromString(value.String())
	} else {
		data.X509CheckFlagNoWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS`); value.Exists() {
		data.X509CheckFlagNoPartialWildcards = tfutils.BoolFromString(value.String())
	} else {
		data.X509CheckFlagNoPartialWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS`); value.Exists() {
		data.X509CheckFlagMultiLabelWildcards = tfutils.BoolFromString(value.String())
	} else {
		data.X509CheckFlagMultiLabelWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS`); value.Exists() {
		data.X509CheckFlagSingleLabelSubdomains = tfutils.BoolFromString(value.String())
	} else {
		data.X509CheckFlagSingleLabelSubdomains = types.BoolNull()
	}
}

func (data *DmSSLHostnameValidationFlags) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT`); value.Exists() && !data.X509CheckFlagAlwaysCheckSubject.IsNull() {
		data.X509CheckFlagAlwaysCheckSubject = tfutils.BoolFromString(value.String())
	} else if data.X509CheckFlagAlwaysCheckSubject.ValueBool() {
		data.X509CheckFlagAlwaysCheckSubject = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_NO_WILDCARDS`); value.Exists() && !data.X509CheckFlagNoWildcards.IsNull() {
		data.X509CheckFlagNoWildcards = tfutils.BoolFromString(value.String())
	} else if data.X509CheckFlagNoWildcards.ValueBool() {
		data.X509CheckFlagNoWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS`); value.Exists() && !data.X509CheckFlagNoPartialWildcards.IsNull() {
		data.X509CheckFlagNoPartialWildcards = tfutils.BoolFromString(value.String())
	} else if data.X509CheckFlagNoPartialWildcards.ValueBool() {
		data.X509CheckFlagNoPartialWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS`); value.Exists() && !data.X509CheckFlagMultiLabelWildcards.IsNull() {
		data.X509CheckFlagMultiLabelWildcards = tfutils.BoolFromString(value.String())
	} else if data.X509CheckFlagMultiLabelWildcards.ValueBool() {
		data.X509CheckFlagMultiLabelWildcards = types.BoolNull()
	}
	if value := res.Get(pathRoot + `X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS`); value.Exists() && !data.X509CheckFlagSingleLabelSubdomains.IsNull() {
		data.X509CheckFlagSingleLabelSubdomains = tfutils.BoolFromString(value.String())
	} else if data.X509CheckFlagSingleLabelSubdomains.ValueBool() {
		data.X509CheckFlagSingleLabelSubdomains = types.BoolNull()
	}
}
