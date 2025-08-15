
resource "datapower_cert_monitor" "test" {
  polling_interval      = 1
  reminder_time         = 30
  log_level             = "warn"
  disable_expired_certs = false
}