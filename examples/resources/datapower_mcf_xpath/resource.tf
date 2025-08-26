
resource "datapower_mcf_xpath" "test" {
  id               = "ResTestMCFXPath"
  app_domain       = "acceptance_test"
  xpath_expression = "*"
  xpath_value      = "value"
}