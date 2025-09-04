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

// Package tfutils provides utilities for Terraform-related functionality.
package tfutils

import (
	"fmt"
	"strings"
)

// AttributeDescription holds a formatted description for a Terraform attribute.
type AttributeDescription struct {
	String string
}

// NewAttributeDescription creates a new AttributeDescription with the given initial description.
func NewAttributeDescription(description string, cliAlias string, referenceTo string) *AttributeDescription {
	s := description + "\n"
	if cliAlias != "" {
		s = fmt.Sprintf("%s  - CLI Alias: `%s`\n", s, cliAlias)
	}
	if referenceTo != "" {
		s = fmt.Sprintf("%s  - Reference to: `datapower_%s:id`\n", s, referenceTo)
	}
	return &AttributeDescription{s}
}

// AddDefault adds a default value to the attribute description.
// It appends a formatted default value to the description and returns the updated AttributeDescription for chaining.
func (d *AttributeDescription) AddDefaultValue(defaultValue string) *AttributeDescription {
	d.String = fmt.Sprintf("%s  - Default value: `%s`\n", d.String, defaultValue)
	return d
}

// AddIntegerRange adds the minimum and maximum values to the attribute description.
// It appends a formatted range to the description and returns the updated AttributeDescription for chaining.
func (d *AttributeDescription) AddIntegerRange(min, max int64) *AttributeDescription {
	d.String = fmt.Sprintf("%s  - Range: `%v`-`%v`\n", d.String, min, max)
	return d
}

// AddStringEnum adds a list of valid enum values to the attribute description.
// It appends a formatted list of choices to the description and returns the updated AttributeDescription for chaining.
func (d *AttributeDescription) AddStringEnum(values ...string) *AttributeDescription {
	v := make([]string, len(values))
	for i, value := range values {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s  - Choices: %s\n", d.String, strings.Join(v, ", "))
	return d
}

// AddActions adds a list of valid actions to the resource description.
// It appends a formatted list of actions to the description and returns the updated AttributeDescription for chaining.
func (d *AttributeDescription) AddActions(actions ...string) *AttributeDescription {
	v := make([]string, len(actions))
	for i, value := range actions {
		v[i] = fmt.Sprintf("`%s`", value)
	}
	d.String = fmt.Sprintf("%s  - Accepted Dependency Actions: %s\n", d.String, strings.Join(v, ", "))
	return d
}

// AddRequiredWhen adds required when conditions.
func (d *AttributeDescription) AddRequiredWhen(condRequired string) *AttributeDescription {
	d.String = fmt.Sprintf("%s  - Required When: %s\n", d.String, condRequired)
	return d
}

// AddNotValidWhen adds required when conditions.
func (d *AttributeDescription) AddNotValidWhen(condNotValid string) *AttributeDescription {
	d.String = fmt.Sprintf("%s  - Not Valid When: %s\n", d.String, condNotValid)
	return d
}
