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

type DmSchemaValidation struct {
	Matching         types.String `tfsdk:"matching"`
	ValidationMode   types.String `tfsdk:"validation_mode"`
	SchemaUrl        types.String `tfsdk:"schema_url"`
	UrlRewritePolicy types.String `tfsdk:"url_rewrite_policy"`
	DynamicSchema    types.String `tfsdk:"dynamic_schema"`
}

var DmSchemaValidationSchemaURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"schema", "schema-rewrite"},
}
var DmSchemaValidationURLRewritePolicyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"attribute-rewrite", "schema-rewrite"},
}
var DmSchemaValidationDynamicSchemaCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"dynamic-schema"},
}
var DmSchemaValidationSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"attribute-rewrite", "dynamic-schema", "default"},
}
var DmSchemaValidationURLRewritePolicyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"schema", "dynamic-schema", "default"},
}
var DmSchemaValidationDynamicSchemaIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "validation_mode",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"schema-rewrite", "schema", "attribute-rewrite", "default"},
}

var DmSchemaValidationObjectType = map[string]attr.Type{
	"matching":           types.StringType,
	"validation_mode":    types.StringType,
	"schema_url":         types.StringType,
	"url_rewrite_policy": types.StringType,
	"dynamic_schema":     types.StringType,
}
var DmSchemaValidationObjectDefault = map[string]attr.Value{
	"matching":           types.StringNull(),
	"validation_mode":    types.StringNull(),
	"schema_url":         types.StringNull(),
	"url_rewrite_policy": types.StringNull(),
	"dynamic_schema":     types.StringNull(),
}

func GetDmSchemaValidationDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmSchemaValidationDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"matching": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the matching rule that is applied to candidate XML documents. Documents whose URL conforms to an expression contained in the matching rule are subject to XML schema validation.", "", "matching").String,
				Computed:            true,
			},
			"validation_mode": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the procedure to use to validate a schema.", "", "").AddStringEnum("default", "schema", "schema-rewrite", "attribute-rewrite", "dynamic-schema").String,
				Computed:            true,
			},
			"schema_url": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the schema to use for validation. Used only when the Validation Mode is schema or schema-rewrite.", "", "").AddRequiredWhen(DmSchemaValidationSchemaURLCondVal.String()).AddNotValidWhen(DmSchemaValidationSchemaURLIgnoreVal.String()).String,
				Computed:            true,
			},
			"url_rewrite_policy": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a URL Rewrite Policy from the list. Used only the Validation Mode is attribute-rewrite or schema-rewrite.", "", "url_rewrite_policy").AddRequiredWhen(DmSchemaValidationURLRewritePolicyCondVal.String()).AddNotValidWhen(DmSchemaValidationURLRewritePolicyIgnoreVal.String()).String,
				Computed:            true,
			},
			"dynamic_schema": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a dynamic schema used to perform document validation regardless of any validation processing instructions contained within the document. Dynamic schemas used Schema Exception Maps or Document Crypto Maps.", "", "").AddRequiredWhen(DmSchemaValidationDynamicSchemaCondVal.String()).AddNotValidWhen(DmSchemaValidationDynamicSchemaIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	return DmSchemaValidationDataSourceSchema
}
func GetDmSchemaValidationResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmSchemaValidationResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"matching": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the matching rule that is applied to candidate XML documents. Documents whose URL conforms to an expression contained in the matching rule are subject to XML schema validation.", "", "matching").String,
				Required:            true,
			},
			"validation_mode": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the procedure to use to validate a schema.", "", "").AddStringEnum("default", "schema", "schema-rewrite", "attribute-rewrite", "dynamic-schema").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "schema", "schema-rewrite", "attribute-rewrite", "dynamic-schema"),
				},
			},
			"schema_url": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the schema to use for validation. Used only when the Validation Mode is schema or schema-rewrite.", "", "").AddRequiredWhen(DmSchemaValidationSchemaURLCondVal.String()).AddNotValidWhen(DmSchemaValidationSchemaURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSchemaValidationSchemaURLCondVal, DmSchemaValidationSchemaURLIgnoreVal, false),
				},
			},
			"url_rewrite_policy": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a URL Rewrite Policy from the list. Used only the Validation Mode is attribute-rewrite or schema-rewrite.", "", "url_rewrite_policy").AddRequiredWhen(DmSchemaValidationURLRewritePolicyCondVal.String()).AddNotValidWhen(DmSchemaValidationURLRewritePolicyIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSchemaValidationURLRewritePolicyCondVal, DmSchemaValidationURLRewritePolicyIgnoreVal, false),
				},
			},
			"dynamic_schema": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a dynamic schema used to perform document validation regardless of any validation processing instructions contained within the document. Dynamic schemas used Schema Exception Maps or Document Crypto Maps.", "", "").AddRequiredWhen(DmSchemaValidationDynamicSchemaCondVal.String()).AddNotValidWhen(DmSchemaValidationDynamicSchemaIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmSchemaValidationDynamicSchemaCondVal, DmSchemaValidationDynamicSchemaIgnoreVal, false),
				},
			},
		},
	}
	return DmSchemaValidationResourceSchema
}

func (data DmSchemaValidation) IsNull() bool {
	if !data.Matching.IsNull() {
		return false
	}
	if !data.ValidationMode.IsNull() {
		return false
	}
	if !data.SchemaUrl.IsNull() {
		return false
	}
	if !data.UrlRewritePolicy.IsNull() {
		return false
	}
	if !data.DynamicSchema.IsNull() {
		return false
	}
	return true
}

func (data DmSchemaValidation) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Matching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Matching`, data.Matching.ValueString())
	}
	if !data.ValidationMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationMode`, data.ValidationMode.ValueString())
	}
	if !data.SchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchemaURL`, data.SchemaUrl.ValueString())
	}
	if !data.UrlRewritePolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRewritePolicy`, data.UrlRewritePolicy.ValueString())
	}
	if !data.DynamicSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicSchema`, data.DynamicSchema.ValueString())
	}
	return body
}

func (data *DmSchemaValidation) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Matching`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Matching = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Matching = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynamicSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicSchema = types.StringNull()
	}
}

func (data *DmSchemaValidation) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Matching`); value.Exists() && !data.Matching.IsNull() {
		data.Matching = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Matching = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationMode`); value.Exists() && !data.ValidationMode.IsNull() {
		data.ValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaURL`); value.Exists() && !data.SchemaUrl.IsNull() {
		data.SchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `URLRewritePolicy`); value.Exists() && !data.UrlRewritePolicy.IsNull() {
		data.UrlRewritePolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRewritePolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicSchema`); value.Exists() && !data.DynamicSchema.IsNull() {
		data.DynamicSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicSchema = types.StringNull()
	}
}
