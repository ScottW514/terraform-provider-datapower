terraform {
  required_providers {
    datapower = {
      source = "scottw514/datapower"
    }
  }
}
resource "datapower_aaajwtgenerator" "acc_test" {
  id         = "AccTest_AAAJWTGenerator"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_aaajwtvalidator" "acc_test" {
  id             = "AccTest_AAAJWTValidator"
  app_domain     = datapower_domain.acc_test.app_domain
  username_claim = "sub"
}
resource "datapower_aaapolicy" "acc_test" {
  id         = "AccTest_AAAPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_accesscontrollist" "acc_test" {
  id         = "AccTest_AccessControlList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_accessprofile" "acc_test" {
  id         = "AccTest_AccessProfile"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_amqpbroker" "acc_test" {
  id            = "AccTest_AMQPBroker"
  app_domain    = datapower_domain.acc_test.app_domain
  host_name     = "host.name"
  port          = 5672
  xml_manager   = "default"
  authorization = "none"
}
resource "datapower_amqpsourceprotocolhandler" "acc_test" {
  id         = "AccTest_AMQPSourceProtocolHandler"
  app_domain = datapower_domain.acc_test.app_domain
  broker     = datapower_amqpbroker.acc_test.id
  from       = "amqpfrom"
}
resource "datapower_analyticsendpoint" "acc_test" {
  id                     = "AccTest_AnalyticsEndpoint"
  app_domain             = datapower_domain.acc_test.app_domain
  analytics_server_url   = "https://localhost"
  ssl_client             = datapower_sslclientprofile.acc_test.id
  max_records            = 1024
  max_records_memory_kb  = 512
  max_delivery_memory_mb = 512
}
resource "datapower_apiapplicationtype" "acc_test" {
  id         = "AccTest_APIApplicationType"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apiauthurlregistry" "acc_test" {
  id         = "AccTest_APIAuthURLRegistry"
  app_domain = datapower_domain.acc_test.app_domain
  auth_url   = "http://localhost"
}
resource "datapower_apiclientidentification" "acc_test" {
  id         = "AccTest_APIClientIdentification"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apicollection" "acc_test" {
  id         = "AccTest_APICollection"
  app_domain = datapower_domain.acc_test.app_domain
  org_id     = "orgid"
  org_name   = "orgname"
  routing_prefix = [{
    type = "host"
  }]
  api_processing_rule = "default-api-rule"
  api_error_rule      = "default-api-error-rule"
  plan                = [datapower_apiplan.acc_test.id]
}
resource "datapower_apicors" "acc_test" {
  id         = "AccTest_APICORS"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apidefinition" "acc_test" {
  id            = "AccTest_APIDefinition"
  app_domain    = datapower_domain.acc_test.app_domain
  base_path     = "/"
  path          = [datapower_apipath.acc_test.id]
  content       = "activity"
  error_content = "payload"
}
resource "datapower_apiexecute" "acc_test" {
  id         = "AccTest_APIExecute"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apifinal" "acc_test" {
  id         = "AccTest_APIFinal"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apigateway" "acc_test" {
  id                       = "AccTest_APIGateway"
  app_domain               = datapower_domain.acc_test.app_domain
  front_protocol           = [datapower_httpsourceprotocolhandler.acc_test.id]
  front_timeout            = 120
  front_persistent_timeout = 180
}
resource "datapower_apildapregistry" "acc_test" {
  id                     = "AccTest_APILDAPRegistry"
  app_domain             = datapower_domain.acc_test.app_domain
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = datapower_ldapsearchparameters.acc_test.id
}
resource "datapower_apioperation" "acc_test" {
  id         = "AccTest_APIOperation"
  app_domain = datapower_domain.acc_test.app_domain
  method     = "GET"
}
resource "datapower_apipath" "acc_test" {
  id         = "AccTest_APIPath"
  app_domain = datapower_domain.acc_test.app_domain
  path       = "/"
}
resource "datapower_apiplan" "acc_test" {
  id               = "AccTest_APIPlan"
  app_domain       = datapower_domain.acc_test.app_domain
  api              = [datapower_apidefinition.acc_test.id]
  rate_limit_scope = "per-application"
}
resource "datapower_apiresult" "acc_test" {
  id         = "AccTest_APIResult"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apirouting" "acc_test" {
  id         = "AccTest_APIRouting"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apirule" "acc_test" {
  id         = "AccTest_APIRule"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apischema" "acc_test" {
  id          = "AccTest_APISchema"
  app_domain  = datapower_domain.acc_test.app_domain
  json_schema = "http://localhost/json"
}
resource "datapower_apisecurity" "acc_test" {
  id         = "AccTest_APISecurity"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apisecurityapikey" "acc_test" {
  id         = "AccTest_APISecurityAPIKey"
  app_domain = datapower_domain.acc_test.app_domain
  where      = "header"
  type       = "id"
  key_name   = "keyname"
}
resource "datapower_apisecuritybasicauth" "acc_test" {
  id            = "AccTest_APISecurityBasicAuth"
  app_domain    = datapower_domain.acc_test.app_domain
  user_registry = datapower_apiauthurlregistry.acc_test.id
}
resource "datapower_apisecurityhttpscheme" "acc_test" {
  id         = "AccTest_APISecurityHTTPScheme"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_apisecurityoauth" "acc_test" {
  id              = "AccTest_APISecurityOAuth"
  app_domain      = datapower_domain.acc_test.app_domain
  o_auth_provider = datapower_oauthprovidersettings.acc_test.id
  o_auth_flow     = "implicit"
}
resource "datapower_apisecurityoauthreq" "acc_test" {
  id                      = "AccTest_APISecurityOAuthReq"
  app_domain              = datapower_domain.acc_test.app_domain
  api_security_o_auth_def = datapower_apisecurityoauth.acc_test.id
}
resource "datapower_apisecurityrequirement" "acc_test" {
  id         = "AccTest_APISecurityRequirement"
  app_domain = datapower_domain.acc_test.app_domain
  security   = [datapower_apisecurityapikey.acc_test.id]
}
resource "datapower_appsecuritypolicy" "acc_test" {
  id         = "AccTest_AppSecurityPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  request_maps = [{
    match = "__default-accept-service-providers__"
    rule  = datapower_webapprequest.acc_test.id
  }]
  response_maps = [{
    match = "__default-accept-service-providers__"
    rule  = datapower_webappresponse.acc_test.id
  }]
}
resource "datapower_as1pollersourceprotocolhandler" "acc_test" {
  id                    = "AccTest_AS1PollerSourceProtocolHandler"
  app_domain            = datapower_domain.acc_test.app_domain
  mail_server           = "smtp.host"
  port                  = 25
  conn_security         = "none"
  account               = "account"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}
