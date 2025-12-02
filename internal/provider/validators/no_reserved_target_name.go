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

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Map = noReservedTargetNameValidator{}

type noReservedTargetNameValidator struct{}

func (v noReservedTargetNameValidator) Description(ctx context.Context) string {
	return `The map key "_base_" is reserved for internal use and cannot be used as a target name.`
}

func (v noReservedTargetNameValidator) MarkdownDescription(ctx context.Context) string {
	return `The map key **_base_** is reserved for internal use and cannot be used as a target name.`
}

func (v noReservedTargetNameValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	// If null or unknown, skip validation (optional attribute)
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// Iterate over the map keys
	for key := range req.ConfigValue.Elements() {
		if key == "_base_" {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path.AtMapKey(key),
				v.Description(ctx),
				key,
			))
			return
		}
	}
}

func NoReservedTargetName() validator.Map {
	return noReservedTargetNameValidator{}
}
