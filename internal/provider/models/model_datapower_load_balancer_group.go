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

type LoadBalancerGroup struct {
	Id                          types.String                `tfsdk:"id"`
	AppDomain                   types.String                `tfsdk:"app_domain"`
	UserSummary                 types.String                `tfsdk:"user_summary"`
	Algorithm                   types.String                `tfsdk:"algorithm"`
	RetrieveInfo                types.Bool                  `tfsdk:"retrieve_info"`
	WlmRetrieval                types.String                `tfsdk:"wlm_retrieval"`
	WebSphereCellConfig         types.String                `tfsdk:"web_sphere_cell_config"`
	WlmGroup                    types.String                `tfsdk:"wlm_group"`
	WlmTransport                types.String                `tfsdk:"wlm_transport"`
	Damp                        types.Int64                 `tfsdk:"damp"`
	NeverReturnSickMember       types.Bool                  `tfsdk:"never_return_sick_member"`
	LbGroupMembers              types.List                  `tfsdk:"lb_group_members"`
	TryEveryServerBeforeFailing types.Bool                  `tfsdk:"try_every_server_before_failing"`
	LbGroupChecks               *DmLBGroupCheck             `tfsdk:"lb_group_checks"`
	MasqueradeMember            types.Bool                  `tfsdk:"masquerade_member"`
	ApplicationRouting          types.Bool                  `tfsdk:"application_routing"`
	LbGroupAffinityConf         *DmLBGroupAffinity          `tfsdk:"lb_group_affinity_conf"`
	MonitoredCookies            types.List                  `tfsdk:"monitored_cookies"`
	DependencyActions           []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var LoadBalancerGroupWLMRetrievalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retrieve_info",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var LoadBalancerGroupWebSphereCellConfigIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "retrieve_info",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "wlm_retrieval",
			AttrType:    "String",
			AttrDefault: "use-websphere",
			Value:       []string{"use-websphere"},
		},
	},
}

var LoadBalancerGroupWLMGroupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "retrieve_info",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var LoadBalancerGroupWLMTransportIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "retrieve_info",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "wlm_retrieval",
			AttrType:    "String",
			AttrDefault: "use-websphere",
			Value:       []string{"use-websphere"},
		},
	},
}

var LoadBalancerGroupApplicationRoutingIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "retrieve_info",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "wlm_retrieval",
			AttrType:    "String",
			AttrDefault: "use-websphere",
			Value:       []string{"use-websphere"},
		},
	},
}

var LoadBalancerGroupMonitoredCookiesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "affinity_wlm_override",
			AttrType:    "Bool",
			AttrDefault: "false",
			AttrPath:    "LBGroupAffinityConf",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "affinity_mode",
			AttrType:    "String",
			AttrDefault: "activeConditional",
			AttrPath:    "LBGroupAffinityConf",
			Value:       []string{"activeConditional"},
		},
	},
}

var LoadBalancerGroupObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"algorithm":                       types.StringType,
	"retrieve_info":                   types.BoolType,
	"wlm_retrieval":                   types.StringType,
	"web_sphere_cell_config":          types.StringType,
	"wlm_group":                       types.StringType,
	"wlm_transport":                   types.StringType,
	"damp":                            types.Int64Type,
	"never_return_sick_member":        types.BoolType,
	"lb_group_members":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmLBGroupMemberObjectType}},
	"try_every_server_before_failing": types.BoolType,
	"lb_group_checks":                 types.ObjectType{AttrTypes: DmLBGroupCheckObjectType},
	"masquerade_member":               types.BoolType,
	"application_routing":             types.BoolType,
	"lb_group_affinity_conf":          types.ObjectType{AttrTypes: DmLBGroupAffinityObjectType},
	"monitored_cookies":               types.ListType{ElemType: types.StringType},
	"dependency_actions":              actions.ActionsListType,
}

func (data LoadBalancerGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LoadBalancerGroup"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LoadBalancerGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Algorithm.IsNull() {
		return false
	}
	if !data.RetrieveInfo.IsNull() {
		return false
	}
	if !data.WlmRetrieval.IsNull() {
		return false
	}
	if !data.WebSphereCellConfig.IsNull() {
		return false
	}
	if !data.WlmGroup.IsNull() {
		return false
	}
	if !data.WlmTransport.IsNull() {
		return false
	}
	if !data.Damp.IsNull() {
		return false
	}
	if !data.NeverReturnSickMember.IsNull() {
		return false
	}
	if !data.LbGroupMembers.IsNull() {
		return false
	}
	if !data.TryEveryServerBeforeFailing.IsNull() {
		return false
	}
	if data.LbGroupChecks != nil {
		if !data.LbGroupChecks.IsNull() {
			return false
		}
	}
	if !data.MasqueradeMember.IsNull() {
		return false
	}
	if !data.ApplicationRouting.IsNull() {
		return false
	}
	if data.LbGroupAffinityConf != nil {
		if !data.LbGroupAffinityConf.IsNull() {
			return false
		}
	}
	if !data.MonitoredCookies.IsNull() {
		return false
	}
	return true
}

