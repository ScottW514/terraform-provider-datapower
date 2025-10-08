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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPMapResource struct {
	MrMethod               types.String `tfsdk:"mr_method"`
	MrCustomUrl            types.String `tfsdk:"mr_custom_url"`
	MrMapUrl               types.String `tfsdk:"mr_map_url"`
	MrMapXpath             types.String `tfsdk:"mr_map_xpath"`
	MrTamMap               types.String `tfsdk:"mr_tam_map"`
	MrTamInstancePrefix    types.String `tfsdk:"mr_tam_instance_prefix"`
	MrTamWebSealDynUrlFile types.String `tfsdk:"mr_tam_web_seal_dyn_url_file"`
}

var DmAAAPMapResourceMRCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mr_method",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"custom"},
}

var DmAAAPMapResourceMRCustomURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPMapResourceMRMapURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mr_method",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"xmlfile"},
}

var DmAAAPMapResourceMRMapURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPMapResourceMRMapXPathCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mr_method",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"xpath"},
}

var DmAAAPMapResourceMRMapXPathIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPMapResourceMRTAMMapCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mr_method",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"tivoli"},
}

var DmAAAPMapResourceMRTAMMapIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPMapResourceMRTAMInstancePrefixCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mr_method",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"tivoli"},
}

var DmAAAPMapResourceMRTAMInstancePrefixIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmAAAPMapResourceMRTAMWebSEALDynURLFileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-not",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "mr_method",
					AttrType:    "String",
					AttrDefault: "none",
					Value:       []string{"tivoli"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "mr_tam_map",
					AttrType:    "String",
					AttrDefault: "WebSEAL",
					Value:       []string{"WebSEAL"},
				},
			},
		},
	},
}

var DmAAAPMapResourceObjectType = map[string]attr.Type{
	"mr_method":                    types.StringType,
	"mr_custom_url":                types.StringType,
	"mr_map_url":                   types.StringType,
	"mr_map_xpath":                 types.StringType,
	"mr_tam_map":                   types.StringType,
	"mr_tam_instance_prefix":       types.StringType,
	"mr_tam_web_seal_dyn_url_file": types.StringType,
}
var DmAAAPMapResourceObjectDefault = map[string]attr.Value{
	"mr_method":                    types.StringValue("none"),
	"mr_custom_url":                types.StringNull(),
	"mr_map_url":                   types.StringNull(),
	"mr_map_xpath":                 types.StringNull(),
	"mr_tam_map":                   types.StringValue("WebSEAL"),
	"mr_tam_instance_prefix":       types.StringNull(),
	"mr_tam_web_seal_dyn_url_file": types.StringNull(),
}

func GetDmAAAPMapResourceDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmAAAPMapResourceDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"mr_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to map resources.", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "tivoli").AddDefaultValue("none").String,
				Computed:            true,
			},
			"mr_custom_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the DynURL file. The file must be in the <tt>local:</tt> or <tt>store:</tt> directory. When configured and an entry matches a request, the DynURL output replaces the output of the extracted resource string that is added to the prefix.", "custom-url", "").AddRequiredWhen(DmAAAPMapResourceMRCustomURLCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRCustomURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"mr_map_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA XML file that defines how to map resources.", "xmlfile-url", "").AddRequiredWhen(DmAAAPMapResourceMRMapURLCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRMapURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"mr_map_xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the extracted resource.", "xpath", "").AddRequiredWhen(DmAAAPMapResourceMRMapXPathCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRMapXPathIgnoreVal.String()).String,
				Computed:            true,
			},
			"mr_tam_map": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Access Manager style-mapping resource. Access Manager organizes resources into a hierarchical protected object space. The default value is WebSEAL.", "tam-mapping", "").AddStringEnum("TFIM", "TAMBI", "WebSEAL", "Custom").AddDefaultValue("WebSEAL").AddRequiredWhen(DmAAAPMapResourceMRTAMMapCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRTAMMapIgnoreVal.String()).String,
				Computed:            true,
			},
			"mr_tam_instance_prefix": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the prefix to add based on the mapping style. Each protected object space has a defined convention. These different options help to map extracted resources to a resource string that follow a Tivoli naming convention.</p><ul><li>For the custom mapping style, specify the string to add as the prefix. In other words, use the <i>prefix</i> format.</li><li>For the TAMBI mapping style, specify the queue manager and queue, and separate the queue manager and the queue with a forward slash. In other words, use the <tt>/PDMQ/ <i>prefix</i></tt> format.</li><li>For the TFIM mapping style, specify the name of the Federated Identity Manager domain. In other words, use the <tt>/itfim-wssm/wssm-default/ <i>prefix</i></tt> format.</li><li>For the WebSEAL mapping style, specify the name of the WebSEAL instance. In other words, use the <tt>/WebSEAL/ <i>prefix</i></tt> format.</li></ul>", "tam-prefix", "").AddRequiredWhen(DmAAAPMapResourceMRTAMInstancePrefixCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRTAMInstancePrefixIgnoreVal.String()).String,
				Computed:            true,
			},
			"mr_tam_web_seal_dyn_url_file": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the DynURL file. The file must be in the <tt>local:</tt> or <tt>store:</tt> directory. When configured and an entry matches a request, the DynURL output replaces the output of the extracted resource string that is added to the prefix.", "webseal-dynurl-file", "").AddNotValidWhen(DmAAAPMapResourceMRTAMWebSEALDynURLFileIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	DmAAAPMapResourceDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPMapResourceDataSourceSchema
}
func GetDmAAAPMapResourceResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmAAAPMapResourceResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmAAAPMapResourceObjectType,
				DmAAAPMapResourceObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"mr_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to map resources.", "method", "").AddStringEnum("none", "custom", "xmlfile", "xpath", "tivoli").AddDefaultValue("none").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "custom", "xmlfile", "xpath", "tivoli"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"mr_custom_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the DynURL file. The file must be in the <tt>local:</tt> or <tt>store:</tt> directory. When configured and an entry matches a request, the DynURL output replaces the output of the extracted resource string that is added to the prefix.", "custom-url", "").AddRequiredWhen(DmAAAPMapResourceMRCustomURLCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRCustomURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPMapResourceMRCustomURLCondVal, DmAAAPMapResourceMRCustomURLIgnoreVal, false),
				},
			},
			"mr_map_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the AAA XML file that defines how to map resources.", "xmlfile-url", "").AddRequiredWhen(DmAAAPMapResourceMRMapURLCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRMapURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPMapResourceMRMapURLCondVal, DmAAAPMapResourceMRMapURLIgnoreVal, false),
				},
			},
			"mr_map_xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the extracted resource.", "xpath", "").AddRequiredWhen(DmAAAPMapResourceMRMapXPathCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRMapXPathIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPMapResourceMRMapXPathCondVal, DmAAAPMapResourceMRMapXPathIgnoreVal, false),
				},
			},
			"mr_tam_map": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Access Manager style-mapping resource. Access Manager organizes resources into a hierarchical protected object space. The default value is WebSEAL.", "tam-mapping", "").AddStringEnum("TFIM", "TAMBI", "WebSEAL", "Custom").AddDefaultValue("WebSEAL").AddRequiredWhen(DmAAAPMapResourceMRTAMMapCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRTAMMapIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TFIM", "TAMBI", "WebSEAL", "Custom"),
					validators.ConditionalRequiredString(DmAAAPMapResourceMRTAMMapCondVal, DmAAAPMapResourceMRTAMMapIgnoreVal, true),
				},
				Default: stringdefault.StaticString("WebSEAL"),
			},
			"mr_tam_instance_prefix": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the prefix to add based on the mapping style. Each protected object space has a defined convention. These different options help to map extracted resources to a resource string that follow a Tivoli naming convention.</p><ul><li>For the custom mapping style, specify the string to add as the prefix. In other words, use the <i>prefix</i> format.</li><li>For the TAMBI mapping style, specify the queue manager and queue, and separate the queue manager and the queue with a forward slash. In other words, use the <tt>/PDMQ/ <i>prefix</i></tt> format.</li><li>For the TFIM mapping style, specify the name of the Federated Identity Manager domain. In other words, use the <tt>/itfim-wssm/wssm-default/ <i>prefix</i></tt> format.</li><li>For the WebSEAL mapping style, specify the name of the WebSEAL instance. In other words, use the <tt>/WebSEAL/ <i>prefix</i></tt> format.</li></ul>", "tam-prefix", "").AddRequiredWhen(DmAAAPMapResourceMRTAMInstancePrefixCondVal.String()).AddNotValidWhen(DmAAAPMapResourceMRTAMInstancePrefixIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmAAAPMapResourceMRTAMInstancePrefixCondVal, DmAAAPMapResourceMRTAMInstancePrefixIgnoreVal, false),
				},
			},
			"mr_tam_web_seal_dyn_url_file": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the DynURL file. The file must be in the <tt>local:</tt> or <tt>store:</tt> directory. When configured and an entry matches a request, the DynURL output replaces the output of the extracted resource string that is added to the prefix.", "webseal-dynurl-file", "").AddNotValidWhen(DmAAAPMapResourceMRTAMWebSEALDynURLFileIgnoreVal.String()).String,
				Optional:            true,
			},
		},
	}
	DmAAAPMapResourceResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmAAAPMapResourceResourceSchema.Required = true
	} else {
		DmAAAPMapResourceResourceSchema.Optional = true
		DmAAAPMapResourceResourceSchema.Computed = true
	}
	return DmAAAPMapResourceResourceSchema
}

