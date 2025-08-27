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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSAMLAttribute struct {
	SourceType   types.String `tfsdk:"source_type"`
	Name         types.String `tfsdk:"name"`
	Format       types.String `tfsdk:"format"`
	Xpath        types.String `tfsdk:"xpath"`
	ValueData    types.String `tfsdk:"value_data"`
	SubValueData types.String `tfsdk:"sub_value_data"`
	FriendlyName types.String `tfsdk:"friendly_name"`
}

var DmSAMLAttributeNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "source_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"var"},
}
var DmSAMLAttributeXPathCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "source_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xpath"},
}
var DmSAMLAttributeXPathIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmSAMLAttributeValueDataIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "source_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"xpath"},
}
var DmSAMLAttributeSubValueDataIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "source_type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"var"},
}

var DmSAMLAttributeObjectType = map[string]attr.Type{
	"source_type":    types.StringType,
	"name":           types.StringType,
	"format":         types.StringType,
	"xpath":          types.StringType,
	"value_data":     types.StringType,
	"sub_value_data": types.StringType,
	"friendly_name":  types.StringType,
}
var DmSAMLAttributeObjectDefault = map[string]attr.Value{
	"source_type":    types.StringNull(),
	"name":           types.StringNull(),
	"format":         types.StringNull(),
	"xpath":          types.StringNull(),
	"value_data":     types.StringNull(),
	"sub_value_data": types.StringNull(),
	"friendly_name":  types.StringNull(),
}

func GetDmSAMLAttributeDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSAMLAttributeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"source_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the data source to retrieve the value for each SAML attribute. <p>If the Data Source Type is a variable, configure a DataPower service or context variable name with Data for Attribute Value. That variable must contain a result element with a list of attribute sub elements. Each attribute element must contain a name attribute, which is used to match the Supplementary Data setting. A sample to describe the format of the variable content:</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p>", "", "").AddStringEnum("var", "xpath", "static").String,
				Computed:            true,
			},
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the SAML attribute to be generated. <p>When the Data Source Type is variable, this setting can be an empty string. In that case, the attribute name that is carried by the variable's content is used.</p>", "", "").String,
				Computed:            true,
			},
			"format": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the optional Namespace URI for the SAML 1.x attribute,</p><p>Optional: Specify the NameFormat value for the SAML 2.0 attribute.</p>", "", "").String,
				Computed:            true,
			},
			"xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPATH information to get the value for the SAML attribute when the attribute value is From Input Message, and specify the XPath expression to locate the value. <p>The XML nodes that the XPath expression points to is the value for the SAML attribute.</p>", "", "").String,
				Computed:            true,
			},
			"value_data": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the information to get the value for the SAML attribute. <p>If the attribute contains a Static Value per each AAA Policy, specify the static string value.</p><p>If the Data Source Type is variable, specify the variable name. You can input an empty string as the variable name to use the default variable var://context/ldap/auxiliary-attributes. That variable is maintained by the LDAP authentication or authorization to query auxiliary LDAP attributes.</p><p>In any case, the variable that is being used here must contain a result element with a list of attribute sub elements. Each attribute element must contain a name attribute, which is used to match the Supplementary Data setting. The following sample describes the format of the variable content.</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p>", "", "").String,
				Computed:            true,
			},
			"sub_value_data": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>When the SAML attribute value is retrieved from a variable as defined in Data for Attribute Value setting, specify the value to match the name attribute of the attribute-value elements that are carried by that variable.</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p><p>When this value is empty, the value of each attribute-value element is treated as one SAML AttributeValue. Therefore, if multiple attribute-values are carried by the DataPower variable, there can be multiple SAML AttributeValues in one SAML Attribute element.</p>", "", "").String,
				Computed:            true,
			},
			"friendly_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a meaningful name for the SAML attribute.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmSAMLAttributeDataSourceSchema
}
func GetDmSAMLAttributeResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSAMLAttributeResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"source_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the data source to retrieve the value for each SAML attribute. <p>If the Data Source Type is a variable, configure a DataPower service or context variable name with Data for Attribute Value. That variable must contain a result element with a list of attribute sub elements. Each attribute element must contain a name attribute, which is used to match the Supplementary Data setting. A sample to describe the format of the variable content:</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p>", "", "").AddStringEnum("var", "xpath", "static").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("var", "xpath", "static"),
				},
			},
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the SAML attribute to be generated. <p>When the Data Source Type is variable, this setting can be an empty string. In that case, the attribute name that is carried by the variable's content is used.</p>", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSAMLAttributeNameCondVal, validators.Evaluation{}, false),
				},
			},
			"format": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the optional Namespace URI for the SAML 1.x attribute,</p><p>Optional: Specify the NameFormat value for the SAML 2.0 attribute.</p>", "", "").String,
				Optional:            true,
			},
			"xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPATH information to get the value for the SAML attribute when the attribute value is From Input Message, and specify the XPath expression to locate the value. <p>The XML nodes that the XPath expression points to is the value for the SAML attribute.</p>", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSAMLAttributeXPathCondVal, validators.Evaluation{}, false),
				},
			},
			"value_data": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the information to get the value for the SAML attribute. <p>If the attribute contains a Static Value per each AAA Policy, specify the static string value.</p><p>If the Data Source Type is variable, specify the variable name. You can input an empty string as the variable name to use the default variable var://context/ldap/auxiliary-attributes. That variable is maintained by the LDAP authentication or authorization to query auxiliary LDAP attributes.</p><p>In any case, the variable that is being used here must contain a result element with a list of attribute sub elements. Each attribute element must contain a name attribute, which is used to match the Supplementary Data setting. The following sample describes the format of the variable content.</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p>", "", "").String,
				Optional:            true,
			},
			"sub_value_data": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>When the SAML attribute value is retrieved from a variable as defined in Data for Attribute Value setting, specify the value to match the name attribute of the attribute-value elements that are carried by that variable.</p><p>&lt;result> &lt;attribute-value name=\"cn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"sn\">Alice&lt;/attribute-value> &lt;attribute-value name=\"creatorsName\"> cn=Manager,dc=datapower,dc=com&lt;/attribute-value> &lt;/result></p><p>When this value is empty, the value of each attribute-value element is treated as one SAML AttributeValue. Therefore, if multiple attribute-values are carried by the DataPower variable, there can be multiple SAML AttributeValues in one SAML Attribute element.</p>", "", "").String,
				Optional:            true,
			},
			"friendly_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a meaningful name for the SAML attribute.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmSAMLAttributeResourceSchema
}

