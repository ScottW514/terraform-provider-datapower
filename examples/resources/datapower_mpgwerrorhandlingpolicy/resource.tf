
resource "datapower_mpgwerrorhandlingpolicy" "test" {
  id         = "MPGWErrorHandlingPolicy_name"
  app_domain = "acc_test_domain"
  policy_maps = [{
    match  = datapower_matching.test.id
    action = datapower_mpgwerroraction.test.id
  }]
}