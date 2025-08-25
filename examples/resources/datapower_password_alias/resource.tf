
resource "datapower_password_alias" "test" {
  id                  = "ResTestPasswordAlias"
  app_domain          = "acceptance_test"
  password_wo         = "password"
  password_wo_version = 1
}