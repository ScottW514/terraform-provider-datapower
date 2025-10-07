
resource "datapower_api_operation" "test" {
  id         = "ResTestAPIOperation"
  app_domain = "acceptance_test"
  method     = "GET"
  response_schema = [{
    status_code = "200"
  }]
  parameter = [{
    name     = "apiparameter"
    required = false
    type     = "string"
    where    = "path"
  }]
}