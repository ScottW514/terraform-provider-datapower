
resource "datapower_url_map" "test" {
  id         = "ResTestURLMap"
  app_domain = "acceptance_test"
  url_map_rule = [{
    pattern = "https://www.company.com/XML/stylesheets/*"
  }]
}