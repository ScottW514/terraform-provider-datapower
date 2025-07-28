
resource "datapower_appsecuritypolicy" "test" {
  id         = "AppSecurityPolicy_name"
  app_domain = "acc_test_domain"
  request_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "TestAccRequestProfile"
  }]
  response_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "TestAccResponseProfile"
  }]
}