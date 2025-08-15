
resource "datapower_jwe_recipient" "test" {
  id          = "ResTestJWERecipient"
  app_domain  = "acceptance_test"
  algorithm   = "RSA1_5"
  certificate = "TestAcc_CryptoCertificate"
}