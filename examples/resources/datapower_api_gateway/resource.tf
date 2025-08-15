
resource "datapower_api_gateway" "test" {
  id                       = "ResTestAPIGateway"
  app_domain               = "acceptance_test"
  front_protocol           = ["AccTest_HTTPSourceProtocolHandler"]
  front_timeout            = 120
  front_persistent_timeout = 180
}