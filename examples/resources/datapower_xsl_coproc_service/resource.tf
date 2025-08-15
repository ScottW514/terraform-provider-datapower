
resource "datapower_xsl_coproc_service" "test" {
  id            = "ResTestXSLCoprocService"
  app_domain    = "acceptance_test"
  local_port    = 8811
  xml_manager   = "default"
  local_address = "0.0.0.0"
}