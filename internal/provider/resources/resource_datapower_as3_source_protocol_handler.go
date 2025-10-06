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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &AS3SourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &AS3SourceProtocolHandlerResource{}
var _ resource.ResourceWithImportState = &AS3SourceProtocolHandlerResource{}

func NewAS3SourceProtocolHandlerResource() resource.Resource {
	return &AS3SourceProtocolHandlerResource{}
}

type AS3SourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *AS3SourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as3_source_protocol_handler"
}

func (r *AS3SourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("AS3 handler", "source-as3", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the local IP address that the service listens. The default value is 0, which indicates that the service is active on all addresses. You can use a local host alias to help ease migration. <p>If the FTP server service is listening on a virtual IP address that is managed by a standby control group with the self-balancing feature enabled, incoming control and data connections always go to the active member of the standby control group.</p>", "address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the port that this service monitors. Enter a value in the range 1 - 65535. The default value is 21. <p>This port is the port on which FTP control connections are established. This port does not control the port for the data connections. If the FTP client uses the <tt>PASV</tt> command, data connections use an arbitrary, unused port.</p>", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("21").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(21),
			},
			"filesystem_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the file system type that the FTP server presents.", "filesystem", "").AddStringEnum("virtual-ephemeral", "virtual-persistent", "transparent").AddDefaultValue("virtual-ephemeral").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("virtual-ephemeral", "virtual-persistent", "transparent"),
				},
				Default: stringdefault.StaticString("virtual-ephemeral"),
			},
			"persistent_filesystem_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to retain a connection to a virtual file system after all FTP control connections from user identities are disconnected. When the timer expires, the virtual file system is destroyed. All response files that were not deleted by the FTP client are deleted from their storage area. Enter a value in the range 1 - 43200. The default value is 600.", "persistent-filesystem-timeout", "").AddIntegerRange(1, 43200).AddDefaultValue("600").AddNotValidWhen(models.AS3SourceProtocolHandlerPersistentFilesystemTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 43200),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AS3SourceProtocolHandlerPersistentFilesystemTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(600),
			},
			"virtual_directories": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directories to create in virtual file system that the FTP server presents. The FTP client can use all of these directories to write file to be processed. The root directory (/) is always present, cannot be created, and is its own response directory.", "virtual-directory", "").AddNotValidWhen(models.AS3SourceProtocolHandlerVirtualDirectoriesIgnoreVal.String()).String,
				NestedObject:        models.GetDmFTPServerVirtualDirectoryResourceSchema(),
				Optional:            true,
			},
			"default_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the initial working directory on the FTP server after users connect and authenticate. For a virtual file system when the working directory is not the root directory, the specified directory must be a configured virtual directories. The default value is the root directory (/).", "default-directory", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"max_filename_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum length for file names on the FTP server. Enter a value in the range 1 - 4000, The default value is 256.", "max-filename-len", "").AddIntegerRange(1, 4000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4000),
				},
				Default: int64default.StaticInt64(256),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "acl", "access_control_list").String,
				Optional:            true,
			},
			"require_tls": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether FTP control connections require TLS encryption. For implicit or explicit FTP, you must complete the configuration by specifying the TLS profile to secure connections.", "require-tls", "").AddStringEnum("off", "explicit", "implicit").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("off", "explicit", "implicit"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"password_aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the AAA policy to validate usernames and passwords provided to the FTP server by the client with the FTP <tt>USER</tt> and <tt>PASS</tt> commands. If authentication succeeds, the FTP client can use all the features of the FTP server. If authentication fails, a 530 error is returned and the user can attempt to authenticate again. Without this AAA policy, any username and password is accepted.", "password-aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"certificate_aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the AAA policy to determine whether a password is required. This AAA policy provides secondary authentication of the information in the presented TLS certificate during TLS negotiation. Primary authentication is done by the TLS profile, which can reject the certificate. This authentication stage controls whether an FTP password is demanded or not. If authentication succeeds, the FTP client can use the <tt>USER</tt> command to login after the <tt>AUTH TLS</tt> command. If authentication fails, the FTP client must use both the <tt>USER</tt> and <b>PASS</b> commands to complete the login process. Without this AAA policy, the FTP client must use the <tt>USER</tt> and <tt>PASS</tt> commands.", "certificate-aaa-policy", "aaa_policy").String,
				Optional:            true,
			},
			"allow_ccc": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the FTP <tt>CCC</tt> command can be used to turn off TLS encryption of the FTP control connection after user authentication. Disabling encryption is necessary when the FTP control connection crosses a firewall or NAT device that must sniff the control connection. Disabling encryption eliminates the secrecy of the files during the transfer and allows TCP packets injection attacks.", "allow-ccc", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify support for passive mode, which controls the use of the FTP <tt>PASV</tt> and <tt>EPSV</tt> commands by the FTP client. The <tt>PASV</tt> and <tt>EPSV</tt> commands are alternatives to the <tt>PORT</tt> and <tt>EPRT</tt> commands. The default behavior is to allow <tt>PORT</tt> , <tt>EPRT</tt> , <tt>PASV</tt> , and <tt>EPSV</tt> commands. Other settings either require the <tt>PASV</tt> or <tt>EPSV</tt> command or require the <tt>PORT</tt> or <tt>EPRT</tt> command. When an acceptable <tt>PORT</tt> , <tt>EPRT</tt> , <tt>PASV</tt> , or <tt>EPSV</tt> command is not received, the <tt>STOR</tt> , <tt>STOU</tt> , <tt>RETR</tt> , <tt>LIST</tt> , and <tt>NLST</tt> data transfer commands fail. In other words, the FTP server never attempts to connect to port 20, which is the deprecated default FTP data port.", "passive", "").AddStringEnum("allow", "disallow", "require").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "disallow", "require"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"use_pasv_port_range": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to limit the port range for passive connections. <p>When enabled and the FTP server receives a <tt>PASV</tt> or <tt>EPSV</tt> command from the FTP client, the available port range is restricted from the listening port range 1024 - 65534. Use this setting when a firewall mandates that incoming FTP data connections on only a limited range of ports when it cannot sniff the FTP control connection.</p><p>The specified range limits how many FTP clients can be in the state between the receipt of the 227 response code to the <tt>PASV</tt> or <tt>EPSV</tt> command and the establishment of the FTP data connection. You can limit the pressure on this limited range by adjusting the idle timeout for passive data connections.</p><p><b>Note:</b> Do not configure a port range that overlaps with other services on the system. The system provides no check for these port conflicts. Generally, the other service allocates the ports, which makes these ports unavailable for the FTP server. The FTP server allocates these listing ports dynamically.</p>", "passive-port-range", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerUsePasvPortRangeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pasv_min_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the lowest port in the passive port range for data connections. This value must be less than the value of the maximum passive port. Enter a value in the range 1024 - 65534.", "passive-port-min", "").AddIntegerRange(1024, 65534).AddDefaultValue("1024").AddNotValidWhen(models.AS3SourceProtocolHandlerPasvMinPortIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 65534),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AS3SourceProtocolHandlerPasvMinPortIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1024),
			},
			"pasv_max_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the highest port in the passive port range for data connections. This value must be greater than the value of the minimum passive port. Enter a value in the range 1024 - 65534.", "passive-port-max", "").AddIntegerRange(1024, 65534).AddDefaultValue("1050").AddNotValidWhen(models.AS3SourceProtocolHandlerPasvMaxPortIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 65534),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AS3SourceProtocolHandlerPasvMaxPortIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1050),
			},
			"pasv_idle_time_out": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that the server waits for a client to establish a passive connection. Enter a value in the range 5 - 300. The default value is 60. <p>This setting controls the amount of time in seconds between when the FTP server issues code 227 (Entering Passive Mode) in response to the <tt>PASV</tt> or <tt>EPSV</tt> command from the FTP client and when the FTP client must establish a TCP data connection to the listening port and issue a data transfer command.</p><ul><li>If the data connection is not established within the timeout period, the listening port will be closed. If a data transfer command is issued after the port is closed, the command fails with code 425 and the <tt>Failed to open data connection</tt> message.</li><li>If the data connection is established but no data transfer command is issued within the timeout period, the TCP data connection will be closed. Any data transfer command after the timeout will be treated as if the <tt>PASV</tt> or <tt>EPSV</tt> command was never issued. The command fails with code 425 and the <tt>Require PASV or PORT command first</tt> message.</li></ul>", "passive-idle-timeout", "").AddIntegerRange(5, 300).AddDefaultValue("60").AddNotValidWhen(models.AS3SourceProtocolHandlerPasvIdleTimeOutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 300),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AS3SourceProtocolHandlerPasvIdleTimeOutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"disable_pasv_ip_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to disable the IP security check for passive data connections. This check verifies that the client IP address that connects to the data connection is the same IP address that established the control connection. This check is the expected behavior for an FTP server. Disable this check only when the incoming connection is not from the same client as the control connection. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>", "passive-promiscuous", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerDisablePASVIPCheckIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_port_ip_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to disable the IP security check for active data connections. This check verifies that the outgoing data connection can connect to only the client. This check is the expected behavior for an FTP server. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>", "port-promiscuous", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerDisablePORTIPCheckIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_alternate_pasv_addr": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use an alternate PASV IP address. When enabled, you can override the IP address that the FTP client presents to the server in passive mode. <p>This setting is supported by only the FTP server handler with a multiprotocol gateway or web service proxy.</p>", "allow-passive-addr", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerUseAlternatePASVAddrIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"alternate_pasv_addr": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IP address to return to the FTP client in response to a <tt>PASV</tt> command. This setting does not change the IP address that the FTP server listens, which is always the IP address for the FTP data connection. This value is used when the FTP server is behind a firewall that is not FTP-aware.", "passive-addr", "").AddRequiredWhen(models.AS3SourceProtocolHandlerAlternatePASVAddrCondVal.String()).AddNotValidWhen(models.AS3SourceProtocolHandlerAlternatePASVAddrIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AS3SourceProtocolHandlerAlternatePASVAddrCondVal, models.AS3SourceProtocolHandlerAlternatePASVAddrIgnoreVal, false),
				},
			},
			"allow_list_cmd": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to support the FTP <tt>LIST</tt> command. When enabled, the FTP server makes a distinction between the <tt>LIST</tt> and <tt>NLST</tt> commands. By default, the server always respond with an <tt>NLST</tt> to list files. <p>This setting is only supported by the FTP server handler for a multiprotocol gateway or web service proxy.</p>", "list-cmd", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerAllowLISTCmdIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_dele_cmd": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to support the FTP <tt>DELE</tt> command. When enabled, the <tt>DELE</tt> command can be passed to the FTP server. The default behavior is to not support the <tt>DELE</tt> command. <p>This setting is only valid in a transparent file system. This setting is only supported by the FTP server handler for a multiprotocol gateway or web service proxy.</p>", "dele-cmd", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerAllowDELECmdIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"data_encryption": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify data encryption for file transfers. Data encryption is controlled by the FTP <tt>PROT P</tt> command.", "data-encryption", "").AddStringEnum("disallow", "allow", "require").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disallow", "allow", "require"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the FTP client can use the FTP <tt>MODE Z</tt> command to compress data transfers. After compression is enabled, the FTP client can use the zlib method to compress data transfers.", "allow-compression", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_stou": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the FTP client can use the FTP <tt>STOU</tt> command to generate unique file names. When enabled, the FTP server generates a unique file name for each transferred file.", "allow-unique-filename", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"unique_filename_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the prefix for file names that the FTP <tt>STOU</tt> command generates. For the prefix, the directory separator (/ character) is not allowed. The default value is an empty string, which indicates to not add a prefix. <p><b>Note:</b> Processing adds a numeric suffix.</p>", "unique-filename-prefix", "").AddNotValidWhen(models.AS3SourceProtocolHandlerUniqueFilenamePrefixIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^/]*$"), "Must match :"+"^[^/]*$"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.AS3SourceProtocolHandlerUniqueFilenamePrefixIgnoreVal, false),
				},
			},
			"allow_rest": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the FTP client can use the FTP <tt>REST</tt> command after an interrupted file transfer. Restart ( <tt>REST</tt> ) is supported in the BSD stream style as described in <tt>draft-ietf-ftpext-mlst-16.txt</tt> . The MODE B style that is described in RFC 959 is not supported. The FTP server must be configured with a virtual persistent file system. <p>For written files, the server delays the processing until a timer expires or until the next FTP command other than a <tt>SIZE</tt> or <tt>REST</tt> command. With this processing, the FTP client can return and resume the transfer by using the <tt>SIZE</tt> , <tt>REST</tt> , and <tt>STOR</tt> commands. The argument to the <tt>REST</tt> command must be the same as the byte count the the <tt>SIZE</tt> command returned.</p>", "allow-restart", "").AddDefaultValue("false").AddNotValidWhen(models.AS3SourceProtocolHandlerAllowRESTIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"restart_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for a restart. When restart ( <tt>REST</tt> ) is enabled, the FTP client must reconnect to the server and use the <tt>SIZE</tt> , <tt>REST</tt> , and <tt>STOR</tt> commands to continue an interrupted file transfer. When this timer elapses, the previously received data on the data connection is passed to the DataPower service. This timer is canceled when a command other than <tt>SIZE</tt> or <tt>REST</tt> is received on the FTP control connection.", "restart-timeout", "").AddDefaultValue("240").AddNotValidWhen(models.AS3SourceProtocolHandlerRestartTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(240),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that the FTP control connection can be idle. After the duration elapses, the FTP server closes the control connection. The default value is 0, which disables the timer.", "idle-timeout", "").AddIntegerRange(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response type", "response-type", "").AddStringEnum("none", "virtual-filesystem").AddDefaultValue("none").AddNotValidWhen(models.AS3SourceProtocolHandlerResponseTypeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "virtual-filesystem"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.AS3SourceProtocolHandlerResponseTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("none"),
			},
			"response_storage": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response storage", "response-storage", "").AddStringEnum("temporary", "nfs").AddDefaultValue("temporary").AddNotValidWhen(models.AS3SourceProtocolHandlerResponseStorageIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("temporary", "nfs"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.AS3SourceProtocolHandlerResponseStorageIgnoreVal, true),
				},
				Default: stringdefault.StaticString("temporary"),
			},
			"temporary_storage_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Temporary storage size", "filesystem-size", "").AddIntegerRange(1, 2048).AddDefaultValue("32").AddNotValidWhen(models.AS3SourceProtocolHandlerTemporaryStorageSizeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2048),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AS3SourceProtocolHandlerTemporaryStorageSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(32),
			},
			"response_nf_sm_ount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response NFS mount", "response-nfs-mount", "nfs_static_mount").AddNotValidWhen(models.AS3SourceProtocolHandlerResponseNFSMountIgnoreVal.String()).String,
				Optional:            true,
			},
			"response_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response suffix", "response-suffix", "").AddNotValidWhen(models.AS3SourceProtocolHandlerResponseSuffixIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^/]*$"), "Must match :"+"^[^/]*$"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.AS3SourceProtocolHandlerResponseSuffixIgnoreVal, false),
				},
			},
			"response_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response URL", "response-url", "").AddNotValidWhen(models.AS3SourceProtocolHandlerResponseURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^ftp://[^/]+(/[^/]+)*$"), "Must match :"+"^ftp://[^/]+(/[^/]+)*$"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.AS3SourceProtocolHandlerResponseURLIgnoreVal, false),
				},
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile type to secure connections from clients. When a TLS profile is assigned, the FTP <tt>AUTH TLS</tt> command is enabled and clients can encrypt FTP control connection.", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").AddRequiredWhen(models.AS3SourceProtocolHandlerSSLServerConfigTypeCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
					validators.ConditionalRequiredString(models.AS3SourceProtocolHandlerSSLServerConfigTypeCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "ssl_server_profile").AddRequiredWhen(models.AS3SourceProtocolHandlerSSLServerCondVal.String()).AddNotValidWhen(models.AS3SourceProtocolHandlerSSLServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AS3SourceProtocolHandlerSSLServerCondVal, models.AS3SourceProtocolHandlerSSLServerIgnoreVal, false),
				},
			},
			"ssl_sni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "ssl_sni_server_profile").AddRequiredWhen(models.AS3SourceProtocolHandlerSSLSNIServerCondVal.String()).AddNotValidWhen(models.AS3SourceProtocolHandlerSSLSNIServerIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AS3SourceProtocolHandlerSSLSNIServerCondVal, models.AS3SourceProtocolHandlerSSLSNIServerIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AS3SourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AS3SourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS3SourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AS3SourceProtocolHandler`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS3SourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS3SourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	data.UpdateFromBody(ctx, `AS3SourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS3SourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS3SourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AS3SourceProtocolHandler`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS3SourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AS3SourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *AS3SourceProtocolHandlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.AS3SourceProtocolHandler
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `AS3SourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS3SourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AS3SourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}
