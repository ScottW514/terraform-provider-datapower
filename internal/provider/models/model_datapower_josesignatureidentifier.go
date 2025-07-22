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

type JOSESignatureIdentifier struct {
	Id              types.String `tfsdk:"id"`
	AppDomain       types.String `tfsdk:"app_domain"`
	UserSummary     types.String `tfsdk:"user_summary"`
	Type            types.String `tfsdk:"type"`
	SsKey           types.String `tfsdk:"ss_key"`
	Certificate     types.String `tfsdk:"certificate"`
	ValidAlgorithms types.List   `tfsdk:"valid_algorithms"`
	HeaderParam     types.List   `tfsdk:"header_param"`
	Verify          types.Bool   `tfsdk:"verify"`
}

var JOSESignatureIdentifierObjectType = map[string]attr.Type{
	"id":               types.StringType,
	"app_domain":       types.StringType,
	"user_summary":     types.StringType,
	"type":             types.StringType,
	"ss_key":           types.StringType,
	"certificate":      types.StringType,
	"valid_algorithms": types.ListType{ElemType: types.StringType},
	"header_param":     types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"verify":           types.BoolType,
}

func (data JOSESignatureIdentifier) GetPath() string {
	rest_path := "/mgmt/config/{domain}/JOSESignatureIdentifier"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data JOSESignatureIdentifier) IsNull() bool {
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
	if !data.SsKey.IsNull() {
		return false
	}
	if !data.Certificate.IsNull() {
		return false
	}
	if !data.ValidAlgorithms.IsNull() {
		return false
	}
	if !data.HeaderParam.IsNull() {
		return false
	}
	if !data.Verify.IsNull() {
		return false
	}
	return true
}

func (data JOSESignatureIdentifier) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.SsKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSKey`, data.SsKey.ValueString())
	}
	if !data.Certificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Certificate`, data.Certificate.ValueString())
	}
	if !data.ValidAlgorithms.IsNull() {
		var values []string
		data.ValidAlgorithms.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`ValidAlgorithms`+".-1", map[string]string{"value": val})
		}
	}
	if !data.HeaderParam.IsNull() {
		var values []DmJOSEHeader
		data.HeaderParam.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderParam`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Verify.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Verify`, tfutils.StringFromBool(data.Verify, false))
	}
	return body
}

func (data *JOSESignatureIdentifier) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SsKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SsKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidAlgorithms`); value.Exists() {
		data.ValidAlgorithms = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ValidAlgorithms = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `HeaderParam`); value.Exists() {
		l := []DmJOSEHeader{}
		if value := res.Get(`HeaderParam`); value.Exists() {
			for _, v := range value.Array() {
				item := DmJOSEHeader{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderParam, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.HeaderParam = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.HeaderParam = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Verify`); value.Exists() {
		data.Verify = tfutils.BoolFromString(value.String())
	} else {
		data.Verify = types.BoolNull()
	}
}

func (data *JOSESignatureIdentifier) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && !data.SsKey.IsNull() {
		data.SsKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SsKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidAlgorithms`); value.Exists() && !data.ValidAlgorithms.IsNull() {
		data.ValidAlgorithms = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ValidAlgorithms = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `HeaderParam`); value.Exists() && !data.HeaderParam.IsNull() {
		l := []DmJOSEHeader{}
		for _, v := range value.Array() {
			item := DmJOSEHeader{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.HeaderParam, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.HeaderParam = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.HeaderParam = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
	if value := res.Get(pathRoot + `Verify`); value.Exists() && !data.Verify.IsNull() {
		data.Verify = tfutils.BoolFromString(value.String())
	} else if !data.Verify.ValueBool() {
		data.Verify = types.BoolNull()
	}
}
