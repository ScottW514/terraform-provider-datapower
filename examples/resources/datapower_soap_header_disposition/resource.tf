
resource "datapower_soap_header_disposition" "test" {
  id         = "ResTestSOAPHeaderDisposition"
  app_domain = "acceptance_test"
  refine = [{
    action = "processed"
  }]
}