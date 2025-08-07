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

package testconfig

type ModelTestConfig struct {
	Name         string
	Model        string
	Resource     string
	Data         string
	ModelOnly    bool
	Dependencies map[string]*ModelTestConfig
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

func (c ModelTestConfig) GetDataConfig() string {
	return c.GetResourceConfig() + c.Data
}

func (c ModelTestConfig) GetResourceConfig() string {
	return c.getPrequisites() + c.Resource
}

func (c *ModelTestConfig) getPrequisites() string {
	referencesTo := make(map[string]*ModelTestConfig)
	preReqs := c.TestPre
	for _, v := range referencesTo {
		if !v.ModelOnly {
			preReqs = preReqs + v.Resource
		}
	}
	return preReqs
}

var FileTestConfig = ModelTestConfig{
	Name: "File",
	Resource: `
resource "datapower_file" "test" {
  app_domain = "acc_test_domain"
  remote_path = "cert:///test_file.txt"
  local_path = "/workspaces/terraform-provider-datapower/testutils/test_file.txt"
  dependency_actions = [
    {
      target_type = "resource_datapower_domain",
      target_domain = "acc_test_domain",
      target_id = "ignored",
      action = "quiesce"
      on_create = true
    }
  ]
}
`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
}
var AAAJWTGeneratorTestConfig = ModelTestConfig{
	Name: "AAAJWTGenerator",
	Resource: `
resource "datapower_aaajwtgenerator" "test" {
  id = "AAAJWTGenerator_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_aaajwtgenerator" "test" {
  depends_on = [ datapower_aaajwtgenerator.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AAAJWTValidatorTestConfig = ModelTestConfig{
	Name: "AAAJWTValidator",
	Resource: `
resource "datapower_aaajwtvalidator" "test" {
  id = "AAAJWTValidator_name"
  app_domain = "acc_test_domain"
  username_claim = "sub"
}`,
	Data: `
data "datapower_aaajwtvalidator" "test" {
  depends_on = [ datapower_aaajwtvalidator.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AAAPolicyTestConfig = ModelTestConfig{
	Name: "AAAPolicy",
	Resource: `
resource "datapower_aaapolicy" "test" {
  id = "AAAPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_aaapolicy" "test" {
  depends_on = [ datapower_aaapolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AMQPBrokerTestConfig = ModelTestConfig{
	Name: "AMQPBroker",
	Resource: `
resource "datapower_amqpbroker" "test" {
  id = "AMQPBroker_name"
  app_domain = "acc_test_domain"
  host_name = "host.name"
  port = 5672
  xml_manager = "default"
  authorization = "none"
}`,
	Data: `
data "datapower_amqpbroker" "test" {
  depends_on = [ datapower_amqpbroker.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var AMQPSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "AMQPSourceProtocolHandler",
	Resource: `
resource "datapower_amqpsourceprotocolhandler" "test" {
  id = "AMQPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  broker = "TestAccAMQPBroker"
  from = "amqpfrom"
}`,
	Data: `
data "datapower_amqpsourceprotocolhandler" "test" {
  depends_on = [ datapower_amqpsourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"AMQPBroker": &AMQPBrokerTestConfig,
	},
	TestPre: `
`,
}
var APIApplicationTypeTestConfig = ModelTestConfig{
	Name: "APIApplicationType",
	Resource: `
resource "datapower_apiapplicationtype" "test" {
  id = "APIApplicationType_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apiapplicationtype" "test" {
  depends_on = [ datapower_apiapplicationtype.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIAuthURLRegistryTestConfig = ModelTestConfig{
	Name: "APIAuthURLRegistry",
	Resource: `
resource "datapower_apiauthurlregistry" "test" {
  id = "APIAuthURLRegistry_name"
  app_domain = "acc_test_domain"
  auth_url = "http://localhost"
}`,
	Data: `
data "datapower_apiauthurlregistry" "test" {
  depends_on = [ datapower_apiauthurlregistry.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APICORSTestConfig = ModelTestConfig{
	Name: "APICORS",
	Resource: `
resource "datapower_apicors" "test" {
  id = "APICORS_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apicors" "test" {
  depends_on = [ datapower_apicors.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIClientIdentificationTestConfig = ModelTestConfig{
	Name: "APIClientIdentification",
	Resource: `
resource "datapower_apiclientidentification" "test" {
  id = "APIClientIdentification_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apiclientidentification" "test" {
  depends_on = [ datapower_apiclientidentification.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APICollectionTestConfig = ModelTestConfig{
	Name: "APICollection",
	Resource: `
resource "datapower_apicollection" "test" {
  id = "APICollection_name"
  app_domain = "acc_test_domain"
  org_id = "orgid"
  org_name = "orgname"
  routing_prefix = ` + DmRoutingPrefixTestConfig.GetModelListConfig() + `
  api_processing_rule = "default-api-rule"
  api_error_rule = "default-api-error-rule"
  plan = ["TestAccAPIPlan"]
}`,
	Data: `
data "datapower_apicollection" "test" {
  depends_on = [ datapower_apicollection.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmRoutingPrefix": &DmRoutingPrefixTestConfig,
		"APIRule":         &APIRuleTestConfig,
		"APIPlan":         &APIPlanTestConfig,
	},
	TestPre: `
`,
}
var APIConnectGatewayServiceTestConfig = ModelTestConfig{
	Name: "APIConnectGatewayService",
	Resource: `
resource "datapower_apiconnectgatewayservice" "test" {
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 3000
  proxy_policy = {proxy_policy_enable = false, remote_address = "localhost", remote_port = 8080}
}`,
	Data: `
data "datapower_apiconnectgatewayservice" "test" {
  depends_on = [ datapower_apiconnectgatewayservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAPICGSProxyPolicy": &DmAPICGSProxyPolicyTestConfig,
	},
	TestPre: `
`,
}
var APIDefinitionTestConfig = ModelTestConfig{
	Name: "APIDefinition",
	Resource: `
resource "datapower_apidefinition" "test" {
  id = "APIDefinition_name"
  app_domain = "acc_test_domain"
  base_path = "/"
  path = ["TestAccAPIPath"]
  content = "activity"
  error_content = "payload"
}`,
	Data: `
data "datapower_apidefinition" "test" {
  depends_on = [ datapower_apidefinition.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APIPath": &APIPathTestConfig,
	},
	TestPre: `
`,
}
var APIExecuteTestConfig = ModelTestConfig{
	Name: "APIExecute",
	Resource: `
resource "datapower_apiexecute" "test" {
  id = "APIExecute_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apiexecute" "test" {
  depends_on = [ datapower_apiexecute.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIFinalTestConfig = ModelTestConfig{
	Name: "APIFinal",
	Resource: `
resource "datapower_apifinal" "test" {
  id = "APIFinal_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apifinal" "test" {
  depends_on = [ datapower_apifinal.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIGatewayTestConfig = ModelTestConfig{
	Name: "APIGateway",
	Resource: `
resource "datapower_apigateway" "test" {
  id = "APIGateway_name"
  app_domain = "acc_test_domain"
  front_protocol = [datapower_httpsourceprotocolhandler.test.id]
  front_timeout = 120
  front_persistent_timeout = 180
}`,
	Data: `
data "datapower_apigateway" "test" {
  depends_on = [ datapower_apigateway.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `resource "datapower_httpsourceprotocolhandler" "test" {
  id = "HTTPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
}

`,
}
var APILDAPRegistryTestConfig = ModelTestConfig{
	Name: "APILDAPRegistry",
	Resource: `
resource "datapower_apildapregistry" "test" {
  id = "APILDAPRegistry_name"
  app_domain = "acc_test_domain"
  ldap_host = "localhost"
  ldap_port = 636
  ldap_search_parameters = "TestAccLDAPSearchParameters"
}`,
	Data: `
data "datapower_apildapregistry" "test" {
  depends_on = [ datapower_apildapregistry.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"LDAPSearchParameters": &LDAPSearchParametersTestConfig,
	},
	TestPre: `
`,
}
var APIOperationTestConfig = ModelTestConfig{
	Name: "APIOperation",
	Resource: `
resource "datapower_apioperation" "test" {
  id = "APIOperation_name"
  app_domain = "acc_test_domain"
  method = "GET"
}`,
	Data: `
data "datapower_apioperation" "test" {
  depends_on = [ datapower_apioperation.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIPathTestConfig = ModelTestConfig{
	Name: "APIPath",
	Resource: `
resource "datapower_apipath" "test" {
  id = "APIPath_name"
  app_domain = "acc_test_domain"
  path = "/"
}`,
	Data: `
data "datapower_apipath" "test" {
  depends_on = [ datapower_apipath.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIPlanTestConfig = ModelTestConfig{
	Name: "APIPlan",
	Resource: `
resource "datapower_apiplan" "test" {
  id = "APIPlan_name"
  app_domain = "acc_test_domain"
  api = ["TestAccAPIDefinition"]
  rate_limit_scope = "per-application"
}`,
	Data: `
data "datapower_apiplan" "test" {
  depends_on = [ datapower_apiplan.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APIDefinition": &APIDefinitionTestConfig,
	},
	TestPre: `
`,
}
var APIResultTestConfig = ModelTestConfig{
	Name: "APIResult",
	Resource: `
resource "datapower_apiresult" "test" {
  id = "APIResult_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apiresult" "test" {
  depends_on = [ datapower_apiresult.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIRoutingTestConfig = ModelTestConfig{
	Name: "APIRouting",
	Resource: `
resource "datapower_apirouting" "test" {
  id = "APIRouting_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apirouting" "test" {
  depends_on = [ datapower_apirouting.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APIRuleTestConfig = ModelTestConfig{
	Name: "APIRule",
	Resource: `
resource "datapower_apirule" "test" {
  id = "APIRule_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apirule" "test" {
  depends_on = [ datapower_apirule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISchemaTestConfig = ModelTestConfig{
	Name: "APISchema",
	Resource: `
resource "datapower_apischema" "test" {
  id = "APISchema_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apischema" "test" {
  depends_on = [ datapower_apischema.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISecurityTestConfig = ModelTestConfig{
	Name: "APISecurity",
	Resource: `
resource "datapower_apisecurity" "test" {
  id = "APISecurity_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apisecurity" "test" {
  depends_on = [ datapower_apisecurity.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISecurityAPIKeyTestConfig = ModelTestConfig{
	Name: "APISecurityAPIKey",
	Resource: `
resource "datapower_apisecurityapikey" "test" {
  id = "APISecurityAPIKey_name"
  app_domain = "acc_test_domain"
  where = "header"
  type = "id"
  key_name = "keyname"
}`,
	Data: `
data "datapower_apisecurityapikey" "test" {
  depends_on = [ datapower_apisecurityapikey.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISecurityBasicAuthTestConfig = ModelTestConfig{
	Name: "APISecurityBasicAuth",
	Resource: `
resource "datapower_apisecuritybasicauth" "test" {
  id = "APISecurityBasicAuth_name"
  app_domain = "acc_test_domain"
  user_registry = datapower_apiauthurlregistry.test.id
}`,
	Data: `
data "datapower_apisecuritybasicauth" "test" {
  depends_on = [ datapower_apisecuritybasicauth.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `resource "datapower_apiauthurlregistry" "test" {
  id = "APIAuthURLRegistry_test"
  app_domain = "acc_test_domain"
  auth_url = "http://localhost"
}

`,
}
var APISecurityHTTPSchemeTestConfig = ModelTestConfig{
	Name: "APISecurityHTTPScheme",
	Resource: `
resource "datapower_apisecurityhttpscheme" "test" {
  id = "APISecurityHTTPScheme_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apisecurityhttpscheme" "test" {
  depends_on = [ datapower_apisecurityhttpscheme.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISecurityOAuthTestConfig = ModelTestConfig{
	Name: "APISecurityOAuth",
	Resource: `
resource "datapower_apisecurityoauth" "test" {
  id = "APISecurityOAuth_name"
  app_domain = "acc_test_domain"
  o_auth_flow = "implicit"
}`,
	Data: `
data "datapower_apisecurityoauth" "test" {
  depends_on = [ datapower_apisecurityoauth.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var APISecurityOAuthReqTestConfig = ModelTestConfig{
	Name: "APISecurityOAuthReq",
	Resource: `
resource "datapower_apisecurityoauthreq" "test" {
  id = "APISecurityOAuthReq_name"
  app_domain = "acc_test_domain"
  api_security_o_auth_def = "TestAccAPISecurityOAuth"
}`,
	Data: `
data "datapower_apisecurityoauthreq" "test" {
  depends_on = [ datapower_apisecurityoauthreq.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APISecurityOAuth": &APISecurityOAuthTestConfig,
	},
	TestPre: `
`,
}
var APISecurityRequirementTestConfig = ModelTestConfig{
	Name: "APISecurityRequirement",
	Resource: `
resource "datapower_apisecurityrequirement" "test" {
  id = "APISecurityRequirement_name"
  app_domain = "acc_test_domain"
  security = [datapower_apisecurityapikey.test.id]
}`,
	Data: `
data "datapower_apisecurityrequirement" "test" {
  depends_on = [ datapower_apisecurityrequirement.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `resource "datapower_apisecurityapikey" "test" {
  id = "APISecurityAPIKey_test"
  app_domain = "acc_test_domain"
  where = "header"
  type = "id"
  key_name = "keyname"
}

`,
}
var APISecurityTokenManagerTestConfig = ModelTestConfig{
	Name: "APISecurityTokenManager",
	Resource: `
resource "datapower_apisecuritytokenmanager" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_apisecuritytokenmanager" "test" {
  depends_on = [ datapower_apisecuritytokenmanager.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AS1PollerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "AS1PollerSourceProtocolHandler",
	Resource: `
resource "datapower_as1pollersourceprotocolhandler" "test" {
  id = "AS1PollerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  mail_server = "smtp.host"
  port = 25
  conn_security = "none"
  account = "account"
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
	Data: `
data "datapower_as1pollersourceprotocolhandler" "test" {
  depends_on = [ datapower_as1pollersourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AS2ProxySourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "AS2ProxySourceProtocolHandler",
	Resource: `
resource "datapower_as2proxysourceprotocolhandler" "test" {
  id = "AS2ProxySourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
  remote_address = "10.10.10.10"
  remote_port = 8888
  remote_connection_timeout = 60
  xml_manager = "default"
}`,
	Data: `
data "datapower_as2proxysourceprotocolhandler" "test" {
  depends_on = [ datapower_as2proxysourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var AS2SourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "AS2SourceProtocolHandler",
	Resource: `
resource "datapower_as2sourceprotocolhandler" "test" {
  id = "AS2SourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
}`,
	Data: `
data "datapower_as2sourceprotocolhandler" "test" {
  depends_on = [ datapower_as2sourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AS3SourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "AS3SourceProtocolHandler",
	Resource: `
resource "datapower_as3sourceprotocolhandler" "test" {
  id = "AS3SourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 21
}`,
	Data: `
data "datapower_as3sourceprotocolhandler" "test" {
  depends_on = [ datapower_as3sourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AccessControlListTestConfig = ModelTestConfig{
	Name: "AccessControlList",
	Resource: `
resource "datapower_accesscontrollist" "test" {
  id = "AccessControlList_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_accesscontrollist" "test" {
  depends_on = [ datapower_accesscontrollist.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AccessProfileTestConfig = ModelTestConfig{
	Name: "AccessProfile",
	Resource: `
resource "datapower_accessprofile" "test" {
  id = "AccessProfile_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_accessprofile" "test" {
  depends_on = [ datapower_accessprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AnalyticsEndpointTestConfig = ModelTestConfig{
	Name: "AnalyticsEndpoint",
	Resource: `
resource "datapower_analyticsendpoint" "test" {
  id = "AnalyticsEndpoint_name"
  app_domain = "acc_test_domain"
  analytics_server_url = "https://localhost"
  max_records = 1024
  max_records_memory_kb = 512
  max_delivery_memory_mb = 512
}`,
	Data: `
data "datapower_analyticsendpoint" "test" {
  depends_on = [ datapower_analyticsendpoint.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AppSecurityPolicyTestConfig = ModelTestConfig{
	Name: "AppSecurityPolicy",
	Resource: `
resource "datapower_appsecuritypolicy" "test" {
  id = "AppSecurityPolicy_name"
  app_domain = "acc_test_domain"
  request_maps = ` + DmWebAppRequestPolicyMapTestConfig.GetModelListConfig() + `
  response_maps = ` + DmWebAppResponsePolicyMapTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_appsecuritypolicy" "test" {
  depends_on = [ datapower_appsecuritypolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmWebAppRequestPolicyMap":  &DmWebAppRequestPolicyMapTestConfig,
		"DmWebAppResponsePolicyMap": &DmWebAppResponsePolicyMapTestConfig,
	},
	TestPre: `
`,
}
var AssemblyTestConfig = ModelTestConfig{
	Name: "Assembly",
	Resource: `
resource "datapower_assembly" "test" {
  id = "Assembly_name"
  app_domain = "acc_test_domain"
  rule = "default-empty-rule"
}`,
	Data: `
data "datapower_assembly" "test" {
  depends_on = [ datapower_assembly.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APIRule": &APIRuleTestConfig,
	},
	TestPre: `
`,
}
var AssemblyActionClientSecurityTestConfig = ModelTestConfig{
	Name: "AssemblyActionClientSecurity",
	Resource: `
resource "datapower_assemblyactionclientsecurity" "test" {
  id = "AssemblyActionClientSecurity_name"
  app_domain = "acc_test_domain"
  authenticate_client_method = "native"
}`,
	Data: `
data "datapower_assemblyactionclientsecurity" "test" {
  depends_on = [ datapower_assemblyactionclientsecurity.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionExtractTestConfig = ModelTestConfig{
	Name: "AssemblyActionExtract",
	Resource: `
resource "datapower_assemblyactionextract" "test" {
  id = "AssemblyActionExtract_name"
  app_domain = "acc_test_domain"
  extract = ` + DmAssemblyActionExtractTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_assemblyactionextract" "test" {
  depends_on = [ datapower_assemblyactionextract.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAssemblyActionExtract": &DmAssemblyActionExtractTestConfig,
	},
	TestPre: `
`,
}
var AssemblyActionFunctionCallTestConfig = ModelTestConfig{
	Name: "AssemblyActionFunctionCall",
	Resource: `
resource "datapower_assemblyactionfunctioncall" "test" {
  id = "AssemblyActionFunctionCall_name"
  app_domain = "acc_test_domain"
  function_call = "default-func-global"
}`,
	Data: `
data "datapower_assemblyactionfunctioncall" "test" {
  depends_on = [ datapower_assemblyactionfunctioncall.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"AssemblyFunction": &AssemblyFunctionTestConfig,
	},
	TestPre: `
`,
}
var AssemblyActionGatewayScriptTestConfig = ModelTestConfig{
	Name: "AssemblyActionGatewayScript",
	Resource: `
resource "datapower_assemblyactiongatewayscript" "test" {
  id = "AssemblyActionGatewayScript_name"
  app_domain = "acc_test_domain"
  source = "gsfile"
}`,
	Data: `
data "datapower_assemblyactiongatewayscript" "test" {
  depends_on = [ datapower_assemblyactiongatewayscript.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionGraphQLCostAnalysisTestConfig = ModelTestConfig{
	Name: "AssemblyActionGraphQLCostAnalysis",
	Resource: `
resource "datapower_assemblyactiongraphqlcostanalysis" "test" {
  id = "AssemblyActionGraphQLCostAnalysis_name"
  app_domain = "acc_test_domain"
  target = "targetquery"
}`,
	Data: `
data "datapower_assemblyactiongraphqlcostanalysis" "test" {
  depends_on = [ datapower_assemblyactiongraphqlcostanalysis.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionGraphQLExecuteTestConfig = ModelTestConfig{
	Name: "AssemblyActionGraphQLExecute",
	Resource: `
resource "datapower_assemblyactiongraphqlexecute" "test" {
  id = "AssemblyActionGraphQLExecute_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactiongraphqlexecute" "test" {
  depends_on = [ datapower_assemblyactiongraphqlexecute.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionGraphQLIntrospectTestConfig = ModelTestConfig{
	Name: "AssemblyActionGraphQLIntrospect",
	Resource: `
resource "datapower_assemblyactiongraphqlintrospect" "test" {
  id = "AssemblyActionGraphQLIntrospect_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactiongraphqlintrospect" "test" {
  depends_on = [ datapower_assemblyactiongraphqlintrospect.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionHtmlPageTestConfig = ModelTestConfig{
	Name: "AssemblyActionHtmlPage",
	Resource: `
resource "datapower_assemblyactionhtmlpage" "test" {
  id = "AssemblyActionHtmlPage_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionhtmlpage" "test" {
  depends_on = [ datapower_assemblyactionhtmlpage.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionInvokeTestConfig = ModelTestConfig{
	Name: "AssemblyActionInvoke",
	Resource: `
resource "datapower_assemblyactioninvoke" "test" {
  id = "AssemblyActionInvoke_name"
  app_domain = "acc_test_domain"
  url = "https://localhost"
  method = "Keep"
  backend_type = "detect"
  cache_type = "Protocol"
}`,
	Data: `
data "datapower_assemblyactioninvoke" "test" {
  depends_on = [ datapower_assemblyactioninvoke.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionJWTGenerateTestConfig = ModelTestConfig{
	Name: "AssemblyActionJWTGenerate",
	Resource: `
resource "datapower_assemblyactionjwtgenerate" "test" {
  id = "AssemblyActionJWTGenerate_name"
  app_domain = "acc_test_domain"
  issuer_claim = "iss.claim"
  validity_period = 3600
}`,
	Data: `
data "datapower_assemblyactionjwtgenerate" "test" {
  depends_on = [ datapower_assemblyactionjwtgenerate.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionJWTValidateTestConfig = ModelTestConfig{
	Name: "AssemblyActionJWTValidate",
	Resource: `
resource "datapower_assemblyactionjwtvalidate" "test" {
  id = "AssemblyActionJWTValidate_name"
  app_domain = "acc_test_domain"
  jwt = "request.headers.authorization"
  output_claims = "decoded.claims"
}`,
	Data: `
data "datapower_assemblyactionjwtvalidate" "test" {
  depends_on = [ datapower_assemblyactionjwtvalidate.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionJson2XmlTestConfig = ModelTestConfig{
	Name: "AssemblyActionJson2Xml",
	Resource: `
resource "datapower_assemblyactionjson2xml" "test" {
  id = "AssemblyActionJson2Xml_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionjson2xml" "test" {
  depends_on = [ datapower_assemblyactionjson2xml.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionLogTestConfig = ModelTestConfig{
	Name: "AssemblyActionLog",
	Resource: `
resource "datapower_assemblyactionlog" "test" {
  id = "AssemblyActionLog_name"
  app_domain = "acc_test_domain"
  mode = "gather-only"
}`,
	Data: `
data "datapower_assemblyactionlog" "test" {
  depends_on = [ datapower_assemblyactionlog.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionMapTestConfig = ModelTestConfig{
	Name: "AssemblyActionMap",
	Resource: `
resource "datapower_assemblyactionmap" "test" {
  id = "AssemblyActionMap_name"
  app_domain = "acc_test_domain"
  location = "local:///file"
}`,
	Data: `
data "datapower_assemblyactionmap" "test" {
  depends_on = [ datapower_assemblyactionmap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionOAuthTestConfig = ModelTestConfig{
	Name: "AssemblyActionOAuth",
	Resource: `
resource "datapower_assemblyactionoauth" "test" {
  id = "AssemblyActionOAuth_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionoauth" "test" {
  depends_on = [ datapower_assemblyactionoauth.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionParseTestConfig = ModelTestConfig{
	Name: "AssemblyActionParse",
	Resource: `
resource "datapower_assemblyactionparse" "test" {
  id = "AssemblyActionParse_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionparse" "test" {
  depends_on = [ datapower_assemblyactionparse.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionRateLimitTestConfig = ModelTestConfig{
	Name: "AssemblyActionRateLimit",
	Resource: `
resource "datapower_assemblyactionratelimit" "test" {
  id = "AssemblyActionRateLimit_name"
  app_domain = "acc_test_domain"
  source = "plan-default"
}`,
	Data: `
data "datapower_assemblyactionratelimit" "test" {
  depends_on = [ datapower_assemblyactionratelimit.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionRateLimitInfoTestConfig = ModelTestConfig{
	Name: "AssemblyActionRateLimitInfo",
	Resource: `
resource "datapower_assemblyactionratelimitinfo" "test" {
  id = "AssemblyActionRateLimitInfo_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionratelimitinfo" "test" {
  depends_on = [ datapower_assemblyactionratelimitinfo.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionRedactTestConfig = ModelTestConfig{
	Name: "AssemblyActionRedact",
	Resource: `
resource "datapower_assemblyactionredact" "test" {
  id = "AssemblyActionRedact_name"
  app_domain = "acc_test_domain"
  redact = ` + DmAssemblyActionRedactTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_assemblyactionredact" "test" {
  depends_on = [ datapower_assemblyactionredact.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAssemblyActionRedact": &DmAssemblyActionRedactTestConfig,
	},
	TestPre: `
`,
}
var AssemblyActionSetVarTestConfig = ModelTestConfig{
	Name: "AssemblyActionSetVar",
	Resource: `
resource "datapower_assemblyactionsetvar" "test" {
  id = "AssemblyActionSetVar_name"
  app_domain = "acc_test_domain"
  variable = ` + DmAssemblyActionSetVarTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_assemblyactionsetvar" "test" {
  depends_on = [ datapower_assemblyactionsetvar.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAssemblyActionSetVar": &DmAssemblyActionSetVarTestConfig,
	},
	TestPre: `
`,
}
var AssemblyActionThrowTestConfig = ModelTestConfig{
	Name: "AssemblyActionThrow",
	Resource: `
resource "datapower_assemblyactionthrow" "test" {
  id = "AssemblyActionThrow_name"
  app_domain = "acc_test_domain"
  error_id = "errorid"
}`,
	Data: `
data "datapower_assemblyactionthrow" "test" {
  depends_on = [ datapower_assemblyactionthrow.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionUserSecurityTestConfig = ModelTestConfig{
	Name: "AssemblyActionUserSecurity",
	Resource: `
resource "datapower_assemblyactionusersecurity" "test" {
  id = "AssemblyActionUserSecurity_name"
  app_domain = "acc_test_domain"
  factor_id = "default"
  extract_identity_method = "basic"
  user_auth_method = "user-registry"
  user_az_method = "authenticated"
}`,
	Data: `
data "datapower_assemblyactionusersecurity" "test" {
  depends_on = [ datapower_assemblyactionusersecurity.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionValidateTestConfig = ModelTestConfig{
	Name: "AssemblyActionValidate",
	Resource: `
resource "datapower_assemblyactionvalidate" "test" {
  id = "AssemblyActionValidate_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionvalidate" "test" {
  depends_on = [ datapower_assemblyactionvalidate.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionWSDLTestConfig = ModelTestConfig{
	Name: "AssemblyActionWSDL",
	Resource: `
resource "datapower_assemblyactionwsdl" "test" {
  id = "AssemblyActionWSDL_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionwsdl" "test" {
  depends_on = [ datapower_assemblyactionwsdl.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionWebSocketUpgradeTestConfig = ModelTestConfig{
	Name: "AssemblyActionWebSocketUpgrade",
	Resource: `
resource "datapower_assemblyactionwebsocketupgrade" "test" {
  id = "AssemblyActionWebSocketUpgrade_name"
  app_domain = "acc_test_domain"
  url = "https://localhost"
}`,
	Data: `
data "datapower_assemblyactionwebsocketupgrade" "test" {
  depends_on = [ datapower_assemblyactionwebsocketupgrade.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionXSLTTestConfig = ModelTestConfig{
	Name: "AssemblyActionXSLT",
	Resource: `
resource "datapower_assemblyactionxslt" "test" {
  id = "AssemblyActionXSLT_name"
  app_domain = "acc_test_domain"
  stylesheet = "local:///stylesheet"
}`,
	Data: `
data "datapower_assemblyactionxslt" "test" {
  depends_on = [ datapower_assemblyactionxslt.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyActionXml2JsonTestConfig = ModelTestConfig{
	Name: "AssemblyActionXml2Json",
	Resource: `
resource "datapower_assemblyactionxml2json" "test" {
  id = "AssemblyActionXml2Json_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyactionxml2json" "test" {
  depends_on = [ datapower_assemblyactionxml2json.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyFunctionTestConfig = ModelTestConfig{
	Name: "AssemblyFunction",
	Resource: `
resource "datapower_assemblyfunction" "test" {
  id = "AssemblyFunction_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_assemblyfunction" "test" {
  depends_on = [ datapower_assemblyfunction.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var AssemblyLogicOperationSwitchTestConfig = ModelTestConfig{
	Name: "AssemblyLogicOperationSwitch",
	Resource: `
resource "datapower_assemblylogicoperationswitch" "test" {
  id = "AssemblyLogicOperationSwitch_name"
  app_domain = "acc_test_domain"
  case = ` + DmAssemblyLogicOperationSwitchCaseTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_assemblylogicoperationswitch" "test" {
  depends_on = [ datapower_assemblylogicoperationswitch.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAssemblyLogicOperationSwitchCase": &DmAssemblyLogicOperationSwitchCaseTestConfig,
	},
	TestPre: `
`,
}
var AssemblyLogicSwitchTestConfig = ModelTestConfig{
	Name: "AssemblyLogicSwitch",
	Resource: `
resource "datapower_assemblylogicswitch" "test" {
  id = "AssemblyLogicSwitch_name"
  app_domain = "acc_test_domain"
  case = ` + DmAssemblyLogicExecuteTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_assemblylogicswitch" "test" {
  depends_on = [ datapower_assemblylogicswitch.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmAssemblyLogicExecute": &DmAssemblyLogicExecuteTestConfig,
	},
	TestPre: `
`,
}
var AuditLogTestConfig = ModelTestConfig{
	Name: "AuditLog",
	Resource: `
resource "datapower_auditlog" "test" {
  audit_level = "full"
}`,
	Data: `
data "datapower_auditlog" "test" {
  depends_on = [ datapower_auditlog.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var B2BCPATestConfig = ModelTestConfig{
	Name: "B2BCPA",
	Resource: `
resource "datapower_b2bcpa" "test" {
  id = "B2BCPA_name"
  app_domain = "acc_test_domain"
  cpa_id = "cpaid"
}`,
	Data: `
data "datapower_b2bcpa" "test" {
  depends_on = [ datapower_b2bcpa.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var B2BCPACollaborationTestConfig = ModelTestConfig{
	Name: "B2BCPACollaboration",
	Resource: `
resource "datapower_b2bcpacollaboration" "test" {
  id = "B2BCPACollaboration_name"
  app_domain = "acc_test_domain"
  service = "service"
}`,
	Data: `
data "datapower_b2bcpacollaboration" "test" {
  depends_on = [ datapower_b2bcpacollaboration.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var B2BCPAReceiverSettingTestConfig = ModelTestConfig{
	Name: "B2BCPAReceiverSetting",
	Resource: `
resource "datapower_b2bcpareceiversetting" "test" {
  id = "B2BCPAReceiverSetting_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_b2bcpareceiversetting" "test" {
  depends_on = [ datapower_b2bcpareceiversetting.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var B2BCPASenderSettingTestConfig = ModelTestConfig{
	Name: "B2BCPASenderSetting",
	Resource: `
resource "datapower_b2bcpasendersetting" "test" {
  id = "B2BCPASenderSetting_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_b2bcpasendersetting" "test" {
  depends_on = [ datapower_b2bcpasendersetting.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var B2BGatewayTestConfig = ModelTestConfig{
	Name: "B2BGateway",
	Resource: `
resource "datapower_b2bgateway" "test" {
  id = "B2BGateway_name"
  app_domain = "acc_test_domain"
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode = "archpurge"
  xml_manager = "default"
}`,
	Data: `
data "datapower_b2bgateway" "test" {
  depends_on = [ datapower_b2bgateway.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var B2BPersistenceTestConfig = ModelTestConfig{
	Name: "B2BPersistence",
	Resource: `
resource "datapower_b2bpersistence" "test" {
  raid_volume = "raid0"
  storage_size = 1024
  ha_enabled = false
}`,
	Data: `
data "datapower_b2bpersistence" "test" {
  depends_on = [ datapower_b2bpersistence.test ]
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"RaidVolume": &RaidVolumeTestConfig,
	},
	TestPre: `
`,
}
var B2BProfileTestConfig = ModelTestConfig{
	Name: "B2BProfile",
	Resource: `
resource "datapower_b2bprofile" "test" {
  id = "B2BProfile_name"
  app_domain = "acc_test_domain"
  destinations = ` + DmB2BDestinationTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_b2bprofile" "test" {
  depends_on = [ datapower_b2bprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmB2BDestination": &DmB2BDestinationTestConfig,
	},
	TestPre: `
`,
}
var B2BProfileGroupTestConfig = ModelTestConfig{
	Name: "B2BProfileGroup",
	Resource: `
resource "datapower_b2bprofilegroup" "test" {
  id = "B2BProfileGroup_name"
  app_domain = "acc_test_domain"
  b2b_profiles = ` + DmB2BGroupedProfileTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_b2bprofilegroup" "test" {
  depends_on = [ datapower_b2bprofilegroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmB2BGroupedProfile": &DmB2BGroupedProfileTestConfig,
	},
	TestPre: `
`,
}
var B2BXPathRoutingPolicyTestConfig = ModelTestConfig{
	Name: "B2BXPathRoutingPolicy",
	Resource: `
resource "datapower_b2bxpathroutingpolicy" "test" {
  id = "B2BXPathRoutingPolicy_name"
  app_domain = "acc_test_domain"
  sender_x_path = "senderpath"
  receiver_x_path = "senderpath"
}`,
	Data: `
data "datapower_b2bxpathroutingpolicy" "test" {
  depends_on = [ datapower_b2bxpathroutingpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CORSPolicyTestConfig = ModelTestConfig{
	Name: "CORSPolicy",
	Resource: `
resource "datapower_corspolicy" "test" {
  id = "CORSPolicy_name"
  app_domain = "acc_test_domain"
  rule = ["TestAccCORSRule"]
}`,
	Data: `
data "datapower_corspolicy" "test" {
  depends_on = [ datapower_corspolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"CORSRule": &CORSRuleTestConfig,
	},
	TestPre: `
`,
}
var CORSRuleTestConfig = ModelTestConfig{
	Name: "CORSRule",
	Resource: `
resource "datapower_corsrule" "test" {
  id = "CORSRule_name"
  app_domain = "acc_test_domain"
  allow_origin = ["origin"]
}`,
	Data: `
data "datapower_corsrule" "test" {
  depends_on = [ datapower_corsrule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CRLFetchTestConfig = ModelTestConfig{
	Name: "CRLFetch",
	Resource: `
resource "datapower_crlfetch" "test" {
}`,
	Data: `
data "datapower_crlfetch" "test" {
  depends_on = [ datapower_crlfetch.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CertMonitorTestConfig = ModelTestConfig{
	Name: "CertMonitor",
	Resource: `
resource "datapower_certmonitor" "test" {
  polling_interval = 1
  reminder_time = 30
  log_level = "warn"
  disable_expired_certs = false
}`,
	Data: `
data "datapower_certmonitor" "test" {
  depends_on = [ datapower_certmonitor.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CompileOptionsPolicyTestConfig = ModelTestConfig{
	Name: "CompileOptionsPolicy",
	Resource: `
resource "datapower_compileoptionspolicy" "test" {
  id = "CompileOptionsPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_compileoptionspolicy" "test" {
  depends_on = [ datapower_compileoptionspolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CompileSettingsTestConfig = ModelTestConfig{
	Name: "CompileSettings",
	Resource: `
resource "datapower_compilesettings" "test" {
  id = "CompileSettings_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_compilesettings" "test" {
  depends_on = [ datapower_compilesettings.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ConfigDeploymentPolicyTestConfig = ModelTestConfig{
	Name: "ConfigDeploymentPolicy",
	Resource: `
resource "datapower_configdeploymentpolicy" "test" {
  id = "ConfigDeploymentPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_configdeploymentpolicy" "test" {
  depends_on = [ datapower_configdeploymentpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ConfigSequenceTestConfig = ModelTestConfig{
	Name: "ConfigSequence",
	Resource: `
resource "datapower_configsequence" "test" {
  id = "ConfigSequence_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_configsequence" "test" {
  depends_on = [ datapower_configsequence.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ConformancePolicyTestConfig = ModelTestConfig{
	Name: "ConformancePolicy",
	Resource: `
resource "datapower_conformancepolicy" "test" {
  id = "ConformancePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_conformancepolicy" "test" {
  depends_on = [ datapower_conformancepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ControlListTestConfig = ModelTestConfig{
	Name: "ControlList",
	Resource: `
resource "datapower_controllist" "test" {
  id = "ControlList_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_controllist" "test" {
  depends_on = [ datapower_controllist.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CookieAttributePolicyTestConfig = ModelTestConfig{
	Name: "CookieAttributePolicy",
	Resource: `
resource "datapower_cookieattributepolicy" "test" {
  id = "CookieAttributePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_cookieattributepolicy" "test" {
  depends_on = [ datapower_cookieattributepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CountMonitorTestConfig = ModelTestConfig{
	Name: "CountMonitor",
	Resource: `
resource "datapower_countmonitor" "test" {
  id = "CookieAttributePolicy_name"
  app_domain = "acc_test_domain"
  measure = "requests"
  source = "all"
  max_sources = 10000
  message_type = "TestAccMessageType"
}`,
	Data: `
data "datapower_countmonitor" "test" {
  depends_on = [ datapower_countmonitor.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"MessageType": &MessageTypeTestConfig,
	},
	TestPre: `
`,
}
var CryptoCertificateTestConfig = ModelTestConfig{
	Name: "CryptoCertificate",
	Resource: `
resource "datapower_cryptocertificate" "test" {
  id = "CryptoCertificate_name"
  app_domain = "acc_test_domain"
  filename = "cert:///acc-test-server.crt"
}`,
	Data: `
data "datapower_cryptocertificate" "test" {
  depends_on = [ datapower_cryptocertificate.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CryptoFWCredTestConfig = ModelTestConfig{
	Name: "CryptoFWCred",
	Resource: `
resource "datapower_cryptofwcred" "test" {
  id = "CryptoFWCred_name"
  app_domain = "acc_test_domain"
  private_key = ["TestAccCryptoKey"]
}`,
	Data: `
data "datapower_cryptofwcred" "test" {
  depends_on = [ datapower_cryptofwcred.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoKey": &CryptoKeyTestConfig,
	},
	TestPre: `
`,
}
var CryptoIdentCredTestConfig = ModelTestConfig{
	Name: "CryptoIdentCred",
	Resource: `
resource "datapower_cryptoidentcred" "test" {
  id = "CookieAttributePolicy_name"
  app_domain = "acc_test_domain"
  key = "TestAccCryptoKey"
  certificate = "TestAccCryptoCertificate"
}`,
	Data: `
data "datapower_cryptoidentcred" "test" {
  depends_on = [ datapower_cryptoidentcred.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoKey":         &CryptoKeyTestConfig,
		"CryptoCertificate": &CryptoCertificateTestConfig,
	},
	TestPre: `
`,
}
var CryptoKerberosKDCTestConfig = ModelTestConfig{
	Name: "CryptoKerberosKDC",
	Resource: `
resource "datapower_cryptokerberoskdc" "test" {
  id = "CryptoKerberosKDC_name"
  app_domain = "acc_test_domain"
  realm = "realm"
  server = "localhost"
}`,
	Data: `
data "datapower_cryptokerberoskdc" "test" {
  depends_on = [ datapower_cryptokerberoskdc.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CryptoKerberosKeytabTestConfig = ModelTestConfig{
	Name: "CryptoKerberosKeytab",
	Resource: `
resource "datapower_cryptokerberoskeytab" "test" {
  id = "CryptoKerberosKeytab_name"
  app_domain = "acc_test_domain"
  filename = "cert:///kerberos-keytab"
}`,
	Data: `
data "datapower_cryptokerberoskeytab" "test" {
  depends_on = [ datapower_cryptokerberoskeytab.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CryptoKeyTestConfig = ModelTestConfig{
	Name: "CryptoKey",
	Resource: `
resource "datapower_cryptokey" "test" {
  id = "CryptoKey_name"
  app_domain = "acc_test_domain"
  filename = "cert:///acc-test-server.key"
  alias = "TestAccPasswordAlias"
}`,
	Data: `
data "datapower_cryptokey" "test" {
  depends_on = [ datapower_cryptokey.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"PasswordAlias": &PasswordAliasTestConfig,
	},
	TestPre: `
`,
}
var CryptoSSKeyTestConfig = ModelTestConfig{
	Name: "CryptoSSKey",
	Resource: `
resource "datapower_cryptosskey" "test" {
  id = "CryptoSSKey_name"
  app_domain = "acc_test_domain"
  filename = "cert:///acc-test-server.key"
}`,
	Data: `
data "datapower_cryptosskey" "test" {
  depends_on = [ datapower_cryptosskey.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var CryptoValCredTestConfig = ModelTestConfig{
	Name: "CryptoValCred",
	Resource: `
resource "datapower_cryptovalcred" "test" {
  id = "CryptoValCred_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_cryptovalcred" "test" {
  depends_on = [ datapower_cryptovalcred.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DNSNameServiceTestConfig = ModelTestConfig{
	Name: "DNSNameService",
	Resource: `
resource "datapower_dnsnameservice" "test" {
  load_balance_algorithm = "first-alive"
}`,
	Data: `
data "datapower_dnsnameservice" "test" {
  depends_on = [ datapower_dnsnameservice.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DeploymentPolicyParametersBindingTestConfig = ModelTestConfig{
	Name: "DeploymentPolicyParametersBinding",
	Resource: `
resource "datapower_deploymentpolicyparametersbinding" "test" {
  id = "DeploymentPolicyParametersBinding_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_deploymentpolicyparametersbinding" "test" {
  depends_on = [ datapower_deploymentpolicyparametersbinding.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DistributedVariableTestConfig = ModelTestConfig{
	Name: "DistributedVariable",
	Resource: `
resource "datapower_distributedvariable" "test" {
  app_domain = "acc_test_domain"
  gateway_peering = "default-gateway-peering"
}`,
	Data: `
data "datapower_distributedvariable" "test" {
  depends_on = [ datapower_distributedvariable.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"GatewayPeering": &GatewayPeeringTestConfig,
	},
	TestPre: `
`,
}
var DmAAAPAuthenticateTestConfig = ModelTestConfig{
	Name: "DmAAAPAuthenticate",
	Model: `{
  au_method = "ldap"
  au_cache_allow = "absolute"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPAuthorizeTestConfig = ModelTestConfig{
	Name: "DmAAAPAuthorize",
	Model: `{
  az_method = "anyauthenticated"
  az_cache_allow = "absolute"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPEIBitmapTestConfig = ModelTestConfig{
	Name: "DmAAAPEIBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPERBitmapTestConfig = ModelTestConfig{
	Name: "DmAAAPERBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPExtractIdentityTestConfig = ModelTestConfig{
	Name: "DmAAAPExtractIdentity",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPExtractResourceTestConfig = ModelTestConfig{
	Name: "DmAAAPExtractResource",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPMapCredentialsTestConfig = ModelTestConfig{
	Name: "DmAAAPMapCredentials",
	Model: `{
  mc_method = "none"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPMapResourceTestConfig = ModelTestConfig{
	Name: "DmAAAPMapResource",
	Model: `{
  mr_method = "none"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAAPPostProcessTestConfig = ModelTestConfig{
	Name: "DmAAAPPostProcess",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAAATransactionPriorityTestConfig = ModelTestConfig{
	Name: "DmAAATransactionPriority",
	Model: `{
  credential = "cred"
  priority = "unknown"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmACETestConfig = ModelTestConfig{
	Name: "DmACE",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIBurstLimitTestConfig = ModelTestConfig{
	Name: "DmAPIBurstLimit",
	Model: `{
  burst = 1000
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPICGSProxyPolicyTestConfig = ModelTestConfig{
	Name: "DmAPICGSProxyPolicy",
	Model: `{
  proxy_policy_enable = false
  remote_address = "localhost"
  remote_port = 8080
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPICountLimitTestConfig = ModelTestConfig{
	Name: "DmAPICountLimit",
	Model: `{
  count = 1000
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIDataTypeDefinitionTestConfig = ModelTestConfig{
	Name: "DmAPIDataTypeDefinition",
	Model: `{
  name = "dtdefname"
  schema = "TestAccAPISchema"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"APISchema": &APISchemaTestConfig,
	},
	TestPre: `
`,
}
var DmAPIParameterTestConfig = ModelTestConfig{
	Name: "DmAPIParameter",
	Model: `{
  required = false
  type = "string"
  where = "path"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIPropertyTestConfig = ModelTestConfig{
	Name: "DmAPIProperty",
	Model: `{
  property_name = "propertyname"
  catalog = "*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIProtocolsTestConfig = ModelTestConfig{
	Name: "DmAPIProtocols",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIProxyPolicyTestConfig = ModelTestConfig{
	Name: "DmAPIProxyPolicy",
	Model: `{
  reg_exp = "*"
  skip = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIRateLimitTestConfig = ModelTestConfig{
	Name: "DmAPIRateLimit",
	Model: `{
  rate = 1000
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAPIResponseSchemaTestConfig = ModelTestConfig{
	Name: "DmAPIResponseSchema",
	Model: `{
  status_code = "200"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmASFrontProtocolTestConfig = ModelTestConfig{
	Name: "DmASFrontProtocol",
	Model: `{
  mdn_receiver = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAddHeaderPolicyTestConfig = ModelTestConfig{
	Name: "DmAddHeaderPolicy",
	Model: `{
  reg_exp = "*"
  add_header = "HEADER"
  add_value = "VALUE"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAllowCompressionPolicyTestConfig = ModelTestConfig{
	Name: "DmAllowCompressionPolicy",
	Model: `{
  reg_exp = "*"
  allow_compression = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAllowedClientTypeTestConfig = ModelTestConfig{
	Name: "DmAllowedClientType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyActionExtractTestConfig = ModelTestConfig{
	Name: "DmAssemblyActionExtract",
	Model: `{
  capture = "capture"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyActionFunctionCallParameterTestConfig = ModelTestConfig{
	Name: "DmAssemblyActionFunctionCallParameter",
	Model: `{
  name = "assemblyactionfunctioncallparametername"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyActionRedactTestConfig = ModelTestConfig{
	Name: "DmAssemblyActionRedact",
	Model: `{
  path = "path"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyActionSetVarTestConfig = ModelTestConfig{
	Name: "DmAssemblyActionSetVar",
	Model: `{
  name = "varname"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyCatchTestConfig = ModelTestConfig{
	Name: "DmAssemblyCatch",
	Model: `{
  error = "errorname"
  handler = "default-api-rule"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"APIRule": &APIRuleTestConfig,
	},
	TestPre: `
`,
}
var DmAssemblyFunctionParameterTestConfig = ModelTestConfig{
	Name: "DmAssemblyFunctionParameter",
	Model: `{
  name = "assemblyfunctionparametername"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyLogicExecuteTestConfig = ModelTestConfig{
	Name: "DmAssemblyLogicExecute",
	Model: `{
  condition = "condition"
  execute = "default-api-rule"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmAssemblyLogicOperationSwitchCaseTestConfig = ModelTestConfig{
	Name: "DmAssemblyLogicOperationSwitchCase",
	Model: `{
  execute = "default-api-rule"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BActiveGroupTestConfig = ModelTestConfig{
	Name: "DmB2BActiveGroup",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BActiveProfileTestConfig = ModelTestConfig{
	Name: "DmB2BActiveProfile",
	Model: `{
  partner_profile = "TestAccB2BProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"B2BProfile": &B2BProfileTestConfig,
	},
	TestPre: `
`,
}
var DmB2BBackupMsgTypeTestConfig = ModelTestConfig{
	Name: "DmB2BBackupMsgType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BCPAEntryTestConfig = ModelTestConfig{
	Name: "DmB2BCPAEntry",
	Model: `{
  cpa = "TestAccB2BCPA"
  collaboration = "TestAccB2BCPACollaboration"
  internal_partner = "TestAccB2BProfile"
  external_partner = "TestAccB2BProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"B2BCPA":              &B2BCPATestConfig,
		"B2BCPACollaboration": &B2BCPACollaborationTestConfig,
		"B2BProfile":          &B2BProfileTestConfig,
	},
	TestPre: `
`,
}
var DmB2BContactTestConfig = ModelTestConfig{
	Name: "DmB2BContact",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BDestinationTestConfig = ModelTestConfig{
	Name: "DmB2BDestination",
	Model: `{
  dest_name = "b2bdestinationname"
  dest_url = "https://localhost"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BEnabledDocTypeTestConfig = ModelTestConfig{
	Name: "DmB2BEnabledDocType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BGroupedProfileTestConfig = ModelTestConfig{
	Name: "DmB2BGroupedProfile",
	Model: `{
  partner_profile = "TestAccB2BProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"B2BProfile": &B2BProfileTestConfig,
	},
	TestPre: `
`,
}
var DmB2BHAHostTestConfig = ModelTestConfig{
	Name: "DmB2BHAHost",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmB2BMessagePropertiesTestConfig = ModelTestConfig{
	Name: "DmB2BMessageProperties",
	Model: `{
  name = "b2bmessagepropertyname"
  value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmBasicAuthPolicyTestConfig = ModelTestConfig{
	Name: "DmBasicAuthPolicy",
	Model: `{
  reg_exp = "*"
  user_name = "someuser"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCORSRuleExposeHeadersTestConfig = ModelTestConfig{
	Name: "DmCORSRuleExposeHeaders",
	Model: `{
  predefined = true
  backend = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCPACollaborationActionTestConfig = ModelTestConfig{
	Name: "DmCPACollaborationAction",
	Model: `{
  name = "cpacollaborationactionname"
  value = "value"
  capability = "cansend"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCRLFetchConfigTestConfig = ModelTestConfig{
	Name: "DmCRLFetchConfig",
	Model: `{
  issuer_valcred = "TestAccCryptoValCred"
  refresh_interval = 240
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoValCred": &CryptoValCredTestConfig,
	},
	TestPre: `
`,
}
var DmClaimTestConfig = ModelTestConfig{
	Name: "DmClaim",
	Model: `{
  value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmConditionTestConfig = ModelTestConfig{
	Name: "DmCondition",
	Model: `{
  expression = "*"
  condition_action = "*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmConfigModifyTypeTestConfig = ModelTestConfig{
	Name: "DmConfigModifyType",
	Model: `{
  match = "10.10.1.1/domainA/services/xslproxy?Value=myhost"
  type = "change"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmConfigSequenceCapabilitiesTestConfig = ModelTestConfig{
	Name: "DmConfigSequenceCapabilities",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmConfigSequenceLocationTestConfig = ModelTestConfig{
	Name: "DmConfigSequenceLocation",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmConformanceProfilesTestConfig = ModelTestConfig{
	Name: "DmConformanceProfiles",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCookieAttributeTestConfig = ModelTestConfig{
	Name: "DmCookieAttribute",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCookieProfileTestConfig = ModelTestConfig{
	Name: "DmCookieProfile",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCountLimitInfoTestConfig = ModelTestConfig{
	Name: "DmCountLimitInfo",
	Model: `{
  name = "countlimitinfoname"
  action = "inc"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmCountMonitorFilterTestConfig = ModelTestConfig{
	Name: "DmCountMonitorFilter",
	Model: `{
  name = "Filter1"
  interval = 1000
  rate_limit = 50
  burst_limit = 100
  action = "TestAccFilterAction"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"FilterAction": &FilterActionTestConfig,
	},
	TestPre: `
`,
}
var DmDefinitionLinkTestConfig = ModelTestConfig{
	Name: "DmDefinitionLink",
	Model: `{
  short_name = "shortname"
  definition = "TestAccRateLimitDefinition"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"RateLimitDefinition": &RateLimitDefinitionTestConfig,
	},
	TestPre: `
`,
}
var DmDeploymentPolicyParameterTestConfig = ModelTestConfig{
	Name: "DmDeploymentPolicyParameter",
	Model: `{
  parameter_name = "parameter_name"
  parameter_value = "parameter_value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDocCachePolicyTestConfig = ModelTestConfig{
	Name: "DmDocCachePolicy",
	Model: `{
  type = "protocol"
  priority = 128
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDomainFileMapTestConfig = ModelTestConfig{
	Name: "DmDomainFileMap",
	Model: `{
  copy_from = true
  copy_to = true
  delete = true
  display = true
  exec = true
  subdir = true
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDomainMonitoringMapTestConfig = ModelTestConfig{
	Name: "DmDomainMonitoringMap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDurationMonitorFilterTestConfig = ModelTestConfig{
	Name: "DmDurationMonitorFilter",
	Model: `{
  name = "DmDurationMonitorFilter_name"
  value = 1
  action = "TestAccFilterAction"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"FilterAction": &FilterActionTestConfig,
	},
	TestPre: `
`,
}
var DmDynamicOAuthProviderSettingsReferenceTestConfig = ModelTestConfig{
	Name: "DmDynamicOAuthProviderSettingsReference",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDynamicParseSettingsReferenceTestConfig = ModelTestConfig{
	Name: "DmDynamicParseSettingsReference",
	Model: `{
  url = "some_url"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmDynamicStylePolicyActionBaseReferenceTestConfig = ModelTestConfig{
	Name: "DmDynamicStylePolicyActionBaseReference",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmExtensionFunctionTestConfig = ModelTestConfig{
	Name: "DmExtensionFunction",
	Model: `{
  extension_function_namespace = "http://www.fredspace/extensions"
  extension_function = "nodeset()"
  local_function = "node-set()"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmExternalAttachedPolicyTestConfig = ModelTestConfig{
	Name: "DmExternalAttachedPolicy",
	Model: `{
  external_attach_policy_url = "some.url"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmFTPPolicyTestConfig = ModelTestConfig{
	Name: "DmFTPPolicy",
	Model: `{
  reg_exp = "*"
  passive = "pasv-req"
  auth_tls = "auth-off"
  use_ccc = "ccc-off"
  encrypt_data = "enc-data-off"
  data_type = "binary"
  slash_stou = "slash-stou-on"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmFTPQuotedCommandTestConfig = ModelTestConfig{
	Name: "DmFTPQuotedCommand",
	Model: `{
  quoted_command = "ls"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmFTPServerVirtualDirectoryTestConfig = ModelTestConfig{
	Name: "DmFTPServerVirtualDirectory",
	Model: `{
  virtual_path = "virtualpath"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmFileSystemUsageTestConfig = ModelTestConfig{
	Name: "DmFileSystemUsage",
	Model: `{
  warning_threshold = 75
  critical_threshold = 90
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmFrontSideTestConfig = ModelTestConfig{
	Name: "DmFrontSide",
	Model: `{
  local_address = "0.0.0.0"
  local_port = 3000
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGatewayInOrderModeTestConfig = ModelTestConfig{
	Name: "DmGatewayInOrderMode",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGatewayPeeringClusterNodeTestConfig = ModelTestConfig{
	Name: "DmGatewayPeeringClusterNode",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGatewayPeeringGroupClusterNodeTestConfig = ModelTestConfig{
	Name: "DmGatewayPeeringGroupClusterNode",
	Model: `{
  host = "localhost"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGatewayPeeringGroupPeerNodeTestConfig = ModelTestConfig{
	Name: "DmGatewayPeeringGroupPeerNode",
	Model: `{
  host = "localhost"
  priority = 100
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGitOpsTemplateEntryTestConfig = ModelTestConfig{
	Name: "DmGitOpsTemplateEntry",
	Model: `{
  template_type = "change"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGitOpsTemplatePolicyTestConfig = ModelTestConfig{
	Name: "DmGitOpsTemplatePolicy",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGitOpsVariableEntryTestConfig = ModelTestConfig{
	Name: "DmGitOpsVariableEntry",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmGssChecksumFlagsTestConfig = ModelTestConfig{
	Name: "DmGssChecksumFlags",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPClientServerVersionTestConfig = ModelTestConfig{
	Name: "DmHTTPClientServerVersion",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPHeaderTestConfig = ModelTestConfig{
	Name: "DmHTTPHeader",
	Model: `{
  name = "HEADER"
  value = "VALUE"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPRequestMethodsTestConfig = ModelTestConfig{
	Name: "DmHTTPRequestMethods",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPResponseCodesTestConfig = ModelTestConfig{
	Name: "DmHTTPResponseCodes",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPVersionMaskTestConfig = ModelTestConfig{
	Name: "DmHTTPVersionMask",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHTTPVersionPolicyTestConfig = ModelTestConfig{
	Name: "DmHTTPVersionPolicy",
	Model: `{
  reg_exp = "*"
  version = "HTTP/1.1"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHeaderInjectionTestConfig = ModelTestConfig{
	Name: "DmHeaderInjection",
	Model: `{
  header_tag_value = "Some Header Value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHeaderRetentionBitmapTestConfig = ModelTestConfig{
	Name: "DmHeaderRetentionBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHeaderRetentionPolicyTestConfig = ModelTestConfig{
	Name: "DmHeaderRetentionPolicy",
	Model: `{
  reg_exp = "*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHeaderSuppressionTestConfig = ModelTestConfig{
	Name: "DmHeaderSuppression",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHostKeyAlgorithmsTestConfig = ModelTestConfig{
	Name: "DmHostKeyAlgorithms",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmHostToSSLServerProfileTestConfig = ModelTestConfig{
	Name: "DmHostToSSLServerProfile",
	Model: `{
  host_name_wildmat = "hostname_wildmat"
  ssl_server = "TestAccSSLServerProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"SSLServerProfile": &SSLServerProfileTestConfig,
	},
	TestPre: `
`,
}
var DmInputEncodingTestConfig = ModelTestConfig{
	Name: "DmInputEncoding",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmInsertionAttributesTestConfig = ModelTestConfig{
	Name: "DmInsertionAttributes",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmInvokeErrorTypeTestConfig = ModelTestConfig{
	Name: "DmInvokeErrorType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmJOSEHeaderTestConfig = ModelTestConfig{
	Name: "DmJOSEHeader",
	Model: `{
  header_value = "VALUE"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmJWTClaimsTestConfig = ModelTestConfig{
	Name: "DmJWTClaims",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmJWTGenMethodTestConfig = ModelTestConfig{
	Name: "DmJWTGenMethod",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmJWTValMethodTestConfig = ModelTestConfig{
	Name: "DmJWTValMethod",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmKafkaEndpointTestConfig = ModelTestConfig{
	Name: "DmKafkaEndpoint",
	Model: `{
  host = "localhost"
  port = 8888
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmKafkaPropertyTestConfig = ModelTestConfig{
	Name: "DmKafkaProperty",
	Model: `{
  name = "name"
  value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLBGroupAffinityTestConfig = ModelTestConfig{
	Name: "DmLBGroupAffinity",
	Model: `{
  enable_sa = true
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLBGroupCheckTestConfig = ModelTestConfig{
	Name: "DmLBGroupCheck",
	Model: `{
  active = false
  port = 80
  ssl = "off"
  timeout = 10
  frequency = 180
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLBGroupMemberTestConfig = ModelTestConfig{
	Name: "DmLBGroupMember",
	Model: `{
  server = "lbhost"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLTPATokenVersionTestConfig = ModelTestConfig{
	Name: "DmLTPATokenVersion",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLTPAUserAttributeNameAndValueTestConfig = ModelTestConfig{
	Name: "DmLTPAUserAttributeNameAndValue",
	Model: `{
  ltpa_user_attribute_name = "name"
  ltpa_user_attribute_type = "static"
  ltpa_user_attribute_static_value = "test"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLogEventTestConfig = ModelTestConfig{
	Name: "DmLogEvent",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLogIPFilterTestConfig = ModelTestConfig{
	Name: "DmLogIPFilter",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLogObjectTestConfig = ModelTestConfig{
	Name: "DmLogObject",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmLogTriggerTestConfig = ModelTestConfig{
	Name: "DmLogTrigger",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMCFilterTestConfig = ModelTestConfig{
	Name: "DmMCFilter",
	Model: `{
  filter_name = "mc_filter_name"
  type = "http"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMQHeadersTestConfig = ModelTestConfig{
	Name: "DmMQHeaders",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMSDebugTriggerTypeTestConfig = ModelTestConfig{
	Name: "DmMSDebugTriggerType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMatchRuleTestConfig = ModelTestConfig{
	Name: "DmMatchRule",
	Model: `{
  type = "url"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMetaItemTestConfig = ModelTestConfig{
	Name: "DmMetaItem",
	Model: `{
  meta_category = "all"
  meta_name = "meta_name"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMetadataFromTypeTestConfig = ModelTestConfig{
	Name: "DmMetadataFromType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMtomRuleTestConfig = ModelTestConfig{
	Name: "DmMtomRule",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmMultipartFormDataProfileTestConfig = ModelTestConfig{
	Name: "DmMultipartFormDataProfile",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmNameServerTestConfig = ModelTestConfig{
	Name: "DmNameServer",
	Model: `{
  udp_port = 53
  tcp_port = 53
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmNamedInOutTestConfig = ModelTestConfig{
	Name: "DmNamedInOut",
	Model: `{
  context = "context"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmNamespaceMappingTestConfig = ModelTestConfig{
	Name: "DmNamespaceMapping",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmNetworkRetryTestConfig = ModelTestConfig{
	Name: "DmNetworkRetry",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthAZGrantTypeTestConfig = ModelTestConfig{
	Name: "DmOAuthAZGrantType",
	Model: `{
  saml20bearer = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthComponentsTestConfig = ModelTestConfig{
	Name: "DmOAuthComponents",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthFeaturesTestConfig = ModelTestConfig{
	Name: "DmOAuthFeatures",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthProviderGrantTypeTestConfig = ModelTestConfig{
	Name: "DmOAuthProviderGrantType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthRSSetHeaderTestConfig = ModelTestConfig{
	Name: "DmOAuthRSSetHeader",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthRoleTestConfig = ModelTestConfig{
	Name: "DmOAuthRole",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOAuthTypeTestConfig = ModelTestConfig{
	Name: "DmOAuthType",
	Model: `{
  client_type = "confidential"
  grant_type = "implicit"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmODRConnPropertyTestConfig = ModelTestConfig{
	Name: "DmODRConnProperty",
	Model: `{
  conn_group_prop_name = "propname"
  conn_group_prop_value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmODRConnectorTestConfig = ModelTestConfig{
	Name: "DmODRConnector",
	Model: `{
  dmgr_hostname = "localhost"
  dmgr_port = 8888
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmODRPropertyTestConfig = ModelTestConfig{
	Name: "DmODRProperty",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOIDCHybridResponseTypeTestConfig = ModelTestConfig{
	Name: "DmOIDCHybridResponseType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOpenTelemetryExporterHeaderTestConfig = ModelTestConfig{
	Name: "DmOpenTelemetryExporterHeader",
	Model: `{
  header_value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmOpenTelemetryResourceAttributeTestConfig = ModelTestConfig{
	Name: "DmOpenTelemetryResourceAttribute",
	Model: `{
  value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmPolicyAttachmentPointTestConfig = ModelTestConfig{
	Name: "DmPolicyAttachmentPoint",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmPolicyMapTestConfig = ModelTestConfig{
	Name: "DmPolicyMap",
	Model: `{
  match = "__default-accept-service-providers__"
  rule = "__dp-policy-begin__"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching":        &MatchingTestConfig,
		"StylePolicyRule": &StylePolicyRuleTestConfig,
	},
	TestPre: `
`,
}
var DmPolicyParameterTestConfig = ModelTestConfig{
	Name: "DmPolicyParameter",
	Model: `{
  parameter_name = "PolicyParameterName"
  parameter_value = "PolicyParameterValue"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmPolicySubjectBitmapTestConfig = ModelTestConfig{
	Name: "DmPolicySubjectBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmProfileCPABindingTestConfig = ModelTestConfig{
	Name: "DmProfileCPABinding",
	Model: `{
  internal_partner = "TestAccB2BProfile"
  cpa = "TestAccB2BCPA"
  collaboration = "TestAccB2BCPACollaboration"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"B2BProfile":          &B2BProfileTestConfig,
		"B2BCPA":              &B2BCPATestConfig,
		"B2BCPACollaboration": &B2BCPACollaborationTestConfig,
	},
	TestPre: `
`,
}
var DmProxyPolicyTestConfig = ModelTestConfig{
	Name: "DmProxyPolicy",
	Model: `{
  reg_exp = "*"
  skip = false
  remote_port = 443
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmPubkeyAuthPolicyTestConfig = ModelTestConfig{
	Name: "DmPubkeyAuthPolicy",
	Model: `{
  reg_exp = "*"
  crypto_key = "TestAccCryptoKey"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoKey": &CryptoKeyTestConfig,
	},
	TestPre: `
`,
}
var DmQMFileSystemUsageTestConfig = ModelTestConfig{
	Name: "DmQMFileSystemUsage",
	Model: `{
  warning_threshold = 75
  critical_threshold = 90
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRBMSSHAuthenticateTypeTestConfig = ModelTestConfig{
	Name: "DmRBMSSHAuthenticateType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRadiusServerTestConfig = ModelTestConfig{
	Name: "DmRadiusServer",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRateLimitConfigurationNameValuePairTestConfig = ModelTestConfig{
	Name: "DmRateLimitConfigurationNameValuePair",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRateLimitDefinitionNameValuePairTestConfig = ModelTestConfig{
	Name: "DmRateLimitDefinitionNameValuePair",
	Model: `{
  name = "name"
  value = "value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRateLimitInfoTestConfig = ModelTestConfig{
	Name: "DmRateLimitInfo",
	Model: `{
  name = "ratelimitinfoname"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmRateLimitInfoDomainNamedTestConfig = ModelTestConfig{
	Name: "DmRateLimitInfoDomainNamed",
	Model: `{
  name = "TestAccRateLimitDefinition"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"RateLimitDefinition": &RateLimitDefinitionTestConfig,
	},
	TestPre: `
`,
}
var DmRoutingPrefixTestConfig = ModelTestConfig{
	Name: "DmRoutingPrefix",
	Model: `{
  type = "host"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSAMLAttributeTestConfig = ModelTestConfig{
	Name: "DmSAMLAttribute",
	Model: `{
  source_type = "var"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSAMLAttributeNameAndValueTestConfig = ModelTestConfig{
	Name: "DmSAMLAttributeNameAndValue",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSAMLStatementTypeTestConfig = ModelTestConfig{
	Name: "DmSAMLStatementType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSFTPPolicyTestConfig = ModelTestConfig{
	Name: "DmSFTPPolicy",
	Model: `{
  reg_exp = "*"
  ssh_client_profile = "TestAccSSHClientProfile"
  use_unique_filenames = false
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"SSHClientProfile": &SSHClientProfileTestConfig,
	},
	TestPre: `
`,
}
var DmSFTPServerVirtualDirectoryTestConfig = ModelTestConfig{
	Name: "DmSFTPServerVirtualDirectory",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSLMStatementTestConfig = ModelTestConfig{
	Name: "DmSLMStatement",
	Model: `{
  action = "notify"
  maximum_total_reporting_records = 5000
  maximum_resources_and_credentials_for_threshold = 5000
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"SLMAction": &SLMActionTestConfig,
	},
	TestPre: `
`,
}
var DmSMFlowTestConfig = ModelTestConfig{
	Name: "DmSMFlow",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSMTPOptionsTestConfig = ModelTestConfig{
	Name: "DmSMTPOptions",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSMTPPolicyTestConfig = ModelTestConfig{
	Name: "DmSMTPPolicy",
	Model: `{
  reg_exp = "dpsmtp://*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSOAPHeaderDispositionItemTestConfig = ModelTestConfig{
	Name: "DmSOAPHeaderDispositionItem",
	Model: `{
  action = "processed"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSQLDataSourceConfigNVPairTestConfig = ModelTestConfig{
	Name: "DmSQLDataSourceConfigNVPair",
	Model: `{
  param_name = "param_name"
  param_value = "param_value"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSQLServerTestConfig = ModelTestConfig{
	Name: "DmSQLServer",
	Model: `{
  host = "sql.host"
  port = 1521
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSHUserAuthenticationMethodsTestConfig = ModelTestConfig{
	Name: "DmSSHUserAuthenticationMethods",
	Model: `{
  publickey = true
  password = true
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLClientFeaturesTestConfig = ModelTestConfig{
	Name: "DmSSLClientFeatures",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLFrontSideTestConfig = ModelTestConfig{
	Name: "DmSSLFrontSide",
	Model: `{
  local_address = "0.0.0.0"
  local_port = 8888
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLHostnameValidationFlagsTestConfig = ModelTestConfig{
	Name: "DmSSLHostnameValidationFlags",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLOptionsTestConfig = ModelTestConfig{
	Name: "DmSSLOptions",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLPolicyTestConfig = ModelTestConfig{
	Name: "DmSSLPolicy",
	Model: `{
  reg_exp = "*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSSLProtoVersionsBitmapTestConfig = ModelTestConfig{
	Name: "DmSSLProtoVersionsBitmap",
	Model: `{
  ssl_v3 = false
  tls_v1d0 = false
  tls_v1d1 = true
  tls_v1d2 = true
  tls_v1d3 = true
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmScheduledRuleTestConfig = ModelTestConfig{
	Name: "DmScheduledRule",
	Model: `{
  rule = "__dp-policy-begin__"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"StylePolicyRule": &StylePolicyRuleTestConfig,
	},
	TestPre: `
`,
}
var DmSchemaExceptionRuleTestConfig = ModelTestConfig{
	Name: "DmSchemaExceptionRule",
	Model: `{
  x_path = "*"
  exception_type = "AllowEncrypted"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSchemaValidationTestConfig = ModelTestConfig{
	Name: "DmSchemaValidation",
	Model: `{
  matching = "__default-accept-service-providers__"
  validation_mode = "default"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching": &MatchingTestConfig,
	},
	TestPre: `
`,
}
var DmSearchDomainTestConfig = ModelTestConfig{
	Name: "DmSearchDomain",
	Model: `{
  search_domain = "datapower.com"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSecurityTypeTestConfig = ModelTestConfig{
	Name: "DmSecurityType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSnmpContextTestConfig = ModelTestConfig{
	Name: "DmSnmpContext",
	Model: `{
  domain = "acc_test_domain"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Domain": &DomainTestConfig,
	},
	TestPre: `
`,
}
var DmSnmpCredTestConfig = ModelTestConfig{
	Name: "DmSnmpCred",
	Model: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSnmpCredMaskedTestConfig = ModelTestConfig{
	Name: "DmSnmpCredMasked",
	Model: `{
  engine_id = "0DEADBEEF0"
  auth_protocol = "sha"
  priv_protocol = "des"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSnmpPolicyTestConfig = ModelTestConfig{
	Name: "DmSnmpPolicy",
	Model: `{
  domain = "acc_test_domain"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Domain": &DomainTestConfig,
	},
	TestPre: `
`,
}
var DmSnmpPolicyMQTestConfig = ModelTestConfig{
	Name: "DmSnmpPolicyMQ",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSnmpTargetTestConfig = ModelTestConfig{
	Name: "DmSnmpTarget",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSoapActionPolicyTestConfig = ModelTestConfig{
	Name: "DmSoapActionPolicy",
	Model: `{
  reg_exp = "*"
  soap_action = "*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSourceAS2FeatureTypeTestConfig = ModelTestConfig{
	Name: "DmSourceAS2FeatureType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmSourceHTTPFeatureTypeTestConfig = ModelTestConfig{
	Name: "DmSourceHTTPFeatureType",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmStaticHostTestConfig = ModelTestConfig{
	Name: "DmStaticHost",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmStylesheetParameterTestConfig = ModelTestConfig{
	Name: "DmStylesheetParameter",
	Model: `{
  parameter_value = "PARAMETER-VALUE"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmTAMAZReplicaTestConfig = ModelTestConfig{
	Name: "DmTAMAZReplica",
	Model: `{
  tamaz_replica = "replica"
  tamaz_replica_port = 7136
  tamaz_replica_weight = 10
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmTAMFedDirTestConfig = ModelTestConfig{
	Name: "DmTAMFedDir",
	Model: `{
  fed_name = "fed_name"
  suffix = "suffix"
  host = "ldap.host"
  port = 389
  bind_dn = "dn"
  bind_pw = "TestAccPasswordAlias"
  use_ssl = false
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"PasswordAlias": &PasswordAliasTestConfig,
	},
	TestPre: `
`,
}
var DmTAMRASTraceTestConfig = ModelTestConfig{
	Name: "DmTAMRASTrace",
	Model: `{
  tam_trace_enable = false
  ldap_trace_enable = false
  gs_kit_trace_enable = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmTableEntryTestConfig = ModelTestConfig{
	Name: "DmTableEntry",
	Model: `{
  name = "tableentryname"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmTargetMapRuleTestConfig = ModelTestConfig{
	Name: "DmTargetMapRule",
	Model: `{
  target = "target"
  execute = "default-func-call-global"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"AssemblyActionFunctionCall": &AssemblyActionFunctionCallTestConfig,
	},
	TestPre: `
`,
}
var DmURLMapRuleTestConfig = ModelTestConfig{
	Name: "DmURLMapRule",
	Model: `{
  pattern = "https://www.company.com/XML/stylesheets/*"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmURLRefreshRuleTestConfig = ModelTestConfig{
	Name: "DmURLRefreshRule",
	Model: `{
  url_map = "default-attempt-stream-all"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"URLMap": &URLMapTestConfig,
	},
	TestPre: `
`,
}
var DmURLRewriteRuleTestConfig = ModelTestConfig{
	Name: "DmURLRewriteRule",
	Model: `{
  type = "header-rewrite"
  match_regexp = "Test"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmUploadChunkedPolicyTestConfig = ModelTestConfig{
	Name: "DmUploadChunkedPolicy",
	Model: `{
  reg_exp = "*"
  upload_chunked = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmValidationFeaturesTestConfig = ModelTestConfig{
	Name: "DmValidationFeatures",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmValidationTypeTestConfig = ModelTestConfig{
	Name: "DmValidationType",
	Model: `{
  fixup = "error"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSBaseWSDLTestConfig = ModelTestConfig{
	Name: "DmWSBaseWSDL",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSDLCachePolicyTestConfig = ModelTestConfig{
	Name: "DmWSDLCachePolicy",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSDLUserPolicyTogglesTestConfig = ModelTestConfig{
	Name: "DmWSDLUserPolicyToggles",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointLocalRewriteRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointLocalRewriteRule",
	Model: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointPublishRewriteRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointPublishRewriteRule",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointRemoteRewriteRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointRemoteRewriteRule",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointSubscriptionLocalRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointSubscriptionLocalRule",
	Model: `{
  local_endpoint_hostname = "0.0.0.0"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointSubscriptionPublishRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointSubscriptionPublishRule",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSEndpointSubscriptionRemoteRuleTestConfig = ModelTestConfig{
	Name: "DmWSEndpointSubscriptionRemoteRule",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSMPolicyMapTestConfig = ModelTestConfig{
	Name: "DmWSMPolicyMap",
	Model: `{
  wsdl_component_type = "all"
  match = "__default-accept-service-providers__"
  rule = "TestAccWSStylePolicyRule"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching":          &MatchingTestConfig,
		"WSStylePolicyRule": &WSStylePolicyRuleTestConfig,
	},
	TestPre: `
`,
}
var DmWSOperationConformancePolicyTestConfig = ModelTestConfig{
	Name: "DmWSOperationConformancePolicy",
	Model: `{
  conformance_policy = "TestAccConformancePolicy"
  conformance_policy_wsdl_component_type = "all"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"ConformancePolicy": &ConformancePolicyTestConfig,
	},
	TestPre: `
`,
}
var DmWSOperationPolicySubjectOptOutTestConfig = ModelTestConfig{
	Name: "DmWSOperationPolicySubjectOptOut",
	Model: `{
  policy_subject_opt_out_wsdl_component_type = "all"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSOperationReliableMessagingTestConfig = ModelTestConfig{
	Name: "DmWSOperationReliableMessaging",
	Model: `{
  reliable_messaging_wsdl_component_type = "all"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSOperationSchedulerPriorityTestConfig = ModelTestConfig{
	Name: "DmWSOperationSchedulerPriority",
	Model: `{
  scheduler_priority_wsdl_component_type = "all"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSPolicyParametersTestConfig = ModelTestConfig{
	Name: "DmWSPolicyParameters",
	Model: `{
  policy_param_parameters = "TestAccPolicyParameters"
  policy_param_wsdl_component_type = "all"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"PolicyParameters": &PolicyParametersTestConfig,
	},
	TestPre: `
`,
}
var DmWSRMPolicyBitmapTestConfig = ModelTestConfig{
	Name: "DmWSRMPolicyBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSRRSavedSearchWSDLSourceTestConfig = ModelTestConfig{
	Name: "DmWSRRSavedSearchWSDLSource",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSRRWSDLSourceTestConfig = ModelTestConfig{
	Name: "DmWSRRWSDLSource",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSSLMOpsTestConfig = ModelTestConfig{
	Name: "DmWSSLMOps",
	Model: `{
  operation = "all"
  target = "front"
  severity = "low"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWSUserTogglesTestConfig = ModelTestConfig{
	Name: "DmWSUserToggles",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWebAppRequestPolicyMapTestConfig = ModelTestConfig{
	Name: "DmWebAppRequestPolicyMap",
	Model: `{
  match = "__default-accept-service-providers__"
  rule = "TestAccRequestProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching":      &MatchingTestConfig,
		"WebAppRequest": &WebAppRequestTestConfig,
	},
	TestPre: `
`,
}
var DmWebAppResponsePolicyMapTestConfig = ModelTestConfig{
	Name: "DmWebAppResponsePolicyMap",
	Model: `{
  match = "__default-accept-service-providers__"
  rule = "TestAccResponseProfile"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching":       &MatchingTestConfig,
		"WebAppResponse": &WebAppResponseTestConfig,
	},
	TestPre: `
`,
}
var DmWebGWErrorPolicyMapTestConfig = ModelTestConfig{
	Name: "DmWebGWErrorPolicyMap",
	Model: `{
  match = "__default-accept-service-providers__"
  action = "TestAccMPGWErrorAction"
}`,
	ModelOnly: true,
	Dependencies: map[string]*ModelTestConfig{
		"Matching":        &MatchingTestConfig,
		"MPGWErrorAction": &MPGWErrorActionTestConfig,
	},
	TestPre: `
`,
}
var DmWebGWErrorRespHeaderInjectionTestConfig = ModelTestConfig{
	Name: "DmWebGWErrorRespHeaderInjection",
	Model: `{
  header_tag_value = "HEADER_VALUE"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWebSphereJMSEndpointTestConfig = ModelTestConfig{
	Name: "DmWebSphereJMSEndpoint",
	Model: `{
  host = "localhost"
  port = 8888
  transport = "TCP"
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmWeekdayBitmapTestConfig = ModelTestConfig{
	Name: "DmWeekdayBitmap",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmXMLMgmtModesTestConfig = ModelTestConfig{
	Name: "DmXMLMgmtModes",
	Model: `{
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DmXPathRoutingRuleTestConfig = ModelTestConfig{
	Name: "DmXPathRoutingRule",
	Model: `{
  x_path = "*"
  host = "localhost"
  port = 8888
  ssl = false
}`,
	ModelOnly:    true,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DocumentCryptoMapTestConfig = ModelTestConfig{
	Name: "DocumentCryptoMap",
	Resource: `
resource "datapower_documentcryptomap" "test" {
  id = "DocumentCryptoMap_name"
  app_domain = "acc_test_domain"
  operation = "encrypt"
  x_path = ["*",]
}`,
	Data: `
data "datapower_documentcryptomap" "test" {
  depends_on = [ datapower_documentcryptomap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DomainTestConfig = ModelTestConfig{
	Name: "Domain",
	Resource: `
resource "datapower_domain" "test" {
  app_domain = "domain_resource_test"
}`,
	Data: `
data "datapower_domain" "test" {
  depends_on = [ datapower_domain.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DomainAvailabilityTestConfig = ModelTestConfig{
	Name: "DomainAvailability",
	Resource: `
resource "datapower_domainavailability" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_domainavailability" "test" {
  depends_on = [ datapower_domainavailability.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DomainSettingsTestConfig = ModelTestConfig{
	Name: "DomainSettings",
	Resource: `
resource "datapower_domainsettings" "test" {
  app_domain = "acc_test_domain"
  password_treatment = "none"
}`,
	Data: `
data "datapower_domainsettings" "test" {
  depends_on = [ datapower_domainsettings.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var DurationMonitorTestConfig = ModelTestConfig{
	Name: "DurationMonitor",
	Resource: `
resource "datapower_durationmonitor" "test" {
  id = "DurationMonitor_name"
  app_domain = "acc_test_domain"
  measure = "messages"
  message_type = "TestAccMessageType"
}`,
	Data: `
data "datapower_durationmonitor" "test" {
  depends_on = [ datapower_durationmonitor.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"MessageType": &MessageTypeTestConfig,
	},
	TestPre: `
`,
}
var EBMS2SourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "EBMS2SourceProtocolHandler",
	Resource: `
resource "datapower_ebms2sourceprotocolhandler" "test" {
  id = "EBMS2SourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
}`,
	Data: `
data "datapower_ebms2sourceprotocolhandler" "test" {
  depends_on = [ datapower_ebms2sourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var EBMS3SourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "EBMS3SourceProtocolHandler",
	Resource: `
resource "datapower_ebms3sourceprotocolhandler" "test" {
  id = "EBMS3SourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
}`,
	Data: `
data "datapower_ebms3sourceprotocolhandler" "test" {
  depends_on = [ datapower_ebms3sourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ErrorReportSettingsTestConfig = ModelTestConfig{
	Name: "ErrorReportSettings",
	Resource: `
resource "datapower_errorreportsettings" "test" {
}`,
	Data: `
data "datapower_errorreportsettings" "test" {
  depends_on = [ datapower_errorreportsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var FTPFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "FTPFilePollerSourceProtocolHandler",
	Resource: `
resource "datapower_ftpfilepollersourceprotocolhandler" "test" {
  id = "FTPFilePollerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  target_directory = "ftp://user:password@host:port/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
	Data: `
data "datapower_ftpfilepollersourceprotocolhandler" "test" {
  depends_on = [ datapower_ftpfilepollersourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var FTPQuoteCommandsTestConfig = ModelTestConfig{
	Name: "FTPQuoteCommands",
	Resource: `
resource "datapower_ftpquotecommands" "test" {
  id = "FTPQuoteCommands_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_ftpquotecommands" "test" {
  depends_on = [ datapower_ftpquotecommands.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var FTPServerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "FTPServerSourceProtocolHandler",
	Resource: `
resource "datapower_ftpserversourceprotocolhandler" "test" {
  id = "FTPServerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 21
}`,
	Data: `
data "datapower_ftpserversourceprotocolhandler" "test" {
  depends_on = [ datapower_ftpserversourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var FileSystemUsageMonitorTestConfig = ModelTestConfig{
	Name: "FileSystemUsageMonitor",
	Resource: `
resource "datapower_filesystemusagemonitor" "test" {
  polling_interval = 60
  all_system = true
}`,
	Data: `
data "datapower_filesystemusagemonitor" "test" {
  depends_on = [ datapower_filesystemusagemonitor.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var FilterActionTestConfig = ModelTestConfig{
	Name: "FilterAction",
	Resource: `
resource "datapower_filteraction" "test" {
  id = "FilterAction_name"
  app_domain = "acc_test_domain"
  type = "notify"
}`,
	Data: `
data "datapower_filteraction" "test" {
  depends_on = [ datapower_filteraction.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var FormsLoginPolicyTestConfig = ModelTestConfig{
	Name: "FormsLoginPolicy",
	Resource: `
resource "datapower_formsloginpolicy" "test" {
  id = "FormsLoginPolicy_name"
  app_domain = "acc_test_domain"
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
data "datapower_formsloginpolicy" "test" {
  depends_on = [ datapower_formsloginpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GWSRemoteDebugTestConfig = ModelTestConfig{
	Name: "GWSRemoteDebug",
	Resource: `
resource "datapower_gwsremotedebug" "test" {
  local_port = 9229
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_gwsremotedebug" "test" {
  depends_on = [ datapower_gwsremotedebug.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GWScriptSettingsTestConfig = ModelTestConfig{
	Name: "GWScriptSettings",
	Resource: `
resource "datapower_gwscriptsettings" "test" {
}`,
	Data: `
data "datapower_gwscriptsettings" "test" {
  depends_on = [ datapower_gwscriptsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GatewayPeeringTestConfig = ModelTestConfig{
	Name: "GatewayPeering",
	Resource: `
resource "datapower_gatewaypeering" "test" {
  id = "___GatewayPeering_name"
  app_domain = "acc_test_domain"
  local_port = 16380
}`,
	Data: `
data "datapower_gatewaypeering" "test" {
  depends_on = [ datapower_gatewaypeering.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GatewayPeeringGroupTestConfig = ModelTestConfig{
	Name: "GatewayPeeringGroup",
	Resource: `
resource "datapower_gatewaypeeringgroup" "test" {
  id = "GatewayPeeringGroup_name"
  app_domain = "acc_test_domain"
  mode = "peer"
  enable_ssl = true
}`,
	Data: `
data "datapower_gatewaypeeringgroup" "test" {
  depends_on = [ datapower_gatewaypeeringgroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GatewayPeeringManagerTestConfig = ModelTestConfig{
	Name: "GatewayPeeringManager",
	Resource: `
resource "datapower_gatewaypeeringmanager" "test" {
  app_domain = "acc_test_domain"
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit = "default-gateway-peering"
  subscription = "default-gateway-peering"
  ratelimit_module = "default-gateway-peering"
}`,
	Data: `
data "datapower_gatewaypeeringmanager" "test" {
  depends_on = [ datapower_gatewaypeeringmanager.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"GatewayPeering": &GatewayPeeringTestConfig,
	},
	TestPre: `
`,
}
var GitOpsTestConfig = ModelTestConfig{
	Name: "GitOps",
	Resource: `
resource "datapower_gitops" "test" {
  app_domain = "acc_test_domain"
  connection_type = "https"
  mode = "read-write"
  commit_identifier_type = "branch"
  commit_identifier = "main"
  remote_location = "https://github.com/ScottW514/terraform-provider-datapower"
}`,
	Data: `
data "datapower_gitops" "test" {
  depends_on = [ datapower_gitops.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GitOpsTemplateTestConfig = ModelTestConfig{
	Name: "GitOpsTemplate",
	Resource: `
resource "datapower_gitopstemplate" "test" {
  id = "GitOpsTemplate_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_gitopstemplate" "test" {
  depends_on = [ datapower_gitopstemplate.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GitOpsVariablesTestConfig = ModelTestConfig{
	Name: "GitOpsVariables",
	Resource: `
resource "datapower_gitopsvariables" "test" {
}`,
	Data: `
data "datapower_gitopsvariables" "test" {
  depends_on = [ datapower_gitopsvariables.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var GraphQLSchemaOptionsTestConfig = ModelTestConfig{
	Name: "GraphQLSchemaOptions",
	Resource: `
resource "datapower_graphqlschemaoptions" "test" {
  id = "GraphQLSchemaOptions_name"
  app_domain = "acc_test_domain"
  api = "TestAccAPIDefinition"
}`,
	Data: `
data "datapower_graphqlschemaoptions" "test" {
  depends_on = [ datapower_graphqlschemaoptions.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APIDefinition": &APIDefinitionTestConfig,
	},
	TestPre: `
`,
}
var HTTPInputConversionMapTestConfig = ModelTestConfig{
	Name: "HTTPInputConversionMap",
	Resource: `
resource "datapower_httpinputconversionmap" "test" {
  id = "___HTTPInputConversionMap_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_httpinputconversionmap" "test" {
  depends_on = [ datapower_httpinputconversionmap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var HTTPSSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "HTTPSSourceProtocolHandler",
	Resource: `
resource "datapower_httpssourceprotocolhandler" "test" {
  id = "HTTPSSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  ssl_server_config_type = "server"
}`,
	Data: `
data "datapower_httpssourceprotocolhandler" "test" {
  depends_on = [ datapower_httpssourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var HTTPSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "HTTPSourceProtocolHandler",
	Resource: `
resource "datapower_httpsourceprotocolhandler" "test" {
  id = "HTTPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 80
}`,
	Data: `
data "datapower_httpsourceprotocolhandler" "test" {
  depends_on = [ datapower_httpsourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var HTTPUserAgentTestConfig = ModelTestConfig{
	Name: "HTTPUserAgent",
	Resource: `
resource "datapower_httpuseragent" "test" {
  id = "___HTTPUserAgent_test"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_httpuseragent" "test" {
  depends_on = [ datapower_httpuseragent.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var HostAliasTestConfig = ModelTestConfig{
	Name: "HostAlias",
	Resource: `
resource "datapower_hostalias" "test" {
  id = "_HostAlias_name"
  ip_address = "10.10.10.10"
}`,
	Data: `
data "datapower_hostalias" "test" {
  depends_on = [ datapower_hostalias.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ImportPackageTestConfig = ModelTestConfig{
	Name: "ImportPackage",
	Resource: `
resource "datapower_importpackage" "test" {
  id = "ImportPackage_name"
  app_domain = "acc_test_domain"
  url = "http://localhost/config.zip"
  import_format = "ZIP"
  overwrite_files = true
  overwrite_objects = true
}`,
	Data: `
data "datapower_importpackage" "test" {
  depends_on = [ datapower_importpackage.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var IncludeConfigTestConfig = ModelTestConfig{
	Name: "IncludeConfig",
	Resource: `
resource "datapower_includeconfig" "test" {
  id = "IncludeConfig_name"
  app_domain = "acc_test_domain"
  url = "http://localhost/config.zip"
}`,
	Data: `
data "datapower_includeconfig" "test" {
  depends_on = [ datapower_includeconfig.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var InteropServiceTestConfig = ModelTestConfig{
	Name: "InteropService",
	Resource: `
resource "datapower_interopservice" "test" {
}`,
	Data: `
data "datapower_interopservice" "test" {
  depends_on = [ datapower_interopservice.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var JOSERecipientIdentifierTestConfig = ModelTestConfig{
	Name: "JOSERecipientIdentifier",
	Resource: `
resource "datapower_joserecipientidentifier" "test" {
  id = "JOSERecipientIdentifier_name"
  app_domain = "acc_test_domain"
  type = "key"
  header_param = ` + DmJOSEHeaderTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_joserecipientidentifier" "test" {
  depends_on = [ datapower_joserecipientidentifier.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmJOSEHeader": &DmJOSEHeaderTestConfig,
	},
	TestPre: `
`,
}
var JOSESignatureIdentifierTestConfig = ModelTestConfig{
	Name: "JOSESignatureIdentifier",
	Resource: `
resource "datapower_josesignatureidentifier" "test" {
  id = "JOSESignatureIdentifier_name"
  app_domain = "acc_test_domain"
  type = "certificate"
  header_param = ` + DmJOSEHeaderTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_josesignatureidentifier" "test" {
  depends_on = [ datapower_josesignatureidentifier.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmJOSEHeader": &DmJOSEHeaderTestConfig,
	},
	TestPre: `
`,
}
var JSONSettingsTestConfig = ModelTestConfig{
	Name: "JSONSettings",
	Resource: `
resource "datapower_jsonsettings" "test" {
  id = "JSONSettings_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_jsonsettings" "test" {
  depends_on = [ datapower_jsonsettings.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var JWEHeaderTestConfig = ModelTestConfig{
	Name: "JWEHeader",
	Resource: `
resource "datapower_jweheader" "test" {
  id = "JWEHeader_name"
  app_domain = "acc_test_domain"
  recipient = "TestAccJWERecipient"
}`,
	Data: `
data "datapower_jweheader" "test" {
  depends_on = [ datapower_jweheader.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"JWERecipient": &JWERecipientTestConfig,
	},
	TestPre: `
`,
}
var JWERecipientTestConfig = ModelTestConfig{
	Name: "JWERecipient",
	Resource: `
resource "datapower_jwerecipient" "test" {
  id = "JWERecipient_name"
  app_domain = "acc_test_domain"
  algorithm = "RSA1_5"
}`,
	Data: `
data "datapower_jwerecipient" "test" {
  depends_on = [ datapower_jwerecipient.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var JWSSignatureTestConfig = ModelTestConfig{
	Name: "JWSSignature",
	Resource: `
resource "datapower_jwssignature" "test" {
  id = "JWSSignature_name"
  app_domain = "acc_test_domain"
  algorithm = "RS256"
  key = "TestAccCryptoKey"
}`,
	Data: `
data "datapower_jwssignature" "test" {
  depends_on = [ datapower_jwssignature.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoKey": &CryptoKeyTestConfig,
	},
	TestPre: `
`,
}
var KafkaClusterTestConfig = ModelTestConfig{
	Name: "KafkaCluster",
	Resource: `
resource "datapower_kafkacluster" "test" {
  id = "KafkaCluster_name"
  app_domain = "acc_test_domain"
  protocol = "plaintext"
  endpoint = ` + DmKafkaEndpointTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_kafkacluster" "test" {
  depends_on = [ datapower_kafkacluster.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmKafkaEndpoint": &DmKafkaEndpointTestConfig,
	},
	TestPre: `
`,
}
var KafkaSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "KafkaSourceProtocolHandler",
	Resource: `
resource "datapower_kafkasourceprotocolhandler" "test" {
  id = "KafkaSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  cluster = "TestAccKafkaCluster"
  request_topic = "topic"
  consumer_group = "consumer"
}`,
	Data: `
data "datapower_kafkasourceprotocolhandler" "test" {
  depends_on = [ datapower_kafkasourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"KafkaCluster": &KafkaClusterTestConfig,
	},
	TestPre: `
`,
}
var LDAPConnectionPoolTestConfig = ModelTestConfig{
	Name: "LDAPConnectionPool",
	Resource: `
resource "datapower_ldapconnectionpool" "test" {
  id = "LDAPConnectionPool_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_ldapconnectionpool" "test" {
  depends_on = [ datapower_ldapconnectionpool.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LDAPSearchParametersTestConfig = ModelTestConfig{
	Name: "LDAPSearchParameters",
	Resource: `
resource "datapower_ldapsearchparameters" "test" {
  id = "LDAPSearchParameters_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_ldapsearchparameters" "test" {
  depends_on = [ datapower_ldapsearchparameters.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LoadBalancerGroupTestConfig = ModelTestConfig{
	Name: "LoadBalancerGroup",
	Resource: `
resource "datapower_loadbalancergroup" "test" {
  id = "LoadBalancerGroup_name"
  app_domain = "acc_test_domain"
  algorithm = "round-robin"
  retrieve_info = false
  damp = 120
}`,
	Data: `
data "datapower_loadbalancergroup" "test" {
  depends_on = [ datapower_loadbalancergroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LogLabelTestConfig = ModelTestConfig{
	Name: "LogLabel",
	Resource: `
resource "datapower_loglabel" "test" {
  id = "___LogLabel_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_loglabel" "test" {
  depends_on = [ datapower_loglabel.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LogTargetTestConfig = ModelTestConfig{
	Name: "LogTarget",
	Resource: `
resource "datapower_logtarget" "test" {
  id = "___LogTarget_name"
  app_domain = "acc_test_domain"
  type = "file"
}`,
	Data: `
data "datapower_logtarget" "test" {
  depends_on = [ datapower_logtarget.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LunaTestConfig = ModelTestConfig{
	Name: "Luna",
	Resource: `
resource "datapower_luna" "test" {
  id = "Luna_name"
  remote_address = "localhost"
  server_cert = "cert:///cert.crt"
  security_option = "none"
}`,
	Data: `
data "datapower_luna" "test" {
  depends_on = [ datapower_luna.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LunaHAGroupTestConfig = ModelTestConfig{
	Name: "LunaHAGroup",
	Resource: `
resource "datapower_lunahagroup" "test" {
  id = "LunaHAGroup_name"
  app_domain = "acc_test_domain"
  group_name = "groupname"
  member = ["TestAccLunaPartition"]
}`,
	Data: `
data "datapower_lunahagroup" "test" {
  depends_on = [ datapower_lunahagroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"LunaPartition": &LunaPartitionTestConfig,
	},
	TestPre: `
`,
}
var LunaHASettingsTestConfig = ModelTestConfig{
	Name: "LunaHASettings",
	Resource: `
resource "datapower_lunahasettings" "test" {
}`,
	Data: `
data "datapower_lunahasettings" "test" {
  depends_on = [ datapower_lunahasettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var LunaPartitionTestConfig = ModelTestConfig{
	Name: "LunaPartition",
	Resource: `
resource "datapower_lunapartition" "test" {
  id = "LunaPartition_name"
  app_domain = "acc_test_domain"
  partition_name = "partitionname"
  partition_serial = "serial"
  password_alias = "TestAccPasswordAlias"
  login_role = "co"
}`,
	Data: `
data "datapower_lunapartition" "test" {
  depends_on = [ datapower_lunapartition.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"PasswordAlias": &PasswordAliasTestConfig,
	},
	TestPre: `
`,
}
var MCFCustomRuleTestConfig = ModelTestConfig{
	Name: "MCFCustomRule",
	Resource: `
resource "datapower_mcfcustomrule" "test" {
  id = "MCFCustomRule_name"
  app_domain = "acc_test_domain"
  custom_rule_name = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}`,
	Data: `
data "datapower_mcfcustomrule" "test" {
  depends_on = [ datapower_mcfcustomrule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"StylePolicyRule": &StylePolicyRuleTestConfig,
	},
	TestPre: `
`,
}
var MCFHttpHeaderTestConfig = ModelTestConfig{
	Name: "MCFHttpHeader",
	Resource: `
resource "datapower_mcfhttpheader" "test" {
  id = "MCFHttpHeader_name"
  app_domain = "acc_test_domain"
  http_name = "HEADERNAME"
  http_value = "HEADERVALUE"
}`,
	Data: `
data "datapower_mcfhttpheader" "test" {
  depends_on = [ datapower_mcfhttpheader.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MCFHttpMethodTestConfig = ModelTestConfig{
	Name: "MCFHttpMethod",
	Resource: `
resource "datapower_mcfhttpmethod" "test" {
  id = "MCFHttpMethod_name"
  app_domain = "acc_test_domain"
  http_method = "GET"
}`,
	Data: `
data "datapower_mcfhttpmethod" "test" {
  depends_on = [ datapower_mcfhttpmethod.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MCFHttpURLTestConfig = ModelTestConfig{
	Name: "MCFHttpURL",
	Resource: `
resource "datapower_mcfhttpurl" "test" {
  id = "MCFHttpURL_name"
  app_domain = "acc_test_domain"
  http_url_expression = "*"
}`,
	Data: `
data "datapower_mcfhttpurl" "test" {
  depends_on = [ datapower_mcfhttpurl.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MCFXPathTestConfig = ModelTestConfig{
	Name: "MCFXPath",
	Resource: `
resource "datapower_mcfxpath" "test" {
  id = "MCFXPath_name"
  app_domain = "acc_test_domain"
  x_path_expression = "*"
  x_path_value = "value"
}`,
	Data: `
data "datapower_mcfxpath" "test" {
  depends_on = [ datapower_mcfxpath.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MPGWErrorActionTestConfig = ModelTestConfig{
	Name: "MPGWErrorAction",
	Resource: `
resource "datapower_mpgwerroraction" "test" {
  id = "MPGWErrorAction_name"
  app_domain = "acc_test_domain"
  remote_url = "http://google.com"
}`,
	Data: `
data "datapower_mpgwerroraction" "test" {
  depends_on = [ datapower_mpgwerroraction.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MPGWErrorHandlingPolicyTestConfig = ModelTestConfig{
	Name: "MPGWErrorHandlingPolicy",
	Resource: `
resource "datapower_mpgwerrorhandlingpolicy" "test" {
  id = "MPGWErrorHandlingPolicy_name"
  app_domain = "acc_test_domain"
  policy_maps = ` + DmWebGWErrorPolicyMapTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_mpgwerrorhandlingpolicy" "test" {
  depends_on = [ datapower_mpgwerrorhandlingpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmWebGWErrorPolicyMap": &DmWebGWErrorPolicyMapTestConfig,
	},
	TestPre: `
`,
}
var MQManagerTestConfig = ModelTestConfig{
	Name: "MQManager",
	Resource: `
resource "datapower_mqmanager" "test" {
  id = "MQManager_name"
  app_domain = "acc_test_domain"
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}`,
	Data: `
data "datapower_mqmanager" "test" {
  depends_on = [ datapower_mqmanager.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var MQManagerGroupTestConfig = ModelTestConfig{
	Name: "MQManagerGroup",
	Resource: `
resource "datapower_mqmanagergroup" "test" {
  id = "MQManagerGroup_name"
  app_domain = "acc_test_domain"
  primary_queue_manager = "TestAccMQManager"
}`,
	Data: `
data "datapower_mqmanagergroup" "test" {
  depends_on = [ datapower_mqmanagergroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"MQManager": &MQManagerTestConfig,
	},
	TestPre: `
`,
}
var MQv9PlusMFTSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "MQv9PlusMFTSourceProtocolHandler",
	Resource: `
resource "datapower_mqv9plusmftsourceprotocolhandler" "test" {
  id = "MQv9PlusMFTSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  queue_manager = datapower_mqmanager.test.id
  get_queue = "queue"
}`,
	Data: `
data "datapower_mqv9plusmftsourceprotocolhandler" "test" {
  depends_on = [ datapower_mqv9plusmftsourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `resource "datapower_mqmanager" "test" {
  id = "MQManager_name"
  app_domain = "acc_test_domain"
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}

`,
}
var MQv9PlusSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "MQv9PlusSourceProtocolHandler",
	Resource: `
resource "datapower_mqv9plussourceprotocolhandler" "test" {
  id = "MQv9PlusSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  queue_manager = datapower_mqmanager.test.id
}`,
	Data: `
data "datapower_mqv9plussourceprotocolhandler" "test" {
  depends_on = [ datapower_mqv9plussourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `resource "datapower_mqmanager" "test" {
  id = "MQManager_name"
  app_domain = "acc_test_domain"
  host_name = "localhost"
  cache_timeout = 60
  xml_manager = "default"
}

`,
}
var MTOMPolicyTestConfig = ModelTestConfig{
	Name: "MTOMPolicy",
	Resource: `
resource "datapower_mtompolicy" "test" {
  id = "MTOMPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_mtompolicy" "test" {
  depends_on = [ datapower_mtompolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MatchingTestConfig = ModelTestConfig{
	Name: "Matching",
	Resource: `
resource "datapower_matching" "test" {
  id = "___Matching_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_matching" "test" {
  depends_on = [ datapower_matching.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MessageContentFiltersTestConfig = ModelTestConfig{
	Name: "MessageContentFilters",
	Resource: `
resource "datapower_messagecontentfilters" "test" {
  id = "MessageContentFilters_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_messagecontentfilters" "test" {
  depends_on = [ datapower_messagecontentfilters.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MessageMatchingTestConfig = ModelTestConfig{
	Name: "MessageMatching",
	Resource: `
resource "datapower_messagematching" "test" {
  id = "MessageMatching_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_messagematching" "test" {
  depends_on = [ datapower_messagematching.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MessageTypeTestConfig = ModelTestConfig{
	Name: "MessageType",
	Resource: `
resource "datapower_messagetype" "test" {
  id = "MessageType_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_messagetype" "test" {
  depends_on = [ datapower_messagetype.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MgmtInterfaceTestConfig = ModelTestConfig{
	Name: "MgmtInterface",
	Resource: `
resource "datapower_mgmtinterface" "test" {
  local_port = 5550
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_mgmtinterface" "test" {
  depends_on = [ datapower_mgmtinterface.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var MultiProtocolGatewayTestConfig = ModelTestConfig{
	Name: "MultiProtocolGateway",
	Resource: `
resource "datapower_multiprotocolgateway" "test" {
  id = "MultiProtocolGateway_name"
  app_domain = "acc_test_domain"
  type = "static-backend"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
	Data: `
data "datapower_multiprotocolgateway" "test" {
  depends_on = [ datapower_multiprotocolgateway.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var NFSClientSettingsTestConfig = ModelTestConfig{
	Name: "NFSClientSettings",
	Resource: `
resource "datapower_nfsclientsettings" "test" {
}`,
	Data: `
data "datapower_nfsclientsettings" "test" {
  depends_on = [ datapower_nfsclientsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var NFSDynamicMountsTestConfig = ModelTestConfig{
	Name: "NFSDynamicMounts",
	Resource: `
resource "datapower_nfsdynamicmounts" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_nfsdynamicmounts" "test" {
  depends_on = [ datapower_nfsdynamicmounts.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var NFSFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "NFSFilePollerSourceProtocolHandler",
	Resource: `
resource "datapower_nfsfilepollersourceprotocolhandler" "test" {
  id = "NFSFilePollerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  target_directory = "dpnfs://static-mount-name/path/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  generate_result_file = false
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
	Data: `
data "datapower_nfsfilepollersourceprotocolhandler" "test" {
  depends_on = [ datapower_nfsfilepollersourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var NFSStaticMountTestConfig = ModelTestConfig{
	Name: "NFSStaticMount",
	Resource: `
resource "datapower_nfsstaticmount" "test" {
  id = "NFSStaticMount_name"
  app_domain = "acc_test_domain"
  remote = "url://test"
}`,
	Data: `
data "datapower_nfsstaticmount" "test" {
  depends_on = [ datapower_nfsstaticmount.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var NameValueProfileTestConfig = ModelTestConfig{
	Name: "NameValueProfile",
	Resource: `
resource "datapower_namevalueprofile" "test" {
  id = "NameValueProfile_name"
  app_domain = "acc_test_domain"
  default_fixup = "strip"
}`,
	Data: `
data "datapower_namevalueprofile" "test" {
  depends_on = [ datapower_namevalueprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var OAuthProviderSettingsTestConfig = ModelTestConfig{
	Name: "OAuthProviderSettings",
	Resource: `
resource "datapower_oauthprovidersettings" "test" {
  id = "OAuthProviderSettings_name"
  app_domain = "acc_test_domain"
  provider_type = "native"
}`,
	Data: `
data "datapower_oauthprovidersettings" "test" {
  depends_on = [ datapower_oauthprovidersettings.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var OAuthSupportedClientTestConfig = ModelTestConfig{
	Name: "OAuthSupportedClient",
	Resource: `
resource "datapower_oauthsupportedclient" "test" {
  id = "OAuthSupportedClient_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_oauthsupportedclient" "test" {
  depends_on = [ datapower_oauthsupportedclient.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var OAuthSupportedClientGroupTestConfig = ModelTestConfig{
	Name: "OAuthSupportedClientGroup",
	Resource: `
resource "datapower_oauthsupportedclientgroup" "test" {
  id = "OAuthSupportedClientGroup_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_oauthsupportedclientgroup" "test" {
  depends_on = [ datapower_oauthsupportedclientgroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ODRTestConfig = ModelTestConfig{
	Name: "ODR",
	Resource: `
resource "datapower_odr" "test" {
  odr_server_name = "dp_set"
}`,
	Data: `
data "datapower_odr" "test" {
  depends_on = [ datapower_odr.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ODRConnectorGroupTestConfig = ModelTestConfig{
	Name: "ODRConnectorGroup",
	Resource: `
resource "datapower_odrconnectorgroup" "test" {
  id = "ODRConnectorGroup_name"
  xml_manager = "default"
}`,
	Data: `
data "datapower_odrconnectorgroup" "test" {
  depends_on = [ datapower_odrconnectorgroup.test ]
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var OpenTelemetryTestConfig = ModelTestConfig{
	Name: "OpenTelemetry",
	Resource: `
resource "datapower_opentelemetry" "test" {
  id = "OpenTelemetry_name"
  app_domain = "acc_test_domain"
  exporter = "TestAccOpenTelemetryExporter"
  sampler = "TestAccOpenTelemetrySampler"
}`,
	Data: `
data "datapower_opentelemetry" "test" {
  depends_on = [ datapower_opentelemetry.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"OpenTelemetryExporter": &OpenTelemetryExporterTestConfig,
		"OpenTelemetrySampler":  &OpenTelemetrySamplerTestConfig,
	},
	TestPre: `
`,
}
var OpenTelemetryExporterTestConfig = ModelTestConfig{
	Name: "OpenTelemetryExporter",
	Resource: `
resource "datapower_opentelemetryexporter" "test" {
  id = "OpenTelemetryExporter_name"
  app_domain = "acc_test_domain"
  type = "http"
  host_name = "localhost"
  port = 4318
  processor = "batch"
}`,
	Data: `
data "datapower_opentelemetryexporter" "test" {
  depends_on = [ datapower_opentelemetryexporter.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var OpenTelemetrySamplerTestConfig = ModelTestConfig{
	Name: "OpenTelemetrySampler",
	Resource: `
resource "datapower_opentelemetrysampler" "test" {
  id = "OpenTelemetrySampler_name"
  app_domain = "acc_test_domain"
  parent_based = true
  type = "always-on"
}`,
	Data: `
data "datapower_opentelemetrysampler" "test" {
  depends_on = [ datapower_opentelemetrysampler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var OperationRateLimitTestConfig = ModelTestConfig{
	Name: "OperationRateLimit",
	Resource: `
resource "datapower_operationratelimit" "test" {
  id = "OperationRateLimit_name"
  app_domain = "acc_test_domain"
  operation = "TestAccAPIOperation"
}`,
	Data: `
data "datapower_operationratelimit" "test" {
  depends_on = [ datapower_operationratelimit.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"APIOperation": &APIOperationTestConfig,
	},
	TestPre: `
`,
}
var POPPollerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "POPPollerSourceProtocolHandler",
	Resource: `
resource "datapower_poppollersourceprotocolhandler" "test" {
  id = "POPPollerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  mail_server = "localhost"
  port = 8888
  conn_security = "none"
  account = "account"
  delay_between_polls = 300
  max_messages_per_poll = 5
}`,
	Data: `
data "datapower_poppollersourceprotocolhandler" "test" {
  depends_on = [ datapower_poppollersourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ParseSettingsTestConfig = ModelTestConfig{
	Name: "ParseSettings",
	Resource: `
resource "datapower_parsesettings" "test" {
  id = "___ParseSettings_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_parsesettings" "test" {
  depends_on = [ datapower_parsesettings.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var PasswordAliasTestConfig = ModelTestConfig{
	Name: "PasswordAlias",
	Resource: `
resource "datapower_passwordalias" "test" {
  id = "PasswordAlias_name"
  app_domain = "acc_test_domain"
  password = "password"
}`,
	Data: `
data "datapower_passwordalias" "test" {
  depends_on = [ datapower_passwordalias.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var PeerGroupTestConfig = ModelTestConfig{
	Name: "PeerGroup",
	Resource: `
resource "datapower_peergroup" "test" {
  id = "PeerGroup_name"
  app_domain = "acc_test_domain"
  type = "slm"
}`,
	Data: `
data "datapower_peergroup" "test" {
  depends_on = [ datapower_peergroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var PolicyAttachmentsTestConfig = ModelTestConfig{
	Name: "PolicyAttachments",
	Resource: `
resource "datapower_policyattachments" "test" {
  id = "PolicyAttachments_name"
  app_domain = "acc_test_domain"
  enforcement_mode = "enforce"
  policy_references = false
  sla_enforcement_mode = "allow-if-no-sla"
}`,
	Data: `
data "datapower_policyattachments" "test" {
  depends_on = [ datapower_policyattachments.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var PolicyParametersTestConfig = ModelTestConfig{
	Name: "PolicyParameters",
	Resource: `
resource "datapower_policyparameters" "test" {
  id = "PolicyParameters_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_policyparameters" "test" {
  depends_on = [ datapower_policyparameters.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ProbeTestConfig = ModelTestConfig{
	Name: "Probe",
	Resource: `
resource "datapower_probe" "test" {
  app_domain = "acc_test_domain"
  max_records = 1000
  gateway_peering = "default-gateway-peering"
}`,
	Data: `
data "datapower_probe" "test" {
  depends_on = [ datapower_probe.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"GatewayPeering": &GatewayPeeringTestConfig,
	},
	TestPre: `
`,
}
var ProcessingMetadataTestConfig = ModelTestConfig{
	Name: "ProcessingMetadata",
	Resource: `
resource "datapower_processingmetadata" "test" {
  id = "___ProcessingMetadata_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_processingmetadata" "test" {
  depends_on = [ datapower_processingmetadata.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var QuotaEnforcementServerTestConfig = ModelTestConfig{
	Name: "QuotaEnforcementServer",
	Resource: `
resource "datapower_quotaenforcementserver" "test" {
  server_port = 16379
  monitor_port = 26379
}`,
	Data: `
data "datapower_quotaenforcementserver" "test" {
  depends_on = [ datapower_quotaenforcementserver.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RADIUSSettingsTestConfig = ModelTestConfig{
	Name: "RADIUSSettings",
	Resource: `
resource "datapower_radiussettings" "test" {
}`,
	Data: `
data "datapower_radiussettings" "test" {
  depends_on = [ datapower_radiussettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RBMSettingsTestConfig = ModelTestConfig{
	Name: "RBMSettings",
	Resource: `
resource "datapower_rbmsettings" "test" {
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
data "datapower_rbmsettings" "test" {
  depends_on = [ datapower_rbmsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RaidVolumeTestConfig = ModelTestConfig{
	Name: "RaidVolume",
	Resource: `
resource "datapower_raidvolume" "test" {
}`,
	Data: `
data "datapower_raidvolume" "test" {
  depends_on = [ datapower_raidvolume.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RateLimitConfigurationTestConfig = ModelTestConfig{
	Name: "RateLimitConfiguration",
	Resource: `
resource "datapower_ratelimitconfiguration" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_ratelimitconfiguration" "test" {
  depends_on = [ datapower_ratelimitconfiguration.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RateLimitDefinitionTestConfig = ModelTestConfig{
	Name: "RateLimitDefinition",
	Resource: `
resource "datapower_ratelimitdefinition" "test" {
  id = "RateLimitDefinition_name"
  app_domain = "acc_test_domain"
  type = "rate"
  rate = 1000
}`,
	Data: `
data "datapower_ratelimitdefinition" "test" {
  depends_on = [ datapower_ratelimitdefinition.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var RateLimitDefinitionGroupTestConfig = ModelTestConfig{
	Name: "RateLimitDefinitionGroup",
	Resource: `
resource "datapower_ratelimitdefinitiongroup" "test" {
  id = "RateLimitDefinitionGroup_name"
  app_domain = "acc_test_domain"
  update_on_exceed = "all"
}`,
	Data: `
data "datapower_ratelimitdefinitiongroup" "test" {
  depends_on = [ datapower_ratelimitdefinitiongroup.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SAMLAttributesTestConfig = ModelTestConfig{
	Name: "SAMLAttributes",
	Resource: `
resource "datapower_samlattributes" "test" {
  id = "SAMLAttributes_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_samlattributes" "test" {
  depends_on = [ datapower_samlattributes.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SFTPFilePollerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "SFTPFilePollerSourceProtocolHandler",
	Resource: `
resource "datapower_sftpfilepollersourceprotocolhandler" "test" {
  id = "SFTPFilePollerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  ssh_client_connection = "TestAccSSHClientProfile"
  target_directory = "/"
  delay_between_polls = 60000
  input_file_match_pattern = ".*"
  processing_seize_timeout = 0
  xml_manager = "default"
}`,
	Data: `
data "datapower_sftpfilepollersourceprotocolhandler" "test" {
  depends_on = [ datapower_sftpfilepollersourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"SSHClientProfile": &SSHClientProfileTestConfig,
		"XMLManager":       &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var SLMActionTestConfig = ModelTestConfig{
	Name: "SLMAction",
	Resource: `
resource "datapower_slmaction" "test" {
  id = "___SLMAction_name"
  app_domain = "acc_test_domain"
  type = "log-only"
}`,
	Data: `
data "datapower_slmaction" "test" {
  depends_on = [ datapower_slmaction.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SLMCredClassTestConfig = ModelTestConfig{
	Name: "SLMCredClass",
	Resource: `
resource "datapower_slmcredclass" "test" {
  id = "SLMCredClass_name"
  app_domain = "acc_test_domain"
  cred_type = "aaa-mapped-credential"
}`,
	Data: `
data "datapower_slmcredclass" "test" {
  depends_on = [ datapower_slmcredclass.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SLMPolicyTestConfig = ModelTestConfig{
	Name: "SLMPolicy",
	Resource: `
resource "datapower_slmpolicy" "test" {
  id = "SLMPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_slmpolicy" "test" {
  depends_on = [ datapower_slmpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SLMRsrcClassTestConfig = ModelTestConfig{
	Name: "SLMRsrcClass",
	Resource: `
resource "datapower_slmrsrcclass" "test" {
  id = "SLMRsrcClass_name"
  app_domain = "acc_test_domain"
  rsrc_type = "aaa-mapped-resource"
}`,
	Data: `
data "datapower_slmrsrcclass" "test" {
  depends_on = [ datapower_slmrsrcclass.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SLMScheduleTestConfig = ModelTestConfig{
	Name: "SLMSchedule",
	Resource: `
resource "datapower_slmschedule" "test" {
  id = "SLMSchedule_name"
  app_domain = "acc_test_domain"
  start_time = "12:34:00"
  duration = 1440
}`,
	Data: `
data "datapower_slmschedule" "test" {
  depends_on = [ datapower_slmschedule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SMTPServerConnectionTestConfig = ModelTestConfig{
	Name: "SMTPServerConnection",
	Resource: `
resource "datapower_smtpserverconnection" "test" {
  id = "___SMTPServerConnection_name"
  app_domain = "acc_test_domain"
  mail_server_host = "localhost"
  mail_server_port = 25
}`,
	Data: `
data "datapower_smtpserverconnection" "test" {
  depends_on = [ datapower_smtpserverconnection.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SNMPSettingsTestConfig = ModelTestConfig{
	Name: "SNMPSettings",
	Resource: `
resource "datapower_snmpsettings" "test" {
  local_port = 161
  security_level = "authPriv"
  access_level = "read-only"
}`,
	Data: `
data "datapower_snmpsettings" "test" {
  depends_on = [ datapower_snmpsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SOAPHeaderDispositionTestConfig = ModelTestConfig{
	Name: "SOAPHeaderDisposition",
	Resource: `
resource "datapower_soapheaderdisposition" "test" {
  id = "SOAPHeaderDisposition_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_soapheaderdisposition" "test" {
  depends_on = [ datapower_soapheaderdisposition.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SQLDataSourceTestConfig = ModelTestConfig{
	Name: "SQLDataSource",
	Resource: `
resource "datapower_sqldatasource" "test" {
  id = "SQLDataSource_name"
  app_domain = "acc_test_domain"
  database = "Oracle"
  username = "username"
  password_alias = "TestAccPasswordAlias"
  data_source_id = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection = 10
  connect_timeout = 15
  query_timeout = 30
  idle_timeout = 180
}`,
	Data: `
data "datapower_sqldatasource" "test" {
  depends_on = [ datapower_sqldatasource.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"PasswordAlias": &PasswordAliasTestConfig,
	},
	TestPre: `
`,
}
var SSHClientProfileTestConfig = ModelTestConfig{
	Name: "SSHClientProfile",
	Resource: `
resource "datapower_sshclientprofile" "test" {
  id = "SSHClientProfile_name"
  app_domain = "acc_test_domain"
  user_name = "someuser"
  profile_usage = "sftp"
}`,
	Data: `
data "datapower_sshclientprofile" "test" {
  depends_on = [ datapower_sshclientprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSHDomainClientProfileTestConfig = ModelTestConfig{
	Name: "SSHDomainClientProfile",
	Resource: `
resource "datapower_sshdomainclientprofile" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_sshdomainclientprofile" "test" {
  depends_on = [ datapower_sshdomainclientprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSHServerProfileTestConfig = ModelTestConfig{
	Name: "SSHServerProfile",
	Resource: `
resource "datapower_sshserverprofile" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_sshserverprofile" "test" {
  depends_on = [ datapower_sshserverprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSHServerSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "SSHServerSourceProtocolHandler",
	Resource: `
resource "datapower_sshserversourceprotocolhandler" "test" {
  id = "SSHServerSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 22
  default_directory = "/"
}`,
	Data: `
data "datapower_sshserversourceprotocolhandler" "test" {
  depends_on = [ datapower_sshserversourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSHServiceTestConfig = ModelTestConfig{
	Name: "SSHService",
	Resource: `
resource "datapower_sshservice" "test" {
  local_port = 22
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_sshservice" "test" {
  depends_on = [ datapower_sshservice.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSLClientProfileTestConfig = ModelTestConfig{
	Name: "SSLClientProfile",
	Resource: `
resource "datapower_sslclientprofile" "test" {
  id = "SSLClientProfile_name"
  app_domain = "acc_test_domain"
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
}`,
	Data: `
data "datapower_sslclientprofile" "test" {
  depends_on = [ datapower_sslclientprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSLProxyServiceTestConfig = ModelTestConfig{
	Name: "SSLProxyService",
	Resource: `
resource "datapower_sslproxyservice" "test" {
  id = "SSLProxyService_name"
  app_domain = "acc_test_domain"
  local_port = 8888
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_sslproxyservice" "test" {
  depends_on = [ datapower_sslproxyservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SSLSNIMappingTestConfig = ModelTestConfig{
	Name: "SSLSNIMapping",
	Resource: `
resource "datapower_sslsnimapping" "test" {
  id = "SSLSNIMapping_name"
  app_domain = "acc_test_domain"
  sni_mapping = ` + DmHostToSSLServerProfileTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_sslsnimapping" "test" {
  depends_on = [ datapower_sslsnimapping.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmHostToSSLServerProfile": &DmHostToSSLServerProfileTestConfig,
	},
	TestPre: `
`,
}
var SSLSNIServerProfileTestConfig = ModelTestConfig{
	Name: "SSLSNIServerProfile",
	Resource: `
resource "datapower_sslsniserverprofile" "test" {
  id = "SSLSNIServerProfile_name"
  app_domain = "acc_test_domain"
  sni_server_mapping = "TestAccSSLSNIMappingHostnameMap"
}`,
	Data: `
data "datapower_sslsniserverprofile" "test" {
  depends_on = [ datapower_sslsniserverprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"SSLSNIMapping": &SSLSNIMappingTestConfig,
	},
	TestPre: `
`,
}
var SSLServerProfileTestConfig = ModelTestConfig{
	Name: "SSLServerProfile",
	Resource: `
resource "datapower_sslserverprofile" "test" {
  id = "SSLServerProfile_name"
  app_domain = "acc_test_domain"
  ciphers = ["AES_256_GCM_SHA384","CHACHA20_POLY1305_SHA256","AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_256_GCM_SHA384","ECDHE_ECDSA_WITH_AES_256_CBC_SHA384","ECDHE_ECDSA_WITH_AES_128_GCM_SHA256","ECDHE_ECDSA_WITH_AES_128_CBC_SHA256","ECDHE_ECDSA_WITH_AES_256_CBC_SHA","ECDHE_ECDSA_WITH_AES_128_CBC_SHA","ECDHE_RSA_WITH_AES_256_GCM_SHA384","ECDHE_RSA_WITH_AES_256_CBC_SHA384","ECDHE_RSA_WITH_AES_128_GCM_SHA256","ECDHE_RSA_WITH_AES_128_CBC_SHA256","ECDHE_RSA_WITH_AES_256_CBC_SHA","ECDHE_RSA_WITH_AES_128_CBC_SHA","DHE_RSA_WITH_AES_256_GCM_SHA384","DHE_RSA_WITH_AES_256_CBC_SHA256","DHE_RSA_WITH_AES_128_GCM_SHA256","DHE_RSA_WITH_AES_128_CBC_SHA256","DHE_RSA_WITH_AES_256_CBC_SHA","DHE_RSA_WITH_AES_128_CBC_SHA",]
  idcred = "TestAccCryptoIdentCred"
}`,
	Data: `
data "datapower_sslserverprofile" "test" {
  depends_on = [ datapower_sslserverprofile.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"CryptoIdentCred": &CryptoIdentCredTestConfig,
	},
	TestPre: `
`,
}
var SchemaExceptionMapTestConfig = ModelTestConfig{
	Name: "SchemaExceptionMap",
	Resource: `
resource "datapower_schemaexceptionmap" "test" {
  id = "SchemaExceptionMap_name"
  app_domain = "acc_test_domain"
  original_schema_url = "http://localhost"
  schema_exception_rules = ` + DmSchemaExceptionRuleTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_schemaexceptionmap" "test" {
  depends_on = [ datapower_schemaexceptionmap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmSchemaExceptionRule": &DmSchemaExceptionRuleTestConfig,
	},
	TestPre: `
`,
}
var SecureBackupModeTestConfig = ModelTestConfig{
	Name: "SecureBackupMode",
	Resource: `
resource "datapower_securebackupmode" "test" {
}`,
	Data: `
data "datapower_securebackupmode" "test" {
  depends_on = [ datapower_securebackupmode.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SocialLoginPolicyTestConfig = ModelTestConfig{
	Name: "SocialLoginPolicy",
	Resource: `
resource "datapower_socialloginpolicy" "test" {
  id = "SocialLoginPolicy_name"
  app_domain = "acc_test_domain"
  client_id = "client_id"
  client_secret = "client_secret"
  client_ssl_profile = "TestAccSSLClientProfile"
  social_provider = "google"
  provider_az_endpoint = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token = true
}`,
	Data: `
data "datapower_socialloginpolicy" "test" {
  depends_on = [ datapower_socialloginpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"SSLClientProfile": &SSLClientProfileTestConfig,
	},
	TestPre: `
`,
}
var StatelessTCPSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "StatelessTCPSourceProtocolHandler",
	Resource: `
resource "datapower_statelesstcpsourceprotocolhandler" "test" {
  id = "StatelessTCPSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 4000
}`,
	Data: `
data "datapower_statelesstcpsourceprotocolhandler" "test" {
  depends_on = [ datapower_statelesstcpsourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var StatisticsTestConfig = ModelTestConfig{
	Name: "Statistics",
	Resource: `
resource "datapower_statistics" "test" {
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_statistics" "test" {
  depends_on = [ datapower_statistics.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var StylePolicyTestConfig = ModelTestConfig{
	Name: "StylePolicy",
	Resource: `
resource "datapower_stylepolicy" "test" {
  id = "___StylePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_stylepolicy" "test" {
  depends_on = [ datapower_stylepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var StylePolicyActionTestConfig = ModelTestConfig{
	Name: "StylePolicyAction",
	Resource: `
resource "datapower_stylepolicyaction" "test" {
  id = "___StylePolicyAction_test"
  app_domain = "acc_test_domain"
  type = "xform"
  named_inputs = null
  named_outputs = null
}`,
	Data: `
data "datapower_stylepolicyaction" "test" {
  depends_on = [ datapower_stylepolicyaction.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmNamedInOut": &DmNamedInOutTestConfig,
	},
	TestPre: `
`,
}
var StylePolicyRuleTestConfig = ModelTestConfig{
	Name: "StylePolicyRule",
	Resource: `
resource "datapower_stylepolicyrule" "test" {
  id = "___StylePolicyRule_name"
  app_domain = "acc_test_domain"
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
	Data: `
data "datapower_stylepolicyrule" "test" {
  depends_on = [ datapower_stylepolicyrule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var SystemSettingsTestConfig = ModelTestConfig{
	Name: "SystemSettings",
	Resource: `
resource "datapower_systemsettings" "test" {
}`,
	Data: `
data "datapower_systemsettings" "test" {
  depends_on = [ datapower_systemsettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var TAMTestConfig = ModelTestConfig{
	Name: "TAM",
	Resource: `
resource "datapower_tam" "test" {
  id = "TAM_name"
  app_domain = "acc_test_domain"
  configuration_file = "local:///tam.config"
  ssl_key_file = "cert:///ssl.key"
  ssl_key_stash_file = "cert:///ssl.stash"
  use_local_mode = false
  tam_use_basic_user = false
}`,
	Data: `
data "datapower_tam" "test" {
  depends_on = [ datapower_tam.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var TCPProxyServiceTestConfig = ModelTestConfig{
	Name: "TCPProxyService",
	Resource: `
resource "datapower_tcpproxyservice" "test" {
  id = "TCPProxyService_name"
  app_domain = "acc_test_domain"
  local_port = 8888
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_tcpproxyservice" "test" {
  depends_on = [ datapower_tcpproxyservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ThrottlerTestConfig = ModelTestConfig{
	Name: "Throttler",
	Resource: `
resource "datapower_throttler" "test" {
  throttle_at = 20
  terminate_at = 5
  temp_fs_throttle_at = 0
  temp_fs_terminate_at = 0
  qname_warn_at = 10
  timeout = 30
}`,
	Data: `
data "datapower_throttler" "test" {
  depends_on = [ datapower_throttler.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var TimeSettingsTestConfig = ModelTestConfig{
	Name: "TimeSettings",
	Resource: `
resource "datapower_timesettings" "test" {
  local_time_zone = "EST5EDT"
}`,
	Data: `
data "datapower_timesettings" "test" {
  depends_on = [ datapower_timesettings.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var URLMapTestConfig = ModelTestConfig{
	Name: "URLMap",
	Resource: `
resource "datapower_urlmap" "test" {
  id = "___URLMap_name"
  app_domain = "acc_test_domain"
  url_map_rule = ` + DmURLMapRuleTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_urlmap" "test" {
  depends_on = [ datapower_urlmap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmURLMapRule": &DmURLMapRuleTestConfig,
	},
	TestPre: `
`,
}
var URLRefreshPolicyTestConfig = ModelTestConfig{
	Name: "URLRefreshPolicy",
	Resource: `
resource "datapower_urlrefreshpolicy" "test" {
  id = "URLRefreshPolicy_name"
  app_domain = "acc_test_domain"
  url_refresh_rule = ` + DmURLRefreshRuleTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_urlrefreshpolicy" "test" {
  depends_on = [ datapower_urlrefreshpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmURLRefreshRule": &DmURLRefreshRuleTestConfig,
	},
	TestPre: `
`,
}
var URLRewritePolicyTestConfig = ModelTestConfig{
	Name: "URLRewritePolicy",
	Resource: `
resource "datapower_urlrewritepolicy" "test" {
  id = "URLRewritePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_urlrewritepolicy" "test" {
  depends_on = [ datapower_urlrewritepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var UserTestConfig = ModelTestConfig{
	Name: "User",
	Resource: `
resource "datapower_user" "test" {
  id = "0user"
  password = "Password$123"
  access_level = "group-defined"
  group_name = "TestAccUserGroup"
  snmp_creds = null
  hashed_snmp_creds = null
}`,
	Data: `
data "datapower_user" "test" {
  depends_on = [ datapower_user.test ]
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"UserGroup":        &UserGroupTestConfig,
		"DmSnmpCred":       &DmSnmpCredTestConfig,
		"DmSnmpCredMasked": &DmSnmpCredMaskedTestConfig,
	},
	TestPre: `
`,
}
var UserGroupTestConfig = ModelTestConfig{
	Name: "UserGroup",
	Resource: `
resource "datapower_usergroup" "test" {
  id = "_UserGroup_name"
  access_policies = ["*/*/*?Access=r"]
}`,
	Data: `
data "datapower_usergroup" "test" {
  depends_on = [ datapower_usergroup.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var VisibilityListTestConfig = ModelTestConfig{
	Name: "VisibilityList",
	Resource: `
resource "datapower_visibilitylist" "test" {
  id = "VisibilityList_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_visibilitylist" "test" {
  depends_on = [ datapower_visibilitylist.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WCCServiceTestConfig = ModelTestConfig{
	Name: "WCCService",
	Resource: `
resource "datapower_wccservice" "test" {
  id = "WCCService_name"
  app_domain = "acc_test_domain"
  odc_info_hostname = "example.com"
  odc_info_port = 1024
  time_interval = 10
}`,
	Data: `
data "datapower_wccservice" "test" {
  depends_on = [ datapower_wccservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WSEndpointRewritePolicyTestConfig = ModelTestConfig{
	Name: "WSEndpointRewritePolicy",
	Resource: `
resource "datapower_wsendpointrewritepolicy" "test" {
  id = "WSEndpointRewritePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_wsendpointrewritepolicy" "test" {
  depends_on = [ datapower_wsendpointrewritepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WSGatewayTestConfig = ModelTestConfig{
	Name: "WSGateway",
	Resource: `
resource "datapower_wsgateway" "test" {
  id = "WSGateway_name"
  app_domain = "acc_test_domain"
  type = "static-from-wsdl"
  style_policy = "default"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
	Data: `
data "datapower_wsgateway" "test" {
  depends_on = [ datapower_wsgateway.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"WSStylePolicy": &WSStylePolicyTestConfig,
		"XMLManager":    &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var WSRRSavedSearchSubscriptionTestConfig = ModelTestConfig{
	Name: "WSRRSavedSearchSubscription",
	Resource: `
resource "datapower_wsrrsavedsearchsubscription" "test" {
  id = "WSRRSavedSearchSubscription_name"
  app_domain = "acc_test_domain"
  server = "TestAccWSRRServer"
  saved_search_name = "saved_search_name"
}`,
	Data: `
data "datapower_wsrrsavedsearchsubscription" "test" {
  depends_on = [ datapower_wsrrsavedsearchsubscription.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"WSRRServer": &WSRRServerTestConfig,
	},
	TestPre: `
`,
}
var WSRRServerTestConfig = ModelTestConfig{
	Name: "WSRRServer",
	Resource: `
resource "datapower_wsrrserver" "test" {
  id = "_WSRRServer_name"
  app_domain = "acc_test_domain"
  soap_url = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}`,
	Data: `
data "datapower_wsrrserver" "test" {
  depends_on = [ datapower_wsrrserver.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WSRRSubscriptionTestConfig = ModelTestConfig{
	Name: "WSRRSubscription",
	Resource: `
resource "datapower_wsrrsubscription" "test" {
  id = "WSRRSubscription_name"
  app_domain = "acc_test_domain"
  server = "TestAccWSRRServer"
  object_type = "wsdl"
  object_name = "object_name"
}`,
	Data: `
data "datapower_wsrrsubscription" "test" {
  depends_on = [ datapower_wsrrsubscription.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"WSRRServer": &WSRRServerTestConfig,
	},
	TestPre: `
`,
}
var WSStylePolicyTestConfig = ModelTestConfig{
	Name: "WSStylePolicy",
	Resource: `
resource "datapower_wsstylepolicy" "test" {
  id = "0WSStylePolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_wsstylepolicy" "test" {
  depends_on = [ datapower_wsstylepolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WSStylePolicyRuleTestConfig = ModelTestConfig{
	Name: "WSStylePolicyRule",
	Resource: `
resource "datapower_wsstylepolicyrule" "test" {
  id = "_WSStylePolicyRule_name"
  app_domain = "acc_test_domain"
  direction = "rule"
  input_format = "none"
  output_format = "none"
}`,
	Data: `
data "datapower_wsstylepolicyrule" "test" {
  depends_on = [ datapower_wsstylepolicyrule.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WXSGridTestConfig = ModelTestConfig{
	Name: "WXSGrid",
	Resource: `
resource "datapower_wxsgrid" "test" {
  id = "WXSGrid_name"
  app_domain = "acc_test_domain"
  collective = "TestAccLoadBalancerGroup"
  grid = "gridname"
  user_name = "username"
  password_alias = "TestAccPasswordAlias"
  timeout = 1000
}`,
	Data: `
data "datapower_wxsgrid" "test" {
  depends_on = [ datapower_wxsgrid.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"LoadBalancerGroup": &LoadBalancerGroupTestConfig,
		"PasswordAlias":     &PasswordAliasTestConfig,
	},
	TestPre: `
`,
}
var WebAppErrorHandlingPolicyTestConfig = ModelTestConfig{
	Name: "WebAppErrorHandlingPolicy",
	Resource: `
resource "datapower_webapperrorhandlingpolicy" "test" {
  id = "WebAppErrorHandlingPolicy_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_webapperrorhandlingpolicy" "test" {
  depends_on = [ datapower_webapperrorhandlingpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebAppFWTestConfig = ModelTestConfig{
	Name: "WebAppFW",
	Resource: `
resource "datapower_webappfw" "test" {
  id = "WebAppFW_name"
  app_domain = "acc_test_domain"
  remote_address = "10.10.10.10"
  style_policy = "TestAccAppSecurityPolicy"
  xml_manager = "default"
  front_timeout = 120
  back_timeout = 120
  front_persistent_timeout = 180
  back_persistent_timeout = 180
}`,
	Data: `
data "datapower_webappfw" "test" {
  depends_on = [ datapower_webappfw.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"AppSecurityPolicy": &AppSecurityPolicyTestConfig,
		"XMLManager":        &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var WebAppRequestTestConfig = ModelTestConfig{
	Name: "WebAppRequest",
	Resource: `
resource "datapower_webapprequest" "test" {
  id = "_WebAppRequest_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_webapprequest" "test" {
  depends_on = [ datapower_webapprequest.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebAppResponseTestConfig = ModelTestConfig{
	Name: "WebAppResponse",
	Resource: `
resource "datapower_webappresponse" "test" {
  id = "_WebAppResponse_name"
  app_domain = "acc_test_domain"
  policy_type = "admission"
}`,
	Data: `
data "datapower_webappresponse" "test" {
  depends_on = [ datapower_webappresponse.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebAppSessionPolicyTestConfig = ModelTestConfig{
	Name: "WebAppSessionPolicy",
	Resource: `
resource "datapower_webappsessionpolicy" "test" {
  id = "WebAppSessionPolicy_name"
  app_domain = "acc_test_domain"
  start_matches = "__default-accept-service-providers__"
}`,
	Data: `
data "datapower_webappsessionpolicy" "test" {
  depends_on = [ datapower_webappsessionpolicy.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"Matching": &MatchingTestConfig,
	},
	TestPre: `
`,
}
var WebB2BViewerTestConfig = ModelTestConfig{
	Name: "WebB2BViewer",
	Resource: `
resource "datapower_webb2bviewer" "test" {
  local_port = 9091
  idle_timeout = 600
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_webb2bviewer" "test" {
  depends_on = [ datapower_webb2bviewer.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebGUITestConfig = ModelTestConfig{
	Name: "WebGUI",
	Resource: `
resource "datapower_webgui" "test" {
  local_port = 9090
  save_config_overwrites = true
  idle_timeout = 600
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_webgui" "test" {
  depends_on = [ datapower_webgui.test ]
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebServiceMonitorTestConfig = ModelTestConfig{
	Name: "WebServiceMonitor",
	Resource: `
resource "datapower_webservicemonitor" "test" {
  id = "WebServiceMonitor_name"
  app_domain = "acc_test_domain"
  wsdlurl = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}`,
	Data: `
data "datapower_webservicemonitor" "test" {
  depends_on = [ datapower_webservicemonitor.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebServicesAgentTestConfig = ModelTestConfig{
	Name: "WebServicesAgent",
	Resource: `
resource "datapower_webservicesagent" "test" {
  app_domain = "acc_test_domain"
  max_records = 3000
  max_memory_kb = 64000
  capture_mode = "faults"
}`,
	Data: `
data "datapower_webservicesagent" "test" {
  depends_on = [ datapower_webservicesagent.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var WebSphereJMSServerTestConfig = ModelTestConfig{
	Name: "WebSphereJMSServer",
	Resource: `
resource "datapower_webspherejmsserver" "test" {
  id = "_WebSphereJMSServer_name"
  app_domain = "acc_test_domain"
  endpoint = ` + DmWebSphereJMSEndpointTestConfig.GetModelListConfig() + `
  messaging_bus = "bus"
}`,
	Data: `
data "datapower_webspherejmsserver" "test" {
  depends_on = [ datapower_webspherejmsserver.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmWebSphereJMSEndpoint": &DmWebSphereJMSEndpointTestConfig,
	},
	TestPre: `
`,
}
var WebSphereJMSSourceProtocolHandlerTestConfig = ModelTestConfig{
	Name: "WebSphereJMSSourceProtocolHandler",
	Resource: `
resource "datapower_webspherejmssourceprotocolhandler" "test" {
  id = "WebSphereJMSSourceProtocolHandler_name"
  app_domain = "acc_test_domain"
  server = "TestAccWebSphereJMSServer"
  get_queue = "getqueue"
}`,
	Data: `
data "datapower_webspherejmssourceprotocolhandler" "test" {
  depends_on = [ datapower_webspherejmssourceprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"WebSphereJMSServer": &WebSphereJMSServerTestConfig,
	},
	TestPre: `
`,
}
var WebTokenServiceTestConfig = ModelTestConfig{
	Name: "WebTokenService",
	Resource: `
resource "datapower_webtokenservice" "test" {
  id = "WebTokenService_name"
  app_domain = "acc_test_domain"
  xml_manager = "default"
  style_policy = "default"
  front_timeout = 120
  front_persistent_timeout = 180
}`,
	Data: `
data "datapower_webtokenservice" "test" {
  depends_on = [ datapower_webtokenservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager":  &XMLManagerTestConfig,
		"StylePolicy": &StylePolicyTestConfig,
	},
	TestPre: `
`,
}
var XACMLPDPTestConfig = ModelTestConfig{
	Name: "XACMLPDP",
	Resource: `
resource "datapower_xacmlpdp" "test" {
  id = "XACMLPDP_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_xacmlpdp" "test" {
  depends_on = [ datapower_xacmlpdp.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var XMLFirewallServiceTestConfig = ModelTestConfig{
	Name: "XMLFirewallService",
	Resource: `
resource "datapower_xmlfirewallservice" "test" {
  id = "XMLFirewallService_name"
  app_domain = "acc_test_domain"
  type = "dynamic-backend"
  xml_manager = "default"
  local_port = 8888
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_xmlfirewallservice" "test" {
  depends_on = [ datapower_xmlfirewallservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var XMLManagerTestConfig = ModelTestConfig{
	Name: "XMLManager",
	Resource: `
resource "datapower_xmlmanager" "test" {
  id = "0_XMLManger_name"
  app_domain = "acc_test_domain"
}`,
	Data: `
data "datapower_xmlmanager" "test" {
  depends_on = [ datapower_xmlmanager.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var XPathRoutingMapTestConfig = ModelTestConfig{
	Name: "XPathRoutingMap",
	Resource: `
resource "datapower_xpathroutingmap" "test" {
  id = "XPathRoutingMap_name"
  app_domain = "acc_test_domain"
  x_path_routing_rules = ` + DmXPathRoutingRuleTestConfig.GetModelListConfig() + `
}`,
	Data: `
data "datapower_xpathroutingmap" "test" {
  depends_on = [ datapower_xpathroutingmap.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"DmXPathRoutingRule": &DmXPathRoutingRuleTestConfig,
	},
	TestPre: `
`,
}
var XSLCoprocServiceTestConfig = ModelTestConfig{
	Name: "XSLCoprocService",
	Resource: `
resource "datapower_xslcoprocservice" "test" {
  id = "XSLCoprocService_name"
  app_domain = "acc_test_domain"
  local_port = 8888
  xml_manager = "default"
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_xslcoprocservice" "test" {
  depends_on = [ datapower_xslcoprocservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var XSLProxyServiceTestConfig = ModelTestConfig{
	Name: "XSLProxyService",
	Resource: `
resource "datapower_xslproxyservice" "test" {
  id = "XSLProxyService_name"
  app_domain = "acc_test_domain"
  type = "static-backend"
  xml_manager = "default"
  local_port = 8899
  remote_address = "10.10.10.10"
  remote_port = 9999
  local_address = "0.0.0.0"
}`,
	Data: `
data "datapower_xslproxyservice" "test" {
  depends_on = [ datapower_xslproxyservice.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly: false,
	Dependencies: map[string]*ModelTestConfig{
		"XMLManager": &XMLManagerTestConfig,
	},
	TestPre: `
`,
}
var XTCProtocolHandlerTestConfig = ModelTestConfig{
	Name: "XTCProtocolHandler",
	Resource: `
resource "datapower_xtcprotocolhandler" "test" {
  id = "XTCProtocolHandler_name"
  app_domain = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port = 3000
  remote_address = "10.10.10.10"
  remote_port = 12000
}`,
	Data: `
data "datapower_xtcprotocolhandler" "test" {
  depends_on = [ datapower_xtcprotocolhandler.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
var ZosNSSClientTestConfig = ModelTestConfig{
	Name: "ZosNSSClient",
	Resource: `
resource "datapower_zosnssclient" "test" {
  id = "ZosNSSClient_name"
  app_domain = "acc_test_domain"
  remote_address = "remote.host"
  remote_port = 443
  client_id = "client_id"
  system_name = "system_name"
  user_name = "username"
}`,
	Data: `
data "datapower_zosnssclient" "test" {
  depends_on = [ datapower_zosnssclient.test ]
  app_domain = "acc_test_domain"
}`,
	ModelOnly:    false,
	Dependencies: map[string]*ModelTestConfig{},
	TestPre: `
`,
}
