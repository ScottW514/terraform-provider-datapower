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

func TestAccDataSourceMQv9PlusMFTSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_MQv9PlusMFTSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_MQv9PlusMFTSourceProtocolHandler")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.MQv9PlusMFTSourceProtocolHandlerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.id", "MQv9PlusMFTSourceProtocolHandler_name"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.get_message_options", "32769"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.concurrent_connections", "1"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.polling_interval", "30"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.retrieve_backout_settings", "false"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.ignore_backout_errors", "false"),
					resource.TestCheckResourceAttr("data.datapower_mqv9plusmftsourceprotocolhandler.test", "result.0.use_qm_name_in_url", "false"),
				}...),
			},
		},
	})
}
