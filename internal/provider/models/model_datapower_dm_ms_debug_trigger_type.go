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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmMSDebugTriggerType struct {
	ClientIp    types.String `tfsdk:"client_ip"`
	InUrlMatch  types.String `tfsdk:"in_url_match"`
	OutUrlMatch types.String `tfsdk:"out_url_match"`
	RuleType    types.String `tfsdk:"rule_type"`
	RuleMatch   types.String `tfsdk:"rule_match"`
	Xpath       types.String `tfsdk:"xpath"`
}

var DmMSDebugTriggerTypeObjectType = map[string]attr.Type{
	"client_ip":     types.StringType,
	"in_url_match":  types.StringType,
	"out_url_match": types.StringType,
	"rule_type":     types.StringType,
	"rule_match":    types.StringType,
	"xpath":         types.StringType,
}
var DmMSDebugTriggerTypeObjectDefault = map[string]attr.Value{
	"client_ip":     types.StringNull(),
	"in_url_match":  types.StringNull(),
	"out_url_match": types.StringNull(),
	"rule_type":     types.StringNull(),
	"rule_match":    types.StringNull(),
	"xpath":         types.StringNull(),
}

func GetDmMSDebugTriggerTypeDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmMSDebugTriggerTypeDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"client_ip": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a PCRE to match against client IP addresses. Requests from clients with matching IP addresses will trigger the probe. To create a match for all IP addresses, specify .* instead of * as the PCRE.",
				Computed:            true,
			},
			"in_url_match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a PCRE to match against the inbound URLs. Requests from clients with matching URLs will trigger the probe. To create a match for all URLs, specify .* instead of * as the PCRE.",
				Computed:            true,
			},
			"out_url_match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a PCRE to match against the outbound URLs. Responses from servers with matching URLs will trigger the probe. To create a match for all URLs, specify .* instead of * as the PCRE.",
				Computed:            true,
			},
			"rule_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Select the rule direction or type that will trigger the probe.",
				Computed:            true,
			},
			"rule_match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify a PCRE to match against names of processing rules. Rules with matching names will trigger the probe. To create a match for all names, specify .* instead of * as the PCRE.",
				Computed:            true,
			},
			"xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify an XPath expression of use the XPath Tool to define an XPath expression to match against messages. Messages that contain the expression will trigger the probe.",
				Computed:            true,
			},
		},
	}
	return DmMSDebugTriggerTypeDataSourceSchema
}
func GetDmMSDebugTriggerTypeResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmMSDebugTriggerTypeResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"client_ip": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE to match against client IP addresses. Requests from clients with matching IP addresses will trigger the probe. To create a match for all IP addresses, specify .* instead of * as the PCRE.", "", "").String,
				Optional:            true,
			},
			"in_url_match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE to match against the inbound URLs. Requests from clients with matching URLs will trigger the probe. To create a match for all URLs, specify .* instead of * as the PCRE.", "", "").String,
				Optional:            true,
			},
			"out_url_match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE to match against the outbound URLs. Responses from servers with matching URLs will trigger the probe. To create a match for all URLs, specify .* instead of * as the PCRE.", "", "").String,
				Optional:            true,
			},
			"rule_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the rule direction or type that will trigger the probe.", "", "").AddStringEnum("all", "response", "request", "call", "error", "scheduled", "lbhealth").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("all", "response", "request", "call", "error", "scheduled", "lbhealth"),
				},
			},
			"rule_match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE to match against names of processing rules. Rules with matching names will trigger the probe. To create a match for all names, specify .* instead of * as the PCRE.", "", "").String,
				Optional:            true,
			},
			"xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify an XPath expression of use the XPath Tool to define an XPath expression to match against messages. Messages that contain the expression will trigger the probe.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmMSDebugTriggerTypeResourceSchema
}

func (data DmMSDebugTriggerType) IsNull() bool {
	if !data.ClientIp.IsNull() {
		return false
	}
	if !data.InUrlMatch.IsNull() {
		return false
	}
	if !data.OutUrlMatch.IsNull() {
		return false
	}
	if !data.RuleType.IsNull() {
		return false
	}
	if !data.RuleMatch.IsNull() {
		return false
	}
	if !data.Xpath.IsNull() {
		return false
	}
	return true
}

func (data DmMSDebugTriggerType) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ClientIp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientIP`, data.ClientIp.ValueString())
	}
	if !data.InUrlMatch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InURLMatch`, data.InUrlMatch.ValueString())
	}
	if !data.OutUrlMatch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutURLMatch`, data.OutUrlMatch.ValueString())
	}
	if !data.RuleType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RuleType`, data.RuleType.ValueString())
	}
	if !data.RuleMatch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RuleMatch`, data.RuleMatch.ValueString())
	}
	if !data.Xpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.Xpath.ValueString())
	}
	return body
}

func (data *DmMSDebugTriggerType) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ClientIP`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientIp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InURLMatch`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InUrlMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InUrlMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutURLMatch`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutUrlMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutUrlMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `RuleType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RuleType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RuleType = types.StringNull()
	}
	if value := res.Get(pathRoot + `RuleMatch`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RuleMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RuleMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
}

func (data *DmMSDebugTriggerType) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ClientIP`); value.Exists() && !data.ClientIp.IsNull() {
		data.ClientIp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientIp = types.StringNull()
	}
	if value := res.Get(pathRoot + `InURLMatch`); value.Exists() && !data.InUrlMatch.IsNull() {
		data.InUrlMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InUrlMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutURLMatch`); value.Exists() && !data.OutUrlMatch.IsNull() {
		data.OutUrlMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutUrlMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `RuleType`); value.Exists() && !data.RuleType.IsNull() {
		data.RuleType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RuleType = types.StringNull()
	}
	if value := res.Get(pathRoot + `RuleMatch`); value.Exists() && !data.RuleMatch.IsNull() {
		data.RuleMatch = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RuleMatch = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.Xpath.IsNull() {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
}
