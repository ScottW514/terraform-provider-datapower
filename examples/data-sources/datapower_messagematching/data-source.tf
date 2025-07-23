
data "datapower_messagematching" "test" {
  depends_on = [datapower_messagematching.test]
  app_domain = "acc_test_domain"
}