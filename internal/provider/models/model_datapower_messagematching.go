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

type MessageMatching struct {
	Id                types.String `tfsdk:"id"`
	AppDomain         types.String `tfsdk:"app_domain"`
	UserSummary       types.String `tfsdk:"user_summary"`
	IpAddress         types.String `tfsdk:"ip_address"`
	IpExclude         types.String `tfsdk:"ip_exclude"`
	HttpMethod        types.String `tfsdk:"http_method"`
	HttpHeader        types.List   `tfsdk:"http_header"`
	HttpHeaderExclude types.List   `tfsdk:"http_header_exclude"`
	RequestUrl        types.String `tfsdk:"request_url"`
}

var MessageMatchingObjectType = map[string]attr.Type{
	"id":                  types.StringType,
	"app_domain":          types.StringType,
	"user_summary":        types.StringType,
	"ip_address":          types.StringType,
	"ip_exclude":          types.StringType,
	"http_method":         types.StringType,
	"http_header":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}},
	"http_header_exclude": types.ListType{ElemType: types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}},
	"request_url":         types.StringType,
}

func (data MessageMatching) GetPath() string {
	rest_path := "/mgmt/config/{domain}/MessageMatching"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data MessageMatching) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.IpAddress.IsNull() {
		return false
	}
	if !data.IpExclude.IsNull() {
		return false
	}
	if !data.HttpMethod.IsNull() {
		return false
	}
	if !data.HttpHeader.IsNull() {
		return false
	}
	if !data.HttpHeaderExclude.IsNull() {
		return false
	}
	if !data.RequestUrl.IsNull() {
		return false
	}
	return true
}

func (data MessageMatching) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPAddress`, data.IpAddress.ValueString())
	}
	if !data.IpExclude.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IPExclude`, data.IpExclude.ValueString())
	}
	if !data.HttpMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HTTPMethod`, data.HttpMethod.ValueString())
	}
	if !data.HttpHeader.IsNull() {
		var values []DmHTTPHeader
		data.HttpHeader.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`HTTPHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.HttpHeaderExclude.IsNull() {
		var values []DmHTTPHeader
		data.HttpHeaderExclude.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`HTTPHeaderExclude`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RequestUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestURL`, data.RequestUrl.ValueString())
	}
	return body
}

func (data *MessageMatching) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPExclude`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IpExclude = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpExclude = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HttpMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HttpMethod = types.StringValue("any")
	}
	if value := res.Get(pathRoot + `HTTPHeader`); value.Exists() {
		l := []DmHTTPHeader{}
		if value := res.Get(`HTTPHeader`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHTTPHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HttpHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}, l)
		} else {
			data.HttpHeader = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
		}
	} else {
		data.HttpHeader = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
	}
	if value := res.Get(pathRoot + `HTTPHeaderExclude`); value.Exists() {
		l := []DmHTTPHeader{}
		if value := res.Get(`HTTPHeaderExclude`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHTTPHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HttpHeaderExclude, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}, l)
		} else {
			data.HttpHeaderExclude = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
		}
	} else {
		data.HttpHeaderExclude = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
	}
	if value := res.Get(pathRoot + `RequestURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestUrl = types.StringNull()
	}
}

func (data *MessageMatching) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `IPAddress`); value.Exists() && !data.IpAddress.IsNull() {
		data.IpAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `IPExclude`); value.Exists() && !data.IpExclude.IsNull() {
		data.IpExclude = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IpExclude = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPMethod`); value.Exists() && !data.HttpMethod.IsNull() {
		data.HttpMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.HttpMethod.ValueString() != "any" {
		data.HttpMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `HTTPHeader`); value.Exists() && !data.HttpHeader.IsNull() {
		l := []DmHTTPHeader{}
		for _, v := range value.Array() {
			item := DmHTTPHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HttpHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}, l)
		} else {
			data.HttpHeader = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
		}
	} else {
		data.HttpHeader = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
	}
	if value := res.Get(pathRoot + `HTTPHeaderExclude`); value.Exists() && !data.HttpHeaderExclude.IsNull() {
		l := []DmHTTPHeader{}
		for _, v := range value.Array() {
			item := DmHTTPHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HttpHeaderExclude, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPHeaderObjectType}, l)
		} else {
			data.HttpHeaderExclude = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
		}
	} else {
		data.HttpHeaderExclude = types.ListNull(types.ObjectType{AttrTypes: DmHTTPHeaderObjectType})
	}
	if value := res.Get(pathRoot + `RequestURL`); value.Exists() && !data.RequestUrl.IsNull() {
		data.RequestUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestUrl = types.StringNull()
	}
}
