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

type OpenTelemetry struct {
	Id                types.String      `tfsdk:"id"`
	AppDomain         types.String      `tfsdk:"app_domain"`
	UserSummary       types.String      `tfsdk:"user_summary"`
	Exporter          types.String      `tfsdk:"exporter"`
	Sampler           types.String      `tfsdk:"sampler"`
	ResourceAttribute types.List        `tfsdk:"resource_attribute"`
	ObjectActions     []*actions.Action `tfsdk:"object_actions"`
}

var OpenTelemetryObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"exporter":           types.StringType,
	"sampler":            types.StringType,
	"resource_attribute": types.ListType{ElemType: types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}},
	"object_actions":     actions.ActionsListType,
}

func (data OpenTelemetry) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OpenTelemetry"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OpenTelemetry) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Exporter.IsNull() {
		return false
	}
	if !data.Sampler.IsNull() {
		return false
	}
	if !data.ResourceAttribute.IsNull() {
		return false
	}
	return true
}

func (data OpenTelemetry) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Exporter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Exporter`, data.Exporter.ValueString())
	}
	if !data.Sampler.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Sampler`, data.Sampler.ValueString())
	}
	if !data.ResourceAttribute.IsNull() {
		var values []DmOpenTelemetryResourceAttribute
		data.ResourceAttribute.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ResourceAttribute`+".-1", val.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *OpenTelemetry) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Exporter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Exporter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Exporter = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sampler`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Sampler = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sampler = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResourceAttribute`); value.Exists() {
		l := []DmOpenTelemetryResourceAttribute{}
		if value := res.Get(`ResourceAttribute`); value.Exists() {
			for _, v := range value.Array() {
				item := DmOpenTelemetryResourceAttribute{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ResourceAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}, l)
		} else {
			data.ResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
		}
	} else {
		data.ResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
	}
}

func (data *OpenTelemetry) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Exporter`); value.Exists() && !data.Exporter.IsNull() {
		data.Exporter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Exporter = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sampler`); value.Exists() && !data.Sampler.IsNull() {
		data.Sampler = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sampler = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResourceAttribute`); value.Exists() && !data.ResourceAttribute.IsNull() {
		l := []DmOpenTelemetryResourceAttribute{}
		for _, v := range value.Array() {
			item := DmOpenTelemetryResourceAttribute{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.ResourceAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType}, l)
		} else {
			data.ResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
		}
	} else {
		data.ResourceAttribute = types.ListNull(types.ObjectType{AttrTypes: DmOpenTelemetryResourceAttributeObjectType})
	}
}
