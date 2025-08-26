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

func TestAccResourceB2BProfile(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_B2BProfile") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_B2BProfile")
	}
	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testutils.B2BProfileTestConfig.GetResourceConfig(),
		Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "profile_type", "internal"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "response_type", "preprocessed"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "inbound_require_signed", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "inbound_require_encrypted", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "outbound_sign", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "outbound_sign_digest_alg", "sha1"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "as_allow_duplicate_message", "never"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "preserve_filename", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_inbound_send_receipt", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_inbound_send_signed_receipt", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_inbound_receipt_reply_pattern", "Response"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_inbound_require_signed", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_inbound_require_encrypted", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_outbound_sign", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_outbound_signature_alg", "dsa-sha1"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_outbound_signature_c14n_alg", "c14n"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_outbound_sign_digest_alg", "sha1"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_enable_cpa_binding", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_start_parameter", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_allow_duplicate_message", "never"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "mdn_ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_ack_ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_outbound_sign", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_outbound_sign_digest_alg", "sha1"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_outbound_signature_alg", "rsa-sha1"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_inbound_require_signed", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_inbound_require_encrypted", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_inbound_require_compressed", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_receipt_ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_notification", "false"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms_notification_ssl_client_config_type", "client"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_allow_duplicate_message", "never"),
			resource.TestCheckResourceAttr("datapower_b2b_profile.test", "ebms3_duplicate_detection_notification", "false"),
		}...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
