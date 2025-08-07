
resource "datapower_assemblyactionjwtvalidate" "test" {
  id            = "ResTestAssemblyActionJWTValidate"
  app_domain    = "acceptance_test"
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}