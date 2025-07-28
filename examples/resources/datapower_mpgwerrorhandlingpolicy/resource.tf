
resource "datapower_mpgwerrorhandlingpolicy" "test" {
  id         = "MPGWErrorHandlingPolicy_name"
  app_domain = "acc_test_domain"
  policy_maps = [{
    match  = "__default-accept-service-providers__"
    action = "TestAccMPGWErrorAction"
  }]
}