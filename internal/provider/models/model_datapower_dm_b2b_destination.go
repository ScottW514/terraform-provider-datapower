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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmB2BDestination struct {
	DestName                                   types.String         `tfsdk:"dest_name"`
	DestUrl                                    types.String         `tfsdk:"dest_url"`
	EnabledDocType                             *DmB2BEnabledDocType `tfsdk:"enabled_doc_type"`
	SmtpServerConnection                       types.String         `tfsdk:"smtp_server_connection"`
	EmailAddress                               types.String         `tfsdk:"email_address"`
	SshClientConnection                        types.String         `tfsdk:"ssh_client_connection"`
	UseUniqueFilenames                         types.Bool           `tfsdk:"use_unique_filenames"`
	OverrideTimeout                            types.Int64          `tfsdk:"override_timeout"`
	EnableFtpSettings                          types.Bool           `tfsdk:"enable_ftp_settings"`
	UserName                                   types.String         `tfsdk:"user_name"`
	PasswordAlias                              types.String         `tfsdk:"password_alias"`
	EbmsmpcAuthMethod                          types.String         `tfsdk:"ebmsmpc_auth_method"`
	UserNameToken                              types.String         `tfsdk:"user_name_token"`
	UserNameTokenPasswordAlias                 types.String         `tfsdk:"user_name_token_password_alias"`
	EbmsmpcVerifyValCred                       types.String         `tfsdk:"ebmsmpc_verify_val_cred"`
	Passive                                    types.String         `tfsdk:"passive"`
	AuthTls                                    types.String         `tfsdk:"auth_tls"`
	UseCcc                                     types.String         `tfsdk:"use_ccc"`
	EncryptData                                types.String         `tfsdk:"encrypt_data"`
	DataType                                   types.String         `tfsdk:"data_type"`
	SlashStou                                  types.String         `tfsdk:"slash_stou"`
	QuotedCommands                             types.String         `tfsdk:"quoted_commands"`
	SizeCheck                                  types.String         `tfsdk:"size_check"`
	BinaryTransferMode                         types.String         `tfsdk:"binary_transfer_mode"`
	AsCompress                                 types.Bool           `tfsdk:"as_compress"`
	AsCompressBeforeSign                       types.Bool           `tfsdk:"as_compress_before_sign"`
	AsSendUnsigned                             types.Bool           `tfsdk:"as_send_unsigned"`
	AsEncrypt                                  types.Bool           `tfsdk:"as_encrypt"`
	AsEncryptCert                              types.String         `tfsdk:"as_encrypt_cert"`
	AsmdnRequest                               types.Bool           `tfsdk:"asmdn_request"`
	AsmdnRequestAsync                          types.Bool           `tfsdk:"asmdn_request_async"`
	As1mdnRedirectEmail                        types.String         `tfsdk:"as1mdn_redirect_email"`
	As2mdnRedirectUrl                          types.String         `tfsdk:"as2mdn_redirect_url"`
	As3mdnRedirectUrl                          types.String         `tfsdk:"as3mdn_redirect_url"`
	AsmdnRequestSigned                         types.Bool           `tfsdk:"asmdn_request_signed"`
	Retransmit                                 types.Bool           `tfsdk:"retransmit"`
	AckTime                                    types.Int64          `tfsdk:"ack_time"`
	MaxResends                                 types.Int64          `tfsdk:"max_resends"`
	AsEncryptAlg                               types.String         `tfsdk:"as_encrypt_alg"`
	AsmdnRequestSignedAlgs                     types.String         `tfsdk:"asmdn_request_signed_algs"`
	EbmsCpaId                                  types.String         `tfsdk:"ebms_cpa_id"`
	EbmsService                                types.String         `tfsdk:"ebms_service"`
	EbmsServiceType                            types.String         `tfsdk:"ebms_service_type"`
	EbmsAction                                 types.String         `tfsdk:"ebms_action"`
	EbmsSendUnsigned                           types.Bool           `tfsdk:"ebms_send_unsigned"`
	EbmsEncrypt                                types.Bool           `tfsdk:"ebms_encrypt"`
	EbmsEncryptCert                            types.String         `tfsdk:"ebms_encrypt_cert"`
	EbmsEncryptAlg                             types.String         `tfsdk:"ebms_encrypt_alg"`
	EbmsKeyEncryptAlg                          types.String         `tfsdk:"ebms_key_encrypt_alg"`
	EbmsDuplicateEliminationRequest            types.Bool           `tfsdk:"ebms_duplicate_elimination_request"`
	EbmsAckRequest                             types.Bool           `tfsdk:"ebms_ack_request"`
	EbmsAckRequestSigned                       types.Bool           `tfsdk:"ebms_ack_request_signed"`
	EbmsSyncReplyMode                          types.String         `tfsdk:"ebms_sync_reply_mode"`
	EbmsRetry                                  types.Bool           `tfsdk:"ebms_retry"`
	EbmsMaxRetries                             types.Int64          `tfsdk:"ebms_max_retries"`
	EbmsRetryInterval                          types.Int64          `tfsdk:"ebms_retry_interval"`
	EbmsIncludeTimeToLive                      types.Bool           `tfsdk:"ebms_include_time_to_live"`
	SslClientConfigType                        types.String         `tfsdk:"ssl_client_config_type"`
	SslClient                                  types.String         `tfsdk:"ssl_client"`
	EbmsMessageExchangePattern                 types.String         `tfsdk:"ebms_message_exchange_pattern"`
	EbmsMessagePartitionChannel                types.String         `tfsdk:"ebms_message_partition_channel"`
	EbmsAgreementRef                           types.String         `tfsdk:"ebms_agreement_ref"`
	EbmspMode                                  types.String         `tfsdk:"ebmsp_mode"`
	EbmsOutboundRequestReceipt                 types.Bool           `tfsdk:"ebms_outbound_request_receipt"`
	EbmsOutboundRequestSignedReceipt           types.Bool           `tfsdk:"ebms_outbound_request_signed_receipt"`
	EbmsOutboundReceiptReplyPattern            types.String         `tfsdk:"ebms_outbound_receipt_reply_pattern"`
	EbmsOutboundReceptionAwarenessNotification types.Bool           `tfsdk:"ebms_outbound_reception_awareness_notification"`
	EbmsOutboundReceptionAwarenessTimeout      types.Int64          `tfsdk:"ebms_outbound_reception_awareness_timeout"`
	EbmsCompress                               types.Bool           `tfsdk:"ebms_compress"`
	EbmssoapBody                               types.Bool           `tfsdk:"ebmssoap_body"`
}

