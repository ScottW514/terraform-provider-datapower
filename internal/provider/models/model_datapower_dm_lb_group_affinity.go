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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmLBGroupAffinity struct {
	EnableSa            types.Bool             `tfsdk:"enable_sa"`
	InsertionCookieName types.String           `tfsdk:"insertion_cookie_name"`
	InsertionPath       types.String           `tfsdk:"insertion_path"`
	InsertionDomain     types.String           `tfsdk:"insertion_domain"`
	AffinityWlmOverride types.Bool             `tfsdk:"affinity_wlm_override"`
	AffinityMode        types.String           `tfsdk:"affinity_mode"`
	InsertionAttributes *DmInsertionAttributes `tfsdk:"insertion_attributes"`
}

var DmLBGroupAffinityAffinityModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "affinity_wlm_override",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var DmLBGroupAffinityObjectType = map[string]attr.Type{
	"enable_sa":             types.BoolType,
	"insertion_cookie_name": types.StringType,
	"insertion_path":        types.StringType,
	"insertion_domain":      types.StringType,
	"affinity_wlm_override": types.BoolType,
	"affinity_mode":         types.StringType,
	"insertion_attributes":  types.ObjectType{AttrTypes: DmInsertionAttributesObjectType},
}
var DmLBGroupAffinityObjectDefault = map[string]attr.Value{
	"enable_sa":             types.BoolValue(true),
	"insertion_cookie_name": types.StringValue("DPJSESSIONID"),
	"insertion_path":        types.StringValue("/"),
	"insertion_domain":      types.StringValue("datapower.com"),
	"affinity_wlm_override": types.BoolValue(false),
	"affinity_mode":         types.StringValue("activeConditional"),
	"insertion_attributes":  types.ObjectValueMust(DmInsertionAttributesObjectType, DmInsertionAttributesObjectDefault),
}

func GetDmLBGroupAffinityDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmLBGroupAffinityDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"enable_sa": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Enables or disables session affinity operations.",
				Computed:            true,
			},
			"insertion_cookie_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Name of the cookie inserted into the response when active or active-conditional session affinity is required.",
				Computed:            true,
			},
			"insertion_path": DataSourceSchema.StringAttribute{
				MarkdownDescription: "The path added to the insertion cookie in the response when active or active-conditional session affinity is required.",
				Computed:            true,
			},
			"insertion_domain": DataSourceSchema.StringAttribute{
				MarkdownDescription: "The domain added to the insertion cookie in the response when active or active-conditional session affinity is required. The domain name cannot begin with a dot.",
				Computed:            true,
			},
			"affinity_wlm_override": DataSourceSchema.BoolAttribute{
				MarkdownDescription: "Overrides the WebSphere Cell Session Affinity cluster configuration with the DataPower Gateway configuration information below.",
				Computed:            true,
			},
			"affinity_mode": DataSourceSchema.StringAttribute{
				MarkdownDescription: "The mode of session affinity applied to this load balancer group.",
				Computed:            true,
			},
			"insertion_attributes": GetDmInsertionAttributesDataSourceSchema("Specifies the attributes to insert in the cookie in the response when active or active-conditional session affinity is required.", "i-cookie-attributes", ""),
		},
	}
	DmLBGroupAffinityDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmLBGroupAffinityDataSourceSchema
}
func GetDmLBGroupAffinityResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmLBGroupAffinityResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmLBGroupAffinityObjectType,
				DmLBGroupAffinityObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"enable_sa": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables or disables session affinity operations.", "enable-affinity", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"insertion_cookie_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the cookie inserted into the response when active or active-conditional session affinity is required.", "i-cookie-name", "").AddDefaultValue("DPJSESSIONID").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("DPJSESSIONID"),
			},
			"insertion_path": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The path added to the insertion cookie in the response when active or active-conditional session affinity is required.", "i-path", "").AddDefaultValue("/").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"insertion_domain": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The domain added to the insertion cookie in the response when active or active-conditional session affinity is required. The domain name cannot begin with a dot.", "i-domain", "").AddDefaultValue("datapower.com").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("datapower.com"),
			},
			"affinity_wlm_override": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Overrides the WebSphere Cell Session Affinity cluster configuration with the DataPower Gateway configuration information below.", "override-wlm-affinity", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"affinity_mode": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The mode of session affinity applied to this load balancer group.", "affinity-mode", "").AddStringEnum("active", "activeConditional").AddDefaultValue("activeConditional").AddNotValidWhen(DmLBGroupAffinityAffinityModeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("active", "activeConditional"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmLBGroupAffinityAffinityModeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("activeConditional"),
			},
			"insertion_attributes": GetDmInsertionAttributesResourceSchema("Specifies the attributes to insert in the cookie in the response when active or active-conditional session affinity is required.", "i-cookie-attributes", "", false),
		},
	}
	DmLBGroupAffinityResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmLBGroupAffinityResourceSchema.Required = true
	} else {
		DmLBGroupAffinityResourceSchema.Optional = true
		DmLBGroupAffinityResourceSchema.Computed = true
	}
	return DmLBGroupAffinityResourceSchema
}

