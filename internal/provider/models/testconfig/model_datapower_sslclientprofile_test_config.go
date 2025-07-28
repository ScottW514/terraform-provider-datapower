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

package testconfig

var SSLClientProfileTestConfig = ModelTestConfig{
	Name: "SSLClientProfile",
	Resource: `
resource "datapower_sslclientprofile" "test" {
  id = "SSLClientProfile_name"
  app_domain = "acc_test_domain"
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
}`,
	Data: `
data "datapower_sslclientprofile" "test" {
  depends_on = [ datapower_sslclientprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
