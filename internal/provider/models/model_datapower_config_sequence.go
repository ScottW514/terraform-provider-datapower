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

type ConfigSequence struct {
	Id                  types.String                  `tfsdk:"id"`
	AppDomain           types.String                  `tfsdk:"app_domain"`
	UserSummary         types.String                  `tfsdk:"user_summary"`
	Locations           types.List                    `tfsdk:"locations"`
	MatchPattern        types.String                  `tfsdk:"match_pattern"`
	ResultNamePattern   types.String                  `tfsdk:"result_name_pattern"`
	StatusNamePattern   types.String                  `tfsdk:"status_name_pattern"`
	Watch               types.Bool                    `tfsdk:"watch"`
	UseOutputLocation   types.Bool                    `tfsdk:"use_output_location"`
	OutputLocation      types.String                  `tfsdk:"output_location"`
	DeleteUnused        types.Bool                    `tfsdk:"delete_unused"`
	RunSequenceInterval types.Int64                   `tfsdk:"run_sequence_interval"`
	Capabilities        *DmConfigSequenceCapabilities `tfsdk:"capabilities"`
	DependencyActions   []*actions.DependencyAction   `tfsdk:"dependency_actions"`
}

var ConfigSequenceOutputLocationCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_output_location",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var ConfigSequenceOutputLocationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_output_location",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var ConfigSequenceObjectType = map[string]attr.Type{
	"id":                    types.StringType,
	"app_domain":            types.StringType,
	"user_summary":          types.StringType,
	"locations":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType}},
	"match_pattern":         types.StringType,
	"result_name_pattern":   types.StringType,
	"status_name_pattern":   types.StringType,
	"watch":                 types.BoolType,
	"use_output_location":   types.BoolType,
	"output_location":       types.StringType,
	"delete_unused":         types.BoolType,
	"run_sequence_interval": types.Int64Type,
	"capabilities":          types.ObjectType{AttrTypes: DmConfigSequenceCapabilitiesObjectType},
	"dependency_actions":    actions.ActionsListType,
}

func (data ConfigSequence) GetPath() string {
	rest_path := "/mgmt/config/{domain}/ConfigSequence"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data ConfigSequence) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Locations.IsNull() {
		return false
	}
	if !data.MatchPattern.IsNull() {
		return false
	}
	if !data.ResultNamePattern.IsNull() {
		return false
	}
	if !data.StatusNamePattern.IsNull() {
		return false
	}
	if !data.Watch.IsNull() {
		return false
	}
	if !data.UseOutputLocation.IsNull() {
		return false
	}
	if !data.OutputLocation.IsNull() {
		return false
	}
	if !data.DeleteUnused.IsNull() {
		return false
	}
	if !data.RunSequenceInterval.IsNull() {
		return false
	}
	if data.Capabilities != nil {
		if !data.Capabilities.IsNull() {
			return false
		}
	}
	return true
}

