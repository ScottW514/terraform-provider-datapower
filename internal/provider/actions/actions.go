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
	RunOnCreate  types.Bool   `tfsdk:"run_on_create"`
	RunOnUpdate  types.Bool   `tfsdk:"run_on_update"`
	RunOnDelete  types.Bool   `tfsdk:"run_on_delete"`
	Action       types.String `tfsdk:"action"`
}

var ActionsSchema = schema.ListNestedAttribute{
	MarkdownDescription: "List of actions to take on dependent objects",
	Optional:            true,
	NestedObject: schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"target_id": schema.StringAttribute{
				MarkdownDescription: "Id of the action target (for `domains`, this must still be set, but the value is ignored)",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"target_domain": schema.StringAttribute{
				MarkdownDescription: "Application domain of the action target",
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
				MarkdownDescription: "Resource type of action target",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"run_on_create": schema.BoolAttribute{
				MarkdownDescription: "Run this action when creating this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"run_on_update": schema.BoolAttribute{
				MarkdownDescription: "Run this action when updating this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"run_on_delete": schema.BoolAttribute{
				MarkdownDescription: "Run this action when deleting this resource.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"action": schema.StringAttribute{
				MarkdownDescription: "Action to take on target",
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
			"run_on_create": types.BoolType,
			"run_on_update": types.BoolType,
			"run_on_delete": types.BoolType,
			"action":        types.StringType,
		},
	},
}

func ValidateConfig(ctx context.Context, diag *diag.Diagnostics, actions []*Action) {
	for _, target := range actions {
		if act, ok := actionMap[target.TargetType.ValueString()]; ok {
			if _, ok := act.ValidActions[target.Action.ValueString()]; !ok {
				diag.AddAttributeError(
					path.Root("object_actions"),
					"Attribute Error",
					fmt.Sprintf("'%s' action is not supported for target_type '%s'", target.Action.ValueString(), target.TargetType.ValueString()),
				)
			}
		} else {
			diag.AddAttributeError(
				path.Root("object_actions"),
				"Attribute Error",
				fmt.Sprintf("target_type '%s' not supported", target.TargetType.ValueString()),
			)
		}
	}
}
