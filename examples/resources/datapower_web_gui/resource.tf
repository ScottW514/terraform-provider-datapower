
resource "datapower_web_gui" "test" {
  local_port             = 9090
  save_config_overwrites = true
  idle_timeout           = 0
  local_address          = "0.0.0.0"
}