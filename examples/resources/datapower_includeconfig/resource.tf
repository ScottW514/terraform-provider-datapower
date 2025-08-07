
resource "datapower_includeconfig" "test" {
  id         = "ResTestIncludeConfig"
  app_domain = "acceptance_test"
  url        = "http://localhost/config.zip"
}