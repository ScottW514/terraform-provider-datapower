
resource "datapower_assembly_action_jwt_generate" "test" {
  id              = "ResTestAssemblyActionJWTGenerate"
  app_domain      = "acceptance_test"
  issuer_claim    = "iss.claim"
  validity_period = 3600
}