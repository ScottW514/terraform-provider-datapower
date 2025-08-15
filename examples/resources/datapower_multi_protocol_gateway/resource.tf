
resource "datapower_multi_protocol_gateway" "test" {
  id                       = "ResTestMultiProtocolGateway"
  app_domain               = "acceptance_test"
  type                     = "static-backend"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}