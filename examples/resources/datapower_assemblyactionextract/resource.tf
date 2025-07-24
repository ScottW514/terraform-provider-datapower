
resource "datapower_assemblyactionextract" "test" {
  id         = "_name"
  app_domain = "acc_test_domain"
  extract = [{
    capture = "capture"
  }]
}