var DmB2BDestinationObjectType = map[string]attr.Type{
	"dest_name":                            types.StringType,
	"dest_url":                             types.StringType,
	"enabled_doc_type":                     types.ObjectType{AttrTypes: DmB2BEnabledDocTypeObjectType},
	"smtp_server_connection":               types.StringType,
	"email_address":                        types.StringType,
	"ssh_client_connection":                types.StringType,
	"use_unique_filenames":                 types.BoolType,
	"override_timeout":                     types.Int64Type,
	"enable_ftp_settings":                  types.BoolType,
	"user_name":                            types.StringType,
	"password_alias":                       types.StringType,
	"ebmsmpc_auth_method":                  types.StringType,
	"user_name_token":                      types.StringType,
	"user_name_token_password_alias":       types.StringType,
	"ebmsmpc_verify_val_cred":              types.StringType,
	"passive":                              types.StringType,
	"auth_tls":                             types.StringType,
	"use_ccc":                              types.StringType,
	"encrypt_data":                         types.StringType,
	"data_type":                            types.StringType,
	"slash_stou":                           types.StringType,
	"quoted_commands":                      types.StringType,
	"size_check":                           types.StringType,
	"binary_transfer_mode":                 types.StringType,
	"as_compress":                          types.BoolType,
	"as_compress_before_sign":              types.BoolType,
	"as_send_unsigned":                     types.BoolType,
	"as_encrypt":                           types.BoolType,
	"as_encrypt_cert":                      types.StringType,
	"asmdn_request":                        types.BoolType,
	"asmdn_request_async":                  types.BoolType,
	"as1mdn_redirect_email":                types.StringType,
	"as2mdn_redirect_url":                  types.StringType,
	"as3mdn_redirect_url":                  types.StringType,
	"asmdn_request_signed":                 types.BoolType,
	"retransmit":                           types.BoolType,
	"ack_time":                             types.Int64Type,
	"max_resends":                          types.Int64Type,
	"as_encrypt_alg":                       types.StringType,
	"asmdn_request_signed_algs":            types.StringType,
	"ebms_cpa_id":                          types.StringType,
	"ebms_service":                         types.StringType,
	"ebms_service_type":                    types.StringType,
	"ebms_action":                          types.StringType,
	"ebms_send_unsigned":                   types.BoolType,
	"ebms_encrypt":                         types.BoolType,
	"ebms_encrypt_cert":                    types.StringType,
	"ebms_encrypt_alg":                     types.StringType,
	"ebms_key_encrypt_alg":                 types.StringType,
	"ebms_duplicate_elimination_request":   types.BoolType,
	"ebms_ack_request":                     types.BoolType,
	"ebms_ack_request_signed":              types.BoolType,
	"ebms_sync_reply_mode":                 types.StringType,
	"ebms_retry":                           types.BoolType,
	"ebms_max_retries":                     types.Int64Type,
	"ebms_retry_interval":                  types.Int64Type,
	"ebms_include_time_to_live":            types.BoolType,
	"ssl_client_config_type":               types.StringType,
	"ssl_client":                           types.StringType,
	"ebms_message_exchange_pattern":        types.StringType,
	"ebms_message_partition_channel":       types.StringType,
	"ebms_agreement_ref":                   types.StringType,
	"ebmsp_mode":                           types.StringType,
	"ebms_outbound_request_receipt":        types.BoolType,
	"ebms_outbound_request_signed_receipt": types.BoolType,
	"ebms_outbound_receipt_reply_pattern":  types.StringType,
	"ebms_outbound_reception_awareness_notification": types.BoolType,
	"ebms_outbound_reception_awareness_timeout":      types.Int64Type,
	"ebms_compress": types.BoolType,
	"ebmssoap_body": types.BoolType,
}
var DmB2BDestinationObjectDefault = map[string]attr.Value{
	"dest_name":                            types.StringNull(),
	"dest_url":                             types.StringNull(),
	"enabled_doc_type":                     types.ObjectValueMust(DmB2BEnabledDocTypeObjectType, DmB2BEnabledDocTypeObjectDefault),
	"smtp_server_connection":               types.StringValue("default"),
	"email_address":                        types.StringNull(),
	"ssh_client_connection":                types.StringNull(),
	"use_unique_filenames":                 types.BoolValue(false),
	"override_timeout":                     types.Int64Value(300),
	"enable_ftp_settings":                  types.BoolValue(false),
	"user_name":                            types.StringNull(),
	"password_alias":                       types.StringNull(),
	"ebmsmpc_auth_method":                  types.StringValue("username-token"),
	"user_name_token":                      types.StringNull(),
	"user_name_token_password_alias":       types.StringNull(),
	"ebmsmpc_verify_val_cred":              types.StringNull(),
	"passive":                              types.StringValue("pasv-req"),
	"auth_tls":                             types.StringValue("auth-off"),
	"use_ccc":                              types.StringValue("ccc-off"),
	"encrypt_data":                         types.StringValue("enc-data-off"),
	"data_type":                            types.StringValue("binary"),
	"slash_stou":                           types.StringValue("slash-stou-on"),
	"quoted_commands":                      types.StringNull(),
	"size_check":                           types.StringValue("size-check-optional"),
	"binary_transfer_mode":                 types.StringValue("auto-detect"),
	"as_compress":                          types.BoolValue(false),
	"as_compress_before_sign":              types.BoolValue(false),
	"as_send_unsigned":                     types.BoolValue(false),
	"as_encrypt":                           types.BoolValue(false),
	"as_encrypt_cert":                      types.StringNull(),
	"asmdn_request":                        types.BoolValue(false),
	"asmdn_request_async":                  types.BoolValue(false),
	"as1mdn_redirect_email":                types.StringNull(),
	"as2mdn_redirect_url":                  types.StringNull(),
	"as3mdn_redirect_url":                  types.StringNull(),
	"asmdn_request_signed":                 types.BoolValue(false),
	"retransmit":                           types.BoolValue(false),
	"ack_time":                             types.Int64Value(1800),
	"max_resends":                          types.Int64Value(3),
	"as_encrypt_alg":                       types.StringValue("3des"),
	"asmdn_request_signed_algs":            types.StringValue("sha1,md5"),
	"ebms_cpa_id":                          types.StringNull(),
	"ebms_service":                         types.StringNull(),
	"ebms_service_type":                    types.StringNull(),
	"ebms_action":                          types.StringNull(),
	"ebms_send_unsigned":                   types.BoolValue(false),
	"ebms_encrypt":                         types.BoolValue(false),
	"ebms_encrypt_cert":                    types.StringNull(),
	"ebms_encrypt_alg":                     types.StringValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
	"ebms_key_encrypt_alg":                 types.StringValue("http://www.w3.org/2001/04/xmlenc#rsa-1_5"),
	"ebms_duplicate_elimination_request":   types.BoolValue(true),
	"ebms_ack_request":                     types.BoolValue(false),
	"ebms_ack_request_signed":              types.BoolValue(false),
	"ebms_sync_reply_mode":                 types.StringValue("none"),
	"ebms_retry":                           types.BoolValue(false),
	"ebms_max_retries":                     types.Int64Value(3),
	"ebms_retry_interval":                  types.Int64Value(1800),
	"ebms_include_time_to_live":            types.BoolValue(true),
	"ssl_client_config_type":               types.StringValue("client"),
	"ssl_client":                           types.StringNull(),
	"ebms_message_exchange_pattern":        types.StringValue("one-way-push"),
	"ebms_message_partition_channel":       types.StringNull(),
	"ebms_agreement_ref":                   types.StringNull(),
	"ebmsp_mode":                           types.StringNull(),
	"ebms_outbound_request_receipt":        types.BoolValue(false),
	"ebms_outbound_request_signed_receipt": types.BoolValue(false),
	"ebms_outbound_receipt_reply_pattern":  types.StringValue("Response"),
	"ebms_outbound_reception_awareness_notification": types.BoolValue(false),
	"ebms_outbound_reception_awareness_timeout":      types.Int64Value(300),
	"ebms_compress": types.BoolValue(false),
	"ebmssoap_body": types.BoolValue(false),
}
var DmB2BDestinationDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"dest_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name to identify this destination in the profile.", "name", "").String,
			Computed:            true,
		},
		"dest_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL for the destination. The URL is parsed to determine the protocol. For load distribution, specify the name of the load balancer instead of the address-port pair. <ul><li>If the URL starts with as1:// or dpsmtp://, the service uses the ESMTP protocol. These protocols require an SMTP server connection and email address. Specify the URL in one of the following formats. <ul><li><tt>as1://mailExchange[:port][/queryParameters]</tt></li><li><tt>dpsmtp://mailExchange[:port][/queryParameters]</tt></li></ul></li><li>If the URL starts with as2://, ebms2://, ebms3://, or http://, the service uses the HTTP protocol. Specify the URL in the <tt>hostname:port/uri</tt> format.</li><li>If the URL starts with as2s://, ebms2s://, ebms3s://, or https://, the service uses the HTTPS protocol. Specify the URL in the <tt>hostname:port/uri</tt> format.</li><li>If the URL starts with as3:// or ftp://, the service uses the FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the <tt>ftp://user:password@host/%2Fpath</tt> format. %2F is the URL encoding of a forward slash.</li><li>For a relative path to the login directory of the user, specify the URL in the <tt>ftp://user:password@host:port/path</tt> format.</li></ul></li><li>If the URL starts with sftp://, the service uses the SSH FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the <tt>sftp://host:port/path</tt> format.</li><li>For a relative path to the login directory of the user, specify the URL in the <tt>sftp://host:port/~/path</tt> format.</li></ul></li><li>If the URL starts with dpnfs://, the service uses the NFS protocol. <ul><li>For static mounts, specify the URL in the <tt>dpnfs://MountName</tt> format, where <tt>MountName</tt> is the name of an NFS mount.</li><li>For dynamic mounts, specify the URL in the <tt>dpnfs://host</tt> or <tt>dpnfs://host/path</tt> format.</li></ul></li><li>If the URL starts with dpims:// or dpimsssl://, the service uses the IMS protocol. Specify the URL in one of the following formats. <ul><li><tt>dpims://ConnectObject/?parameters</tt></li><li><tt>dpimsssl://ConnectObject/?parameters</tt></li></ul></li><li>If the URL starts with dpmq://, the service uses the IBM MQ protocol. The queue manager must exist in current domain. <ul><li>To send a message, specify the URL in the <tt>dpmq://QueueManager/URI?RequestQueue=requestQueue;queryParameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dpmq://QueueManager/URI?ReplyQueue=replyQueue;queryParameters</tt> format.</li></ul><p>If the URL starts with mq://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dpmqfte://, the service uses the IBM MQ FTE protocol. The queue manager must exist in current domain. <ul><li>To send a message, specify the URL in the <tt>dpmqfte://QueueManager/?RequestQueue=request_queue_name;queryParameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dpmqfte://QueueManager/?ReplyQueue=reply_queue_name;queryParameters</tt> format.</li></ul><p>If the URL starts with mqfte://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dptibems://, the service uses the TIBCO EMS protocol. The server must exist. <ul><li>To send a message, specify the URL in the <tt>dptibems://Server/?RequestQueue=queue;RequestReply=queue;query-parameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dptibems://Server/?ReplyQueue=queue;query-parameters</tt> format.</li></ul><p>If the URL starts with tibems://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dpwasjms://, the service uses the WebSphere JMS protocol. The server must exist. <ul><li>To send a message to a request queue, specify the URL in the <tt>dpwasjms://Server/?RequestQueue=queue;RequestReply=queue;query-parameters</tt> format.</li><li>To send a message to a request topic space, specify the URL in the <tt>dpwasjms://Server/?RequestTopicSpace=topic-space;RequestReply=queue;query-parameters</tt> format.</li><li>To retrieve a message from a reply queue, specify the URL in the <tt>dpwasjms://Server/?ReplyQueue=queue;query-parameters</tt> format.</li><li>To retrieve a message from a reply topic space, specify the URL in the <tt>dpwasjms://Server/?ReplyTopicSpace=topic-space;query-parameters</tt> format.</li></ul></li></ul>", "dest-url", "").String,
			Computed:            true,
		},
		"enabled_doc_type": GetDmB2BEnabledDocTypeDataSourceSchema("Specify the document types to support.", "enabled-doc-type", ""),
		"smtp_server_connection": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SMTP server connection to send email messages. By default, the gateway uses the default SMTP server connection.", "smtp-server-connection", "smtp_server_connection").AddDefaultValue("default").String,
			Computed:            true,
		},
		"email_address": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination email address for messages sent to this partner.", "email-address", "").String,
			Computed:            true,
		},
		"ssh_client_connection": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH profile for SSH client connections and authentication.", "ssh-client-connection", "ssh_client_profile").String,
			Computed:            true,
		},
		"use_unique_filenames": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a unique file name for puts to a remote directory.", "use-unique-filenames", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"override_timeout": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to maintain an idle connection in seconds. Enter a value in the range 3 - 7200. The default value is 300.", "timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
			Computed:            true,
		},
		"enable_ftp_settings": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to override the FTP client policy. This policy is defined in user agent of the XML manager for the B2B gateway. <ul><li>When enabled, define the overrides with the advanced AS3 or FTP settings.</li><li>When disabled, uses the original FTP client policy.</li></ul>", "enable-ftp-settings", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"user_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the username to override in the basic authentication policy of the user agent.", "username", "").String,
			Computed:            true,
		},
		"password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to override in the basic authentication policy of the user agent.", "password-alias", "password_alias").String,
			Computed:            true,
		},
		"ebmsmpc_auth_method": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify how the MPC authenticates the incoming pull requests. By default, the MPC authenticates requests through username token.", "embs-mpc-auth-method", "").AddStringEnum("username-token", "cert").AddDefaultValue("username-token").String,
			Computed:            true,
		},
		"user_name_token": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("For ebMS3, the WS-Security UsernameToken authorizes received messages. Enter the username for the wsse:UsernameToken element to package in the SOAP header of the message to send. <ul><li>For one-way push exchange pattern, specify the username token to package in messages to send.</li><li>For one-way pull exchange pattern outbound messages, when the MPC authentication method is username token, specify the username to store messages in the MPC.</li><li>For one-way pull exchange pattern inbound messages, specify the username token to package in pull requests.</li></ul>", "username-token", "").String,
			Computed:            true,
		},
		"user_name_token_password_alias": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias of the username for the wsse:UsernameToken element.", "username-token-password-alias", "password_alias").String,
			Computed:            true,
		},
		"ebmsmpc_verify_val_cred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When the MPC authentication method is validation credentials, specify the certificate to associate with messages that are submitted to the MPC.", "ebms-mpc-verify-valcred", "crypto_val_cred").String,
			Computed:            true,
		},
		"passive": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of FTP passive mode with the FTP <tt>PASV</tt> command.", "ftp-passive", "").AddStringEnum("pasv-off", "pasv-opt", "pasv-req").AddDefaultValue("pasv-req").String,
			Computed:            true,
		},
		"auth_tls": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of TLS to secure connections with the FTP <tt>AUTH TLS</tt> command.", "ftp-auth-tls", "").AddStringEnum("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp").AddDefaultValue("auth-off").String,
			Computed:            true,
		},
		"use_ccc": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of command encryption after authentication with the FTP <tt>CCC</tt> command.", "ftp-use-ccc", "").AddStringEnum("ccc-off", "ccc-opt", "ccc-req").AddDefaultValue("ccc-off").String,
			Computed:            true,
		},
		"encrypt_data": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the encryption of file transfers with the FTP <tt>PROT P</tt> command.", "ftp-encrypt-data", "").AddStringEnum("enc-data-off", "enc-data-opt", "enc-data-req").AddDefaultValue("enc-data-off").String,
			Computed:            true,
		},
		"data_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the data type of file transfers.", "ftp-data-type", "").AddStringEnum("ascii", "binary").AddDefaultValue("binary").String,
			Computed:            true,
		},
		"slash_stou": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage unique files when the name contains a trailing slash with either the FTP <tt>STOU</tt> or <tt>STOR</tt> command. Some FTP servers do not support the <tt>STOU</tt> command.", "ftp-slash-stou", "").AddStringEnum("slash-stou-off", "slash-stou-on").AddDefaultValue("slash-stou-on").String,
			Computed:            true,
		},
		"quoted_commands": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the FTP quoted commands list that defines the command to send to the FTP server. before each <tt>STOU</tt> , <tt>STOR</tt> , or <tt>RETR</tt> command. The commands in the list cannot be data-transfer related commands.", "ftp-quoted-commands", "ftp_quote_commands").String,
			Computed:            true,
		},
		"size_check": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to check the size of file after transfer with the FTP <tt>SIZE</tt> command. Some FTP servers, particularly vsftpd in the default configuration, provide inaccurate responses for files transferred in ASCII mode. If you get such errors, disable this feature or reconfigure vsftpd.", "ftp-size-check", "").AddStringEnum("size-check-optional", "size-check-disabled").AddDefaultValue("size-check-optional").String,
			Computed:            true,
		},
		"binary_transfer_mode": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to transfer the message payload in binary mode.", "binary-transfer-mode", "").AddStringEnum("auto-detect", "enforce").AddDefaultValue("auto-detect").String,
			Computed:            true,
		},
		"as_compress": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound AS messages. The default behavior is disabled.", "as-compress", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"as_compress_before_sign": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound AS messages before signing. The default behavior is disabled. <ul><li>If enabled, compress the message body before signing.</li><li>If disabled, sign the message before compressing the message body.</li></ul>", "as-compress-before-sign", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"as_send_unsigned": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to override the signing of messages to this destination. Whether to sign outbound message is part of the configuration of the internal partner. This property cannot be used to cause a message to be signed. The default behavior is off. <ul><li>If enabled, never sign messages.</li><li>If disabled, sign messages when the sender has signing credentials.</li></ul>", "as-send-unsigned", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"as_encrypt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encrypt the body of outgoing AS messages. The default behavior is disabled.", "as-encrypt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"as_encrypt_cert": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to encrypt outbound AS messages. Use the name of a certificate.", "as-encrypt-cert", "crypto_certificate").String,
			Computed:            true,
		},
		"asmdn_request": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request an MDN for outbound AS messages. The default behavior is disabled.", "as-mdn-request", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"asmdn_request_async": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the MDN request for outbound AS messages is asynchronous. The default behavior is disabled. <ul><li>If enabled, the MDN request is asynchronous</li><li>If disabled, the MDN request is synchronous</li></ul>", "as-mdn-request-async", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"as1mdn_redirect_email": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection email address for the MDN to outbound AS1 messages. The partner that receives the outbound AS1 message sends the MDN to this email address. When blank, the redirection email in the <tt>From</tt> address of the outgoing message (the default email address of the sending internal partner).", "as1-mdn-email", "").String,
			Computed:            true,
		},
		"as2mdn_redirect_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL for the MDN to outbound AS2 messages.", "as2-mdn-url", "").String,
			Computed:            true,
		},
		"as3mdn_redirect_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL for the MDN to outbound AS3 messages.", "as3-mdn-url", "").String,
			Computed:            true,
		},
		"asmdn_request_signed": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed MDN instead of an unsigned one. The default behavior is an unsigned one.", "as-mdn-request-signed", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"retransmit": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retransmit messages. The default behavior is off.", "retransmit", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ack_time": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for the acknowledgment before a message is retransmitted. Enter a value in the range 1 - 3600. The default value is 1800.", "ack-time", "").AddIntegerRange(1, 3600).AddDefaultValue("1800").String,
			Computed:            true,
		},
		"max_resends": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of attempts to retransmit a message. Enter a value in the range 1 - 30. The default value is 3.", "max-resends", "").AddIntegerRange(1, 30).AddDefaultValue("3").String,
			Computed:            true,
		},
		"as_encrypt_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the symmetric encryption algorithm to encrypt outbound AS messages.", "as-encrypt-alg", "").AddStringEnum("3des", "des", "rc2-128", "rc2-64", "rc2-40", "aes-128", "aes-192", "aes-256").AddDefaultValue("3des").String,
			Computed:            true,
		},
		"asmdn_request_signed_algs": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the digest algorithms to request for a signed MDN. The value can be a single algorithm or any combination of algorithms that are separated by a comma. For example, <tt>md5,sha256</tt> . The default value is <tt>sha1,md5</tt> .", "as-mdn-request-signed-algs", "").AddDefaultValue("sha1,md5").String,
			Computed:            true,
		},
		"ebms_cpa_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the CPA ID in the message that is sent to the destination partner. You can use the value to identify the sender and the recipient. This value is used when the B2B gateway is not CPA-enforced.", "ebms-cpa-id", "").String,
			Computed:            true,
		},
		"ebms_service": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of service element in outbound ebMS2 request that acts on the message. For example, <tt>urn:services:SupplierOrderProcessing</tt> . This value is used when the B2B gateway is not CPA-enforced. The value can be a string or a URI. If a non-URI string is specified, you must specify the value of <b>Service Type</b>.", "ebms-service", "").String,
			Computed:            true,
		},
		"ebms_service_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the type attribute in the ebMS SOAP message. When blank, ensure that the service value is a URI.", "ebms-service-type", "").String,
			Computed:            true,
		},
		"ebms_action": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of action element in outbound ebMS2 request. For example, <tt>NewPurchaseOrder</tt> . This value is used when the B2B gateway is not CPA-enforced.", "ebms-action", "").String,
			Computed:            true,
		},
		"ebms_send_unsigned": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Controls whether to override the signing of messages to this destination. Whether to sign outbound message is part of the configuration of the internal partner. This property cannot be used to cause a message to be signed. The default behavior is off. <ul><li>If enabled, never sign messages.</li><li>If disabled, sign messages only if the sender has signing credentials.</li></ul>", "ebms-send-unsigned", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_encrypt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encrypt the body of outgoing ebMS messages.", "ebms-encrypt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_encrypt_cert": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to encrypt outbound messages.", "ebms-encrypt-cert", "crypto_certificate").String,
			Computed:            true,
		},
		"ebms_encrypt_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the encryption algorithm to encrypt outbound ebMS2 messages.", "ebms-encrypt-alg", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc").String,
			Computed:            true,
		},
		"ebms_key_encrypt_alg": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the ebMS3 asymmetric key transport algorithm to encrypt outbound ebMS3 messages. By default, encrypts with RSA Version 1.5.", "ebms-key-encrypt-alg", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#rsa-1_5", "http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p", "http://www.w3.org/2009/xmlenc11#rsa-oaep").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#rsa-1_5").String,
			Computed:            true,
		},
		"ebms_duplicate_elimination_request": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request that the receiving partner checks for duplicates of outbound messages. If enabled, the receiving business partner should ignore duplicate message received. The default behavior is on.", "ebms-duplicate-elimination-request", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ebms_ack_request": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request an acknowledgment for outbound ebMS2 messages. The default behavior is off.", "ebms-ack-request", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_ack_request_signed": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed acknowledgment. The default behavior is off.", "ebms-ack-request-signed", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_sync_reply_mode": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the response/acknowledgment is synchronous or asynchronous. The syncReplyMode parameter indicates to the receiving partner whether to return the business response or acknowledgment in the same connection. None means asynchronous.", "ebms-syncreply-mode", "").AddStringEnum("mshSignalsOnly", "none").AddDefaultValue("none").String,
			Computed:            true,
		},
		"ebms_retry": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retransmit unacknowledged outbound messages. The default behavior is off.", "ebms-retry", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_max_retries": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of retransmit attempts. Enter a value in the range 1 - 30. The default value is 3.", "ebms-max-retries", "").AddIntegerRange(1, 30).AddDefaultValue("3").String,
			Computed:            true,
		},
		"ebms_retry_interval": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds between retry attempts. Enter a value in the range 1 - 3600. This default value is 1800.", "ebms-retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("1800").String,
			Computed:            true,
		},
		"ebms_include_time_to_live": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the sending partner includes the <tt>TimeToLive</tt> element in the outbound ebMS2 message header. The <tt>TimeToLive</tt> element indicates when the message expires. <ul><li>If the receiving partner receives the message before it expires, the receiving partner accepts the message.</li><li>If the receiving partner receives the message after it expires, the receiving partner rejects the message.</li></ul>", "ebms-include-time-to-live", "").AddDefaultValue("true").String,
			Computed:            true,
		},
		"ssl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile type to secure connections with targets.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
		},
		"ssl_client": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections with targets.", "ssl-client", "ssl_client_profile").String,
			Computed:            true,
		},
		"ebms_message_exchange_pattern": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the message exchange pattern (MEP) to use. The MEP defines how a business partner exchanges messages with another business partner. The default value is one-way push.", "ebms-mep", "").AddStringEnum("one-way-push", "one-way-pull").AddDefaultValue("one-way-push").String,
			Computed:            true,
		},
		"ebms_message_partition_channel": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When the MEP is one-way pull, specify the message partition channel (MPC) to pull messages. In the one-way pull mode, a message remains in MPC storage until the B2B gateway receives an authenticated and authorized pull request.", "ebms-mpc", "").String,
			Computed:            true,
		},
		"ebms_agreement_ref": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the reference to the agreement that governs this message exchange. This value maps to <tt>eb:AgreementRef</tt> in the message header.", "ebms-agreementref", "").String,
			Computed:            true,
		},
		"ebmsp_mode": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the PMode identifier for the convenience of PMode management. When specified, the <tt>AgreementRef/@pmode</tt> attribute value is expected in associated messages.", "ebms-pmode", "").String,
			Computed:            true,
		},
		"ebms_outbound_request_receipt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a receipt signal for a sent ebMS3 message. The default behavior is off.", "ebms-outbound-request-receipt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_outbound_request_signed_receipt": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed receipt. The default behavior is off.", "ebms-outbound-request-signed-receipt", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_outbound_receipt_reply_pattern": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the pattern to send the receipt signal. The default behavior is response.", "ebms-outbound-receipt-reply-pattern", "").AddStringEnum("Response", "Callback").AddDefaultValue("Response").String,
			Computed:            true,
		},
		"ebms_outbound_reception_awareness_notification": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to send a reception awareness error notification to the ebMS3 message producer. The default behavior is off. The B2B gateway reports a reception awareness error to the message producer if the receipt signal message is not received for a sent message.", "ebms-reception-awareness-notification", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebms_outbound_reception_awareness_timeout": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for the receipt signal. Enter a value in the range 3 - 7200. The default value is 300. If the B2B gateway does not receive the receipt signal after this duration, the B2B gateway sends a reception awareness error notification to the message producer.", "ebms-reception-awareness-timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
			Computed:            true,
		},
		"ebms_compress": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound ebMS messages. The default behavior is off.", "ebms-compress", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ebmssoap_body": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When compression is not enabled, specify whether to send messages in the SOAP <tt>Body</tt> .", "ebms-soapbody", "").AddDefaultValue("false").String,
			Computed:            true,
		},
	},
}
var DmB2BDestinationResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"dest_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name to identify this destination in the profile.", "name", "").String,
			Required:            true,
		},
		"dest_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL for the destination. The URL is parsed to determine the protocol. For load distribution, specify the name of the load balancer instead of the address-port pair. <ul><li>If the URL starts with as1:// or dpsmtp://, the service uses the ESMTP protocol. These protocols require an SMTP server connection and email address. Specify the URL in one of the following formats. <ul><li><tt>as1://mailExchange[:port][/queryParameters]</tt></li><li><tt>dpsmtp://mailExchange[:port][/queryParameters]</tt></li></ul></li><li>If the URL starts with as2://, ebms2://, ebms3://, or http://, the service uses the HTTP protocol. Specify the URL in the <tt>hostname:port/uri</tt> format.</li><li>If the URL starts with as2s://, ebms2s://, ebms3s://, or https://, the service uses the HTTPS protocol. Specify the URL in the <tt>hostname:port/uri</tt> format.</li><li>If the URL starts with as3:// or ftp://, the service uses the FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the <tt>ftp://user:password@host/%2Fpath</tt> format. %2F is the URL encoding of a forward slash.</li><li>For a relative path to the login directory of the user, specify the URL in the <tt>ftp://user:password@host:port/path</tt> format.</li></ul></li><li>If the URL starts with sftp://, the service uses the SSH FTP protocol. <ul><li>For an absolute path to the root directory, specify the URL in the <tt>sftp://host:port/path</tt> format.</li><li>For a relative path to the login directory of the user, specify the URL in the <tt>sftp://host:port/~/path</tt> format.</li></ul></li><li>If the URL starts with dpnfs://, the service uses the NFS protocol. <ul><li>For static mounts, specify the URL in the <tt>dpnfs://MountName</tt> format, where <tt>MountName</tt> is the name of an NFS mount.</li><li>For dynamic mounts, specify the URL in the <tt>dpnfs://host</tt> or <tt>dpnfs://host/path</tt> format.</li></ul></li><li>If the URL starts with dpims:// or dpimsssl://, the service uses the IMS protocol. Specify the URL in one of the following formats. <ul><li><tt>dpims://ConnectObject/?parameters</tt></li><li><tt>dpimsssl://ConnectObject/?parameters</tt></li></ul></li><li>If the URL starts with dpmq://, the service uses the IBM MQ protocol. The queue manager must exist in current domain. <ul><li>To send a message, specify the URL in the <tt>dpmq://QueueManager/URI?RequestQueue=requestQueue;queryParameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dpmq://QueueManager/URI?ReplyQueue=replyQueue;queryParameters</tt> format.</li></ul><p>If the URL starts with mq://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dpmqfte://, the service uses the IBM MQ FTE protocol. The queue manager must exist in current domain. <ul><li>To send a message, specify the URL in the <tt>dpmqfte://QueueManager/?RequestQueue=request_queue_name;queryParameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dpmqfte://QueueManager/?ReplyQueue=reply_queue_name;queryParameters</tt> format.</li></ul><p>If the URL starts with mqfte://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dptibems://, the service uses the TIBCO EMS protocol. The server must exist. <ul><li>To send a message, specify the URL in the <tt>dptibems://Server/?RequestQueue=queue;RequestReply=queue;query-parameters</tt> format.</li><li>To retrieve a message, specify the URL in the <tt>dptibems://Server/?ReplyQueue=queue;query-parameters</tt> format.</li></ul><p>If the URL starts with tibems://, the protocol is for dynamic routing.</p></li><li>If the URL starts with dpwasjms://, the service uses the WebSphere JMS protocol. The server must exist. <ul><li>To send a message to a request queue, specify the URL in the <tt>dpwasjms://Server/?RequestQueue=queue;RequestReply=queue;query-parameters</tt> format.</li><li>To send a message to a request topic space, specify the URL in the <tt>dpwasjms://Server/?RequestTopicSpace=topic-space;RequestReply=queue;query-parameters</tt> format.</li><li>To retrieve a message from a reply queue, specify the URL in the <tt>dpwasjms://Server/?ReplyQueue=queue;query-parameters</tt> format.</li><li>To retrieve a message from a reply topic space, specify the URL in the <tt>dpwasjms://Server/?ReplyTopicSpace=topic-space;query-parameters</tt> format.</li></ul></li></ul>", "dest-url", "").String,
			Required:            true,
		},
		"enabled_doc_type": GetDmB2BEnabledDocTypeResourceSchema("Specify the document types to support.", "enabled-doc-type", "", false),
		"smtp_server_connection": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SMTP server connection to send email messages. By default, the gateway uses the default SMTP server connection.", "smtp-server-connection", "smtp_server_connection").AddDefaultValue("default").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("default"),
		},
		"email_address": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination email address for messages sent to this partner.", "email-address", "").String,
			Optional:            true,
		},
		"ssh_client_connection": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH profile for SSH client connections and authentication.", "ssh-client-connection", "ssh_client_profile").String,
			Optional:            true,
		},
		"use_unique_filenames": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to generate a unique file name for puts to a remote directory.", "use-unique-filenames", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"override_timeout": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration to maintain an idle connection in seconds. Enter a value in the range 3 - 7200. The default value is 300.", "timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(3, 7200),
			},
			Default: int64default.StaticInt64(300),
		},
		"enable_ftp_settings": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to override the FTP client policy. This policy is defined in user agent of the XML manager for the B2B gateway. <ul><li>When enabled, define the overrides with the advanced AS3 or FTP settings.</li><li>When disabled, uses the original FTP client policy.</li></ul>", "enable-ftp-settings", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"user_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the username to override in the basic authentication policy of the user agent.", "username", "").String,
			Optional:            true,
		},
		"password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to override in the basic authentication policy of the user agent.", "password-alias", "password_alias").String,
			Optional:            true,
		},
		"ebmsmpc_auth_method": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify how the MPC authenticates the incoming pull requests. By default, the MPC authenticates requests through username token.", "embs-mpc-auth-method", "").AddStringEnum("username-token", "cert").AddDefaultValue("username-token").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("username-token", "cert"),
			},
			Default: stringdefault.StaticString("username-token"),
		},
		"user_name_token": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("For ebMS3, the WS-Security UsernameToken authorizes received messages. Enter the username for the wsse:UsernameToken element to package in the SOAP header of the message to send. <ul><li>For one-way push exchange pattern, specify the username token to package in messages to send.</li><li>For one-way pull exchange pattern outbound messages, when the MPC authentication method is username token, specify the username to store messages in the MPC.</li><li>For one-way pull exchange pattern inbound messages, specify the username token to package in pull requests.</li></ul>", "username-token", "").String,
			Optional:            true,
		},
		"user_name_token_password_alias": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias of the username for the wsse:UsernameToken element.", "username-token-password-alias", "password_alias").String,
			Optional:            true,
		},
		"ebmsmpc_verify_val_cred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When the MPC authentication method is validation credentials, specify the certificate to associate with messages that are submitted to the MPC.", "ebms-mpc-verify-valcred", "crypto_val_cred").String,
			Optional:            true,
		},
		"passive": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of FTP passive mode with the FTP <tt>PASV</tt> command.", "ftp-passive", "").AddStringEnum("pasv-off", "pasv-opt", "pasv-req").AddDefaultValue("pasv-req").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("pasv-off", "pasv-opt", "pasv-req"),
			},
			Default: stringdefault.StaticString("pasv-req"),
		},
		"auth_tls": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of TLS to secure connections with the FTP <tt>AUTH TLS</tt> command.", "ftp-auth-tls", "").AddStringEnum("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp").AddDefaultValue("auth-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("auth-off", "auth-tls-opt", "auth-tls-req", "auth-tls-imp"),
			},
			Default: stringdefault.StaticString("auth-off"),
		},
		"use_ccc": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the use of command encryption after authentication with the FTP <tt>CCC</tt> command.", "ftp-use-ccc", "").AddStringEnum("ccc-off", "ccc-opt", "ccc-req").AddDefaultValue("ccc-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("ccc-off", "ccc-opt", "ccc-req"),
			},
			Default: stringdefault.StaticString("ccc-off"),
		},
		"encrypt_data": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the encryption of file transfers with the FTP <tt>PROT P</tt> command.", "ftp-encrypt-data", "").AddStringEnum("enc-data-off", "enc-data-opt", "enc-data-req").AddDefaultValue("enc-data-off").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("enc-data-off", "enc-data-opt", "enc-data-req"),
			},
			Default: stringdefault.StaticString("enc-data-off"),
		},
		"data_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the data type of file transfers.", "ftp-data-type", "").AddStringEnum("ascii", "binary").AddDefaultValue("binary").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("ascii", "binary"),
			},
			Default: stringdefault.StaticString("binary"),
		},
		"slash_stou": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify how to manage unique files when the name contains a trailing slash with either the FTP <tt>STOU</tt> or <tt>STOR</tt> command. Some FTP servers do not support the <tt>STOU</tt> command.", "ftp-slash-stou", "").AddStringEnum("slash-stou-off", "slash-stou-on").AddDefaultValue("slash-stou-on").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("slash-stou-off", "slash-stou-on"),
			},
			Default: stringdefault.StaticString("slash-stou-on"),
		},
		"quoted_commands": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the FTP quoted commands list that defines the command to send to the FTP server. before each <tt>STOU</tt> , <tt>STOR</tt> , or <tt>RETR</tt> command. The commands in the list cannot be data-transfer related commands.", "ftp-quoted-commands", "ftp_quote_commands").String,
			Optional:            true,
		},
		"size_check": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to check the size of file after transfer with the FTP <tt>SIZE</tt> command. Some FTP servers, particularly vsftpd in the default configuration, provide inaccurate responses for files transferred in ASCII mode. If you get such errors, disable this feature or reconfigure vsftpd.", "ftp-size-check", "").AddStringEnum("size-check-optional", "size-check-disabled").AddDefaultValue("size-check-optional").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("size-check-optional", "size-check-disabled"),
			},
			Default: stringdefault.StaticString("size-check-optional"),
		},
		"binary_transfer_mode": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to transfer the message payload in binary mode.", "binary-transfer-mode", "").AddStringEnum("auto-detect", "enforce").AddDefaultValue("auto-detect").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("auto-detect", "enforce"),
			},
			Default: stringdefault.StaticString("auto-detect"),
		},
		"as_compress": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound AS messages. The default behavior is disabled.", "as-compress", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"as_compress_before_sign": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound AS messages before signing. The default behavior is disabled. <ul><li>If enabled, compress the message body before signing.</li><li>If disabled, sign the message before compressing the message body.</li></ul>", "as-compress-before-sign", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"as_send_unsigned": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to override the signing of messages to this destination. Whether to sign outbound message is part of the configuration of the internal partner. This property cannot be used to cause a message to be signed. The default behavior is off. <ul><li>If enabled, never sign messages.</li><li>If disabled, sign messages when the sender has signing credentials.</li></ul>", "as-send-unsigned", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"as_encrypt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encrypt the body of outgoing AS messages. The default behavior is disabled.", "as-encrypt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"as_encrypt_cert": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to encrypt outbound AS messages. Use the name of a certificate.", "as-encrypt-cert", "crypto_certificate").String,
			Optional:            true,
		},
		"asmdn_request": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request an MDN for outbound AS messages. The default behavior is disabled.", "as-mdn-request", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"asmdn_request_async": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the MDN request for outbound AS messages is asynchronous. The default behavior is disabled. <ul><li>If enabled, the MDN request is asynchronous</li><li>If disabled, the MDN request is synchronous</li></ul>", "as-mdn-request-async", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"as1mdn_redirect_email": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection email address for the MDN to outbound AS1 messages. The partner that receives the outbound AS1 message sends the MDN to this email address. When blank, the redirection email in the <tt>From</tt> address of the outgoing message (the default email address of the sending internal partner).", "as1-mdn-email", "").String,
			Optional:            true,
		},
		"as2mdn_redirect_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL for the MDN to outbound AS2 messages.", "as2-mdn-url", "").String,
			Optional:            true,
		},
		"as3mdn_redirect_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the redirection URL for the MDN to outbound AS3 messages.", "as3-mdn-url", "").String,
			Optional:            true,
		},
		"asmdn_request_signed": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed MDN instead of an unsigned one. The default behavior is an unsigned one.", "as-mdn-request-signed", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"retransmit": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retransmit messages. The default behavior is off.", "retransmit", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ack_time": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for the acknowledgment before a message is retransmitted. Enter a value in the range 1 - 3600. The default value is 1800.", "ack-time", "").AddIntegerRange(1, 3600).AddDefaultValue("1800").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 3600),
			},
			Default: int64default.StaticInt64(1800),
		},
		"max_resends": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of attempts to retransmit a message. Enter a value in the range 1 - 30. The default value is 3.", "max-resends", "").AddIntegerRange(1, 30).AddDefaultValue("3").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 30),
			},
			Default: int64default.StaticInt64(3),
		},
		"as_encrypt_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the symmetric encryption algorithm to encrypt outbound AS messages.", "as-encrypt-alg", "").AddStringEnum("3des", "des", "rc2-128", "rc2-64", "rc2-40", "aes-128", "aes-192", "aes-256").AddDefaultValue("3des").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("3des", "des", "rc2-128", "rc2-64", "rc2-40", "aes-128", "aes-192", "aes-256"),
			},
			Default: stringdefault.StaticString("3des"),
		},
		"asmdn_request_signed_algs": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the digest algorithms to request for a signed MDN. The value can be a single algorithm or any combination of algorithms that are separated by a comma. For example, <tt>md5,sha256</tt> . The default value is <tt>sha1,md5</tt> .", "as-mdn-request-signed-algs", "").AddDefaultValue("sha1,md5").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("sha1,md5"),
		},
		"ebms_cpa_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the CPA ID in the message that is sent to the destination partner. You can use the value to identify the sender and the recipient. This value is used when the B2B gateway is not CPA-enforced.", "ebms-cpa-id", "").String,
			Optional:            true,
		},
		"ebms_service": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of service element in outbound ebMS2 request that acts on the message. For example, <tt>urn:services:SupplierOrderProcessing</tt> . This value is used when the B2B gateway is not CPA-enforced. The value can be a string or a URI. If a non-URI string is specified, you must specify the value of <b>Service Type</b>.", "ebms-service", "").String,
			Optional:            true,
		},
		"ebms_service_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of the type attribute in the ebMS SOAP message. When blank, ensure that the service value is a URI.", "ebms-service-type", "").String,
			Optional:            true,
		},
		"ebms_action": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the value of action element in outbound ebMS2 request. For example, <tt>NewPurchaseOrder</tt> . This value is used when the B2B gateway is not CPA-enforced.", "ebms-action", "").String,
			Optional:            true,
		},
		"ebms_send_unsigned": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Controls whether to override the signing of messages to this destination. Whether to sign outbound message is part of the configuration of the internal partner. This property cannot be used to cause a message to be signed. The default behavior is off. <ul><li>If enabled, never sign messages.</li><li>If disabled, sign messages only if the sender has signing credentials.</li></ul>", "ebms-send-unsigned", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_encrypt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encrypt the body of outgoing ebMS messages.", "ebms-encrypt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_encrypt_cert": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the certificate to encrypt outbound messages.", "ebms-encrypt-cert", "crypto_certificate").String,
			Optional:            true,
		},
		"ebms_encrypt_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the encryption algorithm to encrypt outbound ebMS2 messages.", "ebms-encrypt-alg", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm"),
			},
			Default: stringdefault.StaticString("http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
		},
		"ebms_key_encrypt_alg": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the ebMS3 asymmetric key transport algorithm to encrypt outbound ebMS3 messages. By default, encrypts with RSA Version 1.5.", "ebms-key-encrypt-alg", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#rsa-1_5", "http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p", "http://www.w3.org/2009/xmlenc11#rsa-oaep").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#rsa-1_5").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("http://www.w3.org/2001/04/xmlenc#rsa-1_5", "http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p", "http://www.w3.org/2009/xmlenc11#rsa-oaep"),
			},
			Default: stringdefault.StaticString("http://www.w3.org/2001/04/xmlenc#rsa-1_5"),
		},
		"ebms_duplicate_elimination_request": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request that the receiving partner checks for duplicates of outbound messages. If enabled, the receiving business partner should ignore duplicate message received. The default behavior is on.", "ebms-duplicate-elimination-request", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ebms_ack_request": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request an acknowledgment for outbound ebMS2 messages. The default behavior is off.", "ebms-ack-request", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_ack_request_signed": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed acknowledgment. The default behavior is off.", "ebms-ack-request-signed", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_sync_reply_mode": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the response/acknowledgment is synchronous or asynchronous. The syncReplyMode parameter indicates to the receiving partner whether to return the business response or acknowledgment in the same connection. None means asynchronous.", "ebms-syncreply-mode", "").AddStringEnum("mshSignalsOnly", "none").AddDefaultValue("none").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("mshSignalsOnly", "none"),
			},
			Default: stringdefault.StaticString("none"),
		},
		"ebms_retry": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retransmit unacknowledged outbound messages. The default behavior is off.", "ebms-retry", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_max_retries": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of retransmit attempts. Enter a value in the range 1 - 30. The default value is 3.", "ebms-max-retries", "").AddIntegerRange(1, 30).AddDefaultValue("3").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 30),
			},
			Default: int64default.StaticInt64(3),
		},
		"ebms_retry_interval": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds between retry attempts. Enter a value in the range 1 - 3600. This default value is 1800.", "ebms-retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("1800").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(1, 3600),
			},
			Default: int64default.StaticInt64(1800),
		},
		"ebms_include_time_to_live": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the sending partner includes the <tt>TimeToLive</tt> element in the outbound ebMS2 message header. The <tt>TimeToLive</tt> element indicates when the message expires. <ul><li>If the receiving partner receives the message before it expires, the receiving partner accepts the message.</li><li>If the receiving partner receives the message after it expires, the receiving partner rejects the message.</li></ul>", "ebms-include-time-to-live", "").AddDefaultValue("true").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(true),
		},
		"ssl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of TLS profile type to secure connections with targets.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("client"),
			},
			Default: stringdefault.StaticString("client"),
		},
		"ssl_client": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections with targets.", "ssl-client", "ssl_client_profile").String,
			Optional:            true,
		},
		"ebms_message_exchange_pattern": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the message exchange pattern (MEP) to use. The MEP defines how a business partner exchanges messages with another business partner. The default value is one-way push.", "ebms-mep", "").AddStringEnum("one-way-push", "one-way-pull").AddDefaultValue("one-way-push").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("one-way-push", "one-way-pull"),
			},
			Default: stringdefault.StaticString("one-way-push"),
		},
		"ebms_message_partition_channel": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When the MEP is one-way pull, specify the message partition channel (MPC) to pull messages. In the one-way pull mode, a message remains in MPC storage until the B2B gateway receives an authenticated and authorized pull request.", "ebms-mpc", "").String,
			Optional:            true,
		},
		"ebms_agreement_ref": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the reference to the agreement that governs this message exchange. This value maps to <tt>eb:AgreementRef</tt> in the message header.", "ebms-agreementref", "").String,
			Optional:            true,
		},
		"ebmsp_mode": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the PMode identifier for the convenience of PMode management. When specified, the <tt>AgreementRef/@pmode</tt> attribute value is expected in associated messages.", "ebms-pmode", "").String,
			Optional:            true,
		},
		"ebms_outbound_request_receipt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a receipt signal for a sent ebMS3 message. The default behavior is off.", "ebms-outbound-request-receipt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_outbound_request_signed_receipt": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to request a signed receipt. The default behavior is off.", "ebms-outbound-request-signed-receipt", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_outbound_receipt_reply_pattern": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specifies the pattern to send the receipt signal. The default behavior is response.", "ebms-outbound-receipt-reply-pattern", "").AddStringEnum("Response", "Callback").AddDefaultValue("Response").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("Response", "Callback"),
			},
			Default: stringdefault.StaticString("Response"),
		},
		"ebms_outbound_reception_awareness_notification": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to send a reception awareness error notification to the ebMS3 message producer. The default behavior is off. The B2B gateway reports a reception awareness error to the message producer if the receipt signal message is not received for a sent message.", "ebms-reception-awareness-notification", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebms_outbound_reception_awareness_timeout": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait for the receipt signal. Enter a value in the range 3 - 7200. The default value is 300. If the B2B gateway does not receive the receipt signal after this duration, the B2B gateway sends a reception awareness error notification to the message producer.", "ebms-reception-awareness-timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.Int64{
				int64validator.Between(3, 7200),
			},
			Default: int64default.StaticInt64(300),
		},
		"ebms_compress": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to compress the body of outbound ebMS messages. The default behavior is off.", "ebms-compress", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ebmssoap_body": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("When compression is not enabled, specify whether to send messages in the SOAP <tt>Body</tt> .", "ebms-soapbody", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
	},
}

