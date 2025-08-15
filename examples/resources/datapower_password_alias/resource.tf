
resource "datapower_password_alias" "test" {
  id         = "ResTestPasswordAlias"
  app_domain = "acceptance_test"
  password   = "password"
}