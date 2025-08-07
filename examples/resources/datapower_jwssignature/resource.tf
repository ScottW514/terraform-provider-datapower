
resource "datapower_jwssignature" "test" {
  id         = "ResTestJWSSignature"
  app_domain = "acceptance_test"
  algorithm  = "RS256"
  key        = "AccTest_CryptoKey"
}