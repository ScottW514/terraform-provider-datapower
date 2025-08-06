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
	PreBody  string
	PostBody string
}

type target struct {
	ObjectName   string
	ValidActions map[string]action
}

var actionMap = map[string]target{
	"aaapolicy": {
		ObjectName: "AAAPolicy",
		ValidActions: map[string]action{
			"flush_cache": {
				PostBody: `{"FlushAAACache": {"PolicyName": "{name}"}}`,
			},
		},
	},
	"amqpsourceprotocolhandler": {
		ObjectName: "AMQPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"apicollection": {
		ObjectName: "APICollection",
		ValidActions: map[string]action{
			"flush_cache": {
				PostBody: `{"FlushAPISubscriberCache": {"APICollection": "{name}"}}`,
			},
		},
	},
	"apigateway": {
		ObjectName: "APIGateway",
		ValidActions: map[string]action{
			"flush_stylesheet_cache": {
				PostBody: `{"FlushAPIGWStylesheetCache": {"APIGateway": "{name}", "MatchPattern": "*"}}`,
			},
			"flush_document_cache": {
				PostBody: `{"FlushAPIGWDocumentCache": {"APIGateway": "{name}", "MatchPattern": "*"}}`,
			},
		},
	},
	"as1pollersourceprotocolhandler": {
		ObjectName: "AS1PollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"as2proxysourceprotocolhandler": {
		ObjectName: "AS2ProxySourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"as2sourceprotocolhandler": {
		ObjectName: "AS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"as3sourceprotocolhandler": {
		ObjectName: "AS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"b2bgateway": {
		ObjectName: "B2BGateway",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"domain": {
		ObjectName: "Domain",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"DomainQuiesce": {"name": "{domain}", "timeout": 60}}`,
				PostBody: `{"DomainUnquiesce": {"name": "{domain}"}}`,
			},
			"restart": {
				PostBody: `{"RestartDomain": {"Domain": "{domain}"}}`,
			},
		},
	},
	"ebms2sourceprotocolhandler": {
		ObjectName: "EBMS2SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"ebms3sourceprotocolhandler": {
		ObjectName: "EBMS3SourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"ftpfilepollersourceprotocolhandler": {
		ObjectName: "FTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"ftpserversourceprotocolhandler": {
		ObjectName: "FTPServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"httpssourceprotocolhandler": {
		ObjectName: "HTTPSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"httpsourceprotocolhandler": {
		ObjectName: "HTTPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"kafkasourceprotocolhandler": {
		ObjectName: "KafkaSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"mqv9plusmftsourceprotocolhandler": {
		ObjectName: "MQv9PlusMFTSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"mqv9plussourceprotocolhandler": {
		ObjectName: "MQv9PlusSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"multiprotocolgateway": {
		ObjectName: "MultiProtocolGateway",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"nfsfilepollersourceprotocolhandler": {
		ObjectName: "NFSFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"poppollersourceprotocolhandler": {
		ObjectName: "POPPollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"rbmsettings": {
		ObjectName: "RBMSettings",
		ValidActions: map[string]action{
			"flush_cache": {
				PostBody: `{"FlushRBMCache": 0}`,
			},
		},
	},
	"sftpfilepollersourceprotocolhandler": {
		ObjectName: "SFTPFilePollerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"sshserversourceprotocolhandler": {
		ObjectName: "SSHServerSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"sslproxyservice": {
		ObjectName: "SSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"statelesstcpsourceprotocolhandler": {
		ObjectName: "StatelessTCPSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"tcpproxyservice": {
		ObjectName: "TCPProxyService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"wsgateway": {
		ObjectName: "WSGateway",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"webappfw": {
		ObjectName: "WebAppFW",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"webspherejmssourceprotocolhandler": {
		ObjectName: "WebSphereJMSSourceProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"webtokenservice": {
		ObjectName: "WebTokenService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"xmlfirewallservice": {
		ObjectName: "XMLFirewallService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"xmlmanager": {
		ObjectName: "XMLManager",
		ValidActions: map[string]action{
			"flush_stylesheet_cache": {
				PostBody: `{"FlushStylesheetCache": {"XMLManager": "{name}"}}`,
			},
			"flush_document_cache": {
				PostBody: `{"FlushDocumentCache": {"XMLManager": "{name}"}}`,
			},
			"flush_ldap_pool_cache": {
				PostBody: `{"FlushLDAPPoolCache": {"XMLManager": "{name}"}}`,
			},
		},
	},
	"xslcoprocservice": {
		ObjectName: "XSLCoprocService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"xslproxyservice": {
		ObjectName: "XSLProxyService",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
	"xtcprotocolhandler": {
		ObjectName: "XTCProtocolHandler",
		ValidActions: map[string]action{
			"quiesce": {
				PreBody:  `{"ServiceQuiesce": {"type": "{type}", "name": "{name}", "timeout": 60}}`,
				PostBody: `{"ServiceUnquiesce": {"type": "{type}", "name": "{name}"}}`,
			},
		},
	},
}
