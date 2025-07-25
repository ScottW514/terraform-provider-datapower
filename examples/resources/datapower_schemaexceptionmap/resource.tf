
resource "datapower_schemaexceptionmap" "test" {
  id                  = "SchemaExceptionMap_name"
  app_domain          = "acc_test_domain"
  original_schema_url = "http://localhost"
  schema_exception_rules = [{
    x_path         = "*"
    exception_type = "AllowEncrypted"
  }]
}