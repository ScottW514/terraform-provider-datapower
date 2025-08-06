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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceMQManager(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MQManager") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MQManager")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.MQManagerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.id", "MQManager_name"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ccsid", "819"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.channel_name", "SYSTEM.DEF.SVRCONN"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.heartbeat", "300"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.maximum_message_size", "1048576"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.cache_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ffst_size", "500"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ffst_rotate", "3"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.units_of_work", "0"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.automatic_backout", "false"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.total_connection_limit", "250"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.initial_connections", "1"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.sharing_conversations", "1"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.permit_insecure_servers", "false"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ss_lcipher", "none"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.convert_input", "true"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.auto_retry", "true"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.retry_interval", "10"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.retry_attempts", "6"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.long_retry_interval", "600"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.reporting_interval", "10"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.alternate_user", "true"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.outbound_sni", "Channel"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ocsp_check_extensions", "true"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.ocsp_authentication", "REQUIRED"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.cdp_check_extensions", "false"),
					resource.TestCheckResourceAttr("data.datapower_mqmanager.test", "result.0.client_revocation_checks", "REQUIRED"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
