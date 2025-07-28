
resource "datapower_joserecipientidentifier" "test" {
  id         = "JOSERecipientIdentifier_name"
  app_domain = "acc_test_domain"
  type       = "key"
  header_param = [{
    header_value = "VALUE"
  }]
}