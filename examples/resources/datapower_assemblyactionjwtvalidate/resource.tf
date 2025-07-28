
resource "datapower_assemblyactionjwtvalidate" "test" {
  id            = "AssemblyActionJWTValidate_name"
  app_domain    = "acc_test_domain"
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}