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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WebAppRequest struct {
	Id                       types.String                `tfsdk:"id"`
	AppDomain                types.String                `tfsdk:"app_domain"`
	UserSummary              types.String                `tfsdk:"user_summary"`
	PolicyType               types.String                `tfsdk:"policy_type"`
	SslPolicy                types.String                `tfsdk:"ssl_policy"`
	Aaa                      types.String                `tfsdk:"aaa"`
	Sskey                    types.String                `tfsdk:"sskey"`
	RateLimiter              types.String                `tfsdk:"rate_limiter"`
	Acl                      types.String                `tfsdk:"acl"`
	OkMethods                *DmHTTPRequestMethods       `tfsdk:"ok_methods"`
	OkVersions               *DmHTTPVersionMask          `tfsdk:"ok_versions"`
	MinBodySize              types.Int64                 `tfsdk:"min_body_size"`
	MaxBodySize              types.Int64                 `tfsdk:"max_body_size"`
	XmlPolicy                types.String                `tfsdk:"xml_policy"`
	XmlRule                  types.String                `tfsdk:"xml_rule"`
	NonXmlPolicy             types.String                `tfsdk:"non_xml_policy"`
	NonXmlRule               types.String                `tfsdk:"non_xml_rule"`
	ErrorPolicy              types.String                `tfsdk:"error_policy"`
	SessionManagementProfile types.String                `tfsdk:"session_management_profile"`
	HeaderGnvc               types.String                `tfsdk:"header_gnvc"`
	UrlEncodedGnvc           types.String                `tfsdk:"url_encoded_gnvc"`
	QueryStringPolicy        types.String                `tfsdk:"query_string_policy"`
	QueryStringGnvc          types.String                `tfsdk:"query_string_gnvc"`
	SqlInjection             types.Bool                  `tfsdk:"sql_injection"`
	MaxUriSize               types.Int64                 `tfsdk:"max_uri_size"`
	UriFilterUnicode         types.Bool                  `tfsdk:"uri_filter_unicode"`
	UriFilterDotDot          types.Bool                  `tfsdk:"uri_filter_dot_dot"`
	UriFilterExe             types.Bool                  `tfsdk:"uri_filter_exe"`
	UriFilterFragment        types.String                `tfsdk:"uri_filter_fragment"`
	ContentTypes             types.List                  `tfsdk:"content_types"`
	MultipartFormData        *DmMultipartFormDataProfile `tfsdk:"multipart_form_data"`
	CookieProfile            *DmCookieProfile            `tfsdk:"cookie_profile"`
	ProcessAllCookie         types.Bool                  `tfsdk:"process_all_cookie"`
	CookieNameVector         types.List                  `tfsdk:"cookie_name_vector"`
	SqlInjectionPatternsFile types.String                `tfsdk:"sql_injection_patterns_file"`
	DependencyActions        []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebAppRequestXMLRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "xml_policy",
	AttrType:    "String",
	AttrDefault: "nothing",
	Value:       []string{"xml", "soap"},
}
var WebAppRequestNonXMLRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "non_xml_policy",
	AttrType:    "String",
	AttrDefault: "nothing",
	Value:       []string{"side", "binary"},
}
var WebAppRequestQueryStringGNVCCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "query_string_policy",
	AttrType:    "String",
	AttrDefault: "allow",
	Value:       []string{"require"},
}
var WebAppRequestCookieNameVectorCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "process_all_cookie",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}
var WebAppRequestSQLInjectionPatternsFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "sql_injection",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var WebAppRequestObjectType = map[string]attr.Type{
	"id":                          types.StringType,
	"app_domain":                  types.StringType,
	"user_summary":                types.StringType,
	"policy_type":                 types.StringType,
	"ssl_policy":                  types.StringType,
	"aaa":                         types.StringType,
	"sskey":                       types.StringType,
	"rate_limiter":                types.StringType,
	"acl":                         types.StringType,
	"ok_methods":                  types.ObjectType{AttrTypes: DmHTTPRequestMethodsObjectType},
	"ok_versions":                 types.ObjectType{AttrTypes: DmHTTPVersionMaskObjectType},
	"min_body_size":               types.Int64Type,
	"max_body_size":               types.Int64Type,
	"xml_policy":                  types.StringType,
	"xml_rule":                    types.StringType,
	"non_xml_policy":              types.StringType,
	"non_xml_rule":                types.StringType,
	"error_policy":                types.StringType,
	"session_management_profile":  types.StringType,
	"header_gnvc":                 types.StringType,
	"url_encoded_gnvc":            types.StringType,
	"query_string_policy":         types.StringType,
	"query_string_gnvc":           types.StringType,
	"sql_injection":               types.BoolType,
	"max_uri_size":                types.Int64Type,
	"uri_filter_unicode":          types.BoolType,
	"uri_filter_dot_dot":          types.BoolType,
	"uri_filter_exe":              types.BoolType,
	"uri_filter_fragment":         types.StringType,
	"content_types":               types.ListType{ElemType: types.StringType},
	"multipart_form_data":         types.ObjectType{AttrTypes: DmMultipartFormDataProfileObjectType},
	"cookie_profile":              types.ObjectType{AttrTypes: DmCookieProfileObjectType},
	"process_all_cookie":          types.BoolType,
	"cookie_name_vector":          types.ListType{ElemType: types.StringType},
	"sql_injection_patterns_file": types.StringType,
	"dependency_actions":          actions.ActionsListType,
}

