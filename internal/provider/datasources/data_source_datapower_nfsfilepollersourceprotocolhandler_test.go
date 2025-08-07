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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceNFSFilePollerSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_NFSFilePollerSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_NFSFilePollerSourceProtocolHandler")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.NFSFilePollerSourceProtocolHandlerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.id", "AccTest_NFSFilePollerSourceProtocolHandler"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.delay_between_polls", "60000"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.delete_on_success", "false"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.success_rename_pattern", "$1.processed.ok"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.delete_on_error", "false"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.error_rename_pattern", "$0.processed.error"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.generate_result_file", "false"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.processing_seize_timeout", "0"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.xml_manager", "default"),
					resource.TestCheckResourceAttr("data.datapower_nfsfilepollersourceprotocolhandler.test", "result.0.max_transfers_per_poll", "0"),
				}...),
			},
		},
	})
}
