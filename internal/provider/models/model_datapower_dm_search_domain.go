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

type DmSearchDomain struct {
	SearchDomain types.String `tfsdk:"search_domain"`
}

var DmSearchDomainObjectType = map[string]attr.Type{
	"search_domain": types.StringType,
}
var DmSearchDomainObjectDefault = map[string]attr.Value{
	"search_domain": types.StringNull(),
}
var DmSearchDomainDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"search_domain": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the domain name.", "", "").String,
			Computed:            true,
		},
	},
}
var DmSearchDomainResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"search_domain": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the domain name.", "", "").String,
			Required:            true,
		},
	},
}

func (data DmSearchDomain) IsNull() bool {
	if !data.SearchDomain.IsNull() {
		return false
	}
	return true
}

func (data DmSearchDomain) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.SearchDomain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SearchDomain`, data.SearchDomain.ValueString())
	}
	return body
}

func (data *DmSearchDomain) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SearchDomain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SearchDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SearchDomain = types.StringNull()
	}
}

func (data *DmSearchDomain) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `SearchDomain`); value.Exists() && !data.SearchDomain.IsNull() {
		data.SearchDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SearchDomain = types.StringNull()
	}
}
