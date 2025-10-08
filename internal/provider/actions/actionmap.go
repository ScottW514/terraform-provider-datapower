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

package actions

type action struct {
	Body string
}

type target struct {
	ObjectName   string
	ValidActions map[string]action
}

const domainQuiesceBody = `{"DomainQuiesce": {"name": "{domain}", "timeout": 60}}`
const domainUnquiesceBody = `{"DomainUnquiesce": {"name": "{domain}"}}`
const serviceQuiesceBody = `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`
const serviceUnquiesceBody = `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`

var actionMap = map[string]target{
	"datapower_aaa_policy": {
		ObjectName: "AAAPolicy",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushAAACache": {"PolicyName": "{name}"}}`,
			},
		},
	},
	"datapower_amqp_source_protocol_handler": {
		ObjectName: "AMQPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_api_collection": {
		ObjectName: "APICollection",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushAPISubscriberCache": {"APICollection": "{name}"}}`,
			},
		},
	},
	"datapower_api_gateway": {
		ObjectName: "APIGateway",
		ValidActions: map[string]action{
			"flush_stylesheet_cache": {
				Body: `{"FlushAPIGWStylesheetCache": {"APIGateway": "{name}", "MatchPattern": "*"}}`,
			},
			"flush_document_cache": {
				Body: `{"FlushAPIGWDocumentCache": {"APIGateway": "{name}", "MatchPattern": "*"}}`,
			},
		},
	},
	"datapower_as1_poller_source_protocol_handler": {
		ObjectName: "AS1PollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as2_proxy_source_protocol_handler": {
		ObjectName: "AS2ProxySourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as2_source_protocol_handler": {
		ObjectName: "AS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as3_source_protocol_handler": {
		ObjectName: "AS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_b2b_gateway": {
		ObjectName: "B2BGateway",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_domain": {
		ObjectName: "Domain",
		ValidActions: map[string]action{
			"quiesce": {},
			"restart": {
				Body: `{"RestartDomain": {"Domain": "{domain}"}}`,
			},
		},
	},
	"datapower_ebms2_source_protocol_handler": {
		ObjectName: "EBMS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ebms3_source_protocol_handler": {
		ObjectName: "EBMS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ftp_file_poller_source_protocol_handler": {
		ObjectName: "FTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ftp_server_source_protocol_handler": {
		ObjectName: "FTPServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_https_source_protocol_handler": {
		ObjectName: "HTTPSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_http_source_protocol_handler": {
		ObjectName: "HTTPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_kafka_source_protocol_handler": {
		ObjectName: "KafkaSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_mqv9_plus_mft_source_protocol_handler": {
		ObjectName: "MQv9PlusMFTSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_mqv9_plus_source_protocol_handler": {
		ObjectName: "MQv9PlusSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_multi_protocol_gateway": {
		ObjectName: "MultiProtocolGateway",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_nfs_file_poller_source_protocol_handler": {
		ObjectName: "NFSFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_pop_poller_source_protocol_handler": {
		ObjectName: "POPPollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_rbm_settings": {
		ObjectName: "RBMSettings",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushRBMCache": 0}`,
			},
		},
	},
	"datapower_sftp_file_poller_source_protocol_handler": {
		ObjectName: "SFTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ssh_server_source_protocol_handler": {
		ObjectName: "SSHServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ssl_proxy_service": {
		ObjectName: "SSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_stateless_tcp_source_protocol_handler": {
		ObjectName: "StatelessTCPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_tcp_proxy_service": {
		ObjectName: "TCPProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ws_gateway": {
		ObjectName: "WSGateway",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_web_app_fw": {
		ObjectName: "WebAppFW",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_web_sphere_jms_source_protocol_handler": {
		ObjectName: "WebSphereJMSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_web_token_service": {
		ObjectName: "WebTokenService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xml_firewall_service": {
		ObjectName: "XMLFirewallService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xml_manager": {
		ObjectName: "XMLManager",
		ValidActions: map[string]action{
			"flush_stylesheet_cache": {
				Body: `{"FlushStylesheetCache": {"XMLManager": "{name}"}}`,
			},
			"flush_document_cache": {
				Body: `{"FlushDocumentCache": {"XMLManager": "{name}"}}`,
			},
			"flush_ldap_pool_cache": {
				Body: `{"FlushLDAPPoolCache": {"XMLManager": "{name}"}}`,
			},
		},
	},
	"datapower_xsl_coproc_service": {
		ObjectName: "XSLCoprocService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xsl_proxy_service": {
		ObjectName: "XSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xtc_protocol_handler": {
		ObjectName: "XTCProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
}
