
resource "datapower_assemblyactionjwtvalidate" "test" {
  id            = "_name"
  app_domain    = "acc_test_domain"
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}