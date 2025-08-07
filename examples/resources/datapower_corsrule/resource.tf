
resource "datapower_corsrule" "test" {
  id           = "ResTestCORSRule"
  app_domain   = "acceptance_test"
  allow_origin = ["origin"]
}