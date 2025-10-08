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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmExternalAttachedPolicy struct {
	ExternalAttachWsdlComponentType                   types.String `tfsdk:"external_attach_wsdl_component_type"`
	ExternalAttachWsdlComponentValue                  types.String `tfsdk:"external_attach_wsdl_component_value"`
	ExternalAttachPolicyUrl                           types.String `tfsdk:"external_attach_policy_url"`
	ExternalAttachPolicyFragmentId                    types.String `tfsdk:"external_attach_policy_fragment_id"`
	ExternalAttachMessageContentFilter                types.String `tfsdk:"external_attach_message_content_filter"`
	ExternalAttachMessageContentFilterServiceProvider types.String `tfsdk:"external_attach_message_content_filter_service_provider"`
}

var DmExternalAttachedPolicyExternalAttachWSDLComponentValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "external_attach_wsdl_component_type",
	AttrType:    "String",
	AttrDefault: "service",
	Value:       []string{"rest"},
}

var DmExternalAttachedPolicyExternalAttachWSDLComponentValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "external_attach_wsdl_component_type",
	AttrType:    "String",
	AttrDefault: "service",
	Value:       []string{"rest"},
}

var DmExternalAttachedPolicyExternalAttachPolicyFragmentIDIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "external_attach_wsdl_component_type",
	AttrType:    "String",
	AttrDefault: "service",
	Value:       []string{"rest"},
}

var DmExternalAttachedPolicyObjectType = map[string]attr.Type{
	"external_attach_wsdl_component_type":                     types.StringType,
	"external_attach_wsdl_component_value":                    types.StringType,
	"external_attach_policy_url":                              types.StringType,
	"external_attach_policy_fragment_id":                      types.StringType,
	"external_attach_message_content_filter":                  types.StringType,
	"external_attach_message_content_filter_service_provider": types.StringType,
}
var DmExternalAttachedPolicyObjectDefault = map[string]attr.Value{
	"external_attach_wsdl_component_type":                     types.StringValue("service"),
	"external_attach_wsdl_component_value":                    types.StringNull(),
	"external_attach_policy_url":                              types.StringNull(),
	"external_attach_policy_fragment_id":                      types.StringNull(),
	"external_attach_message_content_filter":                  types.StringNull(),
	"external_attach_message_content_filter_service_provider": types.StringNull(),
}

func GetDmExternalAttachedPolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmExternalAttachedPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"external_attach_wsdl_component_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select a type of Component",
				Computed:            true,
			},
			"external_attach_wsdl_component_value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the qname of a WSDL component formatted {ns}ncname",
				Computed:            true,
			},
			"external_attach_policy_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select a document containing policy to be attached",
				Computed:            true,
			},
			"external_attach_policy_fragment_id": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Matches Fragment Identifier",
				Computed:            true,
			},
			"external_attach_message_content_filter": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the name of the message content filters object that specifies the service consumer",
				Computed:            true,
			},
			"external_attach_message_content_filter_service_provider": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Enter the name of the message content filters object that specifies the service provider",
				Computed:            true,
			},
		},
	}
	return DmExternalAttachedPolicyDataSourceSchema
}
func GetDmExternalAttachedPolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmExternalAttachedPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"external_attach_wsdl_component_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a type of Component", "", "").AddStringEnum("service", "port", "fragmentid", "rest").AddDefaultValue("service").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("service", "port", "fragmentid", "rest"),
				},
				Default: stringdefault.StaticString("service"),
			},
			"external_attach_wsdl_component_value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the qname of a WSDL component formatted {ns}ncname", "", "").AddRequiredWhen(DmExternalAttachedPolicyExternalAttachWSDLComponentValueCondVal.String()).AddNotValidWhen(DmExternalAttachedPolicyExternalAttachWSDLComponentValueIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmExternalAttachedPolicyExternalAttachWSDLComponentValueCondVal, DmExternalAttachedPolicyExternalAttachWSDLComponentValueIgnoreVal, false),
				},
			},
			"external_attach_policy_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a document containing policy to be attached", "", "").String,
				Required:            true,
			},
			"external_attach_policy_fragment_id": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").AddNotValidWhen(DmExternalAttachedPolicyExternalAttachPolicyFragmentIDIgnoreVal.String()).String,
				Optional:            true,
			},
			"external_attach_message_content_filter": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the message content filters object that specifies the service consumer", "", "message_content_filters").String,
				Optional:            true,
			},
			"external_attach_message_content_filter_service_provider": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of the message content filters object that specifies the service provider", "", "message_content_filters").String,
				Optional:            true,
			},
		},
	}
	return DmExternalAttachedPolicyResourceSchema
}