func (data WebAppRequest) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebAppRequest"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebAppRequest) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.PolicyType.IsNull() {
		return false
	}
	if !data.SslPolicy.IsNull() {
		return false
	}
	if !data.Aaa.IsNull() {
		return false
	}
	if !data.Sskey.IsNull() {
		return false
	}
	if !data.RateLimiter.IsNull() {
		return false
	}
	if !data.Acl.IsNull() {
		return false
	}
	if data.OkMethods != nil {
		if !data.OkMethods.IsNull() {
			return false
		}
	}
	if data.OkVersions != nil {
		if !data.OkVersions.IsNull() {
			return false
		}
	}
	if !data.MinBodySize.IsNull() {
		return false
	}
	if !data.MaxBodySize.IsNull() {
		return false
	}
	if !data.XmlPolicy.IsNull() {
		return false
	}
	if !data.XmlRule.IsNull() {
		return false
	}
	if !data.NonXmlPolicy.IsNull() {
		return false
	}
	if !data.NonXmlRule.IsNull() {
		return false
	}
	if !data.ErrorPolicy.IsNull() {
		return false
	}
	if !data.SessionManagementProfile.IsNull() {
		return false
	}
	if !data.HeaderGnvc.IsNull() {
		return false
	}
	if !data.UrlEncodedGnvc.IsNull() {
		return false
	}
	if !data.QueryStringPolicy.IsNull() {
		return false
	}
	if !data.QueryStringGnvc.IsNull() {
		return false
	}
	if !data.SqlInjection.IsNull() {
		return false
	}
	if !data.MaxUriSize.IsNull() {
		return false
	}
	if !data.UriFilterUnicode.IsNull() {
		return false
	}
	if !data.UriFilterDotDot.IsNull() {
		return false
	}
	if !data.UriFilterExe.IsNull() {
		return false
	}
	if !data.UriFilterFragment.IsNull() {
		return false
	}
	if !data.ContentTypes.IsNull() {
		return false
	}
	if data.MultipartFormData != nil {
		if !data.MultipartFormData.IsNull() {
			return false
		}
	}
	if data.CookieProfile != nil {
		if !data.CookieProfile.IsNull() {
			return false
		}
	}
	if !data.ProcessAllCookie.IsNull() {
		return false
	}
	if !data.CookieNameVector.IsNull() {
		return false
	}
	if !data.SqlInjectionPatternsFile.IsNull() {
		return false
	}
	return true
}

