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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSBaseWSDL struct {
	WsdlSourceLocation types.String `tfsdk:"wsdl_source_location"`
	WsdlName           types.String `tfsdk:"wsdl_name"`
	PolicyAttachments  types.String `tfsdk:"policy_attachments"`
}

var DmWSBaseWSDLObjectType = map[string]attr.Type{
	"wsdl_source_location": types.StringType,
	"wsdl_name":            types.StringType,
	"policy_attachments":   types.StringType,
}
var DmWSBaseWSDLObjectDefault = map[string]attr.Value{
	"wsdl_source_location": types.StringNull(),
	"wsdl_name":            types.StringNull(),
	"policy_attachments":   types.StringNull(),
}
var DmWSBaseWSDLDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"wsdl_source_location": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Source Location", "", "").String,
			Computed:            true,
		},
		"wsdl_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Local Name", "", "").String,
			Computed:            true,
		},
		"policy_attachments": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy Attachments", "", "policyattachments").String,
			Computed:            true,
		},
	},
}
var DmWSBaseWSDLResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"wsdl_source_location": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("WSDL Source Location", "", "").String,
			Required:            true,
		},
		"wsdl_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Local Name", "", "").String,
			Required:            true,
		},
		"policy_attachments": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy Attachments", "", "policyattachments").String,
			Optional:            true,
		},
	},
}

func (data DmWSBaseWSDL) IsNull() bool {
	if !data.WsdlSourceLocation.IsNull() {
		return false
	}
	if !data.WsdlName.IsNull() {
		return false
	}
	if !data.PolicyAttachments.IsNull() {
		return false
	}
	return true
}

func (data DmWSBaseWSDL) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.WsdlSourceLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLSourceLocation`, data.WsdlSourceLocation.ValueString())
	}
	if !data.WsdlName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLName`, data.WsdlName.ValueString())
	}
	if !data.PolicyAttachments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyAttachments`, data.PolicyAttachments.ValueString())
	}
	return body
}

func (data *DmWSBaseWSDL) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLSourceLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlSourceLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlSourceLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachments`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyAttachments = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachments = types.StringNull()
	}
}

func (data *DmWSBaseWSDL) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `WSDLSourceLocation`); value.Exists() && !data.WsdlSourceLocation.IsNull() {
		data.WsdlSourceLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlSourceLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLName`); value.Exists() && !data.WsdlName.IsNull() {
		data.WsdlName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyAttachments`); value.Exists() && !data.PolicyAttachments.IsNull() {
		data.PolicyAttachments = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyAttachments = types.StringNull()
	}
}
