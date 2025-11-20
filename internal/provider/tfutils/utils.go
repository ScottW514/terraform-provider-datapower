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
	"sync"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)

type ProviderData struct {
	Client *client.DatapowerClient
	Mu     sync.Mutex
}

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

// ToTfName converts a string (possibly with spaces or hyphens) to snake_case.
func ToTfName(s string) string {
	abbreviations := []string{
		"AAA",
		"API",
		"AAA",
		"AS1",
		"AS2",
		"AS3",
		"LTPA",
		"AU",
		"AZ",
		"B2B",
		"BP10",
		"Cache",
		"CRL",
		"EBMS2",
		"EBMS3",
		"EBMS",
		"GraphQL",
		"HMAC",
		"Http500",
		"Source",
		"Scheme",
		"HTTPS",
		"HTTP2",
		"HTTP",
		"Issuer",
		"JWS",
		"JWT",
		"LDAP",
		"MDN",
		"MQv9",
		"SSL",
		"NSS",
		"OAuth",
		"OIDC",
		"PASV",
		"PORT",
		"Remote",
		"SAML2",
		"SAML",
		"Script",
		"SMTP",
		"SM",
		"SOAP",
		"SSKey",
		"TAM",
		"TLS13",
		"TLS",
		"X509",
		"XACML",
		"XPath",
		"WSDL",
		"WSRR",
		"GWS",
		"WS",
	}
	for _, a := range abbreviations {
		s = strings.ReplaceAll(s, a, "-"+strings.ToLower(a)+"-")
	}

	var b strings.Builder
	i := 0
	n := len(s)
	for i < n {
		r := rune(s[i])
		if unicode.IsSpace(r) || r == '-' || r == '_' {
			if b.Len() > 0 && b.String()[b.Len()-1] != '_' {
				b.WriteRune('_')
			}
			i++
			continue
		}

		isUpper := unicode.IsUpper(r)

		insert := false
		if isUpper && i > 0 {
			prev := rune(s[i-1])
			isPrevLower := unicode.IsLower(prev)
			isPrevUpper := unicode.IsUpper(prev)
			if isPrevLower {
				insert = true
			} else if isPrevUpper {
				if i+1 < n {
					next := rune(s[i+1])
					isNextLower := unicode.IsLower(next)
					if isNextLower {
						insert = true
					}
				}
			}
		}

		if insert {
			b.WriteRune('_')
		}

		b.WriteRune(unicode.ToLower(r))
		i++
	}
	return strings.Trim(b.String(), "_")
}

// Test if credentials are good and domain exists
func DomainCredentialTest(c *client.DatapowerClient, diag *diag.Diagnostics, domain string) bool {
	res, err := c.Get("/mgmt/domains/config//")
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to retrieve domain list, got error: %s", err))
		return false
	}
	return res.Get(`domain.#(name=="` + domain + `")`).Exists()
}
