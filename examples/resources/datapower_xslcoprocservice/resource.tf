
resource "datapower_xslcoprocservice" "test" {
  id            = "XSLCoprocService_name"
  app_domain    = "acc_test_domain"
  local_port    = 8888
  xml_manager   = "default"
  local_address = "0.0.0.0"
}