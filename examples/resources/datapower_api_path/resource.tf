
resource "datapower_api_path" "test" {
  id         = "ResTestAPIPath"
  app_domain = "acceptance_test"
  path       = "/"
  parameter = [{
    name     = "apiparameter"
    required = false
    type     = "string"
    where    = "path"
  }]
}