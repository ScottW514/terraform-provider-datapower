
resource "datapower_schema_exception_map" "test" {
  id                  = "ResTestSchemaExceptionMap"
  app_domain          = "acceptance_test"
  original_schema_url = "http://localhost"
  schema_exception_rules = [{
    xpath          = "*"
    exception_type = "AllowEncrypted"
  }]
}