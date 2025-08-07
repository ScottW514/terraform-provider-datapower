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

func TestAccDataSourceLoadBalancerGroup(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_LoadBalancerGroup") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_LoadBalancerGroup")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.LoadBalancerGroupTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.id", "LoadBalancerGroup_name"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.algorithm", "round-robin"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.retrieve_info", "false"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.wlm_retrieval", "use-websphere"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.wlm_transport", "http"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.damp", "120"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.never_return_sick_member", "false"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.try_every_server_before_failing", "false"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.masquerade_member", "false"),
					resource.TestCheckResourceAttr("data.datapower_loadbalancergroup.test", "result.0.application_routing", "false"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