func (data DmSAMLAttribute) IsNull() bool {
	if !data.SourceType.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Format.IsNull() {
		return false
	}
	if !data.Xpath.IsNull() {
		return false
	}
	if !data.ValueData.IsNull() {
		return false
	}
	if !data.SubValueData.IsNull() {
		return false
	}
	if !data.FriendlyName.IsNull() {
		return false
	}
	return true
}

func (data DmSAMLAttribute) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.SourceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SourceType`, data.SourceType.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Format`, data.Format.ValueString())
	}
	if !data.Xpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.Xpath.ValueString())
	}
	if !data.ValueData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueData`, data.ValueData.ValueString())
	}
	if !data.SubValueData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SubValueData`, data.SubValueData.ValueString())
	}
	if !data.FriendlyName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FriendlyName`, data.FriendlyName.ValueString())
	}
	return body
}

func (data *DmSAMLAttribute) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SourceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubValueData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SubValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `FriendlyName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FriendlyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FriendlyName = types.StringNull()
	}
}

func (data *DmSAMLAttribute) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SourceType`); value.Exists() && !data.SourceType.IsNull() {
		data.SourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && !data.Format.IsNull() {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.Xpath.IsNull() {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueData`); value.Exists() && !data.ValueData.IsNull() {
		data.ValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `SubValueData`); value.Exists() && !data.SubValueData.IsNull() {
		data.SubValueData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SubValueData = types.StringNull()
	}
	if value := res.Get(pathRoot + `FriendlyName`); value.Exists() && !data.FriendlyName.IsNull() {
		data.FriendlyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FriendlyName = types.StringNull()
	}
}
