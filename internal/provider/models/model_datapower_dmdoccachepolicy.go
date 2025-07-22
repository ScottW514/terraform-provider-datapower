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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmDocCachePolicy struct {
	Match                 types.String `tfsdk:"match"`
	Type                  types.String `tfsdk:"type"`
	Ttl                   types.Int64  `tfsdk:"ttl"`
	Priority              types.Int64  `tfsdk:"priority"`
	Xc10Grid              types.String `tfsdk:"xc10_grid"`
	CacheBackendResponses types.Bool   `tfsdk:"cache_backend_responses"`
	HttpCacheValidation   types.Bool   `tfsdk:"http_cache_validation"`
	ReturnExpired         types.Bool   `tfsdk:"return_expired"`
	RestInvalidation      types.Bool   `tfsdk:"rest_invalidation"`
	CacheUnsafeResponse   types.Bool   `tfsdk:"cache_unsafe_response"`
}

var DmDocCachePolicyObjectType = map[string]attr.Type{
	"match":                   types.StringType,
	"type":                    types.StringType,
	"ttl":                     types.Int64Type,
	"priority":                types.Int64Type,
	"xc10_grid":               types.StringType,
	"cache_backend_responses": types.BoolType,
	"http_cache_validation":   types.BoolType,
	"return_expired":          types.BoolType,
	"rest_invalidation":       types.BoolType,
	"cache_unsafe_response":   types.BoolType,
}
var DmDocCachePolicyObjectDefault = map[string]attr.Value{
	"match":                   types.StringNull(),
	"type":                    types.StringValue("protocol"),
	"ttl":                     types.Int64Value(900),
	"priority":                types.Int64Value(128),
	"xc10_grid":               types.StringNull(),
	"cache_backend_responses": types.BoolValue(false),
	"http_cache_validation":   types.BoolValue(false),
	"return_expired":          types.BoolValue(false),
	"rest_invalidation":       types.BoolValue(false),
	"cache_unsafe_response":   types.BoolValue(false),
}
var DmDocCachePolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"match": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Match Expression", "match", "").String,
			Computed:            true,
		},
		"type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy Type", "type", "").AddStringEnum("protocol", "no-cache", "fixed").AddDefaultValue("protocol").String,
			Computed:            true,
		},
		"ttl": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TTL", "ttl", "").AddIntegerRange(5, 31708800).AddDefaultValue("900").String,
			Computed:            true,
		},
		"priority": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Priority", "priority", "").AddIntegerRange(1, 255).AddDefaultValue("128").String,
			Computed:            true,
		},
		"xc10_grid": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Grid", "xc10-grid", "").String,
			Computed:            true,
		},
		"cache_backend_responses": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Back-end Responses", "cache-backend-response", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"http_cache_validation": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP Cache Validation", "http-cache-validation", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"return_expired": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Expired Document", "return-expired", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"rest_invalidation": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("RESTful Invalidation", "rest-invalidation", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"cache_unsafe_response": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Response to POST and PUT Requests", "cache-unsafe-response", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmDocCachePolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"match": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("URL Match Expression", "match", "").String,
			Optional:            true,
		},
		"type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Policy Type", "type", "").AddStringEnum("protocol", "no-cache", "fixed").AddDefaultValue("protocol").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("protocol", "no-cache", "fixed"),
			},
			Default: stringdefault.StaticString("protocol"),
		},
		"ttl": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("TTL", "ttl", "").AddIntegerRange(5, 31708800).AddDefaultValue("900").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(5, 31708800),
			},
			Default: int64default.StaticInt64(900),
		},
		"priority": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Priority", "priority", "").AddIntegerRange(1, 255).AddDefaultValue("128").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 255),
			},
			Default: int64default.StaticInt64(128),
		},
		"xc10_grid": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Grid", "xc10-grid", "").String,
			Optional:            true,
		},
		"cache_backend_responses": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Back-end Responses", "cache-backend-response", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"http_cache_validation": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("HTTP Cache Validation", "http-cache-validation", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"return_expired": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Return Expired Document", "return-expired", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"rest_invalidation": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("RESTful Invalidation", "rest-invalidation", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"cache_unsafe_response": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Cache Response to POST and PUT Requests", "cache-unsafe-response", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmDocCachePolicy) IsNull() bool {
	if !data.Match.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Ttl.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.Xc10Grid.IsNull() {
		return false
	}
	if !data.CacheBackendResponses.IsNull() {
		return false
	}
	if !data.HttpCacheValidation.IsNull() {
		return false
	}
	if !data.ReturnExpired.IsNull() {
		return false
	}
	if !data.RestInvalidation.IsNull() {
		return false
	}
	if !data.CacheUnsafeResponse.IsNull() {
		return false
	}
	return true
}

