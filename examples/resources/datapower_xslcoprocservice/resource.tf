
resource "datapower_xslcoprocservice" "test" {
  id            = "ResTestXSLCoprocService"
  app_domain    = "acceptance_test"
  local_port    = 8888
  xml_manager   = "default"
  local_address = "0.0.0.0"
}