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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type TAMList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &TAMDataSource{}
	_ datasource.DataSourceWithConfigure = &TAMDataSource{}
)

func NewTAMDataSource() datasource.DataSource {
	return &TAMDataSource{}
}

type TAMDataSource struct {
	pData *tfutils.ProviderData
}

func (d *TAMDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tam"
}

func (d *TAMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>A DataPower Gateway can support multiple IBM Security Access Manager registry types: LDAP or Microsoft Active Directory. Support for the registry type is established for each Access Manager client that starts according to its configuration.</p>",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
							MarkdownDescription: "Enter the descriptive summary for the configuration.",
							Computed:            true,
						},
						"ad_use_ad": schema.BoolAttribute{
							MarkdownDescription: "<p>Select whether the Access Manager client uses Microsoft Active Directory instead of LDAP as the registry. The default registry for an Access Manager deployment is LDAP. This selection will cause this client to use Microsoft Active Directory. Active Directory type is not supported after ISAM 7.0 .</p><p><b>Note:</b> The type of registry that an Access Manager deployment supports is determined by the configuration of the Access Manager server. The registry that you define in this configuration is for a client and must match the registry of the server.</p>",
							Computed:            true,
						},
						"tam_version": schema.StringAttribute{
							MarkdownDescription: "<p>Select the Access Manager client version to use. The default value is Default.</p>",
							Computed:            true,
						},
						"configuration_file": schema.StringAttribute{
							MarkdownDescription: "Select the location of the configuration file for the Access Manager client. To be available for selection, files must have .conf or .cfg as their file extension.",
							Computed:            true,
						},
						"ad_configuration_file": schema.StringAttribute{
							MarkdownDescription: "Select the location of the configuration file for user directories. To be available for selection, files must have .conf or .cfg as their file extension.",
							Computed:            true,
						},
						"ssl_key_file": schema.StringAttribute{
							MarkdownDescription: "Select the location of the key file for TLS communication. To be available for selection, files must have .kdb as their file extension. Generally, these files are in the cert: directory or the sharedcert: directory.",
							Computed:            true,
						},
						"ssl_key_stash_file": schema.StringAttribute{
							MarkdownDescription: "Select the location of the key stash file for TLS communication. To be available for selection, files must have .sth as their file extension. Generally, these files are in the cert: directory or the sharedcert: directory.",
							Computed:            true,
						},
						"use_local_mode": schema.BoolAttribute{
							MarkdownDescription: "Select whether to create a local copy of the policy database. Set this property to cache the policy database locally instead of accessing the remote policy server. This property must match the behavior defined in the configuration files for the Access Manager client.",
							Computed:            true,
						},
						"poll_interval": schema.StringAttribute{
							MarkdownDescription: "Enter the interval between requests to update the local policy database from the remote policy server. <ul><li><b>default</b> - Uses the default value, which is 600 seconds.</li><li><b>disable</b> - Disables requests to the policy database for requests.</li><li><i>seconds</i> - Specifies the time interval in seconds.</li></ul>",
							Computed:            true,
						},
						"listen_mode": schema.BoolAttribute{
							MarkdownDescription: "Select whether to accept notifications to update the local policy database from the policy server. When you set this property, it overrides the behavior defined in configuration files for the Access Manager client.",
							Computed:            true,
						},
						"listen_port": schema.Int64Attribute{
							MarkdownDescription: "Enter the listening port on the DataPower Gateway to receive update notifications from the remote policy server.",
							Computed:            true,
						},
						"returning_user_attributes": schema.BoolAttribute{
							MarkdownDescription: "Select whether the registry returns users attributes for successful authorization requests. When set, the registry returns user attributes.",
							Computed:            true,
						},
						"ldap_use_ssl": schema.BoolAttribute{
							MarkdownDescription: "Select whether to use TLS communication between the Access Manager client and the LDAP or Active Directory server.",
							Computed:            true,
						},
						"ldapssl_port": schema.Int64Attribute{
							MarkdownDescription: "Enter the listening port that the LDAP server uses for TLS communication. This property does not apply for TLS communication with an Active Directory server.",
							Computed:            true,
						},
						"ldapssl_key_file": schema.StringAttribute{
							MarkdownDescription: "<p>Select the location of the key file that contains the certificates for TLS communication with the registry server.</p><ul><li>For server-only authentication, the key file must contain the signer certificate for the registry server.</li><li>For mutual authentication, the key file must also contain a personal certificate that the registry server can validate. If the personal certificate is not the default personal certificate in the key file, you must enter the label of the personal certificate.</li></ul><p>This file must be in the cert: or the sharedcert: directory.</p>",
							Computed:            true,
						},
						"ldapssl_key_file_password_alias": schema.StringAttribute{
							MarkdownDescription: "Enter the password alias of the password for the key file that contains the certificates for TLS communication with the registry server.",
							Computed:            true,
						},
						"ldapssl_key_file_label": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the label of the personal certificate in the key file for client authentication.</p><ul><li>When using mutual authentication with the registry server and the personal certificate is not the default personal certificate in the key file, enter the label of the personal certificate. The personal certificate allows client authentication.</li><li>For server-only authentication, do not enter a value.</li></ul>",
							Computed:            true,
						},
						"tam_use_fips": schema.BoolAttribute{
							MarkdownDescription: "Select whether the secure communication between the Access Manager client and the authorization server runs in FIPS mode.",
							Computed:            true,
						},
						"tam_choose_nist": schema.StringAttribute{
							MarkdownDescription: "In FIPS mode, there is a mandatory NIST compliance level to select. Select a level that is compatible with the corresponding Access Manager servers and registry servers. Note that the NIST and NSA options are available only in Access Manager versions 7.0 and later.",
							Computed:            true,
						},
						"tam_use_basic_user": schema.BoolAttribute{
							MarkdownDescription: "Control whether to support basic users in the client. When enabled, you can use LDAP user entries for authentication or authorization without importing them into the ISAM domain.",
							Computed:            true,
						},
						"user_principal_attribute": schema.StringAttribute{
							MarkdownDescription: "Specify the attribute that identifies the basic user in the LDAP user entry. The default value is uid.",
							Computed:            true,
						},
						"user_no_duplicates": schema.BoolAttribute{
							MarkdownDescription: "<p>Control whether to allow duplicate principals.</p><ul><li>When enabled, the search for basic users covers all suffixes to ensure that no users with the same name are found. If duplicate principals are found in this client, the system returns an error.</li><li>When disabled, the search for basic users ignores possible duplicates. By default, duplicate principals are not allowed.</li></ul>",
							Computed:            true,
						},
						"user_search_suffixes": schema.ListAttribute{
							MarkdownDescription: "<p>Specify the ordered list of LDAP suffixes to be searched for principals. When specified and suffix optimization is disabled, the suffixes are searched in entry order. If suffix optimization is enabled, this order is overridden by the suffix optimization order.</p><p>If you do not specify any suffixes, the system searches all available suffixes.</p>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"user_suffix_optimiser": schema.BoolAttribute{
							MarkdownDescription: "<p>Control whether to search the suffixes in an optimized order.</p><ul><li>When enabled and uplicate principals are allowed, the suffixes are searched in an optimized order based on hit count, with the most hit of the suffix at the head of the search suffix list. This can help reduce the number of suffixes searched. If duplicate principals are not allowed, the suffix optimization setting is disregarded and all suffixes are searched to check for duplicates.</li><li>When disabled, the search order is provided by the order that is defined by the search suffixes property.</li></ul>",
							Computed:            true,
						},
						"tam_fed_dirs": schema.ListNestedAttribute{
							MarkdownDescription: "Specify a list of federated directories. Each entry describes a unique set of LDAP suffixes and LDAP server. Federated directories define all the suffixes that can be searched for user identities.",
							NestedObject:        models.GetDmTAMFedDirDataSourceSchema(),
							Computed:            true,
						},
						"tamaz_replicas": schema.ListNestedAttribute{
							MarkdownDescription: "<p>Replicas indicate the network location of remote authorization servers. You must configure at least one replica. You can configure additional replicas for failover purposes.</p><p><b>Note:</b> If you uploaded a file that was created previously, it must define at least one replica.</p>",
							NestedObject:        models.GetDmTAMAZReplicaDataSourceSchema(),
							Computed:            true,
						},
						"tamras_trace": models.GetDmTAMRASTraceDataSourceSchema("<p>Trace logging is a useful debugging tool. By default, trace logging is not enabled. Trace logging collects large amounts of data in a short amount of time and might result in a significant performance degradation. Enable trace logging only at the direction of IBM Support.</p><p>When enabled, the DataPower Gateway creates two trace files for each library. The DataPower Gateway writes the files cyclically. Double the size of the files to obtain the total allowable file size.</p>", "tam-ras-trace", ""),
						"auto_retry": schema.BoolAttribute{
							MarkdownDescription: "<p>Select whether to attempt starting the Access Manager client after an initial failure. The DataPower Gateway automatically attempts to start the client after a critical error. This property controls the behavior after the initial attempt to restart.</p><ul><li>When enabled, the DataPower Gateway attempts to start the client with the defined configuration.</li><li>When disabled, the client is marked as <tt>down</tt> .</li></ul><p>The default behavior is to not attempt to start the client after an initial failure.</p>",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "<p>Enter the number of seconds to wait between attempts to start the Access Manager client. Enter a value in the range 1 - 65535. The default value is 180.</p>",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Enter the number of attempts to perform for the Access Manager client. After performing these attempts and the client did not start, each additional attempt waits the number of seconds defined by the long interval. Enter a value in the range 0 - 65535. A value of 0 disables the long interval. The default value is 3.",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Specifies the number of seconds to wait after reaching the number of retry attempts. Enter a value in the range 1 - 65535. The default value is 900.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *TAMDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *TAMDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data TAMList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.TAM{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.TAM{}
	if value := res.Get(`TAM`); value.Exists() {
		for _, v := range value.Array() {
			item := models.TAM{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.TAMObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.TAMObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
