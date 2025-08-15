
resource "datapower_name_value_profile" "test" {
  id            = "ResTestNameValueProfile"
  app_domain    = "acceptance_test"
  default_fixup = "strip"
}