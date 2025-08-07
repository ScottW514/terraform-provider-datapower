
resource "datapower_apigateway" "test" {
  id                       = "ResTestAPIGateway"
  app_domain               = "acceptance_test"
  front_protocol           = ["AccTest_HTTPSourceProtocolHandler"]
  front_timeout            = 120
  front_persistent_timeout = 180
}