
resource "datapower_document_crypto_map" "test" {
  id                  = "ResTestDocumentCryptoMap"
  app_domain          = "acceptance_test"
  operation           = "encrypt"
  xpath               = ["*", ]
  name_space_mappings = null
}