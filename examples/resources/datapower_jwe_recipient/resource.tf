
resource "datapower_jwe_recipient" "test" {
  id          = "ResTestJWERecipient"
  app_domain  = "acceptance_test"
  algorithm   = "RSA1_5"
  certificate = "AccTest_CryptoCertificate"
  unprotected_header = [{
    header_value = "VALUE"
  }]
}