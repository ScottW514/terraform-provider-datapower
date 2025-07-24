
resource "datapower_namevalueprofile" "test" {
  id            = "NameValueProfile_name"
  app_domain    = "acc_test_domain"
  default_fixup = "strip"
}