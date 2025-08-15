
resource "datapower_xacml_pdp" "test" {
  id             = "ResTestXACMLPDP"
  app_domain     = "acceptance_test"
  general_policy = "http://test/uri"
}