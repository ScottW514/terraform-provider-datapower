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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type APIPlan struct {
	Id                           types.String `tfsdk:"id"`
	AppDomain                    types.String `tfsdk:"app_domain"`
	UserSummary                  types.String `tfsdk:"user_summary"`
	Name                         types.String `tfsdk:"name"`
	ProductId                    types.String `tfsdk:"product_id"`
	ProductName                  types.String `tfsdk:"product_name"`
	ProductVersion               types.String `tfsdk:"product_version"`
	ProductTitle                 types.String `tfsdk:"product_title"`
	UseRateLimitGroup            types.Bool   `tfsdk:"use_rate_limit_group"`
	RateLimit                    types.List   `tfsdk:"rate_limit"`
	BurstLimit                   types.List   `tfsdk:"burst_limit"`
	RateLimitGroup               types.String `tfsdk:"rate_limit_group"`
	UseLimitDefinitions          types.Bool   `tfsdk:"use_limit_definitions"`
	AssemblyBurstLimit           types.List   `tfsdk:"assembly_burst_limit"`
	AssemblyBurstLimitDefinition types.List   `tfsdk:"assembly_burst_limit_definition"`
	AssemblyRateLimit            types.List   `tfsdk:"assembly_rate_limit"`
	AssemblyRateLimitDefinition  types.List   `tfsdk:"assembly_rate_limit_definition"`
	AssemblyCountLimit           types.List   `tfsdk:"assembly_count_limit"`
	AssemblyCountLimitDefinition types.List   `tfsdk:"assembly_count_limit_definition"`
	SpaceId                      types.String `tfsdk:"space_id"`
	SpaceName                    types.String `tfsdk:"space_name"`
	Api                          types.List   `tfsdk:"api"`
	ExcludeOperation             types.List   `tfsdk:"exclude_operation"`
	Override                     types.List   `tfsdk:"override"`
	RateLimitScope               types.String `tfsdk:"rate_limit_scope"`
	GraphQlSchemaOptions         types.List   `tfsdk:"graph_ql_schema_options"`
}

var APIPlanObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"name":                            types.StringType,
	"product_id":                      types.StringType,
	"product_name":                    types.StringType,
	"product_version":                 types.StringType,
	"product_title":                   types.StringType,
	"use_rate_limit_group":            types.BoolType,
	"rate_limit":                      types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"burst_limit":                     types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}},
	"rate_limit_group":                types.StringType,
	"use_limit_definitions":           types.BoolType,
	"assembly_burst_limit":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}},
	"assembly_burst_limit_definition": types.ListType{ElemType: types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}},
	"assembly_rate_limit":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"assembly_rate_limit_definition":  types.ListType{ElemType: types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}},
	"assembly_count_limit":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPICountLimitObjectType}},
	"assembly_count_limit_definition": types.ListType{ElemType: types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}},
	"space_id":                        types.StringType,
	"space_name":                      types.StringType,
	"api":                             types.ListType{ElemType: types.StringType},
	"exclude_operation":               types.ListType{ElemType: types.StringType},
	"override":                        types.ListType{ElemType: types.StringType},
	"rate_limit_scope":                types.StringType,
	"graph_ql_schema_options":         types.ListType{ElemType: types.StringType},
}

func (data APIPlan) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIPlan"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIPlan) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.ProductId.IsNull() {
		return false
	}
	if !data.ProductName.IsNull() {
		return false
	}
	if !data.ProductVersion.IsNull() {
		return false
	}
	if !data.ProductTitle.IsNull() {
		return false
	}
	if !data.UseRateLimitGroup.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.BurstLimit.IsNull() {
		return false
	}
	if !data.RateLimitGroup.IsNull() {
		return false
	}
	if !data.UseLimitDefinitions.IsNull() {
		return false
	}
	if !data.AssemblyBurstLimit.IsNull() {
		return false
	}
	if !data.AssemblyBurstLimitDefinition.IsNull() {
		return false
	}
	if !data.AssemblyRateLimit.IsNull() {
		return false
	}
	if !data.AssemblyRateLimitDefinition.IsNull() {
		return false
	}
	if !data.AssemblyCountLimit.IsNull() {
		return false
	}
	if !data.AssemblyCountLimitDefinition.IsNull() {
		return false
	}
	if !data.SpaceId.IsNull() {
		return false
	}
	if !data.SpaceName.IsNull() {
		return false
	}
	if !data.Api.IsNull() {
		return false
	}
	if !data.ExcludeOperation.IsNull() {
		return false
	}
	if !data.Override.IsNull() {
		return false
	}
	if !data.RateLimitScope.IsNull() {
		return false
	}
	if !data.GraphQlSchemaOptions.IsNull() {
		return false
	}
	return true
}

