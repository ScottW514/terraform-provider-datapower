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

type HTTPInputConversionMap struct {
	Id                   types.String      `tfsdk:"id"`
	AppDomain            types.String      `tfsdk:"app_domain"`
	UserSummary          types.String      `tfsdk:"user_summary"`
	DefaultInputEncoding types.String      `tfsdk:"default_input_encoding"`
	InputEncoding        types.List        `tfsdk:"input_encoding"`
	DependencyActions    []*actions.Action `tfsdk:"dependency_actions"`
}

var HTTPInputConversionMapObjectType = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"user_summary":           types.StringType,
	"default_input_encoding": types.StringType,
	"input_encoding":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmInputEncodingObjectType}},
	"dependency_actions":     actions.ActionsListType,
}

func (data HTTPInputConversionMap) GetPath() string {
	rest_path := "/mgmt/config/{domain}/HTTPInputConversionMap"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data HTTPInputConversionMap) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.DefaultInputEncoding.IsNull() {
		return false
	}
	if !data.InputEncoding.IsNull() {
		return false
	}
	return true
}

func (data HTTPInputConversionMap) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.DefaultInputEncoding.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultInputEncoding`, data.DefaultInputEncoding.ValueString())
	}
	if !data.InputEncoding.IsNull() {
		var values []DmInputEncoding
		data.InputEncoding.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`InputEncoding`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *HTTPInputConversionMap) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DefaultInputEncoding`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultInputEncoding = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultInputEncoding = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputEncoding`); value.Exists() {
		l := []DmInputEncoding{}
		if value := res.Get(`InputEncoding`); value.Exists() {
			for _, v := range value.Array() {
				item := DmInputEncoding{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.InputEncoding, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmInputEncodingObjectType}, l)
		} else {
			data.InputEncoding = types.ListNull(types.ObjectType{AttrTypes: DmInputEncodingObjectType})
		}
	} else {
		data.InputEncoding = types.ListNull(types.ObjectType{AttrTypes: DmInputEncodingObjectType})
	}
}

func (data *HTTPInputConversionMap) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `DefaultInputEncoding`); value.Exists() && !data.DefaultInputEncoding.IsNull() {
		data.DefaultInputEncoding = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultInputEncoding = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputEncoding`); value.Exists() && !data.InputEncoding.IsNull() {
		l := []DmInputEncoding{}
		for _, v := range value.Array() {
			item := DmInputEncoding{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.InputEncoding, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmInputEncodingObjectType}, l)
		} else {
			data.InputEncoding = types.ListNull(types.ObjectType{AttrTypes: DmInputEncodingObjectType})
		}
	} else {
		data.InputEncoding = types.ListNull(types.ObjectType{AttrTypes: DmInputEncodingObjectType})
	}
}
