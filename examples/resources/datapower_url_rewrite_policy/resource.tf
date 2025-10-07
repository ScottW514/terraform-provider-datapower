
resource "datapower_url_rewrite_policy" "test" {
  id         = "ResTestURLRewritePolicy"
  app_domain = "acceptance_test"
  url_rewrite_rule = [{
    type                 = "header-rewrite"
    match_regexp         = "Test"
    input_replace_regexp = "Replace"
  }]
}