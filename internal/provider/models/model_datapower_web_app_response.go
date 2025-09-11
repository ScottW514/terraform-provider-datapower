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

type WebAppResponse struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	PolicyType        types.String                `tfsdk:"policy_type"`
	OkCodes           *DmHTTPResponseCodes        `tfsdk:"ok_codes"`
	OkVersions        *DmHTTPVersionMask          `tfsdk:"ok_versions"`
	MinBodySize       types.Int64                 `tfsdk:"min_body_size"`
	MaxBodySize       types.Int64                 `tfsdk:"max_body_size"`
	HeaderGnvc        types.String                `tfsdk:"header_gnvc"`
	ContentTypes      types.List                  `tfsdk:"content_types"`
	XmlPolicy         types.String                `tfsdk:"xml_policy"`
	XmlRule           types.String                `tfsdk:"xml_rule"`
	NonXmlPolicy      types.String                `tfsdk:"non_xml_policy"`
	NonXmlRule        types.String                `tfsdk:"non_xml_rule"`
	ErrorPolicy       types.String                `tfsdk:"error_policy"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebAppResponseXMLRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "xml_policy",
	AttrType:    "String",
	AttrDefault: "nothing",
	Value:       []string{"xml", "soap"},
}
var WebAppResponseNonXMLRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "non_xml_policy",
	AttrType:    "String",
	AttrDefault: "nothing",
	Value:       []string{"side", "binary"},
}
var WebAppResponseXMLRuleIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var WebAppResponseNonXMLRuleIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var WebAppResponseObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"policy_type":        types.StringType,
	"ok_codes":           types.ObjectType{AttrTypes: DmHTTPResponseCodesObjectType},
	"ok_versions":        types.ObjectType{AttrTypes: DmHTTPVersionMaskObjectType},
	"min_body_size":      types.Int64Type,
	"max_body_size":      types.Int64Type,
	"header_gnvc":        types.StringType,
	"content_types":      types.ListType{ElemType: types.StringType},
	"xml_policy":         types.StringType,
	"xml_rule":           types.StringType,
	"non_xml_policy":     types.StringType,
	"non_xml_rule":       types.StringType,
	"error_policy":       types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data WebAppResponse) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebAppResponse"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebAppResponse) IsNull() bool {
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
	if data.OkCodes != nil {
		if !data.OkCodes.IsNull() {
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
	if !data.HeaderGnvc.IsNull() {
		return false
	}
	if !data.ContentTypes.IsNull() {
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
	return true
}

func (data WebAppResponse) ToBody(ctx context.Context, pathRoot string) string {
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
	if data.OkCodes != nil {
		if !data.OkCodes.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OKCodes`, data.OkCodes.ToBody(ctx, ""))
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
	if !data.HeaderGnvc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HeaderGNVC`, data.HeaderGnvc.ValueString())
	}
	if !data.ContentTypes.IsNull() {
		var dataValues []string
		data.ContentTypes.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`ContentTypes`+".-1", map[string]string{"value": val})
		}
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
	return body
}

func (data *WebAppResponse) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OKCodes`); value.Exists() {
		data.OkCodes = &DmHTTPResponseCodes{}
		data.OkCodes.FromBody(ctx, "", value)
	} else {
		data.OkCodes = nil
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
	if value := res.Get(pathRoot + `HeaderGNVC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HeaderGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentTypes`); value.Exists() {
		data.ContentTypes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ContentTypes = types.ListNull(types.StringType)
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
}

func (data *WebAppResponse) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `OKCodes`); value.Exists() {
		data.OkCodes.UpdateFromBody(ctx, "", value)
	} else {
		data.OkCodes = nil
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
	if value := res.Get(pathRoot + `HeaderGNVC`); value.Exists() && !data.HeaderGnvc.IsNull() {
		data.HeaderGnvc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HeaderGnvc = types.StringNull()
	}
	if value := res.Get(pathRoot + `ContentTypes`); value.Exists() && !data.ContentTypes.IsNull() {
		data.ContentTypes = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ContentTypes = types.ListNull(types.StringType)
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
}
