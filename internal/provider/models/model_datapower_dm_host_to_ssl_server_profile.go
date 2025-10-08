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

package models

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmHostToSSLServerProfile struct {
	HostNameWildmat types.String `tfsdk:"host_name_wildmat"`
	SslServer       types.String `tfsdk:"ssl_server"`
}

var DmHostToSSLServerProfileObjectType = map[string]attr.Type{
	"host_name_wildmat": types.StringType,
	"ssl_server":        types.StringType,
}
var DmHostToSSLServerProfileObjectDefault = map[string]attr.Value{
	"host_name_wildmat": types.StringNull(),
	"ssl_server":        types.StringNull(),
}

func GetDmHostToSSLServerProfileDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmHostToSSLServerProfileDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"host_name_wildmat": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the shell-style pattern to match against the hostname.",
				Computed:            true,
			},
			"ssl_server": DataSourceSchema.StringAttribute{
				MarkdownDescription: "Specify the TLS server profile to use when a hostname matches the pattern.",
				Computed:            true,
			},
		},
	}
	return DmHostToSSLServerProfileDataSourceSchema
}
func GetDmHostToSSLServerProfileResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmHostToSSLServerProfileResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"host_name_wildmat": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style pattern to match against the hostname.", "", "").String,
				Required:            true,
			},
			"ssl_server": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS server profile to use when a hostname matches the pattern.", "", "ssl_server_profile").String,
				Required:            true,
			},
		},
	}
	return DmHostToSSLServerProfileResourceSchema
}

func (data DmHostToSSLServerProfile) IsNull() bool {
	if !data.HostNameWildmat.IsNull() {
		return false
	}
	if !data.SslServer.IsNull() {
		return false
	}
	return true
}

func (data DmHostToSSLServerProfile) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.HostNameWildmat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostNameWildmat`, data.HostNameWildmat.ValueString())
	}
	if !data.SslServer.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLServer`, data.SslServer.ValueString())
	}
	return body
}

func (data *DmHostToSSLServerProfile) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HostNameWildmat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HostNameWildmat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostNameWildmat = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
}

func (data *DmHostToSSLServerProfile) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `HostNameWildmat`); value.Exists() && !data.HostNameWildmat.IsNull() {
		data.HostNameWildmat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostNameWildmat = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLServer`); value.Exists() && !data.SslServer.IsNull() {
		data.SslServer = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslServer = types.StringNull()
	}
}
