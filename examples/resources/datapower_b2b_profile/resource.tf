
resource "datapower_b2b_profile" "test" {
  id           = "ResTestB2BProfile"
  app_domain   = "acceptance_test"
  business_ids = ["businessid"]
  destinations = [{
    dest_name = "b2bdestinationname"
    dest_url  = "https://localhost"
    enabled_doc_type = {
    }
    ssl_client = "AccTest_SSLClientProfile"
  }]
  contacts = null
  ebms_profile_cpa_bindings = [{
    internal_partner = "AccTest_B2BProfile"
    cpa              = "AccTest_B2BCPA"
    collaboration    = "AccTest_B2BCPACollaboration"
  }]
  ebms_message_properties = [{
    name  = "b2bmessagepropertyname"
    value = "value"
  }]
}