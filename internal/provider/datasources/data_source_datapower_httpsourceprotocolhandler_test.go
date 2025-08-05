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

func TestAccDataSourceHTTPSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_HTTPSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_HTTPSourceProtocolHandler")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.HTTPSourceProtocolHandlerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.id", "HTTPSourceProtocolHandler_name"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.local_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.local_port", "80"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.http_version", "HTTP/1.1"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.persistent_connections", "true"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.allow_compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.allow_web_socket_upgrade", "false"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.max_url_len", "16384"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.max_total_hdr_len", "128000"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.http2_max_streams", "100"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.http2_max_frame_size", "16384"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.http2_stream_header", "false"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.chunked_encoding", "true"),
					resource.TestCheckResourceAttr("data.datapower_httpsourceprotocolhandler.test", "result.0.header_timeout", "30000"),
				}...),
			},
		},
	})
}
