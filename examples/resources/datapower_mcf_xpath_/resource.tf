
resource "datapower_mcf_xpath_" "test" {
  id                = "ResTestMCFXPath"
  app_domain        = "acceptance_test"
  x_path_expression = "*"
  x_path_value      = "value"
}