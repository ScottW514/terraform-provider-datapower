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

package actions

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Operation int

const (
	Create Operation = iota
	Update
	Delete
)

type Action struct {
	TargetId     types.String `tfsdk:"target_id"`
	TargetDomain types.String `tfsdk:"target_domain"`
	TargetType   types.String `tfsdk:"target_type"`
	OnCreate     types.Bool   `tfsdk:"on_create"`
	OnUpdate     types.Bool   `tfsdk:"on_update"`
	OnDelete     types.Bool   `tfsdk:"on_delete"`
	Action       types.String `tfsdk:"action"`
}

var ActionsSchema = schema.ListNestedAttribute{
	MarkdownDescription: "Actions to take on other resources when operations are performed on this resource.",
	Optional:            true,
	NestedObject: schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"target_id": schema.StringAttribute{
				MarkdownDescription: "Id of the target for the action (required for all resources except `resource_datapower_domain`)",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_domain": schema.StringAttribute{
				MarkdownDescription: "Application domain of the target for the action",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_type": schema.StringAttribute{
				MarkdownDescription: "Resource type of the target for the action",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"on_create": schema.BoolAttribute{
				MarkdownDescription: "Execute this action on the target when creating this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"on_update": schema.BoolAttribute{
				MarkdownDescription: "Execute this action on the target when updating this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"on_delete": schema.BoolAttribute{
				MarkdownDescription: "Execute this action on the target when deleting this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action to take on target resource",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	},
}

var ActionsListType = types.ListType{
	ElemType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"target_id":     types.StringType,
			"target_domain": types.StringType,
			"target_type":   types.StringType,
			"on_create":     types.BoolType,
			"on_update":     types.BoolType,
			"on_delete":     types.BoolType,
			"action":        types.StringType,
		},
	},
}

func ValidateConfig(ctx context.Context, diag *diag.Diagnostics, actions []*Action) {
	for _, target := range actions {
		if act, ok := actionMap[target.TargetType.ValueString()]; ok {
			if _, ok := act.ValidActions[target.Action.ValueString()]; !ok {
				diag.AddAttributeError(
					path.Root("dependency_actions"),
					"Attribute Error",
					fmt.Sprintf("'%s' `action` is not supported for `target_type` '%s'", target.Action.ValueString(), target.TargetType.ValueString()),
				)
			} else {
				if target.TargetType.ValueString() != "resource_datapower_domain" && target.TargetId.IsNull() {
					diag.AddAttributeError(
						path.Root("dependency_actions"), "Attribute Error", "`target_id` is required when `target_type` is not `resource_datapower_domain`")
				}
			}
		} else {
			diag.AddAttributeError(
				path.Root("dependency_actions"),
				"Attribute Error",
				fmt.Sprintf("`target_type` '%s' not supported", target.TargetType.ValueString()),
			)
		}
	}
}
