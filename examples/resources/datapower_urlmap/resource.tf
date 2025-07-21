
resource "datapower_urlmap" "test" {
  id         = "URLMap_name"
  app_domain = "acc_test_domain"
  url_map_rule = [{
    pattern = "https://www.company.com/XML/stylesheets/*"
  }]
}