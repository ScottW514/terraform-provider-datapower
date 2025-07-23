
resource "datapower_webservicesagent" "test" {
  app_domain    = "acc_test_domain"
  max_records   = 3000
  max_memory_kb = 64000
  capture_mode  = "faults"
}