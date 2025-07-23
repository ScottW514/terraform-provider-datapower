
data "datapower_jweheader" "test" {
  depends_on = [datapower_jweheader.test]
  app_domain = "acc_test_domain"
}