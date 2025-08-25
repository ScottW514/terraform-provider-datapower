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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type APISecurityHTTPScheme struct {
	Id                         types.String                `tfsdk:"id"`
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	Scheme                     types.String                `tfsdk:"scheme"`
	BearerFormat               types.String                `tfsdk:"bearer_format"`
	BearerValidationMethod     types.String                `tfsdk:"bearer_validation_method"`
	BearerValidationEndpoint   types.String                `tfsdk:"bearer_validation_endpoint"`
	BearerValidationTlsProfile types.String                `tfsdk:"bearer_validation_tls_profile"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APISecurityHTTPSchemeBearerValidationMethodCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "scheme",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"bearer"},
}
var APISecurityHTTPSchemeBearerValidationEndpointCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "scheme",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"bearer"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "bearer_validation_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"external-url"},
		},
	},
}

var APISecurityHTTPSchemeObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"scheme":                        types.StringType,
	"bearer_format":                 types.StringType,
	"bearer_validation_method":      types.StringType,
	"bearer_validation_endpoint":    types.StringType,
	"bearer_validation_tls_profile": types.StringType,
	"dependency_actions":            actions.ActionsListType,
}

func (data APISecurityHTTPScheme) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISecurityHTTPScheme"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISecurityHTTPScheme) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Scheme.IsNull() {
		return false
	}
	if !data.BearerFormat.IsNull() {
		return false
	}
	if !data.BearerValidationMethod.IsNull() {
		return false
	}
	if !data.BearerValidationEndpoint.IsNull() {
		return false
	}
	if !data.BearerValidationTlsProfile.IsNull() {
		return false
	}
	return true
}

func (data APISecurityHTTPScheme) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Scheme.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Scheme`, data.Scheme.ValueString())
	}
	if !data.BearerFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BearerFormat`, data.BearerFormat.ValueString())
	}
	if !data.BearerValidationMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BearerValidationMethod`, data.BearerValidationMethod.ValueString())
	}
	if !data.BearerValidationEndpoint.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BearerValidationEndpoint`, data.BearerValidationEndpoint.ValueString())
	}
	if !data.BearerValidationTlsProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BearerValidationTLSProfile`, data.BearerValidationTlsProfile.ValueString())
	}
	return body
}

func (data *APISecurityHTTPScheme) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scheme`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scheme = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scheme = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BearerFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BearerValidationMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationEndpoint`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BearerValidationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationTLSProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BearerValidationTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationTlsProfile = types.StringNull()
	}
}

func (data *APISecurityHTTPScheme) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scheme`); value.Exists() && !data.Scheme.IsNull() {
		data.Scheme = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scheme = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerFormat`); value.Exists() && !data.BearerFormat.IsNull() {
		data.BearerFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationMethod`); value.Exists() && !data.BearerValidationMethod.IsNull() {
		data.BearerValidationMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationEndpoint`); value.Exists() && !data.BearerValidationEndpoint.IsNull() {
		data.BearerValidationEndpoint = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationEndpoint = types.StringNull()
	}
	if value := res.Get(pathRoot + `BearerValidationTLSProfile`); value.Exists() && !data.BearerValidationTlsProfile.IsNull() {
		data.BearerValidationTlsProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BearerValidationTlsProfile = types.StringNull()
	}
}
