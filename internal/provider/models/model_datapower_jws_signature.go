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

type JWSSignature struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Algorithm         types.String                `tfsdk:"algorithm"`
	Key               types.String                `tfsdk:"key"`
	SsKey             types.String                `tfsdk:"ss_key"`
	ProtectedHeader   types.List                  `tfsdk:"protected_header"`
	UnprotectedHeader types.List                  `tfsdk:"unprotected_header"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var JWSSignatureObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"algorithm":          types.StringType,
	"key":                types.StringType,
	"ss_key":             types.StringType,
	"protected_header":   types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"unprotected_header": types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data JWSSignature) GetPath() string {
	rest_path := "/mgmt/config/{domain}/JWSSignature"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data JWSSignature) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Algorithm.IsNull() {
		return false
	}
	if !data.Key.IsNull() {
		return false
	}
	if !data.SsKey.IsNull() {
		return false
	}
	if !data.ProtectedHeader.IsNull() {
		return false
	}
	if !data.UnprotectedHeader.IsNull() {
		return false
	}
	return true
}

func (data JWSSignature) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Algorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Algorithm`, data.Algorithm.ValueString())
	}
	if !data.Key.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Key`, data.Key.ValueString())
	}
	if !data.SsKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSKey`, data.SsKey.ValueString())
	}
	if !data.ProtectedHeader.IsNull() {
		var values []DmJOSEHeader
		data.ProtectedHeader.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ProtectedHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UnprotectedHeader.IsNull() {
		var values []DmJOSEHeader
		data.UnprotectedHeader.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`UnprotectedHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *JWSSignature) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Algorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Algorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Algorithm = types.StringValue("RS256")
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SsKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SsKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProtectedHeader`); value.Exists() {
		l := []DmJOSEHeader{}
		if value := res.Get(`ProtectedHeader`); value.Exists() {
			for _, v := range value.Array() {
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ProtectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.ProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.ProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `UnprotectedHeader`); value.Exists() {
		l := []DmJOSEHeader{}
		if value := res.Get(`UnprotectedHeader`); value.Exists() {
			for _, v := range value.Array() {
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UnprotectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
}

func (data *JWSSignature) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Algorithm`); value.Exists() && !data.Algorithm.IsNull() {
		data.Algorithm = tfutils.ParseStringFromGJSON(value)
	} else if data.Algorithm.ValueString() != "RS256" {
		data.Algorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `Key`); value.Exists() && !data.Key.IsNull() {
		data.Key = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Key = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && !data.SsKey.IsNull() {
		data.SsKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SsKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProtectedHeader`); value.Exists() && !data.ProtectedHeader.IsNull() {
		l := []DmJOSEHeader{}
		for _, v := range value.Array() {
			item := DmJOSEHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ProtectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.ProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.ProtectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `UnprotectedHeader`); value.Exists() && !data.UnprotectedHeader.IsNull() {
		l := []DmJOSEHeader{}
		for _, v := range value.Array() {
			item := DmJOSEHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.UnprotectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
}
