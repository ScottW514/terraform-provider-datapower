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

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
)

func TestAccDataSourceDomain(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Domain") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Domain")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.DomainTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_domain.test", "result.0.config_mode", "local"),
					resource.TestCheckResourceAttr("data.datapower_domain.test", "result.0.import_format", "ZIP"),
					resource.TestCheckResourceAttr("data.datapower_domain.test", "result.0.local_ip_rewrite", "true"),
					resource.TestCheckResourceAttr("data.datapower_domain.test", "result.0.max_chkpoints", "3"),
					resource.TestCheckResourceAttr("data.datapower_domain.test", "result.0.config_permissions_mode", "scope-domain"),
				}...),
			},
		},
	})
}
