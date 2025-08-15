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

package datasources_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceWebServicesAgent(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_WebServicesAgent") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_WebServicesAgent")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.WebServicesAgentTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "max_records", "3000"),
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "max_memory_kb", "64000"),
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "capture_mode", "faults"),
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "mediation_metrics", "false"),
					resource.TestCheckResourceAttr("data.datapower_web_services_agent.test", "max_payload_size_kb", "0"),
				}...),
			},
		},
	})
}