func (data DmB2BDestination) IsNull() bool {
	if !data.DestName.IsNull() {
		return false
	}
	if !data.DestUrl.IsNull() {
		return false
	}
	if data.EnabledDocType != nil {
		if !data.EnabledDocType.IsNull() {
			return false
		}
	}
	if !data.SmtpServerConnection.IsNull() {
		return false
	}
	if !data.EmailAddress.IsNull() {
		return false
	}
	if !data.SshClientConnection.IsNull() {
		return false
	}
	if !data.UseUniqueFilenames.IsNull() {
		return false
	}
	if !data.OverrideTimeout.IsNull() {
		return false
	}
	if !data.EnableFtpSettings.IsNull() {
		return false
	}
	if !data.UserName.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.EbmsmpcAuthMethod.IsNull() {
		return false
	}
	if !data.UserNameToken.IsNull() {
		return false
	}
	if !data.UserNameTokenPasswordAlias.IsNull() {
		return false
	}
	if !data.EbmsmpcVerifyValCred.IsNull() {
		return false
	}
	if !data.Passive.IsNull() {
		return false
	}
	if !data.AuthTls.IsNull() {
		return false
	}
	if !data.UseCcc.IsNull() {
		return false
	}
	if !data.EncryptData.IsNull() {
		return false
	}
	if !data.DataType.IsNull() {
		return false
	}
	if !data.SlashStou.IsNull() {
		return false
	}
	if !data.QuotedCommands.IsNull() {
		return false
	}
	if !data.SizeCheck.IsNull() {
		return false
	}
	if !data.BinaryTransferMode.IsNull() {
		return false
	}
	if !data.AsCompress.IsNull() {
		return false
	}
	if !data.AsCompressBeforeSign.IsNull() {
		return false
	}
	if !data.AsSendUnsigned.IsNull() {
		return false
	}
	if !data.AsEncrypt.IsNull() {
		return false
	}
	if !data.AsEncryptCert.IsNull() {
		return false
	}
	if !data.AsmdnRequest.IsNull() {
		return false
	}
	if !data.AsmdnRequestAsync.IsNull() {
		return false
	}
	if !data.As1mdnRedirectEmail.IsNull() {
		return false
	}
	if !data.As2mdnRedirectUrl.IsNull() {
		return false
	}
	if !data.As3mdnRedirectUrl.IsNull() {
		return false
	}
	if !data.AsmdnRequestSigned.IsNull() {
		return false
	}
	if !data.Retransmit.IsNull() {
		return false
	}
	if !data.AckTime.IsNull() {
		return false
	}
	if !data.MaxResends.IsNull() {
		return false
	}
	if !data.AsEncryptAlg.IsNull() {
		return false
	}
	if !data.AsmdnRequestSignedAlgs.IsNull() {
		return false
	}
	if !data.EbmsCpaId.IsNull() {
		return false
	}
	if !data.EbmsService.IsNull() {
		return false
	}
	if !data.EbmsServiceType.IsNull() {
		return false
	}
	if !data.EbmsAction.IsNull() {
		return false
	}
	if !data.EbmsSendUnsigned.IsNull() {
		return false
	}
	if !data.EbmsEncrypt.IsNull() {
		return false
	}
	if !data.EbmsEncryptCert.IsNull() {
		return false
	}
	if !data.EbmsEncryptAlg.IsNull() {
		return false
	}
	if !data.EbmsKeyEncryptAlg.IsNull() {
		return false
	}
	if !data.EbmsDuplicateEliminationRequest.IsNull() {
		return false
	}
	if !data.EbmsAckRequest.IsNull() {
		return false
	}
	if !data.EbmsAckRequestSigned.IsNull() {
		return false
	}
	if !data.EbmsSyncReplyMode.IsNull() {
		return false
	}
	if !data.EbmsRetry.IsNull() {
		return false
	}
	if !data.EbmsMaxRetries.IsNull() {
		return false
	}
	if !data.EbmsRetryInterval.IsNull() {
		return false
	}
	if !data.EbmsIncludeTimeToLive.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	if !data.EbmsMessageExchangePattern.IsNull() {
		return false
	}
	if !data.EbmsMessagePartitionChannel.IsNull() {
		return false
	}
	if !data.EbmsAgreementRef.IsNull() {
		return false
	}
	if !data.EbmspMode.IsNull() {
		return false
	}
	if !data.EbmsOutboundRequestReceipt.IsNull() {
		return false
	}
	if !data.EbmsOutboundRequestSignedReceipt.IsNull() {
		return false
	}
	if !data.EbmsOutboundReceiptReplyPattern.IsNull() {
		return false
	}
	if !data.EbmsOutboundReceptionAwarenessNotification.IsNull() {
		return false
	}
	if !data.EbmsOutboundReceptionAwarenessTimeout.IsNull() {
		return false
	}
	if !data.EbmsCompress.IsNull() {
		return false
	}
	if !data.EbmssoapBody.IsNull() {
		return false
	}
	return true
}

