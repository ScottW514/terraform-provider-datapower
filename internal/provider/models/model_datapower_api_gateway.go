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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type APIGateway struct {
	Id                             types.String                `tfsdk:"id"`
	AppDomain                      types.String                `tfsdk:"app_domain"`
	UserSummary                    types.String                `tfsdk:"user_summary"`
	GatewayServiceName             types.String                `tfsdk:"gateway_service_name"`
	FrontProtocol                  types.List                  `tfsdk:"front_protocol"`
	UrlRefreshPolicy               types.String                `tfsdk:"url_refresh_policy"`
	CacheMemorySize                types.Int64                 `tfsdk:"cache_memory_size"`
	CacheSize                      types.Int64                 `tfsdk:"cache_size"`
	Sha1caching                    types.Bool                  `tfsdk:"sha1caching"`
	StaticDocumentCalls            types.Bool                  `tfsdk:"static_document_calls"`
	VirtualServers                 types.List                  `tfsdk:"virtual_servers"`
	DocCacheMaxDocs                types.Int64                 `tfsdk:"doc_cache_max_docs"`
	DocCacheSize                   types.Int64                 `tfsdk:"doc_cache_size"`
	DocMaxWrites                   types.Int64                 `tfsdk:"doc_max_writes"`
	DocCachePolicy                 types.List                  `tfsdk:"doc_cache_policy"`
	ScheduledRule                  types.List                  `tfsdk:"scheduled_rule"`
	ApiCollection                  types.List                  `tfsdk:"api_collection"`
	ShareRateLimitCount            types.String                `tfsdk:"share_rate_limit_count"`
	AssemblyBurstLimit             types.List                  `tfsdk:"assembly_burst_limit"`
	AssemblyRateLimit              types.List                  `tfsdk:"assembly_rate_limit"`
	AssemblyCountLimit             types.List                  `tfsdk:"assembly_count_limit"`
	LdapConnPool                   types.String                `tfsdk:"ldap_conn_pool"`
	ProxyPolicies                  types.List                  `tfsdk:"proxy_policies"`
	FrontTimeout                   types.Int64                 `tfsdk:"front_timeout"`
	FrontPersistentTimeout         types.Int64                 `tfsdk:"front_persistent_timeout"`
	OpenTelemetry                  types.String                `tfsdk:"open_telemetry"`
	OpenTelemetryResourceAttribute types.List                  `tfsdk:"open_telemetry_resource_attribute"`
	DependencyActions              []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APIGatewayObjectType = map[string]attr.Type{
	"id":                                types.StringType,
	"app_domain":                        types.StringType,
	"user_summary":                      types.StringType,
	"gateway_service_name":              types.StringType,
	"front_protocol":                    types.ListType{ElemType: types.StringType},
	"url_refresh_policy":                types.StringType,
	"cache_memory_size":                 types.Int64Type,
	"cache_size":                        types.Int64Type,
	"sha1caching":                       types.BoolType,
	"static_document_calls":             types.BoolType,
	"virtual_servers":                   types.ListType{ElemType: types.StringType},
	"doc_cache_max_docs":                types.Int64Type,
	"doc_cache_size":                    types.Int64Type,
	"doc_max_writes":                    types.Int64Type,
	"doc_cache_policy":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmDocCachePolicyObjectType}},
	"scheduled_rule":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmScheduledRuleObjectType}},
	"api_collection":                    types.ListType{ElemType: types.StringType},
	"share_rate_limit_count":            types.StringType,
	"assembly_burst_limit":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIBurstLimitObjectType}},
	"assembly_rate_limit":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIRateLimitObjectType}},
	"assembly_count_limit":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPICountLimitObjectType}},
	"ldap_conn_pool":                    types.StringType,
	"proxy_policies":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}},
	"front_timeout":                     types.Int64Type,
	"front_persistent_timeout":          types.Int64Type,
	"open_telemetry":                    types.StringType,
	"open_telemetry_resource_attribute": types.ListType{ElemType: types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}},
	"dependency_actions":                actions.ActionsListType,
}

