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

func TestAccResourceAS2SourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AS2SourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AS2SourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.AS2SourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "local_port", "80"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "allow_compression", "false"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "max_url_len", "16384"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "max_total_hdr_len", "128000"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "max_hdr_count", "0"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "max_name_hdr_len", "0"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "max_value_hdr_len", "0"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "credential_charset", "protocol"),
			resource.TestCheckResourceAttr("datapower_as2sourceprotocolhandler.test", "ssl_server_config_type", "server"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
