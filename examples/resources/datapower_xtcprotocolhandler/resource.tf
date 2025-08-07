
resource "datapower_xtcprotocolhandler" "test" {
  id             = "ResTestXTCProtocolHandler"
  app_domain     = "acceptance_test"
  local_address  = "0.0.0.0"
  local_port     = 3000
  remote_address = "10.10.10.10"
  remote_port    = 12000
}