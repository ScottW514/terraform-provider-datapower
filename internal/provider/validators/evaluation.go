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

package validators

import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Evaluation struct {
	Evaluation  string
	Conditions  []Evaluation
	Attribute   string
	AttrType    string
	AttrDefault string
	AttrPath    string
	Value       []string
}

func (e *Evaluation) matchesConditions(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) bool {
	if e.Evaluation == "" {
		return false
	}
	if e.Evaluation[0:7] != "logical" && len(e.Attribute) == 0 {
		diag.AddAttributeWarning(attrPath, "Attribute Validation Bug", "Attribute Missing. Report this to the provider developers.")
		return false
	}
	if e.Evaluation[0:7] != "logical" && len(e.AttrType) == 0 {
		diag.AddAttributeWarning(attrPath, "Attribute Validation Bug", "AttrType Missing. Report this to the provider developers.")
		return false
	}
	if e.Evaluation[0:7] != "logical" && e.AttrType == "List" {
		for _, v := range e.Value {
			if v != "" {
				diag.AddAttributeWarning(attrPath, "Attribute Validation Bug", "List with value other than ''. Report this to the provider developers.")
				return false
			}
		}
	}
	switch e.Evaluation {
	case "logical-true":
		return true
	case "logical-not":
		return !e.Conditions[0].matchesConditions(ctx, config, diag, attrPath)
	case "logical-and":
		for _, r := range e.Conditions {
			if !r.matchesConditions(ctx, config, diag, attrPath) {
				return false
			}
		}
		return true
	case "logical-or":
		for _, r := range e.Conditions {
			if r.matchesConditions(ctx, config, diag, attrPath) {
				return true
			}
		}
		return false
	case "property-url-protocol-in-list":
		return e.urlInValues(ctx, config, diag, attrPath)
	case "property-url-protocol-not-in-list":
		return !e.urlInValues(ctx, config, diag, attrPath)
	case "property-value-in-list":
		return e.attrInValues(ctx, config, diag, attrPath)
	case "property-value-not-in-list":
		return !e.attrInValues(ctx, config, diag, attrPath)
	case "property-less-than":
		var intAttr types.Int64
		diags := config.GetAttribute(ctx, attrPath.ParentPath().AtName(e.Attribute), &intAttr)
		diag.Append(diags...)
		if diags.HasError() {
			return false
		}
		var a int64 = 0
		v, _ := strconv.ParseInt(e.Value[0], 10, 64)
		if intAttr.IsNull() && e.AttrDefault != "" {
			a, _ = strconv.ParseInt(e.AttrDefault, 10, 64)
		} else if !intAttr.IsNull() {
			a = intAttr.ValueInt64()
		}
		return a < v
	default:
		diag.AddAttributeError(
			attrPath,
			"Invalid Attribute Validation",
			fmt.Sprintf("Attribute evaluation test '%s' not recognized.", e.Evaluation),
		)
	}
	return false
}

func (e *Evaluation) attrInValues(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) bool {
	adjPath := attrPath
	if e.AttrPath != "" {
		for {
			if e.AttrPath != "" && e.AttrPath[0:3] != "../" {
				adjPath = adjPath.AtName(e.AttrPath)
				break
			} else if e.AttrPath[0:3] == "../" {
				e.AttrPath = e.AttrPath[3:]
				adjPath = adjPath.ParentPath()
			} else {
				break
			}
		}
	}
	var attrValues []string
	if e.AttrType == "String" {
		attrValues = append(attrValues, e.getString(ctx, config, diag, adjPath))
	} else if e.AttrType == "Int64" {
		attrValues = append(attrValues, e.getInt64String(ctx, config, diag, adjPath))
	} else if e.AttrType == "Bool" {
		attrValues = append(attrValues, e.getBoolString(ctx, config, diag, adjPath))
	} else if e.AttrType == "List" {
		// The only test on "List" types is if they are empty or not
		// If list has no length, we just add an empty string, otherwise we just "not-empty" so it won't match a test for an empty string
		var listAttr types.List
		diag.Append(config.GetAttribute(ctx, adjPath.ParentPath().AtName(e.Attribute), &listAttr)...)
		if !diag.HasError() {
			if len(listAttr.Elements()) > 0 {
				attrValues = append(attrValues, "not-empty")
			} else {
				attrValues = append(attrValues, "")
			}
		}
	} else if e.AttrType[0:2] == "Dm" {
		// All Dm types that are tested are maps of bools
		for _, v := range e.Value {
			if e.getBoolMapString(ctx, config, diag, adjPath, v) == "true" {
				attrValues = append(attrValues, v)
			}
			if diag.HasError() {
				return false
			}
		}
	} else {
		diag.AddAttributeError(
			attrPath,
			"Invalid Attribute Validation",
			fmt.Sprintf("Property type '%s' not valid to evaluate attribute '%s'.", e.AttrType, e.Attribute),
		)
	}
	if diag.HasError() {
		return false
	}
	for _, v := range e.Value {
		if slices.Contains(attrValues, v) {
			return true
		}
	}
	return false
}

