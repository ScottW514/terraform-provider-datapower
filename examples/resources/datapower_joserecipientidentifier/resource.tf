
resource "datapower_joserecipientidentifier" "test" {
  id         = "JOSERecipientIdentifier_test"
  app_domain = "acc_test_domain"
  type       = "key"
  key        = datapower_cryptokey.test.id
  header_param = [{
    header_value = "VALUE"
  }]
}