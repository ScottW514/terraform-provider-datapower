
resource "datapower_webtokenservice" "test" {
  id                       = "ResTestWebTokenService"
  app_domain               = "acceptance_test"
  xml_manager              = "default"
  style_policy             = "default"
  front_timeout            = 120
  front_persistent_timeout = 180
}