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

func TestAccDataSourceAssemblyActionInvoke(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AssemblyActionInvoke") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AssemblyActionInvoke")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.AssemblyActionInvokeTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.id", "_name"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.timeout", "60"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.method", "Keep"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.backend_type", "detect"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.compression", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.cache_type", "Protocol"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.time_to_live", "900"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.cache_unsafe_response", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.follow_redirects", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.http_version", "HTTP/1.1"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.http2_required", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.do_chunked_upload", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.persistent_connection", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.stop_on_error", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.output", "message"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.decode_request_params", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.encode_plus_char", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.keep_payload", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.inject_user_agent_header", "true"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.inject_proxy_headers", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.header_control_list", "default-accept-all"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.parameter_control_list", "default-reject-all"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactioninvoke.test", "result.0.action_debug", "false"),
				}...),
			},
		},
	})
}