resource "datapower_as2proxysourceprotocolhandler" "acc_test" {
  id                        = "AccTest_AS2ProxySourceProtocolHandler"
  app_domain                = datapower_domain.acc_test.app_domain
  local_address             = "0.0.0.0"
  local_port                = 80
  remote_address            = "10.10.10.10"
  remote_port               = 8888
  remote_connection_timeout = 60
  xml_manager               = "default"
}
resource "datapower_as2sourceprotocolhandler" "acc_test" {
  id            = "AccTest_AS2SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_as3sourceprotocolhandler" "acc_test" {
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
resource "datapower_assemblyactionclientsecurity" "acc_test" {
  id                         = "AccTest_AssemblyActionClientSecurity"
  app_domain                 = datapower_domain.acc_test.app_domain
  id_name                    = "idname"
  secret_name                = "secretname"
  authenticate_client_method = "native"
}
resource "datapower_assemblyactionextract" "acc_test" {
  id         = "AccTest_AssemblyActionExtract"
  app_domain = datapower_domain.acc_test.app_domain
  extract = [{
    capture = "capture"
  }]
}
resource "datapower_assemblyactionfunctioncall" "acc_test" {
  id            = "AccTest_AssemblyActionFunctionCall"
  app_domain    = datapower_domain.acc_test.app_domain
  function_call = "default-func-global"
}
resource "datapower_assemblyactiongatewayscript" "acc_test" {
  id         = "AccTest_AssemblyActionGatewayScript"
  app_domain = datapower_domain.acc_test.app_domain
  source     = "gsfile"
}
resource "datapower_assemblyactiongraphqlcostanalysis" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLCostAnalysis"
  app_domain = datapower_domain.acc_test.app_domain
  target     = "targetquery"
}
resource "datapower_assemblyactiongraphqlexecute" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLExecute"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactiongraphqlintrospect" "acc_test" {
  id         = "AccTest_AssemblyActionGraphQLIntrospect"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionhtmlpage" "acc_test" {
  id         = "AccTest_AssemblyActionHtmlPage"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactioninvoke" "acc_test" {
  id           = "AccTest_AssemblyActionInvoke"
  app_domain   = datapower_domain.acc_test.app_domain
  url          = "https://localhost"
  method       = "Keep"
  backend_type = "detect"
  cache_type   = "Protocol"
}
resource "datapower_assemblyactionjson2xml" "acc_test" {
  id         = "AccTest_AssemblyActionJson2Xml"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionjwtgenerate" "acc_test" {
  id              = "AccTest_AssemblyActionJWTGenerate"
  app_domain      = datapower_domain.acc_test.app_domain
  issuer_claim    = "iss.claim"
  validity_period = 3600
}
resource "datapower_assemblyactionjwtvalidate" "acc_test" {
  id            = "AccTest_AssemblyActionJWTValidate"
  app_domain    = datapower_domain.acc_test.app_domain
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}
resource "datapower_assemblyactionlog" "acc_test" {
  id         = "AccTest_AssemblyActionLog"
  app_domain = datapower_domain.acc_test.app_domain
  mode       = "gather-only"
}
resource "datapower_assemblyactionmap" "acc_test" {
  id         = "AccTest_AssemblyActionMap"
  app_domain = datapower_domain.acc_test.app_domain
  location   = "local:///file"
}
resource "datapower_assemblyactionoauth" "acc_test" {
  id                                 = "AccTest_AssemblyActionOAuth"
  app_domain                         = datapower_domain.acc_test.app_domain
  o_auth_provider_settings_reference = { "Default" : datapower_oauthprovidersettings.acc_test.id }
}
resource "datapower_assemblyactionparse" "acc_test" {
  id         = "AccTest_AssemblyActionParse"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionratelimit" "acc_test" {
  id         = "AccTest_AssemblyActionRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  source     = "plan-default"
}
resource "datapower_assemblyactionratelimitinfo" "acc_test" {
  id         = "AccTest_AssemblyActionRateLimitInfo"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionredact" "acc_test" {
  id         = "AccTest_AssemblyActionRedact"
  app_domain = datapower_domain.acc_test.app_domain
  redact = [{
    path = "path"
  }]
}
resource "datapower_assemblyactionsetvar" "acc_test" {
  id         = "AccTest_AssemblyActionSetVar"
  app_domain = datapower_domain.acc_test.app_domain
  variable = [{
    name  = "varname"
    value = "value"
  }]
}
resource "datapower_assemblyactionthrow" "acc_test" {
  id         = "AccTest_AssemblyActionThrow"
  app_domain = datapower_domain.acc_test.app_domain
  error_id   = "errorid"
}
resource "datapower_assemblyactionusersecurity" "acc_test" {
  id                      = "AccTest_AssemblyActionUserSecurity"
  app_domain              = datapower_domain.acc_test.app_domain
  factor_id               = "default"
  extract_identity_method = "basic"
  user_auth_method        = "user-registry"
  user_registry           = datapower_apiauthurlregistry.acc_test.id
  user_az_method          = "authenticated"
}
resource "datapower_assemblyactionvalidate" "acc_test" {
  id         = "AccTest_AssemblyActionValidate"
  app_domain = datapower_domain.acc_test.app_domain
  schema     = datapower_apischema.acc_test.id
}
resource "datapower_assemblyactionwebsocketupgrade" "acc_test" {
  id         = "AccTest_AssemblyActionWebSocketUpgrade"
  app_domain = datapower_domain.acc_test.app_domain
  url        = "https://localhost"
}
resource "datapower_assemblyactionwsdl" "acc_test" {
  id         = "AccTest_AssemblyActionWSDL"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionxml2json" "acc_test" {
  id         = "AccTest_AssemblyActionXml2Json"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblyactionxslt" "acc_test" {
  id         = "AccTest_AssemblyActionXSLT"
  app_domain = datapower_domain.acc_test.app_domain
  stylesheet = "local:///stylesheet"
}
resource "datapower_assemblyfunction" "acc_test" {
  id         = "AccTest_AssemblyFunction"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_assemblylogicoperationswitch" "acc_test" {
  id         = "AccTest_AssemblyLogicOperationSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = [{
    execute      = "default-api-rule"
    operation_id = "operationid"
  }]
}
resource "datapower_assemblylogicswitch" "acc_test" {
  id         = "AccTest_AssemblyLogicSwitch"
  app_domain = datapower_domain.acc_test.app_domain
  case = [{
    condition = "condition"
    execute   = "default-api-rule"
  }]
}
resource "datapower_b2bcpa" "acc_test" {
  id         = "AccTest_B2BCPA"
  app_domain = datapower_domain.acc_test.app_domain
  cpa_id     = "cpaid"
}
resource "datapower_b2bcpacollaboration" "acc_test" {
  id            = "AccTest_B2BCPACollaboration"
  app_domain    = datapower_domain.acc_test.app_domain
  internal_role = "internal"
  external_role = "external"
  service       = "service"
  actions = [{
    name           = "cpacollaborationactionname"
    value          = "value"
    capability     = "cansend"
    sender_setting = datapower_b2bcpasendersetting.acc_test.id
  }]
}
resource "datapower_b2bcpareceiversetting" "acc_test" {
  id         = "AccTest_B2BCPAReceiverSetting"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_b2bcpasendersetting" "acc_test" {
  id                = "AccTest_B2BCPASenderSetting"
  app_domain        = datapower_domain.acc_test.app_domain
  dest_endpoint_url = "ebms2://somehost/url"
}
resource "datapower_b2bgateway" "acc_test" {
  id                            = "AccTest_B2BGateway"
  app_domain                    = datapower_domain.acc_test.app_domain
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "archpurge"
  xml_manager                   = "default"
}
resource "datapower_b2bprofile" "acc_test" {
  id            = "AccTest_B2BProfile"
  app_domain    = datapower_domain.acc_test.app_domain
  business_i_ds = ["businessid"]
  destinations = [{
    dest_name = "b2bdestinationname"
    dest_url  = "https://localhost"
  }]
}
resource "datapower_b2bprofilegroup" "acc_test" {
  id         = "AccTest_B2BProfileGroup"
  app_domain = datapower_domain.acc_test.app_domain
  b2b_profiles = [{
    partner_profile = datapower_b2bprofile.acc_test.id
  }]
}
resource "datapower_b2bxpathroutingpolicy" "acc_test" {
  id              = "AccTest_B2BXPathRoutingPolicy"
  app_domain      = datapower_domain.acc_test.app_domain
  sender_x_path   = "senderpath"
  receiver_x_path = "senderpath"
}
resource "datapower_compileoptionspolicy" "acc_test" {
  id         = "AccTest_CompileOptionsPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_compilesettings" "acc_test" {
  id         = "AccTest_CompileSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_configdeploymentpolicy" "acc_test" {
  id         = "AccTest_ConfigDeploymentPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_configsequence" "acc_test" {
  id         = "AccTest_ConfigSequence"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_conformancepolicy" "acc_test" {
  id         = "AccTest_ConformancePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_controllist" "acc_test" {
  id         = "AccTest_ControlList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_cookieattributepolicy" "acc_test" {
  id         = "AccTest_CookieAttributePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_corspolicy" "acc_test" {
  id         = "AccTest_CORSPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  rule       = [datapower_corsrule.acc_test.id]
}
resource "datapower_corsrule" "acc_test" {
  id           = "AccTest_CORSRule"
  app_domain   = datapower_domain.acc_test.app_domain
  allow_origin = ["origin"]
}
resource "datapower_countmonitor" "acc_test" {
  id           = "AccTest_CountMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  measure      = "requests"
  source       = "all"
  max_sources  = 10000
  message_type = datapower_messagetype.acc_test.id
}
resource "datapower_cryptocertificate" "acc_test" {
  id         = "AccTest_CryptoCertificate"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.acc_test_server_crt.remote_path
}
resource "datapower_cryptofwcred" "acc_test" {
  id          = "AccTest_CryptoFWCred"
  app_domain  = datapower_domain.acc_test.app_domain
  private_key = [datapower_cryptokey.acc_test.id]
}
resource "datapower_cryptoidentcred" "acc_test" {
  id          = "AccTest_CryptoIdentCred"
  app_domain  = datapower_domain.acc_test.app_domain
  key         = datapower_cryptokey.acc_test.id
  certificate = datapower_cryptocertificate.acc_test.id
}
resource "datapower_cryptokerberoskdc" "acc_test" {
  id         = "AccTest_CryptoKerberosKDC"
  app_domain = datapower_domain.acc_test.app_domain
  realm      = "realm"
  server     = "localhost"
}
resource "datapower_cryptokerberoskeytab" "acc_test" {
  id         = "AccTest_CryptoKerberosKeytab"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.kerberos-keytab.remote_path
}
resource "datapower_cryptokey" "acc_test" {
  id         = "AccTest_CryptoKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.acc_test_server_key.remote_path
  alias      = datapower_passwordalias.acc_test.id
}
resource "datapower_cryptosskey" "acc_test" {
  id         = "AccTest_CryptoSSKey"
  app_domain = datapower_domain.acc_test.app_domain
  filename   = datapower_file.token-secret.remote_path
}
resource "datapower_cryptovalcred" "acc_test" {
  id         = "AccTest_CryptoValCred"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_deploymentpolicyparametersbinding" "acc_test" {
  id         = "AccTest_DeploymentPolicyParametersBinding"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_documentcryptomap" "acc_test" {
  id         = "AccTest_DocumentCryptoMap"
  app_domain = datapower_domain.acc_test.app_domain
  operation  = "encrypt"
  x_path     = ["*", ]
}
resource "datapower_domain" "acc_test" {
  app_domain = "acceptance_test"
}
resource "datapower_durationmonitor" "acc_test" {
  id           = "AccTest_DurationMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  measure      = "messages"
  message_type = datapower_messagetype.acc_test.id
}
resource "datapower_ebms2sourceprotocolhandler" "acc_test" {
  id            = "AccTest_EBMS2SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_ebms3sourceprotocolhandler" "acc_test" {
  id            = "AccTest_EBMS3SourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 80
}
resource "datapower_filteraction" "acc_test" {
  id         = "AccTest_FilterAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "notify"
}
resource "datapower_formsloginpolicy" "acc_test" {
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
resource "datapower_ftpfilepollersourceprotocolhandler" "acc_test" {
  id                       = "AccTest_FTPFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  target_directory         = "ftp://user:password@host:port/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_ftpquotecommands" "acc_test" {
  id         = "AccTest_FTPQuoteCommands"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ftpserversourceprotocolhandler" "acc_test" {
  id            = "AccTest_FTPServerSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 21
}
resource "datapower_gitopstemplate" "acc_test" {
  id         = "AccTest_GitOpsTemplate"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_graphqlschemaoptions" "acc_test" {
  id         = "AccTest_GraphQLSchemaOptions"
  app_domain = datapower_domain.acc_test.app_domain
  api        = datapower_apidefinition.acc_test.id
}
resource "datapower_hostalias" "acc_test" {
  id         = "AccTest_HostAlias"
  ip_address = "10.10.10.10"
}
resource "datapower_httpinputconversionmap" "acc_test" {
  id         = "AccTest_HTTPInputConversionMap"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_httpsourceprotocolhandler" "acc_test" {
  id            = "AccTest_HTTPSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 8088
}
resource "datapower_httpssourceprotocolhandler" "acc_test" {
  id                     = "AccTest_HTTPSSourceProtocolHandler"
  app_domain             = datapower_domain.acc_test.app_domain
  ssl_server_config_type = "server"
  ssl_server             = datapower_sslserverprofile.acc_test.id
}
resource "datapower_httpuseragent" "acc_test" {
  id         = "AccTest_HTTPUserAgent"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_importpackage" "acc_test" {
  id                = "AccTest_ImportPackage"
  app_domain        = datapower_domain.acc_test.app_domain
  url               = "http://localhost/config.zip"
  import_format     = "ZIP"
  overwrite_files   = true
  overwrite_objects = true
}
resource "datapower_includeconfig" "acc_test" {
  id         = "AccTest_IncludeConfig"
  app_domain = datapower_domain.acc_test.app_domain
  url        = "http://localhost/config.zip"
}
resource "datapower_joserecipientidentifier" "acc_test" {
  id         = "AccTest_JOSERecipientIdentifier"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "key"
  key        = datapower_cryptokey.acc_test.id
  header_param = [{
    header_value = "VALUE"
  }]
}
resource "datapower_josesignatureidentifier" "acc_test" {
  id          = "AccTest_JOSESignatureIdentifier"
  app_domain  = datapower_domain.acc_test.app_domain
  type        = "certificate"
  certificate = datapower_cryptocertificate.acc_test.id
  header_param = [{
    header_value = "VALUE"
  }]
}
resource "datapower_jsonsettings" "acc_test" {
  id         = "AccTest_JSONSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_jweheader" "acc_test" {
  id         = "AccTest_JWEHeader"
  app_domain = datapower_domain.acc_test.app_domain
  recipient  = datapower_jwerecipient.acc_test.id
}
resource "datapower_jwerecipient" "acc_test" {
  id          = "AccTest_JWERecipient"
  app_domain  = datapower_domain.acc_test.app_domain
  algorithm   = "RSA1_5"
  certificate = datapower_cryptocertificate.acc_test.id
}
resource "datapower_jwssignature" "acc_test" {
  id         = "AccTest_JWSSignature"
  app_domain = datapower_domain.acc_test.app_domain
  algorithm  = "RS256"
  key        = datapower_cryptokey.acc_test.id
}
resource "datapower_ldapconnectionpool" "acc_test" {
  id         = "AccTest_LDAPConnectionPool"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ldapsearchparameters" "acc_test" {
  id         = "AccTest_LDAPSearchParameters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_loadbalancergroup" "acc_test" {
  id            = "AccTest_LoadBalancerGroup"
  app_domain    = datapower_domain.acc_test.app_domain
  algorithm     = "round-robin"
  retrieve_info = false
  damp          = 120
}
resource "datapower_loglabel" "acc_test" {
  id         = "AccTest_LogLabel"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_logtarget" "acc_test" {
  id         = "AccTest_LogTarget"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "file"
}
resource "datapower_matching" "acc_test" {
  id         = "AccTest_Matching"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_mcfcustomrule" "acc_test" {
  id                = "AccTest_MCFCustomRule"
  app_domain        = datapower_domain.acc_test.app_domain
  custom_rule_name  = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}
resource "datapower_mcfhttpheader" "acc_test" {
  id         = "AccTest_MCFHttpHeader"
  app_domain = datapower_domain.acc_test.app_domain
  http_name  = "HEADERNAME"
  http_value = "HEADERVALUE"
}
resource "datapower_mcfhttpmethod" "acc_test" {
  id          = "AccTest_MCFHttpMethod"
  app_domain  = datapower_domain.acc_test.app_domain
  http_method = "GET"
}
resource "datapower_mcfhttpurl" "acc_test" {
  id                  = "AccTest_MCFHttpURL"
  app_domain          = datapower_domain.acc_test.app_domain
  http_url_expression = "*"
}
resource "datapower_mcfxpath" "acc_test" {
  id                = "AccTest_MCFXPath"
  app_domain        = datapower_domain.acc_test.app_domain
  x_path_expression = "*"
  x_path_value      = "value"
}
resource "datapower_messagecontentfilters" "acc_test" {
  id         = "AccTest_MessageContentFilters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_messagematching" "acc_test" {
  id         = "AccTest_MessageMatching"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_messagetype" "acc_test" {
  id         = "AccTest_MessageType"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_mpgwerroraction" "acc_test" {
  id         = "AccTest_MPGWErrorAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "static"
  local_url  = "store:///schemas/XMLSchema.dtd"
}
resource "datapower_mpgwerrorhandlingpolicy" "acc_test" {
  id         = "AccTest_MPGWErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  policy_maps = [{
    match  = "__default-accept-service-providers__"
    action = datapower_mpgwerroraction.acc_test.id
  }]
}
resource "datapower_mqmanager" "acc_test" {
  id            = "AccTest_MQManager"
  app_domain    = datapower_domain.acc_test.app_domain
  host_name     = "localhost"
  cache_timeout = 60
  xml_manager   = "default"
}
resource "datapower_mqmanagergroup" "acc_test" {
  id                    = "AccTest_MQManagerGroup"
  app_domain            = datapower_domain.acc_test.app_domain
  primary_queue_manager = datapower_mqmanager.acc_test.id
}
resource "datapower_mqv9plusmftsourceprotocolhandler" "acc_test" {
  id            = "AccTest_MQv9PlusMFTSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mqmanager.acc_test.id
  get_queue     = "queue"
}
resource "datapower_mqv9plussourceprotocolhandler" "acc_test" {
  id            = "AccTest_MQv9PlusSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  queue_manager = datapower_mqmanager.acc_test.id
}
resource "datapower_mtompolicy" "acc_test" {
  id         = "AccTest_MTOMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_multiprotocolgateway" "acc_test" {
  id                       = "AccTest_MultiProtocolGateway"
  app_domain               = datapower_domain.acc_test.app_domain
  type                     = "static-backend"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
resource "datapower_namevalueprofile" "acc_test" {
  id            = "AccTest_NameValueProfile"
  app_domain    = datapower_domain.acc_test.app_domain
  default_fixup = "strip"
}
resource "datapower_nfsfilepollersourceprotocolhandler" "acc_test" {
  id                       = "AccTest_NFSFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  target_directory         = "dpnfs://static-mount-name/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_nfsstaticmount" "acc_test" {
  id         = "AccTest_NFSStaticMount"
  app_domain = datapower_domain.acc_test.app_domain
  remote     = "url://test"
}
resource "datapower_oauthprovidersettings" "acc_test" {
  id                = "AccTest_OAuthProviderSettings"
  app_domain        = datapower_domain.acc_test.app_domain
  provider_type     = "native"
  apic_token_secret = datapower_cryptosskey.acc_test.id
}
resource "datapower_oauthsupportedclient" "acc_test" {
  id                     = "AccTest_OAuthSupportedClient"
  app_domain             = datapower_domain.acc_test.app_domain
  o_auth_role            = { "azsvr" : true }
  az_grant               = { "code" : true }
  generate_client_secret = false
  client_secret          = "secret"
  token_secret           = datapower_cryptosskey.acc_test.id
}
resource "datapower_oauthsupportedclientgroup" "acc_test" {
  id         = "AccTest_OAuthSupportedClientGroup"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_odrconnectorgroup" "acc_test" {
  id          = "AccTest_ODRConnectorGroup"
  xml_manager = "default"
}
resource "datapower_opentelemetry" "acc_test" {
  id         = "AccTest_OpenTelemetry"
  app_domain = datapower_domain.acc_test.app_domain
  exporter   = datapower_opentelemetryexporter.acc_test.id
  sampler    = datapower_opentelemetrysampler.acc_test.id
}
resource "datapower_opentelemetryexporter" "acc_test" {
  id         = "AccTest_OpenTelemetryExporter"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "http"
  host_name  = "localhost"
  port       = 4318
  processor  = "batch"
}
resource "datapower_opentelemetrysampler" "acc_test" {
  id           = "AccTest_OpenTelemetrySampler"
  app_domain   = datapower_domain.acc_test.app_domain
  parent_based = true
  type         = "always-on"
}
resource "datapower_operationratelimit" "acc_test" {
  id         = "AccTest_OperationRateLimit"
  app_domain = datapower_domain.acc_test.app_domain
  operation  = datapower_apioperation.acc_test.id
}
resource "datapower_parsesettings" "acc_test" {
  id         = "AccTest_ParseSettings"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_passwordalias" "acc_test" {
  id         = "AccTest_PasswordAlias"
  app_domain = datapower_domain.acc_test.app_domain
  password   = "password"
}
resource "datapower_peergroup" "acc_test" {
  id         = "AccTest_PeerGroup"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "slm"
  url        = ["http://localhost"]
}
resource "datapower_policyattachments" "acc_test" {
  id                   = "AccTest_PolicyAttachments"
  app_domain           = datapower_domain.acc_test.app_domain
  enforcement_mode     = "enforce"
  policy_references    = false
  sla_enforcement_mode = "allow-if-no-sla"
}
resource "datapower_policyparameters" "acc_test" {
  id         = "AccTest_PolicyParameters"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_poppollersourceprotocolhandler" "acc_test" {
  id                    = "AccTest_POPPollerSourceProtocolHandler"
  app_domain            = datapower_domain.acc_test.app_domain
  mail_server           = "localhost"
  port                  = 8888
  conn_security         = "none"
  account               = "account"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}
resource "datapower_processingmetadata" "acc_test" {
  id         = "AccTest_ProcessingMetadata"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_ratelimitdefinition" "acc_test" {
  id         = "AccTest_RateLimitDefinition"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "rate"
  rate       = 1000
}
resource "datapower_ratelimitdefinitiongroup" "acc_test" {
  id               = "AccTest_RateLimitDefinitionGroup"
  app_domain       = datapower_domain.acc_test.app_domain
  update_on_exceed = "all"
}
resource "datapower_samlattributes" "acc_test" {
  id         = "AccTest_SAMLAttributes"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_schemaexceptionmap" "acc_test" {
  id                  = "AccTest_SchemaExceptionMap"
  app_domain          = datapower_domain.acc_test.app_domain
  original_schema_url = "http://localhost"
  schema_exception_rules = [{
    x_path         = "*"
    exception_type = "AllowEncrypted"
  }]
}
resource "datapower_sftpfilepollersourceprotocolhandler" "acc_test" {
  id                       = "AccTest_SFTPFilePollerSourceProtocolHandler"
  app_domain               = datapower_domain.acc_test.app_domain
  ssh_client_connection    = datapower_sshclientprofile.acc_test.id
  target_directory         = "/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}
resource "datapower_slmaction" "acc_test" {
  id         = "AccTest_SLMAction"
  app_domain = datapower_domain.acc_test.app_domain
  type       = "log-only"
}
resource "datapower_slmcredclass" "acc_test" {
  id         = "AccTest_SLMCredClass"
  app_domain = datapower_domain.acc_test.app_domain
  cred_type  = "aaa-mapped-credential"
}
resource "datapower_slmpolicy" "acc_test" {
  id         = "AccTest_SLMPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_slmrsrcclass" "acc_test" {
  id         = "AccTest_SLMRsrcClass"
  app_domain = datapower_domain.acc_test.app_domain
  rsrc_type  = "aaa-mapped-resource"
}
resource "datapower_slmschedule" "acc_test" {
  id         = "AccTest_SLMSchedule"
  app_domain = datapower_domain.acc_test.app_domain
  start_time = "12:34:00"
  duration   = 1440
}
resource "datapower_smtpserverconnection" "acc_test" {
  id               = "AccTest_SMTPServerConnection"
  app_domain       = datapower_domain.acc_test.app_domain
  mail_server_host = "localhost"
  mail_server_port = 25
}
resource "datapower_soapheaderdisposition" "acc_test" {
  id         = "AccTest_SOAPHeaderDisposition"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_socialloginpolicy" "acc_test" {
  id                      = "AccTest_SocialLoginPolicy"
  app_domain              = datapower_domain.acc_test.app_domain
  client_id               = "client_id"
  client_secret           = "client_secret"
  client_ssl_profile      = datapower_sslclientprofile.acc_test.id
  social_provider         = "google"
  provider_az_endpoint    = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token      = false
}
resource "datapower_sqldatasource" "acc_test" {
  id               = "AccTest_SQLDataSource"
  app_domain       = datapower_domain.acc_test.app_domain
  database         = "Oracle"
  username         = "username"
  password_alias   = datapower_passwordalias.acc_test.id
  data_source_id   = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection   = 10
  connect_timeout  = 15
  query_timeout    = 30
  idle_timeout     = 180
}
resource "datapower_sshclientprofile" "acc_test" {
  id            = "AccTest_SSHClientProfile"
  app_domain    = datapower_domain.acc_test.app_domain
  user_name     = "someuser"
  profile_usage = "sftp"
}
resource "datapower_sshserversourceprotocolhandler" "acc_test" {
  id                = "AccTest_SSHServerSourceProtocolHandler"
  app_domain        = datapower_domain.acc_test.app_domain
  local_address     = "0.0.0.0"
  local_port        = 22
  default_directory = "/"
}
resource "datapower_sslclientprofile" "acc_test" {
  id                   = "AccTest_SSLClientProfile"
  app_domain           = datapower_domain.acc_test.app_domain
  ciphers              = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
  validate_server_cert = false
}
resource "datapower_sslproxyservice" "acc_test" {
  id             = "AccTest_SSLProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  local_port     = 4521
  remote_address = "10.10.10.10"
  remote_port    = 9999
  ssl_server     = datapower_sslserverprofile.acc_test.id
  local_address  = "0.0.0.0"
}
resource "datapower_sslserverprofile" "acc_test" {
  id         = "AccTest_SSLServerProfile"
  app_domain = datapower_domain.acc_test.app_domain
  ciphers    = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
  idcred     = datapower_cryptoidentcred.acc_test.id
}
resource "datapower_sslsnimapping" "acc_test" {
  id         = "AccTest_SSLSNIMapping"
  app_domain = datapower_domain.acc_test.app_domain
  sni_mapping = [{
    host_name_wildmat = "hostname_wildmat"
    ssl_server        = datapower_sslserverprofile.acc_test.id
  }]
}
resource "datapower_sslsniserverprofile" "acc_test" {
  id                 = "AccTest_SSLSNIServerProfile"
  app_domain         = datapower_domain.acc_test.app_domain
  sni_server_mapping = datapower_sslsnimapping.acc_test.id
}
resource "datapower_statelesstcpsourceprotocolhandler" "acc_test" {
  id            = "AccTest_StatelessTCPSourceProtocolHandler"
  app_domain    = datapower_domain.acc_test.app_domain
  local_address = "0.0.0.0"
  local_port    = 4000
}
resource "datapower_stylepolicy" "acc_test" {
  id         = "AccTest_StylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_stylepolicyaction" "acc_test" {
  id            = "AccTest_StylePolicyAction"
  app_domain    = datapower_domain.acc_test.app_domain
  type          = "xform"
  input         = "INPUT"
  transform     = "tfile"
  output        = "OUTPUT"
  named_inputs  = null
  named_outputs = null
}
resource "datapower_stylepolicyrule" "acc_test" {
  id            = "AccTest_StylePolicyRule"
  app_domain    = datapower_domain.acc_test.app_domain
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}
resource "datapower_tcpproxyservice" "acc_test" {
  id             = "AccTest_TCPProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  local_port     = 6158
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}
resource "datapower_urlmap" "acc_test" {
  id         = "AccTest_URLMap"
  app_domain = datapower_domain.acc_test.app_domain
  url_map_rule = [{
    pattern = "https://www.company.com/XML/stylesheets/*"
  }]
}
resource "datapower_urlrefreshpolicy" "acc_test" {
  id         = "AccTest_URLRefreshPolicy"
  app_domain = datapower_domain.acc_test.app_domain
  url_refresh_rule = [{
    url_map = "default-attempt-stream-all"
  }]
}
resource "datapower_urlrewritepolicy" "acc_test" {
  id         = "AccTest_URLRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_user" "acc_test" {
  id                = "AccTest_User"
  password          = "Password$123"
  access_level      = "group-defined"
  group_name        = datapower_usergroup.acc_test.id
  snmp_creds        = null
  hashed_snmp_creds = null
}
resource "datapower_usergroup" "acc_test" {
  id              = "AccTest_UserGroup"
  access_policies = ["*/*/*?Access=r"]
}
resource "datapower_visibilitylist" "acc_test" {
  id         = "AccTest_VisibilityList"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_wccservice" "acc_test" {
  id                = "AccTest_WCCService"
  app_domain        = datapower_domain.acc_test.app_domain
  odc_info_hostname = "example.com"
  odc_info_port     = 1024
  time_interval     = 10
}
resource "datapower_webapperrorhandlingpolicy" "acc_test" {
  id         = "AccTest_WebAppErrorHandlingPolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_webappfw" "acc_test" {
  id                       = "AccTest_WebAppFW"
  app_domain               = datapower_domain.acc_test.app_domain
  front_side               = [{ "LocalAddress" : "0.0.0.0" }]
  remote_address           = "10.10.10.10"
  style_policy             = datapower_appsecuritypolicy.acc_test.id
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
resource "datapower_webapprequest" "acc_test" {
  id         = "AccTest_WebAppRequest"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_webappresponse" "acc_test" {
  id          = "AccTest_WebAppResponse"
  app_domain  = datapower_domain.acc_test.app_domain
  policy_type = "admission"
}
resource "datapower_webappsessionpolicy" "acc_test" {
  id            = "AccTest_WebAppSessionPolicy"
  app_domain    = datapower_domain.acc_test.app_domain
  start_matches = "__default-accept-service-providers__"
}
resource "datapower_webservicemonitor" "acc_test" {
  id           = "AccTest_WebServiceMonitor"
  app_domain   = datapower_domain.acc_test.app_domain
  wsdlurl      = "wsdlurl"
  endpoint_url = "endpointurl"
  frontend_url = "frontendurl"
}
resource "datapower_webtokenservice" "acc_test" {
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
resource "datapower_wsendpointrewritepolicy" "acc_test" {
  id         = "AccTest_WSEndpointRewritePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_wsgateway" "acc_test" {
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
resource "datapower_wsrrsavedsearchsubscription" "acc_test" {
  id                = "AccTest_WSRRSavedSearchSubscription"
  app_domain        = datapower_domain.acc_test.app_domain
  server            = datapower_wsrrserver.acc_test.id
  saved_search_name = "ResTestsaved_search"
}
resource "datapower_wsrrserver" "acc_test" {
  id         = "AccTest_WSRRServer"
  app_domain = datapower_domain.acc_test.app_domain
  soap_url   = "https://host:9443/WSRRCoreSDO/services/WSRRCoreSDOPort"
}
resource "datapower_wsrrsubscription" "acc_test" {
  id          = "AccTest_WSRRSubscription"
  app_domain  = datapower_domain.acc_test.app_domain
  server      = datapower_wsrrserver.acc_test.id
  namespace   = "namespace"
  object_type = "wsdl"
  object_name = "ResTestobject"
}
resource "datapower_wsstylepolicy" "acc_test" {
  id         = "AccTest_WSStylePolicy"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_wsstylepolicyrule" "acc_test" {
  id            = "AccTest_WSStylePolicyRule"
  app_domain    = datapower_domain.acc_test.app_domain
  actions       = [datapower_stylepolicyaction.acc_test.id]
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}
resource "datapower_wxsgrid" "acc_test" {
  id             = "AccTest_WXSGrid"
  app_domain     = datapower_domain.acc_test.app_domain
  collective     = datapower_loadbalancergroup.acc_test.id
  grid           = "gridname"
  user_name      = "username"
  password_alias = datapower_passwordalias.acc_test.id
  timeout        = 1000
}
resource "datapower_xacmlpdp" "acc_test" {
  id             = "AccTest_XACMLPDP"
  app_domain     = datapower_domain.acc_test.app_domain
  general_policy = "http://test/uri"
}
resource "datapower_xmlfirewallservice" "acc_test" {
  id            = "AccTest_XMLFirewallService"
  app_domain    = datapower_domain.acc_test.app_domain
  type          = "dynamic-backend"
  xml_manager   = "default"
  local_port    = 7575
  local_address = "0.0.0.0"
}
resource "datapower_xmlmanager" "acc_test" {
  id         = "AccTest_XMLManager"
  app_domain = datapower_domain.acc_test.app_domain
}
resource "datapower_xpathroutingmap" "acc_test" {
  id         = "AccTest_XPathRoutingMap"
  app_domain = datapower_domain.acc_test.app_domain
  x_path_routing_rules = [{
    x_path = "*"
    host   = "localhost"
    port   = 8888
    ssl    = false
  }]
}
resource "datapower_xslcoprocservice" "acc_test" {
  id            = "AccTest_XSLCoprocService"
  app_domain    = datapower_domain.acc_test.app_domain
  local_port    = 8811
  xml_manager   = "default"
  local_address = "0.0.0.0"
}
resource "datapower_xslproxyservice" "acc_test" {
  id             = "AccTest_XSLProxyService"
  app_domain     = datapower_domain.acc_test.app_domain
  type           = "static-backend"
  xml_manager    = "default"
  local_port     = 8899
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}
resource "datapower_xtcprotocolhandler" "acc_test" {
  id             = "AccTest_XTCProtocolHandler"
  app_domain     = datapower_domain.acc_test.app_domain
  local_address  = "0.0.0.0"
  local_port     = 3000
  remote_address = "10.10.10.10"
  remote_port    = 12000
}
resource "datapower_zosnssclient" "acc_test" {
  id             = "AccTest_ZosNSSClient"
  app_domain     = datapower_domain.acc_test.app_domain
  remote_address = "remote.host"
  remote_port    = 443
  client_id      = "client_id"
  system_name    = "ResTestsystem"
  user_name      = "username"
}