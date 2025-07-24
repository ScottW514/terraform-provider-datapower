
data "datapower_mqv9plusmftsourceprotocolhandler" "test" {
  depends_on = [datapower_mqv9plusmftsourceprotocolhandler.test]
  app_domain = "acc_test_domain"
}