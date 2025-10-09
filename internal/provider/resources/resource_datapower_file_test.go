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

package resources_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceFile(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_File") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_File")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.FileTestConfig.GetResourceConfig(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: `
	resource "datapower_file" "content" {
	  app_domain = "acceptance_test"
	  remote_path = "local:///somepath/another/pluseonemore/test_content_file.txt"
	  content = "Test File"
	}
			`,
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: `
	resource "datapower_file" "content" {
	  app_domain = "acceptance_test"
	  remote_path = "local:///newpath/test_content_file.txt"
	  content = "Test File"
	}
			`,
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: `
	resource "datapower_file" "content" {
	  app_domain = "acceptance_test"
	  remote_path = "local:///newpath/test_content_file.txt"
	  content = "Test File2"
	}
			`,
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
