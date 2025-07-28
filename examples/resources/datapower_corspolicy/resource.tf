
resource "datapower_corspolicy" "test" {
  id         = "CORSPolicy_name"
  app_domain = "acc_test_domain"
  rule       = ["TestAccCORSRule"]
}