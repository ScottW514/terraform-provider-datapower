
resource "datapower_namevalueprofile" "test" {
  id            = "NameValueProfile_test"
  app_domain    = "acc_test_domain"
  default_fixup = "strip"
}