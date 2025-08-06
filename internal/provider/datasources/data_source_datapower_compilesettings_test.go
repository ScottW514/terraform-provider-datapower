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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

func TestAccDataSourceCompileSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_CompileSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_CompileSettings")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.CompileSettingsTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.id", "CompileSettings_name"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.xslt_version", "XSLT10"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.strict", "true"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.profile", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.debug", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.stream", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.try_stream", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.minimum_escaping", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.stack_size", "1048576"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsi_validation", "warn"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsdl_validate_body", "strict"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsdl_validate_headers", "lax"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsdl_validate_faults", "strict"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsdl_wrapped_faults", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.allow_soap_enc_array", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.validate_soap_enc_array", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wildcards_ignore_xsi_type", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.wsdl_strict_soap_version", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.xacml_debug", "false"),
					resource.TestCheckResourceAttr("data.datapower_compilesettings.test", "result.0.allow_xop_include", "false"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
