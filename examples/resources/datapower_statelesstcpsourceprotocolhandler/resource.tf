
resource "datapower_statelesstcpsourceprotocolhandler" "test" {
  id            = "StatelessTCPSourceProtocolHandler_name"
  app_domain    = "acc_test_domain"
  local_address = "0.0.0.0"
  local_port    = 4000
}