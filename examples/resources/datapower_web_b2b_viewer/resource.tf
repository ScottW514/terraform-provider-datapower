
resource "datapower_web_b2b_viewer" "test" {
  local_port    = 9091
  idle_timeout  = 600
  local_address = "0.0.0.0"
}