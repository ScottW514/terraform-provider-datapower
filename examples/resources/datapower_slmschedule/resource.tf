
resource "datapower_slmschedule" "test" {
  id         = "SLMSchedule_name"
  app_domain = "acc_test_domain"
  start_time = "12:34:00"
  duration   = 1440
}