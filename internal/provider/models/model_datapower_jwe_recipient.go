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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type JWERecipient struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Algorithm         types.String                `tfsdk:"algorithm"`
	Sskey             types.String                `tfsdk:"sskey"`
	Certificate       types.String                `tfsdk:"certificate"`
	UnprotectedHeader types.List                  `tfsdk:"unprotected_header"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var JWERecipientSSKeyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "algorithm",
	AttrType:    "String",
	AttrDefault: "RSA1_5",
	Value:       []string{"A128KW", "A192KW", "A256KW", "dir"},
}
var JWERecipientCertificateCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "algorithm",
	AttrType:    "String",
	AttrDefault: "RSA1_5",
	Value:       []string{"RSA1_5", "RSA-OAEP", "RSA-OAEP-256"},
}
var JWERecipientSSKeyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "algorithm",
	AttrType:    "String",
	AttrDefault: "RSA1_5",
	Value:       []string{"A128KW", "A192KW", "A256KW", "dir"},
}
var JWERecipientCertificateIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "algorithm",
	AttrType:    "String",
	AttrDefault: "RSA1_5",
	Value:       []string{"RSA1_5", "RSA-OAEP", "RSA-OAEP-256"},
}

var JWERecipientObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"algorithm":          types.StringType,
	"sskey":              types.StringType,
	"certificate":        types.StringType,
	"unprotected_header": types.ListType{ElemType: types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}},
	"dependency_actions": actions.ActionsListType,
}

func (data JWERecipient) GetPath() string {
	rest_path := "/mgmt/config/{domain}/JWERecipient"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data JWERecipient) IsNull() bool {
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
	if !data.Sskey.IsNull() {
		return false
	}
	if !data.Certificate.IsNull() {
		return false
	}
	if !data.UnprotectedHeader.IsNull() {
		return false
	}
	return true
}

func (data JWERecipient) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Sskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSKey`, data.Sskey.ValueString())
	}
	if !data.Certificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Certificate`, data.Certificate.ValueString())
	}
	if !data.UnprotectedHeader.IsNull() {
		var dataValues []DmJOSEHeader
		data.UnprotectedHeader.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`UnprotectedHeader`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *JWERecipient) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
		data.Algorithm = types.StringValue("RSA1_5")
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Sskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
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

func (data *JWERecipient) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	} else if data.Algorithm.ValueString() != "RSA1_5" {
		data.Algorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSKey`); value.Exists() && !data.Sskey.IsNull() {
		data.Sskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `Certificate`); value.Exists() && !data.Certificate.IsNull() {
		data.Certificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Certificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `UnprotectedHeader`); value.Exists() && !data.UnprotectedHeader.IsNull() {
		l := []DmJOSEHeader{}
		e := []DmJOSEHeader{}
		data.UnprotectedHeader.ElementsAs(ctx, &e, false)
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
			data.UnprotectedHeader, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmJOSEHeaderObjectType}, l)
		} else {
			data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
		}
	} else {
		data.UnprotectedHeader = types.ListNull(types.ObjectType{AttrTypes: DmJOSEHeaderObjectType})
	}
}
