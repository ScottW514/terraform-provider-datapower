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

type XPathRoutingMap struct {
	Id                types.String      `tfsdk:"id"`
	AppDomain         types.String      `tfsdk:"app_domain"`
	XPathRoutingRules types.List        `tfsdk:"x_path_routing_rules"`
	NameSpaceMappings types.List        `tfsdk:"name_space_mappings"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	DependencyActions []*actions.Action `tfsdk:"dependency_actions"`
}

var XPathRoutingMapObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"x_path_routing_rules": types.ListType{ElemType: types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType}},
	"name_space_mappings":  types.ListType{ElemType: types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}},
	"user_summary":         types.StringType,
	"dependency_actions":   actions.ActionsListType,
}

func (data XPathRoutingMap) GetPath() string {
	rest_path := "/mgmt/config/{domain}/XPathRoutingMap"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data XPathRoutingMap) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.XPathRoutingRules.IsNull() {
		return false
	}
	if !data.NameSpaceMappings.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data XPathRoutingMap) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.XPathRoutingRules.IsNull() {
		var values []DmXPathRoutingRule
		data.XPathRoutingRules.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`XPathRoutingRules`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.NameSpaceMappings.IsNull() {
		var values []DmNamespaceMapping
		data.NameSpaceMappings.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`NameSpaceMappings`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *XPathRoutingMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathRoutingRules`); value.Exists() {
		l := []DmXPathRoutingRule{}
		if value := res.Get(`XPathRoutingRules`); value.Exists() {
			for _, v := range value.Array() {
				item := DmXPathRoutingRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.XPathRoutingRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType}, l)
		} else {
			data.XPathRoutingRules = types.ListNull(types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType})
		}
	} else {
		data.XPathRoutingRules = types.ListNull(types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType})
	}
	if value := res.Get(pathRoot + `NameSpaceMappings`); value.Exists() {
		l := []DmNamespaceMapping{}
		if value := res.Get(`NameSpaceMappings`); value.Exists() {
			for _, v := range value.Array() {
				item := DmNamespaceMapping{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.NameSpaceMappings, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}, l)
		} else {
			data.NameSpaceMappings = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
		}
	} else {
		data.NameSpaceMappings = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *XPathRoutingMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPathRoutingRules`); value.Exists() && !data.XPathRoutingRules.IsNull() {
		l := []DmXPathRoutingRule{}
		for _, v := range value.Array() {
			item := DmXPathRoutingRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.XPathRoutingRules, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType}, l)
		} else {
			data.XPathRoutingRules = types.ListNull(types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType})
		}
	} else {
		data.XPathRoutingRules = types.ListNull(types.ObjectType{AttrTypes: DmXPathRoutingRuleObjectType})
	}
	if value := res.Get(pathRoot + `NameSpaceMappings`); value.Exists() && !data.NameSpaceMappings.IsNull() {
		l := []DmNamespaceMapping{}
		for _, v := range value.Array() {
			item := DmNamespaceMapping{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.NameSpaceMappings, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}, l)
		} else {
			data.NameSpaceMappings = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
		}
	} else {
		data.NameSpaceMappings = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
