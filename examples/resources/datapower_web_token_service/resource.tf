
resource "datapower_web_token_service" "test" {
  id          = "ResTestWebTokenService"
  app_domain  = "acceptance_test"
  xml_manager = "default"
  front_side = [{
    local_address = "0.0.0.0"
    local_port    = 8888
  }]
  style_policy             = "default"
  front_timeout            = 120
  front_persistent_timeout = 180
}