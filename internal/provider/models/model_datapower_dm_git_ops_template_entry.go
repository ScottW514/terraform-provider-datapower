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

type DmGitOpsTemplateEntry struct {
	TemplateType  types.String `tfsdk:"template_type"`
	ClassName     types.String `tfsdk:"class_name"`
	Name          types.String `tfsdk:"name"`
	Field         types.String `tfsdk:"field"`
	Value         types.String `tfsdk:"value"`
	ValueInverse  types.String `tfsdk:"value_inverse"`
	ValueValidate types.String `tfsdk:"value_validate"`
}

var DmGitOpsTemplateEntryValueCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "template_type",
	AttrType:    "String",
	AttrDefault: "change",
	Value:       []string{"delete"},
}

var DmGitOpsTemplateEntryValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "template_type",
	AttrType:    "String",
	AttrDefault: "change",
	Value:       []string{"delete"},
}

var DmGitOpsTemplateEntryValueInverseCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "template_type",
	AttrType:    "String",
	AttrDefault: "change",
	Value:       []string{"custom"},
}

var DmGitOpsTemplateEntryValueInverseIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "template_type",
	AttrType:    "String",
	AttrDefault: "change",
	Value:       []string{"custom"},
}

var DmGitOpsTemplateEntryValueValidateIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "template_type",
	AttrType:    "String",
	AttrDefault: "change",
	Value:       []string{"custom"},
}

var DmGitOpsTemplateEntryObjectType = map[string]attr.Type{
	"template_type":  types.StringType,
	"class_name":     types.StringType,
	"name":           types.StringType,
	"field":          types.StringType,
	"value":          types.StringType,
	"value_inverse":  types.StringType,
	"value_validate": types.StringType,
}
var DmGitOpsTemplateEntryObjectDefault = map[string]attr.Value{
	"template_type":  types.StringValue("change"),
	"class_name":     types.StringNull(),
	"name":           types.StringNull(),
	"field":          types.StringNull(),
	"value":          types.StringNull(),
	"value_inverse":  types.StringNull(),
	"value_validate": types.StringNull(),
}

func GetDmGitOpsTemplateEntryDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmGitOpsTemplateEntryDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"template_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the type of template. The supported types are change and custom.",
				Computed:            true,
			},
			"class_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the the object class.",
				Computed:            true,
			},
			"name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"field": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the property name.",
				Computed:            true,
			},
			"value": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the value that is specific to the template type. <ul><li>When change, specify the replacement value for the identified property in this specific named object instance.</li><li>When custom, specify the transform to insert a template value</li></ul>",
				Computed:            true,
			},
			"value_inverse": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the value inverse. This value is the inverse transform to replace the template value with the wanted value.",
				Computed:            true,
			},
			"value_validate": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the query to validate the replaced value at the specified location. The query to return <tt>true</tt> when the template value is found at the specified location and <tt>false</tt> when the template value is not found at the specified location. The query is in the following format. <p><code>configuration.%Class[$match(`@name`, /^%Name/)]%ValueValidate</code></p>",
				Computed:            true,
			},
		},
	}
	return DmGitOpsTemplateEntryDataSourceSchema
}
func GetDmGitOpsTemplateEntryResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmGitOpsTemplateEntryResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"template_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of template. The supported types are change and custom.", "type", "").AddStringEnum("change", "add", "delete", "custom").AddDefaultValue("change").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("change", "add", "delete", "custom"),
				},
				Default: stringdefault.StaticString("change"),
			},
			"class_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the the object class.", "class", "").String,
				Required:            true,
			},
			"name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "name", "").String,
				Required:            true,
			},
			"field": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the property name.", "field", "").String,
				Required:            true,
			},
			"value": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value that is specific to the template type. <ul><li>When change, specify the replacement value for the identified property in this specific named object instance.</li><li>When custom, specify the transform to insert a template value</li></ul>", "value", "").AddRequiredWhen(DmGitOpsTemplateEntryValueCondVal.String()).AddNotValidWhen(DmGitOpsTemplateEntryValueIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmGitOpsTemplateEntryValueCondVal, DmGitOpsTemplateEntryValueIgnoreVal, false),
				},
			},
			"value_inverse": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the value inverse. This value is the inverse transform to replace the template value with the wanted value.", "value-inverse", "").AddRequiredWhen(DmGitOpsTemplateEntryValueInverseCondVal.String()).AddNotValidWhen(DmGitOpsTemplateEntryValueInverseIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmGitOpsTemplateEntryValueInverseCondVal, DmGitOpsTemplateEntryValueInverseIgnoreVal, false),
				},
			},
			"value_validate": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the query to validate the replaced value at the specified location. The query to return <tt>true</tt> when the template value is found at the specified location and <tt>false</tt> when the template value is not found at the specified location. The query is in the following format. <p><code>configuration.%Class[$match(`@name`, /^%Name/)]%ValueValidate</code></p>", "value-validate", "").AddNotValidWhen(DmGitOpsTemplateEntryValueValidateIgnoreVal.String()).String,
				Optional:            true,
			},
		},
	}
	return DmGitOpsTemplateEntryResourceSchema
}

func (data DmGitOpsTemplateEntry) IsNull() bool {
	if !data.TemplateType.IsNull() {
		return false
	}
	if !data.ClassName.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Field.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.ValueInverse.IsNull() {
		return false
	}
	if !data.ValueValidate.IsNull() {
		return false
	}
	return true
}

func (data DmGitOpsTemplateEntry) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.TemplateType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TemplateType`, data.TemplateType.ValueString())
	}
	if !data.ClassName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClassName`, data.ClassName.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Field.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Field`, data.Field.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.ValueInverse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueInverse`, data.ValueInverse.ValueString())
	}
	if !data.ValueValidate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValueValidate`, data.ValueValidate.ValueString())
	}
	return body
}

func (data *DmGitOpsTemplateEntry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TemplateType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TemplateType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TemplateType = types.StringValue("change")
	}
	if value := res.Get(pathRoot + `ClassName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClassName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClassName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Field`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Field = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Field = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueInverse`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValueInverse = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueInverse = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueValidate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValueValidate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueValidate = types.StringNull()
	}
}

func (data *DmGitOpsTemplateEntry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TemplateType`); value.Exists() && !data.TemplateType.IsNull() {
		data.TemplateType = tfutils.ParseStringFromGJSON(value)
	} else if data.TemplateType.ValueString() != "change" {
		data.TemplateType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClassName`); value.Exists() && !data.ClassName.IsNull() {
		data.ClassName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClassName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Field`); value.Exists() && !data.Field.IsNull() {
		data.Field = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Field = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueInverse`); value.Exists() && !data.ValueInverse.IsNull() {
		data.ValueInverse = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueInverse = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValueValidate`); value.Exists() && !data.ValueValidate.IsNull() {
		data.ValueValidate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValueValidate = types.StringNull()
	}
}
