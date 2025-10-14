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

type B2BCPACollaboration struct {
	Id                   types.String                `tfsdk:"id"`
	AppDomain            types.String                `tfsdk:"app_domain"`
	UserSummary          types.String                `tfsdk:"user_summary"`
	InternalRole         types.String                `tfsdk:"internal_role"`
	ExternalRole         types.String                `tfsdk:"external_role"`
	ProcessSpecification types.String                `tfsdk:"process_specification"`
	Service              types.String                `tfsdk:"service"`
	ServiceType          types.String                `tfsdk:"service_type"`
	SenderMshSetting     types.String                `tfsdk:"sender_msh_setting"`
	ReceiverMshSetting   types.String                `tfsdk:"receiver_msh_setting"`
	Actions              types.List                  `tfsdk:"actions"`
	DependencyActions    []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BCPACollaborationInternalRoleCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationInternalRoleIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationExternalRoleCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationExternalRoleIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationServiceTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationSenderMshSettingCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationSenderMshSettingIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var B2BCPACollaborationReceiverMshSettingCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationReceiverMshSettingIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var B2BCPACollaborationActionsCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "service",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"urn:oasis:names:tc:ebxml-msg:service"},
}

var B2BCPACollaborationObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"internal_role":         types.StringType,
	"external_role":         types.StringType,
	"process_specification": types.StringType,
	"service":               types.StringType,
	"service_type":          types.StringType,
	"sender_msh_setting":    types.StringType,
	"receiver_msh_setting":  types.StringType,
	"actions":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType}},
	"dependency_actions":    actions.ActionsListType,
}

func (data B2BCPACollaboration) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BCPACollaboration"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BCPACollaboration) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.InternalRole.IsNull() {
		return false
	}
	if !data.ExternalRole.IsNull() {
		return false
	}
	if !data.ProcessSpecification.IsNull() {
		return false
	}
	if !data.Service.IsNull() {
		return false
	}
	if !data.ServiceType.IsNull() {
		return false
	}
	if !data.SenderMshSetting.IsNull() {
		return false
	}
	if !data.ReceiverMshSetting.IsNull() {
		return false
	}
	if !data.Actions.IsNull() {
		return false
	}
	return true
}

func (data B2BCPACollaboration) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.InternalRole.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InternalRole`, data.InternalRole.ValueString())
	}
	if !data.ExternalRole.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalRole`, data.ExternalRole.ValueString())
	}
	if !data.ProcessSpecification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ProcessSpecification`, data.ProcessSpecification.ValueString())
	}
	if !data.Service.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Service`, data.Service.ValueString())
	}
	if !data.ServiceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServiceType`, data.ServiceType.ValueString())
	}
	if !data.SenderMshSetting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SenderMshSetting`, data.SenderMshSetting.ValueString())
	}
	if !data.ReceiverMshSetting.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReceiverMshSetting`, data.ReceiverMshSetting.ValueString())
	}
	if !data.Actions.IsNull() {
		var dataValues []DmCPACollaborationAction
		data.Actions.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`Actions`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Actions`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Actions`, "[]")
	}
	return body
}

func (data *B2BCPACollaboration) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `InternalRole`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InternalRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRole`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessSpecification`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ProcessSpecification = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessSpecification = types.StringNull()
	}
	if value := res.Get(pathRoot + `Service`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Service = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Service = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServiceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServiceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderMshSetting`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SenderMshSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderMshSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverMshSetting`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReceiverMshSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverMshSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() {
		l := []DmCPACollaborationAction{}
		if value := res.Get(`Actions`); value.Exists() {
			for _, v := range value.Array() {
				item := DmCPACollaborationAction{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Actions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType}, l)
		} else {
			data.Actions = types.ListNull(types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType})
		}
	} else {
		data.Actions = types.ListNull(types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType})
	}
}

func (data *B2BCPACollaboration) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `InternalRole`); value.Exists() && !data.InternalRole.IsNull() {
		data.InternalRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InternalRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalRole`); value.Exists() && !data.ExternalRole.IsNull() {
		data.ExternalRole = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalRole = types.StringNull()
	}
	if value := res.Get(pathRoot + `ProcessSpecification`); value.Exists() && !data.ProcessSpecification.IsNull() {
		data.ProcessSpecification = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ProcessSpecification = types.StringNull()
	}
	if value := res.Get(pathRoot + `Service`); value.Exists() && !data.Service.IsNull() {
		data.Service = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Service = types.StringNull()
	}
	if value := res.Get(pathRoot + `ServiceType`); value.Exists() && !data.ServiceType.IsNull() {
		data.ServiceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServiceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderMshSetting`); value.Exists() && !data.SenderMshSetting.IsNull() {
		data.SenderMshSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderMshSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReceiverMshSetting`); value.Exists() && !data.ReceiverMshSetting.IsNull() {
		data.ReceiverMshSetting = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReceiverMshSetting = types.StringNull()
	}
	if value := res.Get(pathRoot + `Actions`); value.Exists() && !data.Actions.IsNull() {
		l := []DmCPACollaborationAction{}
		e := []DmCPACollaborationAction{}
		data.Actions.ElementsAs(ctx, &e, false)
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
				item := DmCPACollaborationAction{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Actions, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType}, l)
		} else {
			data.Actions = types.ListNull(types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType})
		}
	} else {
		data.Actions = types.ListNull(types.ObjectType{AttrTypes: DmCPACollaborationActionObjectType})
	}
}
