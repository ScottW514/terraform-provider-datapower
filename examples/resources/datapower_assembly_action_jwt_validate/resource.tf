
resource "datapower_assembly_action_jwt_validate" "test" {
  id            = "ResTestAssemblyActionJWTValidate"
  app_domain    = "acceptance_test"
  jwt           = "request.headers.authorization"
  output_claims = "decoded.claims"
}