func (data DmB2BDestination) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.DestName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DestName`, data.DestName.ValueString())
	}
	if !data.DestUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DestURL`, data.DestUrl.ValueString())
	}
	if data.EnabledDocType != nil {
		if !data.EnabledDocType.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`EnabledDocType`, data.EnabledDocType.ToBody(ctx, ""))
		}
	}
	if !data.SmtpServerConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SMTPServerConnection`, data.SmtpServerConnection.ValueString())
	}
	if !data.EmailAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EmailAddress`, data.EmailAddress.ValueString())
	}
	if !data.SshClientConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSHClientConnection`, data.SshClientConnection.ValueString())
	}
	if !data.UseUniqueFilenames.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseUniqueFilenames`, tfutils.StringFromBool(data.UseUniqueFilenames, ""))
	}
	if !data.OverrideTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OverrideTimeout`, data.OverrideTimeout.ValueInt64())
	}
	if !data.EnableFtpSettings.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnableFTPSettings`, tfutils.StringFromBool(data.EnableFtpSettings, ""))
	}
	if !data.UserName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserName`, data.UserName.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.EbmsmpcAuthMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSMPCAuthMethod`, data.EbmsmpcAuthMethod.ValueString())
	}
	if !data.UserNameToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserNameToken`, data.UserNameToken.ValueString())
	}
	if !data.UserNameTokenPasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserNameTokenPasswordAlias`, data.UserNameTokenPasswordAlias.ValueString())
	}
	if !data.EbmsmpcVerifyValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSMPCVerifyValCred`, data.EbmsmpcVerifyValCred.ValueString())
	}
	if !data.Passive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Passive`, data.Passive.ValueString())
	}
	if !data.AuthTls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthTLS`, data.AuthTls.ValueString())
	}
	if !data.UseCcc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseCCC`, data.UseCcc.ValueString())
	}
	if !data.EncryptData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptData`, data.EncryptData.ValueString())
	}
	if !data.DataType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataType`, data.DataType.ValueString())
	}
	if !data.SlashStou.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SlashSTOU`, data.SlashStou.ValueString())
	}
	if !data.QuotedCommands.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QuotedCommands`, data.QuotedCommands.ValueString())
	}
	if !data.SizeCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SizeCheck`, data.SizeCheck.ValueString())
	}
	if !data.BinaryTransferMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BinaryTransferMode`, data.BinaryTransferMode.ValueString())
	}
	if !data.AsCompress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASCompress`, tfutils.StringFromBool(data.AsCompress, ""))
	}
	if !data.AsCompressBeforeSign.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASCompressBeforeSign`, tfutils.StringFromBool(data.AsCompressBeforeSign, ""))
	}
	if !data.AsSendUnsigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASSendUnsigned`, tfutils.StringFromBool(data.AsSendUnsigned, ""))
	}
	if !data.AsEncrypt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASEncrypt`, tfutils.StringFromBool(data.AsEncrypt, ""))
	}
	if !data.AsEncryptCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASEncryptCert`, data.AsEncryptCert.ValueString())
	}
	if !data.AsmdnRequest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASMDNRequest`, tfutils.StringFromBool(data.AsmdnRequest, ""))
	}
	if !data.AsmdnRequestAsync.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASMDNRequestAsync`, tfutils.StringFromBool(data.AsmdnRequestAsync, ""))
	}
	if !data.As1mdnRedirectEmail.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS1MDNRedirectEmail`, data.As1mdnRedirectEmail.ValueString())
	}
	if !data.As2mdnRedirectUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS2MDNRedirectURL`, data.As2mdnRedirectUrl.ValueString())
	}
	if !data.As3mdnRedirectUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS3MDNRedirectURL`, data.As3mdnRedirectUrl.ValueString())
	}
	if !data.AsmdnRequestSigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASMDNRequestSigned`, tfutils.StringFromBool(data.AsmdnRequestSigned, ""))
	}
	if !data.Retransmit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Retransmit`, tfutils.StringFromBool(data.Retransmit, ""))
	}
	if !data.AckTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ACKTime`, data.AckTime.ValueInt64())
	}
	if !data.MaxResends.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxResends`, data.MaxResends.ValueInt64())
	}
	if !data.AsEncryptAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASEncryptAlg`, data.AsEncryptAlg.ValueString())
	}
	if !data.AsmdnRequestSignedAlgs.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ASMDNRequestSignedAlgs`, data.AsmdnRequestSignedAlgs.ValueString())
	}
	if !data.EbmsCpaId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSCpaId`, data.EbmsCpaId.ValueString())
	}
	if !data.EbmsService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSService`, data.EbmsService.ValueString())
	}
	if !data.EbmsServiceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSServiceType`, data.EbmsServiceType.ValueString())
	}
	if !data.EbmsAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAction`, data.EbmsAction.ValueString())
	}
	if !data.EbmsSendUnsigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSSendUnsigned`, tfutils.StringFromBool(data.EbmsSendUnsigned, ""))
	}
	if !data.EbmsEncrypt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSEncrypt`, tfutils.StringFromBool(data.EbmsEncrypt, ""))
	}
	if !data.EbmsEncryptCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSEncryptCert`, data.EbmsEncryptCert.ValueString())
	}
	if !data.EbmsEncryptAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSEncryptAlg`, data.EbmsEncryptAlg.ValueString())
	}
	if !data.EbmsKeyEncryptAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSKeyEncryptAlg`, data.EbmsKeyEncryptAlg.ValueString())
	}
	if !data.EbmsDuplicateEliminationRequest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSDuplicateEliminationRequest`, tfutils.StringFromBool(data.EbmsDuplicateEliminationRequest, ""))
	}
	if !data.EbmsAckRequest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAckRequest`, tfutils.StringFromBool(data.EbmsAckRequest, ""))
	}
	if !data.EbmsAckRequestSigned.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAckRequestSigned`, tfutils.StringFromBool(data.EbmsAckRequestSigned, ""))
	}
	if !data.EbmsSyncReplyMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSSyncReplyMode`, data.EbmsSyncReplyMode.ValueString())
	}
	if !data.EbmsRetry.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSRetry`, tfutils.StringFromBool(data.EbmsRetry, ""))
	}
	if !data.EbmsMaxRetries.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSMaxRetries`, data.EbmsMaxRetries.ValueInt64())
	}
	if !data.EbmsRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSRetryInterval`, data.EbmsRetryInterval.ValueInt64())
	}
	if !data.EbmsIncludeTimeToLive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSIncludeTimeToLive`, tfutils.StringFromBool(data.EbmsIncludeTimeToLive, ""))
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	if !data.EbmsMessageExchangePattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSMessageExchangePattern`, data.EbmsMessageExchangePattern.ValueString())
	}
	if !data.EbmsMessagePartitionChannel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSMessagePartitionChannel`, data.EbmsMessagePartitionChannel.ValueString())
	}
	if !data.EbmsAgreementRef.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSAgreementRef`, data.EbmsAgreementRef.ValueString())
	}
	if !data.EbmspMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSPMode`, data.EbmspMode.ValueString())
	}
	if !data.EbmsOutboundRequestReceipt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundRequestReceipt`, tfutils.StringFromBool(data.EbmsOutboundRequestReceipt, ""))
	}
	if !data.EbmsOutboundRequestSignedReceipt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundRequestSignedReceipt`, tfutils.StringFromBool(data.EbmsOutboundRequestSignedReceipt, ""))
	}
	if !data.EbmsOutboundReceiptReplyPattern.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundReceiptReplyPattern`, data.EbmsOutboundReceiptReplyPattern.ValueString())
	}
	if !data.EbmsOutboundReceptionAwarenessNotification.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundReceptionAwarenessNotification`, tfutils.StringFromBool(data.EbmsOutboundReceptionAwarenessNotification, ""))
	}
	if !data.EbmsOutboundReceptionAwarenessTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSOutboundReceptionAwarenessTimeout`, data.EbmsOutboundReceptionAwarenessTimeout.ValueInt64())
	}
	if !data.EbmsCompress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSCompress`, tfutils.StringFromBool(data.EbmsCompress, ""))
	}
	if !data.EbmssoapBody.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EBMSSOAPBody`, tfutils.StringFromBool(data.EbmssoapBody, ""))
	}
	return body
}

