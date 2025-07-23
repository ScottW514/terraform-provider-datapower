
resource "datapower_apidefinition" "test" {
  id            = "APIDefinition_test"
  app_domain    = "acc_test_domain"
  base_path     = "/"
  path          = [datapower_apipath.test.id]
  content       = "activity"
  error_content = "payload"
}