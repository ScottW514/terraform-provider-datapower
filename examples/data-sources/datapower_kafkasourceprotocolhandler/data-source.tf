
data "datapower_kafkasourceprotocolhandler" "test" {
  depends_on = [datapower_kafkasourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}