func (data *DmB2BDestination) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `DestName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DestName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestName = types.StringNull()
	}
	if value := res.Get(pathRoot + `DestURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DestUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnabledDocType`); value.Exists() {
		data.EnabledDocType = &DmB2BEnabledDocType{}
		data.EnabledDocType.FromBody(ctx, "", value)
	} else {
		data.EnabledDocType = nil
	}
	if value := res.Get(pathRoot + `SMTPServerConnection`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SmtpServerConnection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpServerConnection = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHClientConnection`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SshClientConnection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientConnection = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseUniqueFilenames`); value.Exists() {
		data.UseUniqueFilenames = tfutils.BoolFromString(value.String())
	} else {
		data.UseUniqueFilenames = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OverrideTimeout`); value.Exists() {
		data.OverrideTimeout = types.Int64Value(value.Int())
	} else {
		data.OverrideTimeout = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `EnableFTPSettings`); value.Exists() {
		data.EnableFtpSettings = tfutils.BoolFromString(value.String())
	} else {
		data.EnableFtpSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMPCAuthMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsmpcAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsmpcAuthMethod = types.StringValue("username-token")
	}
	if value := res.Get(pathRoot + `UserNameToken`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserNameToken = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserNameToken = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserNameTokenPasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserNameTokenPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserNameTokenPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMPCVerifyValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsmpcVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsmpcVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Passive = types.StringValue("pasv-req")
	}
	if value := res.Get(pathRoot + `AuthTLS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthTls = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthTls = types.StringValue("auth-off")
	}
	if value := res.Get(pathRoot + `UseCCC`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UseCcc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UseCcc = types.StringValue("ccc-off")
	}
	if value := res.Get(pathRoot + `EncryptData`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptData = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptData = types.StringValue("enc-data-off")
	}
	if value := res.Get(pathRoot + `DataType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataType = types.StringValue("binary")
	}
	if value := res.Get(pathRoot + `SlashSTOU`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SlashStou = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SlashStou = types.StringValue("slash-stou-on")
	}
	if value := res.Get(pathRoot + `QuotedCommands`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.QuotedCommands = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommands = types.StringNull()
	}
	if value := res.Get(pathRoot + `SizeCheck`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SizeCheck = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SizeCheck = types.StringValue("size-check-optional")
	}
	if value := res.Get(pathRoot + `BinaryTransferMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BinaryTransferMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BinaryTransferMode = types.StringValue("auto-detect")
	}
	if value := res.Get(pathRoot + `ASCompress`); value.Exists() {
		data.AsCompress = tfutils.BoolFromString(value.String())
	} else {
		data.AsCompress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASCompressBeforeSign`); value.Exists() {
		data.AsCompressBeforeSign = tfutils.BoolFromString(value.String())
	} else {
		data.AsCompressBeforeSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASSendUnsigned`); value.Exists() {
		data.AsSendUnsigned = tfutils.BoolFromString(value.String())
	} else {
		data.AsSendUnsigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASEncrypt`); value.Exists() {
		data.AsEncrypt = tfutils.BoolFromString(value.String())
	} else {
		data.AsEncrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASEncryptCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AsEncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AsEncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequest`); value.Exists() {
		data.AsmdnRequest = tfutils.BoolFromString(value.String())
	} else {
		data.AsmdnRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequestAsync`); value.Exists() {
		data.AsmdnRequestAsync = tfutils.BoolFromString(value.String())
	} else {
		data.AsmdnRequestAsync = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AS1MDNRedirectEmail`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As1mdnRedirectEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1mdnRedirectEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS2MDNRedirectURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As2mdnRedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As2mdnRedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS3MDNRedirectURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As3mdnRedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As3mdnRedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequestSigned`); value.Exists() {
		data.AsmdnRequestSigned = tfutils.BoolFromString(value.String())
	} else {
		data.AsmdnRequestSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Retransmit`); value.Exists() {
		data.Retransmit = tfutils.BoolFromString(value.String())
	} else {
		data.Retransmit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ACKTime`); value.Exists() {
		data.AckTime = types.Int64Value(value.Int())
	} else {
		data.AckTime = types.Int64Value(1800)
	}
	if value := res.Get(pathRoot + `MaxResends`); value.Exists() {
		data.MaxResends = types.Int64Value(value.Int())
	} else {
		data.MaxResends = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `ASEncryptAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AsEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AsEncryptAlg = types.StringValue("3des")
	}
	if value := res.Get(pathRoot + `ASMDNRequestSignedAlgs`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AsmdnRequestSignedAlgs = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AsmdnRequestSignedAlgs = types.StringValue("sha1,md5")
	}
	if value := res.Get(pathRoot + `EBMSCpaId`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsCpaId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsCpaId = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSService`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSServiceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsServiceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsServiceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSSendUnsigned`); value.Exists() {
		data.EbmsSendUnsigned = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsSendUnsigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSEncrypt`); value.Exists() {
		data.EbmsEncrypt = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsEncrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSEncryptCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsEncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsEncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSEncryptAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsEncryptAlg = types.StringValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc")
	}
	if value := res.Get(pathRoot + `EBMSKeyEncryptAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsKeyEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsKeyEncryptAlg = types.StringValue("http://www.w3.org/2001/04/xmlenc#rsa-1_5")
	}
	if value := res.Get(pathRoot + `EBMSDuplicateEliminationRequest`); value.Exists() {
		data.EbmsDuplicateEliminationRequest = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsDuplicateEliminationRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAckRequest`); value.Exists() {
		data.EbmsAckRequest = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsAckRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAckRequestSigned`); value.Exists() {
		data.EbmsAckRequestSigned = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsAckRequestSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSSyncReplyMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsSyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsSyncReplyMode = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `EBMSRetry`); value.Exists() {
		data.EbmsRetry = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSMaxRetries`); value.Exists() {
		data.EbmsMaxRetries = types.Int64Value(value.Int())
	} else {
		data.EbmsMaxRetries = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `EBMSRetryInterval`); value.Exists() {
		data.EbmsRetryInterval = types.Int64Value(value.Int())
	} else {
		data.EbmsRetryInterval = types.Int64Value(1800)
	}
	if value := res.Get(pathRoot + `EBMSIncludeTimeToLive`); value.Exists() {
		data.EbmsIncludeTimeToLive = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsIncludeTimeToLive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMessageExchangePattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsMessageExchangePattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsMessageExchangePattern = types.StringValue("one-way-push")
	}
	if value := res.Get(pathRoot + `EBMSMessagePartitionChannel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsMessagePartitionChannel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsMessagePartitionChannel = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAgreementRef`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsAgreementRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAgreementRef = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSPMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmspMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmspMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundRequestReceipt`); value.Exists() {
		data.EbmsOutboundRequestReceipt = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutboundRequestReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundRequestSignedReceipt`); value.Exists() {
		data.EbmsOutboundRequestSignedReceipt = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutboundRequestSignedReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceiptReplyPattern`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EbmsOutboundReceiptReplyPattern = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsOutboundReceiptReplyPattern = types.StringValue("Response")
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceptionAwarenessNotification`); value.Exists() {
		data.EbmsOutboundReceptionAwarenessNotification = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsOutboundReceptionAwarenessNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceptionAwarenessTimeout`); value.Exists() {
		data.EbmsOutboundReceptionAwarenessTimeout = types.Int64Value(value.Int())
	} else {
		data.EbmsOutboundReceptionAwarenessTimeout = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `EBMSCompress`); value.Exists() {
		data.EbmsCompress = tfutils.BoolFromString(value.String())
	} else {
		data.EbmsCompress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSSOAPBody`); value.Exists() {
		data.EbmssoapBody = tfutils.BoolFromString(value.String())
	} else {
		data.EbmssoapBody = types.BoolNull()
	}
}

