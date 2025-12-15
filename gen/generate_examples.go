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
	"aaa_jwt_generator": testutils.AAAJWTGeneratorTestConfig,
	"aaa_jwt_validator": testutils.AAAJWTValidatorTestConfig,
	"aaa_policy": testutils.AAAPolicyTestConfig,
	"amqp_broker": testutils.AMQPBrokerTestConfig,
	"amqp_source_protocol_handler": testutils.AMQPSourceProtocolHandlerTestConfig,
	"api_application_type": testutils.APIApplicationTypeTestConfig,
	"api_auth_url_registry": testutils.APIAuthURLRegistryTestConfig,
	"api_cors": testutils.APICORSTestConfig,
	"api_client_identification": testutils.APIClientIdentificationTestConfig,
	"api_collection": testutils.APICollectionTestConfig,
	"api_connect_gateway_service": testutils.APIConnectGatewayServiceTestConfig,
	"api_definition": testutils.APIDefinitionTestConfig,
	"api_execute": testutils.APIExecuteTestConfig,
	"api_final": testutils.APIFinalTestConfig,
	"api_gateway": testutils.APIGatewayTestConfig,
	"api_ldap_registry": testutils.APILDAPRegistryTestConfig,
	"api_operation": testutils.APIOperationTestConfig,
	"api_path": testutils.APIPathTestConfig,
	"api_plan": testutils.APIPlanTestConfig,
	"api_result": testutils.APIResultTestConfig,
	"api_routing": testutils.APIRoutingTestConfig,
	"api_rule": testutils.APIRuleTestConfig,
	"api_schema": testutils.APISchemaTestConfig,
	"api_security": testutils.APISecurityTestConfig,
	"api_security_api_key": testutils.APISecurityAPIKeyTestConfig,
	"api_security_basic_auth": testutils.APISecurityBasicAuthTestConfig,
	"api_security_http_scheme": testutils.APISecurityHTTPSchemeTestConfig,
	"api_security_oauth": testutils.APISecurityOAuthTestConfig,
	"api_security_oauth_req": testutils.APISecurityOAuthReqTestConfig,
	"api_security_requirement": testutils.APISecurityRequirementTestConfig,
	"api_security_token_manager": testutils.APISecurityTokenManagerTestConfig,
	"as1_poller_source_protocol_handler": testutils.AS1PollerSourceProtocolHandlerTestConfig,
	"as2_proxy_source_protocol_handler": testutils.AS2ProxySourceProtocolHandlerTestConfig,
	"as2_source_protocol_handler": testutils.AS2SourceProtocolHandlerTestConfig,
	"as3_source_protocol_handler": testutils.AS3SourceProtocolHandlerTestConfig,
	"access_control_list": testutils.AccessControlListTestConfig,
	"access_profile": testutils.AccessProfileTestConfig,
	"analytics_endpoint": testutils.AnalyticsEndpointTestConfig,
	"app_security_policy": testutils.AppSecurityPolicyTestConfig,
	"assembly": testutils.AssemblyTestConfig,
	"assembly_action_client_security": testutils.AssemblyActionClientSecurityTestConfig,
	"assembly_action_extract": testutils.AssemblyActionExtractTestConfig,
	"assembly_action_function_call": testutils.AssemblyActionFunctionCallTestConfig,
	"assembly_action_gateway_script": testutils.AssemblyActionGatewayScriptTestConfig,
	"assembly_action_graphql_cost_analysis": testutils.AssemblyActionGraphQLCostAnalysisTestConfig,
	"assembly_action_graphql_execute": testutils.AssemblyActionGraphQLExecuteTestConfig,
	"assembly_action_graphql_introspect": testutils.AssemblyActionGraphQLIntrospectTestConfig,
	"assembly_action_html_page": testutils.AssemblyActionHtmlPageTestConfig,
	"assembly_action_invoke": testutils.AssemblyActionInvokeTestConfig,
	"assembly_action_jwt_generate": testutils.AssemblyActionJWTGenerateTestConfig,
	"assembly_action_jwt_validate": testutils.AssemblyActionJWTValidateTestConfig,
	"assembly_action_json2xml": testutils.AssemblyActionJson2XmlTestConfig,
	"assembly_action_log": testutils.AssemblyActionLogTestConfig,
	"assembly_action_map": testutils.AssemblyActionMapTestConfig,
	"assembly_action_oauth": testutils.AssemblyActionOAuthTestConfig,
	"assembly_action_parse": testutils.AssemblyActionParseTestConfig,
	"assembly_action_rate_limit": testutils.AssemblyActionRateLimitTestConfig,
	"assembly_action_rate_limit_info": testutils.AssemblyActionRateLimitInfoTestConfig,
	"assembly_action_redact": testutils.AssemblyActionRedactTestConfig,
	"assembly_action_set_var": testutils.AssemblyActionSetVarTestConfig,
	"assembly_action_throw": testutils.AssemblyActionThrowTestConfig,
	"assembly_action_user_security": testutils.AssemblyActionUserSecurityTestConfig,
	"assembly_action_validate": testutils.AssemblyActionValidateTestConfig,
	"assembly_action_wsdl": testutils.AssemblyActionWSDLTestConfig,
	"assembly_action_web_socket_upgrade": testutils.AssemblyActionWebSocketUpgradeTestConfig,
	"assembly_action_xslt": testutils.AssemblyActionXSLTTestConfig,
	"assembly_action_xml2json": testutils.AssemblyActionXml2JsonTestConfig,
	"assembly_function": testutils.AssemblyFunctionTestConfig,
	"assembly_logic_operation_switch": testutils.AssemblyLogicOperationSwitchTestConfig,
	"assembly_logic_switch": testutils.AssemblyLogicSwitchTestConfig,
	"audit_log": testutils.AuditLogTestConfig,
	"b2b_cpa": testutils.B2BCPATestConfig,
	"b2b_cpa_collaboration": testutils.B2BCPACollaborationTestConfig,
	"b2b_cpa_receiver_setting": testutils.B2BCPAReceiverSettingTestConfig,
	"b2b_cpa_sender_setting": testutils.B2BCPASenderSettingTestConfig,
	"b2b_gateway": testutils.B2BGatewayTestConfig,
	"b2b_persistence": testutils.B2BPersistenceTestConfig,
	"b2b_profile": testutils.B2BProfileTestConfig,
	"b2b_profile_group": testutils.B2BProfileGroupTestConfig,
	"b2b_xpath_routing_policy": testutils.B2BXPathRoutingPolicyTestConfig,
	"cors_policy": testutils.CORSPolicyTestConfig,
	"cors_rule": testutils.CORSRuleTestConfig,
	"crl_fetch": testutils.CRLFetchTestConfig,
	"cert_monitor": testutils.CertMonitorTestConfig,
	"compile_options_policy": testutils.CompileOptionsPolicyTestConfig,
	"compile_settings": testutils.CompileSettingsTestConfig,
	"config_deployment_policy": testutils.ConfigDeploymentPolicyTestConfig,
	"config_sequence": testutils.ConfigSequenceTestConfig,
	"conformance_policy": testutils.ConformancePolicyTestConfig,
	"control_list": testutils.ControlListTestConfig,
	"cookie_attribute_policy": testutils.CookieAttributePolicyTestConfig,
	"count_monitor": testutils.CountMonitorTestConfig,
	"crypto_certificate": testutils.CryptoCertificateTestConfig,
	"crypto_fw_cred": testutils.CryptoFWCredTestConfig,
	"crypto_ident_cred": testutils.CryptoIdentCredTestConfig,
	"crypto_kerberos_kdc": testutils.CryptoKerberosKDCTestConfig,
	"crypto_kerberos_keytab": testutils.CryptoKerberosKeytabTestConfig,
	"crypto_key": testutils.CryptoKeyTestConfig,
	"crypto_sskey": testutils.CryptoSSKeyTestConfig,
	"crypto_val_cred": testutils.CryptoValCredTestConfig,
	"dns_name_service": testutils.DNSNameServiceTestConfig,
	"deployment_policy_parameters_binding": testutils.DeploymentPolicyParametersBindingTestConfig,
	"distributed_variable": testutils.DistributedVariableTestConfig,
	"document_crypto_map": testutils.DocumentCryptoMapTestConfig,
	"domain": testutils.DomainTestConfig,
	"domain_availability": testutils.DomainAvailabilityTestConfig,
	"domain_settings": testutils.DomainSettingsTestConfig,
	"duration_monitor": testutils.DurationMonitorTestConfig,
	"ebms2_source_protocol_handler": testutils.EBMS2SourceProtocolHandlerTestConfig,
	"ebms3_source_protocol_handler": testutils.EBMS3SourceProtocolHandlerTestConfig,
	"error_report_settings": testutils.ErrorReportSettingsTestConfig,
	"ftp_file_poller_source_protocol_handler": testutils.FTPFilePollerSourceProtocolHandlerTestConfig,
	"ftp_quote_commands": testutils.FTPQuoteCommandsTestConfig,
	"ftp_server_source_protocol_handler": testutils.FTPServerSourceProtocolHandlerTestConfig,
	"file_system_usage_monitor": testutils.FileSystemUsageMonitorTestConfig,
	"filter_action": testutils.FilterActionTestConfig,
	"forms_login_policy": testutils.FormsLoginPolicyTestConfig,
	"gws_remote_debug": testutils.GWSRemoteDebugTestConfig,
	"gw_script_settings": testutils.GWScriptSettingsTestConfig,
	"gateway_peering": testutils.GatewayPeeringTestConfig,
	"gateway_peering_group": testutils.GatewayPeeringGroupTestConfig,
	"gateway_peering_manager": testutils.GatewayPeeringManagerTestConfig,
	"git_ops": testutils.GitOpsTestConfig,
	"git_ops_template": testutils.GitOpsTemplateTestConfig,
	"git_ops_variables": testutils.GitOpsVariablesTestConfig,
	"graphql_schema_options": testutils.GraphQLSchemaOptionsTestConfig,
	"http_input_conversion_map": testutils.HTTPInputConversionMapTestConfig,
	"https_source_protocol_handler": testutils.HTTPSSourceProtocolHandlerTestConfig,
	"http_source_protocol_handler": testutils.HTTPSourceProtocolHandlerTestConfig,
	"http_user_agent": testutils.HTTPUserAgentTestConfig,
	"host_alias": testutils.HostAliasTestConfig,
	"import_package": testutils.ImportPackageTestConfig,
	"include_config": testutils.IncludeConfigTestConfig,
	"interop_service": testutils.InteropServiceTestConfig,
	"jose_recipient_identifier": testutils.JOSERecipientIdentifierTestConfig,
	"jose_signature_identifier": testutils.JOSESignatureIdentifierTestConfig,
	"json_settings": testutils.JSONSettingsTestConfig,
	"jwe_header": testutils.JWEHeaderTestConfig,
	"jwe_recipient": testutils.JWERecipientTestConfig,
	"jws_signature": testutils.JWSSignatureTestConfig,
	"kafka_cluster": testutils.KafkaClusterTestConfig,
	"kafka_source_protocol_handler": testutils.KafkaSourceProtocolHandlerTestConfig,
	"ldap_connection_pool": testutils.LDAPConnectionPoolTestConfig,
	"ldap_search_parameters": testutils.LDAPSearchParametersTestConfig,
	"load_balancer_group": testutils.LoadBalancerGroupTestConfig,
	"log_label": testutils.LogLabelTestConfig,
	"log_target": testutils.LogTargetTestConfig,
	"luna": testutils.LunaTestConfig,
	"luna_ha_group": testutils.LunaHAGroupTestConfig,
	"luna_ha_settings": testutils.LunaHASettingsTestConfig,
	"luna_partition": testutils.LunaPartitionTestConfig,
	"mcf_custom_rule": testutils.MCFCustomRuleTestConfig,
	"mcf_http_header": testutils.MCFHttpHeaderTestConfig,
	"mcf_http_method": testutils.MCFHttpMethodTestConfig,
	"mcf_http_url": testutils.MCFHttpURLTestConfig,
	"mcf_xpath": testutils.MCFXPathTestConfig,
	"mpgw_error_action": testutils.MPGWErrorActionTestConfig,
	"mpgw_error_handling_policy": testutils.MPGWErrorHandlingPolicyTestConfig,
	"mq_manager": testutils.MQManagerTestConfig,
	"mq_manager_group": testutils.MQManagerGroupTestConfig,
	"mqv9_plus_mft_source_protocol_handler": testutils.MQv9PlusMFTSourceProtocolHandlerTestConfig,
	"mqv9_plus_source_protocol_handler": testutils.MQv9PlusSourceProtocolHandlerTestConfig,
	"mtom_policy": testutils.MTOMPolicyTestConfig,
	"matching": testutils.MatchingTestConfig,
	"message_content_filters": testutils.MessageContentFiltersTestConfig,
	"message_matching": testutils.MessageMatchingTestConfig,
	"message_type": testutils.MessageTypeTestConfig,
	"mgmt_interface": testutils.MgmtInterfaceTestConfig,
	"multi_protocol_gateway": testutils.MultiProtocolGatewayTestConfig,
	"nfs_client_settings": testutils.NFSClientSettingsTestConfig,
	"nfs_dynamic_mounts": testutils.NFSDynamicMountsTestConfig,
	"nfs_file_poller_source_protocol_handler": testutils.NFSFilePollerSourceProtocolHandlerTestConfig,
	"nfs_static_mount": testutils.NFSStaticMountTestConfig,
	"name_value_profile": testutils.NameValueProfileTestConfig,
	"oauth_provider_settings": testutils.OAuthProviderSettingsTestConfig,
	"oauth_supported_client": testutils.OAuthSupportedClientTestConfig,
	"oauth_supported_client_group": testutils.OAuthSupportedClientGroupTestConfig,
	"odr": testutils.ODRTestConfig,
	"odr_connector_group": testutils.ODRConnectorGroupTestConfig,
	"open_telemetry": testutils.OpenTelemetryTestConfig,
	"open_telemetry_exporter": testutils.OpenTelemetryExporterTestConfig,
	"open_telemetry_sampler": testutils.OpenTelemetrySamplerTestConfig,
	"operation_rate_limit": testutils.OperationRateLimitTestConfig,
	"pop_poller_source_protocol_handler": testutils.POPPollerSourceProtocolHandlerTestConfig,
	"parse_settings": testutils.ParseSettingsTestConfig,
	"password_alias": testutils.PasswordAliasTestConfig,
	"peer_group": testutils.PeerGroupTestConfig,
	"policy_attachments": testutils.PolicyAttachmentsTestConfig,
	"policy_parameters": testutils.PolicyParametersTestConfig,
	"probe": testutils.ProbeTestConfig,
	"processing_metadata": testutils.ProcessingMetadataTestConfig,
	"quota_enforcement_server": testutils.QuotaEnforcementServerTestConfig,
	"radius_settings": testutils.RADIUSSettingsTestConfig,
	"rbm_settings": testutils.RBMSettingsTestConfig,
	"raid_volume": testutils.RaidVolumeTestConfig,
	"rate_limit_configuration": testutils.RateLimitConfigurationTestConfig,
	"rate_limit_definition": testutils.RateLimitDefinitionTestConfig,
	"rate_limit_definition_group": testutils.RateLimitDefinitionGroupTestConfig,
	"rest_mgmt_interface": testutils.RestMgmtInterfaceTestConfig,
	"saml_attributes": testutils.SAMLAttributesTestConfig,
	"sftp_file_poller_source_protocol_handler": testutils.SFTPFilePollerSourceProtocolHandlerTestConfig,
	"slm_action": testutils.SLMActionTestConfig,
	"slm_cred_class": testutils.SLMCredClassTestConfig,
	"slm_policy": testutils.SLMPolicyTestConfig,
	"slm_rsrc_class": testutils.SLMRsrcClassTestConfig,
	"slm_schedule": testutils.SLMScheduleTestConfig,
	"smtp_server_connection": testutils.SMTPServerConnectionTestConfig,
	"snmp_settings": testutils.SNMPSettingsTestConfig,
	"soap_header_disposition": testutils.SOAPHeaderDispositionTestConfig,
	"sql_data_source": testutils.SQLDataSourceTestConfig,
	"ssh_client_profile": testutils.SSHClientProfileTestConfig,
	"ssh_domain_client_profile": testutils.SSHDomainClientProfileTestConfig,
	"ssh_server_profile": testutils.SSHServerProfileTestConfig,
	"ssh_server_source_protocol_handler": testutils.SSHServerSourceProtocolHandlerTestConfig,
	"ssh_service": testutils.SSHServiceTestConfig,
	"ssl_client_profile": testutils.SSLClientProfileTestConfig,
	"ssl_proxy_service": testutils.SSLProxyServiceTestConfig,
	"ssl_sni_mapping": testutils.SSLSNIMappingTestConfig,
	"ssl_sni_server_profile": testutils.SSLSNIServerProfileTestConfig,
	"ssl_server_profile": testutils.SSLServerProfileTestConfig,
	"schema_exception_map": testutils.SchemaExceptionMapTestConfig,
	"secure_backup_mode": testutils.SecureBackupModeTestConfig,
	"social_login_policy": testutils.SocialLoginPolicyTestConfig,
	"stateless_tcp_source_protocol_handler": testutils.StatelessTCPSourceProtocolHandlerTestConfig,
	"statistics": testutils.StatisticsTestConfig,
	"style_policy": testutils.StylePolicyTestConfig,
	"style_policy_action": testutils.StylePolicyActionTestConfig,
	"style_policy_rule": testutils.StylePolicyRuleTestConfig,
	"system_settings": testutils.SystemSettingsTestConfig,
	"tam": testutils.TAMTestConfig,
	"tcp_proxy_service": testutils.TCPProxyServiceTestConfig,
	"throttler": testutils.ThrottlerTestConfig,
	"time_settings": testutils.TimeSettingsTestConfig,
	"url_map": testutils.URLMapTestConfig,
	"url_refresh_policy": testutils.URLRefreshPolicyTestConfig,
	"url_rewrite_policy": testutils.URLRewritePolicyTestConfig,
	"user": testutils.UserTestConfig,
	"user_group": testutils.UserGroupTestConfig,
	"visibility_list": testutils.VisibilityListTestConfig,
	"wcc_service": testutils.WCCServiceTestConfig,
	"ws_endpoint_rewrite_policy": testutils.WSEndpointRewritePolicyTestConfig,
	"ws_gateway": testutils.WSGatewayTestConfig,
	"wsrr_saved_search_subscription": testutils.WSRRSavedSearchSubscriptionTestConfig,
	"wsrr_server": testutils.WSRRServerTestConfig,
	"wsrr_subscription": testutils.WSRRSubscriptionTestConfig,
	"ws_style_policy": testutils.WSStylePolicyTestConfig,
	"ws_style_policy_rule": testutils.WSStylePolicyRuleTestConfig,
	"wxs_grid": testutils.WXSGridTestConfig,
	"web_app_error_handling_policy": testutils.WebAppErrorHandlingPolicyTestConfig,
	"web_app_fw": testutils.WebAppFWTestConfig,
	"web_app_request": testutils.WebAppRequestTestConfig,
	"web_app_response": testutils.WebAppResponseTestConfig,
	"web_app_session_policy": testutils.WebAppSessionPolicyTestConfig,
	"web_b2b_viewer": testutils.WebB2BViewerTestConfig,
	"web_gui": testutils.WebGUITestConfig,
	"web_service_monitor": testutils.WebServiceMonitorTestConfig,
	"web_services_agent": testutils.WebServicesAgentTestConfig,
	"web_sphere_jms_server": testutils.WebSphereJMSServerTestConfig,
	"web_sphere_jms_source_protocol_handler": testutils.WebSphereJMSSourceProtocolHandlerTestConfig,
	"web_token_service": testutils.WebTokenServiceTestConfig,
	"xacml_pdp": testutils.XACMLPDPTestConfig,
	"xml_firewall_service": testutils.XMLFirewallServiceTestConfig,
	"xml_manager": testutils.XMLManagerTestConfig,
	"xpath_routing_map": testutils.XPathRoutingMapTestConfig,
	"xsl_coproc_service": testutils.XSLCoprocServiceTestConfig,
	"xsl_proxy_service": testutils.XSLProxyServiceTestConfig,
	"xtc_protocol_handler": testutils.XTCProtocolHandlerTestConfig,
	"zos_nss_client": testutils.ZosNSSClientTestConfig,
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
