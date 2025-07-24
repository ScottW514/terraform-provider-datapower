
resource "datapower_assembly" "test" {
  id         = "_name"
  app_domain = "acc_test_domain"
  rule       = datapower_apirule.test.id
}