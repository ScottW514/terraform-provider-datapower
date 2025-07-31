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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)

// BoolFromString returns a types.Bool from a string.
// It expects a string and converts "on"/"enabled" to true, otherwise false.
func BoolFromString(value string) types.Bool {
	if value == "on" || value == "enabled" || value == "yes" {
		return types.BoolValue(true)
	}
	return types.BoolValue(false)
}

// StringFromBool returns a string from types.Bool.
// If adminState is false, it "on" for true, and "off" for false.
// If adminState is true, it "enabled" for true, and "disabled" for false.
func StringFromBool(value types.Bool, boolType string) string {
	switch boolType {
	case "admin":
		if value.ValueBool() {
			return "enabled"
		}
		return "disabled"
	case "flag":
		if value.ValueBool() {
			return "yes"
		}
		return "no"
	default:
		if value.ValueBool() {
			return "on"
		}
		return "off"
	}
}

// ParseBoolFromGJSON parses a bool from a gjson.Result.
// It expects a string result, and converts "on"/"enabled" to true, otherwise false.
func ParseBoolFromGJSON(result gjson.Result) types.Bool {
	if result.String() == "on" || result.String() == "enabled" || result.String() == "yes" {
		return types.BoolValue(true)
	}
	return types.BoolValue(false)
}

// ParseStringFromGJSON parses a string from a gjson.Result.
// It extracts the "value" field if it exists; otherwise, it uses the result's string representation.
func ParseStringFromGJSON(result gjson.Result) types.String {
	if value := result.Get("value"); value.Exists() {
		return types.StringValue(value.String())
	}
	return types.StringValue(result.String())
}

// ParseStringListFromGJSON converts a gjson.Result to a Terraform List of strings.
// It handles arrays by extracting the "value" field from each item or using the item's string representation.
// For single objects or other types, it extracts the "value" field or uses the string representation.
func ParseStringListFromGJSON(result gjson.Result) types.List {
	var values []attr.Value
	if result.IsArray() {
		for _, item := range result.Array() {
			if value := item.Get("value"); value.Exists() {
				values = append(values, types.StringValue(value.String()))
			} else {
				values = append(values, types.StringValue(item.String()))
			}
		}
	} else {
		var value string
		if v := result.Get("value"); v.Exists() {
			value = v.String()
		} else {
			value = result.String()
		}
		values = append(values, types.StringValue(value))
	}

	return types.ListValueMust(types.StringType, values)
}

// GenerateHash generates a secure bcrypt hash for the given text.
// It returns the hash as a string and an error if the hashing fails.
func GenerateHash(text string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hash)
}

// VerifyHash verifies if the given plaintext matches the provided bcrypt hash.
// It returns true if the plaintext matches the hash, false otherwise.
func VerifyHash(text, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text)) == nil
}

// Restarts domain
func RestartDomain(client *client.DatapowerClient, domain string) error {
	res, err := client.Post(fmt.Sprintf("/mgmt/actionqueue/%s", domain), fmt.Sprintf(`{"RestartDomain": {"Domain": "%s"}}`, domain))
	if err == nil {
		if res.StatusCode() != 202 {
			return fmt.Errorf("client: failed to restart domain. received status %d, expected 202. %s", res.StatusCode(), res.Body())
		}
	}

	// Wait for domain restart to complete for up to 5 seconds
	for attempt := 1; attempt <= 10; attempt++ {
		_, cErr := client.Get(fmt.Sprintf("/mgmt/actionqueue/%s/operations", domain))
		if cErr != nil && !strings.Contains(cErr.Error(), "status 404") {
			return fmt.Errorf("client: failed to restart domain. %s", cErr)
		}

		time.Sleep(time.Millisecond * 500)
	}
	return err
}
