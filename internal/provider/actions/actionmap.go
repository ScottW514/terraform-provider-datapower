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
	"datapower_aaapolicy": {
		ObjectName: "AAAPolicy",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushAAACache": {"PolicyName": "{name}"}}`,
			},
		},
	},
	"datapower_amqpsourceprotocolhandler": {
		ObjectName: "AMQPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_apicollection": {
		ObjectName: "APICollection",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushAPISubscriberCache": {"APICollection": "{name}"}}`,
			},
		},
	},
	"datapower_apigateway": {
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
	"datapower_as1pollersourceprotocolhandler": {
		ObjectName: "AS1PollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as2proxysourceprotocolhandler": {
		ObjectName: "AS2ProxySourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as2sourceprotocolhandler": {
		ObjectName: "AS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_as3sourceprotocolhandler": {
		ObjectName: "AS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_b2bgateway": {
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
	"datapower_ebms2sourceprotocolhandler": {
		ObjectName: "EBMS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ebms3sourceprotocolhandler": {
		ObjectName: "EBMS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ftpfilepollersourceprotocolhandler": {
		ObjectName: "FTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_ftpserversourceprotocolhandler": {
		ObjectName: "FTPServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_httpssourceprotocolhandler": {
		ObjectName: "HTTPSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_httpsourceprotocolhandler": {
		ObjectName: "HTTPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_kafkasourceprotocolhandler": {
		ObjectName: "KafkaSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_mqv9plusmftsourceprotocolhandler": {
		ObjectName: "MQv9PlusMFTSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_mqv9plussourceprotocolhandler": {
		ObjectName: "MQv9PlusSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_multiprotocolgateway": {
		ObjectName: "MultiProtocolGateway",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_nfsfilepollersourceprotocolhandler": {
		ObjectName: "NFSFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_poppollersourceprotocolhandler": {
		ObjectName: "POPPollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_rbmsettings": {
		ObjectName: "RBMSettings",
		ValidActions: map[string]action{
			"flush_cache": {
				Body: `{"FlushRBMCache": 0}`,
			},
		},
	},
	"datapower_sftpfilepollersourceprotocolhandler": {
		ObjectName: "SFTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_sshserversourceprotocolhandler": {
		ObjectName: "SSHServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_sslproxyservice": {
		ObjectName: "SSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_statelesstcpsourceprotocolhandler": {
		ObjectName: "StatelessTCPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_tcpproxyservice": {
		ObjectName: "TCPProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_wsgateway": {
		ObjectName: "WSGateway",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_webappfw": {
		ObjectName: "WebAppFW",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_webspherejmssourceprotocolhandler": {
		ObjectName: "WebSphereJMSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_webtokenservice": {
		ObjectName: "WebTokenService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xmlfirewallservice": {
		ObjectName: "XMLFirewallService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xmlmanager": {
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
	"datapower_xslcoprocservice": {
		ObjectName: "XSLCoprocService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xslproxyservice": {
		ObjectName: "XSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
	"datapower_xtcprotocolhandler": {
		ObjectName: "XTCProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {},
		},
	},
}
