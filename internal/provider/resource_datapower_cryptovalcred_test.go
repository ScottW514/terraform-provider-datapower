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

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
)

func TestAccResourceCryptoValCred(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_CryptoValCred") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_CryptoValCred")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.CryptoValCredTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "cert_validation_mode", "legacy"),
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "use_crl", "true"),
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "require_crl", "false"),
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "crldp_handling", "ignore"),
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "explicit_policy", "false"),
			resource.TestCheckResourceAttr("datapower_cryptovalcred.test", "check_dates", "true"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
