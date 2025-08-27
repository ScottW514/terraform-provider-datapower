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
	"regexp"

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

type DmConfigModifyType struct {
	Match    types.String `tfsdk:"match"`
	Type     types.String `tfsdk:"type"`
	Property types.String `tfsdk:"property"`
	Value    types.String `tfsdk:"value"`
}

var DmConfigModifyTypePropertyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"add"},
}
var DmConfigModifyTypeValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"add", "change"},
}
var DmConfigModifyTypePropertyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var DmConfigModifyTypeValueIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmConfigModifyTypeObjectType = map[string]attr.Type{
	"match":    types.StringType,
	"type":     types.StringType,
	"property": types.StringType,
	"value":    types.StringType,
}
var DmConfigModifyTypeObjectDefault = map[string]attr.Value{
	"match":    types.StringNull(),
	"type":     types.StringNull(),
	"property": types.StringNull(),
	"value":    types.StringNull(),
}

func GetDmConfigModifyTypeDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmConfigModifyTypeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"match": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matching configuration is modified. To create a match statement, type a correctly formatted resource match in the horizontal text box select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>", "match", "").String,
				Computed:            true,
			},
			"type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Selects the type of configuration modification.", "type", "").AddStringEnum("add", "delete", "change").String,
				Computed:            true,
			},
			"property": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of the property to add.", "name", "").String,
				Computed:            true,
			},
			"value": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of the added or changed property. To change a property value, you can use an explicit value or a value that contains a variable in the <tt>${ <i>variable</i> }</tt> format. If you use variables, you need a deployment policy variable configuration to map the variable to its replacement value. For example, when the value is <tt>${newName}</tt> or <tt>${newName}_Service</tt> , a referenced deployment policy variable configuration must map the <tt>newName</tt> variable to an explicit replacement value.", "value", "").String,
				Computed:            true,
			},
		},
	}
	return DmConfigModifyTypeDataSourceSchema
}
func GetDmConfigModifyTypeResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmConfigModifyTypeResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Matching configuration is modified. To create a match statement, type a correctly formatted resource match in the horizontal text box select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>", "match", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[-_a-zA-Z0-9:.*]+/[-_a-zA-Z0-9.*]+/[-a-z0-9/*]+(\\?[^=]+=[^&]+(&[^=]+=[^&]+)*)?$"), "Must match :"+"^[-_a-zA-Z0-9:.*]+/[-_a-zA-Z0-9.*]+/[-a-z0-9/*]+(\\?[^=]+=[^&]+(&[^=]+=[^&]+)*)?$"),
				},
			},
			"type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Selects the type of configuration modification.", "type", "").AddStringEnum("add", "delete", "change").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("add", "delete", "change"),
				},
			},
			"property": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of the property to add.", "name", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmConfigModifyTypePropertyCondVal, validators.Evaluation{}, false),
				},
			},
			"value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the value of the added or changed property. To change a property value, you can use an explicit value or a value that contains a variable in the <tt>${ <i>variable</i> }</tt> format. If you use variables, you need a deployment policy variable configuration to map the variable to its replacement value. For example, when the value is <tt>${newName}</tt> or <tt>${newName}_Service</tt> , a referenced deployment policy variable configuration must map the <tt>newName</tt> variable to an explicit replacement value.", "value", "").String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmConfigModifyTypeValueCondVal, validators.Evaluation{}, false),
				},
			},
		},
	}
	return DmConfigModifyTypeResourceSchema
}

func (data DmConfigModifyType) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Property.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	return true
}

func (data DmConfigModifyType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Property.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Property`, data.Property.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	return body
}

func (data *DmConfigModifyType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Property = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Property = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}

func (data *DmConfigModifyType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() && !data.Property.IsNull() {
		data.Property = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Property = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
}