func (data LoadBalancerGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Algorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Algorithm`, data.Algorithm.ValueString())
	}
	if !data.RetrieveInfo.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetrieveInfo`, tfutils.StringFromBool(data.RetrieveInfo, ""))
	}
	if !data.WlmRetrieval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WLMRetrieval`, data.WlmRetrieval.ValueString())
	}
	if !data.WebSphereCellConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WebSphereCellConfig`, data.WebSphereCellConfig.ValueString())
	}
	if !data.WlmGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WLMGroup`, data.WlmGroup.ValueString())
	}
	if !data.WlmTransport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WLMTransport`, data.WlmTransport.ValueString())
	}
	if !data.Damp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Damp`, data.Damp.ValueInt64())
	}
	if !data.NeverReturnSickMember.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NeverReturnSickMember`, tfutils.StringFromBool(data.NeverReturnSickMember, ""))
	}
	if !data.LbGroupMembers.IsNull() {
		var dataValues []DmLBGroupMember
		data.LbGroupMembers.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`LBGroupMembers`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`LBGroupMembers`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`LBGroupMembers`, "[]")
	}
	if !data.TryEveryServerBeforeFailing.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TryEveryServerBeforeFailing`, tfutils.StringFromBool(data.TryEveryServerBeforeFailing, ""))
	}
	if data.LbGroupChecks != nil {
		if !data.LbGroupChecks.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`LBGroupChecks`, data.LbGroupChecks.ToBody(ctx, ""))
		}
	}
	if !data.MasqueradeMember.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MasqueradeMember`, tfutils.StringFromBool(data.MasqueradeMember, ""))
	}
	if !data.ApplicationRouting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ApplicationRouting`, tfutils.StringFromBool(data.ApplicationRouting, ""))
	}
	if data.LbGroupAffinityConf != nil {
		if !data.LbGroupAffinityConf.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`LBGroupAffinityConf`, data.LbGroupAffinityConf.ToBody(ctx, ""))
		}
	}
	if !data.MonitoredCookies.IsNull() {
		var dataValues []string
		data.MonitoredCookies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`MonitoredCookies`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`MonitoredCookies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`MonitoredCookies`, "[]")
	}
	return body
}