func (e *Evaluation) urlInValues(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) bool {
	urlStr := e.getString(ctx, config, diag, attrPath)
	if diag.HasError() {
		return false
	}
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		diag.AddAttributeError(attrPath, "Error Validating Attribute", fmt.Sprintf("Could not parse url: %s", err))
		return false
	}
	return slices.Contains(e.Value, parsedURL.Scheme)
}

func (e *Evaluation) getString(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) string {
	var stringAttr types.String
	diag.Append(config.GetAttribute(ctx, attrPath.ParentPath().AtName(e.Attribute), &stringAttr)...)
	if !diag.HasError() {
		if stringAttr.IsNull() && e.AttrDefault != "" {
			return e.AttrDefault
		} else {
			return stringAttr.ValueString()
		}
	}
	return ""
}

func (e *Evaluation) getInt64String(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) string {
	var intAttr types.Int64
	diag.Append(config.GetAttribute(ctx, attrPath.ParentPath().AtName(e.Attribute), &intAttr)...)
	if !diag.HasError() {
		if intAttr.IsNull() && e.AttrDefault != "" {
			return e.AttrDefault
		} else {
			return strconv.FormatInt(intAttr.ValueInt64(), 10)
		}
	}
	return ""
}

func (e *Evaluation) getBoolString(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path) string {
	var boolAttr types.Bool
	diag.Append(config.GetAttribute(ctx, attrPath.ParentPath().AtName(e.Attribute), &boolAttr)...)
	if !diag.HasError() {
		if boolAttr.IsNull() && e.AttrDefault != "" {
			return e.AttrDefault
		} else if !boolAttr.IsNull() {
			if boolAttr.ValueBool() {
				return "true"
			} else {
				return "false"
			}
		}
	}
	return ""
}

func (e *Evaluation) getBoolMapString(ctx context.Context, config *tfsdk.Config, diag *diag.Diagnostics, attrPath path.Path, attrName string) string {
	// Some DataPower types are a map of bool values
	var boolAttr types.Bool
	diag.Append(config.GetAttribute(ctx, attrPath.ParentPath().AtName(e.Attribute).AtName(attrName), &boolAttr)...)
	if !diag.HasError() {
		if boolAttr.IsNull() && e.AttrDefault != "" {
			return e.AttrDefault
		} else if !boolAttr.IsNull() {
			if boolAttr.ValueBool() {
				return "true"
			} else {
				return "false"
			}
		}
	}
	return ""
}

func (e *Evaluation) String() string {
	switch e.Evaluation {
	case "logical-true":
		return "attribute is not conditionally required"
	case "logical-not":
		return fmt.Sprintf("NOT%s", e.Conditions[0].String())
	case "logical-and":
		var c []string
		for _, r := range e.Conditions {
			c = append(c, r.String())
		}
		return "(" + strings.Join(c, " AND ") + ")"
	case "logical-or":
		var c []string
		for _, r := range e.Conditions {
			c = append(c, r.String())
		}
		return "(" + strings.Join(c, " OR ") + ")"
	case "property-url-protocol-in-list":
		return fmt.Sprintf("`%s` protocol=%s", e.Attribute, fmtSliceToString(e.Value))
	case "property-url-protocol-not-in-list":
		return fmt.Sprintf("`%s` protocol!=%s", e.Attribute, fmtSliceToString(e.Value))
	case "property-value-in-list":
		return fmt.Sprintf("`%s`=%s", e.Attribute, fmtSliceToString(e.Value))
	case "property-value-not-in-list":
		return fmt.Sprintf("`%s`!=%s", e.Attribute, fmtSliceToString(e.Value))
	case "property-less-than":
		return fmt.Sprintf("`%s`<`%s`", e.Attribute, e.Value[0])
	}
	return ""
}

func fmtSliceToString(s []string) string {
	var q []string
	for _, str := range s {
		q = append(q, fmt.Sprintf("`%s`", strings.Trim(strconv.Quote(str), `"`)))
	}
	return strings.Join(q, `|`)
}
