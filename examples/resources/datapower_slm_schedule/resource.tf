
resource "datapower_slm_schedule" "test" {
  id         = "ResTestSLMSchedule"
  app_domain = "acceptance_test"
  start_time = "12:34:00"
  duration   = 1440
}