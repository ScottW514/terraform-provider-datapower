
resource "datapower_apidefinition" "test" {
  id            = "APIDefinition_name"
  app_domain    = "acc_test_domain"
  base_path     = "/"
  path          = ["TestAccAPIPath"]
  content       = "activity"
  error_content = "payload"
}