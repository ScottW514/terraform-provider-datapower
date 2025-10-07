
resource "datapower_processing_metadata" "test" {
  id         = "ResTestProcessingMetadata"
  app_domain = "acceptance_test"
  meta_item = [{
    meta_category = "all"
    meta_name     = "ResTestmeta"
  }]
}