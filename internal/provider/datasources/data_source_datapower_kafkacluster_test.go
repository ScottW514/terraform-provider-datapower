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

func TestAccDataSourceKafkaCluster(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_KafkaCluster") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_KafkaCluster")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testutils.KafkaClusterTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.id", "AccTest_KafkaCluster"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.protocol", "plaintext"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.sasl_mechanism", "plain"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.autocommit", "true"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.memory_threshold", "268435456"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.maximum_message_size", "1048576"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.auto_retry", "true"),
					resource.TestCheckResourceAttr("data.datapower_kafkacluster.test", "result.0.retry_interval", "10"),
				}...),
			},
		},
	})
}
