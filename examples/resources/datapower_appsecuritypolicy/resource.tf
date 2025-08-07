
resource "datapower_appsecuritypolicy" "test" {
  id         = "ResTestAppSecurityPolicy"
  app_domain = "acceptance_test"
  request_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "AccTest_WebAppRequest"
  }]
  response_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "AccTest_WebAppResponse"
  }]
}