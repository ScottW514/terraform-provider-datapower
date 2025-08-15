
resource "datapower_jwe_header" "test" {
  id         = "ResTestJWEHeader"
  app_domain = "acceptance_test"
  recipient  = "AccTest_JWERecipient"
}