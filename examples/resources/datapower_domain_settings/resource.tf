
resource "datapower_domain_settings" "test" {
  app_domain         = "acceptance_test"
  password_treatment = "masked"
}