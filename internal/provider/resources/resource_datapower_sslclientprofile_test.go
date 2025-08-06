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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccResourceSSLClientProfile(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_SSLClientProfile") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_SSLClientProfile")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testconfig.SSLClientProfileTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "validate_server_cert", "true"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "caching", "true"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "cache_timeout", "300"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "cache_size", "100"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "use_custom_sni_hostname", "false"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "validate_hostname", "false"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "hostname_validation_fail_on_error", "false"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "enable_tls13_compat", "true"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "disable_renegotiation", "true"),
			resource.TestCheckResourceAttr("datapower_sslclientprofile.test", "require_closure_notification", "true"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
	actions.PostProcess()
}