func (data APIPlan) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.ProductId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductID`, data.ProductId.ValueString())
	}
	if !data.ProductName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductName`, data.ProductName.ValueString())
	}
	if !data.ProductVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductVersion`, data.ProductVersion.ValueString())
	}
	if !data.ProductTitle.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProductTitle`, data.ProductTitle.ValueString())
	}
	if !data.UseRateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseRateLimitGroup`, tfutils.StringFromBool(data.UseRateLimitGroup, ""))
	}
	if !data.RateLimit.IsNull() {
		var values []DmAPIRateLimit
		data.RateLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`RateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.BurstLimit.IsNull() {
		var values []DmAPIBurstLimit
		data.BurstLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`BurstLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RateLimitGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimitGroup`, data.RateLimitGroup.ValueString())
	}
	if !data.UseLimitDefinitions.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseLimitDefinitions`, tfutils.StringFromBool(data.UseLimitDefinitions, ""))
	}
	if !data.AssemblyBurstLimit.IsNull() {
		var values []DmAPIBurstLimit
		data.AssemblyBurstLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyBurstLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyBurstLimitDefinition.IsNull() {
		var values []DmDefinitionLink
		data.AssemblyBurstLimitDefinition.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyBurstLimitDefinition`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyRateLimit.IsNull() {
		var values []DmAPIRateLimit
		data.AssemblyRateLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyRateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyRateLimitDefinition.IsNull() {
		var values []DmDefinitionLink
		data.AssemblyRateLimitDefinition.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyRateLimitDefinition`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyCountLimit.IsNull() {
		var values []DmAPICountLimit
		data.AssemblyCountLimit.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyCountLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyCountLimitDefinition.IsNull() {
		var values []DmDefinitionLink
		data.AssemblyCountLimitDefinition.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyCountLimitDefinition`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SpaceId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SpaceID`, data.SpaceId.ValueString())
	}
	if !data.SpaceName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SpaceName`, data.SpaceName.ValueString())
	}
	if !data.Api.IsNull() {
		var values []string
		data.Api.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`API`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ExcludeOperation.IsNull() {
		var values []string
		data.ExcludeOperation.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`ExcludeOperation`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Override.IsNull() {
		var values []string
		data.Override.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Override`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RateLimitScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimitScope`, data.RateLimitScope.ValueString())
	}
	if !data.GraphQlSchemaOptions.IsNull() {
		var values []string
		data.GraphQlSchemaOptions.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`GraphQLSchemaOptions`+".-1", map[string]string{"value": val})
		}
	}
	return body
}

