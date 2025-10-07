
resource "datapower_name_value_profile" "test" {
  id         = "ResTestNameValueProfile"
  app_domain = "acceptance_test"
  validation_list = [{
    name  = "validationname"
    value = "somevalue"
    fixup = "error"
  }]
  default_fixup = "strip"
}