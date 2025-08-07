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

func TestAccDataSourceNFSDynamicMounts(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_NFSDynamicMounts") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_NFSDynamicMounts")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.NFSDynamicMountsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "enabled", "false"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "version", "3"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "transport", "tcp"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "mount_type", "hard"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "read_only", "false"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "read_size", "4096"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "write_size", "4096"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "timeout", "7"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "retransmissions", "3"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "idle_unmount_seconds", "900"),
					resource.TestCheckResourceAttr("data.datapower_nfsdynamicmounts.test", "mount_timeout_seconds", "30"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
