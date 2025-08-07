
resource "datapower_mpgwerrorhandlingpolicy" "test" {
  id         = "ResTestMPGWErrorHandlingPolicy"
  app_domain = "acceptance_test"
  policy_maps = [{
    match  = "__default-accept-service-providers__"
    action = "AccTest_MPGWErrorAction"
  }]
}