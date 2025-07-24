
resource "datapower_jwerecipient" "test" {
  id          = "JWERecipient_name"
  app_domain  = "acc_test_domain"
  algorithm   = "RSA1_5"
  certificate = datapower_cryptocertificate.test.id
}