func (data DmExternalAttachedPolicy) IsNull() bool {
	if !data.ExternalAttachWsdlComponentType.IsNull() {
		return false
	}
	if !data.ExternalAttachWsdlComponentValue.IsNull() {
		return false
	}
	if !data.ExternalAttachPolicyUrl.IsNull() {
		return false
	}
	if !data.ExternalAttachPolicyFragmentId.IsNull() {
		return false
	}
	if !data.ExternalAttachMessageContentFilter.IsNull() {
		return false
	}
	if !data.ExternalAttachMessageContentFilterServiceProvider.IsNull() {
		return false
	}
	return true
}

func (data DmExternalAttachedPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ExternalAttachWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachWSDLComponentType`, data.ExternalAttachWsdlComponentType.ValueString())
	}
	if !data.ExternalAttachWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachWSDLComponentValue`, data.ExternalAttachWsdlComponentValue.ValueString())
	}
	if !data.ExternalAttachPolicyUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachPolicyURL`, data.ExternalAttachPolicyUrl.ValueString())
	}
	if !data.ExternalAttachPolicyFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachPolicyFragmentID`, data.ExternalAttachPolicyFragmentId.ValueString())
	}
	if !data.ExternalAttachMessageContentFilter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachMessageContentFilter`, data.ExternalAttachMessageContentFilter.ValueString())
	}
	if !data.ExternalAttachMessageContentFilterServiceProvider.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAttachMessageContentFilterServiceProvider`, data.ExternalAttachMessageContentFilterServiceProvider.ValueString())
	}
	return body
}

func (data *DmExternalAttachedPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ExternalAttachWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachWsdlComponentType = types.StringValue("service")
	}
	if value := res.Get(pathRoot + `ExternalAttachWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachPolicyURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachPolicyUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachPolicyUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachPolicyFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachPolicyFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachPolicyFragmentId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachMessageContentFilter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachMessageContentFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachMessageContentFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachMessageContentFilterServiceProvider`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAttachMessageContentFilterServiceProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachMessageContentFilterServiceProvider = types.StringNull()
	}
}

func (data *DmExternalAttachedPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ExternalAttachWSDLComponentType`); value.Exists() && !data.ExternalAttachWsdlComponentType.IsNull() {
		data.ExternalAttachWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.ExternalAttachWsdlComponentType.ValueString() != "service" {
		data.ExternalAttachWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachWSDLComponentValue`); value.Exists() && !data.ExternalAttachWsdlComponentValue.IsNull() {
		data.ExternalAttachWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachPolicyURL`); value.Exists() && !data.ExternalAttachPolicyUrl.IsNull() {
		data.ExternalAttachPolicyUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachPolicyUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachPolicyFragmentID`); value.Exists() && !data.ExternalAttachPolicyFragmentId.IsNull() {
		data.ExternalAttachPolicyFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachPolicyFragmentId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachMessageContentFilter`); value.Exists() && !data.ExternalAttachMessageContentFilter.IsNull() {
		data.ExternalAttachMessageContentFilter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachMessageContentFilter = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAttachMessageContentFilterServiceProvider`); value.Exists() && !data.ExternalAttachMessageContentFilterServiceProvider.IsNull() {
		data.ExternalAttachMessageContentFilterServiceProvider = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAttachMessageContentFilterServiceProvider = types.StringNull()
	}
}
