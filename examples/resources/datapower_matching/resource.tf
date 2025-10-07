
resource "datapower_matching" "test" {
  id         = "ResTestMatching"
  app_domain = "acceptance_test"
  match_rules = [{
    type = "url"
    url  = "*"
  }]
}