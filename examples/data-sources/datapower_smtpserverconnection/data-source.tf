
data "datapower_smtpserverconnection" "test" {
  depends_on = [datapower_smtpserverconnection.test]
  app_domain = "acc_test_domain"
}