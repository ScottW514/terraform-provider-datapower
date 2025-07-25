
resource "datapower_documentcryptomap" "test" {
  id         = "DocumentCryptoMap_name"
  app_domain = "acc_test_domain"
  operation  = "encrypt"
  x_path     = ["*", ]
}