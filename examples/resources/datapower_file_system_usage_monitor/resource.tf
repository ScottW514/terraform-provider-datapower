
resource "datapower_file_system_usage_monitor" "test" {
  polling_interval   = 60
  all_system         = true
  all_queue_managers = true
}