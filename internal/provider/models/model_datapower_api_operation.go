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

type APIOperation struct {
	Id                types.String                `tfsdk:"id"`
	AppDomain         types.String                `tfsdk:"app_domain"`
	UserSummary       types.String                `tfsdk:"user_summary"`
	Method            types.String                `tfsdk:"method"`
	OperationId       types.String                `tfsdk:"operation_id"`
	RemoveConsume     types.Bool                  `tfsdk:"remove_consume"`
	Consume           types.List                  `tfsdk:"consume"`
	Produce           types.List                  `tfsdk:"produce"`
	RequestSchema     types.String                `tfsdk:"request_schema"`
	ResponseSchema    types.List                  `tfsdk:"response_schema"`
	Parameter         types.List                  `tfsdk:"parameter"`
	RemoveSecurity    types.Bool                  `tfsdk:"remove_security"`
	Security          types.List                  `tfsdk:"security"`
	SoapAction        types.String                `tfsdk:"soap_action"`
	SoapElementName   types.String                `tfsdk:"soap_element_name"`
	DependencyActions []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var APIOperationConsumeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "remove_consume",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APIOperationSecurityIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "remove_security",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var APIOperationObjectType = map[string]attr.Type{
	"id":                 types.StringType,
	"app_domain":         types.StringType,
	"user_summary":       types.StringType,
	"method":             types.StringType,
	"operation_id":       types.StringType,
	"remove_consume":     types.BoolType,
	"consume":            types.ListType{ElemType: types.StringType},
	"produce":            types.ListType{ElemType: types.StringType},
	"request_schema":     types.StringType,
	"response_schema":    types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType}},
	"parameter":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIParameterObjectType}},
	"remove_security":    types.BoolType,
	"security":           types.ListType{ElemType: types.StringType},
	"soap_action":        types.StringType,
	"soap_element_name":  types.StringType,
	"dependency_actions": actions.ActionsListType,
}

func (data APIOperation) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIOperation"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIOperation) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Method.IsNull() {
		return false
	}
	if !data.OperationId.IsNull() {
		return false
	}
	if !data.RemoveConsume.IsNull() {
		return false
	}
	if !data.Consume.IsNull() {
		return false
	}
	if !data.Produce.IsNull() {
		return false
	}
	if !data.RequestSchema.IsNull() {
		return false
	}
	if !data.ResponseSchema.IsNull() {
		return false
	}
	if !data.Parameter.IsNull() {
		return false
	}
	if !data.RemoveSecurity.IsNull() {
		return false
	}
	if !data.Security.IsNull() {
		return false
	}
	if !data.SoapAction.IsNull() {
		return false
	}
	if !data.SoapElementName.IsNull() {
		return false
	}
	return true
}

