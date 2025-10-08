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

type AnalyticsEndpoint struct {
	Id                     types.String                `tfsdk:"id"`
	AppDomain              types.String                `tfsdk:"app_domain"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	AnalyticsServerUrl     types.String                `tfsdk:"analytics_server_url"`
	SslClient              types.String                `tfsdk:"ssl_client"`
	RequestTopic           types.String                `tfsdk:"request_topic"`
	MaxRecords             types.Int64                 `tfsdk:"max_records"`
	MaxRecordsMemoryKb     types.Int64                 `tfsdk:"max_records_memory_kb"`
	MaxDeliveryMemoryMb    types.Int64                 `tfsdk:"max_delivery_memory_mb"`
	Interval               types.Int64                 `tfsdk:"interval"`
	DeliveryConnections    types.Int64                 `tfsdk:"delivery_connections"`
	EnableJwt              types.Bool                  `tfsdk:"enable_jwt"`
	ManagementUrl          types.String                `tfsdk:"management_url"`
	ManagementUrlSslClient types.String                `tfsdk:"management_url_ssl_client"`
	ClientId               types.String                `tfsdk:"client_id"`
	ClientSecretAlias      types.String                `tfsdk:"client_secret_alias"`
	GrantType              types.String                `tfsdk:"grant_type"`
	Scope                  types.String                `tfsdk:"scope"`
	PersistentConnection   types.Bool                  `tfsdk:"persistent_connection"`
	Timeout                types.Int64                 `tfsdk:"timeout"`
	PersistentTimeout      types.Int64                 `tfsdk:"persistent_timeout"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AnalyticsEndpointSSLClientCondVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"https"},
}

var AnalyticsEndpointSSLClientIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"https"},
}

var AnalyticsEndpointRequestTopicCondVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"dpkafka"},
}

var AnalyticsEndpointRequestTopicIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"dpkafka"},
}

var AnalyticsEndpointDeliveryConnectionsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"dpkafka"},
}

var AnalyticsEndpointEnableJWTIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"http", "https"},
}

var AnalyticsEndpointManagementURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointManagementURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointManagementURL_SSLClientCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "management_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"https"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointManagementURL_SSLClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "management_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"dpkafka"},
		},
	},
}

var AnalyticsEndpointClientIDCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointClientIDIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var AnalyticsEndpointClientSecretAliasCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointClientSecretAliasIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var AnalyticsEndpointGrantTypeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointGrantTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var AnalyticsEndpointScopeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-url-protocol-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointScopeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "enable_jwt",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var AnalyticsEndpointPersistentConnectionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-url-protocol-not-in-list",
	Attribute:   "analytics_server_url",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"http", "https"},
}

var AnalyticsEndpointPersistentTimeoutIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "persistent_connection",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-url-protocol-not-in-list",
			Attribute:   "analytics_server_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"http", "https"},
		},
	},
}

var AnalyticsEndpointObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"app_domain":                types.StringType,
	"user_summary":              types.StringType,
	"analytics_server_url":      types.StringType,
	"ssl_client":                types.StringType,
	"request_topic":             types.StringType,
	"max_records":               types.Int64Type,
	"max_records_memory_kb":     types.Int64Type,
	"max_delivery_memory_mb":    types.Int64Type,
	"interval":                  types.Int64Type,
	"delivery_connections":      types.Int64Type,
	"enable_jwt":                types.BoolType,
	"management_url":            types.StringType,
	"management_url_ssl_client": types.StringType,
	"client_id":                 types.StringType,
	"client_secret_alias":       types.StringType,
	"grant_type":                types.StringType,
	"scope":                     types.StringType,
	"persistent_connection":     types.BoolType,
	"timeout":                   types.Int64Type,
	"persistent_timeout":        types.Int64Type,
	"dependency_actions":        actions.ActionsListType,
}

func (data AnalyticsEndpoint) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AnalyticsEndpoint"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AnalyticsEndpoint) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AnalyticsServerUrl.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.RequestTopic.IsNull() {
		return false
	}
	if !data.MaxRecords.IsNull() {
		return false
	}
	if !data.MaxRecordsMemoryKb.IsNull() {
		return false
	}
	if !data.MaxDeliveryMemoryMb.IsNull() {
		return false
	}
	if !data.Interval.IsNull() {
		return false
	}
	if !data.DeliveryConnections.IsNull() {
		return false
	}
	if !data.EnableJwt.IsNull() {
		return false
	}
	if !data.ManagementUrl.IsNull() {
		return false
	}
	if !data.ManagementUrlSslClient.IsNull() {
		return false
	}
	if !data.ClientId.IsNull() {
		return false
	}
	if !data.ClientSecretAlias.IsNull() {
		return false
	}
	if !data.GrantType.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.PersistentConnection.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.PersistentTimeout.IsNull() {
		return false
	}
	return true
}

