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

type XMLManager struct {
	Id                                   types.String `tfsdk:"id"`
	AppDomain                            types.String `tfsdk:"app_domain"`
	UserSummary                          types.String `tfsdk:"user_summary"`
	UrlRefreshPolicy                     types.String `tfsdk:"url_refresh_policy"`
	CompileOptionsPolicy                 types.String `tfsdk:"compile_options_policy"`
	CacheMemorySize                      types.Int64  `tfsdk:"cache_memory_size"`
	CacheSize                            types.Int64  `tfsdk:"cache_size"`
	Sha1Caching                          types.Bool   `tfsdk:"sha1_caching"`
	StaticDocumentCalls                  types.Bool   `tfsdk:"static_document_calls"`
	SearchResults                        types.Bool   `tfsdk:"search_results"`
	VirtualServers                       types.List   `tfsdk:"virtual_servers"`
	ParserLimitsBytesScanned             types.Int64  `tfsdk:"parser_limits_bytes_scanned"`
	ParserLimitsElementDepth             types.Int64  `tfsdk:"parser_limits_element_depth"`
	ParserLimitsAttributeCount           types.Int64  `tfsdk:"parser_limits_attribute_count"`
	ParserLimitsMaxNodeSize              types.Int64  `tfsdk:"parser_limits_max_node_size"`
	ParserLimitsForbidExternalReferences types.Bool   `tfsdk:"parser_limits_forbid_external_references"`
	ParserLimitsExternalReferences       types.String `tfsdk:"parser_limits_external_references"`
	ParserLimitsMaxPrefixes              types.Int64  `tfsdk:"parser_limits_max_prefixes"`
	ParserLimitsMaxNamespaces            types.Int64  `tfsdk:"parser_limits_max_namespaces"`
	ParserLimitsMaxLocalNames            types.Int64  `tfsdk:"parser_limits_max_local_names"`
	DocCacheMaxDocs                      types.Int64  `tfsdk:"doc_cache_max_docs"`
	DocCacheSize                         types.Int64  `tfsdk:"doc_cache_size"`
	DocMaxWrites                         types.Int64  `tfsdk:"doc_max_writes"`
	DocCachePolicy                       types.List   `tfsdk:"doc_cache_policy"`
	SchemaValidation                     types.List   `tfsdk:"schema_validation"`
	ScheduledRule                        types.List   `tfsdk:"scheduled_rule"`
	UserAgent                            types.String `tfsdk:"user_agent"`
	JsonParserSettings                   types.String `tfsdk:"json_parser_settings"`
	LdapConnPool                         types.String `tfsdk:"ldap_conn_pool"`
}

var XMLManagerObjectType = map[string]attr.Type{
	"id":                                       types.StringType,
	"app_domain":                               types.StringType,
	"user_summary":                             types.StringType,
	"url_refresh_policy":                       types.StringType,
	"compile_options_policy":                   types.StringType,
	"cache_memory_size":                        types.Int64Type,
	"cache_size":                               types.Int64Type,
	"sha1_caching":                             types.BoolType,
	"static_document_calls":                    types.BoolType,
	"search_results":                           types.BoolType,
	"virtual_servers":                          types.ListType{ElemType: types.StringType},
	"parser_limits_bytes_scanned":              types.Int64Type,
	"parser_limits_element_depth":              types.Int64Type,
	"parser_limits_attribute_count":            types.Int64Type,
	"parser_limits_max_node_size":              types.Int64Type,
	"parser_limits_forbid_external_references": types.BoolType,
	"parser_limits_external_references":        types.StringType,
	"parser_limits_max_prefixes":               types.Int64Type,
	"parser_limits_max_namespaces":             types.Int64Type,
	"parser_limits_max_local_names":            types.Int64Type,
	"doc_cache_max_docs":                       types.Int64Type,
	"doc_cache_size":                           types.Int64Type,
	"doc_max_writes":                           types.Int64Type,
	"doc_cache_policy":                         types.ListType{ElemType: types.ObjectType{AttrTypes: DmDocCachePolicyObjectType}},
	"schema_validation":                        types.ListType{ElemType: types.ObjectType{AttrTypes: DmSchemaValidationObjectType}},
	"scheduled_rule":                           types.ListType{ElemType: types.ObjectType{AttrTypes: DmScheduledRuleObjectType}},
	"user_agent":                               types.StringType,
	"json_parser_settings":                     types.StringType,
	"ldap_conn_pool":                           types.StringType,
}

