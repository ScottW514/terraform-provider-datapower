
resource "datapower_multiprotocolgateway" "test" {
  id                       = "MultiProtocolGateway_name"
  app_domain               = "acc_test_domain"
  type                     = "static-backend"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}