func (data WebAppRequest) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.PolicyType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyType`, data.PolicyType.ValueString())
	}
	if !data.SslPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLPolicy`, data.SslPolicy.ValueString())
	}
	if !data.Aaa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAA`, data.Aaa.ValueString())
	}
	if !data.Sskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSKey`, data.Sskey.ValueString())
	}
	if !data.RateLimiter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimiter`, data.RateLimiter.ValueString())
	}
	if !data.Acl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACL`, data.Acl.ValueString())
	}
	if data.OkMethods != nil {
		if !data.OkMethods.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OKMethods`, data.OkMethods.ToBody(ctx, ""))
		}
	}
	if data.OkVersions != nil {
		if !data.OkVersions.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OKVersions`, data.OkVersions.ToBody(ctx, ""))
		}
	}
	if !data.MinBodySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MinBodySize`, data.MinBodySize.ValueInt64())
	}
	if !data.MaxBodySize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxBodySize`, data.MaxBodySize.ValueInt64())
	}
	if !data.XmlPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLPolicy`, data.XmlPolicy.ValueString())
	}
	if !data.XmlRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLRule`, data.XmlRule.ValueString())
	}
	if !data.NonXmlPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NonXMLPolicy`, data.NonXmlPolicy.ValueString())
	}
	if !data.NonXmlRule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NonXMLRule`, data.NonXmlRule.ValueString())
	}
	if !data.ErrorPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorPolicy`, data.ErrorPolicy.ValueString())
	}
	if !data.SessionManagementProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SessionManagementProfile`, data.SessionManagementProfile.ValueString())
	}
	if !data.HeaderGnvc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderGNVC`, data.HeaderGnvc.ValueString())
	}
	if !data.UrlEncodedGnvc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UrlEncodedGNVC`, data.UrlEncodedGnvc.ValueString())
	}
	if !data.QueryStringPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryStringPolicy`, data.QueryStringPolicy.ValueString())
	}
	if !data.QueryStringGnvc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryStringGNVC`, data.QueryStringGnvc.ValueString())
	}
	if !data.SqlInjection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLInjection`, tfutils.StringFromBool(data.SqlInjection, ""))
	}
	if !data.MaxUriSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxURISize`, data.MaxUriSize.ValueInt64())
	}
	if !data.UriFilterUnicode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URIFilterUnicode`, tfutils.StringFromBool(data.UriFilterUnicode, ""))
	}
	if !data.UriFilterDotDot.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URIFilterDotDot`, tfutils.StringFromBool(data.UriFilterDotDot, ""))
	}
	if !data.UriFilterExe.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URIFilterExe`, tfutils.StringFromBool(data.UriFilterExe, ""))
	}
	if !data.UriFilterFragment.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URIFilterFragment`, data.UriFilterFragment.ValueString())
	}
	if !data.ContentTypes.IsNull() {
		var dataValues []string
		data.ContentTypes.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`ContentTypes`+".-1", map[string]string{"value": val})
		}
	}
	if data.MultipartFormData != nil {
		if !data.MultipartFormData.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`MultipartFormData`, data.MultipartFormData.ToBody(ctx, ""))
		}
	}
	if data.CookieProfile != nil {
		if !data.CookieProfile.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`CookieProfile`, data.CookieProfile.ToBody(ctx, ""))
		}
	}
	if !data.ProcessAllCookie.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessAllCookie`, tfutils.StringFromBool(data.ProcessAllCookie, ""))
	}
	if !data.CookieNameVector.IsNull() {
		var dataValues []string
		data.CookieNameVector.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`CookieNameVector`+".-1", map[string]string{"value": val})
		}
	}
	if !data.SqlInjectionPatternsFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLInjectionPatternsFile`, data.SqlInjectionPatternsFile.ValueString())
	}
	return body
}

