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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmB2BContact struct {
	FamilyName types.String `tfsdk:"family_name"`
	GivenName  types.String `tfsdk:"given_name"`
	Title      types.String `tfsdk:"title"`
	Phone      types.String `tfsdk:"phone"`
	Email      types.String `tfsdk:"email"`
}

var DmB2BContactObjectType = map[string]attr.Type{
	"family_name": types.StringType,
	"given_name":  types.StringType,
	"title":       types.StringType,
	"phone":       types.StringType,
	"email":       types.StringType,
}
var DmB2BContactObjectDefault = map[string]attr.Value{
	"family_name": types.StringNull(),
	"given_name":  types.StringNull(),
	"title":       types.StringNull(),
	"phone":       types.StringNull(),
	"email":       types.StringNull(),
}

func GetDmB2BContactDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmB2BContactDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"family_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the family name of the person to contact. A family name is the surname borne by family members.",
				Computed:            true,
			},
			"given_name": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the given name of the person to contact. A given name is the name used to identify an individual within a family.",
				Computed:            true,
			},
			"title": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the title of the person to contact.",
				Computed:            true,
			},
			"phone": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the telephone number to contact the person.",
				Computed:            true,
			},
			"email": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the email address to contact the person.",
				Computed:            true,
			},
		},
	}
	return DmB2BContactDataSourceSchema
}
func GetDmB2BContactResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmB2BContactResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"family_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the family name of the person to contact. A family name is the surname borne by family members.", "family-name", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9-. ]+$"), "Must match :"+"^[_a-zA-Z0-9-. ]+$"),
				},
			},
			"given_name": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the given name of the person to contact. A given name is the name used to identify an individual within a family.", "given-name", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9-. ]+$"), "Must match :"+"^[_a-zA-Z0-9-. ]+$"),
				},
			},
			"title": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the title of the person to contact.", "title", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9-. ]+$"), "Must match :"+"^[_a-zA-Z0-9-. ]+$"),
				},
			},
			"phone": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the telephone number to contact the person.", "phone", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[(0-9][)0-9- ]+[0-9]$"), "Must match :"+"^[(0-9][)0-9- ]+[0-9]$"),
				},
			},
			"email": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the email address to contact the person.", "email", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9-.]+@[_a-zA-Z0-9.]+$"), "Must match :"+"^[_a-zA-Z0-9-.]+@[_a-zA-Z0-9.]+$"),
				},
			},
		},
	}
	return DmB2BContactResourceSchema
}

func (data DmB2BContact) IsNull() bool {
	if !data.FamilyName.IsNull() {
		return false
	}
	if !data.GivenName.IsNull() {
		return false
	}
	if !data.Title.IsNull() {
		return false
	}
	if !data.Phone.IsNull() {
		return false
	}
	if !data.Email.IsNull() {
		return false
	}
	return true
}

func (data DmB2BContact) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.FamilyName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FamilyName`, data.FamilyName.ValueString())
	}
	if !data.GivenName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GivenName`, data.GivenName.ValueString())
	}
	if !data.Title.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Title`, data.Title.ValueString())
	}
	if !data.Phone.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Phone`, data.Phone.ValueString())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Email`, data.Email.ValueString())
	}
	return body
}

func (data *DmB2BContact) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FamilyName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FamilyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FamilyName = types.StringNull()
	}
	if value := res.Get(pathRoot + `GivenName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GivenName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GivenName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `Phone`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Phone = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Phone = types.StringNull()
	}
	if value := res.Get(pathRoot + `Email`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Email = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Email = types.StringNull()
	}
}

func (data *DmB2BContact) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `FamilyName`); value.Exists() && !data.FamilyName.IsNull() {
		data.FamilyName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FamilyName = types.StringNull()
	}
	if value := res.Get(pathRoot + `GivenName`); value.Exists() && !data.GivenName.IsNull() {
		data.GivenName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GivenName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Title`); value.Exists() && !data.Title.IsNull() {
		data.Title = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Title = types.StringNull()
	}
	if value := res.Get(pathRoot + `Phone`); value.Exists() && !data.Phone.IsNull() {
		data.Phone = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Phone = types.StringNull()
	}
	if value := res.Get(pathRoot + `Email`); value.Exists() && !data.Email.IsNull() {
		data.Email = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Email = types.StringNull()
	}
}
