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
	"path"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type APIConnectGatewayService struct {
	AppDomain             types.String                `tfsdk:"app_domain"`
	Enabled               types.Bool                  `tfsdk:"enabled"`
	UserSummary           types.String                `tfsdk:"user_summary"`
	LocalAddress          types.String                `tfsdk:"local_address"`
	LocalPort             types.Int64                 `tfsdk:"local_port"`
	SslServer             types.String                `tfsdk:"ssl_server"`
	ApiGatewayAddress     types.String                `tfsdk:"api_gateway_address"`
	ApiGatewayPort        types.Int64                 `tfsdk:"api_gateway_port"`
	GatewayPeering        types.String                `tfsdk:"gateway_peering"`
	GatewayPeeringManager types.String                `tfsdk:"gateway_peering_manager"`
	V5compatibilityMode   types.Bool                  `tfsdk:"v5compatibility_mode"`
	UserDefinedPolicies   types.List                  `tfsdk:"user_defined_policies"`
	V5cSlmMode            types.String                `tfsdk:"v5c_slm_mode"`
	IpMulticast           types.String                `tfsdk:"ip_multicast"`
	IpUnicast             types.String                `tfsdk:"ip_unicast"`
	JwtValidationMode     types.String                `tfsdk:"jwt_validation_mode"`
	JwtUrl                types.String                `tfsdk:"jwt_url"`
	ProxyPolicy           *DmAPICGSProxyPolicy        `tfsdk:"proxy_policy"`
	DependencyActions     []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APIConnectGatewayServiceGatewayPeeringCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var APIConnectGatewayServiceGatewayPeeringManagerCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var APIConnectGatewayServiceGatewayPeeringManagerIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var APIConnectGatewayServiceUserDefinedPoliciesIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var APIConnectGatewayServiceV5CSlmModeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}

var APIConnectGatewayServiceV5CSlmModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "v5compatibility_mode",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"false"},
}

var APIConnectGatewayServiceIPMulticastCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5c_slm_mode",
			AttrType:    "String",
			AttrDefault: "autounicast",
			Value:       []string{"multicast"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5compatibility_mode",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var APIConnectGatewayServiceIPMulticastIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "v5c_slm_mode",
			AttrType:    "String",
			AttrDefault: "autounicast",
			Value:       []string{"multicast"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5compatibility_mode",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var APIConnectGatewayServiceIPUnicastCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5c_slm_mode",
			AttrType:    "String",
			AttrDefault: "autounicast",
			Value:       []string{"unicast"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5compatibility_mode",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"true"},
		},
	},
}

var APIConnectGatewayServiceIPUnicastIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "v5c_slm_mode",
			AttrType:    "String",
			AttrDefault: "autounicast",
			Value:       []string{"unicast"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "v5compatibility_mode",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
	},
}

var APIConnectGatewayServiceObjectType = map[string]attr.Type{
	"app_domain":              types.StringType,
	"enabled":                 types.BoolType,
	"user_summary":            types.StringType,
	"local_address":           types.StringType,
	"local_port":              types.Int64Type,
	"ssl_server":              types.StringType,
	"api_gateway_address":     types.StringType,
	"api_gateway_port":        types.Int64Type,
	"gateway_peering":         types.StringType,
	"gateway_peering_manager": types.StringType,
	"v5compatibility_mode":    types.BoolType,
	"user_defined_policies":   types.ListType{ElemType: types.StringType},
	"v5c_slm_mode":            types.StringType,
	"ip_multicast":            types.StringType,
	"ip_unicast":              types.StringType,
	"jwt_validation_mode":     types.StringType,
	"jwt_url":                 types.StringType,
	"proxy_policy":            types.ObjectType{AttrTypes: DmAPICGSProxyPolicyObjectType},
	"dependency_actions":      actions.ActionsListType,
}

