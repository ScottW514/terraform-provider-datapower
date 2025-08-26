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

package validators_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/scottw514/terraform-provider-datapower/testutils"
)

// logical-not, logical-and, property-value-in-list
func TestAccValAAAJWTValidator(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Validators") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Validators")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_aaa_jwt_validator" "test" {
  id = "ValTestAAAJWTValidator"
  app_domain = "acceptance_test"
  username_claim = "sub"
  val_method = {"decrypt" = true}
  decrypt_credential_type = "jwk-remote"
  decrypt_fetch_cred_ssl_profile = "AccTest_SSLClientProfile"
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// logical-and, logical-or, property-value-in-list, property-value-not-in-list
func TestAccValOAuthSupportedClient(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Validators") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Validators")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_oauth_supported_client" "test" {
  id = "ValTestOAuthSupportedClient"
  app_domain = "acceptance_test"
  oauth_role = {"azsvr": true}
  az_grant = {"code": true}
  generate_client_secret = false
  client_secret_wo = "secret"
  client_secret_wo_version = 1
  redirect_uri = ["^https://example.com/redirect$"]
  scope = "*"
  token_secret = "AccTest_CryptoSSKey"
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// Empty List
func TestAccValB2BGateway(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_Validators") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_Validators")
	}

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: `
resource "datapower_b2b_gateway" "test" {
  id = "ValTestB2BGateway"
  app_domain = "acceptance_test"
  b2b_profiles = [{partner_profile="AccTest_B2BProfile"}]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode = "purgeonly"
  xml_manager = "default"
}
`,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
