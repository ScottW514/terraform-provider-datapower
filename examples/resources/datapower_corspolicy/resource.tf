
resource "datapower_corspolicy" "test" {
  id         = "CORSPolicy_name"
  app_domain = "acc_test_domain"
  rule       = [datapower_corsrule.test.id]
}