func (data *WebAppRequest) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PolicyType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyType = types.StringValue("admission")
	}
	if value := res.Get(pathRoot + `SSLPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslPolicy = types.StringValue("allow")
	}
	if value := res.Get(pathRoot + `AAA`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Aaa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aaa = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Sskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimiter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RateLimiter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimiter = types.StringNull()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OKMethods`); value.Exists() {
		data.OkMethods = &DmHTTPRequestMethods{}
		data.OkMethods.FromBody(ctx, "", value)
	} else {
		data.OkMethods = nil
	}
	if value := res.Get(pathRoot + `OKVersions`); value.Exists() {
		data.OkVersions = &DmHTTPVersionMask{}
		data.OkVersions.FromBody(ctx, "", value)
	} else {
		data.OkVersions = nil
	}
	if value := res.Get(pathRoot + `MinBodySize`); value.Exists() {
		data.MinBodySize = types.Int64Value(value.Int())
	} else {
		data.MinBodySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxBodySize`); value.Exists() {
		data.MaxBodySize = types.Int64Value(value.Int())
	} else {
		data.MaxBodySize = types.Int64Value(128000000)
	}
	if value := res.Get(pathRoot + `XMLPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlPolicy = types.StringValue("nothing")
	}
	if value := res.Get(pathRoot + `XMLRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `NonXMLPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NonXmlPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NonXmlPolicy = types.StringValue("nothing")
	}
	if value := res.Get(pathRoot + `NonXMLRule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NonXmlRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NonXmlRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SessionManagementProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SessionManagementProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SessionManagementProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderGNVC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `UrlEncodedGNVC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UrlEncodedGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlEncodedGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueryStringPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueryStringPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryStringPolicy = types.StringValue("allow")
	}
	if value := res.Get(pathRoot + `QueryStringGNVC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QueryStringGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryStringGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLInjection`); value.Exists() {
		data.SqlInjection = tfutils.BoolFromString(value.String())
	} else {
		data.SqlInjection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxURISize`); value.Exists() {
		data.MaxUriSize = types.Int64Value(value.Int())
	} else {
		data.MaxUriSize = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `URIFilterUnicode`); value.Exists() {
		data.UriFilterUnicode = tfutils.BoolFromString(value.String())
	} else {
		data.UriFilterUnicode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterDotDot`); value.Exists() {
		data.UriFilterDotDot = tfutils.BoolFromString(value.String())
	} else {
		data.UriFilterDotDot = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterExe`); value.Exists() {
		data.UriFilterExe = tfutils.BoolFromString(value.String())
	} else {
		data.UriFilterExe = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterFragment`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UriFilterFragment = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UriFilterFragment = types.StringValue("truncate")
	}
	if value := res.Get(pathRoot + `ContentTypes`); value.Exists() {
		data.ContentTypes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ContentTypes = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MultipartFormData`); value.Exists() {
		data.MultipartFormData = &DmMultipartFormDataProfile{}
		data.MultipartFormData.FromBody(ctx, "", value)
	} else {
		data.MultipartFormData = nil
	}
	if value := res.Get(pathRoot + `CookieProfile`); value.Exists() {
		data.CookieProfile = &DmCookieProfile{}
		data.CookieProfile.FromBody(ctx, "", value)
	} else {
		data.CookieProfile = nil
	}
	if value := res.Get(pathRoot + `ProcessAllCookie`); value.Exists() {
		data.ProcessAllCookie = tfutils.BoolFromString(value.String())
	} else {
		data.ProcessAllCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieNameVector`); value.Exists() {
		data.CookieNameVector = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CookieNameVector = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SQLInjectionPatternsFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SqlInjectionPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlInjectionPatternsFile = types.StringValue("store:///SQL-Injection-Patterns.xml")
	}
}

