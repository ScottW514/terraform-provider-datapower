
resource "datapower_namevalueprofile" "test" {
  id            = "ResTestNameValueProfile"
  app_domain    = "acceptance_test"
  default_fixup = "strip"
}