func (data DmAAAPMapResource) IsNull() bool {
	if !data.MrMethod.IsNull() {
		return false
	}
	if !data.MrCustomUrl.IsNull() {
		return false
	}
	if !data.MrMapUrl.IsNull() {
		return false
	}
	if !data.MrMapXpath.IsNull() {
		return false
	}
	if !data.MrTamMap.IsNull() {
		return false
	}
	if !data.MrTamInstancePrefix.IsNull() {
		return false
	}
	if !data.MrTamWebSealDynUrlFile.IsNull() {
		return false
	}
	return true
}

func (data DmAAAPMapResource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.MrMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMethod`, data.MrMethod.ValueString())
	}
	if !data.MrCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRCustomURL`, data.MrCustomUrl.ValueString())
	}
	if !data.MrMapUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMapURL`, data.MrMapUrl.ValueString())
	}
	if !data.MrMapXpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRMapXPath`, data.MrMapXpath.ValueString())
	}
	if !data.MrTamMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMMap`, data.MrTamMap.ValueString())
	}
	if !data.MrTamInstancePrefix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMInstancePrefix`, data.MrTamInstancePrefix.ValueString())
	}
	if !data.MrTamWebSealDynUrlFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MRTAMWebSEALDynURLFile`, data.MrTamWebSealDynUrlFile.ValueString())
	}
	return body
}

func (data *DmAAAPMapResource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MRMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMethod = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `MRCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrMapXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrTamMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrTamMap = types.StringValue("WebSEAL")
	}
	if value := res.Get(pathRoot + `MRTAMInstancePrefix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrTamInstancePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrTamInstancePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMWebSEALDynURLFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MrTamWebSealDynUrlFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrTamWebSealDynUrlFile = types.StringNull()
	}
}

func (data *DmAAAPMapResource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `MRMethod`); value.Exists() && !data.MrMethod.IsNull() {
		data.MrMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.MrMethod.ValueString() != "none" {
		data.MrMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRCustomURL`); value.Exists() && !data.MrCustomUrl.IsNull() {
		data.MrCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapURL`); value.Exists() && !data.MrMapUrl.IsNull() {
		data.MrMapUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRMapXPath`); value.Exists() && !data.MrMapXpath.IsNull() {
		data.MrMapXpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrMapXpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMMap`); value.Exists() && !data.MrTamMap.IsNull() {
		data.MrTamMap = tfutils.ParseStringFromGJSON(value)
	} else if data.MrTamMap.ValueString() != "WebSEAL" {
		data.MrTamMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMInstancePrefix`); value.Exists() && !data.MrTamInstancePrefix.IsNull() {
		data.MrTamInstancePrefix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrTamInstancePrefix = types.StringNull()
	}
	if value := res.Get(pathRoot + `MRTAMWebSEALDynURLFile`); value.Exists() && !data.MrTamWebSealDynUrlFile.IsNull() {
		data.MrTamWebSealDynUrlFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MrTamWebSealDynUrlFile = types.StringNull()
	}
}
