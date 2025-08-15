
resource "datapower_include_config" "test" {
  id         = "ResTestIncludeConfig"
  app_domain = "acceptance_test"
  url        = "http://localhost/config.zip"
}