
resource "datapower_assemblyactionextract" "test" {
  id         = "AssemblyActionExtract_name"
  app_domain = "acc_test_domain"
  extract = [{
    capture = "capture"
  }]
}