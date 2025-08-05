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

func TestAccResourceHTTPSSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_HTTPSSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_HTTPSSourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.HTTPSSourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "local_port", "443"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "persistent_connections", "true"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "allow_compression", "false"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "allow_web_socket_upgrade", "false"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "max_url_len", "16384"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "max_total_hdr_len", "128000"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "credential_charset", "protocol"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "ssl_server_config_type", "server"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "http2_max_streams", "100"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "http2_max_frame_size", "16384"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "http2_stream_header", "false"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "chunked_encoding", "true"),
			resource.TestCheckResourceAttr("datapower_httpssourceprotocolhandler.test", "header_timeout", "30000"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
