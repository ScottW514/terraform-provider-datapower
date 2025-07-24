
resource "datapower_webappfw" "test" {
  id                       = "WebAppFW_name"
  app_domain               = "acc_test_domain"
  remote_address           = "10.10.10.10"
  style_policy             = datapower_appsecuritypolicy.test.id
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}