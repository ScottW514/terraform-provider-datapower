
resource "datapower_import_package" "test" {
  id                = "ResTestImportPackage"
  app_domain        = "acceptance_test"
  url               = "http://localhost/config.zip"
  import_format     = "ZIP"
  overwrite_files   = true
  overwrite_objects = true
}