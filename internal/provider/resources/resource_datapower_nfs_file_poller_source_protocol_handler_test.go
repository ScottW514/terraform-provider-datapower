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

func TestAccResourceNFSFilePollerSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_NFSFilePollerSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_NFSFilePollerSourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.NFSFilePollerSourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "delay_between_polls", "60000"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "delete_on_success", "false"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "success_rename_pattern", "$1.processed.ok"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "delete_on_error", "false"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "error_rename_pattern", "$0.processed.error"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "generate_result_file", "false"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "processing_seize_timeout", "0"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "xml_manager", "default"),
			resource.TestCheckResourceAttr("datapower_nfs_file_poller_source_protocol_handler.test", "max_transfers_per_poll", "0"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
