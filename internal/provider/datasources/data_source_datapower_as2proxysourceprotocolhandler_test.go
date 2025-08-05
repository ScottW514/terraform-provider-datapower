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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceAS2ProxySourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AS2ProxySourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AS2ProxySourceProtocolHandler")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.AS2ProxySourceProtocolHandlerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.id", "AS2ProxySourceProtocolHandler_name"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.local_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.local_port", "80"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.http_version", "HTTP/1.1"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.persistent_connections", "true"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.allow_compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.max_url_len", "16384"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.max_total_hdr_len", "128000"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.max_hdr_count", "0"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.max_name_hdr_len", "0"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.max_value_hdr_len", "0"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.credential_charset", "protocol"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.remote_connection_timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.enable_passthrough", "true"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.enable_visibility_event", "true"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.enable_hmac_authentication", "true"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.ssl_server_config_type", "server"),
					resource.TestCheckResourceAttr("data.datapower_as2proxysourceprotocolhandler.test", "result.0.ssl_client_config_type", "client"),
				}...),
			},
		},
	})
}
