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

func TestAccResourceAPIConnectGatewayService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_APIConnectGatewayService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_APIConnectGatewayService")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.APIConnectGatewayServiceTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "enabled", "false"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "local_port", "3000"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "api_gateway_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "api_gateway_port", "9443"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "gateway_peering_manager", "default"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "v5_compatibility_mode", "true"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "v5c_slm_mode", "autounicast"),
			resource.TestCheckResourceAttr("datapower_apiconnectgatewayservice.test", "jwt_validation_mode", "request"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