func (data *WebAppRequest) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `PolicyType`); value.Exists() && !data.PolicyType.IsNull() {
		data.PolicyType = tfutils.ParseStringFromGJSON(value)
	} else if data.PolicyType.ValueString() != "admission" {
		data.PolicyType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLPolicy`); value.Exists() && !data.SslPolicy.IsNull() {
		data.SslPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.SslPolicy.ValueString() != "allow" {
		data.SslPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAA`); value.Exists() && !data.Aaa.IsNull() {
		data.Aaa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aaa = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && !data.Sskey.IsNull() {
		data.Sskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimiter`); value.Exists() && !data.RateLimiter.IsNull() {
		data.RateLimiter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RateLimiter = types.StringNull()
	}
	if value := res.Get(pathRoot + `ACL`); value.Exists() && !data.Acl.IsNull() {
		data.Acl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Acl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OKMethods`); value.Exists() {
		data.OkMethods.UpdateFromBody(ctx, "", value)
	} else {
		data.OkMethods = nil
	}
	if value := res.Get(pathRoot + `OKVersions`); value.Exists() {
		data.OkVersions.UpdateFromBody(ctx, "", value)
	} else {
		data.OkVersions = nil
	}
	if value := res.Get(pathRoot + `MinBodySize`); value.Exists() && !data.MinBodySize.IsNull() {
		data.MinBodySize = types.Int64Value(value.Int())
	} else {
		data.MinBodySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxBodySize`); value.Exists() && !data.MaxBodySize.IsNull() {
		data.MaxBodySize = types.Int64Value(value.Int())
	} else if data.MaxBodySize.ValueInt64() != 128000000 {
		data.MaxBodySize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XMLPolicy`); value.Exists() && !data.XmlPolicy.IsNull() {
		data.XmlPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlPolicy.ValueString() != "nothing" {
		data.XmlPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `XMLRule`); value.Exists() && !data.XmlRule.IsNull() {
		data.XmlRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `NonXMLPolicy`); value.Exists() && !data.NonXmlPolicy.IsNull() {
		data.NonXmlPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.NonXmlPolicy.ValueString() != "nothing" {
		data.NonXmlPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `NonXMLRule`); value.Exists() && !data.NonXmlRule.IsNull() {
		data.NonXmlRule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NonXmlRule = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorPolicy`); value.Exists() && !data.ErrorPolicy.IsNull() {
		data.ErrorPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SessionManagementProfile`); value.Exists() && !data.SessionManagementProfile.IsNull() {
		data.SessionManagementProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SessionManagementProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `HeaderGNVC`); value.Exists() && !data.HeaderGnvc.IsNull() {
		data.HeaderGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `UrlEncodedGNVC`); value.Exists() && !data.UrlEncodedGnvc.IsNull() {
		data.UrlEncodedGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UrlEncodedGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueryStringPolicy`); value.Exists() && !data.QueryStringPolicy.IsNull() {
		data.QueryStringPolicy = tfutils.ParseStringFromGJSON(value)
	} else if data.QueryStringPolicy.ValueString() != "allow" {
		data.QueryStringPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `QueryStringGNVC`); value.Exists() && !data.QueryStringGnvc.IsNull() {
		data.QueryStringGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QueryStringGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLInjection`); value.Exists() && !data.SqlInjection.IsNull() {
		data.SqlInjection = tfutils.BoolFromString(value.String())
	} else if data.SqlInjection.ValueBool() {
		data.SqlInjection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `MaxURISize`); value.Exists() && !data.MaxUriSize.IsNull() {
		data.MaxUriSize = types.Int64Value(value.Int())
	} else if data.MaxUriSize.ValueInt64() != 1024 {
		data.MaxUriSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `URIFilterUnicode`); value.Exists() && !data.UriFilterUnicode.IsNull() {
		data.UriFilterUnicode = tfutils.BoolFromString(value.String())
	} else if !data.UriFilterUnicode.ValueBool() {
		data.UriFilterUnicode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterDotDot`); value.Exists() && !data.UriFilterDotDot.IsNull() {
		data.UriFilterDotDot = tfutils.BoolFromString(value.String())
	} else if !data.UriFilterDotDot.ValueBool() {
		data.UriFilterDotDot = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterExe`); value.Exists() && !data.UriFilterExe.IsNull() {
		data.UriFilterExe = tfutils.BoolFromString(value.String())
	} else if !data.UriFilterExe.ValueBool() {
		data.UriFilterExe = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URIFilterFragment`); value.Exists() && !data.UriFilterFragment.IsNull() {
		data.UriFilterFragment = tfutils.ParseStringFromGJSON(value)
	} else if data.UriFilterFragment.ValueString() != "truncate" {
		data.UriFilterFragment = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentTypes`); value.Exists() && !data.ContentTypes.IsNull() {
		data.ContentTypes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ContentTypes = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MultipartFormData`); value.Exists() {
		data.MultipartFormData.UpdateFromBody(ctx, "", value)
	} else {
		data.MultipartFormData = nil
	}
	if value := res.Get(pathRoot + `CookieProfile`); value.Exists() {
		data.CookieProfile.UpdateFromBody(ctx, "", value)
	} else {
		data.CookieProfile = nil
	}
	if value := res.Get(pathRoot + `ProcessAllCookie`); value.Exists() && !data.ProcessAllCookie.IsNull() {
		data.ProcessAllCookie = tfutils.BoolFromString(value.String())
	} else if !data.ProcessAllCookie.ValueBool() {
		data.ProcessAllCookie = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CookieNameVector`); value.Exists() && !data.CookieNameVector.IsNull() {
		data.CookieNameVector = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.CookieNameVector = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SQLInjectionPatternsFile`); value.Exists() && !data.SqlInjectionPatternsFile.IsNull() {
		data.SqlInjectionPatternsFile = tfutils.ParseStringFromGJSON(value)
	} else if data.SqlInjectionPatternsFile.ValueString() != "store:///SQL-Injection-Patterns.xml" {
		data.SqlInjectionPatternsFile = types.StringNull()
	}
}
