
resource "datapower_jwssignature" "test" {
  id         = "JWSSignature_name"
  app_domain = "acc_test_domain"
  algorithm  = "RS256"
  key        = "TestAccCryptoKey"
}