func (data ConfigSequence) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Locations.IsNull() {
		var dataValues []DmConfigSequenceLocation
		data.Locations.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`Locations`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`Locations`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`Locations`, "[]")
	}
	if !data.MatchPattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MatchPattern`, data.MatchPattern.ValueString())
	}
	if !data.ResultNamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResultNamePattern`, data.ResultNamePattern.ValueString())
	}
	if !data.StatusNamePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`StatusNamePattern`, data.StatusNamePattern.ValueString())
	}
	if !data.Watch.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Watch`, tfutils.StringFromBool(data.Watch, ""))
	}
	if !data.UseOutputLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseOutputLocation`, tfutils.StringFromBool(data.UseOutputLocation, ""))
	}
	if !data.OutputLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputLocation`, data.OutputLocation.ValueString())
	}
	if !data.DeleteUnused.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeleteUnused`, tfutils.StringFromBool(data.DeleteUnused, ""))
	}
	if !data.RunSequenceInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RunSequenceInterval`, data.RunSequenceInterval.ValueInt64())
	}
	if data.Capabilities != nil {
		if !data.Capabilities.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Capabilities`, data.Capabilities.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *ConfigSequence) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Locations`); value.Exists() {
		l := []DmConfigSequenceLocation{}
		if value := res.Get(`Locations`); value.Exists() {
			for _, v := range value.Array() {
				item := DmConfigSequenceLocation{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Locations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType}, l)
		} else {
			data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
		}
	} else {
		data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
	}
	if value := res.Get(pathRoot + `MatchPattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MatchPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MatchPattern = types.StringValue("(.*)\\\\.cfg$")
	}
	if value := res.Get(pathRoot + `ResultNamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResultNamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResultNamePattern = types.StringValue("$1.log")
	}
	if value := res.Get(pathRoot + `StatusNamePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.StatusNamePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.StatusNamePattern = types.StringValue("$1.status")
	}
	if value := res.Get(pathRoot + `Watch`); value.Exists() {
		data.Watch = tfutils.BoolFromString(value.String())
	} else {
		data.Watch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseOutputLocation`); value.Exists() {
		data.UseOutputLocation = tfutils.BoolFromString(value.String())
	} else {
		data.UseOutputLocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutputLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputLocation = types.StringValue("logtemp:///")
	}
	if value := res.Get(pathRoot + `DeleteUnused`); value.Exists() {
		data.DeleteUnused = tfutils.BoolFromString(value.String())
	} else {
		data.DeleteUnused = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RunSequenceInterval`); value.Exists() {
		data.RunSequenceInterval = types.Int64Value(value.Int())
	} else {
		data.RunSequenceInterval = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `Capabilities`); value.Exists() {
		data.Capabilities = &DmConfigSequenceCapabilities{}
		data.Capabilities.FromBody(ctx, "", value)
	} else {
		data.Capabilities = nil
	}
}

