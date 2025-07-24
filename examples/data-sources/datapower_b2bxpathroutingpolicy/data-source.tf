
data "datapower_b2bxpathroutingpolicy" "test" {
  depends_on = [datapower_b2bxpathroutingpolicy.test]
  app_domain = "acc_test_domain"
}