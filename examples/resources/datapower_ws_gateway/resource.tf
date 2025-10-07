
resource "datapower_ws_gateway" "test" {
  id           = "ResTestWSGateway"
  app_domain   = "acceptance_test"
  type         = "static-from-wsdl"
  style_policy = "default"
  remote_fetch_retry = {
  }
  wsdl_cache_policy = [{
  }]
  base_wsdl = [{
    wsdl_source_location = "http:///searchservice.wsdl"
    wsdl_name            = "searchsvc"
  }]
  user_toggles = [{
    wsdl_name         = "\"*\""
    service_name      = "\"*\""
    service_port_name = "\"*\""
    port_type_name    = "\"*\""
    binding_name      = "\"*\""
    operation_name    = "\"*\""
    toggles           = null
  }]
  wsrr_subscriptions              = null
  wsrr_saved_search_subscriptions = null
  operation_priority = [{
    scheduler_priority_wsdl_component_type = "all"
  }]
  operation_conformance_policy = [{
    conformance_policy                     = "AccTest_ConformancePolicy"
    conformance_policy_wsdl_component_type = "all"
  }]
  operation_policy_subject_opt_out = [{
    ignored_subjects = {
    }
    policy_subject_opt_out_wsdl_component_type = "all"
  }]
  policy_parameter = [{
    policy_param_parameters          = "AccTest_PolicyParameters"
    policy_param_wsdl_component_type = "all"
  }]
  reliable_messaging = [{
    options = {
    }
    reliable_messaging_wsdl_component_type = "all"
  }]
  xml_manager = "default"
  header_injection = [{
    header_tag_value = "SomeHeaderValue"
  }]
  header_suppression = null
  stylesheet_parameters = [{
    parameter_value = "PARAMETER-VALUE"
  }]
  debug_trigger            = null
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
  in_order_mode = {
  }
}