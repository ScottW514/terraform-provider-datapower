
resource "datapower_slmaction" "test" {
  id         = "ResTest_SLMAction"
  app_domain = "acceptance_test"
  type       = "log-only"
}