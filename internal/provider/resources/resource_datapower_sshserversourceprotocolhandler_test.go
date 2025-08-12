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

func TestAccResourceSSHServerSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_SSHServerSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_SSHServerSourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.SSHServerSourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "local_port", "22"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_listings", "true"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_delete", "false"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_stat", "false"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_mkdir", "false"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_rmdir", "false"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "allow_backend_rename", "false"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "filesystem_type", "transparent"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "default_directory", "/"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "idle_timeout", "0"),
			resource.TestCheckResourceAttr("datapower_sshserversourceprotocolhandler.test", "persistent_filesystem_timeout", "600"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
