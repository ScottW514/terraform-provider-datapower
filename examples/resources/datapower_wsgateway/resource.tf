
resource "datapower_wsgateway" "test" {
  id                       = "ResTestWSGateway"
  app_domain               = "acceptance_test"
  type                     = "static-from-wsdl"
  style_policy             = "default"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}