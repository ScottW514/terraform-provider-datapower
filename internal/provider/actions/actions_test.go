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

package actions_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccActionsDomainQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_domain" "action_test" {
  app_domain = "test_domain"
}
resource "datapower_file" "action_test" {
  app_domain = datapower_domain.action_test.app_domain
  remote_path = "local:///actions_test_file.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
  dependency_actions = [
    {
      target_domain = datapower_domain.action_test.app_domain
      target_type   = "datapower_domain"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "quiesce"
    }
  ]
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func TestAccActionsDomainRestart(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_domain" "action_test" {
  app_domain = "action_test"
}
resource "datapower_file" "action_test" {
  app_domain = datapower_domain.action_test.app_domain
  remote_path = "local:///actions_test_file.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
  dependency_actions = [
    {
      target_domain = "action_test"
      target_type   = "datapower_domain"
      on_create     = true
      on_update     = false
      on_delete     = true
      action        = "restart"
    }
  ]
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func TestAccActionsMultipleHTTPSSSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_file" "action_test" {
  app_domain = "acceptance_test"
  remote_path = "local:///actions_test_file.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
  dependency_actions = [
    {
      target_id     = "AccTest_HTTPSSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_httpssourceprotocolhandler"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "quiesce"
    }
  ]
}
resource "datapower_file" "action_test2" {
  app_domain = "acceptance_test"
  remote_path = "local:///actions_test_file2.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file2.txt"
  dependency_actions = [
    {
      target_id     = "AccTest_HTTPSSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_httpssourceprotocolhandler"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "quiesce"
    }
  ]
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
