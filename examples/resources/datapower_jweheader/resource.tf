
resource "datapower_jweheader" "test" {
  id         = "ResTestJWEHeader"
  app_domain = "acceptance_test"
  recipient  = "AccTest_JWERecipient"
}