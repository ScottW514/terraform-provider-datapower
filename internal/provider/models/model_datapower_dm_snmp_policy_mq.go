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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSnmpPolicyMQ struct {
	Community types.String `tfsdk:"community"`
	Mode      types.String `tfsdk:"mode"`
	Host      types.String `tfsdk:"host"`
}

var DmSnmpPolicyMQObjectType = map[string]attr.Type{
	"community": types.StringType,
	"mode":      types.StringType,
	"host":      types.StringType,
}
var DmSnmpPolicyMQObjectDefault = map[string]attr.Value{
	"community": types.StringNull(),
	"mode":      types.StringNull(),
	"host":      types.StringValue("0.0.0.0/0"),
}
var DmSnmpPolicyMQDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"community": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a specific SNMP community. All SNMP Version 1 or Version 2c managers identifying themselves as a member of this community will have the Mode permissions set here for the Associated Domains selected here and must originate from the Remote Host Address specified here.", "", "").String,
			Computed:            true,
		},
		"mode": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the access privileges accorded to SNMP managers that belong to this community. Use none to disable SNMP access.", "", "").AddStringEnum("none", "read-only", "read-write").String,
			Computed:            true,
		},
		"host": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The IP address of an SNMP manager belonging to this community. The default value of '0.0.0.0/0' indicates all hosts, or all SNMP managers claiming membership in the community.", "", "").AddDefaultValue("0.0.0.0/0").String,
			Computed:            true,
		},
	},
}
var DmSnmpPolicyMQResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"community": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a specific SNMP community. All SNMP Version 1 or Version 2c managers identifying themselves as a member of this community will have the Mode permissions set here for the Associated Domains selected here and must originate from the Remote Host Address specified here.", "", "").String,
			Required:            true,
		},
		"mode": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select the access privileges accorded to SNMP managers that belong to this community. Use none to disable SNMP access.", "", "").AddStringEnum("none", "read-only", "read-write").String,
			Required:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("none", "read-only", "read-write"),
			},
		},
		"host": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("The IP address of an SNMP manager belonging to this community. The default value of '0.0.0.0/0' indicates all hosts, or all SNMP managers claiming membership in the community.", "", "").AddDefaultValue("0.0.0.0/0").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("0.0.0.0/0"),
		},
	},
}

func (data DmSnmpPolicyMQ) IsNull() bool {
	if !data.Community.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.Host.IsNull() {
		return false
	}
	return true
}

func (data DmSnmpPolicyMQ) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Community.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Community`, data.Community.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Mode`, data.Mode.ValueString())
	}
	if !data.Host.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Host`, data.Host.ValueString())
	}
	return body
}

func (data *DmSnmpPolicyMQ) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Community`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Community = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Community = types.StringNull()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Host = types.StringValue("0.0.0.0/0")
	}
}

func (data *DmSnmpPolicyMQ) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Community`); value.Exists() && !data.Community.IsNull() {
		data.Community = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Community = types.StringNull()
	}
	if value := res.Get(pathRoot + `Mode`); value.Exists() && !data.Mode.IsNull() {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `Host`); value.Exists() && !data.Host.IsNull() {
		data.Host = tfutils.ParseStringFromGJSON(value)
	} else if data.Host.ValueString() != "0.0.0.0/0" {
		data.Host = types.StringNull()
	}
}
