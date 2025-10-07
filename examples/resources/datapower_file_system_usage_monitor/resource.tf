
resource "datapower_file_system_usage_monitor" "test" {
  polling_interval = 60
  all_system       = true
  system = [{
    name               = "system"
    warning_threshold  = 75
    critical_threshold = 90
  }]
}