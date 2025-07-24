
resource "datapower_b2bgateway" "test" {
  id                            = "B2BGateway_name"
  app_domain                    = "acc_test_domain"
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "archpurge"
  xml_manager                   = "default"
}