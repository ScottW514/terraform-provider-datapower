
resource "datapower_ws_style_policy" "test" {
  id         = "ResTestWSStylePolicy"
  app_domain = "acceptance_test"
  policy_maps = [{
    wsdl_component_type = "all"
    match               = "__default-accept-service-providers__"
    rule                = "AccTest_WSStylePolicyRule"
  }]
}