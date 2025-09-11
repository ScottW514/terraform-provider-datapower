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

type OpenTelemetryExporter struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Type                types.String                `tfsdk:"type"`
	HostName            types.String                `tfsdk:"host_name"`
	TracesPath          types.String                `tfsdk:"traces_path"`
	Port                types.Int64                 `tfsdk:"port"`
	HttpContentType     types.String                `tfsdk:"http_content_type"`
	Timeout             types.Int64                 `tfsdk:"timeout"`
	Header              types.List                  `tfsdk:"header"`
	Processor           types.String                `tfsdk:"processor"`
	MaxQueueSize        types.Int64                 `tfsdk:"max_queue_size"`
	MaxExportSize       types.Int64                 `tfsdk:"max_export_size"`
	ExportDelayInterval types.Int64                 `tfsdk:"export_delay_interval"`
	ProxyPolicies       types.List                  `tfsdk:"proxy_policies"`
	SslClient           types.String                `tfsdk:"ssl_client"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var OpenTelemetryExporterObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"type":                  types.StringType,
	"host_name":             types.StringType,
	"traces_path":           types.StringType,
	"port":                  types.Int64Type,
	"http_content_type":     types.StringType,
	"timeout":               types.Int64Type,
	"header":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType}},
	"processor":             types.StringType,
	"max_queue_size":        types.Int64Type,
	"max_export_size":       types.Int64Type,
	"export_delay_interval": types.Int64Type,
	"proxy_policies":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}},
	"ssl_client":            types.StringType,
	"dependency_actions":    actions.ActionsListType,
}

func (data OpenTelemetryExporter) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OpenTelemetryExporter"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OpenTelemetryExporter) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.HostName.IsNull() {
		return false
	}
	if !data.TracesPath.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.HttpContentType.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.Header.IsNull() {
		return false
	}
	if !data.Processor.IsNull() {
		return false
	}
	if !data.MaxQueueSize.IsNull() {
		return false
	}
	if !data.MaxExportSize.IsNull() {
		return false
	}
	if !data.ExportDelayInterval.IsNull() {
		return false
	}
	if !data.ProxyPolicies.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data OpenTelemetryExporter) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostName`, data.HostName.ValueString())
	}
	if !data.TracesPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TracesPath`, data.TracesPath.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.HttpContentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPContentType`, data.HttpContentType.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.Header.IsNull() {
		var dataValues []DmOpenTelemetryExporterHeader
		data.Header.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Header`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Processor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Processor`, data.Processor.ValueString())
	}
	if !data.MaxQueueSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxQueueSize`, data.MaxQueueSize.ValueInt64())
	}
	if !data.MaxExportSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxExportSize`, data.MaxExportSize.ValueInt64())
	}
	if !data.ExportDelayInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExportDelayInterval`, data.ExportDelayInterval.ValueInt64())
	}
	if !data.ProxyPolicies.IsNull() {
		var dataValues []DmAPIProxyPolicy
		data.ProxyPolicies.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicies`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *OpenTelemetryExporter) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("http")
	}
	if value := res.Get(pathRoot + `HostName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `TracesPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TracesPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TracesPath = types.StringValue("/v1/traces")
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(4318)
	}
	if value := res.Get(pathRoot + `HTTPContentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpContentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpContentType = types.StringValue("json")
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() {
		l := []DmOpenTelemetryExporterHeader{}
		if value := res.Get(`Header`); value.Exists() {
			for _, v := range value.Array() {
				item := DmOpenTelemetryExporterHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Header, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType}, l)
		} else {
			data.Header = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType})
		}
	} else {
		data.Header = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Processor`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Processor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Processor = types.StringValue("batch")
	}
	if value := res.Get(pathRoot + `MaxQueueSize`); value.Exists() {
		data.MaxQueueSize = types.Int64Value(value.Int())
	} else {
		data.MaxQueueSize = types.Int64Value(2048)
	}
	if value := res.Get(pathRoot + `MaxExportSize`); value.Exists() {
		data.MaxExportSize = types.Int64Value(value.Int())
	} else {
		data.MaxExportSize = types.Int64Value(512)
	}
	if value := res.Get(pathRoot + `ExportDelayInterval`); value.Exists() {
		data.ExportDelayInterval = types.Int64Value(value.Int())
	} else {
		data.ExportDelayInterval = types.Int64Value(5000)
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() {
		l := []DmAPIProxyPolicy{}
		if value := res.Get(`ProxyPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIProxyPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *OpenTelemetryExporter) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "http" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `HostName`); value.Exists() && !data.HostName.IsNull() {
		data.HostName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get(pathRoot + `TracesPath`); value.Exists() && !data.TracesPath.IsNull() {
		data.TracesPath = tfutils.ParseStringFromGJSON(value)
	} else if data.TracesPath.ValueString() != "/v1/traces" {
		data.TracesPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 4318 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `HTTPContentType`); value.Exists() && !data.HttpContentType.IsNull() {
		data.HttpContentType = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpContentType.ValueString() != "json" {
		data.HttpContentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 10 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && !data.Header.IsNull() {
		l := []DmOpenTelemetryExporterHeader{}
		for _, v := range value.Array() {
			item := DmOpenTelemetryExporterHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Header, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType}, l)
		} else {
			data.Header = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType})
		}
	} else {
		data.Header = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryExporterHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Processor`); value.Exists() && !data.Processor.IsNull() {
		data.Processor = tfutils.ParseStringFromGJSON(value)
	} else if data.Processor.ValueString() != "batch" {
		data.Processor = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxQueueSize`); value.Exists() && !data.MaxQueueSize.IsNull() {
		data.MaxQueueSize = types.Int64Value(value.Int())
	} else if data.MaxQueueSize.ValueInt64() != 2048 {
		data.MaxQueueSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxExportSize`); value.Exists() && !data.MaxExportSize.IsNull() {
		data.MaxExportSize = types.Int64Value(value.Int())
	} else if data.MaxExportSize.ValueInt64() != 512 {
		data.MaxExportSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ExportDelayInterval`); value.Exists() && !data.ExportDelayInterval.IsNull() {
		data.ExportDelayInterval = types.Int64Value(value.Int())
	} else if data.ExportDelayInterval.ValueInt64() != 5000 {
		data.ExportDelayInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() && !data.ProxyPolicies.IsNull() {
		l := []DmAPIProxyPolicy{}
		for _, v := range value.Array() {
			item := DmAPIProxyPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAPIProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}
