
resource "datapower_jws_signature" "test" {
  id         = "ResTestJWSSignature"
  app_domain = "acceptance_test"
  algorithm  = "RS256"
  key        = "AccTest_CryptoKey"
  protected_header = [{
    header_value = "VALUE"
  }]
  unprotected_header = [{
    header_value = "VALUE"
  }]
}