func (data APIConnectGatewayService) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIConnectGatewayService/default"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIConnectGatewayService) IsNull() bool {
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.LocalPort.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	if !data.ApiGatewayAddress.IsNull() {
		return false
	}
	if !data.ApiGatewayPort.IsNull() {
		return false
	}
	if !data.GatewayPeering.IsNull() {
		return false
	}
	if !data.GatewayPeeringManager.IsNull() {
		return false
	}
	if !data.V5compatibilityMode.IsNull() {
		return false
	}
	if !data.UserDefinedPolicies.IsNull() {
		return false
	}
	if !data.V5cSlmMode.IsNull() {
		return false
	}
	if !data.IpMulticast.IsNull() {
		return false
	}
	if !data.IpUnicast.IsNull() {
		return false
	}
	if !data.JwtValidationMode.IsNull() {
		return false
	}
	if !data.JwtUrl.IsNull() {
		return false
	}
	if data.ProxyPolicy != nil {
		if !data.ProxyPolicy.IsNull() {
			return false
		}
	}
	return true
}
func (data *APIConnectGatewayService) ToDefault() {
	data.Enabled = types.BoolValue(false)
	data.UserSummary = types.StringNull()
	data.LocalAddress = types.StringValue("0.0.0.0")
	data.LocalPort = types.Int64Value(3000)
	data.SslServer = types.StringNull()
	data.ApiGatewayAddress = types.StringValue("0.0.0.0")
	data.ApiGatewayPort = types.Int64Value(9443)
	data.GatewayPeering = types.StringNull()
	data.GatewayPeeringManager = types.StringValue("default")
	data.V5compatibilityMode = types.BoolValue(true)
	data.UserDefinedPolicies = types.ListNull(types.StringType)
	data.V5cSlmMode = types.StringValue("autounicast")
	data.IpMulticast = types.StringNull()
	data.IpUnicast = types.StringNull()
	data.JwtValidationMode = types.StringValue("request")
	data.JwtUrl = types.StringNull()
	data.ProxyPolicy = &DmAPICGSProxyPolicy{}
	data.ProxyPolicy.ToDefault()
}

func (data APIConnectGatewayService) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "APIConnectGatewayService.name", path.Base("/mgmt/config/{domain}/APIConnectGatewayService/default"))

	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.LocalPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalPort`, data.LocalPort.ValueInt64())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	if !data.ApiGatewayAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIGatewayAddress`, data.ApiGatewayAddress.ValueString())
	}
	if !data.ApiGatewayPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIGatewayPort`, data.ApiGatewayPort.ValueInt64())
	}
	if !data.GatewayPeering.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayPeering`, data.GatewayPeering.ValueString())
	}
	if !data.GatewayPeeringManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayPeeringManager`, data.GatewayPeeringManager.ValueString())
	}
	if !data.V5compatibilityMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`V5CompatibilityMode`, tfutils.StringFromBool(data.V5compatibilityMode, ""))
	}
	if !data.UserDefinedPolicies.IsNull() {
		var dataValues []string
		data.UserDefinedPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`UserDefinedPolicies`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`UserDefinedPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`UserDefinedPolicies`, "[]")
	}
	if !data.V5cSlmMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`V5CSlmMode`, data.V5cSlmMode.ValueString())
	}
	if !data.IpMulticast.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPMulticast`, data.IpMulticast.ValueString())
	}
	if !data.IpUnicast.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPUnicast`, data.IpUnicast.ValueString())
	}
	if !data.JwtValidationMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTValidationMode`, data.JwtValidationMode.ValueString())
	}
	if !data.JwtUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTURL`, data.JwtUrl.ValueString())
	}
	if data.ProxyPolicy != nil {
		if !data.ProxyPolicy.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicy`, data.ProxyPolicy.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *APIConnectGatewayService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() {
		data.LocalPort = types.Int64Value(value.Int())
	} else {
		data.LocalPort = types.Int64Value(3000)
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIGatewayAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiGatewayAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiGatewayAddress = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `APIGatewayPort`); value.Exists() {
		data.ApiGatewayPort = types.Int64Value(value.Int())
	} else {
		data.ApiGatewayPort = types.Int64Value(9443)
	}
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayPeeringManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayPeeringManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeeringManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `V5CompatibilityMode`); value.Exists() {
		data.V5compatibilityMode = tfutils.BoolFromString(value.String())
	} else {
		data.V5compatibilityMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserDefinedPolicies`); value.Exists() {
		data.UserDefinedPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.UserDefinedPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `V5CSlmMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.V5cSlmMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.V5cSlmMode = types.StringValue("autounicast")
	}
	if value := res.Get(pathRoot + `IPMulticast`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpMulticast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpMulticast = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPUnicast`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpUnicast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpUnicast = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTValidationMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtValidationMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtValidationMode = types.StringValue("request")
	}
	if value := res.Get(pathRoot + `JWTURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProxyPolicy`); value.Exists() {
		data.ProxyPolicy = &DmAPICGSProxyPolicy{}
		data.ProxyPolicy.FromBody(ctx, "", value)
	} else {
		data.ProxyPolicy = nil
	}
}

