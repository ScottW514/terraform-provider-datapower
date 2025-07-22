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
	"accesscontrollist": testconfig.AccessControlListTestConfig,
	"accessprofile": testconfig.AccessProfileTestConfig,
	"auditlog": testconfig.AuditLogTestConfig,
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
	"deploymentpolicyparametersbinding": testconfig.DeploymentPolicyParametersBindingTestConfig,
	"domain": testconfig.DomainTestConfig,
	"domainavailability": testconfig.DomainAvailabilityTestConfig,
	"durationmonitor": testconfig.DurationMonitorTestConfig,
	"ftpquotecommands": testconfig.FTPQuoteCommandsTestConfig,
	"filteraction": testconfig.FilterActionTestConfig,
	"formsloginpolicy": testconfig.FormsLoginPolicyTestConfig,
	"httpinputconversionmap": testconfig.HTTPInputConversionMapTestConfig,
	"httpssourceprotocolhandler": testconfig.HTTPSSourceProtocolHandlerTestConfig,
	"httpsourceprotocolhandler": testconfig.HTTPSourceProtocolHandlerTestConfig,
	"httpuseragent": testconfig.HTTPUserAgentTestConfig,
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
	"multiprotocolgateway": testconfig.MultiProtocolGatewayTestConfig,
	"oauthsupportedclient": testconfig.OAuthSupportedClientTestConfig,
	"oauthsupportedclientgroup": testconfig.OAuthSupportedClientGroupTestConfig,
	"parsesettings": testconfig.ParseSettingsTestConfig,
	"passwordalias": testconfig.PasswordAliasTestConfig,
	"policyattachments": testconfig.PolicyAttachmentsTestConfig,
	"policyparameters": testconfig.PolicyParametersTestConfig,
	"processingmetadata": testconfig.ProcessingMetadataTestConfig,
	"samlattributes": testconfig.SAMLAttributesTestConfig,
	"sqldatasource": testconfig.SQLDataSourceTestConfig,
	"sshclientprofile": testconfig.SSHClientProfileTestConfig,
	"sslclientprofile": testconfig.SSLClientProfileTestConfig,
	"sslsnimapping": testconfig.SSLSNIMappingTestConfig,
	"sslsniserverprofile": testconfig.SSLSNIServerProfileTestConfig,
	"sslserverprofile": testconfig.SSLServerProfileTestConfig,
	"socialloginpolicy": testconfig.SocialLoginPolicyTestConfig,
	"stylepolicy": testconfig.StylePolicyTestConfig,
	"stylepolicyaction": testconfig.StylePolicyActionTestConfig,
	"stylepolicyrule": testconfig.StylePolicyRuleTestConfig,
	"tam": testconfig.TAMTestConfig,
	"tfimendpoint": testconfig.TFIMEndpointTestConfig,
	"urlmap": testconfig.URLMapTestConfig,
	"urlrefreshpolicy": testconfig.URLRefreshPolicyTestConfig,
	"urlrewritepolicy": testconfig.URLRewritePolicyTestConfig,
	"user": testconfig.UserTestConfig,
	"usergroup": testconfig.UserGroupTestConfig,
	"wccservice": testconfig.WCCServiceTestConfig,
	"wsrrsavedsearchsubscription": testconfig.WSRRSavedSearchSubscriptionTestConfig,
	"wsrrserver": testconfig.WSRRServerTestConfig,
	"wsrrsubscription": testconfig.WSRRSubscriptionTestConfig,
	"webservicemonitor": testconfig.WebServiceMonitorTestConfig,
	"xacmlpdp": testconfig.XACMLPDPTestConfig,
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
