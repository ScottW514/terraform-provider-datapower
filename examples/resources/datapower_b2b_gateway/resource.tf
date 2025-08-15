
resource "datapower_b2b_gateway" "test" {
  id                            = "ResTestB2BGateway"
  app_domain                    = "acceptance_test"
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "archpurge"
  xml_manager                   = "default"
}