
resource "datapower_includeconfig" "test" {
  id         = "IncludeConfig_name"
  app_domain = "acc_test_domain"
  url        = "http://localhost/config.zip"
}