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

func TestAccResourceCryptoKerberosKDC(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_CryptoKerberosKDC") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_CryptoKerberosKDC")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.CryptoKerberosKDCTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "use_tcp", "false"),
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "server_port", "88"),
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "udp_timeout", "5"),
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "cache_tickets", "true"),
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "max_cached_tickets", "32"),
			resource.TestCheckResourceAttr("datapower_cryptokerberoskdc.test", "min_cached_ticket_validity", "60"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
