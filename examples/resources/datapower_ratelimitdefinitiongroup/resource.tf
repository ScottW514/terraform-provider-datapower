
resource "datapower_ratelimitdefinitiongroup" "test" {
  id               = "RateLimitDefinitionGroup_name"
  app_domain       = "acc_test_domain"
  update_on_exceed = "all"
}