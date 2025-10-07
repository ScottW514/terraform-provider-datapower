
resource "datapower_multi_protocol_gateway" "test" {
  id                              = "ResTestMultiProtocolGateway"
  app_domain                      = "acceptance_test"
  type                            = "static-backend"
  wsrr_saved_search_subscriptions = null
  wsrr_subscriptions              = null
  policy_parameter = [{
    policy_param_parameters          = "AccTest_PolicyParameters"
    policy_param_wsdl_component_type = "all"
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