
resource "datapower_sslclientprofile" "test" {
  id         = "SSLClientProfile_name"
  app_domain = "acc_test_domain"
  ciphers    = ["RSA_WITH_AES_128_GCM_SHA256", ]
}