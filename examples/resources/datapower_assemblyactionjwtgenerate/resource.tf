
resource "datapower_assemblyactionjwtgenerate" "test" {
  id              = "_name"
  app_domain      = "acc_test_domain"
  issuer_claim    = "iss.claim"
  validity_period = 3600
}