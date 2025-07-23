
data "datapower_messagetype" "test" {
  depends_on = [datapower_messagetype.test]
  app_domain = "acc_test_domain"
}