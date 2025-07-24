
data "datapower_b2bcpacollaboration" "test" {
  depends_on = [datapower_b2bcpacollaboration.test]
  app_domain = "acc_test_domain"
}