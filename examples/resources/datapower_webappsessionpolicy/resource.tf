
resource "datapower_webappsessionpolicy" "test" {
  id            = "ResTestWebAppSessionPolicy"
  app_domain    = "acceptance_test"
  start_matches = "__default-accept-service-providers__"
}