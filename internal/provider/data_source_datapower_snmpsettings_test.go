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

func TestAccDataSourceSNMPSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_SNMPSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_SNMPSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.SNMPSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "local_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "local_port", "161"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "security_level", "authPriv"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "access_level", "read-only"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "enable_default_trap_subscriptions", "true"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "trap_priority", "error"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "config_mib", "/drConfigMIB.txt"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "config_mib_mq", "/mqConfigMIB.txt"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "status_mib", "/drStatusMIB.txt"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "status_mib_mq", "/mqStatusMIB.txt"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "notif_mib", "/drNotificationMIB.txt"),
					resource.TestCheckResourceAttr("data.datapower_snmpsettings.test", "notif_mib_mq", "/mqNotificationMIB.txt"),
				}...),
			},
		},
	})
}
