
resource "datapower_b2b_gateway" "test" {
  id                            = "ResTestB2BGateway"
  app_domain                    = "acceptance_test"
  b2b_profiles                  = [{ partner_profile = "AccTest_B2BProfile" }]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "purgeonly"
  xml_manager                   = "default"
}