func (data AnalyticsEndpoint) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AnalyticsServerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AnalyticsServerURL`, data.AnalyticsServerUrl.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.RequestTopic.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestTopic`, data.RequestTopic.ValueString())
	}
	if !data.MaxRecords.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRecords`, data.MaxRecords.ValueInt64())
	}
	if !data.MaxRecordsMemoryKb.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRecordsMemoryKB`, data.MaxRecordsMemoryKb.ValueInt64())
	}
	if !data.MaxDeliveryMemoryMb.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxDeliveryMemoryMB`, data.MaxDeliveryMemoryMb.ValueInt64())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Interval`, data.Interval.ValueInt64())
	}
	if !data.DeliveryConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeliveryConnections`, data.DeliveryConnections.ValueInt64())
	}
	if !data.EnableJwt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableJWT`, tfutils.StringFromBool(data.EnableJwt, ""))
	}
	if !data.ManagementUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ManagementURL`, data.ManagementUrl.ValueString())
	}
	if !data.ManagementUrlSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ManagementURL_SSLClient`, data.ManagementUrlSslClient.ValueString())
	}
	if !data.ClientId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientID`, data.ClientId.ValueString())
	}
	if !data.ClientSecretAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientSecretAlias`, data.ClientSecretAlias.ValueString())
	}
	if !data.GrantType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GrantType`, data.GrantType.ValueString())
	}
	if !data.Scope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Scope`, data.Scope.ValueString())
	}
	if !data.PersistentConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentConnection`, tfutils.StringFromBool(data.PersistentConnection, ""))
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.PersistentTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PersistentTimeout`, data.PersistentTimeout.ValueInt64())
	}
	return body
}

func (data *AnalyticsEndpoint) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AnalyticsServerURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AnalyticsServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AnalyticsServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopic`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRecords`); value.Exists() {
		data.MaxRecords = types.Int64Value(value.Int())
	} else {
		data.MaxRecords = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `MaxRecordsMemoryKB`); value.Exists() {
		data.MaxRecordsMemoryKb = types.Int64Value(value.Int())
	} else {
		data.MaxRecordsMemoryKb = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `MaxDeliveryMemoryMB`); value.Exists() {
		data.MaxDeliveryMemoryMb = types.Int64Value(value.Int())
	} else {
		data.MaxDeliveryMemoryMb = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Value(600)
	}
	if value := res.Get(pathRoot + `DeliveryConnections`); value.Exists() {
		data.DeliveryConnections = types.Int64Value(value.Int())
	} else {
		data.DeliveryConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `EnableJWT`); value.Exists() {
		data.EnableJwt = tfutils.BoolFromString(value.String())
	} else {
		data.EnableJwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ManagementURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ManagementUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ManagementUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ManagementURL_SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ManagementUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ManagementUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSecretAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientSecretAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecretAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `GrantType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GrantType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GrantType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentConnection`); value.Exists() {
		data.PersistentConnection = tfutils.BoolFromString(value.String())
	} else {
		data.PersistentConnection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(90)
	}
	if value := res.Get(pathRoot + `PersistentTimeout`); value.Exists() {
		data.PersistentTimeout = types.Int64Value(value.Int())
	} else {
		data.PersistentTimeout = types.Int64Value(60)
	}
}

func (data *AnalyticsEndpoint) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AnalyticsServerURL`); value.Exists() && !data.AnalyticsServerUrl.IsNull() {
		data.AnalyticsServerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AnalyticsServerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `RequestTopic`); value.Exists() && !data.RequestTopic.IsNull() {
		data.RequestTopic = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestTopic = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRecords`); value.Exists() && !data.MaxRecords.IsNull() {
		data.MaxRecords = types.Int64Value(value.Int())
	} else if data.MaxRecords.ValueInt64() != 1024 {
		data.MaxRecords = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxRecordsMemoryKB`); value.Exists() && !data.MaxRecordsMemoryKb.IsNull() {
		data.MaxRecordsMemoryKb = types.Int64Value(value.Int())
	} else if data.MaxRecordsMemoryKb.ValueInt64() != 512 {
		data.MaxRecordsMemoryKb = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxDeliveryMemoryMB`); value.Exists() && !data.MaxDeliveryMemoryMb.IsNull() {
		data.MaxDeliveryMemoryMb = types.Int64Value(value.Int())
	} else if data.MaxDeliveryMemoryMb.ValueInt64() != 512 {
		data.MaxDeliveryMemoryMb = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Interval`); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else if data.Interval.ValueInt64() != 600 {
		data.Interval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DeliveryConnections`); value.Exists() && !data.DeliveryConnections.IsNull() {
		data.DeliveryConnections = types.Int64Value(value.Int())
	} else if data.DeliveryConnections.ValueInt64() != 1 {
		data.DeliveryConnections = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EnableJWT`); value.Exists() && !data.EnableJwt.IsNull() {
		data.EnableJwt = tfutils.BoolFromString(value.String())
	} else if data.EnableJwt.ValueBool() {
		data.EnableJwt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ManagementURL`); value.Exists() && !data.ManagementUrl.IsNull() {
		data.ManagementUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ManagementUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ManagementURL_SSLClient`); value.Exists() && !data.ManagementUrlSslClient.IsNull() {
		data.ManagementUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ManagementUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientID`); value.Exists() && !data.ClientId.IsNull() {
		data.ClientId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientId = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientSecretAlias`); value.Exists() && !data.ClientSecretAlias.IsNull() {
		data.ClientSecretAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecretAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `GrantType`); value.Exists() && !data.GrantType.IsNull() {
		data.GrantType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GrantType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && !data.Scope.IsNull() {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `PersistentConnection`); value.Exists() && !data.PersistentConnection.IsNull() {
		data.PersistentConnection = tfutils.BoolFromString(value.String())
	} else if !data.PersistentConnection.ValueBool() {
		data.PersistentConnection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 90 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PersistentTimeout`); value.Exists() && !data.PersistentTimeout.IsNull() {
		data.PersistentTimeout = types.Int64Value(value.Int())
	} else if data.PersistentTimeout.ValueInt64() != 60 {
		data.PersistentTimeout = types.Int64Null()
	}
}
