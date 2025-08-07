
resource "datapower_mcfxpath" "test" {
  id                = "ResTestMCFXPath"
  app_domain        = "acceptance_test"
  x_path_expression = "*"
  x_path_value      = "value"
}