func (data APIOperation) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Method.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Method`, data.Method.ValueString())
	}
	if !data.OperationId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OperationId`, data.OperationId.ValueString())
	}
	if !data.RemoveConsume.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoveConsume`, tfutils.StringFromBool(data.RemoveConsume, ""))
	}
	if !data.Consume.IsNull() {
		var dataValues []string
		data.Consume.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Consume`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Produce.IsNull() {
		var dataValues []string
		data.Produce.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Produce`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RequestSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequestSchema`, data.RequestSchema.ValueString())
	}
	if !data.ResponseSchema.IsNull() {
		var dataValues []DmAPIResponseSchema
		data.ResponseSchema.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ResponseSchema`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Parameter.IsNull() {
		var dataValues []DmAPIParameter
		data.Parameter.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Parameter`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.RemoveSecurity.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoveSecurity`, tfutils.StringFromBool(data.RemoveSecurity, ""))
	}
	if !data.Security.IsNull() {
		var dataValues []string
		data.Security.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`Security`+".-1", map[string]string{"value": val})
		}
	}
	if !data.SoapAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPAction`, data.SoapAction.ValueString())
	}
	if !data.SoapElementName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPElementName`, data.SoapElementName.ValueString())
	}
	return body
}

func (data *APIOperation) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Method`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Method = types.StringValue("GET")
	}
	if value := res.Get(pathRoot + `OperationId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OperationId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationId = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoveConsume`); value.Exists() {
		data.RemoveConsume = tfutils.BoolFromString(value.String())
	} else {
		data.RemoveConsume = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Consume`); value.Exists() {
		data.Consume = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Consume = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Produce`); value.Exists() {
		data.Produce = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Produce = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequestSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RequestSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSchema`); value.Exists() {
		l := []DmAPIResponseSchema{}
		if value := res.Get(`ResponseSchema`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIResponseSchema{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ResponseSchema, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType}, l)
		} else {
			data.ResponseSchema = types.ListNull(types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType})
		}
	} else {
		data.ResponseSchema = types.ListNull(types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType})
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() {
		l := []DmAPIParameter{}
		if value := res.Get(`Parameter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
	}
	if value := res.Get(pathRoot + `RemoveSecurity`); value.Exists() {
		data.RemoveSecurity = tfutils.BoolFromString(value.String())
	} else {
		data.RemoveSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Security`); value.Exists() {
		data.Security = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Security = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SOAPAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `SOAPElementName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapElementName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapElementName = types.StringNull()
	}
}

func (data *APIOperation) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Method`); value.Exists() && !data.Method.IsNull() {
		data.Method = tfutils.ParseStringFromGJSON(value)
	} else if data.Method.ValueString() != "GET" {
		data.Method = types.StringNull()
	}
	if value := res.Get(pathRoot + `OperationId`); value.Exists() && !data.OperationId.IsNull() {
		data.OperationId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OperationId = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoveConsume`); value.Exists() && !data.RemoveConsume.IsNull() {
		data.RemoveConsume = tfutils.BoolFromString(value.String())
	} else if data.RemoveConsume.ValueBool() {
		data.RemoveConsume = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Consume`); value.Exists() && !data.Consume.IsNull() {
		data.Consume = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Consume = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Produce`); value.Exists() && !data.Produce.IsNull() {
		data.Produce = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Produce = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequestSchema`); value.Exists() && !data.RequestSchema.IsNull() {
		data.RequestSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RequestSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResponseSchema`); value.Exists() && !data.ResponseSchema.IsNull() {
		l := []DmAPIResponseSchema{}
		e := []DmAPIResponseSchema{}
		data.ResponseSchema.ElementsAs(ctx, &e, false)
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
				item := DmAPIResponseSchema{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ResponseSchema, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType}, l)
		} else {
			data.ResponseSchema = types.ListNull(types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType})
		}
	} else {
		data.ResponseSchema = types.ListNull(types.ObjectType{AttrTypes: DmAPIResponseSchemaObjectType})
	}
	if value := res.Get(pathRoot + `Parameter`); value.Exists() && !data.Parameter.IsNull() {
		l := []DmAPIParameter{}
		e := []DmAPIParameter{}
		data.Parameter.ElementsAs(ctx, &e, false)
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
				item := DmAPIParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Parameter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIParameterObjectType}, l)
		} else {
			data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
		}
	} else {
		data.Parameter = types.ListNull(types.ObjectType{AttrTypes: DmAPIParameterObjectType})
	}
	if value := res.Get(pathRoot + `RemoveSecurity`); value.Exists() && !data.RemoveSecurity.IsNull() {
		data.RemoveSecurity = tfutils.BoolFromString(value.String())
	} else if data.RemoveSecurity.ValueBool() {
		data.RemoveSecurity = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Security`); value.Exists() && !data.Security.IsNull() {
		data.Security = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Security = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SOAPAction`); value.Exists() && !data.SoapAction.IsNull() {
		data.SoapAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `SOAPElementName`); value.Exists() && !data.SoapElementName.IsNull() {
		data.SoapElementName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapElementName = types.StringNull()
	}
}
