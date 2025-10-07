
resource "datapower_web_app_fw" "test" {
  id                       = "ResTestWebAppFW"
  app_domain               = "acceptance_test"
  front_side               = [{ "LocalAddress" : "0.0.0.0" }]
  remote_address           = "10.10.10.10"
  style_policy             = "AccTest_AppSecurityPolicy"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
  debug_trigger            = null
}