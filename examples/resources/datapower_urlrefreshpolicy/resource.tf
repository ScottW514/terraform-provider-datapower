
resource "datapower_urlrefreshpolicy" "test" {
  id         = "URLRefreshPolicy_name"
  app_domain = "acc_test_domain"
  url_refresh_rule = [{
    url_map = "default-attempt-stream-all"
  }]
}