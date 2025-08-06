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

package resources_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceMQManager(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MQManager") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MQManager")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.MQManagerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ccsid", "819"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "channel_name", "SYSTEM.DEF.SVRCONN"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "heartbeat", "300"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "maximum_message_size", "1048576"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "cache_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ffst_size", "500"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ffst_rotate", "3"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "units_of_work", "0"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "automatic_backout", "false"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "total_connection_limit", "250"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "initial_connections", "1"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "sharing_conversations", "1"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "permit_insecure_servers", "false"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ss_lcipher", "none"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "convert_input", "true"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "auto_retry", "true"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "retry_interval", "10"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "retry_attempts", "6"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "long_retry_interval", "600"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "reporting_interval", "10"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "alternate_user", "true"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "outbound_sni", "Channel"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ocsp_check_extensions", "true"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "ocsp_authentication", "REQUIRED"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "cdp_check_extensions", "false"),
			resource.TestCheckResourceAttr("datapower_mqmanager.test", "client_revocation_checks", "REQUIRED"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
