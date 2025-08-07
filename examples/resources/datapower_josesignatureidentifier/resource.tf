
resource "datapower_josesignatureidentifier" "test" {
  id         = "ResTestJOSESignatureIdentifier"
  app_domain = "acceptance_test"
  type       = "certificate"
  header_param = [{
    header_value = "VALUE"
  }]
}