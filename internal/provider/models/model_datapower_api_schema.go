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

type APISchema struct {
	Id                         types.String                `tfsdk:"id"`
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	JsonSchema                 types.String                `tfsdk:"json_schema"`
	GraphqlSchema              types.String                `tfsdk:"graphql_schema"`
	XmlType                    types.String                `tfsdk:"xml_type"`
	XmlValidationMode          types.String                `tfsdk:"xml_validation_mode"`
	XmlSchemaUrl               types.String                `tfsdk:"xml_schema_url"`
	WsdlSchemaUrl              types.String                `tfsdk:"wsdl_schema_url"`
	WsdlPortQName              types.String                `tfsdk:"wsdl_port_q_name"`
	WsdlOperationName          types.String                `tfsdk:"wsdl_operation_name"`
	WsdlMessageDirectionOrName types.String                `tfsdk:"wsdl_message_direction_or_name"`
	WsdlAttachmentPart         types.String                `tfsdk:"wsdl_attachment_part"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget             types.String                `tfsdk:"provider_target"`
}

var APISchemaXMLValidationModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xml"},
}

var APISchemaXMLSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xml"},
}

var APISchemaWSDLSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"wsdl"},
}

var APISchemaWSDLPortQNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"wsdl"},
}

var APISchemaWSDLOperationNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"wsdl"},
}

var APISchemaWSDLMessageDirectionOrNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"wsdl"},
}

var APISchemaWSDLAttachmentPartIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "xml_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"wsdl"},
}

var APISchemaObjectType = map[string]attr.Type{
	"provider_target":                types.StringType,
	"id":                             types.StringType,
	"app_domain":                     types.StringType,
	"user_summary":                   types.StringType,
	"json_schema":                    types.StringType,
	"graphql_schema":                 types.StringType,
	"xml_type":                       types.StringType,
	"xml_validation_mode":            types.StringType,
	"xml_schema_url":                 types.StringType,
	"wsdl_schema_url":                types.StringType,
	"wsdl_port_q_name":               types.StringType,
	"wsdl_operation_name":            types.StringType,
	"wsdl_message_direction_or_name": types.StringType,
	"wsdl_attachment_part":           types.StringType,
	"dependency_actions":             actions.ActionsListType,
}

func (data APISchema) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APISchema"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APISchema) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.JsonSchema.IsNull() {
		return false
	}
	if !data.GraphqlSchema.IsNull() {
		return false
	}
	if !data.XmlType.IsNull() {
		return false
	}
	if !data.XmlValidationMode.IsNull() {
		return false
	}
	if !data.XmlSchemaUrl.IsNull() {
		return false
	}
	if !data.WsdlSchemaUrl.IsNull() {
		return false
	}
	if !data.WsdlPortQName.IsNull() {
		return false
	}
	if !data.WsdlOperationName.IsNull() {
		return false
	}
	if !data.WsdlMessageDirectionOrName.IsNull() {
		return false
	}
	if !data.WsdlAttachmentPart.IsNull() {
		return false
	}
	return true
}

func (data APISchema) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.JsonSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONSchema`, data.JsonSchema.ValueString())
	}
	if !data.GraphqlSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GraphQLSchema`, data.GraphqlSchema.ValueString())
	}
	if !data.XmlType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLType`, data.XmlType.ValueString())
	}
	if !data.XmlValidationMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLValidationMode`, data.XmlValidationMode.ValueString())
	}
	if !data.XmlSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLSchemaURL`, data.XmlSchemaUrl.ValueString())
	}
	if !data.WsdlSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLSchemaURL`, data.WsdlSchemaUrl.ValueString())
	}
	if !data.WsdlPortQName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLPortQName`, data.WsdlPortQName.ValueString())
	}
	if !data.WsdlOperationName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLOperationName`, data.WsdlOperationName.ValueString())
	}
	if !data.WsdlMessageDirectionOrName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLMessageDirectionOrName`, data.WsdlMessageDirectionOrName.ValueString())
	}
	if !data.WsdlAttachmentPart.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLAttachmentPart`, data.WsdlAttachmentPart.ValueString())
	}
	return body
}

func (data *APISchema) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JSONSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JsonSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GraphqlSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GraphqlSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlType = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLValidationMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLPortQName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlPortQName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlPortQName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLOperationName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlOperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlOperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLMessageDirectionOrName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlMessageDirectionOrName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlMessageDirectionOrName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLAttachmentPart`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlAttachmentPart = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAttachmentPart = types.StringNull()
	}
}

func (data *APISchema) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JSONSchema`); value.Exists() && !data.JsonSchema.IsNull() {
		data.JsonSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSchema`); value.Exists() && !data.GraphqlSchema.IsNull() {
		data.GraphqlSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GraphqlSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLType`); value.Exists() && !data.XmlType.IsNull() {
		data.XmlType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlType = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLValidationMode`); value.Exists() && !data.XmlValidationMode.IsNull() {
		data.XmlValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLSchemaURL`); value.Exists() && !data.XmlSchemaUrl.IsNull() {
		data.XmlSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLSchemaURL`); value.Exists() && !data.WsdlSchemaUrl.IsNull() {
		data.WsdlSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLPortQName`); value.Exists() && !data.WsdlPortQName.IsNull() {
		data.WsdlPortQName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlPortQName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLOperationName`); value.Exists() && !data.WsdlOperationName.IsNull() {
		data.WsdlOperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlOperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLMessageDirectionOrName`); value.Exists() && !data.WsdlMessageDirectionOrName.IsNull() {
		data.WsdlMessageDirectionOrName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlMessageDirectionOrName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLAttachmentPart`); value.Exists() && !data.WsdlAttachmentPart.IsNull() {
		data.WsdlAttachmentPart = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAttachmentPart = types.StringNull()
	}
}
