
data "datapower_mpgwerrorhandlingpolicy" "test" {
  depends_on = [datapower_mpgwerrorhandlingpolicy.test]
  app_domain = "acc_test_domain"
}