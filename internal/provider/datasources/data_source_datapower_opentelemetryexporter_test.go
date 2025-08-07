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

func TestAccDataSourceOpenTelemetryExporter(t *testing.T) {
	if os.Getenv("DP_ACC_ALL") == "" && os.Getenv("DP_ACC_OpenTelemetryExporter") == "" {
		t.Skip("skipping test, set environment variable DP_ACC_ALL DP_ACC_OpenTelemetryExporter")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testutils.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: testutils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testconfig.OpenTelemetryExporterTestConfig.GetDataConfig(),
				Check: resource.ComposeTestCheckFunc([]resource.TestCheckFunc{
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.id", "OpenTelemetryExporter_name"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.type", "http"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.traces_path", "/v1/traces"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.port", "4318"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.http_content_type", "json"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.timeout", "10"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.processor", "batch"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.max_queue_size", "2048"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.max_export_size", "512"),
					resource.TestCheckResourceAttr("data.datapower_opentelemetryexporter.test", "result.0.export_delay_interval", "5000"),
				}...),
			},
		},
	})
	actions.PostProcess()
}