func (data DmDocCachePolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Match.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Match`, data.Match.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Ttl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TTL`, data.Ttl.ValueInt64())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueInt64())
	}
	if !data.Xc10Grid.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XC10Grid`, data.Xc10Grid.ValueString())
	}
	if !data.CacheBackendResponses.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheBackendResponses`, tfutils.StringFromBool(data.CacheBackendResponses, false))
	}
	if !data.HttpCacheValidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPCacheValidation`, tfutils.StringFromBool(data.HttpCacheValidation, false))
	}
	if !data.ReturnExpired.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReturnExpired`, tfutils.StringFromBool(data.ReturnExpired, false))
	}
	if !data.RestInvalidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RESTInvalidation`, tfutils.StringFromBool(data.RestInvalidation, false))
	}
	if !data.CacheUnsafeResponse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheUnsafeResponse`, tfutils.StringFromBool(data.CacheUnsafeResponse, false))
	}
	return body
}

func (data *DmDocCachePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("protocol")
	}
	if value := res.Get(pathRoot + `TTL`); value.Exists() {
		data.Ttl = types.Int64Value(value.Int())
	} else {
		data.Ttl = types.Int64Value(900)
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `XC10Grid`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Xc10Grid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xc10Grid = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheBackendResponses`); value.Exists() {
		data.CacheBackendResponses = tfutils.BoolFromString(value.String())
	} else {
		data.CacheBackendResponses = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPCacheValidation`); value.Exists() {
		data.HttpCacheValidation = tfutils.BoolFromString(value.String())
	} else {
		data.HttpCacheValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReturnExpired`); value.Exists() {
		data.ReturnExpired = tfutils.BoolFromString(value.String())
	} else {
		data.ReturnExpired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RESTInvalidation`); value.Exists() {
		data.RestInvalidation = tfutils.BoolFromString(value.String())
	} else {
		data.RestInvalidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheUnsafeResponse`); value.Exists() {
		data.CacheUnsafeResponse = tfutils.BoolFromString(value.String())
	} else {
		data.CacheUnsafeResponse = types.BoolNull()
	}
}

func (data *DmDocCachePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Match`); value.Exists() && !data.Match.IsNull() {
		data.Match = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Match = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "protocol" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `TTL`); value.Exists() && !data.Ttl.IsNull() {
		data.Ttl = types.Int64Value(value.Int())
	} else if data.Ttl.ValueInt64() != 900 {
		data.Ttl = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else if data.Priority.ValueInt64() != 128 {
		data.Priority = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XC10Grid`); value.Exists() && !data.Xc10Grid.IsNull() {
		data.Xc10Grid = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xc10Grid = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheBackendResponses`); value.Exists() && !data.CacheBackendResponses.IsNull() {
		data.CacheBackendResponses = tfutils.BoolFromString(value.String())
	} else if data.CacheBackendResponses.ValueBool() {
		data.CacheBackendResponses = types.BoolNull()
	}
	if value := res.Get(pathRoot + `HTTPCacheValidation`); value.Exists() && !data.HttpCacheValidation.IsNull() {
		data.HttpCacheValidation = tfutils.BoolFromString(value.String())
	} else if data.HttpCacheValidation.ValueBool() {
		data.HttpCacheValidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ReturnExpired`); value.Exists() && !data.ReturnExpired.IsNull() {
		data.ReturnExpired = tfutils.BoolFromString(value.String())
	} else if data.ReturnExpired.ValueBool() {
		data.ReturnExpired = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RESTInvalidation`); value.Exists() && !data.RestInvalidation.IsNull() {
		data.RestInvalidation = tfutils.BoolFromString(value.String())
	} else if data.RestInvalidation.ValueBool() {
		data.RestInvalidation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CacheUnsafeResponse`); value.Exists() && !data.CacheUnsafeResponse.IsNull() {
		data.CacheUnsafeResponse = tfutils.BoolFromString(value.String())
	} else if data.CacheUnsafeResponse.ValueBool() {
		data.CacheUnsafeResponse = types.BoolNull()
	}
}
