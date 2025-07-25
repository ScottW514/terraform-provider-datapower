
resource "datapower_importpackage" "test" {
  id                = "ImportPackage_name"
  app_domain        = "acc_test_domain"
  url               = "http://localhost/config.zip"
  import_format     = "ZIP"
  overwrite_files   = true
  overwrite_objects = true
}