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

type UserWOList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &UserDataSource{}
	_ datasource.DataSourceWithConfigure = &UserDataSource{}
)

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

type UserDataSource struct {
	pData *tfutils.ProviderData
}

func (d *UserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create or edit a local user account. When you modify a local user account, the changes do not affect remote user accounts in a role-based management XML information file. <p>An administrator can change the password for the account. An administrator can force the account owner to change the password during the next log in attempt.</p>",
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
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"access_level": schema.StringAttribute{
							MarkdownDescription: "Access level",
							Computed:            true,
						},
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Specify the user group for the user account. Each user group has an access profile to manage access rights. A user account inherits access rights from its user group.",
							Computed:            true,
						},
						"snmp_creds": schema.ListNestedAttribute{
							MarkdownDescription: "Specify SNMP V3 credentials for the user account. SNMP V3 users are a type of user account with SNMP V3 credentials. This type of user account creates an account and adds SNMP V3 credentials. Each account can have multiple SNMP V3 credentials, one for each SNMP V3 engine that is identified by an engine ID value. The current implementation supports an SNMP V3 credential for the local engine ID only. Therefore, define only one SNMP V3 credential for each account. <p>SNMP V3 requests can use authentication with the <tt>AuthKey</tt> and can use privacy (encryption and decryption) with the <tt>PrivKey</tt> . The use of these keys is between the user and the local SNMP engine. <ul><li>The <tt>AuthKey</tt> provides data integrity and data origin authentication.</li><li>The <tt>PrivKey</tt> provides data encryption and decryption.</li></ul></p><p>The <tt>AuthKey</tt> and the <tt>PrivKey</tt> can be explicit values or be generated by the DataPower Gateway. Specifying explicit values is useful when the key is created on another system. <ul><li>When the authentication protocol is MD5, enter the hex representation of the 16-byte key.</li><li>When the authentication protocol is SHA, enter the hex representation of the 20-byte key.</li><li>When the DataPower Gateway generates and stores the appropriate length key, enter a plaintext string that is at least 8 characters long as the hash.</li></ul></p>",
							NestedObject:        models.GetDmSnmpCredDataSourceSchema(),
							Computed:            true,
						},
						"hashed_snmp_creds": schema.ListNestedAttribute{
							MarkdownDescription: "",
							NestedObject:        models.GetDmSnmpCredMaskedDataSourceSchema(),
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

func (d *UserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data UserWOList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.UserWO{}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.UserWO{}
	if resFound {
		if value := res.Get(`User`); value.Exists() {
			for _, v := range value.Array() {
				item := models.UserWO{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.UserObjectTypeWO}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.UserObjectTypeWO})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
