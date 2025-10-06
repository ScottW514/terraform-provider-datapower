terraform {
  required_providers {
    datapower = {
      source = "scottw514/datapower"
    }
  }
}
resource "datapower_aaa_jwt_generator" "acc_test" {
  id         = "AccTest_AAAJWTGenerator"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_aaa_jwt_validator" "acc_test" {
  id             = "AccTest_AAAJWTValidator"
  app_domain     = datapower_domain.acc_test.app_domain
  username_claim = "sub"
}
resource "datapower_aaa_policy" "acc_test" {
  id         = "AccTest_AAAPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_access_control_list" "acc_test" {
  id         = "AccTest_AccessControlList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_access_profile" "acc_test" {
  id         = "AccTest_AccessProfile"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_amqp_broker" "acc_test" {
  id            = "AccTest_AMQPBroker"
  app_domain    = datapower_domain.acc_test.app_domain
  host_name     = "host.name"
  port          = 5672
  xml_manager   = "default"
  authorization = "none"
}
resource "datapower_amqp_source_protocol_handler" "acc_test" {
  id         = "AccTest_AMQPSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  broker     = datapower_amqp_broker.acc_test.id
  from       = "amqpfrom"
}
resource "datapower_analytics_endpoint" "acc_test" {
  id                     = "AccTest_AnalyticsEndpoint"
  app_domain             = datapower_domain.acc_test.app_domain
  analytics_server_url   = "https://localhost"
  ssl_client             = datapower_ssl_client_profile.acc_test.id
  max_records            = 1024
  max_records_memory_kb  = 512
  max_delivery_memory_mb = 512
}
resource "datapower_api_application_type" "acc_test" {
  id         = "AccTest_APIApplicationType"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_auth_url_registry" "acc_test" {
  id         = "AccTest_APIAuthURLRegistry"
  app_domain = datapower_domain.acc_test.app_domain
  auth_url   = "http://localhost"
}
resource "datapower_api_client_identification" "acc_test" {
  id         = "AccTest_APIClientIdentification"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_collection" "acc_test" {
  id         = "AccTest_APICollection"
  app_domain = datapower_domain.acc_test.app_domain
  org_id     = "orgid"
  org_name   = "orgname"
  routing_prefix = [{
    type = "host"
  }]
  api_processing_rule = "default-api-rule"
  api_error_rule      = "default-api-error-rule"
  plan                = [datapower_api_plan.acc_test.id]
}
resource "datapower_api_connect_gateway_service" "acc_test" {
  app_domain      = datapower_domain.acc_test.app_domain
  local_address   = "0.0.0.0"
  local_port      = 3000
  gateway_peering = "default-gateway-peering"
}
resource "datapower_api_cors" "acc_test" {
  id         = "AccTest_APICORS"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_definition" "acc_test" {
  id         = "AccTest_APIDefinition"
  app_domain = datapower_domain.acc_test.app_domain
  base_path  = "/"
  path       = [datapower_api_path.acc_test.id]
}
resource "datapower_api_execute" "acc_test" {
  id         = "AccTest_APIExecute"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_final" "acc_test" {
  id         = "AccTest_APIFinal"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_gateway" "acc_test" {
  id                       = "AccTest_APIGateway"
  app_domain               = datapower_domain.acc_test.app_domain
  front_protocol           = [datapower_http_source_protocol_handler.acc_test.id]
  front_timeout            = 120
  front_persistent_timeout = 180
}
resource "datapower_api_ldap_registry" "acc_test" {
  id                     = "AccTest_APILDAPRegistry"
  app_domain             = datapower_domain.acc_test.app_domain
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = datapower_ldap_search_parameters.acc_test.id
}
resource "datapower_api_operation" "acc_test" {
  id         = "AccTest_APIOperation"
  app_domain = datapower_domain.acc_test.app_domain
  method     = "GET"
}
resource "datapower_api_path" "acc_test" {
  id         = "AccTest_APIPath"
  app_domain = datapower_domain.acc_test.app_domain
  path       = "/"
}
resource "datapower_api_plan" "acc_test" {
  id               = "AccTest_APIPlan"
  app_domain       = datapower_domain.acc_test.app_domain
  api              = [datapower_api_definition.acc_test.id]
  rate_limit_scope = "per-application"
}
resource "datapower_api_result" "acc_test" {
  id         = "AccTest_APIResult"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_routing" "acc_test" {
  id         = "AccTest_APIRouting"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_rule" "acc_test" {
  id         = "AccTest_APIRule"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_schema" "acc_test" {
  id          = "AccTest_APISchema"
  app_domain  = datapower_domain.acc_test.app_domain
  json_schema = "http://localhost/json"
}
resource "datapower_api_security" "acc_test" {
  id         = "AccTest_APISecurity"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_security_api_key" "acc_test" {
  id         = "AccTest_APISecurityAPIKey"
  app_domain = datapower_domain.acc_test.app_domain
  where      = "header"
  type       = "id"
  key_name   = "keyname"
}
resource "datapower_api_security_basic_auth" "acc_test" {
  id            = "AccTest_APISecurityBasicAuth"
  app_domain    = datapower_domain.acc_test.app_domain
  user_registry = datapower_api_auth_url_registry.acc_test.id
}
resource "datapower_api_security_http_scheme" "acc_test" {
  id         = "AccTest_APISecurityHTTPScheme"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_api_security_oauth" "acc_test" {
  id             = "AccTest_APISecurityOAuth"
  app_domain     = datapower_domain.acc_test.app_domain
  oauth_provider = datapower_oauth_provider_settings.acc_test.id
  oauth_flow     = "implicit"
}
resource "datapower_api_security_oauth_req" "acc_test" {
  id                     = "AccTest_APISecurityOAuthReq"
  app_domain             = datapower_domain.acc_test.app_domain
  api_security_oauth_def = datapower_api_security_oauth.acc_test.id
}
resource "datapower_api_security_requirement" "acc_test" {
  id         = "AccTest_APISecurityRequirement"
  app_domain = datapower_domain.acc_test.app_domain
  security   = [datapower_api_security_api_key.acc_test.id]
}
resource "datapower_api_security_token_manager" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_app_security_policy" "acc_test" {
  id         = "AccTest_AppSecurityPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  request_maps = [{
    match = "__default-accept-service-providers__"
    rule  = datapower_web_app_request.acc_test.id
  }]
  response_maps = [{
    match = "__default-accept-service-providers__"
    rule  = datapower_web_app_response.acc_test.id
  }]
}
resource "datapower_as1_poller_source_protocol_handler" "acc_test" {
  id                    = "AccTest_AS1PollerSourceProtocolHandler"
  app_domain            = datapower_domain.acc_test.app_domain
  mail_server           = "smtp.host"
  port                  = 25
  conn_security         = "none"
  account               = "account"
  password_alias        = datapower_password_alias.acc_test.id
  delay_between_polls   = 300
  max_messages_per_poll = 5
}
resource "datapower_as2_source_protocol_handler" "acc_test" {
  id            = "AccTest_AS2SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_as3_source_protocol_handler" "acc_test" {
  id            = "AccTest_AS3SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 21
}
resource "datapower_assembly" "acc_test" {
  id         = "AccTest_Assembly"
  app_domain = datapower_domain.acc_test.app_domain
  rule       = "default-empty-rule"
}
resource "datapower_assembly_action_client_security" "acc_test" {
  id                         = "AccTest_AssemblyActionClientSecurity"
  app_domain                 = datapower_domain.acc_test.app_domain
  id_name                    = "idname"
  secret_name                = "secretname"
  authenticate_client_method = "native"
}
resource "datapower_assembly_action_extract" "acc_test" {
  id         = "AccTest_AssemblyActionExtract"
  app_domain = datapower_domain.acc_test.app_domain
  extract = [{
    capture = "capture"
  }]
}
resource "datapower_assembly_action_function_call" "acc_test" {
  id            = "AccTest_AssemblyActionFunctionCall"
  app_domain    = datapower_domain.acc_test.app_domain
  function_call = "default-func-global"
}
resource "datapower_assembly_action_gateway_script" "acc_test" {
  id         = "AccTest_AssemblyActionGatewayScript"
  app_domain = datapower_domain.acc_test.app_domain
  source     = "gsfile"
}
resource "datapower_assembly_action_graphql_cost_analysis" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLCostAnalysis"
  app_domain = datapower_domain.acc_test.app_domain
  target     = "targetquery"
}
resource "datapower_assembly_action_graphql_execute" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLExecute"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_graphql_introspect" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLIntrospect"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_html_page" "acc_test" {
  id         = "AccTest_AssemblyActionHtmlPage"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_invoke" "acc_test" {
  id           = "AccTest_AssemblyActionInvoke"
  app_domain   = datapower_domain.acc_test.app_domain
  url          = "https://localhost"
  method       = "Keep"
  backend_type = "detect"
  cache_type   = "Protocol"
}
resource "datapower_assembly_action_json2xml" "acc_test" {
  id         = "AccTest_AssemblyActionJson2Xml"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_jwt_generate" "acc_test" {
  id              = "AccTest_AssemblyActionJWTGenerate"
  app_domain      = datapower_domain.acc_test.app_domain
  issuer_claim    = "iss.claim"
  validity_period = 3600
}
resource "datapower_assembly_action_jwt_validate" "acc_test" {
  id            = "AccTest_AssemblyActionJWTValidate"
  app_domain    = datapower_domain.acc_test.app_domain
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}
resource "datapower_assembly_action_log" "acc_test" {
  id         = "AccTest_AssemblyActionLog"
  app_domain = datapower_domain.acc_test.app_domain
  mode       = "gather-only"
}
resource "datapower_assembly_action_map" "acc_test" {
  id         = "AccTest_AssemblyActionMap"
  app_domain = datapower_domain.acc_test.app_domain
  location   = "local:///file"
}
resource "datapower_assembly_action_oauth" "acc_test" {
  id                                = "AccTest_AssemblyActionOAuth"
  app_domain                        = datapower_domain.acc_test.app_domain
  oauth_provider_settings_reference = { "Default" : datapower_oauth_provider_settings.acc_test.id }
}
resource "datapower_assembly_action_parse" "acc_test" {
  id         = "AccTest_AssemblyActionParse"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_rate_limit" "acc_test" {
  id         = "AccTest_AssemblyActionRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  source     = "plan-default"
}
resource "datapower_assembly_action_rate_limit_info" "acc_test" {
  id         = "AccTest_AssemblyActionRateLimitInfo"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_redact" "acc_test" {
  id         = "AccTest_AssemblyActionRedact"
  app_domain = datapower_domain.acc_test.app_domain
  redact = [{
    path = "path"
  }]
}
resource "datapower_assembly_action_set_var" "acc_test" {
  id         = "AccTest_AssemblyActionSetVar"
  app_domain = datapower_domain.acc_test.app_domain
  variable = [{
    name  = "varname"
    value = "value"
  }]
}
resource "datapower_assembly_action_throw" "acc_test" {
  id         = "AccTest_AssemblyActionThrow"
  app_domain = datapower_domain.acc_test.app_domain
  error_id   = "errorid"
}
resource "datapower_assembly_action_user_security" "acc_test" {
  id                      = "AccTest_AssemblyActionUserSecurity"
  app_domain              = datapower_domain.acc_test.app_domain
  factor_id               = "default"
  extract_identity_method = "basic"
  user_auth_method        = "user-registry"
  user_registry           = datapower_api_auth_url_registry.acc_test.id
  user_az_method          = "authenticated"
}
resource "datapower_assembly_action_validate" "acc_test" {
  id         = "AccTest_AssemblyActionValidate"
  app_domain = datapower_domain.acc_test.app_domain
  schema     = datapower_api_schema.acc_test.id
}
resource "datapower_assembly_action_web_socket_upgrade" "acc_test" {
  id         = "AccTest_AssemblyActionWebSocketUpgrade"
  app_domain = datapower_domain.acc_test.app_domain
  url        = "https://localhost"
}
resource "datapower_assembly_action_wsdl" "acc_test" {
  id         = "AccTest_AssemblyActionWSDL"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_xml2json" "acc_test" {
  id         = "AccTest_AssemblyActionXml2Json"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_action_xslt" "acc_test" {
  id         = "AccTest_AssemblyActionXSLT"
  app_domain = datapower_domain.acc_test.app_domain
  stylesheet = "local:///stylesheet"
}
resource "datapower_assembly_function" "acc_test" {
  id         = "AccTest_AssemblyFunction"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assembly_logic_operation_switch" "acc_test" {
  id         = "AccTest_AssemblyLogicOperationSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = [{
    execute      = "default-api-rule"
    operation_id = "operationid"
  }]
}
resource "datapower_assembly_logic_switch" "acc_test" {
  id         = "AccTest_AssemblyLogicSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = [{
    condition = "condition"
    execute   = "default-api-rule"
  }]
}
resource "datapower_audit_log" "acc_test" {
}
resource "datapower_b2b_cpa" "acc_test" {
  id         = "AccTest_B2BCPA"
  app_domain = datapower_domain.acc_test.app_domain
  cpa_id     = "cpaid"
}
resource "datapower_b2b_cpa_collaboration" "acc_test" {
  id            = "AccTest_B2BCPACollaboration"
  app_domain    = datapower_domain.acc_test.app_domain
  internal_role = "internal"
  external_role = "external"
  service       = "service"
  actions = [{
    name           = "cpacollaborationactionname"
    value          = "value"
    capability     = "cansend"
    sender_setting = datapower_b2b_cpa_sender_setting.acc_test.id
  }]
}
resource "datapower_b2b_cpa_receiver_setting" "acc_test" {
  id         = "AccTest_B2BCPAReceiverSetting"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_b2b_cpa_sender_setting" "acc_test" {
  id                = "AccTest_B2BCPASenderSetting"
  app_domain        = datapower_domain.acc_test.app_domain
  dest_endpoint_url = "ebms2://somehost/url"
}
resource "datapower_b2b_gateway" "acc_test" {
  id                            = "AccTest_B2BGateway"
  app_domain                    = datapower_domain.acc_test.app_domain
  b2b_profiles                  = [{ partner_profile = datapower_b2b_profile.acc_test.id }]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "purgeonly"
  xml_manager                   = "default"
}
resource "datapower_b2b_persistence" "acc_test" {
  raid_volume  = "raid0"
  storage_size = 1024
  ha_enabled   = false
}
resource "datapower_b2b_profile" "acc_test" {
  id           = "AccTest_B2BProfile"
  app_domain   = datapower_domain.acc_test.app_domain
  business_ids = ["businessid"]
  destinations = [{
    dest_name  = "b2bdestinationname"
    dest_url   = "https://localhost"
    ssl_client = datapower_ssl_client_profile.acc_test.id
  }]
}
resource "datapower_b2b_profile_group" "acc_test" {
  id         = "AccTest_B2BProfileGroup"
  app_domain = datapower_domain.acc_test.app_domain
  b2b_profiles = [{
    partner_profile = datapower_b2b_profile.acc_test.id
  }]
}
resource "datapower_b2b_xpath_routing_policy" "acc_test" {
  id             = "AccTest_B2BXPathRoutingPolicy"
  app_domain     = datapower_domain.acc_test.app_domain
  sender_xpath   = "senderpath"
  receiver_xpath = "senderpath"
}
resource "datapower_cert_monitor" "acc_test" {
  polling_interval      = 1
  reminder_time         = 30
  log_level             = "warn"
  disable_expired_certs = false
}
resource "datapower_compile_options_policy" "acc_test" {
  id         = "AccTest_CompileOptionsPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_compile_settings" "acc_test" {
  id         = "AccTest_CompileSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_config_deployment_policy" "acc_test" {
  id         = "AccTest_ConfigDeploymentPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_config_sequence" "acc_test" {
  id         = "AccTest_ConfigSequence"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_conformance_policy" "acc_test" {
  id         = "AccTest_ConformancePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_control_list" "acc_test" {
  id         = "AccTest_ControlList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_cookie_attribute_policy" "acc_test" {
  id         = "AccTest_CookieAttributePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_cors_policy" "acc_test" {
  id         = "AccTest_CORSPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  rule       = [datapower_cors_rule.acc_test.id]
}
resource "datapower_cors_rule" "acc_test" {
  id           = "AccTest_CORSRule"
  app_domain   = datapower_domain.acc_test.app_domain
  allow_origin = ["origin"]
}
resource "datapower_count_monitor" "acc_test" {
  id           = "AccTest_CountMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  measure      = "requests"
  source       = "all"
  max_source_s = 10000
  message_type = datapower_message_type.acc_test.id
}
resource "datapower_crl_fetch" "acc_test" {
}
resource "datapower_crypto_certificate" "acc_test" {
  id         = "AccTest_CryptoCertificate"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.acc_test_server_crt.remote_path
}
resource "datapower_crypto_fw_cred" "acc_test" {
  id          = "AccTest_CryptoFWCred"
  app_domain  = datapower_domain.acc_test.app_domain
  private_key = [datapower_crypto_key.acc_test.id]
}
resource "datapower_crypto_ident_cred" "acc_test" {
  id          = "AccTest_CryptoIdentCred"
  app_domain  = datapower_domain.acc_test.app_domain
  key         = datapower_crypto_key.acc_test.id
  certificate = datapower_crypto_certificate.acc_test.id
}
resource "datapower_crypto_kerberos_kdc" "acc_test" {
  id         = "AccTest_CryptoKerberosKDC"
  app_domain = datapower_domain.acc_test.app_domain
  realm      = "realm"
  server     = "localhost"
}
resource "datapower_crypto_kerberos_keytab" "acc_test" {
  id         = "AccTest_CryptoKerberosKeytab"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.kerberos-keytab.remote_path
}
resource "datapower_crypto_key" "acc_test" {
  id         = "AccTest_CryptoKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.acc_test_server_key.remote_path
  alias      = datapower_password_alias.acc_test.id
}
resource "datapower_crypto_sskey" "acc_test" {
  id         = "AccTest_CryptoSSKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.token-secret.remote_path
}
resource "datapower_crypto_val_cred" "acc_test" {
  id         = "AccTest_CryptoValCred"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_deployment_policy_parameters_binding" "acc_test" {
  id         = "AccTest_DeploymentPolicyParametersBinding"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_distributed_variable" "acc_test" {
  app_domain      = datapower_domain.acc_test.app_domain
  gateway_peering = "default-gateway-peering"
}
resource "datapower_document_crypto_map" "acc_test" {
  id         = "AccTest_DocumentCryptoMap"
  app_domain = datapower_domain.acc_test.app_domain
  operation  = "encrypt"
  xpath      = ["*", ]
}
resource "datapower_domain" "acc_test" {
  app_domain = "acceptance_test"
}
resource "datapower_domain_availability" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_domain_settings" "acc_test" {
  app_domain         = datapower_domain.acc_test.app_domain
  password_treatment = "masked"
}
resource "datapower_duration_monitor" "acc_test" {
  id           = "AccTest_DurationMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  measure      = "messages"
  message_type = datapower_message_type.acc_test.id
}
resource "datapower_ebms2_source_protocol_handler" "acc_test" {
  id            = "AccTest_EBMS2SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_ebms3_source_protocol_handler" "acc_test" {
  id            = "AccTest_EBMS3SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_error_report_settings" "acc_test" {
}
resource "datapower_file_system_usage_monitor" "acc_test" {
  polling_interval = 60
  all_system       = true
}
resource "datapower_filter_action" "acc_test" {
  id         = "AccTest_FilterAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "notify"
}
resource "datapower_forms_login_policy" "acc_test" {
  id                    = "AccTest_FormsLoginPolicy"
  app_domain            = datapower_domain.acc_test.app_domain
  login_form            = "/LoginPage.htm"
  use_cookie_attributes = false
  enable_migration      = false
  error_page            = "/ErrorPage.htm"
  logout_page           = "/LogoutPage.htm"
  default_url           = "/"
  username_field        = "j_username"
  password_field        = "j_password"
  redirect_field        = "originalUrl"
  form_processing_url   = "/j_security_check"
}
resource "datapower_ftp_file_poller_source_protocol_handler" "acc_test" {
  id                       = "AccTest_FTPFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  target_directory         = "ftp://user:password@host:port/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  result_name_pattern      = "../result/$1"
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_ftp_quote_commands" "acc_test" {
  id         = "AccTest_FTPQuoteCommands"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ftp_server_source_protocol_handler" "acc_test" {
  id            = "AccTest_FTPServerSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 21
}
resource "datapower_gateway_peering_manager" "acc_test" {
  app_domain                  = datapower_domain.acc_test.app_domain
  api_connect_gateway_service = "default-gateway-peering"
  rate_limit                  = "default-gateway-peering"
  subscription                = "default-gateway-peering"
  ratelimit_module            = "default-gateway-peering"
}
resource "datapower_git_ops" "acc_test" {
  app_domain        = datapower_domain.acc_test.app_domain
  connection_type   = "https"
  mode              = "read-write"
  commit_identifier = "main"
  remote_location   = "https://github.com/ScottW514/terraform-provider-datapower"
  username          = "gitusername"
  password          = datapower_password_alias.test_acc.id
  tls_valcred       = datapower_crypto_val_cred.test_acc.id
  git_user          = "Git User"
  git_email         = "git@user.domain"
}
resource "datapower_git_ops_template" "acc_test" {
  id         = "AccTest_GitOpsTemplate"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_git_ops_variables" "acc_test" {
}
resource "datapower_graphql_schema_options" "acc_test" {
  id         = "AccTest_GraphQLSchemaOptions"
  app_domain = datapower_domain.acc_test.app_domain
  api        = datapower_api_definition.acc_test.id
}
resource "datapower_gw_script_settings" "acc_test" {
}
resource "datapower_gws_remote_debug" "acc_test" {
  local_port    = 9229
  local_address = "0.0.0.0"
}
resource "datapower_host_alias" "acc_test" {
  id         = "AccTest_HostAlias"
  ip_address = "10.10.10.10"
}
resource "datapower_http_input_conversion_map" "acc_test" {
  id         = "AccTest_HTTPInputConversionMap"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_http_source_protocol_handler" "acc_test" {
  id            = "AccTest_HTTPSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 8088
}
resource "datapower_http_user_agent" "acc_test" {
  id         = "AccTest_HTTPUserAgent"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_https_source_protocol_handler" "acc_test" {
  id                     = "AccTest_HTTPSSourceProtocolHandler"
  app_domain             = datapower_domain.acc_test.app_domain
  ssl_server_config_type = "server"
  ssl_server             = datapower_ssl_server_profile.acc_test.id
}
resource "datapower_import_package" "acc_test" {
  id                = "AccTest_ImportPackage"
  app_domain        = datapower_domain.acc_test.app_domain
  url               = "http://localhost/config.zip"
  import_format     = "ZIP"
  overwrite_files   = true
  overwrite_objects = true
}
resource "datapower_include_config" "acc_test" {
  id         = "AccTest_IncludeConfig"
  app_domain = datapower_domain.acc_test.app_domain
  url        = "http://localhost/config.zip"
}
resource "datapower_interop_service" "acc_test" {
}
resource "datapower_jose_recipient_identifier" "acc_test" {
  id         = "AccTest_JOSERecipientIdentifier"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "key"
  key        = datapower_crypto_key.acc_test.id
  header_param = [{
    header_value = "VALUE"
  }]
}
resource "datapower_jose_signature_identifier" "acc_test" {
  id          = "AccTest_JOSESignatureIdentifier"
  app_domain  = datapower_domain.acc_test.app_domain
  type        = "certificate"
  certificate = datapower_crypto_certificate.acc_test.id
  header_param = [{
    header_value = "VALUE"
  }]
}
resource "datapower_json_settings" "acc_test" {
  id         = "AccTest_JSONSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_jwe_header" "acc_test" {
  id         = "AccTest_JWEHeader"
  app_domain = datapower_domain.acc_test.app_domain
  recipient  = datapower_jwe_recipient.acc_test.id
}
resource "datapower_jwe_recipient" "acc_test" {
  id          = "AccTest_JWERecipient"
  app_domain  = datapower_domain.acc_test.app_domain
  algorithm   = "RSA1_5"
  certificate = datapower_crypto_certificate.acc_test.id
}
resource "datapower_jws_signature" "acc_test" {
  id         = "AccTest_JWSSignature"
  app_domain = datapower_domain.acc_test.app_domain
  algorithm  = "RS256"
  key        = datapower_crypto_key.acc_test.id
}
resource "datapower_ldap_connection_pool" "acc_test" {
  id         = "AccTest_LDAPConnectionPool"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ldap_search_parameters" "acc_test" {
  id         = "AccTest_LDAPSearchParameters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_load_balancer_group" "acc_test" {
  id            = "AccTest_LoadBalancerGroup"
  app_domain    = datapower_domain.acc_test.app_domain
  algorithm     = "round-robin"
  retrieve_info = false
  damp          = 120
}
resource "datapower_log_label" "acc_test" {
  id         = "AccTest_LogLabel"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_log_target" "acc_test" {
  id         = "AccTest_LogTarget"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "file"
  local_file = "logtemp:///AccTest_LogTarget.log"
}
resource "datapower_matching" "acc_test" {
  id         = "AccTest_Matching"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_mcf_custom_rule" "acc_test" {
  id                = "AccTest_MCFCustomRule"
  app_domain        = datapower_domain.acc_test.app_domain
  custom_rule_name  = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}
resource "datapower_mcf_http_header" "acc_test" {
  id         = "AccTest_MCFHttpHeader"
  app_domain = datapower_domain.acc_test.app_domain
  http_name  = "HEADERNAME"
  http_value = "HEADERVALUE"
}
resource "datapower_mcf_http_method" "acc_test" {
  id          = "AccTest_MCFHttpMethod"
  app_domain  = datapower_domain.acc_test.app_domain
  http_method = "GET"
}
resource "datapower_mcf_http_url" "acc_test" {
  id                  = "AccTest_MCFHttpURL"
  app_domain          = datapower_domain.acc_test.app_domain
  http_url_expression = "*"
}
resource "datapower_mcf_xpath" "acc_test" {
  id               = "AccTest_MCFXPath"
  app_domain       = datapower_domain.acc_test.app_domain
  xpath_expression = "*"
  xpath_value      = "value"
}
resource "datapower_message_content_filters" "acc_test" {
  id         = "AccTest_MessageContentFilters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_message_matching" "acc_test" {
  id         = "AccTest_MessageMatching"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_message_type" "acc_test" {
  id         = "AccTest_MessageType"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_mgmt_interface" "acc_test" {
  local_port    = 5550
  local_address = "0.0.0.0"
}
resource "datapower_mpgw_error_action" "acc_test" {
  id         = "AccTest_MPGWErrorAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "static"
  local_url  = "store:///schemas/XMLSchema.dtd"
}
resource "datapower_mpgw_error_handling_policy" "acc_test" {
  id         = "AccTest_MPGWErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  policy_maps = [{
    match  = "__default-accept-service-providers__"
    action = datapower_mpgw_error_action.acc_test.id
  }]
}
resource "datapower_mq_manager" "acc_test" {
  id            = "AccTest_MQManager"
  app_domain    = datapower_domain.acc_test.app_domain
  host_name     = "localhost"
  cache_timeout = 60
  xml_manager   = "default"
}
resource "datapower_mq_manager_group" "acc_test" {
  id                    = "AccTest_MQManagerGroup"
  app_domain            = datapower_domain.acc_test.app_domain
  primary_queue_manager = datapower_mq_manager.acc_test.id
}
resource "datapower_mqv9_plus_mft_source_protocol_handler" "acc_test" {
  id            = "AccTest_MQv9PlusMFTSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mq_manager.acc_test.id
  get_queue     = "queue"
}
resource "datapower_mqv9_plus_source_protocol_handler" "acc_test" {
  id            = "AccTest_MQv9PlusSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mq_manager.acc_test.id
  get_queue     = "queue"
}
resource "datapower_mtom_policy" "acc_test" {
  id         = "AccTest_MTOMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_multi_protocol_gateway" "acc_test" {
  id                       = "AccTest_MultiProtocolGateway"
  app_domain               = datapower_domain.acc_test.app_domain
  type                     = "static-backend"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
resource "datapower_name_value_profile" "acc_test" {
  id            = "AccTest_NameValueProfile"
  app_domain    = datapower_domain.acc_test.app_domain
  default_fixup = "strip"
}
resource "datapower_nfs_client_settings" "acc_test" {
}
resource "datapower_nfs_dynamic_mounts" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_nfs_file_poller_source_protocol_handler" "acc_test" {
  id                       = "AccTest_NFSFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  target_directory         = "dpnfs://static-mount-name/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_nfs_static_mount" "acc_test" {
  id         = "AccTest_NFSStaticMount"
  app_domain = datapower_domain.acc_test.app_domain
  remote     = "url://test"
}
resource "datapower_oauth_provider_settings" "acc_test" {
  id                 = "AccTest_OAuthProviderSettings"
  app_domain         = datapower_domain.acc_test.app_domain
  provider_type      = "native"
  api_c_token_secret = datapower_crypto_sskey.acc_test.id
}
resource "datapower_oauth_supported_client" "acc_test" {
  id                       = "AccTest_OAuthSupportedClient"
  app_domain               = datapower_domain.acc_test.app_domain
  oauth_role               = { "azsvr" : true }
  az_grant                 = { "code" : true }
  generate_client_secret   = false
  client_secret_wo         = "secret"
  client_secret_wo_version = 1
  redirect_uri             = ["^https://example.com/redirect$"]
  scope                    = "*"
  token_secret             = datapower_crypto_sskey.acc_test.id
}
resource "datapower_oauth_supported_client_group" "acc_test" {
  id         = "AccTest_OAuthSupportedClientGroup"
  app_domain = datapower_domain.acc_test.app_domain
  client     = [datapower_oauth_supported_client.acc_test.id]
}
resource "datapower_odr" "acc_test" {
  odr_server_name = "dp_set"
}
resource "datapower_odr_connector_group" "acc_test" {
  id          = "AccTest_ODRConnectorGroup"
  xml_manager = "default"
}
resource "datapower_open_telemetry" "acc_test" {
  id         = "AccTest_OpenTelemetry"
  app_domain = datapower_domain.acc_test.app_domain
  exporter   = datapower_open_telemetry_exporter.acc_test.id
  sampler    = datapower_open_telemetry_sampler.acc_test.id
}
resource "datapower_open_telemetry_exporter" "acc_test" {
  id         = "AccTest_OpenTelemetryExporter"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "http"
  host_name  = "localhost"
  port       = 4318
  processor  = "batch"
}
resource "datapower_open_telemetry_sampler" "acc_test" {
  id           = "AccTest_OpenTelemetrySampler"
  app_domain   = datapower_domain.acc_test.app_domain
  parent_based = true
  type         = "always-on"
}
resource "datapower_operation_rate_limit" "acc_test" {
  id         = "AccTest_OperationRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  operation  = datapower_api_operation.acc_test.id
  rate_limit = [{ "name" : "RateLimit", "rate" : "1000" }]
}
resource "datapower_parse_settings" "acc_test" {
  id         = "AccTest_ParseSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_password_alias" "acc_test" {
  id                  = "AccTest_PasswordAlias"
  app_domain          = datapower_domain.acc_test.app_domain
  password_wo         = "password"
  password_wo_version = 1
}
resource "datapower_peer_group" "acc_test" {
  id         = "AccTest_PeerGroup"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "slm"
  url        = ["http://localhost"]
}
resource "datapower_policy_attachments" "acc_test" {
  id                   = "AccTest_PolicyAttachments"
  app_domain           = datapower_domain.acc_test.app_domain
  enforcement_mode     = "enforce"
  policy_references    = false
  sla_enforcement_mode = "allow-if-no-sla"
}
resource "datapower_policy_parameters" "acc_test" {
  id         = "AccTest_PolicyParameters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_pop_poller_source_protocol_handler" "acc_test" {
  id                    = "AccTest_POPPollerSourceProtocolHandler"
  app_domain            = datapower_domain.acc_test.app_domain
  mail_server           = "localhost"
  port                  = 8888
  conn_security         = "none"
  account               = "account"
  password_alias        = datapower_password_alias.acc_test.id
  delay_between_polls   = 300
  max_messages_per_poll = 5
}
resource "datapower_probe" "acc_test" {
  app_domain      = datapower_domain.acc_test.app_domain
  max_records     = 1000
  gateway_peering = "default-gateway-peering"
}
resource "datapower_processing_metadata" "acc_test" {
  id         = "AccTest_ProcessingMetadata"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_quota_enforcement_server" "acc_test" {
  server_port  = 16379
  monitor_port = 26379
}
resource "datapower_radius_settings" "acc_test" {
}
resource "datapower_raid_volume" "acc_test" {
}
resource "datapower_rate_limit_configuration" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_rate_limit_definition" "acc_test" {
  id         = "AccTest_RateLimitDefinition"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "rate"
  rate       = 1000
}
resource "datapower_rate_limit_definition_group" "acc_test" {
  id               = "AccTest_RateLimitDefinitionGroup"
  app_domain       = datapower_domain.acc_test.app_domain
  update_on_exceed = "all"
}
resource "datapower_rbm_settings" "acc_test" {
  au_method                   = "local"
  au_cache_allow              = "absolute"
  mc_method                   = "local"
  min_password_length         = 6
  require_mixed_case          = false
  require_digit               = false
  require_non_alpha_numeric   = false
  disallow_username_substring = false
  do_password_aging           = false
  do_password_history         = false
  cli_timeout                 = 0
  max_failed_login            = 0
}
resource "datapower_saml_attributes" "acc_test" {
  id         = "AccTest_SAMLAttributes"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_schema_exception_map" "acc_test" {
  id                  = "AccTest_SchemaExceptionMap"
  app_domain          = datapower_domain.acc_test.app_domain
  original_schema_url = "http://localhost"
  schema_exception_rules = [{
    xpath          = "*"
    exception_type = "AllowEncrypted"
  }]
}
resource "datapower_sftp_file_poller_source_protocol_handler" "acc_test" {
  id                       = "AccTest_SFTPFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  ssh_client_connection    = datapower_ssh_client_profile.acc_test.id
  target_directory         = "/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_slm_action" "acc_test" {
  id         = "AccTest_SLMAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "log-only"
}
resource "datapower_slm_cred_class" "acc_test" {
  id         = "AccTest_SLMCredClass"
  app_domain = datapower_domain.acc_test.app_domain
  cred_type  = "aaa-mapped-credential"
}
resource "datapower_slm_policy" "acc_test" {
  id         = "AccTest_SLMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_slm_rsrc_class" "acc_test" {
  id         = "AccTest_SLMRsrcClass"
  app_domain = datapower_domain.acc_test.app_domain
  rsrc_type  = "aaa-mapped-resource"
}
resource "datapower_slm_schedule" "acc_test" {
  id         = "AccTest_SLMSchedule"
  app_domain = datapower_domain.acc_test.app_domain
  start_time = "12:34:00"
  duration   = 1440
}
resource "datapower_smtp_server_connection" "acc_test" {
  id               = "AccTest_SMTPServerConnection"
  app_domain       = datapower_domain.acc_test.app_domain
  mail_server_host = "localhost"
  mail_server_port = 25
}
resource "datapower_snmp_settings" "acc_test" {
  local_port     = 161
  security_level = "authPriv"
  access_level   = "read-only"
}
resource "datapower_soap_header_disposition" "acc_test" {
  id         = "AccTest_SOAPHeaderDisposition"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_social_login_policy" "acc_test" {
  id                      = "AccTest_SocialLoginPolicy"
  app_domain              = datapower_domain.acc_test.app_domain
  client_id               = "client_id"
  client_secret           = "client_secret"
  client_ssl_profile      = datapower_ssl_client_profile.acc_test.id
  social_provider         = "google"
  provider_az_endpoint    = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token      = false
}
resource "datapower_sql_data_source" "acc_test" {
  id               = "AccTest_SQLDataSource"
  app_domain       = datapower_domain.acc_test.app_domain
  database         = "Oracle"
  username         = "username"
  password_alias   = datapower_password_alias.acc_test.id
  data_source_id   = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection   = 10
  connect_timeout  = 15
  query_timeout    = 30
  idle_timeout     = 180
}
resource "datapower_ssh_client_profile" "acc_test" {
  id            = "AccTest_SSHClientProfile"
  app_domain    = datapower_domain.acc_test.app_domain
  user_name     = "someuser"
  profile_usage = "sftp"
}
resource "datapower_ssh_domain_client_profile" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ssh_server_profile" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ssh_server_source_protocol_handler" "acc_test" {
  id                = "AccTest_SSHServerSourceProtocolHandler"
  app_domain        = datapower_domain.acc_test.app_domain
  local_address     = "0.0.0.0"
  local_port        = 22
  default_directory = "/"
}
resource "datapower_ssh_service" "acc_test" {
  local_port    = 22
  local_address = "0.0.0.0"
}
resource "datapower_ssl_client_profile" "acc_test" {
  id                   = "AccTest_SSLClientProfile"
  app_domain           = datapower_domain.acc_test.app_domain
  ciphers              = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
  validate_server_cert = false
}
resource "datapower_ssl_proxy_service" "acc_test" {
  id             = "AccTest_SSLProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  local_port     = 4521
  remote_address = "10.10.10.10"
  remote_port    = 9999
  ssl_server     = datapower_ssl_server_profile.acc_test.id
  local_address  = "0.0.0.0"
}
resource "datapower_ssl_server_profile" "acc_test" {
  id         = "AccTest_SSLServerProfile"
  app_domain = datapower_domain.acc_test.app_domain
  ciphers    = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
  idcred     = datapower_crypto_ident_cred.acc_test.id
}
resource "datapower_ssl_sni_mapping" "acc_test" {
  id         = "AccTest_SSLSNIMapping"
  app_domain = datapower_domain.acc_test.app_domain
  sni_mapping = [{
    host_name_wildmat = "hostname_wildmat"
    ssl_server        = datapower_ssl_server_profile.acc_test.id
  }]
}
resource "datapower_ssl_sni_server_profile" "acc_test" {
  id                 = "AccTest_SSLSNIServerProfile"
  app_domain         = datapower_domain.acc_test.app_domain
  sni_server_mapping = datapower_ssl_sni_mapping.acc_test.id
}
resource "datapower_stateless_tcp_source_protocol_handler" "acc_test" {
  id            = "AccTest_StatelessTCPSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 4000
}
resource "datapower_statistics" "acc_test" {
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_style_policy" "acc_test" {
  id         = "AccTest_StylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_style_policy_action" "acc_test" {
  id            = "AccTest_StylePolicyAction"
  app_domain    = datapower_domain.acc_test.app_domain
  type          = "xform"
  input         = "INPUT"
  transform     = "tfile"
  output        = "OUTPUT"
  named_inputs  = null
  named_outputs = null
}
resource "datapower_style_policy_rule" "acc_test" {
  id            = "AccTest_StylePolicyRule"
  app_domain    = datapower_domain.acc_test.app_domain
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}
resource "datapower_system_settings" "acc_test" {
}
resource "datapower_tcp_proxy_service" "acc_test" {
  id             = "AccTest_TCPProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  local_port     = 6158
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}
resource "datapower_throttler" "acc_test" {
  throttle_at          = 0
  terminate_at         = 0
  temp_fs_throttle_at  = 0
  temp_fs_terminate_at = 0
  qname_warn_at        = 10
  timeout              = 30
}
resource "datapower_time_settings" "acc_test" {
  local_time_zone = "EST5EDT"
}
resource "datapower_url_map" "acc_test" {
  id         = "AccTest_URLMap"
  app_domain = datapower_domain.acc_test.app_domain
  url_map_rule = [{
    pattern = "https://www.company.com/XML/stylesheets/*"
  }]
}
resource "datapower_url_refresh_policy" "acc_test" {
  id         = "AccTest_URLRefreshPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  url_refresh_rule = [{
    url_map = "default-attempt-stream-all"
  }]
}
resource "datapower_url_rewrite_policy" "acc_test" {
  id         = "AccTest_URLRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_user" "acc_test" {
  id                  = "AccTest_User"
  password_wo         = "Password$123"
  password_wo_version = 1
  access_level        = "group-defined"
  group_name          = datapower_user_group.acc_test.id
  snmp_creds          = null
  hashed_snmp_creds   = null
}
resource "datapower_user_group" "acc_test" {
  id              = "AccTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}
resource "datapower_visibility_list" "acc_test" {
  id         = "AccTest_VisibilityList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_wcc_service" "acc_test" {
  id                = "AccTest_WCCService"
  app_domain        = datapower_domain.acc_test.app_domain
  odc_info_hostname = "example.com"
  odc_info_port     = 1024
  time_interval     = 10
}
resource "datapower_web_app_error_handling_policy" "acc_test" {
  id         = "AccTest_WebAppErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_web_app_fw" "acc_test" {
  id                       = "AccTest_WebAppFW"
  app_domain               = datapower_domain.acc_test.app_domain
  front_side               = [{ "LocalAddress" : "0.0.0.0" }]
  remote_address           = "10.10.10.10"
  style_policy             = datapower_app_security_policy.acc_test.id
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
resource "datapower_web_app_request" "acc_test" {
  id         = "AccTest_WebAppRequest"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_web_app_response" "acc_test" {
  id          = "AccTest_WebAppResponse"
  app_domain  = datapower_domain.acc_test.app_domain
  policy_type = "admission"
}
resource "datapower_web_app_session_policy" "acc_test" {
  id            = "AccTest_WebAppSessionPolicy"
  app_domain    = datapower_domain.acc_test.app_domain
  start_matches = "__default-accept-service-providers__"
}
resource "datapower_web_b2b_viewer" "acc_test" {
  local_port    = 9091
  idle_timeout  = 600
  local_address = "0.0.0.0"
}
resource "datapower_web_gui" "acc_test" {
  local_port             = 9090
  save_config_overwrites = true
  idle_timeout           = 0
  local_address          = "0.0.0.0"
}
resource "datapower_web_service_monitor" "acc_test" {
  id           = "AccTest_WebServiceMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  wsdl_url     = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}
resource "datapower_web_services_agent" "acc_test" {
  app_domain    = datapower_domain.acc_test.app_domain
  max_records   = 3000
  max_memory_kb = 64000
  capture_mode  = "faults"
}
resource "datapower_web_token_service" "acc_test" {
  id          = "AccTest_WebTokenService"
  app_domain  = datapower_domain.acc_test.app_domain
  xml_manager = "default"
  front_side = [{
    local_address = "0.0.0.0"
    local_port    = 8888
  }]
  style_policy             = "default"
  front_timeout            = 120
  front_persistent_timeout = 180
}
resource "datapower_ws_endpoint_rewrite_policy" "acc_test" {
  id         = "AccTest_WSEndpointRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ws_gateway" "acc_test" {
  id                       = "AccTest_WSGateway"
  app_domain               = datapower_domain.acc_test.app_domain
  type                     = "static-from-wsdl"
  style_policy             = "default"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
resource "datapower_ws_style_policy" "acc_test" {
  id         = "AccTest_WSStylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ws_style_policy_rule" "acc_test" {
  id            = "AccTest_WSStylePolicyRule"
  app_domain    = datapower_domain.acc_test.app_domain
  actions       = [datapower_style_policy_action.acc_test.id]
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}
resource "datapower_wsrr_saved_search_subscription" "acc_test" {
  id                = "AccTest_WSRRSavedSearchSubscription"
  app_domain        = datapower_domain.acc_test.app_domain
  server            = datapower_wsrr_server.acc_test.id
  saved_search_name = "ResTestsaved_search"
}
resource "datapower_wsrr_server" "acc_test" {
  id         = "AccTest_WSRRServer"
  app_domain = datapower_domain.acc_test.app_domain
  soap_url   = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}
resource "datapower_wsrr_subscription" "acc_test" {
  id          = "AccTest_WSRRSubscription"
  app_domain  = datapower_domain.acc_test.app_domain
  server      = datapower_wsrr_server.acc_test.id
  namespace   = "namespace"
  object_type = "wsdl"
  object_name = "ResTestobject"
}
resource "datapower_wxs_grid" "acc_test" {
  id             = "AccTest_WXSGrid"
  app_domain     = datapower_domain.acc_test.app_domain
  collective     = datapower_load_balancer_group.acc_test.id
  grid           = "gridname"
  user_name      = "username"
  password_alias = datapower_password_alias.acc_test.id
  timeout        = 1000
}
resource "datapower_xacml_pdp" "acc_test" {
  id             = "AccTest_XACMLPDP"
  app_domain     = datapower_domain.acc_test.app_domain
  general_policy = "http://test/uri"
}
resource "datapower_xml_firewall_service" "acc_test" {
  id            = "AccTest_XMLFirewallService"
  app_domain    = datapower_domain.acc_test.app_domain
  type          = "dynamic-backend"
  xml_manager   = "default"
  local_port    = 7575
  local_address = "0.0.0.0"
}
resource "datapower_xml_manager" "acc_test" {
  id             = "AccTest_XMLManager"
  app_domain     = datapower_domain.acc_test.app_domain
  ldap_conn_pool = datapower_ldap_connection_pool.acc_test.id
}
resource "datapower_xpath_routing_map" "acc_test" {
  id         = "AccTest_XPathRoutingMap"
  app_domain = datapower_domain.acc_test.app_domain
  xpath_routing_rules = [{
    xpath = "*"
    host  = "localhost"
    port  = 8888
    ssl   = false
  }]
}
resource "datapower_xsl_coproc_service" "acc_test" {
  id            = "AccTest_XSLCoprocService"
  app_domain    = datapower_domain.acc_test.app_domain
  local_port    = 8811
  xml_manager   = "default"
  local_address = "0.0.0.0"
}
resource "datapower_xsl_proxy_service" "acc_test" {
  id             = "AccTest_XSLProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  type           = "static-backend"
  xml_manager    = "default"
  local_port     = 8922
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}
resource "datapower_xtc_protocol_handler" "acc_test" {
  id             = "AccTest_XTCProtocolHandler"
  app_domain     = datapower_domain.acc_test.app_domain
  local_address  = "0.0.0.0"
  local_port     = 3000
  remote_address = "10.10.10.10"
  remote_port    = 12000
}
resource "datapower_zos_nss_client" "acc_test" {
  id             = "AccTest_ZosNSSClient"
  app_domain     = datapower_domain.acc_test.app_domain
  remote_address = "remote.host"
  remote_port    = 443
  client_id      = "client_id"
  system_name    = "sysname"
  user_name      = "username"
  password_alias = datapower_password_alias.acc_test.id
  ssl_client     = datapower_ssl_client_profile.acc_test.id
}