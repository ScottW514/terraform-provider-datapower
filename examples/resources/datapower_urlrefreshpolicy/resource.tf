
resource "datapower_urlrefreshpolicy" "test" {
  id         = "URLRefreshPolicy_test"
  app_domain = "acc_test_domain"
  url_refresh_rule = [{
    url_map = datapower_urlmap.test.id
  }]
}