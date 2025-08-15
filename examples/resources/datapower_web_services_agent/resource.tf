
resource "datapower_web_services_agent" "test" {
  app_domain    = "acceptance_test"
  max_records   = 3000
  max_memory_kb = 64000
  capture_mode  = "faults"
}