func (data *LoadBalancerGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Algorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Algorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Algorithm = types.StringValue("round-robin")
	}
	if value := res.Get(pathRoot + `RetrieveInfo`); value.Exists() {
		data.RetrieveInfo = tfutils.BoolFromString(value.String())
	} else {
		data.RetrieveInfo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WLMRetrieval`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WlmRetrieval = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WlmRetrieval = types.StringValue("use-websphere")
	}
	if value := res.Get(pathRoot + `WebSphereCellConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WebSphereCellConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WebSphereCellConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `WLMGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WlmGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WlmGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `WLMTransport`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WlmTransport = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WlmTransport = types.StringValue("http")
	}
	if value := res.Get(pathRoot + `Damp`); value.Exists() {
		data.Damp = types.Int64Value(value.Int())
	} else {
		data.Damp = types.Int64Value(120)
	}
	if value := res.Get(pathRoot + `NeverReturnSickMember`); value.Exists() {
		data.NeverReturnSickMember = tfutils.BoolFromString(value.String())
	} else {
		data.NeverReturnSickMember = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupMembers`); value.Exists() {
		l := []DmLBGroupMember{}
		if value := res.Get(`LBGroupMembers`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLBGroupMember{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LbGroupMembers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLBGroupMemberObjectType}, l)
		} else {
			data.LbGroupMembers = types.ListNull(types.ObjectType{AttrTypes: DmLBGroupMemberObjectType})
		}
	} else {
		data.LbGroupMembers = types.ListNull(types.ObjectType{AttrTypes: DmLBGroupMemberObjectType})
	}
	if value := res.Get(pathRoot + `TryEveryServerBeforeFailing`); value.Exists() {
		data.TryEveryServerBeforeFailing = tfutils.BoolFromString(value.String())
	} else {
		data.TryEveryServerBeforeFailing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupChecks`); value.Exists() {
		data.LbGroupChecks = &DmLBGroupCheck{}
		data.LbGroupChecks.FromBody(ctx, "", value)
	} else {
		data.LbGroupChecks = nil
	}
	if value := res.Get(pathRoot + `MasqueradeMember`); value.Exists() {
		data.MasqueradeMember = tfutils.BoolFromString(value.String())
	} else {
		data.MasqueradeMember = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplicationRouting`); value.Exists() {
		data.ApplicationRouting = tfutils.BoolFromString(value.String())
	} else {
		data.ApplicationRouting = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupAffinityConf`); value.Exists() {
		data.LbGroupAffinityConf = &DmLBGroupAffinity{}
		data.LbGroupAffinityConf.FromBody(ctx, "", value)
	} else {
		data.LbGroupAffinityConf = nil
	}
	if value := res.Get(pathRoot + `MonitoredCookies`); value.Exists() {
		data.MonitoredCookies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MonitoredCookies = types.ListNull(types.StringType)
	}
}

func (data *LoadBalancerGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Algorithm`); value.Exists() && !data.Algorithm.IsNull() {
		data.Algorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.Algorithm.ValueString() != "round-robin" {
		data.Algorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `RetrieveInfo`); value.Exists() && !data.RetrieveInfo.IsNull() {
		data.RetrieveInfo = tfutils.BoolFromString(value.String())
	} else if data.RetrieveInfo.ValueBool() {
		data.RetrieveInfo = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WLMRetrieval`); value.Exists() && !data.WlmRetrieval.IsNull() {
		data.WlmRetrieval = tfutils.ParseStringFromGJSON(value)
	} else if data.WlmRetrieval.ValueString() != "use-websphere" {
		data.WlmRetrieval = types.StringNull()
	}
	if value := res.Get(pathRoot + `WebSphereCellConfig`); value.Exists() && !data.WebSphereCellConfig.IsNull() {
		data.WebSphereCellConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WebSphereCellConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `WLMGroup`); value.Exists() && !data.WlmGroup.IsNull() {
		data.WlmGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WlmGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `WLMTransport`); value.Exists() && !data.WlmTransport.IsNull() {
		data.WlmTransport = tfutils.ParseStringFromGJSON(value)
	} else if data.WlmTransport.ValueString() != "http" {
		data.WlmTransport = types.StringNull()
	}
	if value := res.Get(pathRoot + `Damp`); value.Exists() && !data.Damp.IsNull() {
		data.Damp = types.Int64Value(value.Int())
	} else if data.Damp.ValueInt64() != 120 {
		data.Damp = types.Int64Null()
	}
	if value := res.Get(pathRoot + `NeverReturnSickMember`); value.Exists() && !data.NeverReturnSickMember.IsNull() {
		data.NeverReturnSickMember = tfutils.BoolFromString(value.String())
	} else if data.NeverReturnSickMember.ValueBool() {
		data.NeverReturnSickMember = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupMembers`); value.Exists() && !data.LbGroupMembers.IsNull() {
		l := []DmLBGroupMember{}
		e := []DmLBGroupMember{}
		data.LbGroupMembers.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmLBGroupMember{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LbGroupMembers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLBGroupMemberObjectType}, l)
		} else {
			data.LbGroupMembers = types.ListNull(types.ObjectType{AttrTypes: DmLBGroupMemberObjectType})
		}
	} else {
		data.LbGroupMembers = types.ListNull(types.ObjectType{AttrTypes: DmLBGroupMemberObjectType})
	}
	if value := res.Get(pathRoot + `TryEveryServerBeforeFailing`); value.Exists() && !data.TryEveryServerBeforeFailing.IsNull() {
		data.TryEveryServerBeforeFailing = tfutils.BoolFromString(value.String())
	} else if data.TryEveryServerBeforeFailing.ValueBool() {
		data.TryEveryServerBeforeFailing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupChecks`); value.Exists() {
		data.LbGroupChecks.UpdateFromBody(ctx, "", value)
	} else {
		data.LbGroupChecks = nil
	}
	if value := res.Get(pathRoot + `MasqueradeMember`); value.Exists() && !data.MasqueradeMember.IsNull() {
		data.MasqueradeMember = tfutils.BoolFromString(value.String())
	} else if data.MasqueradeMember.ValueBool() {
		data.MasqueradeMember = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ApplicationRouting`); value.Exists() && !data.ApplicationRouting.IsNull() {
		data.ApplicationRouting = tfutils.BoolFromString(value.String())
	} else if data.ApplicationRouting.ValueBool() {
		data.ApplicationRouting = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LBGroupAffinityConf`); value.Exists() {
		data.LbGroupAffinityConf.UpdateFromBody(ctx, "", value)
	} else {
		data.LbGroupAffinityConf = nil
	}
	if value := res.Get(pathRoot + `MonitoredCookies`); value.Exists() && !data.MonitoredCookies.IsNull() {
		data.MonitoredCookies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.MonitoredCookies = types.ListNull(types.StringType)
	}
}
