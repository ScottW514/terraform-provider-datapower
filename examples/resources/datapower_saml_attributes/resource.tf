
resource "datapower_saml_attributes" "test" {
  id         = "ResTestSAMLAttributes"
  app_domain = "acceptance_test"
  saml_attribute = [{
    source_type = "var"
  }]
}