func (data XMLManager) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XMLManager"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XMLManager) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.UrlRefreshPolicy.IsNull() {
		return false
	}
	if !data.CompileOptionsPolicy.IsNull() {
		return false
	}
	if !data.CacheMemorySize.IsNull() {
		return false
	}
	if !data.CacheSize.IsNull() {
		return false
	}
	if !data.Sha1Caching.IsNull() {
		return false
	}
	if !data.StaticDocumentCalls.IsNull() {
		return false
	}
	if !data.SearchResults.IsNull() {
		return false
	}
	if !data.VirtualServers.IsNull() {
		return false
	}
	if !data.ParserLimitsBytesScanned.IsNull() {
		return false
	}
	if !data.ParserLimitsElementDepth.IsNull() {
		return false
	}
	if !data.ParserLimitsAttributeCount.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxNodeSize.IsNull() {
		return false
	}
	if !data.ParserLimitsForbidExternalReferences.IsNull() {
		return false
	}
	if !data.ParserLimitsExternalReferences.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxPrefixes.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxNamespaces.IsNull() {
		return false
	}
	if !data.ParserLimitsMaxLocalNames.IsNull() {
		return false
	}
	if !data.DocCacheMaxDocs.IsNull() {
		return false
	}
	if !data.DocCacheSize.IsNull() {
		return false
	}
	if !data.DocMaxWrites.IsNull() {
		return false
	}
	if !data.DocCachePolicy.IsNull() {
		return false
	}
	if !data.SchemaValidation.IsNull() {
		return false
	}
	if !data.ScheduledRule.IsNull() {
		return false
	}
	if !data.UserAgent.IsNull() {
		return false
	}
	if !data.JsonParserSettings.IsNull() {
		return false
	}
	if !data.LdapConnPool.IsNull() {
		return false
	}
	return true
}

func (data XMLManager) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.UrlRefreshPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRefreshPolicy`, data.UrlRefreshPolicy.ValueString())
	}
	if !data.CompileOptionsPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CompileOptionsPolicy`, data.CompileOptionsPolicy.ValueString())
	}
	if !data.CacheMemorySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheMemorySize`, data.CacheMemorySize.ValueInt64())
	}
	if !data.CacheSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheSize`, data.CacheSize.ValueInt64())
	}
	if !data.Sha1Caching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SHA1Caching`, tfutils.StringFromBool(data.Sha1Caching, false))
	}
	if !data.StaticDocumentCalls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StaticDocumentCalls`, tfutils.StringFromBool(data.StaticDocumentCalls, false))
	}
	if !data.SearchResults.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SearchResults`, tfutils.StringFromBool(data.SearchResults, false))
	}
	if !data.VirtualServers.IsNull() {
		var values []string
		data.VirtualServers.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`VirtualServers`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ParserLimitsBytesScanned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsBytesScanned`, data.ParserLimitsBytesScanned.ValueInt64())
	}
	if !data.ParserLimitsElementDepth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsElementDepth`, data.ParserLimitsElementDepth.ValueInt64())
	}
	if !data.ParserLimitsAttributeCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsAttributeCount`, data.ParserLimitsAttributeCount.ValueInt64())
	}
	if !data.ParserLimitsMaxNodeSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxNodeSize`, data.ParserLimitsMaxNodeSize.ValueInt64())
	}
	if !data.ParserLimitsForbidExternalReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsForbidExternalReferences`, tfutils.StringFromBool(data.ParserLimitsForbidExternalReferences, false))
	}
	if !data.ParserLimitsExternalReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsExternalReferences`, data.ParserLimitsExternalReferences.ValueString())
	}
	if !data.ParserLimitsMaxPrefixes.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxPrefixes`, data.ParserLimitsMaxPrefixes.ValueInt64())
	}
	if !data.ParserLimitsMaxNamespaces.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxNamespaces`, data.ParserLimitsMaxNamespaces.ValueInt64())
	}
	if !data.ParserLimitsMaxLocalNames.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParserLimitsMaxLocalNames`, data.ParserLimitsMaxLocalNames.ValueInt64())
	}
	if !data.DocCacheMaxDocs.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocCacheMaxDocs`, data.DocCacheMaxDocs.ValueInt64())
	}
	if !data.DocCacheSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocCacheSize`, data.DocCacheSize.ValueInt64())
	}
	if !data.DocMaxWrites.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocMaxWrites`, data.DocMaxWrites.ValueInt64())
	}
	if !data.DocCachePolicy.IsNull() {
		var values []DmDocCachePolicy
		data.DocCachePolicy.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`DocCachePolicy`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SchemaValidation.IsNull() {
		var values []DmSchemaValidation
		data.SchemaValidation.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`SchemaValidation`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ScheduledRule.IsNull() {
		var values []DmScheduledRule
		data.ScheduledRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ScheduledRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UserAgent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserAgent`, data.UserAgent.ValueString())
	}
	if !data.JsonParserSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONParserSettings`, data.JsonParserSettings.ValueString())
	}
	if !data.LdapConnPool.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPConnPool`, data.LdapConnPool.ValueString())
	}
	return body
}

