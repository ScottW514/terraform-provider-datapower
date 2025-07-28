
resource "datapower_assemblyactionjwtgenerate" "test" {
  id              = "AssemblyActionJWTGenerate_name"
  app_domain      = "acc_test_domain"
  issuer_claim    = "iss.claim"
  validity_period = 3600
}