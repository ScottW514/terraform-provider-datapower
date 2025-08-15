
resource "datapower_tcp_proxy_service" "test" {
  id             = "ResTestTCPProxyService"
  app_domain     = "acceptance_test"
  local_port     = 6158
  remote_address = "10.10.10.10"
  remote_port    = 9999
  local_address  = "0.0.0.0"
}