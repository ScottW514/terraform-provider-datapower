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

	"github.com/scottw514/terraform-provider-datapower/internal/provider/models/testconfig"
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
	"apiratelimit": testconfig.APIRateLimitTestConfig,
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
	"crlfetch": testconfig.CRLFetchTestConfig,
	"certmonitor": testconfig.CertMonitorTestConfig,
	"compileoptionspolicy": testconfig.CompileOptionsPolicyTestConfig,
	"configdeploymentpolicy": testconfig.ConfigDeploymentPolicyTestConfig,
	"cookieattributepolicy": testconfig.CookieAttributePolicyTestConfig,
	"countmonitor": testconfig.CountMonitorTestConfig,
	"cryptocertificate": testconfig.CryptoCertificateTestConfig,
	"cryptofwcred": testconfig.CryptoFWCredTestConfig,
	"cryptoidentcred": testconfig.CryptoIdentCredTestConfig,
	"cryptokerberoskeytab": testconfig.CryptoKerberosKeytabTestConfig,
	"cryptokey": testconfig.CryptoKeyTestConfig,
	"cryptosskey": testconfig.CryptoSSKeyTestConfig,
	"cryptovalcred": testconfig.CryptoValCredTestConfig,
	"dnsnameservice": testconfig.DNSNameServiceTestConfig,
	"deploymentpolicyparametersbinding": testconfig.DeploymentPolicyParametersBindingTestConfig,
	"distributedvariable": testconfig.DistributedVariableTestConfig,
	"domain": testconfig.DomainTestConfig,
	"domainavailability": testconfig.DomainAvailabilityTestConfig,
	"domainsettings": testconfig.DomainSettingsTestConfig,
	"durationmonitor": testconfig.DurationMonitorTestConfig,
	"errorreportsettings": testconfig.ErrorReportSettingsTestConfig,
	"ftpquotecommands": testconfig.FTPQuoteCommandsTestConfig,
	"filesystemusagemonitor": testconfig.FileSystemUsageMonitorTestConfig,
	"filteraction": testconfig.FilterActionTestConfig,
	"formsloginpolicy": testconfig.FormsLoginPolicyTestConfig,
	"gwsremotedebug": testconfig.GWSRemoteDebugTestConfig,
	"gwscriptsettings": testconfig.GWScriptSettingsTestConfig,
	"gatewaypeering": testconfig.GatewayPeeringTestConfig,
	"gatewaypeeringmanager": testconfig.GatewayPeeringManagerTestConfig,
	"gitops": testconfig.GitOpsTestConfig,
	"gitopsvariables": testconfig.GitOpsVariablesTestConfig,
	"httpinputconversionmap": testconfig.HTTPInputConversionMapTestConfig,
	"httpssourceprotocolhandler": testconfig.HTTPSSourceProtocolHandlerTestConfig,
	"httpsourceprotocolhandler": testconfig.HTTPSourceProtocolHandlerTestConfig,
	"httpuseragent": testconfig.HTTPUserAgentTestConfig,
	"interopservice": testconfig.InteropServiceTestConfig,
	"joserecipientidentifier": testconfig.JOSERecipientIdentifierTestConfig,
	"josesignatureidentifier": testconfig.JOSESignatureIdentifierTestConfig,
	"jsonsettings": testconfig.JSONSettingsTestConfig,
	"jweheader": testconfig.JWEHeaderTestConfig,
	"jwerecipient": testconfig.JWERecipientTestConfig,
	"jwssignature": testconfig.JWSSignatureTestConfig,
	"ldapconnectionpool": testconfig.LDAPConnectionPoolTestConfig,
	"ldapsearchparameters": testconfig.LDAPSearchParametersTestConfig,
	"loadbalancergroup": testconfig.LoadBalancerGroupTestConfig,
	"loglabel": testconfig.LogLabelTestConfig,
	"mpgwerroraction": testconfig.MPGWErrorActionTestConfig,
	"mpgwerrorhandlingpolicy": testconfig.MPGWErrorHandlingPolicyTestConfig,
	"matching": testconfig.MatchingTestConfig,
	"messagecontentfilters": testconfig.MessageContentFiltersTestConfig,
	"messagematching": testconfig.MessageMatchingTestConfig,
	"messagetype": testconfig.MessageTypeTestConfig,
	"mgmtinterface": testconfig.MgmtInterfaceTestConfig,
	"multiprotocolgateway": testconfig.MultiProtocolGatewayTestConfig,
	"nfsclientsettings": testconfig.NFSClientSettingsTestConfig,
	"nfsdynamicmounts": testconfig.NFSDynamicMountsTestConfig,
	"namevalueprofile": testconfig.NameValueProfileTestConfig,
	"oauthsupportedclient": testconfig.OAuthSupportedClientTestConfig,
	"oauthsupportedclientgroup": testconfig.OAuthSupportedClientGroupTestConfig,
	"odr": testconfig.ODRTestConfig,
	"parsesettings": testconfig.ParseSettingsTestConfig,
	"passwordalias": testconfig.PasswordAliasTestConfig,
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
	"samlattributes": testconfig.SAMLAttributesTestConfig,
	"snmpsettings": testconfig.SNMPSettingsTestConfig,
	"sqldatasource": testconfig.SQLDataSourceTestConfig,
	"sshclientprofile": testconfig.SSHClientProfileTestConfig,
	"sshdomainclientprofile": testconfig.SSHDomainClientProfileTestConfig,
	"sshserverprofile": testconfig.SSHServerProfileTestConfig,
	"sshservice": testconfig.SSHServiceTestConfig,
	"sslclientprofile": testconfig.SSLClientProfileTestConfig,
	"sslsnimapping": testconfig.SSLSNIMappingTestConfig,
	"sslsniserverprofile": testconfig.SSLSNIServerProfileTestConfig,
	"sslserverprofile": testconfig.SSLServerProfileTestConfig,
	"socialloginpolicy": testconfig.SocialLoginPolicyTestConfig,
	"statistics": testconfig.StatisticsTestConfig,
	"stylepolicy": testconfig.StylePolicyTestConfig,
	"stylepolicyaction": testconfig.StylePolicyActionTestConfig,
	"stylepolicyrule": testconfig.StylePolicyRuleTestConfig,
	"systemsettings": testconfig.SystemSettingsTestConfig,
	"tam": testconfig.TAMTestConfig,
	"tfimendpoint": testconfig.TFIMEndpointTestConfig,
	"throttler": testconfig.ThrottlerTestConfig,
	"timesettings": testconfig.TimeSettingsTestConfig,
	"urlmap": testconfig.URLMapTestConfig,
	"urlrefreshpolicy": testconfig.URLRefreshPolicyTestConfig,
	"urlrewritepolicy": testconfig.URLRewritePolicyTestConfig,
	"user": testconfig.UserTestConfig,
	"usergroup": testconfig.UserGroupTestConfig,
	"wccservice": testconfig.WCCServiceTestConfig,
	"wsrrsavedsearchsubscription": testconfig.WSRRSavedSearchSubscriptionTestConfig,
	"wsrrserver": testconfig.WSRRServerTestConfig,
	"wsrrsubscription": testconfig.WSRRSubscriptionTestConfig,
	"webapperrorhandlingpolicy": testconfig.WebAppErrorHandlingPolicyTestConfig,
	"webappfw": testconfig.WebAppFWTestConfig,
	"webapprequest": testconfig.WebAppRequestTestConfig,
	"webappresponse": testconfig.WebAppResponseTestConfig,
	"webappsessionpolicy": testconfig.WebAppSessionPolicyTestConfig,
	"webb2bviewer": testconfig.WebB2BViewerTestConfig,
	"webgui": testconfig.WebGUITestConfig,
	"webservicemonitor": testconfig.WebServiceMonitorTestConfig,
	"webservicesagent": testconfig.WebServicesAgentTestConfig,
	"xacmlpdp": testconfig.XACMLPDPTestConfig,
	"xmlfirewallservice": testconfig.XMLFirewallServiceTestConfig,
	"xmlmanager": testconfig.XMLManagerTestConfig,
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
	for k, v := range endPoints {
		if v.Data != "" {
			writeOutput("./examples/data-sources/datapower_"+k+"/data-source.tf", v.Data)
		}
		writeOutput("./examples/resources/datapower_"+k+"/resource.tf", v.Resource)
	}
}
