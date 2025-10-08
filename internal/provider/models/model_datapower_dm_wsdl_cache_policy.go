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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSDLCachePolicy struct {
	Match types.String `tfsdk:"match"`
	Ttl   types.Int64  `tfsdk:"ttl"`
}

var DmWSDLCachePolicyObjectType = map[string]attr.Type{
	"match": types.StringType,
	"ttl":   types.Int64Type,
}
var DmWSDLCachePolicyObjectDefault = map[string]attr.Value{
	"match": types.StringNull(),
	"ttl":   types.Int64Value(900),
}

func GetDmWSDLCachePolicyDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSDLCachePolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"match": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Provide a literal or wildcard expression to define a URL set included in this cache policy.",
				Computed:            true,
			},
			"ttl": DataSourceSchema.Int64Attribute{
				MarkdownDescription: "Configures lifetime in seconds of document. Enter an integer between 5 and 86400. The default value is 900.",
				Computed:            true,
			},
		},
	}
	return DmWSDLCachePolicyDataSourceSchema
}
func GetDmWSDLCachePolicyResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSDLCachePolicyResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"match": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Provide a literal or wildcard expression to define a URL set included in this cache policy.", "", "").String,
				Optional:            true,
			},
			"ttl": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configures lifetime in seconds of document. Enter an integer between 5 and 86400. The default value is 900.", "", "").AddIntegerRange(5, 86400).AddDefaultValue("900").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 86400),
				},
				Default: int64default.StaticInt64(900),
			},
		},
	}
	return DmWSDLCachePolicyResourceSchema
}

func (data DmWSDLCachePolicy) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Ttl.IsNull() {
		return false
	}
	return true
}

func (data DmWSDLCachePolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Ttl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TTL`, data.Ttl.ValueInt64())
	}
	return body
}

func (data *DmWSDLCachePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `TTL`); value.Exists() {
		data.Ttl = types.Int64Value(value.Int())
	} else {
		data.Ttl = types.Int64Value(900)
	}
}

func (data *DmWSDLCachePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `TTL`); value.Exists() && !data.Ttl.IsNull() {
		data.Ttl = types.Int64Value(value.Int())
	} else if data.Ttl.ValueInt64() != 900 {
		data.Ttl = types.Int64Null()
	}
}
