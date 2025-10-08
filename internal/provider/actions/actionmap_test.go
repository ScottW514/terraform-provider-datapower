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

package actions_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccActionsAAAPolicyFlushCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_AAAPolicy"
      target_domain = "acceptance_test"
      target_type   = "datapower_aaa_policy"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_cache"
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

func TestAccActionsAMQPSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_AMQPSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_amqp_source_protocol_handler"
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

func TestAccActionsAPICollectionFlushCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_APICollection"
      target_domain = "acceptance_test"
      target_type   = "datapower_api_collection"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_cache"
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

func TestAccActionsAPIGatewayFlushStylesheetCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_APIGateway"
      target_domain = "acceptance_test"
      target_type   = "datapower_api_gateway"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_stylesheet_cache"
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
func TestAccActionsAPIGatewayFlushDocumentCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_APIGateway"
      target_domain = "acceptance_test"
      target_type   = "datapower_api_gateway"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_document_cache"
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

func TestAccActionsAS1PollerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_AS1PollerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_as1_poller_source_protocol_handler"
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

func TestAccActionsAS2SourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_AS2SourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_as2_source_protocol_handler"
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

func TestAccActionsAS3SourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_AS3SourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_as3_source_protocol_handler"
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

func TestAccActionsB2BGatewayQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_B2BGateway"
      target_domain = "acceptance_test"
      target_type   = "datapower_b2b_gateway"
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

func TestAccActionsEBMS2SourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_EBMS2SourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_ebms2_source_protocol_handler"
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

func TestAccActionsEBMS3SourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_EBMS3SourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_ebms3_source_protocol_handler"
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

func TestAccActionsFTPFilePollerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_FTPFilePollerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_ftp_file_poller_source_protocol_handler"
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

func TestAccActionsFTPServerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_FTPServerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_ftp_server_source_protocol_handler"
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

func TestAccActionsHTTPSSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_HTTPSSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_https_source_protocol_handler"
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

func TestAccActionsHTTPSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_HTTPSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_http_source_protocol_handler"
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

func TestAccActionsMQv9PlusMFTSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_MQv9PlusMFTSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_mqv9_plus_mft_source_protocol_handler"
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

func TestAccActionsMQv9PlusSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_MQv9PlusSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_mqv9_plus_source_protocol_handler"
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

func TestAccActionsMultiProtocolGatewayQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_MultiProtocolGateway"
      target_domain = "acceptance_test"
      target_type   = "datapower_multi_protocol_gateway"
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

func TestAccActionsNFSFilePollerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_NFSFilePollerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_nfs_file_poller_source_protocol_handler"
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

func TestAccActionsPOPPollerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_POPPollerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_pop_poller_source_protocol_handler"
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

func TestAccActionsRBMSettingsFlushCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_RBMSettings"
      target_domain = "acceptance_test"
      target_type   = "datapower_rbm_settings"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_cache"
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

func TestAccActionsSFTPFilePollerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_SFTPFilePollerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_sftp_file_poller_source_protocol_handler"
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

func TestAccActionsSSHServerSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_SSHServerSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_ssh_server_source_protocol_handler"
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

func TestAccActionsSSLProxyServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_SSLProxyService"
      target_domain = "acceptance_test"
      target_type   = "datapower_ssl_proxy_service"
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

func TestAccActionsStatelessTCPSourceProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_StatelessTCPSourceProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_stateless_tcp_source_protocol_handler"
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

func TestAccActionsTCPProxyServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_TCPProxyService"
      target_domain = "acceptance_test"
      target_type   = "datapower_tcp_proxy_service"
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

func TestAccActionsWSGatewayQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_WSGateway"
      target_domain = "acceptance_test"
      target_type   = "datapower_ws_gateway"
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

func TestAccActionsWebAppFWQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_WebAppFW"
      target_domain = "acceptance_test"
      target_type   = "datapower_web_app_fw"
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

func TestAccActionsWebTokenServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_WebTokenService"
      target_domain = "acceptance_test"
      target_type   = "datapower_web_token_service"
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

func TestAccActionsXMLFirewallServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XMLFirewallService"
      target_domain = "acceptance_test"
      target_type   = "datapower_xml_firewall_service"
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

func TestAccActionsXMLManagerFlushStylesheetCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XMLManager"
      target_domain = "acceptance_test"
      target_type   = "datapower_xml_manager"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_stylesheet_cache"
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
func TestAccActionsXMLManagerFlushDocumentCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XMLManager"
      target_domain = "acceptance_test"
      target_type   = "datapower_xml_manager"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_document_cache"
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
func TestAccActionsXMLManagerFlushLdapPoolCache(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XMLManager"
      target_domain = "acceptance_test"
      target_type   = "datapower_xml_manager"
      on_create     = true
      on_update     = false
      on_delete     = false
      action        = "flush_ldap_pool_cache"
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

func TestAccActionsXSLCoprocServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XSLCoprocService"
      target_domain = "acceptance_test"
      target_type   = "datapower_xsl_coproc_service"
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

func TestAccActionsXSLProxyServiceQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XSLProxyService"
      target_domain = "acceptance_test"
      target_type   = "datapower_xsl_proxy_service"
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

func TestAccActionsXTCProtocolHandlerQuiesce(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Actions") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Actions")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `resource "datapower_log_label" "action_test" {
  id         = "ActionTest_LogLabel"
  app_domain = "acceptance_test"
  dependency_actions = [
    {
      target_id     = "AccTest_XTCProtocolHandler"
      target_domain = "acceptance_test"
      target_type   = "datapower_xtc_protocol_handler"
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
