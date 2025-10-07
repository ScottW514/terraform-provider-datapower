
resource "datapower_api_definition" "test" {
  id         = "ResTestAPIDefinition"
  app_domain = "acceptance_test"
  base_path  = "/"
  path       = ["AccTest_APIPath"]
  properties = [{
    property_name = "propertyname"
    catalog       = "*"
  }]
  schemas = [{
    name   = "dtdefname"
    schema = "AccTest_APISchema"
  }]
  allowed_api_protocols = {
  }
}