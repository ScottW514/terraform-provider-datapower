
resource "datapower_api_definition" "test" {
  id            = "ResTestAPIDefinition"
  app_domain    = "acceptance_test"
  base_path     = "/"
  path          = ["AccTest_APIPath"]
  content       = "activity"
  error_content = "payload"
}