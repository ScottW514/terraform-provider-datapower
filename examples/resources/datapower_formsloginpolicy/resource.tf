
resource "datapower_formsloginpolicy" "test" {
  id                    = "ResTestFormsLoginPolicy"
  app_domain            = "acceptance_test"
  login_form            = "/LoginPage.htm"
  use_cookie_attributes = false
  enable_migration      = false
  error_page            = "/ErrorPage.htm"
  logout_page           = "/LogoutPage.htm"
  default_url           = "/"
  username_field        = "j_username"
  password_field        = "j_password"
  redirect_field        = "originalUrl"
  form_processing_url   = "/j_security_check"
}