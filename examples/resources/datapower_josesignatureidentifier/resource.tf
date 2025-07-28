
resource "datapower_josesignatureidentifier" "test" {
  id         = "JOSESignatureIdentifier_name"
  app_domain = "acc_test_domain"
  type       = "certificate"
  header_param = [{
    header_value = "VALUE"
  }]
}