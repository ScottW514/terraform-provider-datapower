
resource "datapower_documentcryptomap" "test" {
  id         = "ResTestDocumentCryptoMap"
  app_domain = "acceptance_test"
  operation  = "encrypt"
  x_path     = ["*", ]
}