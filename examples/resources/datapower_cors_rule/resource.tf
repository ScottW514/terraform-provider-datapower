
resource "datapower_cors_rule" "test" {
  id           = "ResTestCORSRule"
  app_domain   = "acceptance_test"
  allow_origin = ["origin"]
}