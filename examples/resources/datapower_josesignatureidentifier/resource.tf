
resource "datapower_josesignatureidentifier" "test" {
  id          = "JOSESignatureIdentifier_test"
  app_domain  = "acc_test_domain"
  type        = "certificate"
  certificate = datapower_cryptocertificate.test.id
  header_param = [{
    header_value = "VALUE"
  }]
  verify = true
}