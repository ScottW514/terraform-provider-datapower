
data "datapower_messagecontentfilters" "test" {
  depends_on = [datapower_messagecontentfilters.test]
  app_domain = "acc_test_domain"
}