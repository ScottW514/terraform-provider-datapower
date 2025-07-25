
resource "datapower_corsrule" "test" {
  id           = "CORSRule_name"
  app_domain   = "acc_test_domain"
  allow_origin = ["origin"]
}