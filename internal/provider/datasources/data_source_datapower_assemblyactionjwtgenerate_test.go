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
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceAssemblyActionJWTGenerate(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_AssemblyActionJWTGenerate") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_AssemblyActionJWTGenerate")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.AssemblyActionJWTGenerateTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.id", "AccTest_AssemblyActionJWTGenerate"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.jwt", "generated.jwt"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.jwtid_claims", "false"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.issuer_claim", "iss.claim"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.validity_period", "3600"),
					resource.TestCheckResourceAttr("data.datapower_assemblyactionjwtgenerate.test", "result.0.action_debug", "false"),
				}...),
			},
		},
	})
}