func (data DmLBGroupAffinity) IsNull() bool {
	if !data.EnableSa.IsNull() {
		return false
	}
	if !data.InsertionCookieName.IsNull() {
		return false
	}
	if !data.InsertionPath.IsNull() {
		return false
	}
	if !data.InsertionDomain.IsNull() {
		return false
	}
	if !data.AffinityWlmOverride.IsNull() {
		return false
	}
	if !data.AffinityMode.IsNull() {
		return false
	}
	if data.InsertionAttributes != nil {
		if !data.InsertionAttributes.IsNull() {
			return false
		}
	}
	return true
}

func (data DmLBGroupAffinity) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.EnableSa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableSA`, tfutils.StringFromBool(data.EnableSa, ""))
	}
	if !data.InsertionCookieName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InsertionCookieName`, data.InsertionCookieName.ValueString())
	}
	if !data.InsertionPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InsertionPath`, data.InsertionPath.ValueString())
	}
	if !data.InsertionDomain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InsertionDomain`, data.InsertionDomain.ValueString())
	}
	if !data.AffinityWlmOverride.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AffinityWLMOverride`, tfutils.StringFromBool(data.AffinityWlmOverride, ""))
	}
	if !data.AffinityMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AffinityMode`, data.AffinityMode.ValueString())
	}
	if data.InsertionAttributes != nil {
		if !data.InsertionAttributes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`InsertionAttributes`, data.InsertionAttributes.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *DmLBGroupAffinity) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EnableSA`); value.Exists() {
		data.EnableSa = tfutils.BoolFromString(value.String())
	} else {
		data.EnableSa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InsertionCookieName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InsertionCookieName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InsertionCookieName = types.StringValue("DPJSESSIONID")
	}
	if value := res.Get(pathRoot + `InsertionPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InsertionPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InsertionPath = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `InsertionDomain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InsertionDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InsertionDomain = types.StringValue("datapower.com")
	}
	if value := res.Get(pathRoot + `AffinityWLMOverride`); value.Exists() {
		data.AffinityWlmOverride = tfutils.BoolFromString(value.String())
	} else {
		data.AffinityWlmOverride = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AffinityMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AffinityMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AffinityMode = types.StringValue("activeConditional")
	}
	if value := res.Get(pathRoot + `InsertionAttributes`); value.Exists() {
		data.InsertionAttributes = &DmInsertionAttributes{}
		data.InsertionAttributes.FromBody(ctx, "", value)
	} else {
		data.InsertionAttributes = nil
	}
}

func (data *DmLBGroupAffinity) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EnableSA`); value.Exists() && !data.EnableSa.IsNull() {
		data.EnableSa = tfutils.BoolFromString(value.String())
	} else if !data.EnableSa.ValueBool() {
		data.EnableSa = types.BoolNull()
	}
	if value := res.Get(pathRoot + `InsertionCookieName`); value.Exists() && !data.InsertionCookieName.IsNull() {
		data.InsertionCookieName = tfutils.ParseStringFromGJSON(value)
	} else if data.InsertionCookieName.ValueString() != "DPJSESSIONID" {
		data.InsertionCookieName = types.StringNull()
	}
	if value := res.Get(pathRoot + `InsertionPath`); value.Exists() && !data.InsertionPath.IsNull() {
		data.InsertionPath = tfutils.ParseStringFromGJSON(value)
	} else if data.InsertionPath.ValueString() != "/" {
		data.InsertionPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `InsertionDomain`); value.Exists() && !data.InsertionDomain.IsNull() {
		data.InsertionDomain = tfutils.ParseStringFromGJSON(value)
	} else if data.InsertionDomain.ValueString() != "datapower.com" {
		data.InsertionDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `AffinityWLMOverride`); value.Exists() && !data.AffinityWlmOverride.IsNull() {
		data.AffinityWlmOverride = tfutils.BoolFromString(value.String())
	} else if data.AffinityWlmOverride.ValueBool() {
		data.AffinityWlmOverride = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AffinityMode`); value.Exists() && !data.AffinityMode.IsNull() {
		data.AffinityMode = tfutils.ParseStringFromGJSON(value)
	} else if data.AffinityMode.ValueString() != "activeConditional" {
		data.AffinityMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `InsertionAttributes`); value.Exists() {
		data.InsertionAttributes.UpdateFromBody(ctx, "", value)
	} else {
		data.InsertionAttributes = nil
	}
}