func (data *APIPlan) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductVersion = types.StringValue("1.0.0")
	}
	if value := res.Get(pathRoot + `ProductTitle`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProductTitle = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductTitle = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		l := []DmAPIRateLimit{}
		if value := res.Get(`RateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() {
		l := []DmAPIBurstLimit{}
		if value := res.Get(`BurstLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIBurstLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.BurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.BurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.BurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseLimitDefinitions`); value.Exists() {
		data.UseLimitDefinitions = tfutils.BoolFromString(value.String())
	} else {
		data.UseLimitDefinitions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimit`); value.Exists() {
		l := []DmAPIBurstLimit{}
		if value := res.Get(`AssemblyBurstLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIBurstLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimitDefinition`); value.Exists() {
		l := []DmDefinitionLink{}
		if value := res.Get(`AssemblyBurstLimitDefinition`); value.Exists() {
			for _, v := range value.Array() {
				item := DmDefinitionLink{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyBurstLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyBurstLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimit`); value.Exists() {
		l := []DmAPIRateLimit{}
		if value := res.Get(`AssemblyRateLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIRateLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimitDefinition`); value.Exists() {
		l := []DmDefinitionLink{}
		if value := res.Get(`AssemblyRateLimitDefinition`); value.Exists() {
			for _, v := range value.Array() {
				item := DmDefinitionLink{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyRateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyRateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimit`); value.Exists() {
		l := []DmAPICountLimit{}
		if value := res.Get(`AssemblyCountLimit`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPICountLimit{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPICountLimitObjectType}, l)
		} else {
			data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
		}
	} else {
		data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimitDefinition`); value.Exists() {
		l := []DmDefinitionLink{}
		if value := res.Get(`AssemblyCountLimitDefinition`); value.Exists() {
			for _, v := range value.Array() {
				item := DmDefinitionLink{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyCountLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyCountLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `SpaceID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SpaceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SpaceId = types.StringNull()
	}
	if value := res.Get(pathRoot + `SpaceName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SpaceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SpaceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `API`); value.Exists() {
		data.Api = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Api = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ExcludeOperation`); value.Exists() {
		data.ExcludeOperation = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ExcludeOperation = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Override`); value.Exists() {
		data.Override = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Override = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RateLimitScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimitScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitScope = types.StringValue("per-application")
	}
	if value := res.Get(pathRoot + `GraphQLSchemaOptions`); value.Exists() {
		data.GraphQlSchemaOptions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.GraphQlSchemaOptions = types.ListNull(types.StringType)
	}
}

func (data *APIPlan) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductID`); value.Exists() && !data.ProductId.IsNull() {
		data.ProductId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductName`); value.Exists() && !data.ProductName.IsNull() {
		data.ProductName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductVersion`); value.Exists() && !data.ProductVersion.IsNull() {
		data.ProductVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.ProductVersion.ValueString() != "1.0.0" {
		data.ProductVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProductTitle`); value.Exists() && !data.ProductTitle.IsNull() {
		data.ProductTitle = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProductTitle = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseRateLimitGroup`); value.Exists() && !data.UseRateLimitGroup.IsNull() {
		data.UseRateLimitGroup = tfutils.BoolFromString(value.String())
	} else if data.UseRateLimitGroup.ValueBool() {
		data.UseRateLimitGroup = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		l := []DmAPIRateLimit{}
		for _, v := range value.Array() {
			item := DmAPIRateLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.RateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.RateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `BurstLimit`); value.Exists() && !data.BurstLimit.IsNull() {
		l := []DmAPIBurstLimit{}
		for _, v := range value.Array() {
			item := DmAPIBurstLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.BurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.BurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.BurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `RateLimitGroup`); value.Exists() && !data.RateLimitGroup.IsNull() {
		data.RateLimitGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimitGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseLimitDefinitions`); value.Exists() && !data.UseLimitDefinitions.IsNull() {
		data.UseLimitDefinitions = tfutils.BoolFromString(value.String())
	} else if data.UseLimitDefinitions.ValueBool() {
		data.UseLimitDefinitions = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimit`); value.Exists() && !data.AssemblyBurstLimit.IsNull() {
		l := []DmAPIBurstLimit{}
		for _, v := range value.Array() {
			item := DmAPIBurstLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}, l)
		} else {
			data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
		}
	} else {
		data.AssemblyBurstLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyBurstLimitDefinition`); value.Exists() && !data.AssemblyBurstLimitDefinition.IsNull() {
		l := []DmDefinitionLink{}
		for _, v := range value.Array() {
			item := DmDefinitionLink{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyBurstLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyBurstLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyBurstLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimit`); value.Exists() && !data.AssemblyRateLimit.IsNull() {
		l := []DmAPIRateLimit{}
		for _, v := range value.Array() {
			item := DmAPIRateLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}, l)
		} else {
			data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
		}
	} else {
		data.AssemblyRateLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPIRateLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyRateLimitDefinition`); value.Exists() && !data.AssemblyRateLimitDefinition.IsNull() {
		l := []DmDefinitionLink{}
		for _, v := range value.Array() {
			item := DmDefinitionLink{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyRateLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyRateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyRateLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimit`); value.Exists() && !data.AssemblyCountLimit.IsNull() {
		l := []DmAPICountLimit{}
		for _, v := range value.Array() {
			item := DmAPICountLimit{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimit, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPICountLimitObjectType}, l)
		} else {
			data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
		}
	} else {
		data.AssemblyCountLimit = types.ListNull(types.ObjectType{AttrTypes: DmAPICountLimitObjectType})
	}
	if value := res.Get(pathRoot + `AssemblyCountLimitDefinition`); value.Exists() && !data.AssemblyCountLimitDefinition.IsNull() {
		l := []DmDefinitionLink{}
		for _, v := range value.Array() {
			item := DmDefinitionLink{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AssemblyCountLimitDefinition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDefinitionLinkObjectType}, l)
		} else {
			data.AssemblyCountLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
		}
	} else {
		data.AssemblyCountLimitDefinition = types.ListNull(types.ObjectType{AttrTypes: DmDefinitionLinkObjectType})
	}
	if value := res.Get(pathRoot + `SpaceID`); value.Exists() && !data.SpaceId.IsNull() {
		data.SpaceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SpaceId = types.StringNull()
	}
	if value := res.Get(pathRoot + `SpaceName`); value.Exists() && !data.SpaceName.IsNull() {
		data.SpaceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SpaceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `API`); value.Exists() && !data.Api.IsNull() {
		data.Api = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Api = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ExcludeOperation`); value.Exists() && !data.ExcludeOperation.IsNull() {
		data.ExcludeOperation = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ExcludeOperation = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Override`); value.Exists() && !data.Override.IsNull() {
		data.Override = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Override = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RateLimitScope`); value.Exists() && !data.RateLimitScope.IsNull() {
		data.RateLimitScope = tfutils.ParseStringFromGJSON(value)
	} else if data.RateLimitScope.ValueString() != "per-application" {
		data.RateLimitScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSchemaOptions`); value.Exists() && !data.GraphQlSchemaOptions.IsNull() {
		data.GraphQlSchemaOptions = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.GraphQlSchemaOptions = types.ListNull(types.StringType)
	}
}
