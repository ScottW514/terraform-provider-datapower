
resource "datapower_jweheader" "test" {
  id         = "JWEHeader_test"
  app_domain = "acc_test_domain"
  recipient  = datapower_jwerecipient.test.id
}