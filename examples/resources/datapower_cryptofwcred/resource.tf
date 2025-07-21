
resource "datapower_cryptofwcred" "test" {
  id          = "CryptoFWCred_name"
  app_domain  = "acc_test_domain"
  private_key = [datapower_cryptokey.test.id]
}