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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceXSLCoprocService(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_XSLCoprocService") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_XSLCoprocService")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.XSLCoprocServiceTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "priority", "normal"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "connection_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "intermediate_result_timeout", "20"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "cache_relative_url", "false"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "use_client_uri_resolver", "true"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "crypto_extensions", "false"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "default_param_namespace", "http://www.datapower.com/param/config"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "debug_mode", "off"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "debug_history", "25"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "ssl_server_config_type", "server"),
			resource.TestCheckResourceAttr("datapower_xslcoprocservice.test", "local_address", "0.0.0.0"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
