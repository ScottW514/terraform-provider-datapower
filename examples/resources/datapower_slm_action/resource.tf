
resource "datapower_slm_action" "test" {
  id         = "ResTest_SLMAction"
  app_domain = "acceptance_test"
  type       = "log-only"
}