func (data *APIConnectGatewayService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAddress.ValueString() != "0.0.0.0" {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalPort`); value.Exists() && !data.LocalPort.IsNull() {
		data.LocalPort = types.Int64Value(value.Int())
	} else if data.LocalPort.ValueInt64() != 3000 {
		data.LocalPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIGatewayAddress`); value.Exists() && !data.ApiGatewayAddress.IsNull() {
		data.ApiGatewayAddress = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiGatewayAddress.ValueString() != "0.0.0.0" {
		data.ApiGatewayAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `APIGatewayPort`); value.Exists() && !data.ApiGatewayPort.IsNull() {
		data.ApiGatewayPort = types.Int64Value(value.Int())
	} else if data.ApiGatewayPort.ValueInt64() != 9443 {
		data.ApiGatewayPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `GatewayPeering`); value.Exists() && !data.GatewayPeering.IsNull() {
		data.GatewayPeering = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayPeering = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayPeeringManager`); value.Exists() && !data.GatewayPeeringManager.IsNull() {
		data.GatewayPeeringManager = tfutils.ParseStringFromGJSON(value)
	} else if data.GatewayPeeringManager.ValueString() != "default" {
		data.GatewayPeeringManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `V5CompatibilityMode`); value.Exists() && !data.V5compatibilityMode.IsNull() {
		data.V5compatibilityMode = tfutils.BoolFromString(value.String())
	} else if !data.V5compatibilityMode.ValueBool() {
		data.V5compatibilityMode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserDefinedPolicies`); value.Exists() && !data.UserDefinedPolicies.IsNull() {
		data.UserDefinedPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.UserDefinedPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `V5CSlmMode`); value.Exists() && !data.V5cSlmMode.IsNull() {
		data.V5cSlmMode = tfutils.ParseStringFromGJSON(value)
	} else if data.V5cSlmMode.ValueString() != "autounicast" {
		data.V5cSlmMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPMulticast`); value.Exists() && !data.IpMulticast.IsNull() {
		data.IpMulticast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpMulticast = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPUnicast`); value.Exists() && !data.IpUnicast.IsNull() {
		data.IpUnicast = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpUnicast = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTValidationMode`); value.Exists() && !data.JwtValidationMode.IsNull() {
		data.JwtValidationMode = tfutils.ParseStringFromGJSON(value)
	} else if data.JwtValidationMode.ValueString() != "request" {
		data.JwtValidationMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTURL`); value.Exists() && !data.JwtUrl.IsNull() {
		data.JwtUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProxyPolicy`); value.Exists() {
		data.ProxyPolicy.UpdateFromBody(ctx, "", value)
	} else {
		data.ProxyPolicy = nil
	}
}
