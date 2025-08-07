
resource "datapower_assemblyactionjwtgenerate" "test" {
  id              = "ResTestAssemblyActionJWTGenerate"
  app_domain      = "acceptance_test"
  issuer_claim    = "iss.claim"
  validity_period = 3600
}