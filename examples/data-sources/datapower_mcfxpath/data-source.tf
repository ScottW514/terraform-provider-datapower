
data "datapower_mcfxpath" "test" {
  depends_on = [datapower_mcfxpath.test]
  app_domain = "acc_test_domain"
}