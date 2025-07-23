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
	"path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CRLFetch struct {
	Enabled        types.Bool `tfsdk:"enabled"`
	CrlFetchConfig types.List `tfsdk:"crl_fetch_config"`
}

var CRLFetchObjectType = map[string]attr.Type{
	"enabled":          types.BoolType,
	"crl_fetch_config": types.ListType{ElemType: types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType}},
}

func (data CRLFetch) GetPath() string {
	rest_path := "/mgmt/config/default/CRLFetch/CRL-Settings"
	return rest_path
}

func (data CRLFetch) IsNull() bool {
	if !data.Enabled.IsNull() {
		return false
	}
	if !data.CrlFetchConfig.IsNull() {
		return false
	}
	return true
}

func (data CRLFetch) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	body, _ = sjson.Set(body, "CRLFetch.name", path.Base("/mgmt/config/default/CRLFetch/CRL-Settings"))
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`mAdminState`, tfutils.StringFromBool(data.Enabled, "admin"))
	}
	if !data.CrlFetchConfig.IsNull() {
		var values []DmCRLFetchConfig
		data.CrlFetchConfig.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`CRLFetchConfig`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *CRLFetch) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CRLFetchConfig`); value.Exists() {
		l := []DmCRLFetchConfig{}
		if value := res.Get(`CRLFetchConfig`); value.Exists() {
			for _, v := range value.Array() {
				item := DmCRLFetchConfig{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.CrlFetchConfig, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType}, l)
		} else {
			data.CrlFetchConfig = types.ListNull(types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType})
		}
	} else {
		data.CrlFetchConfig = types.ListNull(types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType})
	}
}

func (data *CRLFetch) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `mAdminState`); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = tfutils.BoolFromString(value.String())
	} else if !data.Enabled.ValueBool() {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CRLFetchConfig`); value.Exists() && !data.CrlFetchConfig.IsNull() {
		l := []DmCRLFetchConfig{}
		for _, v := range value.Array() {
			item := DmCRLFetchConfig{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.CrlFetchConfig, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType}, l)
		} else {
			data.CrlFetchConfig = types.ListNull(types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType})
		}
	} else {
		data.CrlFetchConfig = types.ListNull(types.ObjectType{AttrTypes: DmCRLFetchConfigObjectType})
	}
}
