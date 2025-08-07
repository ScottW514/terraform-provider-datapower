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

	"github.com/scottw514/terraform-provider-datapower/internal/provider/testconfig"
)

var endPoints = map[string]testconfig.ModelTestConfig{
	"file": testconfig.FileTestConfig,
	"aaajwtgenerator": testconfig.AAAJWTGeneratorTestConfig,
	"aaajwtvalidator": testconfig.AAAJWTValidatorTestConfig,
	"aaapolicy": testconfig.AAAPolicyTestConfig,
	"amqpbroker": testconfig.AMQPBrokerTestConfig,
	"amqpsourceprotocolhandler": testconfig.AMQPSourceProtocolHandlerTestConfig,
	"apiapplicationtype": testconfig.APIApplicationTypeTestConfig,
	"apiauthurlregistry": testconfig.APIAuthURLRegistryTestConfig,
	"apicors": testconfig.APICORSTestConfig,
	"apiclientidentification": testconfig.APIClientIdentificationTestConfig,
	"apicollection": testconfig.APICollectionTestConfig,
	"apiconnectgatewayservice": testconfig.APIConnectGatewayServiceTestConfig,
	"apidefinition": testconfig.APIDefinitionTestConfig,
	"apiexecute": testconfig.APIExecuteTestConfig,
	"apifinal": testconfig.APIFinalTestConfig,
	"apigateway": testconfig.APIGatewayTestConfig,
	"apildapregistry": testconfig.APILDAPRegistryTestConfig,
	"apioperation": testconfig.APIOperationTestConfig,
	"apipath": testconfig.APIPathTestConfig,
	"apiplan": testconfig.APIPlanTestConfig,
	"apiresult": testconfig.APIResultTestConfig,
	"apirouting": testconfig.APIRoutingTestConfig,
	"apirule": testconfig.APIRuleTestConfig,
	"apischema": testconfig.APISchemaTestConfig,
	"apisecurity": testconfig.APISecurityTestConfig,
	"apisecurityapikey": testconfig.APISecurityAPIKeyTestConfig,
	"apisecuritybasicauth": testconfig.APISecurityBasicAuthTestConfig,
	"apisecurityhttpscheme": testconfig.APISecurityHTTPSchemeTestConfig,
	"apisecurityoauth": testconfig.APISecurityOAuthTestConfig,
	"apisecurityoauthreq": testconfig.APISecurityOAuthReqTestConfig,
	"apisecurityrequirement": testconfig.APISecurityRequirementTestConfig,
	"apisecuritytokenmanager": testconfig.APISecurityTokenManagerTestConfig,
	"as1pollersourceprotocolhandler": testconfig.AS1PollerSourceProtocolHandlerTestConfig,
	"as2proxysourceprotocolhandler": testconfig.AS2ProxySourceProtocolHandlerTestConfig,
	"as2sourceprotocolhandler": testconfig.AS2SourceProtocolHandlerTestConfig,
	"as3sourceprotocolhandler": testconfig.AS3SourceProtocolHandlerTestConfig,
	"accesscontrollist": testconfig.AccessControlListTestConfig,
	"accessprofile": testconfig.AccessProfileTestConfig,
	"analyticsendpoint": testconfig.AnalyticsEndpointTestConfig,
	"appsecuritypolicy": testconfig.AppSecurityPolicyTestConfig,
	"assembly": testconfig.AssemblyTestConfig,
	"assemblyactionclientsecurity": testconfig.AssemblyActionClientSecurityTestConfig,
	"assemblyactionextract": testconfig.AssemblyActionExtractTestConfig,
	"assemblyactionfunctioncall": testconfig.AssemblyActionFunctionCallTestConfig,
	"assemblyactiongatewayscript": testconfig.AssemblyActionGatewayScriptTestConfig,
	"assemblyactiongraphqlcostanalysis": testconfig.AssemblyActionGraphQLCostAnalysisTestConfig,
	"assemblyactiongraphqlexecute": testconfig.AssemblyActionGraphQLExecuteTestConfig,
	"assemblyactiongraphqlintrospect": testconfig.AssemblyActionGraphQLIntrospectTestConfig,
	"assemblyactionhtmlpage": testconfig.AssemblyActionHtmlPageTestConfig,
	"assemblyactioninvoke": testconfig.AssemblyActionInvokeTestConfig,
	"assemblyactionjwtgenerate": testconfig.AssemblyActionJWTGenerateTestConfig,
	"assemblyactionjwtvalidate": testconfig.AssemblyActionJWTValidateTestConfig,
	"assemblyactionjson2xml": testconfig.AssemblyActionJson2XmlTestConfig,
	"assemblyactionlog": testconfig.AssemblyActionLogTestConfig,
	"assemblyactionmap": testconfig.AssemblyActionMapTestConfig,
	"assemblyactionoauth": testconfig.AssemblyActionOAuthTestConfig,
	"assemblyactionparse": testconfig.AssemblyActionParseTestConfig,
	"assemblyactionratelimit": testconfig.AssemblyActionRateLimitTestConfig,
	"assemblyactionratelimitinfo": testconfig.AssemblyActionRateLimitInfoTestConfig,
	"assemblyactionredact": testconfig.AssemblyActionRedactTestConfig,
	"assemblyactionsetvar": testconfig.AssemblyActionSetVarTestConfig,
	"assemblyactionthrow": testconfig.AssemblyActionThrowTestConfig,
	"assemblyactionusersecurity": testconfig.AssemblyActionUserSecurityTestConfig,
	"assemblyactionvalidate": testconfig.AssemblyActionValidateTestConfig,
	"assemblyactionwsdl": testconfig.AssemblyActionWSDLTestConfig,
	"assemblyactionwebsocketupgrade": testconfig.AssemblyActionWebSocketUpgradeTestConfig,
	"assemblyactionxslt": testconfig.AssemblyActionXSLTTestConfig,
	"assemblyactionxml2json": testconfig.AssemblyActionXml2JsonTestConfig,
	"assemblyfunction": testconfig.AssemblyFunctionTestConfig,
	"assemblylogicoperationswitch": testconfig.AssemblyLogicOperationSwitchTestConfig,
	"assemblylogicswitch": testconfig.AssemblyLogicSwitchTestConfig,
	"auditlog": testconfig.AuditLogTestConfig,
	"b2bcpa": testconfig.B2BCPATestConfig,
	"b2bcpacollaboration": testconfig.B2BCPACollaborationTestConfig,
	"b2bcpareceiversetting": testconfig.B2BCPAReceiverSettingTestConfig,
	"b2bcpasendersetting": testconfig.B2BCPASenderSettingTestConfig,
	"b2bgateway": testconfig.B2BGatewayTestConfig,
	"b2bpersistence": testconfig.B2BPersistenceTestConfig,
	"b2bprofile": testconfig.B2BProfileTestConfig,
	"b2bprofilegroup": testconfig.B2BProfileGroupTestConfig,
	"b2bxpathroutingpolicy": testconfig.B2BXPathRoutingPolicyTestConfig,
	"corspolicy": testconfig.CORSPolicyTestConfig,
	"corsrule": testconfig.CORSRuleTestConfig,
	"crlfetch": testconfig.CRLFetchTestConfig,
	"certmonitor": testconfig.CertMonitorTestConfig,
	"compileoptionspolicy": testconfig.CompileOptionsPolicyTestConfig,
	"compilesettings": testconfig.CompileSettingsTestConfig,
	"configdeploymentpolicy": testconfig.ConfigDeploymentPolicyTestConfig,
	"configsequence": testconfig.ConfigSequenceTestConfig,
	"conformancepolicy": testconfig.ConformancePolicyTestConfig,
	"controllist": testconfig.ControlListTestConfig,
	"cookieattributepolicy": testconfig.CookieAttributePolicyTestConfig,
	"countmonitor": testconfig.CountMonitorTestConfig,
	"cryptocertificate": testconfig.CryptoCertificateTestConfig,
	"cryptofwcred": testconfig.CryptoFWCredTestConfig,
	"cryptoidentcred": testconfig.CryptoIdentCredTestConfig,
	"cryptokerberoskdc": testconfig.CryptoKerberosKDCTestConfig,
	"cryptokerberoskeytab": testconfig.CryptoKerberosKeytabTestConfig,
	"cryptokey": testconfig.CryptoKeyTestConfig,
	"cryptosskey": testconfig.CryptoSSKeyTestConfig,
	"cryptovalcred": testconfig.CryptoValCredTestConfig,
	"dnsnameservice": testconfig.DNSNameServiceTestConfig,
	"deploymentpolicyparametersbinding": testconfig.DeploymentPolicyParametersBindingTestConfig,
	"distributedvariable": testconfig.DistributedVariableTestConfig,
	"documentcryptomap": testconfig.DocumentCryptoMapTestConfig,
	"domain": testconfig.DomainTestConfig,
	"domainavailability": testconfig.DomainAvailabilityTestConfig,
	"domainsettings": testconfig.DomainSettingsTestConfig,
	"durationmonitor": testconfig.DurationMonitorTestConfig,
	"ebms2sourceprotocolhandler": testconfig.EBMS2SourceProtocolHandlerTestConfig,
	"ebms3sourceprotocolhandler": testconfig.EBMS3SourceProtocolHandlerTestConfig,
	"errorreportsettings": testconfig.ErrorReportSettingsTestConfig,
	"ftpfilepollersourceprotocolhandler": testconfig.FTPFilePollerSourceProtocolHandlerTestConfig,
	"ftpquotecommands": testconfig.FTPQuoteCommandsTestConfig,
	"ftpserversourceprotocolhandler": testconfig.FTPServerSourceProtocolHandlerTestConfig,
	"filesystemusagemonitor": testconfig.FileSystemUsageMonitorTestConfig,
	"filteraction": testconfig.FilterActionTestConfig,
	"formsloginpolicy": testconfig.FormsLoginPolicyTestConfig,
	"gwsremotedebug": testconfig.GWSRemoteDebugTestConfig,
	"gwscriptsettings": testconfig.GWScriptSettingsTestConfig,
	"gatewaypeering": testconfig.GatewayPeeringTestConfig,
	"gatewaypeeringgroup": testconfig.GatewayPeeringGroupTestConfig,
	"gatewaypeeringmanager": testconfig.GatewayPeeringManagerTestConfig,
	"gitops": testconfig.GitOpsTestConfig,
	"gitopstemplate": testconfig.GitOpsTemplateTestConfig,
	"gitopsvariables": testconfig.GitOpsVariablesTestConfig,
	"graphqlschemaoptions": testconfig.GraphQLSchemaOptionsTestConfig,
	"httpinputconversionmap": testconfig.HTTPInputConversionMapTestConfig,
	"httpssourceprotocolhandler": testconfig.HTTPSSourceProtocolHandlerTestConfig,
	"httpsourceprotocolhandler": testconfig.HTTPSourceProtocolHandlerTestConfig,
	"httpuseragent": testconfig.HTTPUserAgentTestConfig,
	"hostalias": testconfig.HostAliasTestConfig,
	"importpackage": testconfig.ImportPackageTestConfig,
	"includeconfig": testconfig.IncludeConfigTestConfig,
	"interopservice": testconfig.InteropServiceTestConfig,
	"joserecipientidentifier": testconfig.JOSERecipientIdentifierTestConfig,
	"josesignatureidentifier": testconfig.JOSESignatureIdentifierTestConfig,
	"jsonsettings": testconfig.JSONSettingsTestConfig,
	"jweheader": testconfig.JWEHeaderTestConfig,
	"jwerecipient": testconfig.JWERecipientTestConfig,
	"jwssignature": testconfig.JWSSignatureTestConfig,
	"kafkacluster": testconfig.KafkaClusterTestConfig,
	"kafkasourceprotocolhandler": testconfig.KafkaSourceProtocolHandlerTestConfig,
	"ldapconnectionpool": testconfig.LDAPConnectionPoolTestConfig,
	"ldapsearchparameters": testconfig.LDAPSearchParametersTestConfig,
	"loadbalancergroup": testconfig.LoadBalancerGroupTestConfig,
	"loglabel": testconfig.LogLabelTestConfig,
	"logtarget": testconfig.LogTargetTestConfig,
	"luna": testconfig.LunaTestConfig,
	"lunahagroup": testconfig.LunaHAGroupTestConfig,
	"lunahasettings": testconfig.LunaHASettingsTestConfig,
	"lunapartition": testconfig.LunaPartitionTestConfig,
	"mcfcustomrule": testconfig.MCFCustomRuleTestConfig,
	"mcfhttpheader": testconfig.MCFHttpHeaderTestConfig,
	"mcfhttpmethod": testconfig.MCFHttpMethodTestConfig,
	"mcfhttpurl": testconfig.MCFHttpURLTestConfig,
	"mcfxpath": testconfig.MCFXPathTestConfig,
	"mpgwerroraction": testconfig.MPGWErrorActionTestConfig,
	"mpgwerrorhandlingpolicy": testconfig.MPGWErrorHandlingPolicyTestConfig,
	"mqmanager": testconfig.MQManagerTestConfig,
	"mqmanagergroup": testconfig.MQManagerGroupTestConfig,
	"mqv9plusmftsourceprotocolhandler": testconfig.MQv9PlusMFTSourceProtocolHandlerTestConfig,
	"mqv9plussourceprotocolhandler": testconfig.MQv9PlusSourceProtocolHandlerTestConfig,
	"mtompolicy": testconfig.MTOMPolicyTestConfig,
	"matching": testconfig.MatchingTestConfig,
	"messagecontentfilters": testconfig.MessageContentFiltersTestConfig,
	"messagematching": testconfig.MessageMatchingTestConfig,
	"messagetype": testconfig.MessageTypeTestConfig,
	"mgmtinterface": testconfig.MgmtInterfaceTestConfig,
	"multiprotocolgateway": testconfig.MultiProtocolGatewayTestConfig,
	"nfsclientsettings": testconfig.NFSClientSettingsTestConfig,
	"nfsdynamicmounts": testconfig.NFSDynamicMountsTestConfig,
	"nfsfilepollersourceprotocolhandler": testconfig.NFSFilePollerSourceProtocolHandlerTestConfig,
	"nfsstaticmount": testconfig.NFSStaticMountTestConfig,
	"namevalueprofile": testconfig.NameValueProfileTestConfig,
	"oauthprovidersettings": testconfig.OAuthProviderSettingsTestConfig,
	"oauthsupportedclient": testconfig.OAuthSupportedClientTestConfig,
	"oauthsupportedclientgroup": testconfig.OAuthSupportedClientGroupTestConfig,
	"odr": testconfig.ODRTestConfig,
	"odrconnectorgroup": testconfig.ODRConnectorGroupTestConfig,
	"opentelemetry": testconfig.OpenTelemetryTestConfig,
	"opentelemetryexporter": testconfig.OpenTelemetryExporterTestConfig,
	"opentelemetrysampler": testconfig.OpenTelemetrySamplerTestConfig,
	"operationratelimit": testconfig.OperationRateLimitTestConfig,
	"poppollersourceprotocolhandler": testconfig.POPPollerSourceProtocolHandlerTestConfig,
	"parsesettings": testconfig.ParseSettingsTestConfig,
	"passwordalias": testconfig.PasswordAliasTestConfig,
	"peergroup": testconfig.PeerGroupTestConfig,
	"policyattachments": testconfig.PolicyAttachmentsTestConfig,
	"policyparameters": testconfig.PolicyParametersTestConfig,
	"probe": testconfig.ProbeTestConfig,
	"processingmetadata": testconfig.ProcessingMetadataTestConfig,
	"quotaenforcementserver": testconfig.QuotaEnforcementServerTestConfig,
	"radiussettings": testconfig.RADIUSSettingsTestConfig,
	"rbmsettings": testconfig.RBMSettingsTestConfig,
	"raidvolume": testconfig.RaidVolumeTestConfig,
	"ratelimitconfiguration": testconfig.RateLimitConfigurationTestConfig,
	"ratelimitdefinition": testconfig.RateLimitDefinitionTestConfig,
	"ratelimitdefinitiongroup": testconfig.RateLimitDefinitionGroupTestConfig,
	"samlattributes": testconfig.SAMLAttributesTestConfig,
	"sftpfilepollersourceprotocolhandler": testconfig.SFTPFilePollerSourceProtocolHandlerTestConfig,
	"slmaction": testconfig.SLMActionTestConfig,
	"slmcredclass": testconfig.SLMCredClassTestConfig,
	"slmpolicy": testconfig.SLMPolicyTestConfig,
	"slmrsrcclass": testconfig.SLMRsrcClassTestConfig,
	"slmschedule": testconfig.SLMScheduleTestConfig,
	"smtpserverconnection": testconfig.SMTPServerConnectionTestConfig,
	"snmpsettings": testconfig.SNMPSettingsTestConfig,
	"soapheaderdisposition": testconfig.SOAPHeaderDispositionTestConfig,
	"sqldatasource": testconfig.SQLDataSourceTestConfig,
	"sshclientprofile": testconfig.SSHClientProfileTestConfig,
	"sshdomainclientprofile": testconfig.SSHDomainClientProfileTestConfig,
	"sshserverprofile": testconfig.SSHServerProfileTestConfig,
	"sshserversourceprotocolhandler": testconfig.SSHServerSourceProtocolHandlerTestConfig,
	"sshservice": testconfig.SSHServiceTestConfig,
	"sslclientprofile": testconfig.SSLClientProfileTestConfig,
	"sslproxyservice": testconfig.SSLProxyServiceTestConfig,
	"sslsnimapping": testconfig.SSLSNIMappingTestConfig,
	"sslsniserverprofile": testconfig.SSLSNIServerProfileTestConfig,
	"sslserverprofile": testconfig.SSLServerProfileTestConfig,
	"schemaexceptionmap": testconfig.SchemaExceptionMapTestConfig,
	"securebackupmode": testconfig.SecureBackupModeTestConfig,
	"socialloginpolicy": testconfig.SocialLoginPolicyTestConfig,
	"statelesstcpsourceprotocolhandler": testconfig.StatelessTCPSourceProtocolHandlerTestConfig,
	"statistics": testconfig.StatisticsTestConfig,
	"stylepolicy": testconfig.StylePolicyTestConfig,
	"stylepolicyaction": testconfig.StylePolicyActionTestConfig,
	"stylepolicyrule": testconfig.StylePolicyRuleTestConfig,
	"systemsettings": testconfig.SystemSettingsTestConfig,
	"tam": testconfig.TAMTestConfig,
	"tcpproxyservice": testconfig.TCPProxyServiceTestConfig,
	"throttler": testconfig.ThrottlerTestConfig,
	"timesettings": testconfig.TimeSettingsTestConfig,
	"urlmap": testconfig.URLMapTestConfig,
	"urlrefreshpolicy": testconfig.URLRefreshPolicyTestConfig,
	"urlrewritepolicy": testconfig.URLRewritePolicyTestConfig,
	"user": testconfig.UserTestConfig,
	"usergroup": testconfig.UserGroupTestConfig,
	"visibilitylist": testconfig.VisibilityListTestConfig,
	"wccservice": testconfig.WCCServiceTestConfig,
	"wsendpointrewritepolicy": testconfig.WSEndpointRewritePolicyTestConfig,
	"wsgateway": testconfig.WSGatewayTestConfig,
	"wsrrsavedsearchsubscription": testconfig.WSRRSavedSearchSubscriptionTestConfig,
	"wsrrserver": testconfig.WSRRServerTestConfig,
	"wsrrsubscription": testconfig.WSRRSubscriptionTestConfig,
	"wsstylepolicy": testconfig.WSStylePolicyTestConfig,
	"wsstylepolicyrule": testconfig.WSStylePolicyRuleTestConfig,
	"wxsgrid": testconfig.WXSGridTestConfig,
	"webapperrorhandlingpolicy": testconfig.WebAppErrorHandlingPolicyTestConfig,
	"webappfw": testconfig.WebAppFWTestConfig,
	"webapprequest": testconfig.WebAppRequestTestConfig,
	"webappresponse": testconfig.WebAppResponseTestConfig,
	"webappsessionpolicy": testconfig.WebAppSessionPolicyTestConfig,
	"webb2bviewer": testconfig.WebB2BViewerTestConfig,
	"webgui": testconfig.WebGUITestConfig,
	"webservicemonitor": testconfig.WebServiceMonitorTestConfig,
	"webservicesagent": testconfig.WebServicesAgentTestConfig,
	"webspherejmsserver": testconfig.WebSphereJMSServerTestConfig,
	"webspherejmssourceprotocolhandler": testconfig.WebSphereJMSSourceProtocolHandlerTestConfig,
	"webtokenservice": testconfig.WebTokenServiceTestConfig,
	"xacmlpdp": testconfig.XACMLPDPTestConfig,
	"xmlfirewallservice": testconfig.XMLFirewallServiceTestConfig,
	"xmlmanager": testconfig.XMLManagerTestConfig,
	"xpathroutingmap": testconfig.XPathRoutingMapTestConfig,
	"xslcoprocservice": testconfig.XSLCoprocServiceTestConfig,
	"xslproxyservice": testconfig.XSLProxyServiceTestConfig,
	"xtcprotocolhandler": testconfig.XTCProtocolHandlerTestConfig,
	"zosnssclient": testconfig.ZosNSSClientTestConfig,
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
