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

var _ validator.String = conditionalRequiredString{}

type conditionalRequiredString struct {
	RequiredConditions Evaluation
	IgnoredConditions  Evaluation // Not currently tested
	HasDefault         bool
}

func (v conditionalRequiredString) Description(ctx context.Context) string {
	return "Validates if conditions are present for the attribute to be required."
}

func (v conditionalRequiredString) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v conditionalRequiredString) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() && !v.HasDefault && v.RequiredConditions.matchesConditions(ctx, &req.Config, &resp.Diagnostics, req.Path) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Attribute Configuration",
			fmt.Sprintf("Attribute '%s' is required when %s", req.Path, v.RequiredConditions.String()),
		)
	}
}

// Helper function to create the validator.
func ConditionalRequiredString(evalRequired Evaluation, evalIgnored Evaluation, hasDefault bool) validator.String {
	return conditionalRequiredString{
		RequiredConditions: evalRequired,
		IgnoredConditions:  evalIgnored,
		HasDefault:         hasDefault,
	}
}
