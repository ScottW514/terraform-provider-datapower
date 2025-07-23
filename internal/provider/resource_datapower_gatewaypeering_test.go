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

func TestAccResourceGatewayPeering(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_GatewayPeering") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_GatewayPeering")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.GatewayPeeringTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "local_port", "16380"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "primary_count", "1"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "monitor_port", "26380"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "cluster_auto_config", "true"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "enable_peer_group", "true"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "priority", "100"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "enable_ssl", "true"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "persistence_location", "memory"),
			resource.TestCheckResourceAttr("datapower_gatewaypeering.test", "local_directory", "local:///"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