func (data *ConfigSequence) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Locations`); value.Exists() && !data.Locations.IsNull() {
		l := []DmConfigSequenceLocation{}
		e := []DmConfigSequenceLocation{}
		data.Locations.ElementsAs(ctx, &e, false)
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
				item := DmConfigSequenceLocation{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Locations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType}, l)
		} else {
			data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
		}
	} else {
		data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
	}
	if value := res.Get(pathRoot + `MatchPattern`); value.Exists() && !data.MatchPattern.IsNull() {
		data.MatchPattern = tfutils.ParseStringFromGJSON(value)
	} else if data.MatchPattern.ValueString() != "(.*)\\\\.cfg$" {
		data.MatchPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `ResultNamePattern`); value.Exists() && !data.ResultNamePattern.IsNull() {
		data.ResultNamePattern = tfutils.ParseStringFromGJSON(value)
	} else if data.ResultNamePattern.ValueString() != "$1.log" {
		data.ResultNamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `StatusNamePattern`); value.Exists() && !data.StatusNamePattern.IsNull() {
		data.StatusNamePattern = tfutils.ParseStringFromGJSON(value)
	} else if data.StatusNamePattern.ValueString() != "$1.status" {
		data.StatusNamePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `Watch`); value.Exists() && !data.Watch.IsNull() {
		data.Watch = tfutils.BoolFromString(value.String())
	} else if !data.Watch.ValueBool() {
		data.Watch = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseOutputLocation`); value.Exists() && !data.UseOutputLocation.IsNull() {
		data.UseOutputLocation = tfutils.BoolFromString(value.String())
	} else if data.UseOutputLocation.ValueBool() {
		data.UseOutputLocation = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OutputLocation`); value.Exists() && !data.OutputLocation.IsNull() {
		data.OutputLocation = tfutils.ParseStringFromGJSON(value)
	} else if data.OutputLocation.ValueString() != "logtemp:///" {
		data.OutputLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `DeleteUnused`); value.Exists() && !data.DeleteUnused.IsNull() {
		data.DeleteUnused = tfutils.BoolFromString(value.String())
	} else if !data.DeleteUnused.ValueBool() {
		data.DeleteUnused = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RunSequenceInterval`); value.Exists() && !data.RunSequenceInterval.IsNull() {
		data.RunSequenceInterval = types.Int64Value(value.Int())
	} else if data.RunSequenceInterval.ValueInt64() != 100 {
		data.RunSequenceInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Capabilities`); value.Exists() {
		data.Capabilities.UpdateFromBody(ctx, "", value)
	} else {
		data.Capabilities = nil
	}
}
func (data *ConfigSequence) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Id.IsUnknown() {
		if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
			data.Id = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Id = types.StringNull()
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.Locations.IsUnknown() {
		if value := res.Get(pathRoot + `Locations`); value.Exists() && !data.Locations.IsNull() {
			l := []DmConfigSequenceLocation{}
			if value := res.Get(`Locations`); value.Exists() {
				for _, v := range value.Array() {
					item := DmConfigSequenceLocation{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.Locations, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType}, l)
			} else {
				data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
			}
		} else {
			data.Locations = types.ListNull(types.ObjectType{AttrTypes: DmConfigSequenceLocationObjectType})
		}
	}
	if data.MatchPattern.IsUnknown() {
		if value := res.Get(pathRoot + `MatchPattern`); value.Exists() && !data.MatchPattern.IsNull() {
			data.MatchPattern = tfutils.ParseStringFromGJSON(value)
		} else if data.MatchPattern.ValueString() != "(.*)\\\\.cfg$" {
			data.MatchPattern = types.StringNull()
		}
	}
	if data.ResultNamePattern.IsUnknown() {
		if value := res.Get(pathRoot + `ResultNamePattern`); value.Exists() && !data.ResultNamePattern.IsNull() {
			data.ResultNamePattern = tfutils.ParseStringFromGJSON(value)
		} else if data.ResultNamePattern.ValueString() != "$1.log" {
			data.ResultNamePattern = types.StringNull()
		}
	}
	if data.StatusNamePattern.IsUnknown() {
		if value := res.Get(pathRoot + `StatusNamePattern`); value.Exists() && !data.StatusNamePattern.IsNull() {
			data.StatusNamePattern = tfutils.ParseStringFromGJSON(value)
		} else if data.StatusNamePattern.ValueString() != "$1.status" {
			data.StatusNamePattern = types.StringNull()
		}
	}
	if data.Watch.IsUnknown() {
		if value := res.Get(pathRoot + `Watch`); value.Exists() && !data.Watch.IsNull() {
			data.Watch = tfutils.BoolFromString(value.String())
		} else {
			data.Watch = types.BoolNull()
		}
	}
	if data.UseOutputLocation.IsUnknown() {
		if value := res.Get(pathRoot + `UseOutputLocation`); value.Exists() && !data.UseOutputLocation.IsNull() {
			data.UseOutputLocation = tfutils.BoolFromString(value.String())
		} else {
			data.UseOutputLocation = types.BoolNull()
		}
	}
	if data.OutputLocation.IsUnknown() {
		if value := res.Get(pathRoot + `OutputLocation`); value.Exists() && !data.OutputLocation.IsNull() {
			data.OutputLocation = tfutils.ParseStringFromGJSON(value)
		} else if data.OutputLocation.ValueString() != "logtemp:///" {
			data.OutputLocation = types.StringNull()
		}
	}
	if data.DeleteUnused.IsUnknown() {
		if value := res.Get(pathRoot + `DeleteUnused`); value.Exists() && !data.DeleteUnused.IsNull() {
			data.DeleteUnused = tfutils.BoolFromString(value.String())
		} else {
			data.DeleteUnused = types.BoolNull()
		}
	}
	if data.RunSequenceInterval.IsUnknown() {
		if value := res.Get(pathRoot + `RunSequenceInterval`); value.Exists() && !data.RunSequenceInterval.IsNull() {
			data.RunSequenceInterval = types.Int64Value(value.Int())
		} else if data.RunSequenceInterval.ValueInt64() != 100 {
			data.RunSequenceInterval = types.Int64Null()
		}
	}
	if data.Capabilities == nil {
		if value := res.Get(pathRoot + `Capabilities`); value.Exists() {
			d := DmConfigSequenceCapabilities{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.Capabilities = &d
			}
		}
	}
}
