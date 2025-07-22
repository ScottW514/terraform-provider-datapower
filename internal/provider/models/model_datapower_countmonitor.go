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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CountMonitor struct {
	Id          types.String `tfsdk:"id"`
	AppDomain   types.String `tfsdk:"app_domain"`
	Measure     types.String `tfsdk:"measure"`
	Source      types.String `tfsdk:"source"`
	Header      types.String `tfsdk:"header"`
	Filter      types.List   `tfsdk:"filter"`
	MaxSources  types.Int64  `tfsdk:"max_sources"`
	UserSummary types.String `tfsdk:"user_summary"`
	MessageType types.String `tfsdk:"message_type"`
}

var CountMonitorObjectType = map[string]attr.Type{
	"id":           types.StringType,
	"app_domain":   types.StringType,
	"measure":      types.StringType,
	"source":       types.StringType,
	"header":       types.StringType,
	"filter":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType}},
	"max_sources":  types.Int64Type,
	"user_summary": types.StringType,
	"message_type": types.StringType,
}

func (data CountMonitor) GetPath() string {
	rest_path := "/mgmt/config/{domain}/CountMonitor"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data CountMonitor) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Measure.IsNull() {
		return false
	}
	if !data.Source.IsNull() {
		return false
	}
	if !data.Header.IsNull() {
		return false
	}
	if !data.Filter.IsNull() {
		return false
	}
	if !data.MaxSources.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MessageType.IsNull() {
		return false
	}
	return true
}

func (data CountMonitor) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Measure.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Measure`, data.Measure.ValueString())
	}
	if !data.Source.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Source`, data.Source.ValueString())
	}
	if !data.Header.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Header`, data.Header.ValueString())
	}
	if !data.Filter.IsNull() {
		var values []DmCountMonitorFilter
		data.Filter.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Filter`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.MaxSources.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxSources`, data.MaxSources.ValueInt64())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.MessageType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessageType`, data.MessageType.ValueString())
	}
	return body
}

func (data *CountMonitor) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Measure`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Measure = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Measure = types.StringValue("requests")
	}
	if value := res.Get(pathRoot + `Source`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Source = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Source = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Header = types.StringValue("X-Client-IP")
	}
	if value := res.Get(pathRoot + `Filter`); value.Exists() {
		l := []DmCountMonitorFilter{}
		if value := res.Get(`Filter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmCountMonitorFilter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Filter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType}, l)
		} else {
			data.Filter = types.ListNull(types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType})
		}
	} else {
		data.Filter = types.ListNull(types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType})
	}
	if value := res.Get(pathRoot + `MaxSources`); value.Exists() {
		data.MaxSources = types.Int64Value(value.Int())
	} else {
		data.MaxSources = types.Int64Value(10000)
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `MessageType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MessageType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageType = types.StringNull()
	}
}

func (data *CountMonitor) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Measure`); value.Exists() && !data.Measure.IsNull() {
		data.Measure = tfutils.ParseStringFromGJSON(value)
	} else if data.Measure.ValueString() != "requests" {
		data.Measure = types.StringNull()
	}
	if value := res.Get(pathRoot + `Source`); value.Exists() && !data.Source.IsNull() {
		data.Source = tfutils.ParseStringFromGJSON(value)
	} else if data.Source.ValueString() != "all" {
		data.Source = types.StringNull()
	}
	if value := res.Get(pathRoot + `Header`); value.Exists() && !data.Header.IsNull() {
		data.Header = tfutils.ParseStringFromGJSON(value)
	} else if data.Header.ValueString() != "X-Client-IP" {
		data.Header = types.StringNull()
	}
	if value := res.Get(pathRoot + `Filter`); value.Exists() && !data.Filter.IsNull() {
		l := []DmCountMonitorFilter{}
		for _, v := range value.Array() {
			item := DmCountMonitorFilter{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Filter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType}, l)
		} else {
			data.Filter = types.ListNull(types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType})
		}
	} else {
		data.Filter = types.ListNull(types.ObjectType{AttrTypes: DmCountMonitorFilterObjectType})
	}
	if value := res.Get(pathRoot + `MaxSources`); value.Exists() && !data.MaxSources.IsNull() {
		data.MaxSources = types.Int64Value(value.Int())
	} else if data.MaxSources.ValueInt64() != 10000 {
		data.MaxSources = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `MessageType`); value.Exists() && !data.MessageType.IsNull() {
		data.MessageType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MessageType = types.StringNull()
	}
}
