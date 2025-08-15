
resource "datapower_web_app_session_policy" "test" {
  id            = "ResTestWebAppSessionPolicy"
  app_domain    = "acceptance_test"
  start_matches = "__default-accept-service-providers__"
}