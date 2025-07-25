
resource "datapower_xtcprotocolhandler" "test" {
  id             = "XTCProtocolHandler_name"
  app_domain     = "acc_test_domain"
  local_address  = "0.0.0.0"
  local_port     = 3000
  remote_address = "10.10.10.10"
  remote_port    = 12000
}