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

type AppSecurityPolicy struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	RequestMaps       types.List                  `tfsdk:"request_maps"`
	ResponseMaps      types.List                  `tfsdk:"response_maps"`
	ErrorMaps         types.List                  `tfsdk:"error_maps"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget    types.String                `tfsdk:"provider_target"`
}

var AppSecurityPolicyObjectType = map[string]attr.Type{
	"provider_target":    types.StringType,
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"request_maps":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType}},
	"response_maps":      types.ListType{ElemType: types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType}},
	"error_maps":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmPolicyMapObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data AppSecurityPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AppSecurityPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AppSecurityPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.RequestMaps.IsNull() {
		return false
	}
	if !data.ResponseMaps.IsNull() {
		return false
	}
	if !data.ErrorMaps.IsNull() {
		return false
	}
	return true
}

func (data AppSecurityPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.RequestMaps.IsNull() {
		var dataValues []DmWebAppRequestPolicyMap
		data.RequestMaps.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`RequestMaps`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`RequestMaps`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`RequestMaps`, "[]")
	}
	if !data.ResponseMaps.IsNull() {
		var dataValues []DmWebAppResponsePolicyMap
		data.ResponseMaps.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`ResponseMaps`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`ResponseMaps`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`ResponseMaps`, "[]")
	}
	if !data.ErrorMaps.IsNull() {
		var dataValues []DmPolicyMap
		data.ErrorMaps.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`ErrorMaps`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`ErrorMaps`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`ErrorMaps`, "[]")
	}
	return body
}

func (data *AppSecurityPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RequestMaps`); value.Exists() {
		l := []DmWebAppRequestPolicyMap{}
		if value := res.Get(`RequestMaps`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWebAppRequestPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RequestMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType}, l)
		} else {
			data.RequestMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType})
		}
	} else {
		data.RequestMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType})
	}
	if value := res.Get(pathRoot + `ResponseMaps`); value.Exists() {
		l := []DmWebAppResponsePolicyMap{}
		if value := res.Get(`ResponseMaps`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWebAppResponsePolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ResponseMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType}, l)
		} else {
			data.ResponseMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType})
		}
	} else {
		data.ResponseMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType})
	}
	if value := res.Get(pathRoot + `ErrorMaps`); value.Exists() {
		l := []DmPolicyMap{}
		if value := res.Get(`ErrorMaps`); value.Exists() {
			for _, v := range value.Array() {
				item := DmPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ErrorMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyMapObjectType}, l)
		} else {
			data.ErrorMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
		}
	} else {
		data.ErrorMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
	}
}

func (data *AppSecurityPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `RequestMaps`); value.Exists() && !data.RequestMaps.IsNull() {
		l := []DmWebAppRequestPolicyMap{}
		e := []DmWebAppRequestPolicyMap{}
		data.RequestMaps.ElementsAs(ctx, &e, false)
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
				item := DmWebAppRequestPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.RequestMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType}, l)
		} else {
			data.RequestMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType})
		}
	} else {
		data.RequestMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppRequestPolicyMapObjectType})
	}
	if value := res.Get(pathRoot + `ResponseMaps`); value.Exists() && !data.ResponseMaps.IsNull() {
		l := []DmWebAppResponsePolicyMap{}
		e := []DmWebAppResponsePolicyMap{}
		data.ResponseMaps.ElementsAs(ctx, &e, false)
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
				item := DmWebAppResponsePolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ResponseMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType}, l)
		} else {
			data.ResponseMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType})
		}
	} else {
		data.ResponseMaps = types.ListNull(types.ObjectType{AttrTypes: DmWebAppResponsePolicyMapObjectType})
	}
	if value := res.Get(pathRoot + `ErrorMaps`); value.Exists() && !data.ErrorMaps.IsNull() {
		l := []DmPolicyMap{}
		e := []DmPolicyMap{}
		data.ErrorMaps.ElementsAs(ctx, &e, false)
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
				item := DmPolicyMap{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ErrorMaps, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyMapObjectType}, l)
		} else {
			data.ErrorMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
		}
	} else {
		data.ErrorMaps = types.ListNull(types.ObjectType{AttrTypes: DmPolicyMapObjectType})
	}
}
