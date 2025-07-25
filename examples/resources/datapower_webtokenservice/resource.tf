
resource "datapower_webtokenservice" "test" {
  id                       = "WebTokenService_name"
  app_domain               = "acc_test_domain"
  xml_manager              = "default"
  style_policy             = "default"
  front_timeout            = 120
  front_persistent_timeout = 180
}