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

type ProcessingMetadata struct {
	Id          types.String `tfsdk:"id"`
	AppDomain   types.String `tfsdk:"app_domain"`
	UserSummary types.String `tfsdk:"user_summary"`
	MetaItem    types.List   `tfsdk:"meta_item"`
}

var ProcessingMetadataObjectType = map[string]attr.Type{
	"id":           types.StringType,
	"app_domain":   types.StringType,
	"user_summary": types.StringType,
	"meta_item":    types.ListType{ElemType: types.ObjectType{AttrTypes: DmMetaItemObjectType}},
}

func (data ProcessingMetadata) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ProcessingMetadata"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ProcessingMetadata) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.MetaItem.IsNull() {
		return false
	}
	return true
}

func (data ProcessingMetadata) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.MetaItem.IsNull() {
		var values []DmMetaItem
		data.MetaItem.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`MetaItem`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *ProcessingMetadata) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MetaItem`); value.Exists() {
		l := []DmMetaItem{}
		if value := res.Get(`MetaItem`); value.Exists() {
			for _, v := range value.Array() {
				item := DmMetaItem{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.MetaItem, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMetaItemObjectType}, l)
		} else {
			data.MetaItem = types.ListNull(types.ObjectType{AttrTypes: DmMetaItemObjectType})
		}
	} else {
		data.MetaItem = types.ListNull(types.ObjectType{AttrTypes: DmMetaItemObjectType})
	}
}

func (data *ProcessingMetadata) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `MetaItem`); value.Exists() && !data.MetaItem.IsNull() {
		l := []DmMetaItem{}
		for _, v := range value.Array() {
			item := DmMetaItem{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.MetaItem, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmMetaItemObjectType}, l)
		} else {
			data.MetaItem = types.ListNull(types.ObjectType{AttrTypes: DmMetaItemObjectType})
		}
	} else {
		data.MetaItem = types.ListNull(types.ObjectType{AttrTypes: DmMetaItemObjectType})
	}
}
