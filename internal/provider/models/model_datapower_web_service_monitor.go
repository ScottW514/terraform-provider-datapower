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

type WebServiceMonitor struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	WsdlUrl           types.String                `tfsdk:"wsdl_url"`
	Operations        types.List                  `tfsdk:"operations"`
	EndpointName      types.String                `tfsdk:"endpoint_name"`
	EndpointUrl       types.String                `tfsdk:"endpoint_url"`
	FrontendUrl       types.String                `tfsdk:"frontend_url"`
	Transport         types.String                `tfsdk:"transport"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var WebServiceMonitorObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"wsdl_url":           types.StringType,
	"operations":         types.ListType{ElemType: types.ObjectType{AttrTypes: DmWSSLMOpsObjectType}},
	"endpoint_name":      types.StringType,
	"endpoint_url":       types.StringType,
	"frontend_url":       types.StringType,
	"transport":          types.StringType,
	"user_summary":       types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data WebServiceMonitor) GetPath() string {
	rest_path := "/mgmt/config/{domain}/WebServiceMonitor"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data WebServiceMonitor) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.WsdlUrl.IsNull() {
		return false
	}
	if !data.Operations.IsNull() {
		return false
	}
	if !data.EndpointName.IsNull() {
		return false
	}
	if !data.EndpointUrl.IsNull() {
		return false
	}
	if !data.FrontendUrl.IsNull() {
		return false
	}
	if !data.Transport.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	return true
}

func (data WebServiceMonitor) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.WsdlUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLURL`, data.WsdlUrl.ValueString())
	}
	if !data.Operations.IsNull() {
		var dataValues []DmWSSLMOps
		data.Operations.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Operations`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.EndpointName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EndpointName`, data.EndpointName.ValueString())
	}
	if !data.EndpointUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EndpointURL`, data.EndpointUrl.ValueString())
	}
	if !data.FrontendUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontendURL`, data.FrontendUrl.ValueString())
	}
	if !data.Transport.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transport`, data.Transport.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	return body
}

func (data *WebServiceMonitor) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Operations`); value.Exists() {
		l := []DmWSSLMOps{}
		if value := res.Get(`Operations`); value.Exists() {
			for _, v := range value.Array() {
				item := DmWSSLMOps{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Operations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSSLMOpsObjectType}, l)
		} else {
			data.Operations = types.ListNull(types.ObjectType{AttrTypes: DmWSSLMOpsObjectType})
		}
	} else {
		data.Operations = types.ListNull(types.ObjectType{AttrTypes: DmWSSLMOpsObjectType})
	}
	if value := res.Get(pathRoot + `EndpointName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EndpointName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointName = types.StringNull()
	}
	if value := res.Get(pathRoot + `EndpointURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EndpointUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontendURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontendUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontendUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transport`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transport = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}

func (data *WebServiceMonitor) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLURL`); value.Exists() && !data.WsdlUrl.IsNull() {
		data.WsdlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Operations`); value.Exists() && !data.Operations.IsNull() {
		l := []DmWSSLMOps{}
		e := []DmWSSLMOps{}
		data.Operations.ElementsAs(ctx, &e, false)
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
				item := DmWSSLMOps{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Operations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmWSSLMOpsObjectType}, l)
		} else {
			data.Operations = types.ListNull(types.ObjectType{AttrTypes: DmWSSLMOpsObjectType})
		}
	} else {
		data.Operations = types.ListNull(types.ObjectType{AttrTypes: DmWSSLMOpsObjectType})
	}
	if value := res.Get(pathRoot + `EndpointName`); value.Exists() && !data.EndpointName.IsNull() {
		data.EndpointName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointName = types.StringNull()
	}
	if value := res.Get(pathRoot + `EndpointURL`); value.Exists() && !data.EndpointUrl.IsNull() {
		data.EndpointUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EndpointUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontendURL`); value.Exists() && !data.FrontendUrl.IsNull() {
		data.FrontendUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontendUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transport`); value.Exists() && !data.Transport.IsNull() {
		data.Transport = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transport = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
}
