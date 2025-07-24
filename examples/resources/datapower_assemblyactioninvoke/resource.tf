
resource "datapower_assemblyactioninvoke" "test" {
  id           = "_name"
  app_domain   = "acc_test_domain"
  url          = "https://localhost"
  method       = "Keep"
  backend_type = "detect"
  cache_type   = "Protocol"
}