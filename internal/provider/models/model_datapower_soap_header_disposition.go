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

type SOAPHeaderDisposition struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Refine            types.List                  `tfsdk:"refine"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var SOAPHeaderDispositionObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"refine":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data SOAPHeaderDisposition) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SOAPHeaderDisposition"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SOAPHeaderDisposition) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Refine.IsNull() {
		return false
	}
	return true
}

func (data SOAPHeaderDisposition) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Refine.IsNull() {
		var dataValues []DmSOAPHeaderDispositionItem
		data.Refine.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Refine`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *SOAPHeaderDisposition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Refine`); value.Exists() {
		l := []DmSOAPHeaderDispositionItem{}
		if value := res.Get(`Refine`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSOAPHeaderDispositionItem{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Refine, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType}, l)
		} else {
			data.Refine = types.ListNull(types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType})
		}
	} else {
		data.Refine = types.ListNull(types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType})
	}
}

func (data *SOAPHeaderDisposition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Refine`); value.Exists() && !data.Refine.IsNull() {
		l := []DmSOAPHeaderDispositionItem{}
		e := []DmSOAPHeaderDispositionItem{}
		data.Refine.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSOAPHeaderDispositionItem{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Refine, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType}, l)
		} else {
			data.Refine = types.ListNull(types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType})
		}
	} else {
		data.Refine = types.ListNull(types.ObjectType{AttrTypes: DmSOAPHeaderDispositionItemObjectType})
	}
}
