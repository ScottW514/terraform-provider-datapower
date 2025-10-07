
resource "datapower_jwe_header" "test" {
  id         = "ResTestJWEHeader"
  app_domain = "acceptance_test"
  jwe_protected_header = [{
    header_value = "VALUE"
  }]
  jwe_shared_unprotected_header = [{
    header_value = "VALUE"
  }]
  recipient = "AccTest_JWERecipient"
}