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

type WSEndpointRewritePolicy struct {
	Id                                types.String      `tfsdk:"id"`
	AppDomain                         types.String      `tfsdk:"app_domain"`
	UserSummary                       types.String      `tfsdk:"user_summary"`
	WsEndpointLocalRewriteRule        types.List        `tfsdk:"ws_endpoint_local_rewrite_rule"`
	WsEndpointRemoteRewriteRule       types.List        `tfsdk:"ws_endpoint_remote_rewrite_rule"`
	WsEndpointPublishRewriteRule      types.List        `tfsdk:"ws_endpoint_publish_rewrite_rule"`
	WsEndpointSubscriptionLocalRule   types.List        `tfsdk:"ws_endpoint_subscription_local_rule"`
	WsEndpointSubscriptionRemoteRule  types.List        `tfsdk:"ws_endpoint_subscription_remote_rule"`
	WsEndpointSubscriptionPublishRule types.List        `tfsdk:"ws_endpoint_subscription_publish_rule"`
	DependencyActions                 []*actions.Action `tfsdk:"dependency_actions"`
}

var WSEndpointRewritePolicyObjectType = map[string]attr.Type{
	"id":                                    types.StringType,
	"app_domain":                            types.StringType,
	"user_summary":                          types.StringType,
	"ws_endpoint_local_rewrite_rule":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType}},
	"ws_endpoint_remote_rewrite_rule":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType}},
	"ws_endpoint_publish_rewrite_rule":      types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType}},
	"ws_endpoint_subscription_local_rule":   types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType}},
	"ws_endpoint_subscription_remote_rule":  types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType}},
	"ws_endpoint_subscription_publish_rule": types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType}},
	"dependency_actions":                    actions.ActionsListType,
}

func (data WSEndpointRewritePolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WSEndpointRewritePolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WSEndpointRewritePolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.WsEndpointLocalRewriteRule.IsNull() {
		return false
	}
	if !data.WsEndpointRemoteRewriteRule.IsNull() {
		return false
	}
	if !data.WsEndpointPublishRewriteRule.IsNull() {
		return false
	}
	if !data.WsEndpointSubscriptionLocalRule.IsNull() {
		return false
	}
	if !data.WsEndpointSubscriptionRemoteRule.IsNull() {
		return false
	}
	if !data.WsEndpointSubscriptionPublishRule.IsNull() {
		return false
	}
	return true
}

func (data WSEndpointRewritePolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.WsEndpointLocalRewriteRule.IsNull() {
		var values []DmWSEndpointLocalRewriteRule
		data.WsEndpointLocalRewriteRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointLocalRewriteRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsEndpointRemoteRewriteRule.IsNull() {
		var values []DmWSEndpointRemoteRewriteRule
		data.WsEndpointRemoteRewriteRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointRemoteRewriteRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsEndpointPublishRewriteRule.IsNull() {
		var values []DmWSEndpointPublishRewriteRule
		data.WsEndpointPublishRewriteRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointPublishRewriteRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsEndpointSubscriptionLocalRule.IsNull() {
		var values []DmWSEndpointSubscriptionLocalRule
		data.WsEndpointSubscriptionLocalRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointSubscriptionLocalRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsEndpointSubscriptionRemoteRule.IsNull() {
		var values []DmWSEndpointSubscriptionRemoteRule
		data.WsEndpointSubscriptionRemoteRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointSubscriptionRemoteRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.WsEndpointSubscriptionPublishRule.IsNull() {
		var values []DmWSEndpointSubscriptionPublishRule
		data.WsEndpointSubscriptionPublishRule.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`WSEndpointSubscriptionPublishRule`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *WSEndpointRewritePolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `WSEndpointLocalRewriteRule`); value.Exists() {
		l := []DmWSEndpointLocalRewriteRule{}
		if value := res.Get(`WSEndpointLocalRewriteRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointLocalRewriteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointLocalRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointLocalRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointLocalRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointRemoteRewriteRule`); value.Exists() {
		l := []DmWSEndpointRemoteRewriteRule{}
		if value := res.Get(`WSEndpointRemoteRewriteRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointRemoteRewriteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointRemoteRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointRemoteRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointRemoteRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointPublishRewriteRule`); value.Exists() {
		l := []DmWSEndpointPublishRewriteRule{}
		if value := res.Get(`WSEndpointPublishRewriteRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointPublishRewriteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointPublishRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointPublishRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointPublishRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionLocalRule`); value.Exists() {
		l := []DmWSEndpointSubscriptionLocalRule{}
		if value := res.Get(`WSEndpointSubscriptionLocalRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointSubscriptionLocalRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionLocalRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionLocalRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionLocalRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionRemoteRule`); value.Exists() {
		l := []DmWSEndpointSubscriptionRemoteRule{}
		if value := res.Get(`WSEndpointSubscriptionRemoteRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointSubscriptionRemoteRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionRemoteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionRemoteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionRemoteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionPublishRule`); value.Exists() {
		l := []DmWSEndpointSubscriptionPublishRule{}
		if value := res.Get(`WSEndpointSubscriptionPublishRule`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSEndpointSubscriptionPublishRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionPublishRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionPublishRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionPublishRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType})
	}
}

func (data *WSEndpointRewritePolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `WSEndpointLocalRewriteRule`); value.Exists() && !data.WsEndpointLocalRewriteRule.IsNull() {
		l := []DmWSEndpointLocalRewriteRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointLocalRewriteRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointLocalRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointLocalRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointLocalRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointLocalRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointRemoteRewriteRule`); value.Exists() && !data.WsEndpointRemoteRewriteRule.IsNull() {
		l := []DmWSEndpointRemoteRewriteRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointRemoteRewriteRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointRemoteRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointRemoteRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointRemoteRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointRemoteRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointPublishRewriteRule`); value.Exists() && !data.WsEndpointPublishRewriteRule.IsNull() {
		l := []DmWSEndpointPublishRewriteRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointPublishRewriteRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointPublishRewriteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType}, l)
		} else {
			data.WsEndpointPublishRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType})
		}
	} else {
		data.WsEndpointPublishRewriteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointPublishRewriteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionLocalRule`); value.Exists() && !data.WsEndpointSubscriptionLocalRule.IsNull() {
		l := []DmWSEndpointSubscriptionLocalRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointSubscriptionLocalRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionLocalRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionLocalRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionLocalRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionLocalRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionRemoteRule`); value.Exists() && !data.WsEndpointSubscriptionRemoteRule.IsNull() {
		l := []DmWSEndpointSubscriptionRemoteRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointSubscriptionRemoteRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionRemoteRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionRemoteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionRemoteRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionRemoteRuleObjectType})
	}
	if value := res.Get(pathRoot + `WSEndpointSubscriptionPublishRule`); value.Exists() && !data.WsEndpointSubscriptionPublishRule.IsNull() {
		l := []DmWSEndpointSubscriptionPublishRule{}
		for _, v := range value.Array() {
			item := DmWSEndpointSubscriptionPublishRule{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.WsEndpointSubscriptionPublishRule, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType}, l)
		} else {
			data.WsEndpointSubscriptionPublishRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType})
		}
	} else {
		data.WsEndpointSubscriptionPublishRule = types.ListNull(types.ObjectType{AttrTypes: DmWSEndpointSubscriptionPublishRuleObjectType})
	}
}
