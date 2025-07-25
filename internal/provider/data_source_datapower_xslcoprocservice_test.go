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

func TestAccDataSourceXSLCoprocService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XSLCoprocService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XSLCoprocService")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.XSLCoprocServiceTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.id", "_name"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.priority", "normal"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.connection_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.intermediate_result_timeout", "20"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.cache_relative_url", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.use_client_uri_resolver", "true"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.crypto_extensions", "false"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.default_param_namespace", "http://www.datapower.com/param/config"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.debug_mode", "off"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.debug_history", "25"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.ssl_server_config_type", "server"),
					resource.TestCheckResourceAttr("data.datapower_xslcoprocservice.test", "result.0.local_address", "0.0.0.0"),
				}...),
			},
		},
	})
}
