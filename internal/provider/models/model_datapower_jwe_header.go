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

type JWEHeader struct {
	Id                         types.String                `tfsdk:"id"`
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	JweProtectedHeader         types.List                  `tfsdk:"jwe_protected_header"`
	JweSharedUnprotectedHeader types.List                  `tfsdk:"jwe_shared_unprotected_header"`
	Recipient                  types.String                `tfsdk:"recipient"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var JWEHeaderObjectType = map[string]attr.Type{
	"id":                            types.StringType,
	"app_domain":                    types.StringType,
	"user_summary":                  types.StringType,
	"jwe_protected_header":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"jwe_shared_unprotected_header": types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"recipient":                     types.StringType,
	"dependency_actions":            actions.ActionsListType,
}

func (data JWEHeader) GetPath() string {
	rest_path := "/mgmt/config/{domain}/JWEHeader"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data JWEHeader) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.JweProtectedHeader.IsNull() {
		return false
	}
	if !data.JweSharedUnprotectedHeader.IsNull() {
		return false
	}
	if !data.Recipient.IsNull() {
		return false
	}
	return true
}

func (data JWEHeader) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.JweProtectedHeader.IsNull() {
		var dataValues []DmJOSEHeader
		data.JweProtectedHeader.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`JWEProtectedHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.JweSharedUnprotectedHeader.IsNull() {
		var dataValues []DmJOSEHeader
		data.JweSharedUnprotectedHeader.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`JWESharedUnprotectedHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Recipient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Recipient`, data.Recipient.ValueString())
	}
	return body
}

func (data *JWEHeader) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JWEProtectedHeader`); value.Exists() {
		l := []DmJOSEHeader{}
		if value := res.Get(`JWEProtectedHeader`); value.Exists() {
			for _, v := range value.Array() {
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.JweProtectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.JweProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.JweProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `JWESharedUnprotectedHeader`); value.Exists() {
		l := []DmJOSEHeader{}
		if value := res.Get(`JWESharedUnprotectedHeader`); value.Exists() {
			for _, v := range value.Array() {
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.JweSharedUnprotectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.JweSharedUnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.JweSharedUnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Recipient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Recipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Recipient = types.StringNull()
	}
}

func (data *JWEHeader) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `JWEProtectedHeader`); value.Exists() && !data.JweProtectedHeader.IsNull() {
		l := []DmJOSEHeader{}
		e := []DmJOSEHeader{}
		data.JweProtectedHeader.ElementsAs(ctx, &e, false)
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
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.JweProtectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.JweProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.JweProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `JWESharedUnprotectedHeader`); value.Exists() && !data.JweSharedUnprotectedHeader.IsNull() {
		l := []DmJOSEHeader{}
		e := []DmJOSEHeader{}
		data.JweSharedUnprotectedHeader.ElementsAs(ctx, &e, false)
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
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.JweSharedUnprotectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.JweSharedUnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.JweSharedUnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Recipient`); value.Exists() && !data.Recipient.IsNull() {
		data.Recipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Recipient = types.StringNull()
	}
}