func (data *XMLManager) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRefreshPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CompileOptionsPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CompileOptionsPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CompileOptionsPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheMemorySize`); value.Exists() {
		data.CacheMemorySize = types.Int64Value(value.Int())
	} else {
		data.CacheMemorySize = types.Int64Value(2147483647)
	}
	if value := res.Get(pathRoot + `CacheSize`); value.Exists() {
		data.CacheSize = types.Int64Value(value.Int())
	} else {
		data.CacheSize = types.Int64Value(256)
	}
	if value := res.Get(pathRoot + `SHA1Caching`); value.Exists() {
		data.Sha1Caching = tfutils.BoolFromString(value.String())
	} else {
		data.Sha1Caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StaticDocumentCalls`); value.Exists() {
		data.StaticDocumentCalls = tfutils.BoolFromString(value.String())
	} else {
		data.StaticDocumentCalls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SearchResults`); value.Exists() {
		data.SearchResults = tfutils.BoolFromString(value.String())
	} else {
		data.SearchResults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VirtualServers`); value.Exists() {
		data.VirtualServers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.VirtualServers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ParserLimitsBytesScanned`); value.Exists() {
		data.ParserLimitsBytesScanned = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsBytesScanned = types.Int64Value(4194304)
	}
	if value := res.Get(pathRoot + `ParserLimitsElementDepth`); value.Exists() {
		data.ParserLimitsElementDepth = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsElementDepth = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `ParserLimitsAttributeCount`); value.Exists() {
		data.ParserLimitsAttributeCount = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsAttributeCount = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNodeSize`); value.Exists() {
		data.ParserLimitsMaxNodeSize = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxNodeSize = types.Int64Value(33554432)
	}
	if value := res.Get(pathRoot + `ParserLimitsForbidExternalReferences`); value.Exists() {
		data.ParserLimitsForbidExternalReferences = tfutils.BoolFromString(value.String())
	} else {
		data.ParserLimitsForbidExternalReferences = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParserLimitsExternalReferences = types.StringValue("forbid")
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxPrefixes`); value.Exists() {
		data.ParserLimitsMaxPrefixes = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxPrefixes = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNamespaces`); value.Exists() {
		data.ParserLimitsMaxNamespaces = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxNamespaces = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxLocalNames`); value.Exists() {
		data.ParserLimitsMaxLocalNames = types.Int64Value(value.Int())
	} else {
		data.ParserLimitsMaxLocalNames = types.Int64Value(60000)
	}
	if value := res.Get(pathRoot + `DocCacheMaxDocs`); value.Exists() {
		data.DocCacheMaxDocs = types.Int64Value(value.Int())
	} else {
		data.DocCacheMaxDocs = types.Int64Value(5000)
	}
	if value := res.Get(pathRoot + `DocCacheSize`); value.Exists() {
		data.DocCacheSize = types.Int64Value(value.Int())
	} else {
		data.DocCacheSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DocMaxWrites`); value.Exists() {
		data.DocMaxWrites = types.Int64Value(value.Int())
	} else {
		data.DocMaxWrites = types.Int64Value(32768)
	}
	if value := res.Get(pathRoot + `DocCachePolicy`); value.Exists() {
		l := []DmDocCachePolicy{}
		if value := res.Get(`DocCachePolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := DmDocCachePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.DocCachePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDocCachePolicyObjectType}, l)
		} else {
			data.DocCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmDocCachePolicyObjectType})
		}
	} else {
		data.DocCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmDocCachePolicyObjectType})
	}
	if value := res.Get(pathRoot + `SchemaValidation`); value.Exists() {
		l := []DmSchemaValidation{}
		if value := res.Get(`SchemaValidation`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSchemaValidation{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SchemaValidation, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSchemaValidationObjectType}, l)
		} else {
			data.SchemaValidation = types.ListNull(types.ObjectType{AttrTypes: DmSchemaValidationObjectType})
		}
	} else {
		data.SchemaValidation = types.ListNull(types.ObjectType{AttrTypes: DmSchemaValidationObjectType})
	}
	if value := res.Get(pathRoot + `ScheduledRule`); value.Exists() {
		l := []DmScheduledRule{}
		if value := res.Get(`ScheduledRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmScheduledRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ScheduledRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmScheduledRuleObjectType}, l)
		} else {
			data.ScheduledRule = types.ListNull(types.ObjectType{AttrTypes: DmScheduledRuleObjectType})
		}
	} else {
		data.ScheduledRule = types.ListNull(types.ObjectType{AttrTypes: DmScheduledRuleObjectType})
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserAgent = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `JSONParserSettings`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JsonParserSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonParserSettings = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPConnPool`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapConnPool = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapConnPool = types.StringNull()
	}
}

func (data *XMLManager) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && !data.UrlRefreshPolicy.IsNull() {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRefreshPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CompileOptionsPolicy`); value.Exists() && !data.CompileOptionsPolicy.IsNull() {
		data.CompileOptionsPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CompileOptionsPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `CacheMemorySize`); value.Exists() && !data.CacheMemorySize.IsNull() {
		data.CacheMemorySize = types.Int64Value(value.Int())
	} else if data.CacheMemorySize.ValueInt64() != 2147483647 {
		data.CacheMemorySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CacheSize`); value.Exists() && !data.CacheSize.IsNull() {
		data.CacheSize = types.Int64Value(value.Int())
	} else if data.CacheSize.ValueInt64() != 256 {
		data.CacheSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SHA1Caching`); value.Exists() && !data.Sha1Caching.IsNull() {
		data.Sha1Caching = tfutils.BoolFromString(value.String())
	} else if !data.Sha1Caching.ValueBool() {
		data.Sha1Caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StaticDocumentCalls`); value.Exists() && !data.StaticDocumentCalls.IsNull() {
		data.StaticDocumentCalls = tfutils.BoolFromString(value.String())
	} else if !data.StaticDocumentCalls.ValueBool() {
		data.StaticDocumentCalls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SearchResults`); value.Exists() && !data.SearchResults.IsNull() {
		data.SearchResults = tfutils.BoolFromString(value.String())
	} else if !data.SearchResults.ValueBool() {
		data.SearchResults = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VirtualServers`); value.Exists() && !data.VirtualServers.IsNull() {
		data.VirtualServers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.VirtualServers = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ParserLimitsBytesScanned`); value.Exists() && !data.ParserLimitsBytesScanned.IsNull() {
		data.ParserLimitsBytesScanned = types.Int64Value(value.Int())
	} else if data.ParserLimitsBytesScanned.ValueInt64() != 4194304 {
		data.ParserLimitsBytesScanned = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsElementDepth`); value.Exists() && !data.ParserLimitsElementDepth.IsNull() {
		data.ParserLimitsElementDepth = types.Int64Value(value.Int())
	} else if data.ParserLimitsElementDepth.ValueInt64() != 512 {
		data.ParserLimitsElementDepth = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsAttributeCount`); value.Exists() && !data.ParserLimitsAttributeCount.IsNull() {
		data.ParserLimitsAttributeCount = types.Int64Value(value.Int())
	} else if data.ParserLimitsAttributeCount.ValueInt64() != 128 {
		data.ParserLimitsAttributeCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNodeSize`); value.Exists() && !data.ParserLimitsMaxNodeSize.IsNull() {
		data.ParserLimitsMaxNodeSize = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxNodeSize.ValueInt64() != 33554432 {
		data.ParserLimitsMaxNodeSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsForbidExternalReferences`); value.Exists() && !data.ParserLimitsForbidExternalReferences.IsNull() {
		data.ParserLimitsForbidExternalReferences = tfutils.BoolFromString(value.String())
	} else if !data.ParserLimitsForbidExternalReferences.ValueBool() {
		data.ParserLimitsForbidExternalReferences = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsExternalReferences`); value.Exists() && !data.ParserLimitsExternalReferences.IsNull() {
		data.ParserLimitsExternalReferences = tfutils.ParseStringFromGJSON(value)
	} else if data.ParserLimitsExternalReferences.ValueString() != "forbid" {
		data.ParserLimitsExternalReferences = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxPrefixes`); value.Exists() && !data.ParserLimitsMaxPrefixes.IsNull() {
		data.ParserLimitsMaxPrefixes = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxPrefixes.ValueInt64() != 1024 {
		data.ParserLimitsMaxPrefixes = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxNamespaces`); value.Exists() && !data.ParserLimitsMaxNamespaces.IsNull() {
		data.ParserLimitsMaxNamespaces = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxNamespaces.ValueInt64() != 1024 {
		data.ParserLimitsMaxNamespaces = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ParserLimitsMaxLocalNames`); value.Exists() && !data.ParserLimitsMaxLocalNames.IsNull() {
		data.ParserLimitsMaxLocalNames = types.Int64Value(value.Int())
	} else if data.ParserLimitsMaxLocalNames.ValueInt64() != 60000 {
		data.ParserLimitsMaxLocalNames = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DocCacheMaxDocs`); value.Exists() && !data.DocCacheMaxDocs.IsNull() {
		data.DocCacheMaxDocs = types.Int64Value(value.Int())
	} else if data.DocCacheMaxDocs.ValueInt64() != 5000 {
		data.DocCacheMaxDocs = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DocCacheSize`); value.Exists() && !data.DocCacheSize.IsNull() {
		data.DocCacheSize = types.Int64Value(value.Int())
	} else {
		data.DocCacheSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DocMaxWrites`); value.Exists() && !data.DocMaxWrites.IsNull() {
		data.DocMaxWrites = types.Int64Value(value.Int())
	} else if data.DocMaxWrites.ValueInt64() != 32768 {
		data.DocMaxWrites = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DocCachePolicy`); value.Exists() && !data.DocCachePolicy.IsNull() {
		l := []DmDocCachePolicy{}
		for _, v := range value.Array() {
			item := DmDocCachePolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.DocCachePolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmDocCachePolicyObjectType}, l)
		} else {
			data.DocCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmDocCachePolicyObjectType})
		}
	} else {
		data.DocCachePolicy = types.ListNull(types.ObjectType{AttrTypes: DmDocCachePolicyObjectType})
	}
	if value := res.Get(pathRoot + `SchemaValidation`); value.Exists() && !data.SchemaValidation.IsNull() {
		l := []DmSchemaValidation{}
		for _, v := range value.Array() {
			item := DmSchemaValidation{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SchemaValidation, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSchemaValidationObjectType}, l)
		} else {
			data.SchemaValidation = types.ListNull(types.ObjectType{AttrTypes: DmSchemaValidationObjectType})
		}
	} else {
		data.SchemaValidation = types.ListNull(types.ObjectType{AttrTypes: DmSchemaValidationObjectType})
	}
	if value := res.Get(pathRoot + `ScheduledRule`); value.Exists() && !data.ScheduledRule.IsNull() {
		l := []DmScheduledRule{}
		for _, v := range value.Array() {
			item := DmScheduledRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ScheduledRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmScheduledRuleObjectType}, l)
		} else {
			data.ScheduledRule = types.ListNull(types.ObjectType{AttrTypes: DmScheduledRuleObjectType})
		}
	} else {
		data.ScheduledRule = types.ListNull(types.ObjectType{AttrTypes: DmScheduledRuleObjectType})
	}
	if value := res.Get(pathRoot + `UserAgent`); value.Exists() && !data.UserAgent.IsNull() {
		data.UserAgent = tfutils.ParseStringFromGJSON(value)
	} else if data.UserAgent.ValueString() != "default" {
		data.UserAgent = types.StringNull()
	}
	if value := res.Get(pathRoot + `JSONParserSettings`); value.Exists() && !data.JsonParserSettings.IsNull() {
		data.JsonParserSettings = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonParserSettings = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPConnPool`); value.Exists() && !data.LdapConnPool.IsNull() {
		data.LdapConnPool = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapConnPool = types.StringNull()
	}
}
