
resource "datapower_joserecipientidentifier" "test" {
  id         = "ResTestJOSERecipientIdentifier"
  app_domain = "acceptance_test"
  type       = "key"
  key        = "AccTest_CryptoKey"
  header_param = [{
    header_value = "VALUE"
  }]
}