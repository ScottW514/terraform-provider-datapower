
resource "datapower_b2b_gateway" "test" {
  id         = "ResTestB2BGateway"
  app_domain = "acceptance_test"
  as_front_protocol = [{
    front_protocol = "AccTest_HTTPSourceProtocolHandler"
    mdn_receiver   = false
  }]
  b2b_profiles = [{ partner_profile = "AccTest_B2BProfile" }]
  b2b_groups = [{
  }]
  document_routing_preprocessor = "store:///b2b-routing.xsl"
  archive_mode                  = "purgeonly"
  archive_backup_documents = {
  }
  xml_manager = "default"
  cpa_entries = [{
    cpa              = "AccTest_B2BCPA"
    collaboration    = "AccTest_B2BCPACollaboration"
    internal_partner = "AccTest_B2BProfile"
    external_partner = "AccTest_B2BProfile"
  }]
}