
data "datapower_assemblyactionratelimit" "test" {
  depends_on = [datapower_assemblyactionratelimit.test]
  app_domain = "acc_test_domain"
}