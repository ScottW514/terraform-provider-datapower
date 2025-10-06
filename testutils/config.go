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

package testutils

type ModelTestConfig struct {
	Name         string
	Model        string
	Resource     string
	Data         string
    TestBed      string
    ModelTestBed string
	ModelOnly    bool
	TestPre      string
}

func (c ModelTestConfig) GetModelListConfig() string {
	if c.Model == "\n" {
		return "null"
	} else {
		return "[" + c.Model + "]"
	}
}

func (c ModelTestConfig) GetModelConfig() string {
	if c.Model == "{}" {
		return "null"
	} else {
		return c.Model
	}
}

func (c ModelTestConfig) GetModelTestBedListConfig() string {
	if c.Model == "\n" {
		return "null"
	} else {
		return "[" + c.ModelTestBed + "]"
	}
}

func (c ModelTestConfig) GetModelTestBedConfig() string {
	if c.Model == "{}" {
		return "null"
	} else {
		return c.ModelTestBed
	}
}

func (c ModelTestConfig) GetDataConfig() string {
	return c.Data
}

func (c ModelTestConfig) GetResourceConfig() string {
	return c.getPrequisites() + c.Resource
}

func (c *ModelTestConfig) getPrequisites() string {
	return c.TestPre
}

var FileTestConfig = ModelTestConfig{
	Name: "File",
	Resource: `
resource "datapower_file" "acc_test" {
  app_domain = "acceptance_test"
  remote_path = "cert:///test_file.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
}
`,
	ModelOnly:    false,
}
var AAAJWTGeneratorTestConfig = ModelTestConfig{
    Name:         "AAAJWTGenerator",
    Resource: `
resource "datapower_aaa_jwt_generator" "test" {
  id = "ResTestAAAJWTGenerator"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_aaa_jwt_generator" "test" {
  id         = "AccTest_AAAJWTGenerator"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_aaa_jwt_generator" "acc_test" {
  id = "AccTest_AAAJWTGenerator"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AAAJWTValidatorTestConfig = ModelTestConfig{
    Name:         "AAAJWTValidator",
    Resource: `
resource "datapower_aaa_jwt_validator" "test" {
  id = "ResTestAAAJWTValidator"
  app_domain = "acceptance_test"
  username_claim = "sub"
}`,
    Data: `
data "datapower_aaa_jwt_validator" "test" {
  id         = "AccTest_AAAJWTValidator"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_aaa_jwt_validator" "acc_test" {
  id = "AccTest_AAAJWTValidator"
  app_domain = datapower_domain.acc_test.app_domain
  username_claim = "sub"
}`,
    ModelOnly:    false,
}
var AAAPolicyTestConfig = ModelTestConfig{
    Name:         "AAAPolicy",
    Resource: `
resource "datapower_aaa_policy" "test" {
  id = "ResTestAAAPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_aaa_policy" "test" {
  id         = "AccTest_AAAPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_aaa_policy" "acc_test" {
  id = "AccTest_AAAPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AMQPBrokerTestConfig = ModelTestConfig{
    Name:         "AMQPBroker",
    Resource: `
resource "datapower_amqp_broker" "test" {
  id = "ResTestAMQPBroker"
  app_domain = "acceptance_test"
  host_name = "host.name"
  port = 5672
  xml_manager = "default"
  authorization = "none"
}`,
    Data: `
data "datapower_amqp_broker" "test" {
  id         = "AccTest_AMQPBroker"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_amqp_broker" "acc_test" {
  id = "AccTest_AMQPBroker"
  app_domain = datapower_domain.acc_test.app_domain
  host_name = "host.name"
  port = 5672
  xml_manager = "default"
  authorization = "none"
}`,
    ModelOnly:    false,
}
var AMQPSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "AMQPSourceProtocolHandler",
    Resource: `
resource "datapower_amqp_source_protocol_handler" "test" {
  id = "ResTestAMQPSourceProtocolHandler"
  app_domain = "acceptance_test"
  broker = "AccTest_AMQPBroker"
  from = "amqpfrom"
}`,
    Data: `
data "datapower_amqp_source_protocol_handler" "test" {
  id         = "AccTest_AMQPSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_amqp_source_protocol_handler" "acc_test" {
  id = "AccTest_AMQPSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  broker = datapower_amqp_broker.acc_test.id
  from = "amqpfrom"
}`,
    ModelOnly:    false,
}
var APIApplicationTypeTestConfig = ModelTestConfig{
    Name:         "APIApplicationType",
    Resource: `
resource "datapower_api_application_type" "test" {
  id = "ResTestAPIApplicationType"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_application_type" "test" {
  id         = "AccTest_APIApplicationType"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_application_type" "acc_test" {
  id = "AccTest_APIApplicationType"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIAuthURLRegistryTestConfig = ModelTestConfig{
    Name:         "APIAuthURLRegistry",
    Resource: `
resource "datapower_api_auth_url_registry" "test" {
  id = "ResTestAPIAuthURLRegistry"
  app_domain = "acceptance_test"
  auth_url = "http://localhost"
}`,
    Data: `
data "datapower_api_auth_url_registry" "test" {
  id         = "AccTest_APIAuthURLRegistry"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_auth_url_registry" "acc_test" {
  id = "AccTest_APIAuthURLRegistry"
  app_domain = datapower_domain.acc_test.app_domain
  auth_url = "http://localhost"
}`,
    ModelOnly:    false,
}
var APICORSTestConfig = ModelTestConfig{
    Name:         "APICORS",
    Resource: `
resource "datapower_api_cors" "test" {
  id = "ResTestAPICORS"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_cors" "test" {
  id         = "AccTest_APICORS"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_cors" "acc_test" {
  id = "AccTest_APICORS"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIClientIdentificationTestConfig = ModelTestConfig{
    Name:         "APIClientIdentification",
    Resource: `
resource "datapower_api_client_identification" "test" {
  id = "ResTestAPIClientIdentification"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_client_identification" "test" {
  id         = "AccTest_APIClientIdentification"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_client_identification" "acc_test" {
  id = "AccTest_APIClientIdentification"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APICollectionTestConfig = ModelTestConfig{
    Name:         "APICollection",
    Resource: `
resource "datapower_api_collection" "test" {
  id = "ResTestAPICollection"
  app_domain = "acceptance_test"
  org_id = "orgid"
  org_name = "orgname"
  routing_prefix = ` + DmRoutingPrefixTestConfig.GetModelListConfig() + `
  api_processing_rule = "default-api-rule"
  api_error_rule = "default-api-error-rule"
  plan = ["AccTest_APIPlan"]
}`,
    Data: `
data "datapower_api_collection" "test" {
  id         = "AccTest_APICollection"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_collection" "acc_test" {
  id = "AccTest_APICollection"
  app_domain = datapower_domain.acc_test.app_domain
  org_id = "orgid"
  org_name = "orgname"
  routing_prefix = ` + DmRoutingPrefixTestConfig.GetModelTestBedListConfig() + `
  api_processing_rule = "default-api-rule"
  api_error_rule = "default-api-error-rule"
  plan = [datapower_api_plan.acc_test.id]
}`,
    ModelOnly:    false,
}
var APIConnectGatewayServiceTestConfig = ModelTestConfig{
    Name:         "APIConnectGatewayService",
    Resource: `
resource "datapower_api_connect_gateway_service" "test" {
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 3000
  gateway_peering = "default-gateway-peering"
}`,
    Data: `
data "datapower_api_connect_gateway_service" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_connect_gateway_service" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 3000
  gateway_peering = "default-gateway-peering"
}`,
    ModelOnly:    false,
}
var APIDefinitionTestConfig = ModelTestConfig{
    Name:         "APIDefinition",
    Resource: `
resource "datapower_api_definition" "test" {
  id = "ResTestAPIDefinition"
  app_domain = "acceptance_test"
  base_path = "/"
  path = ["AccTest_APIPath"]
}`,
    Data: `
data "datapower_api_definition" "test" {
  id         = "AccTest_APIDefinition"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_definition" "acc_test" {
  id = "AccTest_APIDefinition"
  app_domain = datapower_domain.acc_test.app_domain
  base_path = "/"
  path = [datapower_api_path.acc_test.id]
}`,
    ModelOnly:    false,
}
var APIExecuteTestConfig = ModelTestConfig{
    Name:         "APIExecute",
    Resource: `
resource "datapower_api_execute" "test" {
  id = "ResTestAPIExecute"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_execute" "test" {
  id         = "AccTest_APIExecute"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_execute" "acc_test" {
  id = "AccTest_APIExecute"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIFinalTestConfig = ModelTestConfig{
    Name:         "APIFinal",
    Resource: `
resource "datapower_api_final" "test" {
  id = "ResTestAPIFinal"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_final" "test" {
  id         = "AccTest_APIFinal"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_final" "acc_test" {
  id = "AccTest_APIFinal"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIGatewayTestConfig = ModelTestConfig{
    Name:         "APIGateway",
    Resource: `
resource "datapower_api_gateway" "test" {
  id = "ResTestAPIGateway"
  app_domain = "acceptance_test"
  front_protocol = ["AccTest_HTTPSourceProtocolHandler"]
  front_timeout = 120
  front_persistent_timeout = 180
}`,
    Data: `
data "datapower_api_gateway" "test" {
  id         = "AccTest_APIGateway"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_gateway" "acc_test" {
  id = "AccTest_APIGateway"
  app_domain = datapower_domain.acc_test.app_domain
  front_protocol = [datapower_http_source_protocol_handler.acc_test.id]
  front_timeout = 120
  front_persistent_timeout = 180
}`,
    ModelOnly:    false,
}
var APILDAPRegistryTestConfig = ModelTestConfig{
    Name:         "APILDAPRegistry",
    Resource: `
resource "datapower_api_ldap_registry" "test" {
  id = "ResTestAPILDAPRegistry"
  app_domain = "acceptance_test"
  ldap_host = "localhost"
  ldap_port = 636
  ldap_search_parameters = "AccTest_LDAPSearchParameters"
}`,
    Data: `
data "datapower_api_ldap_registry" "test" {
  id         = "AccTest_APILDAPRegistry"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_ldap_registry" "acc_test" {
  id = "AccTest_APILDAPRegistry"
  app_domain = datapower_domain.acc_test.app_domain
  ldap_host = "localhost"
  ldap_port = 636
  ldap_search_parameters = datapower_ldap_search_parameters.acc_test.id
}`,
    ModelOnly:    false,
}
var APIOperationTestConfig = ModelTestConfig{
    Name:         "APIOperation",
    Resource: `
resource "datapower_api_operation" "test" {
  id = "ResTestAPIOperation"
  app_domain = "acceptance_test"
  method = "GET"
}`,
    Data: `
data "datapower_api_operation" "test" {
  id         = "AccTest_APIOperation"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_operation" "acc_test" {
  id = "AccTest_APIOperation"
  app_domain = datapower_domain.acc_test.app_domain
  method = "GET"
}`,
    ModelOnly:    false,
}
var APIPathTestConfig = ModelTestConfig{
    Name:         "APIPath",
    Resource: `
resource "datapower_api_path" "test" {
  id = "ResTestAPIPath"
  app_domain = "acceptance_test"
  path = "/"
}`,
    Data: `
data "datapower_api_path" "test" {
  id         = "AccTest_APIPath"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_path" "acc_test" {
  id = "AccTest_APIPath"
  app_domain = datapower_domain.acc_test.app_domain
  path = "/"
}`,
    ModelOnly:    false,
}
var APIPlanTestConfig = ModelTestConfig{
    Name:         "APIPlan",
    Resource: `
resource "datapower_api_plan" "test" {
  id = "ResTestAPIPlan"
  app_domain = "acceptance_test"
  api = ["AccTest_APIDefinition"]
  rate_limit_scope = "per-application"
}`,
    Data: `
data "datapower_api_plan" "test" {
  id         = "AccTest_APIPlan"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_plan" "acc_test" {
  id = "AccTest_APIPlan"
  app_domain = datapower_domain.acc_test.app_domain
  api = [datapower_api_definition.acc_test.id]
  rate_limit_scope = "per-application"
}`,
    ModelOnly:    false,
}
var APIResultTestConfig = ModelTestConfig{
    Name:         "APIResult",
    Resource: `
resource "datapower_api_result" "test" {
  id = "ResTestAPIResult"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_result" "test" {
  id         = "AccTest_APIResult"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_result" "acc_test" {
  id = "AccTest_APIResult"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIRoutingTestConfig = ModelTestConfig{
    Name:         "APIRouting",
    Resource: `
resource "datapower_api_routing" "test" {
  id = "ResTestAPIRouting"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_routing" "test" {
  id         = "AccTest_APIRouting"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_routing" "acc_test" {
  id = "AccTest_APIRouting"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APIRuleTestConfig = ModelTestConfig{
    Name:         "APIRule",
    Resource: `
resource "datapower_api_rule" "test" {
  id = "ResTestAPIRule"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_rule" "test" {
  id         = "AccTest_APIRule"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_rule" "acc_test" {
  id = "AccTest_APIRule"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APISchemaTestConfig = ModelTestConfig{
    Name:         "APISchema",
    Resource: `
resource "datapower_api_schema" "test" {
  id = "ResTestAPISchema"
  app_domain = "acceptance_test"
  json_schema = "http://localhost/json"
}`,
    Data: `
data "datapower_api_schema" "test" {
  id         = "AccTest_APISchema"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_schema" "acc_test" {
  id = "AccTest_APISchema"
  app_domain = datapower_domain.acc_test.app_domain
  json_schema = "http://localhost/json"
}`,
    ModelOnly:    false,
}
var APISecurityTestConfig = ModelTestConfig{
    Name:         "APISecurity",
    Resource: `
resource "datapower_api_security" "test" {
  id = "ResTestAPISecurity"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_security" "test" {
  id         = "AccTest_APISecurity"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security" "acc_test" {
  id = "AccTest_APISecurity"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APISecurityAPIKeyTestConfig = ModelTestConfig{
    Name:         "APISecurityAPIKey",
    Resource: `
resource "datapower_api_security_api_key" "test" {
  id = "ResTestAPISecurityAPIKey"
  app_domain = "acceptance_test"
  where = "header"
  type = "id"
  key_name = "keyname"
}`,
    Data: `
data "datapower_api_security_api_key" "test" {
  id         = "AccTest_APISecurityAPIKey"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_api_key" "acc_test" {
  id = "AccTest_APISecurityAPIKey"
  app_domain = datapower_domain.acc_test.app_domain
  where = "header"
  type = "id"
  key_name = "keyname"
}`,
    ModelOnly:    false,
}
var APISecurityBasicAuthTestConfig = ModelTestConfig{
    Name:         "APISecurityBasicAuth",
    Resource: `
resource "datapower_api_security_basic_auth" "test" {
  id = "ResTestAPISecurityBasicAuth"
  app_domain = "acceptance_test"
  user_registry = "AccTest_APIAuthURLRegistry"
}`,
    Data: `
data "datapower_api_security_basic_auth" "test" {
  id         = "AccTest_APISecurityBasicAuth"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_basic_auth" "acc_test" {
  id = "AccTest_APISecurityBasicAuth"
  app_domain = datapower_domain.acc_test.app_domain
  user_registry = datapower_api_auth_url_registry.acc_test.id
}`,
    ModelOnly:    false,
}
var APISecurityHTTPSchemeTestConfig = ModelTestConfig{
    Name:         "APISecurityHTTPScheme",
    Resource: `
resource "datapower_api_security_http_scheme" "test" {
  id = "ResTestAPISecurityHTTPScheme"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_security_http_scheme" "test" {
  id         = "AccTest_APISecurityHTTPScheme"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_http_scheme" "acc_test" {
  id = "AccTest_APISecurityHTTPScheme"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var APISecurityOAuthTestConfig = ModelTestConfig{
    Name:         "APISecurityOAuth",
    Resource: `
resource "datapower_api_security_oauth" "test" {
  id = "ResTestAPISecurityOAuth"
  app_domain = "acceptance_test"
  oauth_provider = "AccTest_OAuthProviderSettings"
  oauth_flow = "implicit"
}`,
    Data: `
data "datapower_api_security_oauth" "test" {
  id         = "AccTest_APISecurityOAuth"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_oauth" "acc_test" {
  id = "AccTest_APISecurityOAuth"
  app_domain = datapower_domain.acc_test.app_domain
  oauth_provider = datapower_oauth_provider_settings.acc_test.id
  oauth_flow = "implicit"
}`,
    ModelOnly:    false,
}
var APISecurityOAuthReqTestConfig = ModelTestConfig{
    Name:         "APISecurityOAuthReq",
    Resource: `
resource "datapower_api_security_oauth_req" "test" {
  id = "ResTestAPISecurityOAuthReq"
  app_domain = "acceptance_test"
  api_security_oauth_def = "AccTest_APISecurityOAuth"
}`,
    Data: `
data "datapower_api_security_oauth_req" "test" {
  id         = "AccTest_APISecurityOAuthReq"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_oauth_req" "acc_test" {
  id = "AccTest_APISecurityOAuthReq"
  app_domain = datapower_domain.acc_test.app_domain
  api_security_oauth_def = datapower_api_security_oauth.acc_test.id
}`,
    ModelOnly:    false,
}
var APISecurityRequirementTestConfig = ModelTestConfig{
    Name:         "APISecurityRequirement",
    Resource: `
resource "datapower_api_security_requirement" "test" {
  id = "ResTestAPISecurityRequirement"
  app_domain = "acceptance_test"
  security = ["AccTest_APISecurityAPIKey"]
}`,
    Data: `
data "datapower_api_security_requirement" "test" {
  id         = "AccTest_APISecurityRequirement"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_requirement" "acc_test" {
  id = "AccTest_APISecurityRequirement"
  app_domain = datapower_domain.acc_test.app_domain
  security = [datapower_api_security_api_key.acc_test.id]
}`,
    ModelOnly:    false,
}
var APISecurityTokenManagerTestConfig = ModelTestConfig{
    Name:         "APISecurityTokenManager",
    Resource: `
resource "datapower_api_security_token_manager" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_api_security_token_manager" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_api_security_token_manager" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AS1PollerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "AS1PollerSourceProtocolHandler",
    Resource: `
resource "datapower_as1_poller_source_protocol_handler" "test" {
  id = "ResTestAS1PollerSourceProtocolHandler"
  app_domain = "acceptance_test"
  mail_server = "smtp.host"
  port = 25
  conn_security = "none"
  account = "account"
  password_alias = "AccTest_PasswordAlias"
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
    Data: `
data "datapower_as1_poller_source_protocol_handler" "test" {
  id         = "AccTest_AS1PollerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_as1_poller_source_protocol_handler" "acc_test" {
  id = "AccTest_AS1PollerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  mail_server = "smtp.host"
  port = 25
  conn_security = "none"
  account = "account"
  password_alias = datapower_password_alias.acc_test.id
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
    ModelOnly:    false,
}
var AS2ProxySourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "AS2ProxySourceProtocolHandler",
    Resource: `
resource "datapower_as2_proxy_source_protocol_handler" "test" {
  id = "ResTestAS2ProxySourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 80
  remote_address = "10.10.10.10"
  remote_port = 8888
  remote_connection_timeout = 60
  xml_manager = "default"
}`,
    Data: `
data "datapower_as2_proxy_source_protocol_handler" "test" {
  id         = "AccTest_AS2ProxySourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_AS2ProxySourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 80
  remote_address = "10.10.10.10"
  remote_port = 8888
  remote_connection_timeout = 60
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var AS2SourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "AS2SourceProtocolHandler",
    Resource: `
resource "datapower_as2_source_protocol_handler" "test" {
  id = "ResTestAS2SourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 80
}`,
    Data: `
data "datapower_as2_source_protocol_handler" "test" {
  id         = "AccTest_AS2SourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_as2_source_protocol_handler" "acc_test" {
  id = "AccTest_AS2SourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 80
}`,
    ModelOnly:    false,
}
var AS3SourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "AS3SourceProtocolHandler",
    Resource: `
resource "datapower_as3_source_protocol_handler" "test" {
  id = "ResTestAS3SourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 21
}`,
    Data: `
data "datapower_as3_source_protocol_handler" "test" {
  id         = "AccTest_AS3SourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_as3_source_protocol_handler" "acc_test" {
  id = "AccTest_AS3SourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 21
}`,
    ModelOnly:    false,
}
var AccessControlListTestConfig = ModelTestConfig{
    Name:         "AccessControlList",
    Resource: `
resource "datapower_access_control_list" "test" {
  id = "ResTestAccessControlList"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_access_control_list" "test" {
  id         = "AccTest_AccessControlList"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_access_control_list" "acc_test" {
  id = "AccTest_AccessControlList"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AccessProfileTestConfig = ModelTestConfig{
    Name:         "AccessProfile",
    Resource: `
resource "datapower_access_profile" "test" {
  id = "ResTestAccessProfile"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_access_profile" "test" {
  id         = "AccTest_AccessProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_access_profile" "acc_test" {
  id = "AccTest_AccessProfile"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AnalyticsEndpointTestConfig = ModelTestConfig{
    Name:         "AnalyticsEndpoint",
    Resource: `
resource "datapower_analytics_endpoint" "test" {
  id = "ResTestAnalyticsEndpoint"
  app_domain = "acceptance_test"
  analytics_server_url = "https://localhost"
  ssl_client = "AccTest_SSLClientProfile"
  max_records = 1024
  max_records_memory_kb = 512
  max_delivery_memory_mb = 512
}`,
    Data: `
data "datapower_analytics_endpoint" "test" {
  id         = "AccTest_AnalyticsEndpoint"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_analytics_endpoint" "acc_test" {
  id = "AccTest_AnalyticsEndpoint"
  app_domain = datapower_domain.acc_test.app_domain
  analytics_server_url = "https://localhost"
  ssl_client = datapower_ssl_client_profile.acc_test.id
  max_records = 1024
  max_records_memory_kb = 512
  max_delivery_memory_mb = 512
}`,
    ModelOnly:    false,
}
var AppSecurityPolicyTestConfig = ModelTestConfig{
    Name:         "AppSecurityPolicy",
    Resource: `
resource "datapower_app_security_policy" "test" {
  id = "ResTestAppSecurityPolicy"
  app_domain = "acceptance_test"
  request_maps = ` + DmWebAppRequestPolicyMapTestConfig.GetModelListConfig() + `
  response_maps = ` + DmWebAppResponsePolicyMapTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_app_security_policy" "test" {
  id         = "AccTest_AppSecurityPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_app_security_policy" "acc_test" {
  id = "AccTest_AppSecurityPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  request_maps = ` + DmWebAppRequestPolicyMapTestConfig.GetModelTestBedListConfig() + `
  response_maps = ` + DmWebAppResponsePolicyMapTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AssemblyTestConfig = ModelTestConfig{
    Name:         "Assembly",
    Resource: `
resource "datapower_assembly" "test" {
  id = "ResTestAssembly"
  app_domain = "acceptance_test"
  rule = "default-empty-rule"
}`,
    Data: `
data "datapower_assembly" "test" {
  id         = "AccTest_Assembly"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly" "acc_test" {
  id = "AccTest_Assembly"
  app_domain = datapower_domain.acc_test.app_domain
  rule = "default-empty-rule"
}`,
    ModelOnly:    false,
}
var AssemblyActionClientSecurityTestConfig = ModelTestConfig{
    Name:         "AssemblyActionClientSecurity",
    Resource: `
resource "datapower_assembly_action_client_security" "test" {
  id = "ResTestAssemblyActionClientSecurity"
  app_domain = "acceptance_test"
  id_name = "idname"
  secret_name = "secretname"
  authenticate_client_method = "native"
}`,
    Data: `
data "datapower_assembly_action_client_security" "test" {
  id         = "AccTest_AssemblyActionClientSecurity"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_client_security" "acc_test" {
  id = "AccTest_AssemblyActionClientSecurity"
  app_domain = datapower_domain.acc_test.app_domain
  id_name = "idname"
  secret_name = "secretname"
  authenticate_client_method = "native"
}`,
    ModelOnly:    false,
}
var AssemblyActionExtractTestConfig = ModelTestConfig{
    Name:         "AssemblyActionExtract",
    Resource: `
resource "datapower_assembly_action_extract" "test" {
  id = "ResTestAssemblyActionExtract"
  app_domain = "acceptance_test"
  extract = ` + DmAssemblyActionExtractTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_assembly_action_extract" "test" {
  id         = "AccTest_AssemblyActionExtract"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_extract" "acc_test" {
  id = "AccTest_AssemblyActionExtract"
  app_domain = datapower_domain.acc_test.app_domain
  extract = ` + DmAssemblyActionExtractTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AssemblyActionFunctionCallTestConfig = ModelTestConfig{
    Name:         "AssemblyActionFunctionCall",
    Resource: `
resource "datapower_assembly_action_function_call" "test" {
  id = "ResTestAssemblyActionFunctionCall"
  app_domain = "acceptance_test"
  function_call = "default-func-global"
}`,
    Data: `
data "datapower_assembly_action_function_call" "test" {
  id         = "AccTest_AssemblyActionFunctionCall"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_function_call" "acc_test" {
  id = "AccTest_AssemblyActionFunctionCall"
  app_domain = datapower_domain.acc_test.app_domain
  function_call = "default-func-global"
}`,
    ModelOnly:    false,
}
var AssemblyActionGatewayScriptTestConfig = ModelTestConfig{
    Name:         "AssemblyActionGatewayScript",
    Resource: `
resource "datapower_assembly_action_gateway_script" "test" {
  id = "ResTestAssemblyActionGatewayScript"
  app_domain = "acceptance_test"
  source = "gsfile"
}`,
    Data: `
data "datapower_assembly_action_gateway_script" "test" {
  id         = "AccTest_AssemblyActionGatewayScript"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_gateway_script" "acc_test" {
  id = "AccTest_AssemblyActionGatewayScript"
  app_domain = datapower_domain.acc_test.app_domain
  source = "gsfile"
}`,
    ModelOnly:    false,
}
var AssemblyActionGraphQLCostAnalysisTestConfig = ModelTestConfig{
    Name:         "AssemblyActionGraphQLCostAnalysis",
    Resource: `
resource "datapower_assembly_action_graphql_cost_analysis" "test" {
  id = "ResTestAssemblyActionGraphQLCostAnalysis"
  app_domain = "acceptance_test"
  target = "targetquery"
}`,
    Data: `
data "datapower_assembly_action_graphql_cost_analysis" "test" {
  id         = "AccTest_AssemblyActionGraphQLCostAnalysis"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_graphql_cost_analysis" "acc_test" {
  id = "AccTest_AssemblyActionGraphQLCostAnalysis"
  app_domain = datapower_domain.acc_test.app_domain
  target = "targetquery"
}`,
    ModelOnly:    false,
}
var AssemblyActionGraphQLExecuteTestConfig = ModelTestConfig{
    Name:         "AssemblyActionGraphQLExecute",
    Resource: `
resource "datapower_assembly_action_graphql_execute" "test" {
  id = "ResTestAssemblyActionGraphQLExecute"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_graphql_execute" "test" {
  id         = "AccTest_AssemblyActionGraphQLExecute"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_graphql_execute" "acc_test" {
  id = "AccTest_AssemblyActionGraphQLExecute"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionGraphQLIntrospectTestConfig = ModelTestConfig{
    Name:         "AssemblyActionGraphQLIntrospect",
    Resource: `
resource "datapower_assembly_action_graphql_introspect" "test" {
  id = "ResTestAssemblyActionGraphQLIntrospect"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_graphql_introspect" "test" {
  id         = "AccTest_AssemblyActionGraphQLIntrospect"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_graphql_introspect" "acc_test" {
  id = "AccTest_AssemblyActionGraphQLIntrospect"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionHtmlPageTestConfig = ModelTestConfig{
    Name:         "AssemblyActionHtmlPage",
    Resource: `
resource "datapower_assembly_action_html_page" "test" {
  id = "ResTestAssemblyActionHtmlPage"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_html_page" "test" {
  id         = "AccTest_AssemblyActionHtmlPage"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_html_page" "acc_test" {
  id = "AccTest_AssemblyActionHtmlPage"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionInvokeTestConfig = ModelTestConfig{
    Name:         "AssemblyActionInvoke",
    Resource: `
resource "datapower_assembly_action_invoke" "test" {
  id = "ResTestAssemblyActionInvoke"
  app_domain = "acceptance_test"
  url = "https://localhost"
  method = "Keep"
  backend_type = "detect"
  cache_type = "Protocol"
}`,
    Data: `
data "datapower_assembly_action_invoke" "test" {
  id         = "AccTest_AssemblyActionInvoke"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_invoke" "acc_test" {
  id = "AccTest_AssemblyActionInvoke"
  app_domain = datapower_domain.acc_test.app_domain
  url = "https://localhost"
  method = "Keep"
  backend_type = "detect"
  cache_type = "Protocol"
}`,
    ModelOnly:    false,
}
var AssemblyActionJWTGenerateTestConfig = ModelTestConfig{
    Name:         "AssemblyActionJWTGenerate",
    Resource: `
resource "datapower_assembly_action_jwt_generate" "test" {
  id = "ResTestAssemblyActionJWTGenerate"
  app_domain = "acceptance_test"
  issuer_claim = "iss.claim"
  validity_period = 3600
}`,
    Data: `
data "datapower_assembly_action_jwt_generate" "test" {
  id         = "AccTest_AssemblyActionJWTGenerate"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_jwt_generate" "acc_test" {
  id = "AccTest_AssemblyActionJWTGenerate"
  app_domain = datapower_domain.acc_test.app_domain
  issuer_claim = "iss.claim"
  validity_period = 3600
}`,
    ModelOnly:    false,
}
var AssemblyActionJWTValidateTestConfig = ModelTestConfig{
    Name:         "AssemblyActionJWTValidate",
    Resource: `
resource "datapower_assembly_action_jwt_validate" "test" {
  id = "ResTestAssemblyActionJWTValidate"
  app_domain = "acceptance_test"
  jwt = "request.headers.authorization"
  output_claims = "decoded.claims"
}`,
    Data: `
data "datapower_assembly_action_jwt_validate" "test" {
  id         = "AccTest_AssemblyActionJWTValidate"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_jwt_validate" "acc_test" {
  id = "AccTest_AssemblyActionJWTValidate"
  app_domain = datapower_domain.acc_test.app_domain
  jwt = "request.headers.authorization"
  output_claims = "decoded.claims"
}`,
    ModelOnly:    false,
}
var AssemblyActionJson2XmlTestConfig = ModelTestConfig{
    Name:         "AssemblyActionJson2Xml",
    Resource: `
resource "datapower_assembly_action_json2xml" "test" {
  id = "ResTestAssemblyActionJson2Xml"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_json2xml" "test" {
  id         = "AccTest_AssemblyActionJson2Xml"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_json2xml" "acc_test" {
  id = "AccTest_AssemblyActionJson2Xml"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionLogTestConfig = ModelTestConfig{
    Name:         "AssemblyActionLog",
    Resource: `
resource "datapower_assembly_action_log" "test" {
  id = "ResTestAssemblyActionLog"
  app_domain = "acceptance_test"
  mode = "gather-only"
}`,
    Data: `
data "datapower_assembly_action_log" "test" {
  id         = "AccTest_AssemblyActionLog"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_log" "acc_test" {
  id = "AccTest_AssemblyActionLog"
  app_domain = datapower_domain.acc_test.app_domain
  mode = "gather-only"
}`,
    ModelOnly:    false,
}
var AssemblyActionMapTestConfig = ModelTestConfig{
    Name:         "AssemblyActionMap",
    Resource: `
resource "datapower_assembly_action_map" "test" {
  id = "ResTestAssemblyActionMap"
  app_domain = "acceptance_test"
  location = "local:///file"
}`,
    Data: `
data "datapower_assembly_action_map" "test" {
  id         = "AccTest_AssemblyActionMap"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_map" "acc_test" {
  id = "AccTest_AssemblyActionMap"
  app_domain = datapower_domain.acc_test.app_domain
  location = "local:///file"
}`,
    ModelOnly:    false,
}
var AssemblyActionOAuthTestConfig = ModelTestConfig{
    Name:         "AssemblyActionOAuth",
    Resource: `
resource "datapower_assembly_action_oauth" "test" {
  id = "ResTestAssemblyActionOAuth"
  app_domain = "acceptance_test"
  oauth_provider_settings_reference = {"Default": "AccTest_OAuthProviderSettings"}
}`,
    Data: `
data "datapower_assembly_action_oauth" "test" {
  id         = "AccTest_AssemblyActionOAuth"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_oauth" "acc_test" {
  id = "AccTest_AssemblyActionOAuth"
  app_domain = datapower_domain.acc_test.app_domain
  oauth_provider_settings_reference = {"Default": datapower_oauth_provider_settings.acc_test.id}
}`,
    ModelOnly:    false,
}
var AssemblyActionParseTestConfig = ModelTestConfig{
    Name:         "AssemblyActionParse",
    Resource: `
resource "datapower_assembly_action_parse" "test" {
  id = "ResTestAssemblyActionParse"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_parse" "test" {
  id         = "AccTest_AssemblyActionParse"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_parse" "acc_test" {
  id = "AccTest_AssemblyActionParse"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionRateLimitTestConfig = ModelTestConfig{
    Name:         "AssemblyActionRateLimit",
    Resource: `
resource "datapower_assembly_action_rate_limit" "test" {
  id = "ResTestAssemblyActionRateLimit"
  app_domain = "acceptance_test"
  source = "plan-default"
}`,
    Data: `
data "datapower_assembly_action_rate_limit" "test" {
  id         = "AccTest_AssemblyActionRateLimit"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_rate_limit" "acc_test" {
  id = "AccTest_AssemblyActionRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  source = "plan-default"
}`,
    ModelOnly:    false,
}
var AssemblyActionRateLimitInfoTestConfig = ModelTestConfig{
    Name:         "AssemblyActionRateLimitInfo",
    Resource: `
resource "datapower_assembly_action_rate_limit_info" "test" {
  id = "ResTestAssemblyActionRateLimitInfo"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_rate_limit_info" "test" {
  id         = "AccTest_AssemblyActionRateLimitInfo"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_rate_limit_info" "acc_test" {
  id = "AccTest_AssemblyActionRateLimitInfo"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionRedactTestConfig = ModelTestConfig{
    Name:         "AssemblyActionRedact",
    Resource: `
resource "datapower_assembly_action_redact" "test" {
  id = "ResTestAssemblyActionRedact"
  app_domain = "acceptance_test"
  redact = ` + DmAssemblyActionRedactTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_assembly_action_redact" "test" {
  id         = "AccTest_AssemblyActionRedact"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_redact" "acc_test" {
  id = "AccTest_AssemblyActionRedact"
  app_domain = datapower_domain.acc_test.app_domain
  redact = ` + DmAssemblyActionRedactTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AssemblyActionSetVarTestConfig = ModelTestConfig{
    Name:         "AssemblyActionSetVar",
    Resource: `
resource "datapower_assembly_action_set_var" "test" {
  id = "ResTestAssemblyActionSetVar"
  app_domain = "acceptance_test"
  variable = ` + DmAssemblyActionSetVarTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_assembly_action_set_var" "test" {
  id         = "AccTest_AssemblyActionSetVar"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_set_var" "acc_test" {
  id = "AccTest_AssemblyActionSetVar"
  app_domain = datapower_domain.acc_test.app_domain
  variable = ` + DmAssemblyActionSetVarTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AssemblyActionThrowTestConfig = ModelTestConfig{
    Name:         "AssemblyActionThrow",
    Resource: `
resource "datapower_assembly_action_throw" "test" {
  id = "ResTestAssemblyActionThrow"
  app_domain = "acceptance_test"
  error_id = "errorid"
}`,
    Data: `
data "datapower_assembly_action_throw" "test" {
  id         = "AccTest_AssemblyActionThrow"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_throw" "acc_test" {
  id = "AccTest_AssemblyActionThrow"
  app_domain = datapower_domain.acc_test.app_domain
  error_id = "errorid"
}`,
    ModelOnly:    false,
}
var AssemblyActionUserSecurityTestConfig = ModelTestConfig{
    Name:         "AssemblyActionUserSecurity",
    Resource: `
resource "datapower_assembly_action_user_security" "test" {
  id = "ResTestAssemblyActionUserSecurity"
  app_domain = "acceptance_test"
  factor_id = "default"
  extract_identity_method = "basic"
  user_auth_method = "user-registry"
  user_registry = "AccTest_APIAuthURLRegistry"
  user_az_method = "authenticated"
}`,
    Data: `
data "datapower_assembly_action_user_security" "test" {
  id         = "AccTest_AssemblyActionUserSecurity"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_user_security" "acc_test" {
  id = "AccTest_AssemblyActionUserSecurity"
  app_domain = datapower_domain.acc_test.app_domain
  factor_id = "default"
  extract_identity_method = "basic"
  user_auth_method = "user-registry"
  user_registry = datapower_api_auth_url_registry.acc_test.id
  user_az_method = "authenticated"
}`,
    ModelOnly:    false,
}
var AssemblyActionValidateTestConfig = ModelTestConfig{
    Name:         "AssemblyActionValidate",
    Resource: `
resource "datapower_assembly_action_validate" "test" {
  id = "ResTestAssemblyActionValidate"
  app_domain = "acceptance_test"
  schema = "AccTest_APISchema"
}`,
    Data: `
data "datapower_assembly_action_validate" "test" {
  id         = "AccTest_AssemblyActionValidate"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_validate" "acc_test" {
  id = "AccTest_AssemblyActionValidate"
  app_domain = datapower_domain.acc_test.app_domain
  schema = datapower_api_schema.acc_test.id
}`,
    ModelOnly:    false,
}
var AssemblyActionWSDLTestConfig = ModelTestConfig{
    Name:         "AssemblyActionWSDL",
    Resource: `
resource "datapower_assembly_action_wsdl" "test" {
  id = "ResTestAssemblyActionWSDL"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_wsdl" "test" {
  id         = "AccTest_AssemblyActionWSDL"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_wsdl" "acc_test" {
  id = "AccTest_AssemblyActionWSDL"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyActionWebSocketUpgradeTestConfig = ModelTestConfig{
    Name:         "AssemblyActionWebSocketUpgrade",
    Resource: `
resource "datapower_assembly_action_web_socket_upgrade" "test" {
  id = "ResTestAssemblyActionWebSocketUpgrade"
  app_domain = "acceptance_test"
  url = "https://localhost"
}`,
    Data: `
data "datapower_assembly_action_web_socket_upgrade" "test" {
  id         = "AccTest_AssemblyActionWebSocketUpgrade"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_web_socket_upgrade" "acc_test" {
  id = "AccTest_AssemblyActionWebSocketUpgrade"
  app_domain = datapower_domain.acc_test.app_domain
  url = "https://localhost"
}`,
    ModelOnly:    false,
}
var AssemblyActionXSLTTestConfig = ModelTestConfig{
    Name:         "AssemblyActionXSLT",
    Resource: `
resource "datapower_assembly_action_xslt" "test" {
  id = "ResTestAssemblyActionXSLT"
  app_domain = "acceptance_test"
  stylesheet = "local:///stylesheet"
}`,
    Data: `
data "datapower_assembly_action_xslt" "test" {
  id         = "AccTest_AssemblyActionXSLT"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_xslt" "acc_test" {
  id = "AccTest_AssemblyActionXSLT"
  app_domain = datapower_domain.acc_test.app_domain
  stylesheet = "local:///stylesheet"
}`,
    ModelOnly:    false,
}
var AssemblyActionXml2JsonTestConfig = ModelTestConfig{
    Name:         "AssemblyActionXml2Json",
    Resource: `
resource "datapower_assembly_action_xml2json" "test" {
  id = "ResTestAssemblyActionXml2Json"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_action_xml2json" "test" {
  id         = "AccTest_AssemblyActionXml2Json"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_action_xml2json" "acc_test" {
  id = "AccTest_AssemblyActionXml2Json"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyFunctionTestConfig = ModelTestConfig{
    Name:         "AssemblyFunction",
    Resource: `
resource "datapower_assembly_function" "test" {
  id = "ResTestAssemblyFunction"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_assembly_function" "test" {
  id         = "AccTest_AssemblyFunction"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_function" "acc_test" {
  id = "AccTest_AssemblyFunction"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var AssemblyLogicOperationSwitchTestConfig = ModelTestConfig{
    Name:         "AssemblyLogicOperationSwitch",
    Resource: `
resource "datapower_assembly_logic_operation_switch" "test" {
  id = "ResTestAssemblyLogicOperationSwitch"
  app_domain = "acceptance_test"
  case = ` + DmAssemblyLogicOperationSwitchCaseTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_assembly_logic_operation_switch" "test" {
  id         = "AccTest_AssemblyLogicOperationSwitch"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_logic_operation_switch" "acc_test" {
  id = "AccTest_AssemblyLogicOperationSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = ` + DmAssemblyLogicOperationSwitchCaseTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AssemblyLogicSwitchTestConfig = ModelTestConfig{
    Name:         "AssemblyLogicSwitch",
    Resource: `
resource "datapower_assembly_logic_switch" "test" {
  id = "ResTestAssemblyLogicSwitch"
  app_domain = "acceptance_test"
  case = ` + DmAssemblyLogicExecuteTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_assembly_logic_switch" "test" {
  id         = "AccTest_AssemblyLogicSwitch"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_assembly_logic_switch" "acc_test" {
  id = "AccTest_AssemblyLogicSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = ` + DmAssemblyLogicExecuteTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var AuditLogTestConfig = ModelTestConfig{
    Name:         "AuditLog",
    Resource: `
resource "datapower_audit_log" "test" {
}`,
    Data: `
data "datapower_audit_log" "test" {
}`,
    TestBed: `
resource "datapower_audit_log" "acc_test" {
}`,
    ModelOnly:    false,
}
var B2BCPATestConfig = ModelTestConfig{
    Name:         "B2BCPA",
    Resource: `
resource "datapower_b2b_cpa" "test" {
  id = "ResTestB2BCPA"
  app_domain = "acceptance_test"
  cpa_id = "cpaid"
}`,
    Data: `
data "datapower_b2b_cpa" "test" {
  id         = "AccTest_B2BCPA"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_cpa" "acc_test" {
  id = "AccTest_B2BCPA"
  app_domain = datapower_domain.acc_test.app_domain
  cpa_id = "cpaid"
}`,
    ModelOnly:    false,
}
var B2BCPACollaborationTestConfig = ModelTestConfig{
    Name:         "B2BCPACollaboration",
    Resource: `
resource "datapower_b2b_cpa_collaboration" "test" {
  id = "ResTestB2BCPACollaboration"
  app_domain = "acceptance_test"
  internal_role = "internal"
  external_role = "external"
  service = "service"
  actions = ` + DmCPACollaborationActionTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_b2b_cpa_collaboration" "test" {
  id         = "AccTest_B2BCPACollaboration"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_cpa_collaboration" "acc_test" {
  id = "AccTest_B2BCPACollaboration"
  app_domain = datapower_domain.acc_test.app_domain
  internal_role = "internal"
  external_role = "external"
  service = "service"
  actions = ` + DmCPACollaborationActionTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var B2BCPAReceiverSettingTestConfig = ModelTestConfig{
    Name:         "B2BCPAReceiverSetting",
    Resource: `
resource "datapower_b2b_cpa_receiver_setting" "test" {
  id = "ResTestB2BCPAReceiverSetting"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_b2b_cpa_receiver_setting" "test" {
  id         = "AccTest_B2BCPAReceiverSetting"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_cpa_receiver_setting" "acc_test" {
  id = "AccTest_B2BCPAReceiverSetting"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var B2BCPASenderSettingTestConfig = ModelTestConfig{
    Name:         "B2BCPASenderSetting",
    Resource: `
resource "datapower_b2b_cpa_sender_setting" "test" {
  id = "ResTestB2BCPASenderSetting"
  app_domain = "acceptance_test"
  dest_endpoint_url = "ebms2://somehost/url"
}`,
    Data: `
data "datapower_b2b_cpa_sender_setting" "test" {
  id         = "AccTest_B2BCPASenderSetting"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_cpa_sender_setting" "acc_test" {
  id = "AccTest_B2BCPASenderSetting"
  app_domain = datapower_domain.acc_test.app_domain
  dest_endpoint_url = "ebms2://somehost/url"
}`,
    ModelOnly:    false,
}
var B2BGatewayTestConfig = ModelTestConfig{
    Name:         "B2BGateway",
    Resource: `
resource "datapower_b2b_gateway" "test" {
  id = "ResTestB2BGateway"
  app_domain = "acceptance_test"
  b2b_profiles = [{partner_profile="AccTest_B2BProfile"}]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode = "purgeonly"
  xml_manager = "default"
}`,
    Data: `
data "datapower_b2b_gateway" "test" {
  id         = "AccTest_B2BGateway"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_gateway" "acc_test" {
  id = "AccTest_B2BGateway"
  app_domain = datapower_domain.acc_test.app_domain
  b2b_profiles = [{partner_profile=datapower_b2b_profile.acc_test.id}]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode = "purgeonly"
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var B2BPersistenceTestConfig = ModelTestConfig{
    Name:         "B2BPersistence",
    Resource: `
resource "datapower_b2b_persistence" "test" {
  raid_volume = "raid0"
  storage_size = 1024
  ha_enabled = false
}`,
    Data: `
data "datapower_b2b_persistence" "test" {
}`,
    TestBed: `
resource "datapower_b2b_persistence" "acc_test" {
  raid_volume = "raid0"
  storage_size = 1024
  ha_enabled = false
}`,
    ModelOnly:    false,
}
var B2BProfileTestConfig = ModelTestConfig{
    Name:         "B2BProfile",
    Resource: `
resource "datapower_b2b_profile" "test" {
  id = "ResTestB2BProfile"
  app_domain = "acceptance_test"
  business_ids = ["businessid"]
  destinations = ` + DmB2BDestinationTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_b2b_profile" "test" {
  id         = "AccTest_B2BProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_profile" "acc_test" {
  id = "AccTest_B2BProfile"
  app_domain = datapower_domain.acc_test.app_domain
  business_ids = ["businessid"]
  destinations = ` + DmB2BDestinationTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var B2BProfileGroupTestConfig = ModelTestConfig{
    Name:         "B2BProfileGroup",
    Resource: `
resource "datapower_b2b_profile_group" "test" {
  id = "ResTestB2BProfileGroup"
  app_domain = "acceptance_test"
  b2b_profiles = ` + DmB2BGroupedProfileTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_b2b_profile_group" "test" {
  id         = "AccTest_B2BProfileGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_profile_group" "acc_test" {
  id = "AccTest_B2BProfileGroup"
  app_domain = datapower_domain.acc_test.app_domain
  b2b_profiles = ` + DmB2BGroupedProfileTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var B2BXPathRoutingPolicyTestConfig = ModelTestConfig{
    Name:         "B2BXPathRoutingPolicy",
    Resource: `
resource "datapower_b2b_xpath_routing_policy" "test" {
  id = "ResTestB2BXPathRoutingPolicy"
  app_domain = "acceptance_test"
  sender_xpath = "senderpath"
  receiver_xpath = "senderpath"
}`,
    Data: `
data "datapower_b2b_xpath_routing_policy" "test" {
  id         = "AccTest_B2BXPathRoutingPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_b2b_xpath_routing_policy" "acc_test" {
  id = "AccTest_B2BXPathRoutingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  sender_xpath = "senderpath"
  receiver_xpath = "senderpath"
}`,
    ModelOnly:    false,
}
var CORSPolicyTestConfig = ModelTestConfig{
    Name:         "CORSPolicy",
    Resource: `
resource "datapower_cors_policy" "test" {
  id = "ResTestCORSPolicy"
  app_domain = "acceptance_test"
  rule = ["AccTest_CORSRule"]
}`,
    Data: `
data "datapower_cors_policy" "test" {
  id         = "AccTest_CORSPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_cors_policy" "acc_test" {
  id = "AccTest_CORSPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  rule = [datapower_cors_rule.acc_test.id]
}`,
    ModelOnly:    false,
}
var CORSRuleTestConfig = ModelTestConfig{
    Name:         "CORSRule",
    Resource: `
resource "datapower_cors_rule" "test" {
  id = "ResTestCORSRule"
  app_domain = "acceptance_test"
  allow_origin = ["origin"]
}`,
    Data: `
data "datapower_cors_rule" "test" {
  id         = "AccTest_CORSRule"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_cors_rule" "acc_test" {
  id = "AccTest_CORSRule"
  app_domain = datapower_domain.acc_test.app_domain
  allow_origin = ["origin"]
}`,
    ModelOnly:    false,
}
var CRLFetchTestConfig = ModelTestConfig{
    Name:         "CRLFetch",
    Resource: `
resource "datapower_crl_fetch" "test" {
}`,
    Data: `
data "datapower_crl_fetch" "test" {
}`,
    TestBed: `
resource "datapower_crl_fetch" "acc_test" {
}`,
    ModelOnly:    false,
}
var CertMonitorTestConfig = ModelTestConfig{
    Name:         "CertMonitor",
    Resource: `
resource "datapower_cert_monitor" "test" {
  polling_interval = 1
  reminder_time = 30
  log_level = "warn"
  disable_expired_certs = false
}`,
    Data: `
data "datapower_cert_monitor" "test" {
}`,
    TestBed: `
resource "datapower_cert_monitor" "acc_test" {
  polling_interval = 1
  reminder_time = 30
  log_level = "warn"
  disable_expired_certs = false
}`,
    ModelOnly:    false,
}
var CompileOptionsPolicyTestConfig = ModelTestConfig{
    Name:         "CompileOptionsPolicy",
    Resource: `
resource "datapower_compile_options_policy" "test" {
  id = "ResTestCompileOptionsPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_compile_options_policy" "test" {
  id         = "AccTest_CompileOptionsPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_compile_options_policy" "acc_test" {
  id = "AccTest_CompileOptionsPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var CompileSettingsTestConfig = ModelTestConfig{
    Name:         "CompileSettings",
    Resource: `
resource "datapower_compile_settings" "test" {
  id = "ResTestCompileSettings"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_compile_settings" "test" {
  id         = "AccTest_CompileSettings"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_compile_settings" "acc_test" {
  id = "AccTest_CompileSettings"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var ConfigDeploymentPolicyTestConfig = ModelTestConfig{
    Name:         "ConfigDeploymentPolicy",
    Resource: `
resource "datapower_config_deployment_policy" "test" {
  id = "ResTestConfigDeploymentPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_config_deployment_policy" "test" {
  id         = "AccTest_ConfigDeploymentPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_config_deployment_policy" "acc_test" {
  id = "AccTest_ConfigDeploymentPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var ConfigSequenceTestConfig = ModelTestConfig{
    Name:         "ConfigSequence",
    Resource: `
resource "datapower_config_sequence" "test" {
  id = "ResTestConfigSequence"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_config_sequence" "test" {
  id         = "AccTest_ConfigSequence"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_config_sequence" "acc_test" {
  id = "AccTest_ConfigSequence"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var ConformancePolicyTestConfig = ModelTestConfig{
    Name:         "ConformancePolicy",
    Resource: `
resource "datapower_conformance_policy" "test" {
  id = "ResTestConformancePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_conformance_policy" "test" {
  id         = "AccTest_ConformancePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_conformance_policy" "acc_test" {
  id = "AccTest_ConformancePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var ControlListTestConfig = ModelTestConfig{
    Name:         "ControlList",
    Resource: `
resource "datapower_control_list" "test" {
  id = "ResTestControlList"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_control_list" "test" {
  id         = "AccTest_ControlList"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_control_list" "acc_test" {
  id = "AccTest_ControlList"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var CookieAttributePolicyTestConfig = ModelTestConfig{
    Name:         "CookieAttributePolicy",
    Resource: `
resource "datapower_cookie_attribute_policy" "test" {
  id = "ResTestCookieAttributePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_cookie_attribute_policy" "test" {
  id         = "AccTest_CookieAttributePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_cookie_attribute_policy" "acc_test" {
  id = "AccTest_CookieAttributePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var CountMonitorTestConfig = ModelTestConfig{
    Name:         "CountMonitor",
    Resource: `
resource "datapower_count_monitor" "test" {
  id = "ResTestCountMonitor"
  app_domain = "acceptance_test"
  measure = "requests"
  source = "all"
  max_source_s = 10000
  message_type = "AccTest_MessageType"
}`,
    Data: `
data "datapower_count_monitor" "test" {
  id         = "AccTest_CountMonitor"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_count_monitor" "acc_test" {
  id = "AccTest_CountMonitor"
  app_domain = datapower_domain.acc_test.app_domain
  measure = "requests"
  source = "all"
  max_source_s = 10000
  message_type = datapower_message_type.acc_test.id
}`,
    ModelOnly:    false,
}
var CryptoCertificateTestConfig = ModelTestConfig{
    Name:         "CryptoCertificate",
    Resource: `
resource "datapower_crypto_certificate" "test" {
  id = "ResTestCryptoCertificate"
  app_domain = "acceptance_test"
  filename = "cert:///acc-test-server.crt"
}`,
    Data: `
data "datapower_crypto_certificate" "test" {
  id         = "AccTest_CryptoCertificate"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_certificate" "acc_test" {
  id = "AccTest_CryptoCertificate"
  app_domain = datapower_domain.acc_test.app_domain
  filename = datapower_file.acc_test_server_crt.remote_path
}`,
    ModelOnly:    false,
}
var CryptoFWCredTestConfig = ModelTestConfig{
    Name:         "CryptoFWCred",
    Resource: `
resource "datapower_crypto_fw_cred" "test" {
  id = "ResTestCryptoFWCred"
  app_domain = "acceptance_test"
  private_key = ["AccTest_CryptoKey"]
}`,
    Data: `
data "datapower_crypto_fw_cred" "test" {
  id         = "AccTest_CryptoFWCred"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_fw_cred" "acc_test" {
  id = "AccTest_CryptoFWCred"
  app_domain = datapower_domain.acc_test.app_domain
  private_key = [datapower_crypto_key.acc_test.id]
}`,
    ModelOnly:    false,
}
var CryptoIdentCredTestConfig = ModelTestConfig{
    Name:         "CryptoIdentCred",
    Resource: `
resource "datapower_crypto_ident_cred" "test" {
  id = "ResTestCryptoIdentCred"
  app_domain = "acceptance_test"
  key = "AccTest_CryptoKey"
  certificate = "AccTest_CryptoCertificate"
}`,
    Data: `
data "datapower_crypto_ident_cred" "test" {
  id         = "AccTest_CryptoIdentCred"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_ident_cred" "acc_test" {
  id = "AccTest_CryptoIdentCred"
  app_domain = datapower_domain.acc_test.app_domain
  key = datapower_crypto_key.acc_test.id
  certificate = datapower_crypto_certificate.acc_test.id
}`,
    ModelOnly:    false,
}
var CryptoKerberosKDCTestConfig = ModelTestConfig{
    Name:         "CryptoKerberosKDC",
    Resource: `
resource "datapower_crypto_kerberos_kdc" "test" {
  id = "ResTestCryptoKerberosKDC"
  app_domain = "acceptance_test"
  realm = "realm"
  server = "localhost"
}`,
    Data: `
data "datapower_crypto_kerberos_kdc" "test" {
  id         = "AccTest_CryptoKerberosKDC"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_kerberos_kdc" "acc_test" {
  id = "AccTest_CryptoKerberosKDC"
  app_domain = datapower_domain.acc_test.app_domain
  realm = "realm"
  server = "localhost"
}`,
    ModelOnly:    false,
}
var CryptoKerberosKeytabTestConfig = ModelTestConfig{
    Name:         "CryptoKerberosKeytab",
    Resource: `
resource "datapower_crypto_kerberos_keytab" "test" {
  id = "ResTestCryptoKerberosKeytab"
  app_domain = "acceptance_test"
  filename = "cert:///kerberos-keytab"
}`,
    Data: `
data "datapower_crypto_kerberos_keytab" "test" {
  id         = "AccTest_CryptoKerberosKeytab"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_kerberos_keytab" "acc_test" {
  id = "AccTest_CryptoKerberosKeytab"
  app_domain = datapower_domain.acc_test.app_domain
  filename = datapower_file.kerberos-keytab.remote_path
}`,
    ModelOnly:    false,
}
var CryptoKeyTestConfig = ModelTestConfig{
    Name:         "CryptoKey",
    Resource: `
resource "datapower_crypto_key" "test" {
  id = "ResTestCryptoKey"
  app_domain = "acceptance_test"
  filename = "cert:///acc-test-server.key"
  alias = "AccTest_PasswordAlias"
}`,
    Data: `
data "datapower_crypto_key" "test" {
  id         = "AccTest_CryptoKey"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_key" "acc_test" {
  id = "AccTest_CryptoKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename = datapower_file.acc_test_server_key.remote_path
  alias = datapower_password_alias.acc_test.id
}`,
    ModelOnly:    false,
}
var CryptoSSKeyTestConfig = ModelTestConfig{
    Name:         "CryptoSSKey",
    Resource: `
resource "datapower_crypto_sskey" "test" {
  id = "ResTestCryptoSSKey"
  app_domain = "acceptance_test"
  filename = "cert://tokensecret"
}`,
    Data: `
data "datapower_crypto_sskey" "test" {
  id         = "AccTest_CryptoSSKey"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_sskey" "acc_test" {
  id = "AccTest_CryptoSSKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename = datapower_file.token-secret.remote_path
}`,
    ModelOnly:    false,
}
var CryptoValCredTestConfig = ModelTestConfig{
    Name:         "CryptoValCred",
    Resource: `
resource "datapower_crypto_val_cred" "test" {
  id = "ResTestCryptoValCred"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_crypto_val_cred" "test" {
  id         = "AccTest_CryptoValCred"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_crypto_val_cred" "acc_test" {
  id = "AccTest_CryptoValCred"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var DNSNameServiceTestConfig = ModelTestConfig{
    Name:         "DNSNameService",
    Resource: `
resource "datapower_dns_name_service" "test" {
  load_balance_algorithm = "first-alive"
}`,
    Data: `
data "datapower_dns_name_service" "test" {
}`,
    ModelTestBed: `{
  load_balance_algorithm = "first-alive"
}`,
    ModelOnly:    false,
}
var DeploymentPolicyParametersBindingTestConfig = ModelTestConfig{
    Name:         "DeploymentPolicyParametersBinding",
    Resource: `
resource "datapower_deployment_policy_parameters_binding" "test" {
  id = "ResTestDeploymentPolicyParametersBinding"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_deployment_policy_parameters_binding" "test" {
  id         = "AccTest_DeploymentPolicyParametersBinding"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_deployment_policy_parameters_binding" "acc_test" {
  id = "AccTest_DeploymentPolicyParametersBinding"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var DistributedVariableTestConfig = ModelTestConfig{
    Name:         "DistributedVariable",
    Resource: `
resource "datapower_distributed_variable" "test" {
  app_domain = "acceptance_test"
  gateway_peering = "default-gateway-peering"
}`,
    Data: `
data "datapower_distributed_variable" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_distributed_variable" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  gateway_peering = "default-gateway-peering"
}`,
    ModelOnly:    false,
}
var DmAAAPAuthenticateTestConfig = ModelTestConfig{
    Name:         "DmAAAPAuthenticate",
    Model: `{
  au_method = "ldap"
  au_cache_allow = "absolute"
}`,
    ModelTestBed: `{
  au_method = "ldap"
  au_cache_allow = "absolute"
}`,
    ModelOnly:    true,
}
var DmAAAPAuthorizeTestConfig = ModelTestConfig{
    Name:         "DmAAAPAuthorize",
    Model: `{
  az_method = "anyauthenticated"
  az_cache_allow = "absolute"
}`,
    ModelTestBed: `{
  az_method = "anyauthenticated"
  az_cache_allow = "absolute"
}`,
    ModelOnly:    true,
}
var DmAAAPEIBitmapTestConfig = ModelTestConfig{
    Name:         "DmAAAPEIBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAAAPERBitmapTestConfig = ModelTestConfig{
    Name:         "DmAAAPERBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAAAPExtractIdentityTestConfig = ModelTestConfig{
    Name:         "DmAAAPExtractIdentity",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAAAPExtractResourceTestConfig = ModelTestConfig{
    Name:         "DmAAAPExtractResource",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAAAPMapCredentialsTestConfig = ModelTestConfig{
    Name:         "DmAAAPMapCredentials",
    Model: `{
  mc_method = "none"
}`,
    ModelTestBed: `{
  mc_method = "none"
}`,
    ModelOnly:    true,
}
var DmAAAPMapResourceTestConfig = ModelTestConfig{
    Name:         "DmAAAPMapResource",
    Model: `{
  mr_method = "none"
}`,
    ModelTestBed: `{
  mr_method = "none"
}`,
    ModelOnly:    true,
}
var DmAAAPPostProcessTestConfig = ModelTestConfig{
    Name:         "DmAAAPPostProcess",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAAATransactionPriorityTestConfig = ModelTestConfig{
    Name:         "DmAAATransactionPriority",
    Model: `{
  credential = "cred"
  priority = "unknown"
}`,
    ModelTestBed: `{
  credential = "cred"
  priority = "unknown"
}`,
    ModelOnly:    true,
}
var DmACETestConfig = ModelTestConfig{
    Name:         "DmACE",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAPIBurstLimitTestConfig = ModelTestConfig{
    Name:         "DmAPIBurstLimit",
    Model: `{
  burst = 1000
}`,
    ModelTestBed: `{
  burst = 1000
}`,
    ModelOnly:    true,
}
var DmAPICGSProxyPolicyTestConfig = ModelTestConfig{
    Name:         "DmAPICGSProxyPolicy",
    Model: `{
  remote_address = "localhost"
  remote_port = 8080
}`,
    ModelTestBed: `{
  remote_address = "localhost"
  remote_port = 8080
}`,
    ModelOnly:    true,
}
var DmAPICountLimitTestConfig = ModelTestConfig{
    Name:         "DmAPICountLimit",
    Model: `{
  count = 1000
}`,
    ModelTestBed: `{
  count = 1000
}`,
    ModelOnly:    true,
}
var DmAPIDataTypeDefinitionTestConfig = ModelTestConfig{
    Name:         "DmAPIDataTypeDefinition",
    Model: `{
  name = "dtdefname"
  schema = "AccTest_APISchema"
}`,
    ModelTestBed: `{
  name = "dtdefname"
  schema = datapower_api_schema.acc_test.id
}`,
    ModelOnly:    true,
}
var DmAPIParameterTestConfig = ModelTestConfig{
    Name:         "DmAPIParameter",
    Model: `{
  required = false
  type = "string"
  where = "path"
}`,
    ModelTestBed: `{
  required = false
  type = "string"
  where = "path"
}`,
    ModelOnly:    true,
}
var DmAPIPropertyTestConfig = ModelTestConfig{
    Name:         "DmAPIProperty",
    Model: `{
  property_name = "propertyname"
  catalog = "*"
}`,
    ModelTestBed: `{
  property_name = "propertyname"
  catalog = "*"
}`,
    ModelOnly:    true,
}
var DmAPIProtocolsTestConfig = ModelTestConfig{
    Name:         "DmAPIProtocols",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAPIProxyPolicyTestConfig = ModelTestConfig{
    Name:         "DmAPIProxyPolicy",
    Model: `{
  reg_exp = "*"
  skip = false
}`,
    ModelTestBed: `{
  reg_exp = "*"
  skip = false
}`,
    ModelOnly:    true,
}
var DmAPIRateLimitTestConfig = ModelTestConfig{
    Name:         "DmAPIRateLimit",
    Model: `{
  rate = 1000
}`,
    ModelTestBed: `{
  rate = 1000
}`,
    ModelOnly:    true,
}
var DmAPIResponseSchemaTestConfig = ModelTestConfig{
    Name:         "DmAPIResponseSchema",
    Model: `{
  status_code = "200"
}`,
    ModelTestBed: `{
  status_code = "200"
}`,
    ModelOnly:    true,
}
var DmASFrontProtocolTestConfig = ModelTestConfig{
    Name:         "DmASFrontProtocol",
    Model: `{
  mdn_receiver = false
}`,
    ModelTestBed: `{
  mdn_receiver = false
}`,
    ModelOnly:    true,
}
var DmAddHeaderPolicyTestConfig = ModelTestConfig{
    Name:         "DmAddHeaderPolicy",
    Model: `{
  reg_exp = "*"
  add_header = "HEADER"
  add_value = "VALUE"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  add_header = "HEADER"
  add_value = "VALUE"
}`,
    ModelOnly:    true,
}
var DmAllowCompressionPolicyTestConfig = ModelTestConfig{
    Name:         "DmAllowCompressionPolicy",
    Model: `{
  reg_exp = "*"
  allow_compression = false
}`,
    ModelTestBed: `{
  reg_exp = "*"
  allow_compression = false
}`,
    ModelOnly:    true,
}
var DmAllowedClientTypeTestConfig = ModelTestConfig{
    Name:         "DmAllowedClientType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmAssemblyActionExtractTestConfig = ModelTestConfig{
    Name:         "DmAssemblyActionExtract",
    Model: `{
  capture = "capture"
}`,
    ModelTestBed: `{
  capture = "capture"
}`,
    ModelOnly:    true,
}
var DmAssemblyActionFunctionCallParameterTestConfig = ModelTestConfig{
    Name:         "DmAssemblyActionFunctionCallParameter",
    Model: `{
  name = "assemblyactionfunctioncallparametername"
}`,
    ModelTestBed: `{
  name = "assemblyactionfunctioncallparametername"
}`,
    ModelOnly:    true,
}
var DmAssemblyActionRedactTestConfig = ModelTestConfig{
    Name:         "DmAssemblyActionRedact",
    Model: `{
  path = "path"
}`,
    ModelTestBed: `{
  path = "path"
}`,
    ModelOnly:    true,
}
var DmAssemblyActionSetVarTestConfig = ModelTestConfig{
    Name:         "DmAssemblyActionSetVar",
    Model: `{
  name = "varname"
  value = "value"
}`,
    ModelTestBed: `{
  name = "varname"
  value = "value"
}`,
    ModelOnly:    true,
}
var DmAssemblyCatchTestConfig = ModelTestConfig{
    Name:         "DmAssemblyCatch",
    Model: `{
  error = "errorname"
  handler = "default-api-rule"
}`,
    ModelTestBed: `{
  error = "errorname"
  handler = "default-api-rule"
}`,
    ModelOnly:    true,
}
var DmAssemblyFunctionParameterTestConfig = ModelTestConfig{
    Name:         "DmAssemblyFunctionParameter",
    Model: `{
  name = "assemblyfunctionparametername"
}`,
    ModelTestBed: `{
  name = "assemblyfunctionparametername"
}`,
    ModelOnly:    true,
}
var DmAssemblyLogicExecuteTestConfig = ModelTestConfig{
    Name:         "DmAssemblyLogicExecute",
    Model: `{
  condition = "condition"
  execute = "default-api-rule"
}`,
    ModelTestBed: `{
  condition = "condition"
  execute = "default-api-rule"
}`,
    ModelOnly:    true,
}
var DmAssemblyLogicOperationSwitchCaseTestConfig = ModelTestConfig{
    Name:         "DmAssemblyLogicOperationSwitchCase",
    Model: `{
  execute = "default-api-rule"
  operation_id = "operationid"
}`,
    ModelTestBed: `{
  execute = "default-api-rule"
  operation_id = "operationid"
}`,
    ModelOnly:    true,
}
var DmB2BActiveGroupTestConfig = ModelTestConfig{
    Name:         "DmB2BActiveGroup",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmB2BActiveProfileTestConfig = ModelTestConfig{
    Name:         "DmB2BActiveProfile",
    Model: `{
  partner_profile = "AccTest_B2BProfile"
}`,
    ModelTestBed: `{
  partner_profile = datapower_b2b_profile.acc_test.id
}`,
    ModelOnly:    true,
}
var DmB2BBackupMsgTypeTestConfig = ModelTestConfig{
    Name:         "DmB2BBackupMsgType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmB2BCPAEntryTestConfig = ModelTestConfig{
    Name:         "DmB2BCPAEntry",
    Model: `{
  cpa = "AccTest_B2BCPA"
  collaboration = "AccTest_B2BCPACollaboration"
  internal_partner = "AccTest_B2BProfile"
  external_partner = "AccTest_B2BProfile"
}`,
    ModelTestBed: `{
  cpa = datapower_b2b_cpa.acc_test.id
  collaboration = datapower_b2b_cpa_collaboration.acc_test.id
  internal_partner = datapower_b2b_profile.acc_test.id
  external_partner = datapower_b2b_profile.acc_test.id
}`,
    ModelOnly:    true,
}
var DmB2BContactTestConfig = ModelTestConfig{
    Name:         "DmB2BContact",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmB2BDestinationTestConfig = ModelTestConfig{
    Name:         "DmB2BDestination",
    Model: `{
  dest_name = "b2bdestinationname"
  dest_url = "https://localhost"
  ssl_client = "AccTest_SSLClientProfile"
}`,
    ModelTestBed: `{
  dest_name = "b2bdestinationname"
  dest_url = "https://localhost"
  ssl_client = datapower_ssl_client_profile.acc_test.id
}`,
    ModelOnly:    true,
}
var DmB2BEnabledDocTypeTestConfig = ModelTestConfig{
    Name:         "DmB2BEnabledDocType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmB2BGroupedProfileTestConfig = ModelTestConfig{
    Name:         "DmB2BGroupedProfile",
    Model: `{
  partner_profile = "AccTest_B2BProfile"
}`,
    ModelTestBed: `{
  partner_profile = datapower_b2b_profile.acc_test.id
}`,
    ModelOnly:    true,
}
var DmB2BHAHostTestConfig = ModelTestConfig{
    Name:         "DmB2BHAHost",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmB2BMessagePropertiesTestConfig = ModelTestConfig{
    Name:         "DmB2BMessageProperties",
    Model: `{
  name = "b2bmessagepropertyname"
  value = "value"
}`,
    ModelTestBed: `{
  name = "b2bmessagepropertyname"
  value = "value"
}`,
    ModelOnly:    true,
}
var DmBasicAuthPolicyTestConfig = ModelTestConfig{
    Name:         "DmBasicAuthPolicy",
    Model: `{
  reg_exp = "*"
  user_name = "someuser"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  user_name = "someuser"
}`,
    ModelOnly:    true,
}
var DmCORSRuleExposeHeadersTestConfig = ModelTestConfig{
    Name:         "DmCORSRuleExposeHeaders",
    Model: `{
  predefined = true
  backend = false
}`,
    ModelTestBed: `{
  predefined = true
  backend = false
}`,
    ModelOnly:    true,
}
var DmCPACollaborationActionTestConfig = ModelTestConfig{
    Name:         "DmCPACollaborationAction",
    Model: `{
  name = "cpacollaborationactionname"
  value = "value"
  capability = "cansend"
  sender_setting = "AccTest_B2BCPASenderSetting"
}`,
    ModelTestBed: `{
  name = "cpacollaborationactionname"
  value = "value"
  capability = "cansend"
  sender_setting = datapower_b2b_cpa_sender_setting.acc_test.id
}`,
    ModelOnly:    true,
}
var DmCRLFetchConfigTestConfig = ModelTestConfig{
    Name:         "DmCRLFetchConfig",
    Model: `{
  issuer_valcred = "AccTest_CryptoValCred"
  refresh_interval = 240
}`,
    ModelTestBed: `{
  issuer_valcred = datapower_crypto_valcred.acc_test.id
  refresh_interval = 240
}`,
    ModelOnly:    true,
}
var DmClaimTestConfig = ModelTestConfig{
    Name:         "DmClaim",
    Model: `{
  value = "value"
}`,
    ModelTestBed: `{
  value = "value"
}`,
    ModelOnly:    true,
}
var DmConditionTestConfig = ModelTestConfig{
    Name:         "DmCondition",
    Model: `{
  expression = "*"
  condition_action = "*"
}`,
    ModelTestBed: `{
  expression = "*"
  condition_action = "*"
}`,
    ModelOnly:    true,
}
var DmConfigModifyTypeTestConfig = ModelTestConfig{
    Name:         "DmConfigModifyType",
    Model: `{
  match = "10.10.1.1/domainA/services/xslproxy?Value=myhost"
  type = "change"
}`,
    ModelTestBed: `{
  match = "10.10.1.1/domainA/services/xslproxy?Value=myhost"
  type = "change"
}`,
    ModelOnly:    true,
}
var DmConfigSequenceCapabilitiesTestConfig = ModelTestConfig{
    Name:         "DmConfigSequenceCapabilities",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmConfigSequenceLocationTestConfig = ModelTestConfig{
    Name:         "DmConfigSequenceLocation",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmConformanceProfilesTestConfig = ModelTestConfig{
    Name:         "DmConformanceProfiles",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmCookieAttributeTestConfig = ModelTestConfig{
    Name:         "DmCookieAttribute",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmCookieProfileTestConfig = ModelTestConfig{
    Name:         "DmCookieProfile",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmCountLimitInfoTestConfig = ModelTestConfig{
    Name:         "DmCountLimitInfo",
    Model: `{
  name = "countlimitinfoname"
  action = "inc"
}`,
    ModelTestBed: `{
  name = "countlimitinfoname"
  action = "inc"
}`,
    ModelOnly:    true,
}
var DmCountMonitorFilterTestConfig = ModelTestConfig{
    Name:         "DmCountMonitorFilter",
    Model: `{
  name = "Filter1"
  interval = 1000
  rate_limit = 50
  burst_limit = 100
  action = "AccTest_FilterAction"
}`,
    ModelTestBed: `{
  name = "Filter1"
  interval = 1000
  rate_limit = 50
  burst_limit = 100
  action = datapower_filter_action.acc_test.id
}`,
    ModelOnly:    true,
}
var DmDefinitionLinkTestConfig = ModelTestConfig{
    Name:         "DmDefinitionLink",
    Model: `{
  short_name = "shortname"
  definition = "AccTest_RateLimitDefinition"
}`,
    ModelTestBed: `{
  short_name = "shortname"
  definition = datapower_rate_limit_definition.acc_test.id
}`,
    ModelOnly:    true,
}
var DmDeploymentPolicyParameterTestConfig = ModelTestConfig{
    Name:         "DmDeploymentPolicyParameter",
    Model: `{
  parameter_name = "ResTestparameter"
  parameter_value = "parameter_value"
}`,
    ModelTestBed: `{
  parameter_name = "ResTestparameter"
  parameter_value = "parameter_value"
}`,
    ModelOnly:    true,
}
var DmDocCachePolicyTestConfig = ModelTestConfig{
    Name:         "DmDocCachePolicy",
    Model: `{
  type = "protocol"
  priority = 128
}`,
    ModelTestBed: `{
  type = "protocol"
  priority = 128
}`,
    ModelOnly:    true,
}
var DmDomainFileMapTestConfig = ModelTestConfig{
    Name:         "DmDomainFileMap",
    Model: `{
  copy_from = true
  copy_to = true
  delete = true
  display = true
  exec = true
  subdir = true
}`,
    ModelTestBed: `{
  copy_from = true
  copy_to = true
  delete = true
  display = true
  exec = true
  subdir = true
}`,
    ModelOnly:    true,
}
var DmDomainMonitoringMapTestConfig = ModelTestConfig{
    Name:         "DmDomainMonitoringMap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmDurationMonitorFilterTestConfig = ModelTestConfig{
    Name:         "DmDurationMonitorFilter",
    Model: `{
  name = "ResTestDmDurationMonitorFilter"
  value = 1
  action = "AccTest_FilterAction"
}`,
    ModelTestBed: `{
  name = "ResTestDmDurationMonitorFilter"
  value = 1
  action = datapower_filter_action.acc_test.id
}`,
    ModelOnly:    true,
}
var DmDynamicOAuthProviderSettingsReferenceTestConfig = ModelTestConfig{
    Name:         "DmDynamicOAuthProviderSettingsReference",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmDynamicParseSettingsReferenceTestConfig = ModelTestConfig{
    Name:         "DmDynamicParseSettingsReference",
    Model: `{
  url = "some_url"
}`,
    ModelTestBed: `{
  url = "some_url"
}`,
    ModelOnly:    true,
}
var DmDynamicStylePolicyActionBaseReferenceTestConfig = ModelTestConfig{
    Name:         "DmDynamicStylePolicyActionBaseReference",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmExtensionFunctionTestConfig = ModelTestConfig{
    Name:         "DmExtensionFunction",
    Model: `{
  extension_function_namespace = "http://www.fredspace/extensions"
  extension_function = "nodeset()"
  local_function = "node-set()"
}`,
    ModelTestBed: `{
  extension_function_namespace = "http://www.fredspace/extensions"
  extension_function = "nodeset()"
  local_function = "node-set()"
}`,
    ModelOnly:    true,
}
var DmExternalAttachedPolicyTestConfig = ModelTestConfig{
    Name:         "DmExternalAttachedPolicy",
    Model: `{
  external_attach_policy_url = "some.url"
}`,
    ModelTestBed: `{
  external_attach_policy_url = "some.url"
}`,
    ModelOnly:    true,
}
var DmFTPPolicyTestConfig = ModelTestConfig{
    Name:         "DmFTPPolicy",
    Model: `{
  reg_exp = "*"
  passive = "pasv-req"
  auth_tls = "auth-off"
  use_ccc = "ccc-off"
  encrypt_data = "enc-data-off"
  data_type = "binary"
  slash_stou = "slash-stou-on"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  passive = "pasv-req"
  auth_tls = "auth-off"
  use_ccc = "ccc-off"
  encrypt_data = "enc-data-off"
  data_type = "binary"
  slash_stou = "slash-stou-on"
}`,
    ModelOnly:    true,
}
var DmFTPQuotedCommandTestConfig = ModelTestConfig{
    Name:         "DmFTPQuotedCommand",
    Model: `{
  quoted_command = "ls"
}`,
    ModelTestBed: `{
  quoted_command = "ls"
}`,
    ModelOnly:    true,
}
var DmFTPServerVirtualDirectoryTestConfig = ModelTestConfig{
    Name:         "DmFTPServerVirtualDirectory",
    Model: `{
  virtual_path = "virtualpath"
}`,
    ModelTestBed: `{
  virtual_path = "virtualpath"
}`,
    ModelOnly:    true,
}
var DmFileSystemUsageTestConfig = ModelTestConfig{
    Name:         "DmFileSystemUsage",
    Model: `{
  warning_threshold = 75
  critical_threshold = 90
}`,
    ModelTestBed: `{
  warning_threshold = 75
  critical_threshold = 90
}`,
    ModelOnly:    true,
}
var DmFrontSideTestConfig = ModelTestConfig{
    Name:         "DmFrontSide",
    Model: `{
  local_address = "0.0.0.0"
  local_port = 3000
}`,
    ModelTestBed: `{
  local_address = "0.0.0.0"
  local_port = 3000
}`,
    ModelOnly:    true,
}
var DmGatewayInOrderModeTestConfig = ModelTestConfig{
    Name:         "DmGatewayInOrderMode",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmGatewayPeeringClusterNodeTestConfig = ModelTestConfig{
    Name:         "DmGatewayPeeringClusterNode",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmGatewayPeeringGroupClusterNodeTestConfig = ModelTestConfig{
    Name:         "DmGatewayPeeringGroupClusterNode",
    Model: `{
  host = "localhost"
}`,
    ModelTestBed: `{
  host = "localhost"
}`,
    ModelOnly:    true,
}
var DmGatewayPeeringGroupPeerNodeTestConfig = ModelTestConfig{
    Name:         "DmGatewayPeeringGroupPeerNode",
    Model: `{
  host = "10.10.10.10"
  priority = 100
}`,
    ModelTestBed: `{
  host = "10.10.10.10"
  priority = 100
}`,
    ModelOnly:    true,
}
var DmGitOpsTemplateEntryTestConfig = ModelTestConfig{
    Name:         "DmGitOpsTemplateEntry",
    Model: `{
  template_type = "change"
}`,
    ModelTestBed: `{
  template_type = "change"
}`,
    ModelOnly:    true,
}
var DmGitOpsTemplatePolicyTestConfig = ModelTestConfig{
    Name:         "DmGitOpsTemplatePolicy",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmGitOpsVariableEntryTestConfig = ModelTestConfig{
    Name:         "DmGitOpsVariableEntry",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmGssChecksumFlagsTestConfig = ModelTestConfig{
    Name:         "DmGssChecksumFlags",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHTTPClientServerVersionTestConfig = ModelTestConfig{
    Name:         "DmHTTPClientServerVersion",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHTTPHeaderTestConfig = ModelTestConfig{
    Name:         "DmHTTPHeader",
    Model: `{
  name = "HEADER"
  value = "VALUE"
}`,
    ModelTestBed: `{
  name = "HEADER"
  value = "VALUE"
}`,
    ModelOnly:    true,
}
var DmHTTPRequestMethodsTestConfig = ModelTestConfig{
    Name:         "DmHTTPRequestMethods",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHTTPResponseCodesTestConfig = ModelTestConfig{
    Name:         "DmHTTPResponseCodes",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHTTPVersionMaskTestConfig = ModelTestConfig{
    Name:         "DmHTTPVersionMask",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHTTPVersionPolicyTestConfig = ModelTestConfig{
    Name:         "DmHTTPVersionPolicy",
    Model: `{
  reg_exp = "*"
  version = "HTTP/1.1"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  version = "HTTP/1.1"
}`,
    ModelOnly:    true,
}
var DmHeaderInjectionTestConfig = ModelTestConfig{
    Name:         "DmHeaderInjection",
    Model: `{
  header_tag_value = "Some Header Value"
}`,
    ModelTestBed: `{
  header_tag_value = "Some Header Value"
}`,
    ModelOnly:    true,
}
var DmHeaderRetentionBitmapTestConfig = ModelTestConfig{
    Name:         "DmHeaderRetentionBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHeaderRetentionPolicyTestConfig = ModelTestConfig{
    Name:         "DmHeaderRetentionPolicy",
    Model: `{
  reg_exp = "*"
}`,
    ModelTestBed: `{
  reg_exp = "*"
}`,
    ModelOnly:    true,
}
var DmHeaderSuppressionTestConfig = ModelTestConfig{
    Name:         "DmHeaderSuppression",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHostKeyAlgorithmsTestConfig = ModelTestConfig{
    Name:         "DmHostKeyAlgorithms",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmHostToSSLServerProfileTestConfig = ModelTestConfig{
    Name:         "DmHostToSSLServerProfile",
    Model: `{
  host_name_wildmat = "hostname_wildmat"
  ssl_server = "AccTest_SSLServerProfile"
}`,
    ModelTestBed: `{
  host_name_wildmat = "hostname_wildmat"
  ssl_server = datapower_ssl_server_profile.acc_test.id
}`,
    ModelOnly:    true,
}
var DmInputEncodingTestConfig = ModelTestConfig{
    Name:         "DmInputEncoding",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmInsertionAttributesTestConfig = ModelTestConfig{
    Name:         "DmInsertionAttributes",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmInvokeErrorTypeTestConfig = ModelTestConfig{
    Name:         "DmInvokeErrorType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmJOSEHeaderTestConfig = ModelTestConfig{
    Name:         "DmJOSEHeader",
    Model: `{
  header_value = "VALUE"
}`,
    ModelTestBed: `{
  header_value = "VALUE"
}`,
    ModelOnly:    true,
}
var DmJWTClaimsTestConfig = ModelTestConfig{
    Name:         "DmJWTClaims",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmJWTGenMethodTestConfig = ModelTestConfig{
    Name:         "DmJWTGenMethod",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmJWTValMethodTestConfig = ModelTestConfig{
    Name:         "DmJWTValMethod",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmKafkaEndpointTestConfig = ModelTestConfig{
    Name:         "DmKafkaEndpoint",
    Model: `{
  host = "localhost"
  port = 8888
}`,
    ModelTestBed: `{
  host = "localhost"
  port = 8888
}`,
    ModelOnly:    true,
}
var DmKafkaPropertyTestConfig = ModelTestConfig{
    Name:         "DmKafkaProperty",
    Model: `{
  name = "name"
  value = "value"
}`,
    ModelTestBed: `{
  name = "name"
  value = "value"
}`,
    ModelOnly:    true,
}
var DmLBGroupAffinityTestConfig = ModelTestConfig{
    Name:         "DmLBGroupAffinity",
    Model: `{
  enable_sa = true
}`,
    ModelTestBed: `{
  enable_sa = true
}`,
    ModelOnly:    true,
}
var DmLBGroupCheckTestConfig = ModelTestConfig{
    Name:         "DmLBGroupCheck",
    Model: `{
  active = false
  port = 80
  ssl = "off"
  timeout = 10
  frequency = 180
}`,
    ModelTestBed: `{
  active = false
  port = 80
  ssl = "off"
  timeout = 10
  frequency = 180
}`,
    ModelOnly:    true,
}
var DmLBGroupMemberTestConfig = ModelTestConfig{
    Name:         "DmLBGroupMember",
    Model: `{
  server = "lbhost"
}`,
    ModelTestBed: `{
  server = "lbhost"
}`,
    ModelOnly:    true,
}
var DmLTPATokenVersionTestConfig = ModelTestConfig{
    Name:         "DmLTPATokenVersion",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmLTPAUserAttributeNameAndValueTestConfig = ModelTestConfig{
    Name:         "DmLTPAUserAttributeNameAndValue",
    Model: `{
  ltpa_user_attribute_name = "name"
  ltpa_user_attribute_type = "static"
  ltpa_user_attribute_static_value = "test"
}`,
    ModelTestBed: `{
  ltpa_user_attribute_name = "name"
  ltpa_user_attribute_type = "static"
  ltpa_user_attribute_static_value = "test"
}`,
    ModelOnly:    true,
}
var DmLogEventTestConfig = ModelTestConfig{
    Name:         "DmLogEvent",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmLogIPFilterTestConfig = ModelTestConfig{
    Name:         "DmLogIPFilter",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmLogObjectTestConfig = ModelTestConfig{
    Name:         "DmLogObject",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmLogTriggerTestConfig = ModelTestConfig{
    Name:         "DmLogTrigger",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmMCFilterTestConfig = ModelTestConfig{
    Name:         "DmMCFilter",
    Model: `{
  filter_name = "ResTestmc_filter"
  type = "http"
}`,
    ModelTestBed: `{
  filter_name = "ResTestmc_filter"
  type = "http"
}`,
    ModelOnly:    true,
}
var DmMQHeadersTestConfig = ModelTestConfig{
    Name:         "DmMQHeaders",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmMSDebugTriggerTypeTestConfig = ModelTestConfig{
    Name:         "DmMSDebugTriggerType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmMatchRuleTestConfig = ModelTestConfig{
    Name:         "DmMatchRule",
    Model: `{
  type = "url"
  error_code = "*"
}`,
    ModelTestBed: `{
  type = "url"
  error_code = "*"
}`,
    ModelOnly:    true,
}
var DmMetaItemTestConfig = ModelTestConfig{
    Name:         "DmMetaItem",
    Model: `{
  meta_category = "all"
  meta_name = "ResTestmeta"
}`,
    ModelTestBed: `{
  meta_category = "all"
  meta_name = "ResTestmeta"
}`,
    ModelOnly:    true,
}
var DmMetadataFromTypeTestConfig = ModelTestConfig{
    Name:         "DmMetadataFromType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmMtomRuleTestConfig = ModelTestConfig{
    Name:         "DmMtomRule",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmMultipartFormDataProfileTestConfig = ModelTestConfig{
    Name:         "DmMultipartFormDataProfile",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmNameServerTestConfig = ModelTestConfig{
    Name:         "DmNameServer",
    Model: `{
  udp_port = 53
  tcp_port = 53
}`,
    ModelTestBed: `{
  udp_port = 53
  tcp_port = 53
}`,
    ModelOnly:    true,
}
var DmNamedInOutTestConfig = ModelTestConfig{
    Name:         "DmNamedInOut",
    Model: `{
  context = "context"
}`,
    ModelTestBed: `{
  context = "context"
}`,
    ModelOnly:    true,
}
var DmNamespaceMappingTestConfig = ModelTestConfig{
    Name:         "DmNamespaceMapping",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmNetworkRetryTestConfig = ModelTestConfig{
    Name:         "DmNetworkRetry",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthAZGrantTypeTestConfig = ModelTestConfig{
    Name:         "DmOAuthAZGrantType",
    Model: `{
  saml20bearer = false
}`,
    ModelTestBed: `{
  saml20bearer = false
}`,
    ModelOnly:    true,
}
var DmOAuthComponentsTestConfig = ModelTestConfig{
    Name:         "DmOAuthComponents",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthFeaturesTestConfig = ModelTestConfig{
    Name:         "DmOAuthFeatures",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthProviderGrantTypeTestConfig = ModelTestConfig{
    Name:         "DmOAuthProviderGrantType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthRSSetHeaderTestConfig = ModelTestConfig{
    Name:         "DmOAuthRSSetHeader",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthRoleTestConfig = ModelTestConfig{
    Name:         "DmOAuthRole",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOAuthTypeTestConfig = ModelTestConfig{
    Name:         "DmOAuthType",
    Model: `{
  client_type = "confidential"
  grant_type = "implicit"
}`,
    ModelTestBed: `{
  client_type = "confidential"
  grant_type = "implicit"
}`,
    ModelOnly:    true,
}
var DmODRConnPropertyTestConfig = ModelTestConfig{
    Name:         "DmODRConnProperty",
    Model: `{
  conn_group_prop_name = "propname"
  conn_group_prop_value = "value"
}`,
    ModelTestBed: `{
  conn_group_prop_name = "propname"
  conn_group_prop_value = "value"
}`,
    ModelOnly:    true,
}
var DmODRConnectorTestConfig = ModelTestConfig{
    Name:         "DmODRConnector",
    Model: `{
  dmgr_hostname = "localhost"
  dmgr_port = 8888
}`,
    ModelTestBed: `{
  dmgr_hostname = "localhost"
  dmgr_port = 8888
}`,
    ModelOnly:    true,
}
var DmODRPropertyTestConfig = ModelTestConfig{
    Name:         "DmODRProperty",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOIDCHybridResponseTypeTestConfig = ModelTestConfig{
    Name:         "DmOIDCHybridResponseType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmOpenTelemetryExporterHeaderTestConfig = ModelTestConfig{
    Name:         "DmOpenTelemetryExporterHeader",
    Model: `{
  header_value = "value"
}`,
    ModelTestBed: `{
  header_value = "value"
}`,
    ModelOnly:    true,
}
var DmOpenTelemetryResourceAttributeTestConfig = ModelTestConfig{
    Name:         "DmOpenTelemetryResourceAttribute",
    Model: `{
  value = "value"
}`,
    ModelTestBed: `{
  value = "value"
}`,
    ModelOnly:    true,
}
var DmPolicyAttachmentPointTestConfig = ModelTestConfig{
    Name:         "DmPolicyAttachmentPoint",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmPolicyMapTestConfig = ModelTestConfig{
    Name:         "DmPolicyMap",
    Model: `{
  match = "__default-accept-service-providers__"
  rule = "__dp-policy-begin__"
}`,
    ModelTestBed: `{
  match = "__default-accept-service-providers__"
  rule = "__dp-policy-begin__"
}`,
    ModelOnly:    true,
}
var DmPolicyParameterTestConfig = ModelTestConfig{
    Name:         "DmPolicyParameter",
    Model: `{
  parameter_name = "PolicyParameterName"
  parameter_value = "PolicyParameterValue"
}`,
    ModelTestBed: `{
  parameter_name = "PolicyParameterName"
  parameter_value = "PolicyParameterValue"
}`,
    ModelOnly:    true,
}
var DmPolicySubjectBitmapTestConfig = ModelTestConfig{
    Name:         "DmPolicySubjectBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmProfileCPABindingTestConfig = ModelTestConfig{
    Name:         "DmProfileCPABinding",
    Model: `{
  internal_partner = "AccTest_B2BProfile"
  cpa = "AccTest_B2BCPA"
  collaboration = "AccTest_B2BCPACollaboration"
}`,
    ModelTestBed: `{
  internal_partner = datapower_b2b_profile.acc_test.id
  cpa = datapower_b2b_cpa.acc_test.id
  collaboration = datapower_b2b_cpa_collaboration.acc_test.id
}`,
    ModelOnly:    true,
}
var DmProxyPolicyTestConfig = ModelTestConfig{
    Name:         "DmProxyPolicy",
    Model: `{
  reg_exp = "*"
  skip = false
  remote_port = 443
}`,
    ModelTestBed: `{
  reg_exp = "*"
  skip = false
  remote_port = 443
}`,
    ModelOnly:    true,
}
var DmPubkeyAuthPolicyTestConfig = ModelTestConfig{
    Name:         "DmPubkeyAuthPolicy",
    Model: `{
  reg_exp = "*"
  crypto_key = "AccTest_CryptoKey"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  crypto_key = datapower_crypto_key.acc_test.id
}`,
    ModelOnly:    true,
}
var DmRBMSSHAuthenticateTypeTestConfig = ModelTestConfig{
    Name:         "DmRBMSSHAuthenticateType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmRadiusServerTestConfig = ModelTestConfig{
    Name:         "DmRadiusServer",
    Model: `{
  secret_wo_version = 1
}`,
    ModelTestBed: `{
  secret_wo_version = 1
}`,
    ModelOnly:    true,
}
var DmRateLimitConfigurationNameValuePairTestConfig = ModelTestConfig{
    Name:         "DmRateLimitConfigurationNameValuePair",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmRateLimitDefinitionNameValuePairTestConfig = ModelTestConfig{
    Name:         "DmRateLimitDefinitionNameValuePair",
    Model: `{
  name = "name"
  value = "value"
}`,
    ModelTestBed: `{
  name = "name"
  value = "value"
}`,
    ModelOnly:    true,
}
var DmRateLimitInfoTestConfig = ModelTestConfig{
    Name:         "DmRateLimitInfo",
    Model: `{
  name = "ratelimitinfoname"
}`,
    ModelTestBed: `{
  name = "ratelimitinfoname"
}`,
    ModelOnly:    true,
}
var DmRateLimitInfoDomainNamedTestConfig = ModelTestConfig{
    Name:         "DmRateLimitInfoDomainNamed",
    Model: `{
  name = "AccTest_RateLimitDefinition"
}`,
    ModelTestBed: `{
  name = datapower_rate_limit_definition.acc_test.id
}`,
    ModelOnly:    true,
}
var DmRoutingPrefixTestConfig = ModelTestConfig{
    Name:         "DmRoutingPrefix",
    Model: `{
  type = "host"
}`,
    ModelTestBed: `{
  type = "host"
}`,
    ModelOnly:    true,
}
var DmSAMLAttributeTestConfig = ModelTestConfig{
    Name:         "DmSAMLAttribute",
    Model: `{
  source_type = "var"
}`,
    ModelTestBed: `{
  source_type = "var"
}`,
    ModelOnly:    true,
}
var DmSAMLAttributeNameAndValueTestConfig = ModelTestConfig{
    Name:         "DmSAMLAttributeNameAndValue",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSAMLStatementTypeTestConfig = ModelTestConfig{
    Name:         "DmSAMLStatementType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSFTPPolicyTestConfig = ModelTestConfig{
    Name:         "DmSFTPPolicy",
    Model: `{
  reg_exp = "*"
  ssh_client_profile = "AccTest_SSHClientProfile"
  use_unique_filenames = false
}`,
    ModelTestBed: `{
  reg_exp = "*"
  ssh_client_profile = datapower_ssh_client_profile.acc_test.id
  use_unique_filenames = false
}`,
    ModelOnly:    true,
}
var DmSFTPServerVirtualDirectoryTestConfig = ModelTestConfig{
    Name:         "DmSFTPServerVirtualDirectory",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSLMStatementTestConfig = ModelTestConfig{
    Name:         "DmSLMStatement",
    Model: `{
  action = "notify"
  maximum_total_reporting_records = 5000
  maximum_resources_and_credentials_for_threshold = 5000
}`,
    ModelTestBed: `{
  action = "notify"
  maximum_total_reporting_records = 5000
  maximum_resources_and_credentials_for_threshold = 5000
}`,
    ModelOnly:    true,
}
var DmSMFlowTestConfig = ModelTestConfig{
    Name:         "DmSMFlow",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSMTPOptionsTestConfig = ModelTestConfig{
    Name:         "DmSMTPOptions",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSMTPPolicyTestConfig = ModelTestConfig{
    Name:         "DmSMTPPolicy",
    Model: `{
  reg_exp = "dpsmtp://*"
}`,
    ModelTestBed: `{
  reg_exp = "dpsmtp://*"
}`,
    ModelOnly:    true,
}
var DmSOAPHeaderDispositionItemTestConfig = ModelTestConfig{
    Name:         "DmSOAPHeaderDispositionItem",
    Model: `{
  action = "processed"
}`,
    ModelTestBed: `{
  action = "processed"
}`,
    ModelOnly:    true,
}
var DmSQLDataSourceConfigNVPairTestConfig = ModelTestConfig{
    Name:         "DmSQLDataSourceConfigNVPair",
    Model: `{
  param_name = "ResTestparam"
  param_value = "param_value"
}`,
    ModelTestBed: `{
  param_name = "ResTestparam"
  param_value = "param_value"
}`,
    ModelOnly:    true,
}
var DmSQLServerTestConfig = ModelTestConfig{
    Name:         "DmSQLServer",
    Model: `{
  host = "sql.host"
  port = 1521
}`,
    ModelTestBed: `{
  host = "sql.host"
  port = 1521
}`,
    ModelOnly:    true,
}
var DmSSHUserAuthenticationMethodsTestConfig = ModelTestConfig{
    Name:         "DmSSHUserAuthenticationMethods",
    Model: `{
  publickey = true
  password = true
}`,
    ModelTestBed: `{
  publickey = true
  password = true
}`,
    ModelOnly:    true,
}
var DmSSLClientFeaturesTestConfig = ModelTestConfig{
    Name:         "DmSSLClientFeatures",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSSLFrontSideTestConfig = ModelTestConfig{
    Name:         "DmSSLFrontSide",
    Model: `{
  local_address = "0.0.0.0"
  local_port = 8888
}`,
    ModelTestBed: `{
  local_address = "0.0.0.0"
  local_port = 8888
}`,
    ModelOnly:    true,
}
var DmSSLHostnameValidationFlagsTestConfig = ModelTestConfig{
    Name:         "DmSSLHostnameValidationFlags",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSSLOptionsTestConfig = ModelTestConfig{
    Name:         "DmSSLOptions",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSSLPolicyTestConfig = ModelTestConfig{
    Name:         "DmSSLPolicy",
    Model: `{
  reg_exp = "*"
}`,
    ModelTestBed: `{
  reg_exp = "*"
}`,
    ModelOnly:    true,
}
var DmSSLProtoVersionsBitmapTestConfig = ModelTestConfig{
    Name:         "DmSSLProtoVersionsBitmap",
    Model: `{
  ssl_v3 = false
  tls_v1d0 = false
  tls_v1d1 = true
  tls_v1d2 = true
  tls_v1d3 = true
}`,
    ModelTestBed: `{
  ssl_v3 = false
  tls_v1d0 = false
  tls_v1d1 = true
  tls_v1d2 = true
  tls_v1d3 = true
}`,
    ModelOnly:    true,
}
var DmScheduledRuleTestConfig = ModelTestConfig{
    Name:         "DmScheduledRule",
    Model: `{
  rule = "__dp-policy-begin__"
}`,
    ModelTestBed: `{
  rule = "__dp-policy-begin__"
}`,
    ModelOnly:    true,
}
var DmSchemaExceptionRuleTestConfig = ModelTestConfig{
    Name:         "DmSchemaExceptionRule",
    Model: `{
  xpath = "*"
  exception_type = "AllowEncrypted"
}`,
    ModelTestBed: `{
  xpath = "*"
  exception_type = "AllowEncrypted"
}`,
    ModelOnly:    true,
}
var DmSchemaValidationTestConfig = ModelTestConfig{
    Name:         "DmSchemaValidation",
    Model: `{
  matching = "__default-accept-service-providers__"
  validation_mode = "default"
}`,
    ModelTestBed: `{
  matching = "__default-accept-service-providers__"
  validation_mode = "default"
}`,
    ModelOnly:    true,
}
var DmSearchDomainTestConfig = ModelTestConfig{
    Name:         "DmSearchDomain",
    Model: `{
  search_domain = "datapower.com"
}`,
    ModelTestBed: `{
  search_domain = "datapower.com"
}`,
    ModelOnly:    true,
}
var DmSecurityTypeTestConfig = ModelTestConfig{
    Name:         "DmSecurityType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSnmpContextTestConfig = ModelTestConfig{
    Name:         "DmSnmpContext",
    Model: `{
  domain = "acc_test_domain"
}`,
    ModelTestBed: `{
  domain = "acc_test_domain"
}`,
    ModelOnly:    true,
}
var DmSnmpCredTestConfig = ModelTestConfig{
    Name:         "DmSnmpCred",
    Model: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
    ModelTestBed: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
    ModelOnly:    true,
}
var DmSnmpCredMaskedTestConfig = ModelTestConfig{
    Name:         "DmSnmpCredMasked",
    Model: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
    ModelTestBed: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
    ModelOnly:    true,
}
var DmSnmpPolicyTestConfig = ModelTestConfig{
    Name:         "DmSnmpPolicy",
    Model: `{
  domain = "acc_test_domain"
}`,
    ModelTestBed: `{
  domain = "acc_test_domain"
}`,
    ModelOnly:    true,
}
var DmSnmpPolicyMQTestConfig = ModelTestConfig{
    Name:         "DmSnmpPolicyMQ",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSnmpTargetTestConfig = ModelTestConfig{
    Name:         "DmSnmpTarget",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSoapActionPolicyTestConfig = ModelTestConfig{
    Name:         "DmSoapActionPolicy",
    Model: `{
  reg_exp = "*"
  soap_action = "*"
}`,
    ModelTestBed: `{
  reg_exp = "*"
  soap_action = "*"
}`,
    ModelOnly:    true,
}
var DmSourceAS2FeatureTypeTestConfig = ModelTestConfig{
    Name:         "DmSourceAS2FeatureType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmSourceHTTPFeatureTypeTestConfig = ModelTestConfig{
    Name:         "DmSourceHTTPFeatureType",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmStaticHostTestConfig = ModelTestConfig{
    Name:         "DmStaticHost",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmStylesheetParameterTestConfig = ModelTestConfig{
    Name:         "DmStylesheetParameter",
    Model: `{
  parameter_value = "PARAMETER-VALUE"
}`,
    ModelTestBed: `{
  parameter_value = "PARAMETER-VALUE"
}`,
    ModelOnly:    true,
}
var DmTAMAZReplicaTestConfig = ModelTestConfig{
    Name:         "DmTAMAZReplica",
    Model: `{
  tam_az_replica = "replica"
  tam_az_replica_port = 7136
  tam_az_replica_weight = 10
}`,
    ModelTestBed: `{
  tam_az_replica = "replica"
  tam_az_replica_port = 7136
  tam_az_replica_weight = 10
}`,
    ModelOnly:    true,
}
var DmTAMFedDirTestConfig = ModelTestConfig{
    Name:         "DmTAMFedDir",
    Model: `{
  fed_name = "ResTestfed"
  suffix = "suffix"
  host = "ldap.host"
  port = 389
  bind_dn = "dn"
  bind_pw = "AccTest_PasswordAlias"
  use_ssl = false
}`,
    ModelTestBed: `{
  fed_name = "ResTestfed"
  suffix = "suffix"
  host = "ldap.host"
  port = 389
  bind_dn = "dn"
  bind_pw = datapower_password_alias.acc_test.id
  use_ssl = false
}`,
    ModelOnly:    true,
}
var DmTAMRASTraceTestConfig = ModelTestConfig{
    Name:         "DmTAMRASTrace",
    Model: `{
  tam_trace_enable = false
  ldap_trace_enable = false
  gs_kit_trace_enable = false
}`,
    ModelTestBed: `{
  tam_trace_enable = false
  ldap_trace_enable = false
  gs_kit_trace_enable = false
}`,
    ModelOnly:    true,
}
var DmTableEntryTestConfig = ModelTestConfig{
    Name:         "DmTableEntry",
    Model: `{
  name = "tableentryname"
}`,
    ModelTestBed: `{
  name = "tableentryname"
}`,
    ModelOnly:    true,
}
var DmTargetMapRuleTestConfig = ModelTestConfig{
    Name:         "DmTargetMapRule",
    Model: `{
  target = "target"
  execute = "default-func-call-global"
}`,
    ModelTestBed: `{
  target = "target"
  execute = "default-func-call-global"
}`,
    ModelOnly:    true,
}
var DmURLMapRuleTestConfig = ModelTestConfig{
    Name:         "DmURLMapRule",
    Model: `{
  pattern = "https://www.company.com/XML/stylesheets/*"
}`,
    ModelTestBed: `{
  pattern = "https://www.company.com/XML/stylesheets/*"
}`,
    ModelOnly:    true,
}
var DmURLRefreshRuleTestConfig = ModelTestConfig{
    Name:         "DmURLRefreshRule",
    Model: `{
  url_map = "default-attempt-stream-all"
}`,
    ModelTestBed: `{
  url_map = "default-attempt-stream-all"
}`,
    ModelOnly:    true,
}
var DmURLRewriteRuleTestConfig = ModelTestConfig{
    Name:         "DmURLRewriteRule",
    Model: `{
  type = "header-rewrite"
  match_regexp = "Test"
}`,
    ModelTestBed: `{
  type = "header-rewrite"
  match_regexp = "Test"
}`,
    ModelOnly:    true,
}
var DmUploadChunkedPolicyTestConfig = ModelTestConfig{
    Name:         "DmUploadChunkedPolicy",
    Model: `{
  reg_exp = "*"
  upload_chunked = false
}`,
    ModelTestBed: `{
  reg_exp = "*"
  upload_chunked = false
}`,
    ModelOnly:    true,
}
var DmValidationFeaturesTestConfig = ModelTestConfig{
    Name:         "DmValidationFeatures",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmValidationTypeTestConfig = ModelTestConfig{
    Name:         "DmValidationType",
    Model: `{
  fixup = "error"
}`,
    ModelTestBed: `{
  fixup = "error"
}`,
    ModelOnly:    true,
}
var DmWSBaseWSDLTestConfig = ModelTestConfig{
    Name:         "DmWSBaseWSDL",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSDLCachePolicyTestConfig = ModelTestConfig{
    Name:         "DmWSDLCachePolicy",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSDLUserPolicyTogglesTestConfig = ModelTestConfig{
    Name:         "DmWSDLUserPolicyToggles",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSEndpointLocalRewriteRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointLocalRewriteRule",
    Model: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
    ModelTestBed: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
    ModelOnly:    true,
}
var DmWSEndpointPublishRewriteRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointPublishRewriteRule",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSEndpointRemoteRewriteRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointRemoteRewriteRule",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSEndpointSubscriptionLocalRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointSubscriptionLocalRule",
    Model: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
    ModelTestBed: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
    ModelOnly:    true,
}
var DmWSEndpointSubscriptionPublishRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointSubscriptionPublishRule",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSEndpointSubscriptionRemoteRuleTestConfig = ModelTestConfig{
    Name:         "DmWSEndpointSubscriptionRemoteRule",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSMPolicyMapTestConfig = ModelTestConfig{
    Name:         "DmWSMPolicyMap",
    Model: `{
  wsdl_component_type = "all"
  match = "__default-accept-service-providers__"
  rule = "AccTest_WSStylePolicyRule"
}`,
    ModelTestBed: `{
  wsdl_component_type = "all"
  match = "__default-accept-service-providers__"
  rule = datapower_ws_style_policy_rule.acc_test.id
}`,
    ModelOnly:    true,
}
var DmWSOperationConformancePolicyTestConfig = ModelTestConfig{
    Name:         "DmWSOperationConformancePolicy",
    Model: `{
  conformance_policy = "AccTest_ConformancePolicy"
  conformance_policy_wsdl_component_type = "all"
}`,
    ModelTestBed: `{
  conformance_policy = datapower_conformance_policy.acc_test.id
  conformance_policy_wsdl_component_type = "all"
}`,
    ModelOnly:    true,
}
var DmWSOperationPolicySubjectOptOutTestConfig = ModelTestConfig{
    Name:         "DmWSOperationPolicySubjectOptOut",
    Model: `{
  policy_subject_opt_out_wsdl_component_type = "all"
}`,
    ModelTestBed: `{
  policy_subject_opt_out_wsdl_component_type = "all"
}`,
    ModelOnly:    true,
}
var DmWSOperationReliableMessagingTestConfig = ModelTestConfig{
    Name:         "DmWSOperationReliableMessaging",
    Model: `{
  reliable_messaging_wsdl_component_type = "all"
}`,
    ModelTestBed: `{
  reliable_messaging_wsdl_component_type = "all"
}`,
    ModelOnly:    true,
}
var DmWSOperationSchedulerPriorityTestConfig = ModelTestConfig{
    Name:         "DmWSOperationSchedulerPriority",
    Model: `{
  scheduler_priority_wsdl_component_type = "all"
}`,
    ModelTestBed: `{
  scheduler_priority_wsdl_component_type = "all"
}`,
    ModelOnly:    true,
}
var DmWSPolicyParametersTestConfig = ModelTestConfig{
    Name:         "DmWSPolicyParameters",
    Model: `{
  policy_param_parameters = "AccTest_PolicyParameters"
  policy_param_wsdl_component_type = "all"
}`,
    ModelTestBed: `{
  policy_param_parameters = datapower_policy_parameters.acc_test.id
  policy_param_wsdl_component_type = "all"
}`,
    ModelOnly:    true,
}
var DmWSRMPolicyBitmapTestConfig = ModelTestConfig{
    Name:         "DmWSRMPolicyBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSRRSavedSearchWSDLSourceTestConfig = ModelTestConfig{
    Name:         "DmWSRRSavedSearchWSDLSource",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSRRWSDLSourceTestConfig = ModelTestConfig{
    Name:         "DmWSRRWSDLSource",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWSSLMOpsTestConfig = ModelTestConfig{
    Name:         "DmWSSLMOps",
    Model: `{
  operation = "all"
  target = "front"
  severity = "low"
}`,
    ModelTestBed: `{
  operation = "all"
  target = "front"
  severity = "low"
}`,
    ModelOnly:    true,
}
var DmWSUserTogglesTestConfig = ModelTestConfig{
    Name:         "DmWSUserToggles",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmWebAppRequestPolicyMapTestConfig = ModelTestConfig{
    Name:         "DmWebAppRequestPolicyMap",
    Model: `{
  match = "__default-accept-service-providers__"
  rule = "AccTest_WebAppRequest"
}`,
    ModelTestBed: `{
  match = "__default-accept-service-providers__"
  rule = datapower_web_app_request.acc_test.id
}`,
    ModelOnly:    true,
}
var DmWebAppResponsePolicyMapTestConfig = ModelTestConfig{
    Name:         "DmWebAppResponsePolicyMap",
    Model: `{
  match = "__default-accept-service-providers__"
  rule = "AccTest_WebAppResponse"
}`,
    ModelTestBed: `{
  match = "__default-accept-service-providers__"
  rule = datapower_web_app_response.acc_test.id
}`,
    ModelOnly:    true,
}
var DmWebGWErrorPolicyMapTestConfig = ModelTestConfig{
    Name:         "DmWebGWErrorPolicyMap",
    Model: `{
  match = "__default-accept-service-providers__"
  action = "AccTest_MPGWErrorAction"
}`,
    ModelTestBed: `{
  match = "__default-accept-service-providers__"
  action = datapower_mpgw_error_action.acc_test.id
}`,
    ModelOnly:    true,
}
var DmWebGWErrorRespHeaderInjectionTestConfig = ModelTestConfig{
    Name:         "DmWebGWErrorRespHeaderInjection",
    Model: `{
  header_tag_value = "HEADER_VALUE"
}`,
    ModelTestBed: `{
  header_tag_value = "HEADER_VALUE"
}`,
    ModelOnly:    true,
}
var DmWebSphereJMSEndpointTestConfig = ModelTestConfig{
    Name:         "DmWebSphereJMSEndpoint",
    Model: `{
  host = "localhost"
  port = 8888
  transport = "TCP"
}`,
    ModelTestBed: `{
  host = "localhost"
  port = 8888
  transport = "TCP"
}`,
    ModelOnly:    true,
}
var DmWeekdayBitmapTestConfig = ModelTestConfig{
    Name:         "DmWeekdayBitmap",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmXMLMgmtModesTestConfig = ModelTestConfig{
    Name:         "DmXMLMgmtModes",
    Model: `{
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    true,
}
var DmXPathRoutingRuleTestConfig = ModelTestConfig{
    Name:         "DmXPathRoutingRule",
    Model: `{
  xpath = "*"
  host = "localhost"
  port = 8888
  ssl = false
}`,
    ModelTestBed: `{
  xpath = "*"
  host = "localhost"
  port = 8888
  ssl = false
}`,
    ModelOnly:    true,
}
var DocumentCryptoMapTestConfig = ModelTestConfig{
    Name:         "DocumentCryptoMap",
    Resource: `
resource "datapower_document_crypto_map" "test" {
  id = "ResTestDocumentCryptoMap"
  app_domain = "acceptance_test"
  operation = "encrypt"
  xpath = ["*",]
}`,
    Data: `
data "datapower_document_crypto_map" "test" {
  id         = "AccTest_DocumentCryptoMap"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_document_crypto_map" "acc_test" {
  id = "AccTest_DocumentCryptoMap"
  app_domain = datapower_domain.acc_test.app_domain
  operation = "encrypt"
  xpath = ["*",]
}`,
    ModelOnly:    false,
}
var DomainTestConfig = ModelTestConfig{
    Name:         "Domain",
    Resource: `
resource "datapower_domain" "test" {
  app_domain = "domain_resource_test"
}`,
    Data: `
data "datapower_domain" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_domain" "acc_test" {
  app_domain = "acceptance_test"
}`,
    ModelOnly:    false,
}
var DomainAvailabilityTestConfig = ModelTestConfig{
    Name:         "DomainAvailability",
    Resource: `
resource "datapower_domain_availability" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_domain_availability" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_domain_availability" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var DomainSettingsTestConfig = ModelTestConfig{
    Name:         "DomainSettings",
    Resource: `
resource "datapower_domain_settings" "test" {
  app_domain = "acceptance_test"
  password_treatment = "masked"
}`,
    Data: `
data "datapower_domain_settings" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_domain_settings" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  password_treatment = "masked"
}`,
    ModelOnly:    false,
}
var DurationMonitorTestConfig = ModelTestConfig{
    Name:         "DurationMonitor",
    Resource: `
resource "datapower_duration_monitor" "test" {
  id = "ResTestDurationMonitor"
  app_domain = "acceptance_test"
  measure = "messages"
  message_type = "AccTest_MessageType"
}`,
    Data: `
data "datapower_duration_monitor" "test" {
  id         = "AccTest_DurationMonitor"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_duration_monitor" "acc_test" {
  id = "AccTest_DurationMonitor"
  app_domain = datapower_domain.acc_test.app_domain
  measure = "messages"
  message_type = datapower_message_type.acc_test.id
}`,
    ModelOnly:    false,
}
var EBMS2SourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "EBMS2SourceProtocolHandler",
    Resource: `
resource "datapower_ebms2_source_protocol_handler" "test" {
  id = "ResTestEBMS2SourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 80
}`,
    Data: `
data "datapower_ebms2_source_protocol_handler" "test" {
  id         = "AccTest_EBMS2SourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ebms2_source_protocol_handler" "acc_test" {
  id = "AccTest_EBMS2SourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 80
}`,
    ModelOnly:    false,
}
var EBMS3SourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "EBMS3SourceProtocolHandler",
    Resource: `
resource "datapower_ebms3_source_protocol_handler" "test" {
  id = "ResTestEBMS3SourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 80
}`,
    Data: `
data "datapower_ebms3_source_protocol_handler" "test" {
  id         = "AccTest_EBMS3SourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ebms3_source_protocol_handler" "acc_test" {
  id = "AccTest_EBMS3SourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 80
}`,
    ModelOnly:    false,
}
var ErrorReportSettingsTestConfig = ModelTestConfig{
    Name:         "ErrorReportSettings",
    Resource: `
resource "datapower_error_report_settings" "test" {
}`,
    Data: `
data "datapower_error_report_settings" "test" {
}`,
    TestBed: `
resource "datapower_error_report_settings" "acc_test" {
}`,
    ModelOnly:    false,
}
var FTPFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "FTPFilePollerSourceProtocolHandler",
    Resource: `
resource "datapower_ftp_file_poller_source_protocol_handler" "test" {
  id = "ResTestFTPFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
  target_directory = "ftp://user:password@host:port/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  result_name_pattern = "../result/$1"
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    Data: `
data "datapower_ftp_file_poller_source_protocol_handler" "test" {
  id         = "AccTest_FTPFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ftp_file_poller_source_protocol_handler" "acc_test" {
  id = "AccTest_FTPFilePollerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  target_directory = "ftp://user:password@host:port/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  result_name_pattern = "../result/$1"
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var FTPQuoteCommandsTestConfig = ModelTestConfig{
    Name:         "FTPQuoteCommands",
    Resource: `
resource "datapower_ftp_quote_commands" "test" {
  id = "ResTestFTPQuoteCommands"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ftp_quote_commands" "test" {
  id         = "AccTest_FTPQuoteCommands"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ftp_quote_commands" "acc_test" {
  id = "AccTest_FTPQuoteCommands"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var FTPServerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "FTPServerSourceProtocolHandler",
    Resource: `
resource "datapower_ftp_server_source_protocol_handler" "test" {
  id = "ResTestFTPServerSourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 21
}`,
    Data: `
data "datapower_ftp_server_source_protocol_handler" "test" {
  id         = "AccTest_FTPServerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ftp_server_source_protocol_handler" "acc_test" {
  id = "AccTest_FTPServerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 21
}`,
    ModelOnly:    false,
}
var FileSystemUsageMonitorTestConfig = ModelTestConfig{
    Name:         "FileSystemUsageMonitor",
    Resource: `
resource "datapower_file_system_usage_monitor" "test" {
  polling_interval = 60
  all_system = true
}`,
    Data: `
data "datapower_file_system_usage_monitor" "test" {
}`,
    TestBed: `
resource "datapower_file_system_usage_monitor" "acc_test" {
  polling_interval = 60
  all_system = true
}`,
    ModelOnly:    false,
}
var FilterActionTestConfig = ModelTestConfig{
    Name:         "FilterAction",
    Resource: `
resource "datapower_filter_action" "test" {
  id = "ResTestFilterAction"
  app_domain = "acceptance_test"
  type = "notify"
}`,
    Data: `
data "datapower_filter_action" "test" {
  id         = "AccTest_FilterAction"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_filter_action" "acc_test" {
  id = "AccTest_FilterAction"
  app_domain = datapower_domain.acc_test.app_domain
  type = "notify"
}`,
    ModelOnly:    false,
}
var FormsLoginPolicyTestConfig = ModelTestConfig{
    Name:         "FormsLoginPolicy",
    Resource: `
resource "datapower_forms_login_policy" "test" {
  id = "ResTestFormsLoginPolicy"
  app_domain = "acceptance_test"
  login_form = "/LoginPage.htm"
  use_cookie_attributes = false
  enable_migration = false
  error_page = "/ErrorPage.htm"
  logout_page = "/LogoutPage.htm"
  default_url = "/"
  username_field = "j_username"
  password_field = "j_password"
  redirect_field = "originalUrl"
  form_processing_url = "/j_security_check"
}`,
    Data: `
data "datapower_forms_login_policy" "test" {
  id         = "AccTest_FormsLoginPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_forms_login_policy" "acc_test" {
  id = "AccTest_FormsLoginPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  login_form = "/LoginPage.htm"
  use_cookie_attributes = false
  enable_migration = false
  error_page = "/ErrorPage.htm"
  logout_page = "/LogoutPage.htm"
  default_url = "/"
  username_field = "j_username"
  password_field = "j_password"
  redirect_field = "originalUrl"
  form_processing_url = "/j_security_check"
}`,
    ModelOnly:    false,
}
var GWSRemoteDebugTestConfig = ModelTestConfig{
    Name:         "GWSRemoteDebug",
    Resource: `
resource "datapower_gws_remote_debug" "test" {
  local_port = 9229
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_gws_remote_debug" "test" {
}`,
    TestBed: `
resource "datapower_gws_remote_debug" "acc_test" {
  local_port = 9229
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var GWScriptSettingsTestConfig = ModelTestConfig{
    Name:         "GWScriptSettings",
    Resource: `
resource "datapower_gw_script_settings" "test" {
}`,
    Data: `
data "datapower_gw_script_settings" "test" {
}`,
    TestBed: `
resource "datapower_gw_script_settings" "acc_test" {
}`,
    ModelOnly:    false,
}
var GatewayPeeringTestConfig = ModelTestConfig{
    Name:         "GatewayPeering",
    Resource: `
resource "datapower_gateway_peering" "test" {
  id = "ResTestGatewayPeering"
  app_domain = "acceptance_test"
  local_address = "5.5.5.5"
  local_port = 16380
  enable_ssl = false
}`,
    Data: `
data "datapower_gateway_peering" "test" {
  id         = "AccTest_GatewayPeering"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_GatewayPeering"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "5.5.5.5"
  local_port = 16380
  enable_ssl = false
}`,
    ModelOnly:    false,
}
var GatewayPeeringGroupTestConfig = ModelTestConfig{
    Name:         "GatewayPeeringGroup",
    Resource: `
resource "datapower_gateway_peering_group" "test" {
  id = "ResTestGatewayPeeringGroup"
  app_domain = "acceptance_test"
  mode = "peer"
  enable_ssl = false
}`,
    Data: `
data "datapower_gateway_peering_group" "test" {
  id         = "AccTest_GatewayPeeringGroup"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_GatewayPeeringGroup"
  app_domain = datapower_domain.acc_test.app_domain
  mode = "peer"
  enable_ssl = false
}`,
    ModelOnly:    false,
}
var GatewayPeeringManagerTestConfig = ModelTestConfig{
    Name:         "GatewayPeeringManager",
    Resource: `
resource "datapower_gateway_peering_manager" "test" {
  app_domain = "acceptance_test"
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit = "default-gateway-peering"
  subscription = "default-gateway-peering"
  ratelimit_module = "default-gateway-peering"
}`,
    Data: `
data "datapower_gateway_peering_manager" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_gateway_peering_manager" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit = "default-gateway-peering"
  subscription = "default-gateway-peering"
  ratelimit_module = "default-gateway-peering"
}`,
    ModelOnly:    false,
}
var GitOpsTestConfig = ModelTestConfig{
    Name:         "GitOps",
    Resource: `
resource "datapower_git_ops" "test" {
  app_domain = "acceptance_test"
  connection_type = "https"
  mode = "read-write"
  commit_identifier = "main"
  remote_location = "https://github.com/ScottW514/terraform-provider-datapower"
  username = "gitusername"
  password = "AccTest_PasswordAlias"
  tls_valcred = "AccTest_CryptoValCred"
  git_user = "Git User"
  git_email = "git@user.domain"
}`,
    Data: `
data "datapower_git_ops" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_git_ops" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  connection_type = "https"
  mode = "read-write"
  commit_identifier = "main"
  remote_location = "https://github.com/ScottW514/terraform-provider-datapower"
  username = "gitusername"
  password = datapower_password_alias.test_acc.id
  tls_valcred = datapower_crypto_val_cred.test_acc.id
  git_user = "Git User"
  git_email = "git@user.domain"
}`,
    ModelOnly:    false,
}
var GitOpsTemplateTestConfig = ModelTestConfig{
    Name:         "GitOpsTemplate",
    Resource: `
resource "datapower_git_ops_template" "test" {
  id = "ResTestGitOpsTemplate"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_git_ops_template" "test" {
  id         = "AccTest_GitOpsTemplate"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_git_ops_template" "acc_test" {
  id = "AccTest_GitOpsTemplate"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var GitOpsVariablesTestConfig = ModelTestConfig{
    Name:         "GitOpsVariables",
    Resource: `
resource "datapower_git_ops_variables" "test" {
}`,
    Data: `
data "datapower_git_ops_variables" "test" {
}`,
    TestBed: `
resource "datapower_git_ops_variables" "acc_test" {
}`,
    ModelOnly:    false,
}
var GraphQLSchemaOptionsTestConfig = ModelTestConfig{
    Name:         "GraphQLSchemaOptions",
    Resource: `
resource "datapower_graphql_schema_options" "test" {
  id = "ResTestGraphQLSchemaOptions"
  app_domain = "acceptance_test"
  api = "AccTest_APIDefinition"
}`,
    Data: `
data "datapower_graphql_schema_options" "test" {
  id         = "AccTest_GraphQLSchemaOptions"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_graphql_schema_options" "acc_test" {
  id = "AccTest_GraphQLSchemaOptions"
  app_domain = datapower_domain.acc_test.app_domain
  api = datapower_api_definition.acc_test.id
}`,
    ModelOnly:    false,
}
var HTTPInputConversionMapTestConfig = ModelTestConfig{
    Name:         "HTTPInputConversionMap",
    Resource: `
resource "datapower_http_input_conversion_map" "test" {
  id = "ResTestHTTPInputConversionMap"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_http_input_conversion_map" "test" {
  id         = "_webappfw_sql-json"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_http_input_conversion_map" "acc_test" {
  id = "AccTest_HTTPInputConversionMap"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var HTTPSSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "HTTPSSourceProtocolHandler",
    Resource: `
resource "datapower_https_source_protocol_handler" "test" {
  id = "ResTestHTTPSSourceProtocolHandler"
  app_domain = "acceptance_test"
  ssl_server_config_type = "server"
  ssl_server = "AccTest_SSLServerProfile"
}`,
    Data: `
data "datapower_https_source_protocol_handler" "test" {
  id         = "AccTest_HTTPSSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_https_source_protocol_handler" "acc_test" {
  id = "AccTest_HTTPSSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  ssl_server_config_type = "server"
  ssl_server = datapower_ssl_server_profile.acc_test.id
}`,
    ModelOnly:    false,
}
var HTTPSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "HTTPSourceProtocolHandler",
    Resource: `
resource "datapower_http_source_protocol_handler" "test" {
  id = "ResTestHTTPSourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 8088
}`,
    Data: `
data "datapower_http_source_protocol_handler" "test" {
  id         = "AccTest_HTTPSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_http_source_protocol_handler" "acc_test" {
  id = "AccTest_HTTPSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 8088
}`,
    ModelOnly:    false,
}
var HTTPUserAgentTestConfig = ModelTestConfig{
    Name:         "HTTPUserAgent",
    Resource: `
resource "datapower_http_user_agent" "test" {
  id = "ResTestHTTPUserAgent"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_http_user_agent" "test" {
  id         = "AccTest_HTTPUserAgent"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_http_user_agent" "acc_test" {
  id = "AccTest_HTTPUserAgent"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var HostAliasTestConfig = ModelTestConfig{
    Name:         "HostAlias",
    Resource: `
resource "datapower_host_alias" "test" {
  id = "ResTest_HostAlias"
  ip_address = "10.10.10.10"
}`,
    Data: `
data "datapower_host_alias" "test" {
  id         = "AccTest_HostAlias"
}`,
    TestBed: `
resource "datapower_host_alias" "acc_test" {
  id = "AccTest_HostAlias"
  ip_address = "10.10.10.10"
}`,
    ModelOnly:    false,
}
var ImportPackageTestConfig = ModelTestConfig{
    Name:         "ImportPackage",
    Resource: `
resource "datapower_import_package" "test" {
  id = "ResTestImportPackage"
  app_domain = "acceptance_test"
  url = "http://localhost/config.zip"
  import_format = "ZIP"
  overwrite_files = true
  overwrite_objects = true
}`,
    Data: `
data "datapower_import_package" "test" {
  id         = "AccTest_ImportPackage"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_import_package" "acc_test" {
  id = "AccTest_ImportPackage"
  app_domain = datapower_domain.acc_test.app_domain
  url = "http://localhost/config.zip"
  import_format = "ZIP"
  overwrite_files = true
  overwrite_objects = true
}`,
    ModelOnly:    false,
}
var IncludeConfigTestConfig = ModelTestConfig{
    Name:         "IncludeConfig",
    Resource: `
resource "datapower_include_config" "test" {
  id = "ResTestIncludeConfig"
  app_domain = "acceptance_test"
  url = "http://localhost/config.zip"
}`,
    Data: `
data "datapower_include_config" "test" {
  id         = "AccTest_IncludeConfig"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_include_config" "acc_test" {
  id = "AccTest_IncludeConfig"
  app_domain = datapower_domain.acc_test.app_domain
  url = "http://localhost/config.zip"
}`,
    ModelOnly:    false,
}
var InteropServiceTestConfig = ModelTestConfig{
    Name:         "InteropService",
    Resource: `
resource "datapower_interop_service" "test" {
}`,
    Data: `
data "datapower_interop_service" "test" {
}`,
    TestBed: `
resource "datapower_interop_service" "acc_test" {
}`,
    ModelOnly:    false,
}
var JOSERecipientIdentifierTestConfig = ModelTestConfig{
    Name:         "JOSERecipientIdentifier",
    Resource: `
resource "datapower_jose_recipient_identifier" "test" {
  id = "ResTestJOSERecipientIdentifier"
  app_domain = "acceptance_test"
  type = "key"
  key = "AccTest_CryptoKey"
  header_param = ` + DmJOSEHeaderTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_jose_recipient_identifier" "test" {
  id         = "AccTest_JOSERecipientIdentifier"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_jose_recipient_identifier" "acc_test" {
  id = "AccTest_JOSERecipientIdentifier"
  app_domain = datapower_domain.acc_test.app_domain
  type = "key"
  key = datapower_crypto_key.acc_test.id
  header_param = ` + DmJOSEHeaderTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var JOSESignatureIdentifierTestConfig = ModelTestConfig{
    Name:         "JOSESignatureIdentifier",
    Resource: `
resource "datapower_jose_signature_identifier" "test" {
  id = "ResTestJOSESignatureIdentifier"
  app_domain = "acceptance_test"
  type = "certificate"
  certificate = "AccTest_CryptoCertificate"
  header_param = ` + DmJOSEHeaderTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_jose_signature_identifier" "test" {
  id         = "AccTest_JOSESignatureIdentifier"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_jose_signature_identifier" "acc_test" {
  id = "AccTest_JOSESignatureIdentifier"
  app_domain = datapower_domain.acc_test.app_domain
  type = "certificate"
  certificate = datapower_crypto_certificate.acc_test.id
  header_param = ` + DmJOSEHeaderTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var JSONSettingsTestConfig = ModelTestConfig{
    Name:         "JSONSettings",
    Resource: `
resource "datapower_json_settings" "test" {
  id = "ResTestJSONSettings"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_json_settings" "test" {
  id         = "AccTest_JSONSettings"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_json_settings" "acc_test" {
  id = "AccTest_JSONSettings"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var JWEHeaderTestConfig = ModelTestConfig{
    Name:         "JWEHeader",
    Resource: `
resource "datapower_jwe_header" "test" {
  id = "ResTestJWEHeader"
  app_domain = "acceptance_test"
  recipient = "AccTest_JWERecipient"
}`,
    Data: `
data "datapower_jwe_header" "test" {
  id         = "AccTest_JWEHeader"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_jwe_header" "acc_test" {
  id = "AccTest_JWEHeader"
  app_domain = datapower_domain.acc_test.app_domain
  recipient = datapower_jwe_recipient.acc_test.id
}`,
    ModelOnly:    false,
}
var JWERecipientTestConfig = ModelTestConfig{
    Name:         "JWERecipient",
    Resource: `
resource "datapower_jwe_recipient" "test" {
  id = "ResTestJWERecipient"
  app_domain = "acceptance_test"
  algorithm = "RSA1_5"
  certificate = "TestAcc_CryptoCertificate"
}`,
    Data: `
data "datapower_jwe_recipient" "test" {
  id         = "AccTest_JWERecipient"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_jwe_recipient" "acc_test" {
  id = "AccTest_JWERecipient"
  app_domain = datapower_domain.acc_test.app_domain
  algorithm = "RSA1_5"
  certificate = datapower_crypto_certificate.acc_test.id
}`,
    ModelOnly:    false,
}
var JWSSignatureTestConfig = ModelTestConfig{
    Name:         "JWSSignature",
    Resource: `
resource "datapower_jws_signature" "test" {
  id = "ResTestJWSSignature"
  app_domain = "acceptance_test"
  algorithm = "RS256"
  key = "AccTest_CryptoKey"
}`,
    Data: `
data "datapower_jws_signature" "test" {
  id         = "AccTest_JWSSignature"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_jws_signature" "acc_test" {
  id = "AccTest_JWSSignature"
  app_domain = datapower_domain.acc_test.app_domain
  algorithm = "RS256"
  key = datapower_crypto_key.acc_test.id
}`,
    ModelOnly:    false,
}
var KafkaClusterTestConfig = ModelTestConfig{
    Name:         "KafkaCluster",
    Resource: `
resource "datapower_kafka_cluster" "test" {
  id = "ResTestKafkaCluster"
  app_domain = "acceptance_test"
  protocol = "plaintext"
  endpoint = ` + DmKafkaEndpointTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_kafka_cluster" "test" {
  id         = "AccTest_KafkaCluster"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_KafkaCluster"
  app_domain = datapower_domain.acc_test.app_domain
  protocol = "plaintext"
  endpoint = ` + DmKafkaEndpointTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var KafkaSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "KafkaSourceProtocolHandler",
    Resource: `
resource "datapower_kafka_source_protocol_handler" "test" {
  id = "ResTestKafkaSourceProtocolHandler"
  app_domain = "acceptance_test"
  cluster = "AccTest_KafkaCluster"
  request_topic = "topic"
  consumer_group = "consumer"
}`,
    Data: `
data "datapower_kafka_source_protocol_handler" "test" {
  id         = "AccTest_KafkaSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_KafkaSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  cluster = datapower_kafka_cluster.acc_test.id
  request_topic = "topic"
  consumer_group = "consumer"
}`,
    ModelOnly:    false,
}
var LDAPConnectionPoolTestConfig = ModelTestConfig{
    Name:         "LDAPConnectionPool",
    Resource: `
resource "datapower_ldap_connection_pool" "test" {
  id = "ResTestLDAPConnectionPool"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ldap_connection_pool" "test" {
  id         = "AccTest_LDAPConnectionPool"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ldap_connection_pool" "acc_test" {
  id = "AccTest_LDAPConnectionPool"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var LDAPSearchParametersTestConfig = ModelTestConfig{
    Name:         "LDAPSearchParameters",
    Resource: `
resource "datapower_ldap_search_parameters" "test" {
  id = "ResTestLDAPSearchParameters"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ldap_search_parameters" "test" {
  id         = "AccTest_LDAPSearchParameters"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ldap_search_parameters" "acc_test" {
  id = "AccTest_LDAPSearchParameters"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var LoadBalancerGroupTestConfig = ModelTestConfig{
    Name:         "LoadBalancerGroup",
    Resource: `
resource "datapower_load_balancer_group" "test" {
  id = "ResTestLoadBalancerGroup"
  app_domain = "acceptance_test"
  algorithm = "round-robin"
  retrieve_info = false
  damp = 120
}`,
    Data: `
data "datapower_load_balancer_group" "test" {
  id         = "AccTest_LoadBalancerGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_load_balancer_group" "acc_test" {
  id = "AccTest_LoadBalancerGroup"
  app_domain = datapower_domain.acc_test.app_domain
  algorithm = "round-robin"
  retrieve_info = false
  damp = 120
}`,
    ModelOnly:    false,
}
var LogLabelTestConfig = ModelTestConfig{
    Name:         "LogLabel",
    Resource: `
resource "datapower_log_label" "test" {
  id = "ResTestLogLabel"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_log_label" "test" {
  id         = "aaa"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_log_label" "acc_test" {
  id = "AccTest_LogLabel"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var LogTargetTestConfig = ModelTestConfig{
    Name:         "LogTarget",
    Resource: `
resource "datapower_log_target" "test" {
  id = "ResTest_LogTarget"
  app_domain = "acceptance_test"
  type = "file"
  local_file = "logtemp:///AccTest_LogTarget.log"
}`,
    Data: `
data "datapower_log_target" "test" {
  id         = "AccTest_LogTarget"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_log_target" "acc_test" {
  id = "AccTest_LogTarget"
  app_domain = datapower_domain.acc_test.app_domain
  type = "file"
  local_file = "logtemp:///AccTest_LogTarget.log"
}`,
    ModelOnly:    false,
}
var LunaTestConfig = ModelTestConfig{
    Name:         "Luna",
    Resource: `
resource "datapower_luna" "test" {
  id = "ResTestLuna"
  remote_address = "localhost"
  server_cert = "cert:///cert.crt"
  security_option = "none"
}`,
    Data: `
data "datapower_luna" "test" {
  id         = "AccTest_Luna"
}`,
    ModelTestBed: `{
  id = "AccTest_Luna"
  remote_address = "localhost"
  server_cert = "cert:///cert.crt"
  security_option = "none"
}`,
    ModelOnly:    false,
}
var LunaHAGroupTestConfig = ModelTestConfig{
    Name:         "LunaHAGroup",
    Resource: `
resource "datapower_luna_ha_group" "test" {
  id = "ResTestLunaHAGroup"
  app_domain = "acceptance_test"
  group_name = "groupname"
  member = ["AccTest_LunaPartition"]
}`,
    Data: `
data "datapower_luna_ha_group" "test" {
  id         = "AccTest_LunaHAGroup"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_LunaHAGroup"
  app_domain = datapower_domain.acc_test.app_domain
  group_name = "groupname"
  member = [datapower_luna_partition.acc_test.id]
}`,
    ModelOnly:    false,
}
var LunaHASettingsTestConfig = ModelTestConfig{
    Name:         "LunaHASettings",
    Resource: `
resource "datapower_luna_ha_settings" "test" {
}`,
    Data: `
data "datapower_luna_ha_settings" "test" {
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    false,
}
var LunaPartitionTestConfig = ModelTestConfig{
    Name:         "LunaPartition",
    Resource: `
resource "datapower_luna_partition" "test" {
  id = "ResTestLunaPartition"
  app_domain = "acceptance_test"
  partition_name = "partitionname"
  partition_serial = "serial"
  password_alias = "AccTest_PasswordAlias"
  login_role = "co"
}`,
    Data: `
data "datapower_luna_partition" "test" {
  id         = "AccTest_LunaPartition"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_LunaPartition"
  app_domain = datapower_domain.acc_test.app_domain
  partition_name = "partitionname"
  partition_serial = "serial"
  password_alias = datapower_password_alias.acc_test.id
  login_role = "co"
}`,
    ModelOnly:    false,
}
var MCFCustomRuleTestConfig = ModelTestConfig{
    Name:         "MCFCustomRule",
    Resource: `
resource "datapower_mcf_custom_rule" "test" {
  id = "ResTestMCFCustomRule"
  app_domain = "acceptance_test"
  custom_rule_name = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}`,
    Data: `
data "datapower_mcf_custom_rule" "test" {
  id         = "AccTest_MCFCustomRule"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mcf_custom_rule" "acc_test" {
  id = "AccTest_MCFCustomRule"
  app_domain = datapower_domain.acc_test.app_domain
  custom_rule_name = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}`,
    ModelOnly:    false,
}
var MCFHttpHeaderTestConfig = ModelTestConfig{
    Name:         "MCFHttpHeader",
    Resource: `
resource "datapower_mcf_http_header" "test" {
  id = "ResTestMCFHttpHeader"
  app_domain = "acceptance_test"
  http_name = "HEADERNAME"
  http_value = "HEADERVALUE"
}`,
    Data: `
data "datapower_mcf_http_header" "test" {
  id         = "AccTest_MCFHttpHeader"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mcf_http_header" "acc_test" {
  id = "AccTest_MCFHttpHeader"
  app_domain = datapower_domain.acc_test.app_domain
  http_name = "HEADERNAME"
  http_value = "HEADERVALUE"
}`,
    ModelOnly:    false,
}
var MCFHttpMethodTestConfig = ModelTestConfig{
    Name:         "MCFHttpMethod",
    Resource: `
resource "datapower_mcf_http_method" "test" {
  id = "ResTestMCFHttpMethod"
  app_domain = "acceptance_test"
  http_method = "GET"
}`,
    Data: `
data "datapower_mcf_http_method" "test" {
  id         = "AccTest_MCFHttpMethod"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mcf_http_method" "acc_test" {
  id = "AccTest_MCFHttpMethod"
  app_domain = datapower_domain.acc_test.app_domain
  http_method = "GET"
}`,
    ModelOnly:    false,
}
var MCFHttpURLTestConfig = ModelTestConfig{
    Name:         "MCFHttpURL",
    Resource: `
resource "datapower_mcf_http_url" "test" {
  id = "ResTestMCFHttpURL"
  app_domain = "acceptance_test"
  http_url_expression = "*"
}`,
    Data: `
data "datapower_mcf_http_url" "test" {
  id         = "AccTest_MCFHttpURL"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mcf_http_url" "acc_test" {
  id = "AccTest_MCFHttpURL"
  app_domain = datapower_domain.acc_test.app_domain
  http_url_expression = "*"
}`,
    ModelOnly:    false,
}
var MCFXPathTestConfig = ModelTestConfig{
    Name:         "MCFXPath",
    Resource: `
resource "datapower_mcf_xpath" "test" {
  id = "ResTestMCFXPath"
  app_domain = "acceptance_test"
  xpath_expression = "*"
  xpath_value = "value"
}`,
    Data: `
data "datapower_mcf_xpath" "test" {
  id         = "AccTest_MCFXPath"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mcf_xpath" "acc_test" {
  id = "AccTest_MCFXPath"
  app_domain = datapower_domain.acc_test.app_domain
  xpath_expression = "*"
  xpath_value = "value"
}`,
    ModelOnly:    false,
}
var MPGWErrorActionTestConfig = ModelTestConfig{
    Name:         "MPGWErrorAction",
    Resource: `
resource "datapower_mpgw_error_action" "test" {
  id = "ResTestMPGWErrorAction"
  app_domain = "acceptance_test"
  type = "static"
  local_url = "store:///schemas/XMLSchema.dtd"
}`,
    Data: `
data "datapower_mpgw_error_action" "test" {
  id         = "AccTest_MPGWErrorAction"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mpgw_error_action" "acc_test" {
  id = "AccTest_MPGWErrorAction"
  app_domain = datapower_domain.acc_test.app_domain
  type = "static"
  local_url = "store:///schemas/XMLSchema.dtd"
}`,
    ModelOnly:    false,
}
var MPGWErrorHandlingPolicyTestConfig = ModelTestConfig{
    Name:         "MPGWErrorHandlingPolicy",
    Resource: `
resource "datapower_mpgw_error_handling_policy" "test" {
  id = "ResTestMPGWErrorHandlingPolicy"
  app_domain = "acceptance_test"
  policy_maps = ` + DmWebGWErrorPolicyMapTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_mpgw_error_handling_policy" "test" {
  id         = "AccTest_MPGWErrorHandlingPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mpgw_error_handling_policy" "acc_test" {
  id = "AccTest_MPGWErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  policy_maps = ` + DmWebGWErrorPolicyMapTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var MQManagerTestConfig = ModelTestConfig{
    Name:         "MQManager",
    Resource: `
resource "datapower_mq_manager" "test" {
  id = "ResTestMQManager"
  app_domain = "acceptance_test"
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}`,
    Data: `
data "datapower_mq_manager" "test" {
  id         = "AccTest_MQManager"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mq_manager" "acc_test" {
  id = "AccTest_MQManager"
  app_domain = datapower_domain.acc_test.app_domain
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var MQManagerGroupTestConfig = ModelTestConfig{
    Name:         "MQManagerGroup",
    Resource: `
resource "datapower_mq_manager_group" "test" {
  id = "ResTestMQManagerGroup"
  app_domain = "acceptance_test"
  primary_queue_manager = "AccTest_MQManager"
}`,
    Data: `
data "datapower_mq_manager_group" "test" {
  id         = "AccTest_MQManagerGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mq_manager_group" "acc_test" {
  id = "AccTest_MQManagerGroup"
  app_domain = datapower_domain.acc_test.app_domain
  primary_queue_manager = datapower_mq_manager.acc_test.id
}`,
    ModelOnly:    false,
}
var MQv9PlusMFTSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "MQv9PlusMFTSourceProtocolHandler",
    Resource: `
resource "datapower_mqv9_plus_mft_source_protocol_handler" "test" {
  id = "ResTestMQv9PlusMFTSourceProtocolHandler"
  app_domain = "acceptance_test"
  queue_manager = "AccTest_MQManager"
  get_queue = "queue"
}`,
    Data: `
data "datapower_mqv9_plus_mft_source_protocol_handler" "test" {
  id         = "AccTest_MQv9PlusMFTSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mqv9_plus_mft_source_protocol_handler" "acc_test" {
  id = "AccTest_MQv9PlusMFTSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mq_manager.acc_test.id
  get_queue = "queue"
}`,
    ModelOnly:    false,
}
var MQv9PlusSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "MQv9PlusSourceProtocolHandler",
    Resource: `
resource "datapower_mqv9_plus_source_protocol_handler" "test" {
  id = "ResTestMQv9PlusSourceProtocolHandler"
  app_domain = "acceptance_test"
  queue_manager = "AccTest_MQManager"
  get_queue = "queue"
}`,
    Data: `
data "datapower_mqv9_plus_source_protocol_handler" "test" {
  id         = "AccTest_MQv9PlusSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mqv9_plus_source_protocol_handler" "acc_test" {
  id = "AccTest_MQv9PlusSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mq_manager.acc_test.id
  get_queue = "queue"
}`,
    ModelOnly:    false,
}
var MTOMPolicyTestConfig = ModelTestConfig{
    Name:         "MTOMPolicy",
    Resource: `
resource "datapower_mtom_policy" "test" {
  id = "ResTestMTOMPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_mtom_policy" "test" {
  id         = "AccTest_MTOMPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_mtom_policy" "acc_test" {
  id = "AccTest_MTOMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var MatchingTestConfig = ModelTestConfig{
    Name:         "Matching",
    Resource: `
resource "datapower_matching" "test" {
  id = "ResTestMatching"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_matching" "test" {
  id         = "__default-accept-service-providers__"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_matching" "acc_test" {
  id = "AccTest_Matching"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var MessageContentFiltersTestConfig = ModelTestConfig{
    Name:         "MessageContentFilters",
    Resource: `
resource "datapower_message_content_filters" "test" {
  id = "ResTestMessageContentFilters"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_message_content_filters" "test" {
  id         = "AccTest_MessageContentFilters"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_message_content_filters" "acc_test" {
  id = "AccTest_MessageContentFilters"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var MessageMatchingTestConfig = ModelTestConfig{
    Name:         "MessageMatching",
    Resource: `
resource "datapower_message_matching" "test" {
  id = "ResTestMessageMatching"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_message_matching" "test" {
  id         = "AccTest_MessageMatching"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_message_matching" "acc_test" {
  id = "AccTest_MessageMatching"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var MessageTypeTestConfig = ModelTestConfig{
    Name:         "MessageType",
    Resource: `
resource "datapower_message_type" "test" {
  id = "ResTestMessageType"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_message_type" "test" {
  id         = "AccTest_MessageType"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_message_type" "acc_test" {
  id = "AccTest_MessageType"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var MgmtInterfaceTestConfig = ModelTestConfig{
    Name:         "MgmtInterface",
    Resource: `
resource "datapower_mgmt_interface" "test" {
  local_port = 5550
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_mgmt_interface" "test" {
}`,
    TestBed: `
resource "datapower_mgmt_interface" "acc_test" {
  local_port = 5550
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var MultiProtocolGatewayTestConfig = ModelTestConfig{
    Name:         "MultiProtocolGateway",
    Resource: `
resource "datapower_multi_protocol_gateway" "test" {
  id = "ResTestMultiProtocolGateway"
  app_domain = "acceptance_test"
  type = "static-backend"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    Data: `
data "datapower_multi_protocol_gateway" "test" {
  id         = "AccTest_MultiProtocolGateway"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_multi_protocol_gateway" "acc_test" {
  id = "AccTest_MultiProtocolGateway"
  app_domain = datapower_domain.acc_test.app_domain
  type = "static-backend"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    ModelOnly:    false,
}
var NFSClientSettingsTestConfig = ModelTestConfig{
    Name:         "NFSClientSettings",
    Resource: `
resource "datapower_nfs_client_settings" "test" {
}`,
    Data: `
data "datapower_nfs_client_settings" "test" {
}`,
    TestBed: `
resource "datapower_nfs_client_settings" "acc_test" {
}`,
    ModelOnly:    false,
}
var NFSDynamicMountsTestConfig = ModelTestConfig{
    Name:         "NFSDynamicMounts",
    Resource: `
resource "datapower_nfs_dynamic_mounts" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_nfs_dynamic_mounts" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_nfs_dynamic_mounts" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var NFSFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "NFSFilePollerSourceProtocolHandler",
    Resource: `
resource "datapower_nfs_file_poller_source_protocol_handler" "test" {
  id = "ResTestNFSFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
  target_directory = "dpnfs://static-mount-name/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  generate_result_file = false
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    Data: `
data "datapower_nfs_file_poller_source_protocol_handler" "test" {
  id         = "AccTest_NFSFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_nfs_file_poller_source_protocol_handler" "acc_test" {
  id = "AccTest_NFSFilePollerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  target_directory = "dpnfs://static-mount-name/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  generate_result_file = false
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var NFSStaticMountTestConfig = ModelTestConfig{
    Name:         "NFSStaticMount",
    Resource: `
resource "datapower_nfs_static_mount" "test" {
  id = "ResTestNFSStaticMount"
  app_domain = "acceptance_test"
  remote = "url://test"
}`,
    Data: `
data "datapower_nfs_static_mount" "test" {
  id         = "AccTest_NFSStaticMount"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_nfs_static_mount" "acc_test" {
  id = "AccTest_NFSStaticMount"
  app_domain = datapower_domain.acc_test.app_domain
  remote = "url://test"
}`,
    ModelOnly:    false,
}
var NameValueProfileTestConfig = ModelTestConfig{
    Name:         "NameValueProfile",
    Resource: `
resource "datapower_name_value_profile" "test" {
  id = "ResTestNameValueProfile"
  app_domain = "acceptance_test"
  default_fixup = "strip"
}`,
    Data: `
data "datapower_name_value_profile" "test" {
  id         = "AccTest_NameValueProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_name_value_profile" "acc_test" {
  id = "AccTest_NameValueProfile"
  app_domain = datapower_domain.acc_test.app_domain
  default_fixup = "strip"
}`,
    ModelOnly:    false,
}
var OAuthProviderSettingsTestConfig = ModelTestConfig{
    Name:         "OAuthProviderSettings",
    Resource: `
resource "datapower_oauth_provider_settings" "test" {
  id = "ResTestOAuthProviderSettings"
  app_domain = "acceptance_test"
  provider_type = "native"
  api_c_token_secret = "AccTest_CryptoSSKey"
}`,
    Data: `
data "datapower_oauth_provider_settings" "test" {
  id         = "AccTest_OAuthProviderSettings"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_oauth_provider_settings" "acc_test" {
  id = "AccTest_OAuthProviderSettings"
  app_domain = datapower_domain.acc_test.app_domain
  provider_type = "native"
  api_c_token_secret = datapower_crypto_sskey.acc_test.id
}`,
    ModelOnly:    false,
}
var OAuthSupportedClientTestConfig = ModelTestConfig{
    Name:         "OAuthSupportedClient",
    Resource: `
resource "datapower_oauth_supported_client" "test" {
  id = "ResTestOAuthSupportedClient"
  app_domain = "acceptance_test"
  oauth_role = {"azsvr": true}
  az_grant = {"code": true}
  generate_client_secret = false
  client_secret_wo = "secret"
  client_secret_wo_version = 1
  redirect_uri = ["^https://example.com/redirect$"]
  scope = "*"
  token_secret = "AccTest_CryptoSSKey"
}`,
    Data: `
data "datapower_oauth_supported_client" "test" {
  id         = "AccTest_OAuthSupportedClient"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_oauth_supported_client" "acc_test" {
  id = "AccTest_OAuthSupportedClient"
  app_domain = datapower_domain.acc_test.app_domain
  oauth_role = {"azsvr": true}
  az_grant = {"code": true}
  generate_client_secret = false
  client_secret_wo = "secret"
  client_secret_wo_version = 1
  redirect_uri = ["^https://example.com/redirect$"]
  scope = "*"
  token_secret = datapower_crypto_sskey.acc_test.id
}`,
    ModelOnly:    false,
}
var OAuthSupportedClientGroupTestConfig = ModelTestConfig{
    Name:         "OAuthSupportedClientGroup",
    Resource: `
resource "datapower_oauth_supported_client_group" "test" {
  id = "ResTestOAuthSupportedClientGroup"
  app_domain = "acceptance_test"
  client = ["AccTest_OAuthSupportedClient"]
}`,
    Data: `
data "datapower_oauth_supported_client_group" "test" {
  id         = "AccTest_OAuthSupportedClientGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_oauth_supported_client_group" "acc_test" {
  id = "AccTest_OAuthSupportedClientGroup"
  app_domain = datapower_domain.acc_test.app_domain
  client = [datapower_oauth_supported_client.acc_test.id]
}`,
    ModelOnly:    false,
}
var ODRTestConfig = ModelTestConfig{
    Name:         "ODR",
    Resource: `
resource "datapower_odr" "test" {
  odr_server_name = "dp_set"
}`,
    Data: `
data "datapower_odr" "test" {
}`,
    TestBed: `
resource "datapower_odr" "acc_test" {
  odr_server_name = "dp_set"
}`,
    ModelOnly:    false,
}
var ODRConnectorGroupTestConfig = ModelTestConfig{
    Name:         "ODRConnectorGroup",
    Resource: `
resource "datapower_odr_connector_group" "test" {
  id = "ResTestODRConnectorGroup"
  xml_manager = "default"
}`,
    Data: `
data "datapower_odr_connector_group" "test" {
  id         = "AccTest_ODRConnectorGroup"
}`,
    TestBed: `
resource "datapower_odr_connector_group" "acc_test" {
  id = "AccTest_ODRConnectorGroup"
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var OpenTelemetryTestConfig = ModelTestConfig{
    Name:         "OpenTelemetry",
    Resource: `
resource "datapower_open_telemetry" "test" {
  id = "ResTestOpenTelemetry"
  app_domain = "acceptance_test"
  exporter = "AccTest_OpenTelemetryExporter"
  sampler = "AccTest_OpenTelemetrySampler"
}`,
    Data: `
data "datapower_open_telemetry" "test" {
  id         = "AccTest_OpenTelemetry"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_open_telemetry" "acc_test" {
  id = "AccTest_OpenTelemetry"
  app_domain = datapower_domain.acc_test.app_domain
  exporter = datapower_open_telemetry_exporter.acc_test.id
  sampler = datapower_open_telemetry_sampler.acc_test.id
}`,
    ModelOnly:    false,
}
var OpenTelemetryExporterTestConfig = ModelTestConfig{
    Name:         "OpenTelemetryExporter",
    Resource: `
resource "datapower_open_telemetry_exporter" "test" {
  id = "ResTestOpenTelemetryExporter"
  app_domain = "acceptance_test"
  type = "http"
  host_name = "localhost"
  port = 4318
  processor = "batch"
}`,
    Data: `
data "datapower_open_telemetry_exporter" "test" {
  id         = "AccTest_OpenTelemetryExporter"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_open_telemetry_exporter" "acc_test" {
  id = "AccTest_OpenTelemetryExporter"
  app_domain = datapower_domain.acc_test.app_domain
  type = "http"
  host_name = "localhost"
  port = 4318
  processor = "batch"
}`,
    ModelOnly:    false,
}
var OpenTelemetrySamplerTestConfig = ModelTestConfig{
    Name:         "OpenTelemetrySampler",
    Resource: `
resource "datapower_open_telemetry_sampler" "test" {
  id = "ResTestOpenTelemetrySampler"
  app_domain = "acceptance_test"
  parent_based = true
  type = "always-on"
}`,
    Data: `
data "datapower_open_telemetry_sampler" "test" {
  id         = "AccTest_OpenTelemetrySampler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_open_telemetry_sampler" "acc_test" {
  id = "AccTest_OpenTelemetrySampler"
  app_domain = datapower_domain.acc_test.app_domain
  parent_based = true
  type = "always-on"
}`,
    ModelOnly:    false,
}
var OperationRateLimitTestConfig = ModelTestConfig{
    Name:         "OperationRateLimit",
    Resource: `
resource "datapower_operation_rate_limit" "test" {
  id = "ResTestOperationRateLimit"
  app_domain = "acceptance_test"
  operation = "AccTest_APIOperation"
  rate_limit = [{"name": "RateLimit", "rate": "1000"}]
}`,
    Data: `
data "datapower_operation_rate_limit" "test" {
  id         = "AccTest_OperationRateLimit"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_operation_rate_limit" "acc_test" {
  id = "AccTest_OperationRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  operation = datapower_api_operation.acc_test.id
  rate_limit = [{"name": "RateLimit", "rate": "1000"}]
}`,
    ModelOnly:    false,
}
var POPPollerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "POPPollerSourceProtocolHandler",
    Resource: `
resource "datapower_pop_poller_source_protocol_handler" "test" {
  id = "ResTestPOPPollerSourceProtocolHandler"
  app_domain = "acceptance_test"
  mail_server = "localhost"
  port = 8888
  conn_security = "none"
  account = "account"
  password_alias = "AccTest_PasswordAlias"
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
    Data: `
data "datapower_pop_poller_source_protocol_handler" "test" {
  id         = "AccTest_POPPollerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_pop_poller_source_protocol_handler" "acc_test" {
  id = "AccTest_POPPollerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  mail_server = "localhost"
  port = 8888
  conn_security = "none"
  account = "account"
  password_alias = datapower_password_alias.acc_test.id
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
    ModelOnly:    false,
}
var ParseSettingsTestConfig = ModelTestConfig{
    Name:         "ParseSettings",
    Resource: `
resource "datapower_parse_settings" "test" {
  id = "ResTestParseSettings"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_parse_settings" "test" {
  id         = "AccTest_ParseSettings"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_parse_settings" "acc_test" {
  id = "AccTest_ParseSettings"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var PasswordAliasTestConfig = ModelTestConfig{
    Name:         "PasswordAlias",
    Resource: `
resource "datapower_password_alias" "test" {
  id = "ResTestPasswordAlias"
  app_domain = "acceptance_test"
  password_wo = "password"
  password_wo_version = 1
}`,
    Data: `
data "datapower_password_alias" "test" {
  id         = "AccTest_PasswordAlias"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_password_alias" "acc_test" {
  id = "AccTest_PasswordAlias"
  app_domain = datapower_domain.acc_test.app_domain
  password_wo = "password"
  password_wo_version = 1
}`,
    ModelOnly:    false,
}
var PeerGroupTestConfig = ModelTestConfig{
    Name:         "PeerGroup",
    Resource: `
resource "datapower_peer_group" "test" {
  id = "ResTestPeerGroup"
  app_domain = "acceptance_test"
  type = "slm"
  url = ["http://localhost"]
}`,
    Data: `
data "datapower_peer_group" "test" {
  id         = "AccTest_PeerGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_peer_group" "acc_test" {
  id = "AccTest_PeerGroup"
  app_domain = datapower_domain.acc_test.app_domain
  type = "slm"
  url = ["http://localhost"]
}`,
    ModelOnly:    false,
}
var PolicyAttachmentsTestConfig = ModelTestConfig{
    Name:         "PolicyAttachments",
    Resource: `
resource "datapower_policy_attachments" "test" {
  id = "ResTestPolicyAttachments"
  app_domain = "acceptance_test"
  enforcement_mode = "enforce"
  policy_references = false
  sla_enforcement_mode = "allow-if-no-sla"
}`,
    Data: `
data "datapower_policy_attachments" "test" {
  id         = "AccTest_PolicyAttachments"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_policy_attachments" "acc_test" {
  id = "AccTest_PolicyAttachments"
  app_domain = datapower_domain.acc_test.app_domain
  enforcement_mode = "enforce"
  policy_references = false
  sla_enforcement_mode = "allow-if-no-sla"
}`,
    ModelOnly:    false,
}
var PolicyParametersTestConfig = ModelTestConfig{
    Name:         "PolicyParameters",
    Resource: `
resource "datapower_policy_parameters" "test" {
  id = "ResTestPolicyParameters"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_policy_parameters" "test" {
  id         = "AccTest_PolicyParameters"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_policy_parameters" "acc_test" {
  id = "AccTest_PolicyParameters"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var ProbeTestConfig = ModelTestConfig{
    Name:         "Probe",
    Resource: `
resource "datapower_probe" "test" {
  app_domain = "acceptance_test"
  max_records = 1000
  gateway_peering = "default-gateway-peering"
}`,
    Data: `
data "datapower_probe" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_probe" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  max_records = 1000
  gateway_peering = "default-gateway-peering"
}`,
    ModelOnly:    false,
}
var ProcessingMetadataTestConfig = ModelTestConfig{
    Name:         "ProcessingMetadata",
    Resource: `
resource "datapower_processing_metadata" "test" {
  id = "ResTestProcessingMetadata"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_processing_metadata" "test" {
  id         = "AccTest_ProcessingMetadata"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_processing_metadata" "acc_test" {
  id = "AccTest_ProcessingMetadata"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var QuotaEnforcementServerTestConfig = ModelTestConfig{
    Name:         "QuotaEnforcementServer",
    Resource: `
resource "datapower_quota_enforcement_server" "test" {
  server_port = 16379
  monitor_port = 26379
}`,
    Data: `
data "datapower_quota_enforcement_server" "test" {
}`,
    TestBed: `
resource "datapower_quota_enforcement_server" "acc_test" {
  server_port = 16379
  monitor_port = 26379
}`,
    ModelOnly:    false,
}
var RADIUSSettingsTestConfig = ModelTestConfig{
    Name:         "RADIUSSettings",
    Resource: `
resource "datapower_radius_settings" "test" {
}`,
    Data: `
data "datapower_radius_settings" "test" {
}`,
    TestBed: `
resource "datapower_radius_settings" "acc_test" {
}`,
    ModelOnly:    false,
}
var RBMSettingsTestConfig = ModelTestConfig{
    Name:         "RBMSettings",
    Resource: `
resource "datapower_rbm_settings" "test" {
  au_method = "local"
  au_cache_allow = "absolute"
  mc_method = "local"
  min_password_length = 6
  require_mixed_case = false
  require_digit = false
  require_non_alpha_numeric = false
  disallow_username_substring = false
  do_password_aging = false
  do_password_history = false
  cli_timeout = 0
  max_failed_login = 0
}`,
    Data: `
data "datapower_rbm_settings" "test" {
}`,
    TestBed: `
resource "datapower_rbm_settings" "acc_test" {
  au_method = "local"
  au_cache_allow = "absolute"
  mc_method = "local"
  min_password_length = 6
  require_mixed_case = false
  require_digit = false
  require_non_alpha_numeric = false
  disallow_username_substring = false
  do_password_aging = false
  do_password_history = false
  cli_timeout = 0
  max_failed_login = 0
}`,
    ModelOnly:    false,
}
var RaidVolumeTestConfig = ModelTestConfig{
    Name:         "RaidVolume",
    Resource: `
resource "datapower_raid_volume" "test" {
}`,
    Data: `
data "datapower_raid_volume" "test" {
}`,
    TestBed: `
resource "datapower_raid_volume" "acc_test" {
}`,
    ModelOnly:    false,
}
var RateLimitConfigurationTestConfig = ModelTestConfig{
    Name:         "RateLimitConfiguration",
    Resource: `
resource "datapower_rate_limit_configuration" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_rate_limit_configuration" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_rate_limit_configuration" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var RateLimitDefinitionTestConfig = ModelTestConfig{
    Name:         "RateLimitDefinition",
    Resource: `
resource "datapower_rate_limit_definition" "test" {
  id = "ResTestRateLimitDefinition"
  app_domain = "acceptance_test"
  type = "rate"
  rate = 1000
}`,
    Data: `
data "datapower_rate_limit_definition" "test" {
  id         = "AccTest_RateLimitDefinition"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_rate_limit_definition" "acc_test" {
  id = "AccTest_RateLimitDefinition"
  app_domain = datapower_domain.acc_test.app_domain
  type = "rate"
  rate = 1000
}`,
    ModelOnly:    false,
}
var RateLimitDefinitionGroupTestConfig = ModelTestConfig{
    Name:         "RateLimitDefinitionGroup",
    Resource: `
resource "datapower_rate_limit_definition_group" "test" {
  id = "ResTestRateLimitDefinitionGroup"
  app_domain = "acceptance_test"
  update_on_exceed = "all"
}`,
    Data: `
data "datapower_rate_limit_definition_group" "test" {
  id         = "AccTest_RateLimitDefinitionGroup"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_rate_limit_definition_group" "acc_test" {
  id = "AccTest_RateLimitDefinitionGroup"
  app_domain = datapower_domain.acc_test.app_domain
  update_on_exceed = "all"
}`,
    ModelOnly:    false,
}
var SAMLAttributesTestConfig = ModelTestConfig{
    Name:         "SAMLAttributes",
    Resource: `
resource "datapower_saml_attributes" "test" {
  id = "ResTestSAMLAttributes"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_saml_attributes" "test" {
  id         = "AccTest_SAMLAttributes"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_saml_attributes" "acc_test" {
  id = "AccTest_SAMLAttributes"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var SFTPFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "SFTPFilePollerSourceProtocolHandler",
    Resource: `
resource "datapower_sftp_file_poller_source_protocol_handler" "test" {
  id = "ResTestSFTPFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
  ssh_client_connection = "AccTest_SSHClientProfile"
  target_directory = "/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  generate_result_file = false
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    Data: `
data "datapower_sftp_file_poller_source_protocol_handler" "test" {
  id         = "AccTest_SFTPFilePollerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_sftp_file_poller_source_protocol_handler" "acc_test" {
  id = "AccTest_SFTPFilePollerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  ssh_client_connection = datapower_ssh_client_profile.acc_test.id
  target_directory = "/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  generate_result_file = false
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
    ModelOnly:    false,
}
var SLMActionTestConfig = ModelTestConfig{
    Name:         "SLMAction",
    Resource: `
resource "datapower_slm_action" "test" {
  id = "ResTest_SLMAction"
  app_domain = "acceptance_test"
  type = "log-only"
}`,
    Data: `
data "datapower_slm_action" "test" {
  id         = "AccTest_SLMAction"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_slm_action" "acc_test" {
  id = "AccTest_SLMAction"
  app_domain = datapower_domain.acc_test.app_domain
  type = "log-only"
}`,
    ModelOnly:    false,
}
var SLMCredClassTestConfig = ModelTestConfig{
    Name:         "SLMCredClass",
    Resource: `
resource "datapower_slm_cred_class" "test" {
  id = "ResTestSLMCredClass"
  app_domain = "acceptance_test"
  cred_type = "aaa-mapped-credential"
}`,
    Data: `
data "datapower_slm_cred_class" "test" {
  id         = "AccTest_SLMCredClass"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_slm_cred_class" "acc_test" {
  id = "AccTest_SLMCredClass"
  app_domain = datapower_domain.acc_test.app_domain
  cred_type = "aaa-mapped-credential"
}`,
    ModelOnly:    false,
}
var SLMPolicyTestConfig = ModelTestConfig{
    Name:         "SLMPolicy",
    Resource: `
resource "datapower_slm_policy" "test" {
  id = "ResTestSLMPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_slm_policy" "test" {
  id         = "AccTest_SLMPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_slm_policy" "acc_test" {
  id = "AccTest_SLMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var SLMRsrcClassTestConfig = ModelTestConfig{
    Name:         "SLMRsrcClass",
    Resource: `
resource "datapower_slm_rsrc_class" "test" {
  id = "ResTestSLMRsrcClass"
  app_domain = "acceptance_test"
  rsrc_type = "aaa-mapped-resource"
}`,
    Data: `
data "datapower_slm_rsrc_class" "test" {
  id         = "AccTest_SLMRsrcClass"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_slm_rsrc_class" "acc_test" {
  id = "AccTest_SLMRsrcClass"
  app_domain = datapower_domain.acc_test.app_domain
  rsrc_type = "aaa-mapped-resource"
}`,
    ModelOnly:    false,
}
var SLMScheduleTestConfig = ModelTestConfig{
    Name:         "SLMSchedule",
    Resource: `
resource "datapower_slm_schedule" "test" {
  id = "ResTestSLMSchedule"
  app_domain = "acceptance_test"
  start_time = "12:34:00"
  duration = 1440
}`,
    Data: `
data "datapower_slm_schedule" "test" {
  id         = "AccTest_SLMSchedule"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_slm_schedule" "acc_test" {
  id = "AccTest_SLMSchedule"
  app_domain = datapower_domain.acc_test.app_domain
  start_time = "12:34:00"
  duration = 1440
}`,
    ModelOnly:    false,
}
var SMTPServerConnectionTestConfig = ModelTestConfig{
    Name:         "SMTPServerConnection",
    Resource: `
resource "datapower_smtp_server_connection" "test" {
  id = "ResTest_SMTPServerConnection"
  app_domain = "acceptance_test"
  mail_server_host = "localhost"
  mail_server_port = 25
}`,
    Data: `
data "datapower_smtp_server_connection" "test" {
  id         = "AccTest_SMTPServerConnection"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_smtp_server_connection" "acc_test" {
  id = "AccTest_SMTPServerConnection"
  app_domain = datapower_domain.acc_test.app_domain
  mail_server_host = "localhost"
  mail_server_port = 25
}`,
    ModelOnly:    false,
}
var SNMPSettingsTestConfig = ModelTestConfig{
    Name:         "SNMPSettings",
    Resource: `
resource "datapower_snmp_settings" "test" {
  local_port = 161
  security_level = "authPriv"
  access_level = "read-only"
}`,
    Data: `
data "datapower_snmp_settings" "test" {
}`,
    TestBed: `
resource "datapower_snmp_settings" "acc_test" {
  local_port = 161
  security_level = "authPriv"
  access_level = "read-only"
}`,
    ModelOnly:    false,
}
var SOAPHeaderDispositionTestConfig = ModelTestConfig{
    Name:         "SOAPHeaderDisposition",
    Resource: `
resource "datapower_soap_header_disposition" "test" {
  id = "ResTestSOAPHeaderDisposition"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_soap_header_disposition" "test" {
  id         = "AccTest_SOAPHeaderDisposition"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_soap_header_disposition" "acc_test" {
  id = "AccTest_SOAPHeaderDisposition"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var SQLDataSourceTestConfig = ModelTestConfig{
    Name:         "SQLDataSource",
    Resource: `
resource "datapower_sql_data_source" "test" {
  id = "ResTestSQLDataSource"
  app_domain = "acceptance_test"
  database = "Oracle"
  username = "username"
  password_alias = "AccTest_PasswordAlias"
  data_source_id = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection = 10
  connect_timeout = 15
  query_timeout = 30
  idle_timeout = 180
}`,
    Data: `
data "datapower_sql_data_source" "test" {
  id         = "AccTest_SQLDataSource"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_sql_data_source" "acc_test" {
  id = "AccTest_SQLDataSource"
  app_domain = datapower_domain.acc_test.app_domain
  database = "Oracle"
  username = "username"
  password_alias = datapower_password_alias.acc_test.id
  data_source_id = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection = 10
  connect_timeout = 15
  query_timeout = 30
  idle_timeout = 180
}`,
    ModelOnly:    false,
}
var SSHClientProfileTestConfig = ModelTestConfig{
    Name:         "SSHClientProfile",
    Resource: `
resource "datapower_ssh_client_profile" "test" {
  id = "ResTestSSHClientProfile"
  app_domain = "acceptance_test"
  user_name = "someuser"
  profile_usage = "sftp"
}`,
    Data: `
data "datapower_ssh_client_profile" "test" {
  id         = "AccTest_SSHClientProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssh_client_profile" "acc_test" {
  id = "AccTest_SSHClientProfile"
  app_domain = datapower_domain.acc_test.app_domain
  user_name = "someuser"
  profile_usage = "sftp"
}`,
    ModelOnly:    false,
}
var SSHDomainClientProfileTestConfig = ModelTestConfig{
    Name:         "SSHDomainClientProfile",
    Resource: `
resource "datapower_ssh_domain_client_profile" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ssh_domain_client_profile" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssh_domain_client_profile" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var SSHServerProfileTestConfig = ModelTestConfig{
    Name:         "SSHServerProfile",
    Resource: `
resource "datapower_ssh_server_profile" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ssh_server_profile" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssh_server_profile" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var SSHServerSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "SSHServerSourceProtocolHandler",
    Resource: `
resource "datapower_ssh_server_source_protocol_handler" "test" {
  id = "ResTestSSHServerSourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 22
  default_directory = "/"
}`,
    Data: `
data "datapower_ssh_server_source_protocol_handler" "test" {
  id         = "AccTest_SSHServerSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssh_server_source_protocol_handler" "acc_test" {
  id = "AccTest_SSHServerSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 22
  default_directory = "/"
}`,
    ModelOnly:    false,
}
var SSHServiceTestConfig = ModelTestConfig{
    Name:         "SSHService",
    Resource: `
resource "datapower_ssh_service" "test" {
  local_port = 22
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_ssh_service" "test" {
}`,
    TestBed: `
resource "datapower_ssh_service" "acc_test" {
  local_port = 22
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var SSLClientProfileTestConfig = ModelTestConfig{
    Name:         "SSLClientProfile",
    Resource: `
resource "datapower_ssl_client_profile" "test" {
  id = "ResTestSSLClientProfile"
  app_domain = "acceptance_test"
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
  validate_server_cert = false
}`,
    Data: `
data "datapower_ssl_client_profile" "test" {
  id         = "AccTest_SSLClientProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssl_client_profile" "acc_test" {
  id = "AccTest_SSLClientProfile"
  app_domain = datapower_domain.acc_test.app_domain
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
  validate_server_cert = false
}`,
    ModelOnly:    false,
}
var SSLProxyServiceTestConfig = ModelTestConfig{
    Name:         "SSLProxyService",
    Resource: `
resource "datapower_ssl_proxy_service" "test" {
  id = "ResTestSSLProxyService"
  app_domain = "acceptance_test"
  local_port = 4521
  remote_address = "10.10.10.10"
  remote_port = 9999
  ssl_server = "AccTest_SSLServerProfile"
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_ssl_proxy_service" "test" {
  id         = "AccTest_SSLProxyService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssl_proxy_service" "acc_test" {
  id = "AccTest_SSLProxyService"
  app_domain = datapower_domain.acc_test.app_domain
  local_port = 4521
  remote_address = "10.10.10.10"
  remote_port = 9999
  ssl_server = datapower_ssl_server_profile.acc_test.id
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var SSLSNIMappingTestConfig = ModelTestConfig{
    Name:         "SSLSNIMapping",
    Resource: `
resource "datapower_ssl_sni_mapping" "test" {
  id = "ResTestSSLSNIMapping"
  app_domain = "acceptance_test"
  sni_mapping = ` + DmHostToSSLServerProfileTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_ssl_sni_mapping" "test" {
  id         = "AccTest_SSLSNIMapping"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssl_sni_mapping" "acc_test" {
  id = "AccTest_SSLSNIMapping"
  app_domain = datapower_domain.acc_test.app_domain
  sni_mapping = ` + DmHostToSSLServerProfileTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var SSLSNIServerProfileTestConfig = ModelTestConfig{
    Name:         "SSLSNIServerProfile",
    Resource: `
resource "datapower_ssl_sni_server_profile" "test" {
  id = "ResTestSSLSNIServerProfile"
  app_domain = "acceptance_test"
  sni_server_mapping = "AccTest_SSLSNIMapping"
}`,
    Data: `
data "datapower_ssl_sni_server_profile" "test" {
  id         = "AccTest_SSLSNIServerProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssl_sni_server_profile" "acc_test" {
  id = "AccTest_SSLSNIServerProfile"
  app_domain = datapower_domain.acc_test.app_domain
  sni_server_mapping = datapower_ssl_sni_mapping.acc_test.id
}`,
    ModelOnly:    false,
}
var SSLServerProfileTestConfig = ModelTestConfig{
    Name:         "SSLServerProfile",
    Resource: `
resource "datapower_ssl_server_profile" "test" {
  id = "ResTestSSLServerProfile"
  app_domain = "acceptance_test"
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
  idcred = "AccTest_CryptoIdentCred"
}`,
    Data: `
data "datapower_ssl_server_profile" "test" {
  id         = "AccTest_SSLServerProfile"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ssl_server_profile" "acc_test" {
  id = "AccTest_SSLServerProfile"
  app_domain = datapower_domain.acc_test.app_domain
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
  idcred = datapower_crypto_ident_cred.acc_test.id
}`,
    ModelOnly:    false,
}
var SchemaExceptionMapTestConfig = ModelTestConfig{
    Name:         "SchemaExceptionMap",
    Resource: `
resource "datapower_schema_exception_map" "test" {
  id = "ResTestSchemaExceptionMap"
  app_domain = "acceptance_test"
  original_schema_url = "http://localhost"
  schema_exception_rules = ` + DmSchemaExceptionRuleTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_schema_exception_map" "test" {
  id         = "AccTest_SchemaExceptionMap"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_schema_exception_map" "acc_test" {
  id = "AccTest_SchemaExceptionMap"
  app_domain = datapower_domain.acc_test.app_domain
  original_schema_url = "http://localhost"
  schema_exception_rules = ` + DmSchemaExceptionRuleTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var SecureBackupModeTestConfig = ModelTestConfig{
    Name:         "SecureBackupMode",
    Resource: `
resource "datapower_secure_backup_mode" "test" {
}`,
    Data: `
data "datapower_secure_backup_mode" "test" {
}`,
    ModelTestBed: `{
}`,
    ModelOnly:    false,
}
var SocialLoginPolicyTestConfig = ModelTestConfig{
    Name:         "SocialLoginPolicy",
    Resource: `
resource "datapower_social_login_policy" "test" {
  id = "ResTestSocialLoginPolicy"
  app_domain = "acceptance_test"
  client_id = "client_id"
  client_secret = "client_secret"
  client_ssl_profile = "AccTest_SSLClientProfile"
  social_provider = "google"
  provider_az_endpoint = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token = false
}`,
    Data: `
data "datapower_social_login_policy" "test" {
  id         = "AccTest_SocialLoginPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_social_login_policy" "acc_test" {
  id = "AccTest_SocialLoginPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  client_id = "client_id"
  client_secret = "client_secret"
  client_ssl_profile = datapower_ssl_client_profile.acc_test.id
  social_provider = "google"
  provider_az_endpoint = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token = false
}`,
    ModelOnly:    false,
}
var StatelessTCPSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "StatelessTCPSourceProtocolHandler",
    Resource: `
resource "datapower_stateless_tcp_source_protocol_handler" "test" {
  id = "ResTestStatelessTCPSourceProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 4000
}`,
    Data: `
data "datapower_stateless_tcp_source_protocol_handler" "test" {
  id         = "AccTest_StatelessTCPSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_stateless_tcp_source_protocol_handler" "acc_test" {
  id = "AccTest_StatelessTCPSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 4000
}`,
    ModelOnly:    false,
}
var StatisticsTestConfig = ModelTestConfig{
    Name:         "Statistics",
    Resource: `
resource "datapower_statistics" "test" {
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_statistics" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_statistics" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var StylePolicyTestConfig = ModelTestConfig{
    Name:         "StylePolicy",
    Resource: `
resource "datapower_style_policy" "test" {
  id = "ResTestStylePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_style_policy" "test" {
  id         = "AccTest_StylePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_style_policy" "acc_test" {
  id = "AccTest_StylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var StylePolicyActionTestConfig = ModelTestConfig{
    Name:         "StylePolicyAction",
    Resource: `
resource "datapower_style_policy_action" "test" {
  id = "ResTestStylePolicyAction"
  app_domain = "acceptance_test"
  type = "xform"
  input = "INPUT"
  transform = "tfile"
  output = "OUTPUT"
  named_inputs = null
  named_outputs = null
}`,
    Data: `
data "datapower_style_policy_action" "test" {
  id         = "__default-accept-service-providers-filter-action__"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_style_policy_action" "acc_test" {
  id = "AccTest_StylePolicyAction"
  app_domain = datapower_domain.acc_test.app_domain
  type = "xform"
  input = "INPUT"
  transform = "tfile"
  output = "OUTPUT"
  named_inputs = null
  named_outputs = null
}`,
    ModelOnly:    false,
}
var StylePolicyRuleTestConfig = ModelTestConfig{
    Name:         "StylePolicyRule",
    Resource: `
resource "datapower_style_policy_rule" "test" {
  id = "ResTestStylePolicyRule"
  app_domain = "acceptance_test"
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
    Data: `
data "datapower_style_policy_rule" "test" {
  id         = "__default-accept-service-providers-request__"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_style_policy_rule" "acc_test" {
  id = "AccTest_StylePolicyRule"
  app_domain = datapower_domain.acc_test.app_domain
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
    ModelOnly:    false,
}
var SystemSettingsTestConfig = ModelTestConfig{
    Name:         "SystemSettings",
    Resource: `
resource "datapower_system_settings" "test" {
}`,
    Data: `
data "datapower_system_settings" "test" {
}`,
    TestBed: `
resource "datapower_system_settings" "acc_test" {
}`,
    ModelOnly:    false,
}
var TAMTestConfig = ModelTestConfig{
    Name:         "TAM",
    Resource: `
resource "datapower_tam" "test" {
  id = "ResTestTAM"
  app_domain = "acceptance_test"
  configuration_file = "local:///tam.config"
  ssl_key_file = "cert:///ssl.key"
  ssl_key_stash_file = "cert:///ssl.stash"
  use_local_mode = false
  tam_use_basic_user = false
}`,
    Data: `
data "datapower_tam" "test" {
  id         = "AccTest_TAM"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_TAM"
  app_domain = datapower_domain.acc_test.app_domain
  configuration_file = "local:///tam.config"
  ssl_key_file = "cert:///ssl.key"
  ssl_key_stash_file = "cert:///ssl.stash"
  use_local_mode = false
  tam_use_basic_user = false
}`,
    ModelOnly:    false,
}
var TCPProxyServiceTestConfig = ModelTestConfig{
    Name:         "TCPProxyService",
    Resource: `
resource "datapower_tcp_proxy_service" "test" {
  id = "ResTestTCPProxyService"
  app_domain = "acceptance_test"
  local_port = 6158
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_tcp_proxy_service" "test" {
  id         = "AccTest_TCPProxyService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_tcp_proxy_service" "acc_test" {
  id = "AccTest_TCPProxyService"
  app_domain = datapower_domain.acc_test.app_domain
  local_port = 6158
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var ThrottlerTestConfig = ModelTestConfig{
    Name:         "Throttler",
    Resource: `
resource "datapower_throttler" "test" {
  throttle_at = 0
  terminate_at = 0
  temp_fs_throttle_at = 0
  temp_fs_terminate_at = 0
  qname_warn_at = 10
  timeout = 30
}`,
    Data: `
data "datapower_throttler" "test" {
}`,
    TestBed: `
resource "datapower_throttler" "acc_test" {
  throttle_at = 0
  terminate_at = 0
  temp_fs_throttle_at = 0
  temp_fs_terminate_at = 0
  qname_warn_at = 10
  timeout = 30
}`,
    ModelOnly:    false,
}
var TimeSettingsTestConfig = ModelTestConfig{
    Name:         "TimeSettings",
    Resource: `
resource "datapower_time_settings" "test" {
  local_time_zone = "EST5EDT"
}`,
    Data: `
data "datapower_time_settings" "test" {
}`,
    TestBed: `
resource "datapower_time_settings" "acc_test" {
  local_time_zone = "EST5EDT"
}`,
    ModelOnly:    false,
}
var URLMapTestConfig = ModelTestConfig{
    Name:         "URLMap",
    Resource: `
resource "datapower_url_map" "test" {
  id = "ResTestURLMap"
  app_domain = "acceptance_test"
  url_map_rule = ` + DmURLMapRuleTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_url_map" "test" {
  id         = "AccTest_URLMap"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_url_map" "acc_test" {
  id = "AccTest_URLMap"
  app_domain = datapower_domain.acc_test.app_domain
  url_map_rule = ` + DmURLMapRuleTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var URLRefreshPolicyTestConfig = ModelTestConfig{
    Name:         "URLRefreshPolicy",
    Resource: `
resource "datapower_url_refresh_policy" "test" {
  id = "ResTestURLRefreshPolicy"
  app_domain = "acceptance_test"
  url_refresh_rule = ` + DmURLRefreshRuleTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_url_refresh_policy" "test" {
  id         = "AccTest_URLRefreshPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_url_refresh_policy" "acc_test" {
  id = "AccTest_URLRefreshPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  url_refresh_rule = ` + DmURLRefreshRuleTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var URLRewritePolicyTestConfig = ModelTestConfig{
    Name:         "URLRewritePolicy",
    Resource: `
resource "datapower_url_rewrite_policy" "test" {
  id = "ResTestURLRewritePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_url_rewrite_policy" "test" {
  id         = "AccTest_URLRewritePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_url_rewrite_policy" "acc_test" {
  id = "AccTest_URLRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var UserTestConfig = ModelTestConfig{
    Name:         "User",
    Resource: `
resource "datapower_user" "test" {
  id = "ResTest_User"
  password_wo = "Password$123"
  password_wo_version = 1
  access_level = "group-defined"
  group_name = "AccTest_UserGroup"
  snmp_creds = null
  hashed_snmp_creds = null
}`,
    Data: `
data "datapower_user" "test" {
  id         = "AccTest_User"
}`,
    TestBed: `
resource "datapower_user" "acc_test" {
  id = "AccTest_User"
  password_wo = "Password$123"
  password_wo_version = 1
  access_level = "group-defined"
  group_name = datapower_user_group.acc_test.id
  snmp_creds = null
  hashed_snmp_creds = null
}`,
    ModelOnly:    false,
}
var UserGroupTestConfig = ModelTestConfig{
    Name:         "UserGroup",
    Resource: `
resource "datapower_user_group" "test" {
  id = "ResTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}`,
    Data: `
data "datapower_user_group" "test" {
  id         = "AccTest_UserGroup"
}`,
    TestBed: `
resource "datapower_user_group" "acc_test" {
  id = "AccTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}`,
    ModelOnly:    false,
}
var VisibilityListTestConfig = ModelTestConfig{
    Name:         "VisibilityList",
    Resource: `
resource "datapower_visibility_list" "test" {
  id = "ResTestVisibilityList"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_visibility_list" "test" {
  id         = "AccTest_VisibilityList"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_visibility_list" "acc_test" {
  id = "AccTest_VisibilityList"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var WCCServiceTestConfig = ModelTestConfig{
    Name:         "WCCService",
    Resource: `
resource "datapower_wcc_service" "test" {
  id = "ResTestWCCService"
  app_domain = "acceptance_test"
  odc_info_hostname = "example.com"
  odc_info_port = 1024
  time_interval = 10
}`,
    Data: `
data "datapower_wcc_service" "test" {
  id         = "AccTest_WCCService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_wcc_service" "acc_test" {
  id = "AccTest_WCCService"
  app_domain = datapower_domain.acc_test.app_domain
  odc_info_hostname = "example.com"
  odc_info_port = 1024
  time_interval = 10
}`,
    ModelOnly:    false,
}
var WSEndpointRewritePolicyTestConfig = ModelTestConfig{
    Name:         "WSEndpointRewritePolicy",
    Resource: `
resource "datapower_ws_endpoint_rewrite_policy" "test" {
  id = "ResTestWSEndpointRewritePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ws_endpoint_rewrite_policy" "test" {
  id         = "AccTest_WSEndpointRewritePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ws_endpoint_rewrite_policy" "acc_test" {
  id = "AccTest_WSEndpointRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var WSGatewayTestConfig = ModelTestConfig{
    Name:         "WSGateway",
    Resource: `
resource "datapower_ws_gateway" "test" {
  id = "ResTestWSGateway"
  app_domain = "acceptance_test"
  type = "static-from-wsdl"
  style_policy = "default"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    Data: `
data "datapower_ws_gateway" "test" {
  id         = "AccTest_WSGateway"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ws_gateway" "acc_test" {
  id = "AccTest_WSGateway"
  app_domain = datapower_domain.acc_test.app_domain
  type = "static-from-wsdl"
  style_policy = "default"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    ModelOnly:    false,
}
var WSRRSavedSearchSubscriptionTestConfig = ModelTestConfig{
    Name:         "WSRRSavedSearchSubscription",
    Resource: `
resource "datapower_wsrr_saved_search_subscription" "test" {
  id = "ResTestWSRRSavedSearchSubscription"
  app_domain = "acceptance_test"
  server = "AccTest_WSRRServer"
  saved_search_name = "ResTestsaved_search"
}`,
    Data: `
data "datapower_wsrr_saved_search_subscription" "test" {
  id         = "AccTest_WSRRSavedSearchSubscription"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_wsrr_saved_search_subscription" "acc_test" {
  id = "AccTest_WSRRSavedSearchSubscription"
  app_domain = datapower_domain.acc_test.app_domain
  server = datapower_wsrr_server.acc_test.id
  saved_search_name = "ResTestsaved_search"
}`,
    ModelOnly:    false,
}
var WSRRServerTestConfig = ModelTestConfig{
    Name:         "WSRRServer",
    Resource: `
resource "datapower_wsrr_server" "test" {
  id = "ResTest_WSRRServer"
  app_domain = "acceptance_test"
  soap_url = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}`,
    Data: `
data "datapower_wsrr_server" "test" {
  id         = "AccTest_WSRRServer"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_wsrr_server" "acc_test" {
  id = "AccTest_WSRRServer"
  app_domain = datapower_domain.acc_test.app_domain
  soap_url = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}`,
    ModelOnly:    false,
}
var WSRRSubscriptionTestConfig = ModelTestConfig{
    Name:         "WSRRSubscription",
    Resource: `
resource "datapower_wsrr_subscription" "test" {
  id = "ResTestWSRRSubscription"
  app_domain = "acceptance_test"
  server = "AccTest_WSRRServer"
  namespace = "namespace"
  object_type = "wsdl"
  object_name = "ResTestobject"
}`,
    Data: `
data "datapower_wsrr_subscription" "test" {
  id         = "AccTest_WSRRSubscription"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_wsrr_subscription" "acc_test" {
  id = "AccTest_WSRRSubscription"
  app_domain = datapower_domain.acc_test.app_domain
  server = datapower_wsrr_server.acc_test.id
  namespace = "namespace"
  object_type = "wsdl"
  object_name = "ResTestobject"
}`,
    ModelOnly:    false,
}
var WSStylePolicyTestConfig = ModelTestConfig{
    Name:         "WSStylePolicy",
    Resource: `
resource "datapower_ws_style_policy" "test" {
  id = "ResTestWSStylePolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_ws_style_policy" "test" {
  id         = "AccTest_WSStylePolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ws_style_policy" "acc_test" {
  id = "AccTest_WSStylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var WSStylePolicyRuleTestConfig = ModelTestConfig{
    Name:         "WSStylePolicyRule",
    Resource: `
resource "datapower_ws_style_policy_rule" "test" {
  id = "ResTest_WSStylePolicyRule"
  app_domain = "acceptance_test"
  actions = ["AccTest_StylePolicyAction"]
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
    Data: `
data "datapower_ws_style_policy_rule" "test" {
  id         = "AccTest_WSStylePolicyRule"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_ws_style_policy_rule" "acc_test" {
  id = "AccTest_WSStylePolicyRule"
  app_domain = datapower_domain.acc_test.app_domain
  actions = [datapower_style_policy_action.acc_test.id]
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
    ModelOnly:    false,
}
var WXSGridTestConfig = ModelTestConfig{
    Name:         "WXSGrid",
    Resource: `
resource "datapower_wxs_grid" "test" {
  id = "ResTestWXSGrid"
  app_domain = "acceptance_test"
  collective = "AccTest_LoadBalancerGroup"
  grid = "gridname"
  user_name = "username"
  password_alias = "AccTest_PasswordAlias"
  timeout = 1000
}`,
    Data: `
data "datapower_wxs_grid" "test" {
  id         = "AccTest_WXSGrid"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_wxs_grid" "acc_test" {
  id = "AccTest_WXSGrid"
  app_domain = datapower_domain.acc_test.app_domain
  collective = datapower_load_balancer_group.acc_test.id
  grid = "gridname"
  user_name = "username"
  password_alias = datapower_password_alias.acc_test.id
  timeout = 1000
}`,
    ModelOnly:    false,
}
var WebAppErrorHandlingPolicyTestConfig = ModelTestConfig{
    Name:         "WebAppErrorHandlingPolicy",
    Resource: `
resource "datapower_web_app_error_handling_policy" "test" {
  id = "ResTestWebAppErrorHandlingPolicy"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_web_app_error_handling_policy" "test" {
  id         = "AccTest_WebAppErrorHandlingPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_app_error_handling_policy" "acc_test" {
  id = "AccTest_WebAppErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var WebAppFWTestConfig = ModelTestConfig{
    Name:         "WebAppFW",
    Resource: `
resource "datapower_web_app_fw" "test" {
  id = "ResTestWebAppFW"
  app_domain = "acceptance_test"
  front_side = [{"LocalAddress": "0.0.0.0"}]
  remote_address = "10.10.10.10"
  style_policy = "AccTest_AppSecurityPolicy"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    Data: `
data "datapower_web_app_fw" "test" {
  id         = "AccTest_WebAppFW"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_app_fw" "acc_test" {
  id = "AccTest_WebAppFW"
  app_domain = datapower_domain.acc_test.app_domain
  front_side = [{"LocalAddress": "0.0.0.0"}]
  remote_address = "10.10.10.10"
  style_policy = datapower_app_security_policy.acc_test.id
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
    ModelOnly:    false,
}
var WebAppRequestTestConfig = ModelTestConfig{
    Name:         "WebAppRequest",
    Resource: `
resource "datapower_web_app_request" "test" {
  id = "ResTest_WebAppRequest"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_web_app_request" "test" {
  id         = "AccTest_WebAppRequest"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_app_request" "acc_test" {
  id = "AccTest_WebAppRequest"
  app_domain = datapower_domain.acc_test.app_domain
}`,
    ModelOnly:    false,
}
var WebAppResponseTestConfig = ModelTestConfig{
    Name:         "WebAppResponse",
    Resource: `
resource "datapower_web_app_response" "test" {
  id = "ResTest_WebAppResponse"
  app_domain = "acceptance_test"
  policy_type = "admission"
}`,
    Data: `
data "datapower_web_app_response" "test" {
  id         = "AccTest_WebAppResponse"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_app_response" "acc_test" {
  id = "AccTest_WebAppResponse"
  app_domain = datapower_domain.acc_test.app_domain
  policy_type = "admission"
}`,
    ModelOnly:    false,
}
var WebAppSessionPolicyTestConfig = ModelTestConfig{
    Name:         "WebAppSessionPolicy",
    Resource: `
resource "datapower_web_app_session_policy" "test" {
  id = "ResTestWebAppSessionPolicy"
  app_domain = "acceptance_test"
  start_matches = "__default-accept-service-providers__"
}`,
    Data: `
data "datapower_web_app_session_policy" "test" {
  id         = "AccTest_WebAppSessionPolicy"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_app_session_policy" "acc_test" {
  id = "AccTest_WebAppSessionPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  start_matches = "__default-accept-service-providers__"
}`,
    ModelOnly:    false,
}
var WebB2BViewerTestConfig = ModelTestConfig{
    Name:         "WebB2BViewer",
    Resource: `
resource "datapower_web_b2b_viewer" "test" {
  local_port = 9091
  idle_timeout = 600
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_web_b2b_viewer" "test" {
}`,
    TestBed: `
resource "datapower_web_b2b_viewer" "acc_test" {
  local_port = 9091
  idle_timeout = 600
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var WebGUITestConfig = ModelTestConfig{
    Name:         "WebGUI",
    Resource: `
resource "datapower_web_gui" "test" {
  local_port = 9090
  save_config_overwrites = true
  idle_timeout = 0
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_web_gui" "test" {
}`,
    TestBed: `
resource "datapower_web_gui" "acc_test" {
  local_port = 9090
  save_config_overwrites = true
  idle_timeout = 0
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var WebServiceMonitorTestConfig = ModelTestConfig{
    Name:         "WebServiceMonitor",
    Resource: `
resource "datapower_web_service_monitor" "test" {
  id = "ResTestWebServiceMonitor"
  app_domain = "acceptance_test"
  wsdl_url = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}`,
    Data: `
data "datapower_web_service_monitor" "test" {
  id         = "AccTest_WebServiceMonitor"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_service_monitor" "acc_test" {
  id = "AccTest_WebServiceMonitor"
  app_domain = datapower_domain.acc_test.app_domain
  wsdl_url = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}`,
    ModelOnly:    false,
}
var WebServicesAgentTestConfig = ModelTestConfig{
    Name:         "WebServicesAgent",
    Resource: `
resource "datapower_web_services_agent" "test" {
  app_domain = "acceptance_test"
  max_records = 3000
  max_memory_kb = 64000
  capture_mode = "faults"
}`,
    Data: `
data "datapower_web_services_agent" "test" {
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_services_agent" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
  max_records = 3000
  max_memory_kb = 64000
  capture_mode = "faults"
}`,
    ModelOnly:    false,
}
var WebSphereJMSServerTestConfig = ModelTestConfig{
    Name:         "WebSphereJMSServer",
    Resource: `
resource "datapower_web_sphere_jms_server" "test" {
  id = "ResTest_WebSphereJMSServer"
  app_domain = "acceptance_test"
  endpoint = ` + DmWebSphereJMSEndpointTestConfig.GetModelListConfig() + `
  messaging_bus = "bus"
}`,
    Data: `
data "datapower_web_sphere_jms_server" "test" {
  id         = "AccTest_WebSphereJMSServer"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_WebSphereJMSServer"
  app_domain = datapower_domain.acc_test.app_domain
  endpoint = ` + DmWebSphereJMSEndpointTestConfig.GetModelTestBedListConfig() + `
  messaging_bus = "bus"
}`,
    ModelOnly:    false,
}
var WebSphereJMSSourceProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "WebSphereJMSSourceProtocolHandler",
    Resource: `
resource "datapower_web_sphere_jms_source_protocol_handler" "test" {
  id = "ResTestWebSphereJMSSourceProtocolHandler"
  app_domain = "acceptance_test"
  server = "AccTest_WebSphereJMSServer"
  get_queue = "getqueue"
}`,
    Data: `
data "datapower_web_sphere_jms_source_protocol_handler" "test" {
  id         = "AccTest_WebSphereJMSSourceProtocolHandler"
  app_domain = "acceptance_test"
}`,
    ModelTestBed: `{
  id = "AccTest_WebSphereJMSSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  server = datapower_web_sphere_jms_server.acc_test.id
  get_queue = "getqueue"
}`,
    ModelOnly:    false,
}
var WebTokenServiceTestConfig = ModelTestConfig{
    Name:         "WebTokenService",
    Resource: `
resource "datapower_web_token_service" "test" {
  id = "ResTestWebTokenService"
  app_domain = "acceptance_test"
  xml_manager = "default"
  front_side = ` + DmSSLFrontSideTestConfig.GetModelListConfig() + `
  style_policy = "default"
  front_timeout = 120
  front_persistent_timeout = 180
}`,
    Data: `
data "datapower_web_token_service" "test" {
  id         = "AccTest_WebTokenService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_web_token_service" "acc_test" {
  id = "AccTest_WebTokenService"
  app_domain = datapower_domain.acc_test.app_domain
  xml_manager = "default"
  front_side = ` + DmSSLFrontSideTestConfig.GetModelTestBedListConfig() + `
  style_policy = "default"
  front_timeout = 120
  front_persistent_timeout = 180
}`,
    ModelOnly:    false,
}
var XACMLPDPTestConfig = ModelTestConfig{
    Name:         "XACMLPDP",
    Resource: `
resource "datapower_xacml_pdp" "test" {
  id = "ResTestXACMLPDP"
  app_domain = "acceptance_test"
  general_policy = "http://test/uri"
}`,
    Data: `
data "datapower_xacml_pdp" "test" {
  id         = "AccTest_XACMLPDP"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xacml_pdp" "acc_test" {
  id = "AccTest_XACMLPDP"
  app_domain = datapower_domain.acc_test.app_domain
  general_policy = "http://test/uri"
}`,
    ModelOnly:    false,
}
var XMLFirewallServiceTestConfig = ModelTestConfig{
    Name:         "XMLFirewallService",
    Resource: `
resource "datapower_xml_firewall_service" "test" {
  id = "ResTestXMLFirewallService"
  app_domain = "acceptance_test"
  type = "dynamic-backend"
  xml_manager = "default"
  local_port = 7575
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_xml_firewall_service" "test" {
  id         = "AccTest_XMLFirewallService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xml_firewall_service" "acc_test" {
  id = "AccTest_XMLFirewallService"
  app_domain = datapower_domain.acc_test.app_domain
  type = "dynamic-backend"
  xml_manager = "default"
  local_port = 7575
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var XMLManagerTestConfig = ModelTestConfig{
    Name:         "XMLManager",
    Resource: `
resource "datapower_xml_manager" "test" {
  id = "ResTestXMLManager"
  app_domain = "acceptance_test"
}`,
    Data: `
data "datapower_xml_manager" "test" {
  id         = "AccTest_XMLManager"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xml_manager" "acc_test" {
  id = "AccTest_XMLManager"
  app_domain = datapower_domain.acc_test.app_domain
  ldap_conn_pool = datapower_ldap_connection_pool.acc_test.id
}`,
    ModelOnly:    false,
}
var XPathRoutingMapTestConfig = ModelTestConfig{
    Name:         "XPathRoutingMap",
    Resource: `
resource "datapower_xpath_routing_map" "test" {
  id = "ResTestXPathRoutingMap"
  app_domain = "acceptance_test"
  xpath_routing_rules = ` + DmXPathRoutingRuleTestConfig.GetModelListConfig() + `
}`,
    Data: `
data "datapower_xpath_routing_map" "test" {
  id         = "AccTest_XPathRoutingMap"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xpath_routing_map" "acc_test" {
  id = "AccTest_XPathRoutingMap"
  app_domain = datapower_domain.acc_test.app_domain
  xpath_routing_rules = ` + DmXPathRoutingRuleTestConfig.GetModelTestBedListConfig() + `
}`,
    ModelOnly:    false,
}
var XSLCoprocServiceTestConfig = ModelTestConfig{
    Name:         "XSLCoprocService",
    Resource: `
resource "datapower_xsl_coproc_service" "test" {
  id = "ResTestXSLCoprocService"
  app_domain = "acceptance_test"
  local_port = 8811
  xml_manager = "default"
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_xsl_coproc_service" "test" {
  id         = "AccTest_XSLCoprocService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xsl_coproc_service" "acc_test" {
  id = "AccTest_XSLCoprocService"
  app_domain = datapower_domain.acc_test.app_domain
  local_port = 8811
  xml_manager = "default"
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var XSLProxyServiceTestConfig = ModelTestConfig{
    Name:         "XSLProxyService",
    Resource: `
resource "datapower_xsl_proxy_service" "test" {
  id = "ResTestXSLProxyService"
  app_domain = "acceptance_test"
  type = "static-backend"
  xml_manager = "default"
  local_port = 8922
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
    Data: `
data "datapower_xsl_proxy_service" "test" {
  id         = "AccTest_XSLProxyService"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xsl_proxy_service" "acc_test" {
  id = "AccTest_XSLProxyService"
  app_domain = datapower_domain.acc_test.app_domain
  type = "static-backend"
  xml_manager = "default"
  local_port = 8922
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
    ModelOnly:    false,
}
var XTCProtocolHandlerTestConfig = ModelTestConfig{
    Name:         "XTCProtocolHandler",
    Resource: `
resource "datapower_xtc_protocol_handler" "test" {
  id = "ResTestXTCProtocolHandler"
  app_domain = "acceptance_test"
  local_address = "0.0.0.0"
  local_port = 3000
  remote_address = "10.10.10.10"
  remote_port = 12000
}`,
    Data: `
data "datapower_xtc_protocol_handler" "test" {
  id         = "AccTest_XTCProtocolHandler"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_xtc_protocol_handler" "acc_test" {
  id = "AccTest_XTCProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port = 3000
  remote_address = "10.10.10.10"
  remote_port = 12000
}`,
    ModelOnly:    false,
}
var ZosNSSClientTestConfig = ModelTestConfig{
    Name:         "ZosNSSClient",
    Resource: `
resource "datapower_zos_nss_client" "test" {
  id = "ResTestZosNSSClient"
  app_domain = "acceptance_test"
  remote_address = "remote.host"
  remote_port = 443
  client_id = "client_id"
  system_name = "sysname"
  user_name = "username"
  password_alias = "AccTest_PasswordAlias"
  ssl_client = "AccTest_SSLClientProfile"
}`,
    Data: `
data "datapower_zos_nss_client" "test" {
  id         = "AccTest_ZosNSSClient"
  app_domain = "acceptance_test"
}`,
    TestBed: `
resource "datapower_zos_nss_client" "acc_test" {
  id = "AccTest_ZosNSSClient"
  app_domain = datapower_domain.acc_test.app_domain
  remote_address = "remote.host"
  remote_port = 443
  client_id = "client_id"
  system_name = "sysname"
  user_name = "username"
  password_alias = datapower_password_alias.acc_test.id
  ssl_client = datapower_ssl_client_profile.acc_test.id
}`,
    ModelOnly:    false,
}
