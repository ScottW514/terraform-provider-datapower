
resource "datapower_urlrefreshpolicy" "test" {
  id         = "ResTestURLRefreshPolicy"
  app_domain = "acceptance_test"
  url_refresh_rule = [{
    url_map = "default-attempt-stream-all"
  }]
}