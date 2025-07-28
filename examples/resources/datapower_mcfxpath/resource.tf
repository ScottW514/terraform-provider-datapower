
resource "datapower_mcfxpath" "test" {
  id                = "MCFXPath_name"
  app_domain        = "acc_test_domain"
  x_path_expression = "*"
  x_path_value      = "value"
}