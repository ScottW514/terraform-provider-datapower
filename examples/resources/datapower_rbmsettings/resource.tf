
resource "datapower_rbmsettings" "test" {
  au_method                   = "local"
  au_cache_allow              = "absolute"
  mc_method                   = "local"
  min_password_length         = 6
  require_mixed_case          = false
  require_digit               = false
  require_non_alpha_numeric   = false
  disallow_username_substring = false
  do_password_aging           = false
  do_password_history         = false
  cli_timeout                 = 0
  max_failed_login            = 0
}