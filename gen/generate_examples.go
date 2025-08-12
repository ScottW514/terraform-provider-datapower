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

//go:build ignore

package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/scottw514/terraform-provider-datapower/testutils"
)

var endPoints = map[string]testutils.ModelTestConfig{
	"file": testutils.FileTestConfig,
	"aaajwtgenerator": testutils.AAAJWTGeneratorTestConfig,
	"aaajwtvalidator": testutils.AAAJWTValidatorTestConfig,
	"aaapolicy": testutils.AAAPolicyTestConfig,
	"amqpbroker": testutils.AMQPBrokerTestConfig,
	"amqpsourceprotocolhandler": testutils.AMQPSourceProtocolHandlerTestConfig,
	"apiapplicationtype": testutils.APIApplicationTypeTestConfig,
	"apiauthurlregistry": testutils.APIAuthURLRegistryTestConfig,
	"apicors": testutils.APICORSTestConfig,
	"apiclientidentification": testutils.APIClientIdentificationTestConfig,
	"apicollection": testutils.APICollectionTestConfig,
	"apiconnectgatewayservice": testutils.APIConnectGatewayServiceTestConfig,
	"apidefinition": testutils.APIDefinitionTestConfig,
	"apiexecute": testutils.APIExecuteTestConfig,
	"apifinal": testutils.APIFinalTestConfig,
	"apigateway": testutils.APIGatewayTestConfig,
	"apildapregistry": testutils.APILDAPRegistryTestConfig,
	"apioperation": testutils.APIOperationTestConfig,
	"apipath": testutils.APIPathTestConfig,
	"apiplan": testutils.APIPlanTestConfig,
	"apiresult": testutils.APIResultTestConfig,
	"apirouting": testutils.APIRoutingTestConfig,
	"apirule": testutils.APIRuleTestConfig,
	"apischema": testutils.APISchemaTestConfig,
	"apisecurity": testutils.APISecurityTestConfig,
	"apisecurityapikey": testutils.APISecurityAPIKeyTestConfig,
	"apisecuritybasicauth": testutils.APISecurityBasicAuthTestConfig,
	"apisecurityhttpscheme": testutils.APISecurityHTTPSchemeTestConfig,
	"apisecurityoauth": testutils.APISecurityOAuthTestConfig,
	"apisecurityoauthreq": testutils.APISecurityOAuthReqTestConfig,
	"apisecurityrequirement": testutils.APISecurityRequirementTestConfig,
	"apisecuritytokenmanager": testutils.APISecurityTokenManagerTestConfig,
	"as1pollersourceprotocolhandler": testutils.AS1PollerSourceProtocolHandlerTestConfig,
	"as2proxysourceprotocolhandler": testutils.AS2ProxySourceProtocolHandlerTestConfig,
	"as2sourceprotocolhandler": testutils.AS2SourceProtocolHandlerTestConfig,
	"as3sourceprotocolhandler": testutils.AS3SourceProtocolHandlerTestConfig,
	"accesscontrollist": testutils.AccessControlListTestConfig,
	"accessprofile": testutils.AccessProfileTestConfig,
	"analyticsendpoint": testutils.AnalyticsEndpointTestConfig,
	"appsecuritypolicy": testutils.AppSecurityPolicyTestConfig,
	"assembly": testutils.AssemblyTestConfig,
	"assemblyactionclientsecurity": testutils.AssemblyActionClientSecurityTestConfig,
	"assemblyactionextract": testutils.AssemblyActionExtractTestConfig,
	"assemblyactionfunctioncall": testutils.AssemblyActionFunctionCallTestConfig,
	"assemblyactiongatewayscript": testutils.AssemblyActionGatewayScriptTestConfig,
	"assemblyactiongraphqlcostanalysis": testutils.AssemblyActionGraphQLCostAnalysisTestConfig,
	"assemblyactiongraphqlexecute": testutils.AssemblyActionGraphQLExecuteTestConfig,
	"assemblyactiongraphqlintrospect": testutils.AssemblyActionGraphQLIntrospectTestConfig,
	"assemblyactionhtmlpage": testutils.AssemblyActionHtmlPageTestConfig,
	"assemblyactioninvoke": testutils.AssemblyActionInvokeTestConfig,
	"assemblyactionjwtgenerate": testutils.AssemblyActionJWTGenerateTestConfig,
	"assemblyactionjwtvalidate": testutils.AssemblyActionJWTValidateTestConfig,
	"assemblyactionjson2xml": testutils.AssemblyActionJson2XmlTestConfig,
	"assemblyactionlog": testutils.AssemblyActionLogTestConfig,
	"assemblyactionmap": testutils.AssemblyActionMapTestConfig,
	"assemblyactionoauth": testutils.AssemblyActionOAuthTestConfig,
	"assemblyactionparse": testutils.AssemblyActionParseTestConfig,
	"assemblyactionratelimit": testutils.AssemblyActionRateLimitTestConfig,
	"assemblyactionratelimitinfo": testutils.AssemblyActionRateLimitInfoTestConfig,
	"assemblyactionredact": testutils.AssemblyActionRedactTestConfig,
	"assemblyactionsetvar": testutils.AssemblyActionSetVarTestConfig,
	"assemblyactionthrow": testutils.AssemblyActionThrowTestConfig,
	"assemblyactionusersecurity": testutils.AssemblyActionUserSecurityTestConfig,
	"assemblyactionvalidate": testutils.AssemblyActionValidateTestConfig,
	"assemblyactionwsdl": testutils.AssemblyActionWSDLTestConfig,
	"assemblyactionwebsocketupgrade": testutils.AssemblyActionWebSocketUpgradeTestConfig,
	"assemblyactionxslt": testutils.AssemblyActionXSLTTestConfig,
	"assemblyactionxml2json": testutils.AssemblyActionXml2JsonTestConfig,
	"assemblyfunction": testutils.AssemblyFunctionTestConfig,
	"assemblylogicoperationswitch": testutils.AssemblyLogicOperationSwitchTestConfig,
	"assemblylogicswitch": testutils.AssemblyLogicSwitchTestConfig,
	"auditlog": testutils.AuditLogTestConfig,
	"b2bcpa": testutils.B2BCPATestConfig,
	"b2bcpacollaboration": testutils.B2BCPACollaborationTestConfig,
	"b2bcpareceiversetting": testutils.B2BCPAReceiverSettingTestConfig,
	"b2bcpasendersetting": testutils.B2BCPASenderSettingTestConfig,
	"b2bgateway": testutils.B2BGatewayTestConfig,
	"b2bpersistence": testutils.B2BPersistenceTestConfig,
	"b2bprofile": testutils.B2BProfileTestConfig,
	"b2bprofilegroup": testutils.B2BProfileGroupTestConfig,
	"b2bxpathroutingpolicy": testutils.B2BXPathRoutingPolicyTestConfig,
	"corspolicy": testutils.CORSPolicyTestConfig,
	"corsrule": testutils.CORSRuleTestConfig,
	"crlfetch": testutils.CRLFetchTestConfig,
	"certmonitor": testutils.CertMonitorTestConfig,
	"compileoptionspolicy": testutils.CompileOptionsPolicyTestConfig,
	"compilesettings": testutils.CompileSettingsTestConfig,
	"configdeploymentpolicy": testutils.ConfigDeploymentPolicyTestConfig,
	"configsequence": testutils.ConfigSequenceTestConfig,
	"conformancepolicy": testutils.ConformancePolicyTestConfig,
	"controllist": testutils.ControlListTestConfig,
	"cookieattributepolicy": testutils.CookieAttributePolicyTestConfig,
	"countmonitor": testutils.CountMonitorTestConfig,
	"cryptocertificate": testutils.CryptoCertificateTestConfig,
	"cryptofwcred": testutils.CryptoFWCredTestConfig,
	"cryptoidentcred": testutils.CryptoIdentCredTestConfig,
	"cryptokerberoskdc": testutils.CryptoKerberosKDCTestConfig,
	"cryptokerberoskeytab": testutils.CryptoKerberosKeytabTestConfig,
	"cryptokey": testutils.CryptoKeyTestConfig,
	"cryptosskey": testutils.CryptoSSKeyTestConfig,
	"cryptovalcred": testutils.CryptoValCredTestConfig,
	"dnsnameservice": testutils.DNSNameServiceTestConfig,
	"deploymentpolicyparametersbinding": testutils.DeploymentPolicyParametersBindingTestConfig,
	"distributedvariable": testutils.DistributedVariableTestConfig,
	"documentcryptomap": testutils.DocumentCryptoMapTestConfig,
	"domain": testutils.DomainTestConfig,
	"domainavailability": testutils.DomainAvailabilityTestConfig,
	"domainsettings": testutils.DomainSettingsTestConfig,
	"durationmonitor": testutils.DurationMonitorTestConfig,
	"ebms2sourceprotocolhandler": testutils.EBMS2SourceProtocolHandlerTestConfig,
	"ebms3sourceprotocolhandler": testutils.EBMS3SourceProtocolHandlerTestConfig,
	"errorreportsettings": testutils.ErrorReportSettingsTestConfig,
	"ftpfilepollersourceprotocolhandler": testutils.FTPFilePollerSourceProtocolHandlerTestConfig,
	"ftpquotecommands": testutils.FTPQuoteCommandsTestConfig,
	"ftpserversourceprotocolhandler": testutils.FTPServerSourceProtocolHandlerTestConfig,
	"filesystemusagemonitor": testutils.FileSystemUsageMonitorTestConfig,
	"filteraction": testutils.FilterActionTestConfig,
	"formsloginpolicy": testutils.FormsLoginPolicyTestConfig,
	"gwsremotedebug": testutils.GWSRemoteDebugTestConfig,
	"gwscriptsettings": testutils.GWScriptSettingsTestConfig,
	"gatewaypeering": testutils.GatewayPeeringTestConfig,
	"gatewaypeeringgroup": testutils.GatewayPeeringGroupTestConfig,
	"gatewaypeeringmanager": testutils.GatewayPeeringManagerTestConfig,
	"gitops": testutils.GitOpsTestConfig,
	"gitopstemplate": testutils.GitOpsTemplateTestConfig,
	"gitopsvariables": testutils.GitOpsVariablesTestConfig,
	"graphqlschemaoptions": testutils.GraphQLSchemaOptionsTestConfig,
	"httpinputconversionmap": testutils.HTTPInputConversionMapTestConfig,
	"httpssourceprotocolhandler": testutils.HTTPSSourceProtocolHandlerTestConfig,
	"httpsourceprotocolhandler": testutils.HTTPSourceProtocolHandlerTestConfig,
	"httpuseragent": testutils.HTTPUserAgentTestConfig,
	"hostalias": testutils.HostAliasTestConfig,
	"importpackage": testutils.ImportPackageTestConfig,
	"includeconfig": testutils.IncludeConfigTestConfig,
	"interopservice": testutils.InteropServiceTestConfig,
	"joserecipientidentifier": testutils.JOSERecipientIdentifierTestConfig,
	"josesignatureidentifier": testutils.JOSESignatureIdentifierTestConfig,
	"jsonsettings": testutils.JSONSettingsTestConfig,
	"jweheader": testutils.JWEHeaderTestConfig,
	"jwerecipient": testutils.JWERecipientTestConfig,
	"jwssignature": testutils.JWSSignatureTestConfig,
	"kafkacluster": testutils.KafkaClusterTestConfig,
	"kafkasourceprotocolhandler": testutils.KafkaSourceProtocolHandlerTestConfig,
	"ldapconnectionpool": testutils.LDAPConnectionPoolTestConfig,
	"ldapsearchparameters": testutils.LDAPSearchParametersTestConfig,
	"loadbalancergroup": testutils.LoadBalancerGroupTestConfig,
	"loglabel": testutils.LogLabelTestConfig,
	"logtarget": testutils.LogTargetTestConfig,
	"luna": testutils.LunaTestConfig,
	"lunahagroup": testutils.LunaHAGroupTestConfig,
	"lunahasettings": testutils.LunaHASettingsTestConfig,
	"lunapartition": testutils.LunaPartitionTestConfig,
	"mcfcustomrule": testutils.MCFCustomRuleTestConfig,
	"mcfhttpheader": testutils.MCFHttpHeaderTestConfig,
	"mcfhttpmethod": testutils.MCFHttpMethodTestConfig,
	"mcfhttpurl": testutils.MCFHttpURLTestConfig,
	"mcfxpath": testutils.MCFXPathTestConfig,
	"mpgwerroraction": testutils.MPGWErrorActionTestConfig,
	"mpgwerrorhandlingpolicy": testutils.MPGWErrorHandlingPolicyTestConfig,
	"mqmanager": testutils.MQManagerTestConfig,
	"mqmanagergroup": testutils.MQManagerGroupTestConfig,
	"mqv9plusmftsourceprotocolhandler": testutils.MQv9PlusMFTSourceProtocolHandlerTestConfig,
	"mqv9plussourceprotocolhandler": testutils.MQv9PlusSourceProtocolHandlerTestConfig,
	"mtompolicy": testutils.MTOMPolicyTestConfig,
	"matching": testutils.MatchingTestConfig,
	"messagecontentfilters": testutils.MessageContentFiltersTestConfig,
	"messagematching": testutils.MessageMatchingTestConfig,
	"messagetype": testutils.MessageTypeTestConfig,
	"mgmtinterface": testutils.MgmtInterfaceTestConfig,
	"multiprotocolgateway": testutils.MultiProtocolGatewayTestConfig,
	"nfsclientsettings": testutils.NFSClientSettingsTestConfig,
	"nfsdynamicmounts": testutils.NFSDynamicMountsTestConfig,
	"nfsfilepollersourceprotocolhandler": testutils.NFSFilePollerSourceProtocolHandlerTestConfig,
	"nfsstaticmount": testutils.NFSStaticMountTestConfig,
	"namevalueprofile": testutils.NameValueProfileTestConfig,
	"oauthprovidersettings": testutils.OAuthProviderSettingsTestConfig,
	"oauthsupportedclient": testutils.OAuthSupportedClientTestConfig,
	"oauthsupportedclientgroup": testutils.OAuthSupportedClientGroupTestConfig,
	"odr": testutils.ODRTestConfig,
	"odrconnectorgroup": testutils.ODRConnectorGroupTestConfig,
	"opentelemetry": testutils.OpenTelemetryTestConfig,
	"opentelemetryexporter": testutils.OpenTelemetryExporterTestConfig,
	"opentelemetrysampler": testutils.OpenTelemetrySamplerTestConfig,
	"operationratelimit": testutils.OperationRateLimitTestConfig,
	"poppollersourceprotocolhandler": testutils.POPPollerSourceProtocolHandlerTestConfig,
	"parsesettings": testutils.ParseSettingsTestConfig,
	"passwordalias": testutils.PasswordAliasTestConfig,
	"peergroup": testutils.PeerGroupTestConfig,
	"policyattachments": testutils.PolicyAttachmentsTestConfig,
	"policyparameters": testutils.PolicyParametersTestConfig,
	"probe": testutils.ProbeTestConfig,
	"processingmetadata": testutils.ProcessingMetadataTestConfig,
	"quotaenforcementserver": testutils.QuotaEnforcementServerTestConfig,
	"radiussettings": testutils.RADIUSSettingsTestConfig,
	"rbmsettings": testutils.RBMSettingsTestConfig,
	"raidvolume": testutils.RaidVolumeTestConfig,
	"ratelimitconfiguration": testutils.RateLimitConfigurationTestConfig,
	"ratelimitdefinition": testutils.RateLimitDefinitionTestConfig,
	"ratelimitdefinitiongroup": testutils.RateLimitDefinitionGroupTestConfig,
	"samlattributes": testutils.SAMLAttributesTestConfig,
	"sftpfilepollersourceprotocolhandler": testutils.SFTPFilePollerSourceProtocolHandlerTestConfig,
	"slmaction": testutils.SLMActionTestConfig,
	"slmcredclass": testutils.SLMCredClassTestConfig,
	"slmpolicy": testutils.SLMPolicyTestConfig,
	"slmrsrcclass": testutils.SLMRsrcClassTestConfig,
	"slmschedule": testutils.SLMScheduleTestConfig,
	"smtpserverconnection": testutils.SMTPServerConnectionTestConfig,
	"snmpsettings": testutils.SNMPSettingsTestConfig,
	"soapheaderdisposition": testutils.SOAPHeaderDispositionTestConfig,
	"sqldatasource": testutils.SQLDataSourceTestConfig,
	"sshclientprofile": testutils.SSHClientProfileTestConfig,
	"sshdomainclientprofile": testutils.SSHDomainClientProfileTestConfig,
	"sshserverprofile": testutils.SSHServerProfileTestConfig,
	"sshserversourceprotocolhandler": testutils.SSHServerSourceProtocolHandlerTestConfig,
	"sshservice": testutils.SSHServiceTestConfig,
	"sslclientprofile": testutils.SSLClientProfileTestConfig,
	"sslproxyservice": testutils.SSLProxyServiceTestConfig,
	"sslsnimapping": testutils.SSLSNIMappingTestConfig,
	"sslsniserverprofile": testutils.SSLSNIServerProfileTestConfig,
	"sslserverprofile": testutils.SSLServerProfileTestConfig,
	"schemaexceptionmap": testutils.SchemaExceptionMapTestConfig,
	"securebackupmode": testutils.SecureBackupModeTestConfig,
	"socialloginpolicy": testutils.SocialLoginPolicyTestConfig,
	"statelesstcpsourceprotocolhandler": testutils.StatelessTCPSourceProtocolHandlerTestConfig,
	"statistics": testutils.StatisticsTestConfig,
	"stylepolicy": testutils.StylePolicyTestConfig,
	"stylepolicyaction": testutils.StylePolicyActionTestConfig,
	"stylepolicyrule": testutils.StylePolicyRuleTestConfig,
	"systemsettings": testutils.SystemSettingsTestConfig,
	"tam": testutils.TAMTestConfig,
	"tcpproxyservice": testutils.TCPProxyServiceTestConfig,
	"throttler": testutils.ThrottlerTestConfig,
	"timesettings": testutils.TimeSettingsTestConfig,
	"urlmap": testutils.URLMapTestConfig,
	"urlrefreshpolicy": testutils.URLRefreshPolicyTestConfig,
	"urlrewritepolicy": testutils.URLRewritePolicyTestConfig,
	"user": testutils.UserTestConfig,
	"usergroup": testutils.UserGroupTestConfig,
	"visibilitylist": testutils.VisibilityListTestConfig,
	"wccservice": testutils.WCCServiceTestConfig,
	"wsendpointrewritepolicy": testutils.WSEndpointRewritePolicyTestConfig,
	"wsgateway": testutils.WSGatewayTestConfig,
	"wsrrsavedsearchsubscription": testutils.WSRRSavedSearchSubscriptionTestConfig,
	"wsrrserver": testutils.WSRRServerTestConfig,
	"wsrrsubscription": testutils.WSRRSubscriptionTestConfig,
	"wsstylepolicy": testutils.WSStylePolicyTestConfig,
	"wsstylepolicyrule": testutils.WSStylePolicyRuleTestConfig,
	"wxsgrid": testutils.WXSGridTestConfig,
	"webapperrorhandlingpolicy": testutils.WebAppErrorHandlingPolicyTestConfig,
	"webappfw": testutils.WebAppFWTestConfig,
	"webapprequest": testutils.WebAppRequestTestConfig,
	"webappresponse": testutils.WebAppResponseTestConfig,
	"webappsessionpolicy": testutils.WebAppSessionPolicyTestConfig,
	"webb2bviewer": testutils.WebB2BViewerTestConfig,
	"webgui": testutils.WebGUITestConfig,
	"webservicemonitor": testutils.WebServiceMonitorTestConfig,
	"webservicesagent": testutils.WebServicesAgentTestConfig,
	"webspherejmsserver": testutils.WebSphereJMSServerTestConfig,
	"webspherejmssourceprotocolhandler": testutils.WebSphereJMSSourceProtocolHandlerTestConfig,
	"webtokenservice": testutils.WebTokenServiceTestConfig,
	"xacmlpdp": testutils.XACMLPDPTestConfig,
	"xmlfirewallservice": testutils.XMLFirewallServiceTestConfig,
	"xmlmanager": testutils.XMLManagerTestConfig,
	"xpathroutingmap": testutils.XPathRoutingMapTestConfig,
	"xslcoprocservice": testutils.XSLCoprocServiceTestConfig,
	"xslproxyservice": testutils.XSLProxyServiceTestConfig,
	"xtcprotocolhandler": testutils.XTCProtocolHandlerTestConfig,
	"zosnssclient": testutils.ZosNSSClientTestConfig,
}

func writeOutput(outputPath, content string) {
	// Create directories if needed.
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	// Write to output file.
	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
}

func main() {
	if err := os.MkdirAll("./examples/test_bed", 0755); err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}
	f, err := os.Create("./examples/test_bed/resources.tf")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	if _, err := f.WriteString(`terraform {
  required_providers {
    datapower = {
      source  = "scottw514/datapower"
    }
  }
}`); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}
	defer f.Close()

	resources := make([]string, 0, len(endPoints))
	for k := range endPoints {
		resources = append(resources, k)
	}
	sort.Strings(resources)

	for _, k := range resources {
		if endPoints[k].Data != "" {
			writeOutput("./examples/data-sources/datapower_"+k+"/data-source.tf", endPoints[k].Data)
		}
		writeOutput("./examples/resources/datapower_"+k+"/resource.tf", endPoints[k].Resource)
		if endPoints[k].TestBed != "" {
			if _, err := f.WriteString(endPoints[k].TestBed); err != nil {
				log.Fatalf("Error writing output file: %v", err)
			}
		}
	}
}
