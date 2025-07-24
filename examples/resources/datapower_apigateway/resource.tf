
resource "datapower_apigateway" "test" {
  id                       = "APIGateway_name"
  app_domain               = "acc_test_domain"
  front_protocol           = [datapower_httpsourceprotocolhandler.test.id]
  front_timeout            = 120
  front_persistent_timeout = 180
}