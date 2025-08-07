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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceAssemblyActionInvoke(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AssemblyActionInvoke") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AssemblyActionInvoke")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.AssemblyActionInvokeTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "timeout", "60"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "method", "Keep"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "backend_type", "detect"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "compression", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "cache_type", "Protocol"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "time_to_live", "900"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "cache_unsafe_response", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "follow_redirects", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "http_version", "HTTP/1.1"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "http2_required", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "do_chunked_upload", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "persistent_connection", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "stop_on_error", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "output", "message"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "decode_request_params", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "encode_plus_char", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "keep_payload", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "inject_user_agent_header", "true"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "inject_proxy_headers", "false"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "header_control_list", "default-accept-all"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "parameter_control_list", "default-reject-all"),
			resource.TestCheckResourceAttr("datapower_assemblyactioninvoke.test", "action_debug", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