func (data APIGateway) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIGateway"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIGateway) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.GatewayServiceName.IsNull() {
		return false
	}
	if !data.FrontProtocol.IsNull() {
		return false
	}
	if !data.UrlRefreshPolicy.IsNull() {
		return false
	}
	if !data.CacheMemorySize.IsNull() {
		return false
	}
	if !data.CacheSize.IsNull() {
		return false
	}
	if !data.Sha1caching.IsNull() {
		return false
	}
	if !data.StaticDocumentCalls.IsNull() {
		return false
	}
	if !data.VirtualServers.IsNull() {
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
	if !data.ScheduledRule.IsNull() {
		return false
	}
	if !data.ApiCollection.IsNull() {
		return false
	}
	if !data.ShareRateLimitCount.IsNull() {
		return false
	}
	if !data.AssemblyBurstLimit.IsNull() {
		return false
	}
	if !data.AssemblyRateLimit.IsNull() {
		return false
	}
	if !data.AssemblyCountLimit.IsNull() {
		return false
	}
	if !data.LdapConnPool.IsNull() {
		return false
	}
	if !data.ProxyPolicies.IsNull() {
		return false
	}
	if !data.FrontTimeout.IsNull() {
		return false
	}
	if !data.FrontPersistentTimeout.IsNull() {
		return false
	}
	if !data.OpenTelemetry.IsNull() {
		return false
	}
	if !data.OpenTelemetryResourceAttribute.IsNull() {
		return false
	}
	return true
}

func (data APIGateway) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.GatewayServiceName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayServiceName`, data.GatewayServiceName.ValueString())
	}
	if !data.FrontProtocol.IsNull() {
		var dataValues []string
		data.FrontProtocol.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`FrontProtocol`+".-1", map[string]string{"value": val})
		}
	}
	if !data.UrlRefreshPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URLRefreshPolicy`, data.UrlRefreshPolicy.ValueString())
	}
	if !data.CacheMemorySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheMemorySize`, data.CacheMemorySize.ValueInt64())
	}
	if !data.CacheSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CacheSize`, data.CacheSize.ValueInt64())
	}
	if !data.Sha1caching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SHA1Caching`, tfutils.StringFromBool(data.Sha1caching, ""))
	}
	if !data.StaticDocumentCalls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StaticDocumentCalls`, tfutils.StringFromBool(data.StaticDocumentCalls, ""))
	}
	if !data.VirtualServers.IsNull() {
		var dataValues []string
		data.VirtualServers.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`VirtualServers`+".-1", map[string]string{"value": val})
		}
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
		var dataValues []DmDocCachePolicy
		data.DocCachePolicy.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`DocCachePolicy`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ScheduledRule.IsNull() {
		var dataValues []DmScheduledRule
		data.ScheduledRule.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ScheduledRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ApiCollection.IsNull() {
		var dataValues []string
		data.ApiCollection.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`APICollection`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ShareRateLimitCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ShareRateLimitCount`, data.ShareRateLimitCount.ValueString())
	}
	if !data.AssemblyBurstLimit.IsNull() {
		var dataValues []DmAPIBurstLimit
		data.AssemblyBurstLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyBurstLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyRateLimit.IsNull() {
		var dataValues []DmAPIRateLimit
		data.AssemblyRateLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyRateLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.AssemblyCountLimit.IsNull() {
		var dataValues []DmAPICountLimit
		data.AssemblyCountLimit.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`AssemblyCountLimit`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.LdapConnPool.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPConnPool`, data.LdapConnPool.ValueString())
	}
	if !data.ProxyPolicies.IsNull() {
		var dataValues []DmAPIProxyPolicy
		data.ProxyPolicies.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicies`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.FrontTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontTimeout`, data.FrontTimeout.ValueInt64())
	}
	if !data.FrontPersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontPersistentTimeout`, data.FrontPersistentTimeout.ValueInt64())
	}
	if !data.OpenTelemetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OpenTelemetry`, data.OpenTelemetry.ValueString())
	}
	if !data.OpenTelemetryResourceAttribute.IsNull() {
		var dataValues []DmOpenTelemetryResourceAttribute
		data.OpenTelemetryResourceAttribute.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`OpenTelemetryResourceAttribute`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *APIGateway) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `GatewayServiceName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayServiceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayServiceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() {
		data.FrontProtocol = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FrontProtocol = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRefreshPolicy = types.StringNull()
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
		data.Sha1caching = tfutils.BoolFromString(value.String())
	} else {
		data.Sha1caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StaticDocumentCalls`); value.Exists() {
		data.StaticDocumentCalls = tfutils.BoolFromString(value.String())
	} else {
		data.StaticDocumentCalls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VirtualServers`); value.Exists() {
		data.VirtualServers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.VirtualServers = types.ListNull(types.StringType)
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
	if value := res.Get(pathRoot + `APICollection`); value.Exists() {
		data.ApiCollection = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApiCollection = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ShareRateLimitCount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ShareRateLimitCount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShareRateLimitCount = types.StringValue("yes")
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
	if value := res.Get(pathRoot + `LDAPConnPool`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapConnPool = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapConnPool = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() {
		l := []DmAPIProxyPolicy{}
		if value := res.Get(`ProxyPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIProxyPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontTimeout = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontPersistentTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `OpenTelemetry`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OpenTelemetry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OpenTelemetry = types.StringNull()
	}
	if value := res.Get(pathRoot + `OpenTelemetryResourceAttribute`); value.Exists() {
		l := []DmOpenTelemetryResourceAttribute{}
		if value := res.Get(`OpenTelemetryResourceAttribute`); value.Exists() {
			for _, v := range value.Array() {
				item := DmOpenTelemetryResourceAttribute{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OpenTelemetryResourceAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}, l)
		} else {
			data.OpenTelemetryResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
		}
	} else {
		data.OpenTelemetryResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
	}
}

