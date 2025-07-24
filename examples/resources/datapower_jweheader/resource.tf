
resource "datapower_jweheader" "test" {
  id         = "JWEHeader_name"
  app_domain = "acc_test_domain"
  recipient  = datapower_jwerecipient.test.id
}