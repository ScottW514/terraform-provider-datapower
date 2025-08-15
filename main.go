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

package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/scottw514/terraform-provider-datapower/internal/provider"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
)

//go:generate go run gen/generator.go

var (
	version string = "0.9.1"
)

func main() {
	p := provider.New(version)
	defer func() {
		err := actions.SaveDomains()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/ScottW514/datapower",
	}

	err := providerserver.Serve(context.Background(), p, opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