func (data *APIGateway) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `GatewayServiceName`); value.Exists() && !data.GatewayServiceName.IsNull() {
		data.GatewayServiceName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayServiceName = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && !data.FrontProtocol.IsNull() {
		data.FrontProtocol = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.FrontProtocol = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `URLRefreshPolicy`); value.Exists() && !data.UrlRefreshPolicy.IsNull() {
		data.UrlRefreshPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlRefreshPolicy = types.StringNull()
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
	if value := res.Get(pathRoot + `SHA1Caching`); value.Exists() && !data.Sha1caching.IsNull() {
		data.Sha1caching = tfutils.BoolFromString(value.String())
	} else if !data.Sha1caching.ValueBool() {
		data.Sha1caching = types.BoolNull()
	}
	if value := res.Get(pathRoot + `StaticDocumentCalls`); value.Exists() && !data.StaticDocumentCalls.IsNull() {
		data.StaticDocumentCalls = tfutils.BoolFromString(value.String())
	} else if !data.StaticDocumentCalls.ValueBool() {
		data.StaticDocumentCalls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `VirtualServers`); value.Exists() && !data.VirtualServers.IsNull() {
		data.VirtualServers = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.VirtualServers = types.ListNull(types.StringType)
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
	if value := res.Get(pathRoot + `APICollection`); value.Exists() && !data.ApiCollection.IsNull() {
		data.ApiCollection = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApiCollection = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `ShareRateLimitCount`); value.Exists() && !data.ShareRateLimitCount.IsNull() {
		data.ShareRateLimitCount = tfutils.ParseStringFromGJSON(value)
	} else if data.ShareRateLimitCount.ValueString() != "yes" {
		data.ShareRateLimitCount = types.StringNull()
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
	if value := res.Get(pathRoot + `LDAPConnPool`); value.Exists() && !data.LdapConnPool.IsNull() {
		data.LdapConnPool = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapConnPool = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() && !data.ProxyPolicies.IsNull() {
		l := []DmAPIProxyPolicy{}
		for _, v := range value.Array() {
			item := DmAPIProxyPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `FrontTimeout`); value.Exists() && !data.FrontTimeout.IsNull() {
		data.FrontTimeout = types.Int64Value(value.Int())
	} else if data.FrontTimeout.ValueInt64() != 120 {
		data.FrontTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FrontPersistentTimeout`); value.Exists() && !data.FrontPersistentTimeout.IsNull() {
		data.FrontPersistentTimeout = types.Int64Value(value.Int())
	} else if data.FrontPersistentTimeout.ValueInt64() != 180 {
		data.FrontPersistentTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `OpenTelemetry`); value.Exists() && !data.OpenTelemetry.IsNull() {
		data.OpenTelemetry = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OpenTelemetry = types.StringNull()
	}
	if value := res.Get(pathRoot + `OpenTelemetryResourceAttribute`); value.Exists() && !data.OpenTelemetryResourceAttribute.IsNull() {
		l := []DmOpenTelemetryResourceAttribute{}
		for _, v := range value.Array() {
			item := DmOpenTelemetryResourceAttribute{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OpenTelemetryResourceAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}, l)
		} else {
			data.OpenTelemetryResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
		}
	} else {
		data.OpenTelemetryResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
	}
}
