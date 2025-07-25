
data "datapower_importpackage" "test" {
  depends_on = [datapower_importpackage.test]
  app_domain = "acc_test_domain"
}