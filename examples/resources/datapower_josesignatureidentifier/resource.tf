
resource "datapower_josesignatureidentifier" "test" {
  id          = "ResTestJOSESignatureIdentifier"
  app_domain  = "acceptance_test"
  type        = "certificate"
  certificate = "AccTest_CryptoCertificate"
  header_param = [{
    header_value = "VALUE"
  }]
}