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

func TestAccResourceCompileSettings(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_CompileSettings") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_CompileSettings")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.CompileSettingsTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "xslt_version", "XSLT10"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "strict", "true"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "profile", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "debug", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "stream", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "try_stream", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "minimum_escaping", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "stack_size", "1048576"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsi_validation", "warn"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsdl_validate_body", "strict"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsdl_validate_headers", "lax"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsdl_validate_faults", "strict"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsdl_wrapped_faults", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "allow_soap_enc_array", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "validate_soap_enc_array", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wildcards_ignore_xsi_type", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "wsdl_strict_soap_version", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "xacml_debug", "false"),
			resource.TestCheckResourceAttr("datapower_compilesettings.test", "allow_xop_include", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
