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

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Int64 = conditionalRequiredInt64{}

type conditionalRequiredInt64 struct {
	RequiredConditions Evaluation
	IgnoredConditions  Evaluation // Not currently tested
	HasDefault         bool
}

func (v conditionalRequiredInt64) Description(ctx context.Context) string {
	return "Validates if conditions are present for the attribute to be required."
}

func (v conditionalRequiredInt64) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v conditionalRequiredInt64) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	matchesRequired := v.RequiredConditions.matchesConditions(ctx, &req.Config, &resp.Diagnostics, req.Path)
	if !req.ConfigValue.IsNull() && matchesRequired {
		return
	} else if req.ConfigValue.IsNull() && !v.HasDefault && matchesRequired {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Attribute Configuration",
			fmt.Sprintf("Attribute '%s' is required when %s", req.Path, v.RequiredConditions.String()),
		)
	} else if !req.ConfigValue.IsNull() && v.IgnoredConditions.matchesConditions(ctx, &req.Config, &resp.Diagnostics, req.Path) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Attribute Configuration",
			fmt.Sprintf("Attribute '%s' is not expected when %s", req.Path, v.IgnoredConditions.String()),
		)
	}
}

// Helper function to create the validator.
func ConditionalRequiredInt64(evalRequired Evaluation, evalIgnored Evaluation, hasDefault bool) validator.Int64 {
	return conditionalRequiredInt64{
		RequiredConditions: evalRequired,
		IgnoredConditions:  evalIgnored,
		HasDefault:         hasDefault,
	}
}
