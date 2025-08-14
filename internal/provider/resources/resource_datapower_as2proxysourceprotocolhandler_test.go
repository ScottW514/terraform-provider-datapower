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
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceAS2ProxySourceProtocolHandler(t *testing.T) {
	t.Skip("skipping test - configured in definition")
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AS2ProxySourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AS2ProxySourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.AS2ProxySourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "local_port", "80"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "allow_compression", "false"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "max_url_len", "16384"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "max_total_hdr_len", "128000"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "max_hdr_count", "0"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "max_name_hdr_len", "0"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "max_value_hdr_len", "0"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "credential_charset", "protocol"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "remote_connection_timeout", "60"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "enable_passthrough", "true"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "enable_visibility_event", "true"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "enable_hmac_authentication", "true"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "ssl_server_config_type", "server"),
			resource.TestCheckResourceAttr("datapower_as2proxysourceprotocolhandler.test", "ssl_client_config_type", "client"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
