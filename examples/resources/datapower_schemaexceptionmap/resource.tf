
resource "datapower_schemaexceptionmap" "test" {
  id                  = "ResTestSchemaExceptionMap"
  app_domain          = "acceptance_test"
  original_schema_url = "http://localhost"
  schema_exception_rules = [{
    x_path         = "*"
    exception_type = "AllowEncrypted"
  }]
}