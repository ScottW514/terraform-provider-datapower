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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DNSNameService struct {
	Enabled              types.Bool                  `tfsdk:"enabled"`
	UserSummary          types.String                `tfsdk:"user_summary"`
	SearchDomains        types.List                  `tfsdk:"search_domains"`
	NameServers          types.List                  `tfsdk:"name_servers"`
	StaticHosts          types.List                  `tfsdk:"static_hosts"`
	IpPreference         types.String                `tfsdk:"ip_preference"`
	ForceIpPreference    types.Bool                  `tfsdk:"force_ip_preference"`
	LoadBalanceAlgorithm types.String                `tfsdk:"load_balance_algorithm"`
	MaxRetries           types.Int64                 `tfsdk:"max_retries"`
	Timeout              types.Int64                 `tfsdk:"timeout"`
	DependencyActions    []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var DNSNameServiceObjectType = map[string]attr.Type{
	"enabled":                types.BoolType,
	"user_summary":           types.StringType,
	"search_domains":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmSearchDomainObjectType}},
	"name_servers":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmNameServerObjectType}},
	"static_hosts":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmStaticHostObjectType}},
	"ip_preference":          types.StringType,
	"force_ip_preference":    types.BoolType,
	"load_balance_algorithm": types.StringType,
	"max_retries":            types.Int64Type,
	"timeout":                types.Int64Type,
	"dependency_actions":     actions.ActionsListType,
}

func (data DNSNameService) GetPath() string {
	rest_path := "/mgmt/config/default/DNSNameService/Main-Name-Service"
	return rest_path
}

func (data DNSNameService) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.SearchDomains.IsNull() {
		return false
	}
	if !data.NameServers.IsNull() {
		return false
	}
	if !data.StaticHosts.IsNull() {
		return false
	}
	if !data.IpPreference.IsNull() {
		return false
	}
	if !data.ForceIpPreference.IsNull() {
		return false
	}
	if !data.LoadBalanceAlgorithm.IsNull() {
		return false
	}
	if !data.MaxRetries.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	return true
}

func (data DNSNameService) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "DNSNameService.name", path.Base("/mgmt/config/default/DNSNameService/Main-Name-Service"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.SearchDomains.IsNull() {
		var values []DmSearchDomain
		data.SearchDomains.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`SearchDomains`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.NameServers.IsNull() {
		var values []DmNameServer
		data.NameServers.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`NameServers`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.StaticHosts.IsNull() {
		var values []DmStaticHost
		data.StaticHosts.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`StaticHosts`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.IpPreference.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPPreference`, data.IpPreference.ValueString())
	}
	if !data.ForceIpPreference.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ForceIPPreference`, tfutils.StringFromBool(data.ForceIpPreference, ""))
	}
	if !data.LoadBalanceAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoadBalanceAlgorithm`, data.LoadBalanceAlgorithm.ValueString())
	}
	if !data.MaxRetries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRetries`, data.MaxRetries.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	return body
}

func (data *DNSNameService) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `SearchDomains`); value.Exists() {
		l := []DmSearchDomain{}
		if value := res.Get(`SearchDomains`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSearchDomain{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SearchDomains, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSearchDomainObjectType}, l)
		} else {
			data.SearchDomains = types.ListNull(types.ObjectType{AttrTypes: DmSearchDomainObjectType})
		}
	} else {
		data.SearchDomains = types.ListNull(types.ObjectType{AttrTypes: DmSearchDomainObjectType})
	}
	if value := res.Get(pathRoot + `NameServers`); value.Exists() {
		l := []DmNameServer{}
		if value := res.Get(`NameServers`); value.Exists() {
			for _, v := range value.Array() {
				item := DmNameServer{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.NameServers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNameServerObjectType}, l)
		} else {
			data.NameServers = types.ListNull(types.ObjectType{AttrTypes: DmNameServerObjectType})
		}
	} else {
		data.NameServers = types.ListNull(types.ObjectType{AttrTypes: DmNameServerObjectType})
	}
	if value := res.Get(pathRoot + `StaticHosts`); value.Exists() {
		l := []DmStaticHost{}
		if value := res.Get(`StaticHosts`); value.Exists() {
			for _, v := range value.Array() {
				item := DmStaticHost{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.StaticHosts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStaticHostObjectType}, l)
		} else {
			data.StaticHosts = types.ListNull(types.ObjectType{AttrTypes: DmStaticHostObjectType})
		}
	} else {
		data.StaticHosts = types.ListNull(types.ObjectType{AttrTypes: DmStaticHostObjectType})
	}
	if value := res.Get(pathRoot + `IPPreference`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpPreference = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpPreference = types.StringNull()
	}
	if value := res.Get(pathRoot + `ForceIPPreference`); value.Exists() {
		data.ForceIpPreference = tfutils.BoolFromString(value.String())
	} else {
		data.ForceIpPreference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoadBalanceAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LoadBalanceAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoadBalanceAlgorithm = types.StringValue("first-alive")
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else {
		data.MaxRetries = types.Int64Value(2)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(5)
	}
}

func (data *DNSNameService) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `SearchDomains`); value.Exists() && !data.SearchDomains.IsNull() {
		l := []DmSearchDomain{}
		for _, v := range value.Array() {
			item := DmSearchDomain{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SearchDomains, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSearchDomainObjectType}, l)
		} else {
			data.SearchDomains = types.ListNull(types.ObjectType{AttrTypes: DmSearchDomainObjectType})
		}
	} else {
		data.SearchDomains = types.ListNull(types.ObjectType{AttrTypes: DmSearchDomainObjectType})
	}
	if value := res.Get(pathRoot + `NameServers`); value.Exists() && !data.NameServers.IsNull() {
		l := []DmNameServer{}
		for _, v := range value.Array() {
			item := DmNameServer{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.NameServers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNameServerObjectType}, l)
		} else {
			data.NameServers = types.ListNull(types.ObjectType{AttrTypes: DmNameServerObjectType})
		}
	} else {
		data.NameServers = types.ListNull(types.ObjectType{AttrTypes: DmNameServerObjectType})
	}
	if value := res.Get(pathRoot + `StaticHosts`); value.Exists() && !data.StaticHosts.IsNull() {
		l := []DmStaticHost{}
		for _, v := range value.Array() {
			item := DmStaticHost{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.StaticHosts, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStaticHostObjectType}, l)
		} else {
			data.StaticHosts = types.ListNull(types.ObjectType{AttrTypes: DmStaticHostObjectType})
		}
	} else {
		data.StaticHosts = types.ListNull(types.ObjectType{AttrTypes: DmStaticHostObjectType})
	}
	if value := res.Get(pathRoot + `IPPreference`); value.Exists() && !data.IpPreference.IsNull() {
		data.IpPreference = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpPreference = types.StringNull()
	}
	if value := res.Get(pathRoot + `ForceIPPreference`); value.Exists() && !data.ForceIpPreference.IsNull() {
		data.ForceIpPreference = tfutils.BoolFromString(value.String())
	} else if data.ForceIpPreference.ValueBool() {
		data.ForceIpPreference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LoadBalanceAlgorithm`); value.Exists() && !data.LoadBalanceAlgorithm.IsNull() {
		data.LoadBalanceAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.LoadBalanceAlgorithm.ValueString() != "first-alive" {
		data.LoadBalanceAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRetries`); value.Exists() && !data.MaxRetries.IsNull() {
		data.MaxRetries = types.Int64Value(value.Int())
	} else if data.MaxRetries.ValueInt64() != 2 {
		data.MaxRetries = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 5 {
		data.Timeout = types.Int64Null()
	}
}