func (data *DmB2BDestination) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `DestName`); value.Exists() && !data.DestName.IsNull() {
		data.DestName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestName = types.StringNull()
	}
	if value := res.Get(pathRoot + `DestURL`); value.Exists() && !data.DestUrl.IsNull() {
		data.DestUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DestUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnabledDocType`); value.Exists() {
		data.EnabledDocType.UpdateFromBody(ctx, "", value)
	} else {
		data.EnabledDocType = nil
	}
	if value := res.Get(pathRoot + `SMTPServerConnection`); value.Exists() && !data.SmtpServerConnection.IsNull() {
		data.SmtpServerConnection = tfutils.ParseStringFromGJSON(value)
	} else if data.SmtpServerConnection.ValueString() != "default" {
		data.SmtpServerConnection = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && !data.EmailAddress.IsNull() {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSHClientConnection`); value.Exists() && !data.SshClientConnection.IsNull() {
		data.SshClientConnection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SshClientConnection = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseUniqueFilenames`); value.Exists() && !data.UseUniqueFilenames.IsNull() {
		data.UseUniqueFilenames = tfutils.BoolFromString(value.String())
	} else if data.UseUniqueFilenames.ValueBool() {
		data.UseUniqueFilenames = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OverrideTimeout`); value.Exists() && !data.OverrideTimeout.IsNull() {
		data.OverrideTimeout = types.Int64Value(value.Int())
	} else if data.OverrideTimeout.ValueInt64() != 300 {
		data.OverrideTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EnableFTPSettings`); value.Exists() && !data.EnableFtpSettings.IsNull() {
		data.EnableFtpSettings = tfutils.BoolFromString(value.String())
	} else if data.EnableFtpSettings.ValueBool() {
		data.EnableFtpSettings = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UserName`); value.Exists() && !data.UserName.IsNull() {
		data.UserName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserName = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMPCAuthMethod`); value.Exists() && !data.EbmsmpcAuthMethod.IsNull() {
		data.EbmsmpcAuthMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsmpcAuthMethod.ValueString() != "username-token" {
		data.EbmsmpcAuthMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserNameToken`); value.Exists() && !data.UserNameToken.IsNull() {
		data.UserNameToken = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserNameToken = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserNameTokenPasswordAlias`); value.Exists() && !data.UserNameTokenPasswordAlias.IsNull() {
		data.UserNameTokenPasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserNameTokenPasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMPCVerifyValCred`); value.Exists() && !data.EbmsmpcVerifyValCred.IsNull() {
		data.EbmsmpcVerifyValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsmpcVerifyValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `Passive`); value.Exists() && !data.Passive.IsNull() {
		data.Passive = tfutils.ParseStringFromGJSON(value)
	} else if data.Passive.ValueString() != "pasv-req" {
		data.Passive = types.StringNull()
	}
	if value := res.Get(pathRoot + `AuthTLS`); value.Exists() && !data.AuthTls.IsNull() {
		data.AuthTls = tfutils.ParseStringFromGJSON(value)
	} else if data.AuthTls.ValueString() != "auth-off" {
		data.AuthTls = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseCCC`); value.Exists() && !data.UseCcc.IsNull() {
		data.UseCcc = tfutils.ParseStringFromGJSON(value)
	} else if data.UseCcc.ValueString() != "ccc-off" {
		data.UseCcc = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptData`); value.Exists() && !data.EncryptData.IsNull() {
		data.EncryptData = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptData.ValueString() != "enc-data-off" {
		data.EncryptData = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataType`); value.Exists() && !data.DataType.IsNull() {
		data.DataType = tfutils.ParseStringFromGJSON(value)
	} else if data.DataType.ValueString() != "binary" {
		data.DataType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SlashSTOU`); value.Exists() && !data.SlashStou.IsNull() {
		data.SlashStou = tfutils.ParseStringFromGJSON(value)
	} else if data.SlashStou.ValueString() != "slash-stou-on" {
		data.SlashStou = types.StringNull()
	}
	if value := res.Get(pathRoot + `QuotedCommands`); value.Exists() && !data.QuotedCommands.IsNull() {
		data.QuotedCommands = tfutils.ParseStringFromGJSON(value)
	} else {
		data.QuotedCommands = types.StringNull()
	}
	if value := res.Get(pathRoot + `SizeCheck`); value.Exists() && !data.SizeCheck.IsNull() {
		data.SizeCheck = tfutils.ParseStringFromGJSON(value)
	} else if data.SizeCheck.ValueString() != "size-check-optional" {
		data.SizeCheck = types.StringNull()
	}
	if value := res.Get(pathRoot + `BinaryTransferMode`); value.Exists() && !data.BinaryTransferMode.IsNull() {
		data.BinaryTransferMode = tfutils.ParseStringFromGJSON(value)
	} else if data.BinaryTransferMode.ValueString() != "auto-detect" {
		data.BinaryTransferMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASCompress`); value.Exists() && !data.AsCompress.IsNull() {
		data.AsCompress = tfutils.BoolFromString(value.String())
	} else if data.AsCompress.ValueBool() {
		data.AsCompress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASCompressBeforeSign`); value.Exists() && !data.AsCompressBeforeSign.IsNull() {
		data.AsCompressBeforeSign = tfutils.BoolFromString(value.String())
	} else if data.AsCompressBeforeSign.ValueBool() {
		data.AsCompressBeforeSign = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASSendUnsigned`); value.Exists() && !data.AsSendUnsigned.IsNull() {
		data.AsSendUnsigned = tfutils.BoolFromString(value.String())
	} else if data.AsSendUnsigned.ValueBool() {
		data.AsSendUnsigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASEncrypt`); value.Exists() && !data.AsEncrypt.IsNull() {
		data.AsEncrypt = tfutils.BoolFromString(value.String())
	} else if data.AsEncrypt.ValueBool() {
		data.AsEncrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASEncryptCert`); value.Exists() && !data.AsEncryptCert.IsNull() {
		data.AsEncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AsEncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequest`); value.Exists() && !data.AsmdnRequest.IsNull() {
		data.AsmdnRequest = tfutils.BoolFromString(value.String())
	} else if data.AsmdnRequest.ValueBool() {
		data.AsmdnRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequestAsync`); value.Exists() && !data.AsmdnRequestAsync.IsNull() {
		data.AsmdnRequestAsync = tfutils.BoolFromString(value.String())
	} else if data.AsmdnRequestAsync.ValueBool() {
		data.AsmdnRequestAsync = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AS1MDNRedirectEmail`); value.Exists() && !data.As1mdnRedirectEmail.IsNull() {
		data.As1mdnRedirectEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1mdnRedirectEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS2MDNRedirectURL`); value.Exists() && !data.As2mdnRedirectUrl.IsNull() {
		data.As2mdnRedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As2mdnRedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS3MDNRedirectURL`); value.Exists() && !data.As3mdnRedirectUrl.IsNull() {
		data.As3mdnRedirectUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As3mdnRedirectUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequestSigned`); value.Exists() && !data.AsmdnRequestSigned.IsNull() {
		data.AsmdnRequestSigned = tfutils.BoolFromString(value.String())
	} else if data.AsmdnRequestSigned.ValueBool() {
		data.AsmdnRequestSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Retransmit`); value.Exists() && !data.Retransmit.IsNull() {
		data.Retransmit = tfutils.BoolFromString(value.String())
	} else if data.Retransmit.ValueBool() {
		data.Retransmit = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ACKTime`); value.Exists() && !data.AckTime.IsNull() {
		data.AckTime = types.Int64Value(value.Int())
	} else if data.AckTime.ValueInt64() != 1800 {
		data.AckTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxResends`); value.Exists() && !data.MaxResends.IsNull() {
		data.MaxResends = types.Int64Value(value.Int())
	} else if data.MaxResends.ValueInt64() != 3 {
		data.MaxResends = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ASEncryptAlg`); value.Exists() && !data.AsEncryptAlg.IsNull() {
		data.AsEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.AsEncryptAlg.ValueString() != "3des" {
		data.AsEncryptAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASMDNRequestSignedAlgs`); value.Exists() && !data.AsmdnRequestSignedAlgs.IsNull() {
		data.AsmdnRequestSignedAlgs = tfutils.ParseStringFromGJSON(value)
	} else if data.AsmdnRequestSignedAlgs.ValueString() != "sha1,md5" {
		data.AsmdnRequestSignedAlgs = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSCpaId`); value.Exists() && !data.EbmsCpaId.IsNull() {
		data.EbmsCpaId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsCpaId = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSService`); value.Exists() && !data.EbmsService.IsNull() {
		data.EbmsService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSServiceType`); value.Exists() && !data.EbmsServiceType.IsNull() {
		data.EbmsServiceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsServiceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAction`); value.Exists() && !data.EbmsAction.IsNull() {
		data.EbmsAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSSendUnsigned`); value.Exists() && !data.EbmsSendUnsigned.IsNull() {
		data.EbmsSendUnsigned = tfutils.BoolFromString(value.String())
	} else if data.EbmsSendUnsigned.ValueBool() {
		data.EbmsSendUnsigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSEncrypt`); value.Exists() && !data.EbmsEncrypt.IsNull() {
		data.EbmsEncrypt = tfutils.BoolFromString(value.String())
	} else if data.EbmsEncrypt.ValueBool() {
		data.EbmsEncrypt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSEncryptCert`); value.Exists() && !data.EbmsEncryptCert.IsNull() {
		data.EbmsEncryptCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsEncryptCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSEncryptAlg`); value.Exists() && !data.EbmsEncryptAlg.IsNull() {
		data.EbmsEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsEncryptAlg.ValueString() != "http://www.w3.org/2001/04/xmlenc#tripledes-cbc" {
		data.EbmsEncryptAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSKeyEncryptAlg`); value.Exists() && !data.EbmsKeyEncryptAlg.IsNull() {
		data.EbmsKeyEncryptAlg = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsKeyEncryptAlg.ValueString() != "http://www.w3.org/2001/04/xmlenc#rsa-1_5" {
		data.EbmsKeyEncryptAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSDuplicateEliminationRequest`); value.Exists() && !data.EbmsDuplicateEliminationRequest.IsNull() {
		data.EbmsDuplicateEliminationRequest = tfutils.BoolFromString(value.String())
	} else if !data.EbmsDuplicateEliminationRequest.ValueBool() {
		data.EbmsDuplicateEliminationRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAckRequest`); value.Exists() && !data.EbmsAckRequest.IsNull() {
		data.EbmsAckRequest = tfutils.BoolFromString(value.String())
	} else if data.EbmsAckRequest.ValueBool() {
		data.EbmsAckRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSAckRequestSigned`); value.Exists() && !data.EbmsAckRequestSigned.IsNull() {
		data.EbmsAckRequestSigned = tfutils.BoolFromString(value.String())
	} else if data.EbmsAckRequestSigned.ValueBool() {
		data.EbmsAckRequestSigned = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSSyncReplyMode`); value.Exists() && !data.EbmsSyncReplyMode.IsNull() {
		data.EbmsSyncReplyMode = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsSyncReplyMode.ValueString() != "none" {
		data.EbmsSyncReplyMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSRetry`); value.Exists() && !data.EbmsRetry.IsNull() {
		data.EbmsRetry = tfutils.BoolFromString(value.String())
	} else if data.EbmsRetry.ValueBool() {
		data.EbmsRetry = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSMaxRetries`); value.Exists() && !data.EbmsMaxRetries.IsNull() {
		data.EbmsMaxRetries = types.Int64Value(value.Int())
	} else if data.EbmsMaxRetries.ValueInt64() != 3 {
		data.EbmsMaxRetries = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EBMSRetryInterval`); value.Exists() && !data.EbmsRetryInterval.IsNull() {
		data.EbmsRetryInterval = types.Int64Value(value.Int())
	} else if data.EbmsRetryInterval.ValueInt64() != 1800 {
		data.EbmsRetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EBMSIncludeTimeToLive`); value.Exists() && !data.EbmsIncludeTimeToLive.IsNull() {
		data.EbmsIncludeTimeToLive = tfutils.BoolFromString(value.String())
	} else if !data.EbmsIncludeTimeToLive.ValueBool() {
		data.EbmsIncludeTimeToLive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMessageExchangePattern`); value.Exists() && !data.EbmsMessageExchangePattern.IsNull() {
		data.EbmsMessageExchangePattern = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsMessageExchangePattern.ValueString() != "one-way-push" {
		data.EbmsMessageExchangePattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSMessagePartitionChannel`); value.Exists() && !data.EbmsMessagePartitionChannel.IsNull() {
		data.EbmsMessagePartitionChannel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsMessagePartitionChannel = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSAgreementRef`); value.Exists() && !data.EbmsAgreementRef.IsNull() {
		data.EbmsAgreementRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmsAgreementRef = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSPMode`); value.Exists() && !data.EbmspMode.IsNull() {
		data.EbmspMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EbmspMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundRequestReceipt`); value.Exists() && !data.EbmsOutboundRequestReceipt.IsNull() {
		data.EbmsOutboundRequestReceipt = tfutils.BoolFromString(value.String())
	} else if data.EbmsOutboundRequestReceipt.ValueBool() {
		data.EbmsOutboundRequestReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundRequestSignedReceipt`); value.Exists() && !data.EbmsOutboundRequestSignedReceipt.IsNull() {
		data.EbmsOutboundRequestSignedReceipt = tfutils.BoolFromString(value.String())
	} else if data.EbmsOutboundRequestSignedReceipt.ValueBool() {
		data.EbmsOutboundRequestSignedReceipt = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceiptReplyPattern`); value.Exists() && !data.EbmsOutboundReceiptReplyPattern.IsNull() {
		data.EbmsOutboundReceiptReplyPattern = tfutils.ParseStringFromGJSON(value)
	} else if data.EbmsOutboundReceiptReplyPattern.ValueString() != "Response" {
		data.EbmsOutboundReceiptReplyPattern = types.StringNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceptionAwarenessNotification`); value.Exists() && !data.EbmsOutboundReceptionAwarenessNotification.IsNull() {
		data.EbmsOutboundReceptionAwarenessNotification = tfutils.BoolFromString(value.String())
	} else if data.EbmsOutboundReceptionAwarenessNotification.ValueBool() {
		data.EbmsOutboundReceptionAwarenessNotification = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSOutboundReceptionAwarenessTimeout`); value.Exists() && !data.EbmsOutboundReceptionAwarenessTimeout.IsNull() {
		data.EbmsOutboundReceptionAwarenessTimeout = types.Int64Value(value.Int())
	} else if data.EbmsOutboundReceptionAwarenessTimeout.ValueInt64() != 300 {
		data.EbmsOutboundReceptionAwarenessTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `EBMSCompress`); value.Exists() && !data.EbmsCompress.IsNull() {
		data.EbmsCompress = tfutils.BoolFromString(value.String())
	} else if data.EbmsCompress.ValueBool() {
		data.EbmsCompress = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EBMSSOAPBody`); value.Exists() && !data.EbmssoapBody.IsNull() {
		data.EbmssoapBody = tfutils.BoolFromString(value.String())
	} else if data.EbmssoapBody.ValueBool() {
		data.EbmssoapBody = types.BoolNull()
	}
}
