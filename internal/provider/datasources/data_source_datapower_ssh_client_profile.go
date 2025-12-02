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

package datasources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type SSHClientProfileList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &SSHClientProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SSHClientProfileDataSource{}
)

func NewSSHClientProfileDataSource() datasource.DataSource {
	return &SSHClientProfileDataSource{}
}

type SSHClientProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *SSHClientProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_client_profile"
}

func (d *SSHClientProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An SSH client profile defines the authentication type, credentials, and cipher suites to use for an SSH client connection.",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"user_name": schema.StringAttribute{
							MarkdownDescription: "User",
							Computed:            true,
						},
						"profile_usage": schema.StringAttribute{
							MarkdownDescription: "Specify the usage of the profile. Only SFTP is supported.",
							Computed:            true,
						},
						"ssh_user_authentication": models.GetDmSSHUserAuthenticationMethodsDataSourceSchema("Specify the available types of SSH user authentication for the SSH client. Authentication can be public key or password or both public key and password. When both methods are defined, public key authentication is attempted first.", "user-auth", ""),
						"user_private_key": schema.StringAttribute{
							MarkdownDescription: "Specify the private key for public key authentication. User private keys must not be password protected.",
							Computed:            true,
						},
						"password_alias": schema.StringAttribute{
							MarkdownDescription: "Password Alias",
							Computed:            true,
						},
						"persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to support persistent connections. By default, persistent connections are enabled. <ul><li>When enabled, new requests reuse the connection of a previous session without reauthentication.</li><li>When not enabled, new request must reauthenticate.</li></ul>",
							Computed:            true,
						},
						"persistent_connection_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the idle duration in seconds for a persistent connection. When the connection remains idle for the specified duration, the connection is closed. Enter any value in the range 1 - 86000. The default value is 120.",
							Computed:            true,
						},
						"strict_host_key_checking": schema.BoolAttribute{
							MarkdownDescription: "Specify how to check host keys during the connection and authentication phases. By default strict host key checking is not enabled. <ul><li>When enabled, checks the host key against the known hosts list. Host keys that are not in the known host list are rejected.</li><li>When not enabled, checks the host key against the known hosts list. Host keys that are not in the known host list are added to the known hosts list and accepted.</li></ul>",
							Computed:            true,
						},
						"ciphers": schema.ListAttribute{
							MarkdownDescription: "Specify the SSH cipher suites to support.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"kex_alg": schema.ListAttribute{
							MarkdownDescription: "Specify the key exchange (KEX) algorithms to support.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"mac_alg": schema.ListAttribute{
							MarkdownDescription: "Specify the message authentication codes (MAC) to support.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SSHClientProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *SSHClientProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data SSHClientProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.SSHClientProfile{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.SSHClientProfile{}
	if resFound {
		if value := res.Get(`SSHClientProfile`); value.Exists() {
			for _, v := range value.Array() {
				item := models.SSHClientProfile{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.SSHClientProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.SSHClientProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
