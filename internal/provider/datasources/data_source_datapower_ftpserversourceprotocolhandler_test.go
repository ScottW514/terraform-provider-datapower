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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceFTPServerSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_FTPServerSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_FTPServerSourceProtocolHandler")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.FTPServerSourceProtocolHandlerTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.id", "FTPServerSourceProtocolHandler_name"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.local_address", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.local_port", "21"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.filesystem_type", "virtual-ephemeral"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.persistent_filesystem_timeout", "600"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.default_directory", "/"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.max_filename_length", "256"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.require_tls", "off"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_ccc", "true"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.passive", "allow"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.use_pasv_port_range", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.pasv_min_port", "1024"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.pasv_max_port", "1050"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.pasv_idle_time_out", "60"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.disable_pasvip_check", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.disable_portip_check", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.use_alternate_pasv_addr", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_list_cmd", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_dele_cmd", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.data_encryption", "allow"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_compression", "true"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_stou", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.allow_rest", "false"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.restart_timeout", "240"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.idle_timeout", "0"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.response_type", "none"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.response_storage", "temporary"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.temporary_storage_size", "32"),
					resource.TestCheckResourceAttr("data.datapower_ftpserversourceprotocolhandler.test", "result.0.ssl_server_config_type", "server"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
