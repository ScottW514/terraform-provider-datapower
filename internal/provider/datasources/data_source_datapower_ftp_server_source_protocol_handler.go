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

type FTPServerSourceProtocolHandlerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &FTPServerSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &FTPServerSourceProtocolHandlerDataSource{}
)

func NewFTPServerSourceProtocolHandlerDataSource() datasource.DataSource {
	return &FTPServerSourceProtocolHandlerDataSource{}
}

type FTPServerSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *FTPServerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftp_server_source_protocol_handler"
}

func (d *FTPServerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>The FTP server handler provides an FTP server that can be used to submit files for processing by the system. Each file that is written results in one transaction.</p><p>There can be multiple FTP servers, but only one server can listen on the default port 21 on a given IP address. There can be multiple simultaneous connections from FTP clients to the same FTP server.</p><p><b>Notes:</b></p><ul><li>The 226 FTP response code at the end of an FTP <tt>STOR</tt> or <tt>STOU</tt> command is conditional on successful completion of the internal steps and backside operation of the transaction.</li><li>Changes in the configuration affect only new connections to this FTP server. Existing connections continue to use their current configuration until they disconnect.</li></ul>",
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
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Specify the local address that the FTP server service listens. If the service is listening on a virtual IP address that is managed by a standby control group with the self-balancing feature enabled, incoming control and data connections always go to the active member of the standby control group. The default value is 0.0.0.0, which indicates that the service is active on all IP4v addresses.",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the local port that the FTP server service listens. This port is to establish FTP control connections. This port does not control the port for data connections. If the FTP client uses the <tt>PASV</tt> command, data connections use an arbitrary, unused port. Enter a value in the range 1 - 65535. The default value is 21.",
							Computed:            true,
						},
						"filesystem_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of file system that the FTP server shows. <ul><li>When virtual ephemeral or virtual persistent, the client can write files to all directories. These files are shown in directory listings but cannot be retrieved. For file system responses, the responses are shown and can be retrieved, renamed, and deleted by the client.</li><li>When transparent, the file system shows the contents of the equivalent path of the remote server.</li></ul><p>This setting is supported by only the FTP server handler and a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"persistent_filesystem_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to retain a connection to a virtual file system after all FTP control connections from user identities are disconnected. When the timer expires, the virtual file system is destroyed. All response files that were not deleted by the FTP client are deleted from their storage area. Enter a value in the range 1 - 43200. The default value is 600.",
							Computed:            true,
						},
						"virtual_directories": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the directories to create in virtual file system that the FTP server presents. The FTP client can use all of these directories to write file to be processed. The root directory (/) is always present, cannot be created, and is its own response directory.",
							NestedObject:        models.GetDmFTPServerVirtualDirectoryDataSourceSchema(),
							Computed:            true,
						},
						"default_directory": schema.StringAttribute{
							MarkdownDescription: "Specify the initial working directory on the FTP server after users connect and authenticate. For a virtual file system when the working directory is not the root directory, the specified directory must be a configured virtual directories. The default value is the root directory (/).",
							Computed:            true,
						},
						"max_filename_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum length for file names on the FTP server. Enter a value in the range 1 - 4000, The default value is 256.",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access control list",
							Computed:            true,
						},
						"require_tls": schema.StringAttribute{
							MarkdownDescription: "Specify whether FTP control connections require TLS encryption. For implicit or explicit FTP, you must complete the configuration by specifying the TLS profile to secure connections.",
							Computed:            true,
						},
						"password_aaa_policy": schema.StringAttribute{
							MarkdownDescription: "Specify the AAA policy to validate usernames and passwords provided to the FTP server by the client with the FTP <tt>USER</tt> and <tt>PASS</tt> commands. If authentication succeeds, the FTP client can use all the features of the FTP server. If authentication fails, a 530 error is returned and the user can attempt to authenticate again. Without this AAA policy, any username and password is accepted.",
							Computed:            true,
						},
						"certificate_aaa_policy": schema.StringAttribute{
							MarkdownDescription: "Specify the AAA policy to determine whether a password is required. This AAA policy provides secondary authentication of the information in the presented TLS certificate during TLS negotiation. Primary authentication is done by the TLS profile, which can reject the certificate. This authentication stage controls whether an FTP password is demanded or not. If authentication succeeds, the FTP client can use the <tt>USER</tt> command to login after the <tt>AUTH TLS</tt> command. If authentication fails, the FTP client must use both the <tt>USER</tt> and <b>PASS</b> commands to complete the login process. Without this AAA policy, the FTP client must use the <tt>USER</tt> and <tt>PASS</tt> commands.",
							Computed:            true,
						},
						"allow_ccc": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the FTP <tt>CCC</tt> command can be used to turn off TLS encryption of the FTP control connection after user authentication. Disabling encryption is necessary when the FTP control connection crosses a firewall or NAT device that must sniff the control connection. Disabling encryption eliminates the secrecy of the files during the transfer and allows TCP packets injection attacks.",
							Computed:            true,
						},
						"passive": schema.StringAttribute{
							MarkdownDescription: "Specify support for passive mode, which controls the use of the FTP <tt>PASV</tt> and <tt>EPSV</tt> commands by the FTP client. The <tt>PASV</tt> and <tt>EPSV</tt> commands are alternatives to the <tt>PORT</tt> and <tt>EPRT</tt> commands. The default behavior is to allow <tt>PORT</tt> , <tt>EPRT</tt> , <tt>PASV</tt> , and <tt>EPSV</tt> commands. Other settings either require the <tt>PASV</tt> or <tt>EPSV</tt> command or require the <tt>PORT</tt> or <tt>EPRT</tt> command. When an acceptable <tt>PORT</tt> , <tt>EPRT</tt> , <tt>PASV</tt> , or <tt>EPSV</tt> command is not received, the <tt>STOR</tt> , <tt>STOU</tt> , <tt>RETR</tt> , <tt>LIST</tt> , and <tt>NLST</tt> data transfer commands fail. In other words, the FTP server never attempts to connect to port 20, which is the deprecated default FTP data port.",
							Computed:            true,
						},
						"use_pasv_port_range": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to limit the port range for passive connections. <p>When enabled and the FTP server receives a <tt>PASV</tt> or <tt>EPSV</tt> command from the FTP client, the available port range is restricted from the listening port range 1024 - 65534. Use this setting when a firewall mandates that incoming FTP data connections on only a limited range of ports when it cannot sniff the FTP control connection.</p><p>The specified range limits how many FTP clients can be in the state between the receipt of the 227 response code to the <tt>PASV</tt> or <tt>EPSV</tt> command and the establishment of the FTP data connection. You can limit the pressure on this limited range by adjusting the idle timeout for passive data connections.</p><p><b>Note:</b> Do not configure a port range that overlaps with other services on the system. The system provides no check for these port conflicts. Generally, the other service allocates the ports, which makes these ports unavailable for the FTP server. The FTP server allocates these listing ports dynamically.</p>",
							Computed:            true,
						},
						"pasv_min_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the lowest port in the passive port range for data connections. This value must be less than the value of the maximum passive port. Enter a value in the range 1024 - 65534.",
							Computed:            true,
						},
						"pasv_max_port": schema.Int64Attribute{
							MarkdownDescription: "Specify the highest port in the passive port range for data connections. This value must be greater than the value of the minimum passive port. Enter a value in the range 1024 - 65534.",
							Computed:            true,
						},
						"pasv_idle_time_out": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds that the server waits for a client to establish a passive connection. Enter a value in the range 5 - 300. The default value is 60. <p>This setting controls the amount of time in seconds between when the FTP server issues code 227 (Entering Passive Mode) in response to the <tt>PASV</tt> or <tt>EPSV</tt> command from the FTP client and when the FTP client must establish a TCP data connection to the listening port and issue a data transfer command.</p><ul><li>If the data connection is not established within the timeout period, the listening port will be closed. If a data transfer command is issued after the port is closed, the command fails with code 425 and the <tt>Failed to open data connection</tt> message.</li><li>If the data connection is established but no data transfer command is issued within the timeout period, the TCP data connection will be closed. Any data transfer command after the timeout will be treated as if the <tt>PASV</tt> or <tt>EPSV</tt> command was never issued. The command fails with code 425 and the <tt>Require PASV or PORT command first</tt> message.</li></ul>",
							Computed:            true,
						},
						"disable_pasvip_check": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to disable the IP security check for passive data connections. This check verifies that the client IP address that connects to the data connection is the same IP address that established the control connection. This check is the expected behavior for an FTP server. Disable this check only when the incoming connection is not from the same client as the control connection. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"disable_portip_check": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to disable the IP security check for active data connections. This check verifies that the outgoing data connection can connect to only the client. This check is the expected behavior for an FTP server. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"use_alternate_pasv_addr": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use an alternate PASV IP address. When enabled, you can override the IP address that the FTP client presents to the server in passive mode. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"alternate_pasv_addr": schema.StringAttribute{
							MarkdownDescription: "Specify the IP address to return to the FTP client in response to a <tt>PASV</tt> command. This setting does not change the IP address that the FTP server listens, which is always the IP address for the FTP data connection. This value is used when the FTP server is behind a firewall that is not FTP-aware.",
							Computed:            true,
						},
						"allow_list_cmd": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to support the FTP <tt>LIST</tt> command. When enabled, the FTP server makes a distinction between the <tt>LIST</tt> and <tt>NLST</tt> commands. By default, the server always respond with an <tt>NLST</tt> to list files. <p>This setting is only supported by the FTP server handler for a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"allow_dele_cmd": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to support the FTP <tt>DELE</tt> command. When enabled, the <tt>DELE</tt> command can be passed to the FTP server. The default behavior is to not support the <tt>DELE</tt> command. <p>This setting is only valid in a transparent file system. This setting is only supported by the FTP server handler for a multiprotocol gateway or web service proxy.</p>",
							Computed:            true,
						},
						"data_encryption": schema.StringAttribute{
							MarkdownDescription: "Specify data encryption for file transfers. Data encryption is controlled by the FTP <tt>PROT P</tt> command.",
							Computed:            true,
						},
						"allow_compression": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the FTP client can use the FTP <tt>MODE Z</tt> command to compress data transfers. After compression is enabled, the FTP client can use the zlib method to compress data transfers.",
							Computed:            true,
						},
						"allow_stou": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the FTP client can use the FTP <tt>STOU</tt> command to generate unique file names. When enabled, the FTP server generates a unique file name for each transferred file.",
							Computed:            true,
						},
						"unique_filename_prefix": schema.StringAttribute{
							MarkdownDescription: "Specify the prefix for file names that the FTP <tt>STOU</tt> command generates. For the prefix, the directory separator (/ character) is not allowed. The default value is an empty string, which indicates to not add a prefix. <p><b>Note:</b> Processing adds a numeric suffix.</p>",
							Computed:            true,
						},
						"allow_rest": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the FTP client can use the FTP <tt>REST</tt> command after an interrupted file transfer. Restart ( <tt>REST</tt> ) is supported in the BSD stream style as described in <tt>draft-ietf-ftpext-mlst-16.txt</tt> . The MODE B style that is described in RFC 959 is not supported. The FTP server must be configured with a virtual persistent file system. <p>For written files, the server delays the processing until a timer expires or until the next FTP command other than a <tt>SIZE</tt> or <tt>REST</tt> command. With this processing, the FTP client can return and resume the transfer by using the <tt>SIZE</tt> , <tt>REST</tt> , and <tt>STOR</tt> commands. The argument to the <tt>REST</tt> command must be the same as the byte count the the <tt>SIZE</tt> command returned.</p>",
							Computed:            true,
						},
						"restart_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds to wait for a restart. When restart ( <tt>REST</tt> ) is enabled, the FTP client must reconnect to the server and use the <tt>SIZE</tt> , <tt>REST</tt> , and <tt>STOR</tt> commands to continue an interrupted file transfer. When this timer elapses, the previously received data on the data connection is passed to the DataPower service. This timer is canceled when a command other than <tt>SIZE</tt> or <tt>REST</tt> is received on the FTP control connection.",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration in seconds that the FTP control connection can be idle. After the duration elapses, the FTP server closes the control connection. The default value is 0, which disables the timer.",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Specify how to make response files available to FTP client for gateway transactions that are started by using an FTP STOR or SOUT operation.",
							Computed:            true,
						},
						"response_storage": schema.StringAttribute{
							MarkdownDescription: "Response storage",
							Computed:            true,
						},
						"temporary_storage_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size in MB for the temporary file system. Enter a value in the range 1 - 2048. The default value is 32.",
							Computed:            true,
						},
						"response_nfs_mount": schema.StringAttribute{
							MarkdownDescription: "When the response type is virtual file system and response storage is NFS, specify the NFS static mount to store response files. Each response file has a unique file name in the NFS directory. The name of the response file is not related to the file name that the virtual file system presents to the FTP client. Generally, this NFS directory is not made available through the FTP server. Do not use this directory for any other purpose.",
							Computed:            true,
						},
						"response_suffix": schema.StringAttribute{
							MarkdownDescription: "When the response type is a virtual file system, specify the suffix to add when generating response files. The directory separator (/ character) is not allowed. If the response type is a virtual file system and the FTP client is writing in virtual directories that do not have a response directory or if the response directory is the same as the virtual directory, this value must have a non-empty value. If empty, the response attempts to overwrite the request, which is not allowed. The default value is an empty string.",
							Computed:            true,
						},
						"ssl_server_config_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of TLS profile type to secure connections from clients. When a TLS profile is assigned, the FTP <tt>AUTH TLS</tt> command is enabled and clients can encrypt FTP control connection.",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *FTPServerSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *FTPServerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data FTPServerSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.FTPServerSourceProtocolHandler{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.FTPServerSourceProtocolHandler{}
	if value := res.Get(`FTPServerSourceProtocolHandler`); value.Exists() {
		for _, v := range value.Array() {
			item := models.FTPServerSourceProtocolHandler{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.FTPServerSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.FTPServerSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
