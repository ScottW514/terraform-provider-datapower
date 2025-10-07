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

type PolicyAttachments struct {
	Id                            types.String                `tfsdk:"id"`
	AppDomain                     types.String                `tfsdk:"app_domain"`
	UserSummary                   types.String                `tfsdk:"user_summary"`
	EnforcementMode               types.String                `tfsdk:"enforcement_mode"`
	PolicyReferences              types.Bool                  `tfsdk:"policy_references"`
	IgnoredPolicyAttachmentPoints types.List                  `tfsdk:"ignored_policy_attachment_points"`
	ExternalPolicy                types.List                  `tfsdk:"external_policy"`
	SlaEnforcementMode            types.String                `tfsdk:"sla_enforcement_mode"`
	DependencyActions             []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var PolicyAttachmentsIgnoredPolicyAttachmentPointsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "policy_references",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var PolicyAttachmentsObjectType = map[string]attr.Type{
	"id":                               types.StringType,
	"app_domain":                       types.StringType,
	"user_summary":                     types.StringType,
	"enforcement_mode":                 types.StringType,
	"policy_references":                types.BoolType,
	"ignored_policy_attachment_points": types.ListType{ElemType: types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType}},
	"external_policy":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType}},
	"sla_enforcement_mode":             types.StringType,
	"dependency_actions":               actions.ActionsListType,
}

func (data PolicyAttachments) GetPath() string {
	rest_path := "/mgmt/config/{domain}/PolicyAttachments"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data PolicyAttachments) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.EnforcementMode.IsNull() {
		return false
	}
	if !data.PolicyReferences.IsNull() {
		return false
	}
	if !data.IgnoredPolicyAttachmentPoints.IsNull() {
		return false
	}
	if !data.ExternalPolicy.IsNull() {
		return false
	}
	if !data.SlaEnforcementMode.IsNull() {
		return false
	}
	return true
}

func (data PolicyAttachments) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.EnforcementMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforcementMode`, data.EnforcementMode.ValueString())
	}
	if !data.PolicyReferences.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyReferences`, tfutils.StringFromBool(data.PolicyReferences, ""))
	}
	if !data.IgnoredPolicyAttachmentPoints.IsNull() {
		var dataValues []DmPolicyAttachmentPoint
		data.IgnoredPolicyAttachmentPoints.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`IgnoredPolicyAttachmentPoints`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ExternalPolicy.IsNull() {
		var dataValues []DmExternalAttachedPolicy
		data.ExternalPolicy.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ExternalPolicy`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SlaEnforcementMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SLAEnforcementMode`, data.SlaEnforcementMode.ValueString())
	}
	return body
}

func (data *PolicyAttachments) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EnforcementMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EnforcementMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EnforcementMode = types.StringValue("enforce")
	}
	if value := res.Get(pathRoot + `PolicyReferences`); value.Exists() {
		data.PolicyReferences = tfutils.BoolFromString(value.String())
	} else {
		data.PolicyReferences = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IgnoredPolicyAttachmentPoints`); value.Exists() {
		l := []DmPolicyAttachmentPoint{}
		if value := res.Get(`IgnoredPolicyAttachmentPoints`); value.Exists() {
			for _, v := range value.Array() {
				item := DmPolicyAttachmentPoint{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.IgnoredPolicyAttachmentPoints, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType}, l)
		} else {
			data.IgnoredPolicyAttachmentPoints = types.ListNull(types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType})
		}
	} else {
		data.IgnoredPolicyAttachmentPoints = types.ListNull(types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType})
	}
	if value := res.Get(pathRoot + `ExternalPolicy`); value.Exists() {
		l := []DmExternalAttachedPolicy{}
		if value := res.Get(`ExternalPolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := DmExternalAttachedPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ExternalPolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType}, l)
		} else {
			data.ExternalPolicy = types.ListNull(types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType})
		}
	} else {
		data.ExternalPolicy = types.ListNull(types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SLAEnforcementMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SlaEnforcementMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SlaEnforcementMode = types.StringValue("allow-if-no-sla")
	}
}

func (data *PolicyAttachments) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `EnforcementMode`); value.Exists() && !data.EnforcementMode.IsNull() {
		data.EnforcementMode = tfutils.ParseStringFromGJSON(value)
	} else if data.EnforcementMode.ValueString() != "enforce" {
		data.EnforcementMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyReferences`); value.Exists() && !data.PolicyReferences.IsNull() {
		data.PolicyReferences = tfutils.BoolFromString(value.String())
	} else if data.PolicyReferences.ValueBool() {
		data.PolicyReferences = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IgnoredPolicyAttachmentPoints`); value.Exists() && !data.IgnoredPolicyAttachmentPoints.IsNull() {
		l := []DmPolicyAttachmentPoint{}
		e := []DmPolicyAttachmentPoint{}
		data.IgnoredPolicyAttachmentPoints.ElementsAs(ctx, &e, false)
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
				item := DmPolicyAttachmentPoint{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.IgnoredPolicyAttachmentPoints, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType}, l)
		} else {
			data.IgnoredPolicyAttachmentPoints = types.ListNull(types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType})
		}
	} else {
		data.IgnoredPolicyAttachmentPoints = types.ListNull(types.ObjectType{AttrTypes: DmPolicyAttachmentPointObjectType})
	}
	if value := res.Get(pathRoot + `ExternalPolicy`); value.Exists() && !data.ExternalPolicy.IsNull() {
		l := []DmExternalAttachedPolicy{}
		e := []DmExternalAttachedPolicy{}
		data.ExternalPolicy.ElementsAs(ctx, &e, false)
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
				item := DmExternalAttachedPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ExternalPolicy, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType}, l)
		} else {
			data.ExternalPolicy = types.ListNull(types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType})
		}
	} else {
		data.ExternalPolicy = types.ListNull(types.ObjectType{AttrTypes: DmExternalAttachedPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SLAEnforcementMode`); value.Exists() && !data.SlaEnforcementMode.IsNull() {
		data.SlaEnforcementMode = tfutils.ParseStringFromGJSON(value)
	} else if data.SlaEnforcementMode.ValueString() != "allow-if-no-sla" {
		data.SlaEnforcementMode = types.StringNull()
	}
}
