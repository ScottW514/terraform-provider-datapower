
resource "datapower_passwordalias" "test" {
  id         = "PasswordAlias_name"
  app_domain = "acc_test_domain"
  password   = "password"
}