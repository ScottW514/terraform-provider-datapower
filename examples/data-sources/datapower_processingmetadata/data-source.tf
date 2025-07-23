
data "datapower_processingmetadata" "test" {
  depends_on = [datapower_processingmetadata.test]
  app_domain = "acc_test_domain"
}