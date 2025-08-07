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

type KafkaCluster struct {
	Id                 types.String      `tfsdk:"id"`
	AppDomain          types.String      `tfsdk:"app_domain"`
	UserSummary        types.String      `tfsdk:"user_summary"`
	Protocol           types.String      `tfsdk:"protocol"`
	Endpoint           types.List        `tfsdk:"endpoint"`
	SaslMechanism      types.String      `tfsdk:"sasl_mechanism"`
	UserName           types.String      `tfsdk:"user_name"`
	PasswordAlias      types.String      `tfsdk:"password_alias"`
	Autocommit         types.Bool        `tfsdk:"autocommit"`
	SslClient          types.String      `tfsdk:"ssl_client"`
	MemoryThreshold    types.Int64       `tfsdk:"memory_threshold"`
	MaximumMessageSize types.Int64       `tfsdk:"maximum_message_size"`
	AutoRetry          types.Bool        `tfsdk:"auto_retry"`
	RetryInterval      types.Int64       `tfsdk:"retry_interval"`
	Property           types.List        `tfsdk:"property"`
	DependencyActions  []*actions.Action `tfsdk:"dependency_actions"`
}

var KafkaClusterObjectType = map[string]attr.Type{
	"id":                   types.StringType,
	"app_domain":           types.StringType,
	"user_summary":         types.StringType,
	"protocol":             types.StringType,
	"endpoint":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmKafkaEndpointObjectType}},
	"sasl_mechanism":       types.StringType,
	"user_name":            types.StringType,
	"password_alias":       types.StringType,
	"autocommit":           types.BoolType,
	"ssl_client":           types.StringType,
	"memory_threshold":     types.Int64Type,
	"maximum_message_size": types.Int64Type,
	"auto_retry":           types.BoolType,
	"retry_interval":       types.Int64Type,
	"property":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmKafkaPropertyObjectType}},
	"dependency_actions":   actions.ActionsListType,
}

func (data KafkaCluster) GetPath() string {
	rest_path := "/mgmt/config/{domain}/KafkaCluster"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data KafkaCluster) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Protocol.IsNull() {
		return false
	}
	if !data.Endpoint.IsNull() {
		return false
	}
	if !data.SaslMechanism.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.Autocommit.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.MemoryThreshold.IsNull() {
		return false
	}
	if !data.MaximumMessageSize.IsNull() {
		return false
	}
	if !data.AutoRetry.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.Property.IsNull() {
		return false
	}
	return true
}

func (data KafkaCluster) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Protocol`, data.Protocol.ValueString())
	}
	if !data.Endpoint.IsNull() {
		var values []DmKafkaEndpoint
		data.Endpoint.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Endpoint`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SaslMechanism.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SASLMechanism`, data.SaslMechanism.ValueString())
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.Autocommit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Autocommit`, tfutils.StringFromBool(data.Autocommit, ""))
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.MemoryThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MemoryThreshold`, data.MemoryThreshold.ValueInt64())
	}
	if !data.MaximumMessageSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaximumMessageSize`, data.MaximumMessageSize.ValueInt64())
	}
	if !data.AutoRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AutoRetry`, tfutils.StringFromBool(data.AutoRetry, ""))
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.Property.IsNull() {
		var values []DmKafkaProperty
		data.Property.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Property`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *KafkaCluster) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Protocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Protocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Protocol = types.StringValue("plaintext")
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() {
		l := []DmKafkaEndpoint{}
		if value := res.Get(`Endpoint`); value.Exists() {
			for _, v := range value.Array() {
				item := DmKafkaEndpoint{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Endpoint, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmKafkaEndpointObjectType}, l)
		} else {
			data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmKafkaEndpointObjectType})
		}
	} else {
		data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmKafkaEndpointObjectType})
	}
	if value := res.Get(pathRoot + `SASLMechanism`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SaslMechanism = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SaslMechanism = types.StringValue("plain")
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Autocommit`); value.Exists() {
		data.Autocommit = tfutils.BoolFromString(value.String())
	} else {
		data.Autocommit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `MemoryThreshold`); value.Exists() {
		data.MemoryThreshold = types.Int64Value(value.Int())
	} else {
		data.MemoryThreshold = types.Int64Value(268435456)
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else {
		data.MaximumMessageSize = types.Int64Value(1048576)
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() {
		l := []DmKafkaProperty{}
		if value := res.Get(`Property`); value.Exists() {
			for _, v := range value.Array() {
				item := DmKafkaProperty{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Property, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmKafkaPropertyObjectType}, l)
		} else {
			data.Property = types.ListNull(types.ObjectType{AttrTypes: DmKafkaPropertyObjectType})
		}
	} else {
		data.Property = types.ListNull(types.ObjectType{AttrTypes: DmKafkaPropertyObjectType})
	}
}

func (data *KafkaCluster) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Protocol`); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = tfutils.ParseStringFromGJSON(value)
	} else if data.Protocol.ValueString() != "plaintext" {
		data.Protocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `Endpoint`); value.Exists() && !data.Endpoint.IsNull() {
		l := []DmKafkaEndpoint{}
		for _, v := range value.Array() {
			item := DmKafkaEndpoint{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Endpoint, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmKafkaEndpointObjectType}, l)
		} else {
			data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmKafkaEndpointObjectType})
		}
	} else {
		data.Endpoint = types.ListNull(types.ObjectType{AttrTypes: DmKafkaEndpointObjectType})
	}
	if value := res.Get(pathRoot + `SASLMechanism`); value.Exists() && !data.SaslMechanism.IsNull() {
		data.SaslMechanism = tfutils.ParseStringFromGJSON(value)
	} else if data.SaslMechanism.ValueString() != "plain" {
		data.SaslMechanism = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `Autocommit`); value.Exists() && !data.Autocommit.IsNull() {
		data.Autocommit = tfutils.BoolFromString(value.String())
	} else if !data.Autocommit.ValueBool() {
		data.Autocommit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `MemoryThreshold`); value.Exists() && !data.MemoryThreshold.IsNull() {
		data.MemoryThreshold = types.Int64Value(value.Int())
	} else if data.MemoryThreshold.ValueInt64() != 268435456 {
		data.MemoryThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaximumMessageSize`); value.Exists() && !data.MaximumMessageSize.IsNull() {
		data.MaximumMessageSize = types.Int64Value(value.Int())
	} else if data.MaximumMessageSize.ValueInt64() != 1048576 {
		data.MaximumMessageSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AutoRetry`); value.Exists() && !data.AutoRetry.IsNull() {
		data.AutoRetry = tfutils.BoolFromString(value.String())
	} else if !data.AutoRetry.ValueBool() {
		data.AutoRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 10 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Property`); value.Exists() && !data.Property.IsNull() {
		l := []DmKafkaProperty{}
		for _, v := range value.Array() {
			item := DmKafkaProperty{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Property, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmKafkaPropertyObjectType}, l)
		} else {
			data.Property = types.ListNull(types.ObjectType{AttrTypes: DmKafkaPropertyObjectType})
		}
	} else {
		data.Property = types.ListNull(types.ObjectType{AttrTypes: DmKafkaPropertyObjectType})
	}
}
