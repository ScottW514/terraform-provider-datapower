
resource "datapower_jwssignature" "test" {
  id         = "JWSSignature_test"
  app_domain = "acc_test_domain"
  algorithm  = "RS256"
  key        = datapower_cryptokey.test.id
}