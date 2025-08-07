
resource "datapower_passwordalias" "test" {
  id         = "ResTestPasswordAlias"
  app_domain = "acceptance_test"
  password   = "password"
}