
data "datapower_kafkacluster" "test" {
  depends_on = [datapower_kafkacluster.test]
  app_domain = "acc_test_domain"
}