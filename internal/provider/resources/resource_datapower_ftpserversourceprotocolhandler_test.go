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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceFTPServerSourceProtocolHandler(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_FTPServerSourceProtocolHandler") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_FTPServerSourceProtocolHandler")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.FTPServerSourceProtocolHandlerTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "local_address", "0.0.0.0"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "local_port", "21"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "filesystem_type", "virtual-ephemeral"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "persistent_filesystem_timeout", "600"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "default_directory", "/"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "max_filename_length", "256"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "require_tls", "off"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_ccc", "true"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "passive", "allow"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "use_pasv_port_range", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "pasv_min_port", "1024"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "pasv_max_port", "1050"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "pasv_idle_time_out", "60"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "disable_pasvip_check", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "disable_portip_check", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "use_alternate_pasv_addr", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_list_cmd", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_dele_cmd", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "data_encryption", "allow"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_compression", "true"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_stou", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "allow_rest", "false"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "restart_timeout", "240"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "idle_timeout", "0"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "response_type", "none"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "response_storage", "temporary"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "temporary_storage_size", "32"),
			resource.TestCheckResourceAttr("datapower_ftpserversourceprotocolhandler.test", "ssl_server_config_type", "server"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
