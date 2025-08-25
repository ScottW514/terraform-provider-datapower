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

type GatewayPeeringGroup struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Mode                types.String                `tfsdk:"mode"`
	PeerNodes           types.List                  `tfsdk:"peer_nodes"`
	ClusterPrimaryCount types.String                `tfsdk:"cluster_primary_count"`
	ClusterNodes        types.List                  `tfsdk:"cluster_nodes"`
	ClusterAutoConfig   types.Bool                  `tfsdk:"cluster_auto_config"`
	EnableSsl           types.Bool                  `tfsdk:"enable_ssl"`
	Idcred              types.String                `tfsdk:"idcred"`
	Valcred             types.String                `tfsdk:"valcred"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var GatewayPeeringGroupPeerNodesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "peer",
	Value:       []string{"peer"},
}
var GatewayPeeringGroupClusterPrimaryCountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "peer",
	Value:       []string{"cluster"},
}
var GatewayPeeringGroupClusterNodesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "peer",
	Value:       []string{"cluster"},
}
var GatewayPeeringGroupClusterAutoConfigCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "mode",
	AttrType:    "String",
	AttrDefault: "peer",
	Value:       []string{"cluster"},
}
var GatewayPeeringGroupIdcredCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "enable_ssl",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var GatewayPeeringGroupObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"mode":                  types.StringType,
	"peer_nodes":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType}},
	"cluster_primary_count": types.StringType,
	"cluster_nodes":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType}},
	"cluster_auto_config":   types.BoolType,
	"enable_ssl":            types.BoolType,
	"idcred":                types.StringType,
	"valcred":               types.StringType,
	"dependency_actions":    actions.ActionsListType,
}

func (data GatewayPeeringGroup) GetPath() string {
	rest_path := "/mgmt/config/{domain}/GatewayPeeringGroup"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data GatewayPeeringGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.PeerNodes.IsNull() {
		return false
	}
	if !data.ClusterPrimaryCount.IsNull() {
		return false
	}
	if !data.ClusterNodes.IsNull() {
		return false
	}
	if !data.ClusterAutoConfig.IsNull() {
		return false
	}
	if !data.EnableSsl.IsNull() {
		return false
	}
	if !data.Idcred.IsNull() {
		return false
	}
	if !data.Valcred.IsNull() {
		return false
	}
	return true
}

func (data GatewayPeeringGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Mode`, data.Mode.ValueString())
	}
	if !data.PeerNodes.IsNull() {
		var dataValues []DmGatewayPeeringGroupPeerNode
		data.PeerNodes.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`PeerNodes`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ClusterPrimaryCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClusterPrimaryCount`, data.ClusterPrimaryCount.ValueString())
	}
	if !data.ClusterNodes.IsNull() {
		var dataValues []DmGatewayPeeringGroupClusterNode
		data.ClusterNodes.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ClusterNodes`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ClusterAutoConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClusterAutoConfig`, tfutils.StringFromBool(data.ClusterAutoConfig, ""))
	}
	if !data.EnableSsl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableSSL`, tfutils.StringFromBool(data.EnableSsl, ""))
	}
	if !data.Idcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Idcred`, data.Idcred.ValueString())
	}
	if !data.Valcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Valcred`, data.Valcred.ValueString())
	}
	return body
}

func (data *GatewayPeeringGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Mode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Mode = types.StringValue("peer")
	}
	if value := res.Get(pathRoot + `PeerNodes`); value.Exists() {
		l := []DmGatewayPeeringGroupPeerNode{}
		if value := res.Get(`PeerNodes`); value.Exists() {
			for _, v := range value.Array() {
				item := DmGatewayPeeringGroupPeerNode{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PeerNodes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType}, l)
		} else {
			data.PeerNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType})
		}
	} else {
		data.PeerNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType})
	}
	if value := res.Get(pathRoot + `ClusterPrimaryCount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClusterPrimaryCount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClusterPrimaryCount = types.StringValue("3")
	}
	if value := res.Get(pathRoot + `ClusterNodes`); value.Exists() {
		l := []DmGatewayPeeringGroupClusterNode{}
		if value := res.Get(`ClusterNodes`); value.Exists() {
			for _, v := range value.Array() {
				item := DmGatewayPeeringGroupClusterNode{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ClusterNodes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType}, l)
		} else {
			data.ClusterNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType})
		}
	} else {
		data.ClusterNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType})
	}
	if value := res.Get(pathRoot + `ClusterAutoConfig`); value.Exists() {
		data.ClusterAutoConfig = tfutils.BoolFromString(value.String())
	} else {
		data.ClusterAutoConfig = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableSSL`); value.Exists() {
		data.EnableSsl = tfutils.BoolFromString(value.String())
	} else {
		data.EnableSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Idcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Idcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Idcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Valcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Valcred = types.StringNull()
	}
}

func (data *GatewayPeeringGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Mode`); value.Exists() && !data.Mode.IsNull() {
		data.Mode = tfutils.ParseStringFromGJSON(value)
	} else if data.Mode.ValueString() != "peer" {
		data.Mode = types.StringNull()
	}
	if value := res.Get(pathRoot + `PeerNodes`); value.Exists() && !data.PeerNodes.IsNull() {
		l := []DmGatewayPeeringGroupPeerNode{}
		for _, v := range value.Array() {
			item := DmGatewayPeeringGroupPeerNode{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.PeerNodes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType}, l)
		} else {
			data.PeerNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType})
		}
	} else {
		data.PeerNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupPeerNodeObjectType})
	}
	if value := res.Get(pathRoot + `ClusterPrimaryCount`); value.Exists() && !data.ClusterPrimaryCount.IsNull() {
		data.ClusterPrimaryCount = tfutils.ParseStringFromGJSON(value)
	} else if data.ClusterPrimaryCount.ValueString() != "3" {
		data.ClusterPrimaryCount = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClusterNodes`); value.Exists() && !data.ClusterNodes.IsNull() {
		l := []DmGatewayPeeringGroupClusterNode{}
		for _, v := range value.Array() {
			item := DmGatewayPeeringGroupClusterNode{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ClusterNodes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType}, l)
		} else {
			data.ClusterNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType})
		}
	} else {
		data.ClusterNodes = types.ListNull(types.ObjectType{AttrTypes: DmGatewayPeeringGroupClusterNodeObjectType})
	}
	if value := res.Get(pathRoot + `ClusterAutoConfig`); value.Exists() && !data.ClusterAutoConfig.IsNull() {
		data.ClusterAutoConfig = tfutils.BoolFromString(value.String())
	} else if !data.ClusterAutoConfig.ValueBool() {
		data.ClusterAutoConfig = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnableSSL`); value.Exists() && !data.EnableSsl.IsNull() {
		data.EnableSsl = tfutils.BoolFromString(value.String())
	} else if !data.EnableSsl.ValueBool() {
		data.EnableSsl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Idcred`); value.Exists() && !data.Idcred.IsNull() {
		data.Idcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Idcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Valcred`); value.Exists() && !data.Valcred.IsNull() {
		data.Valcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Valcred = types.StringNull()
	}
}
