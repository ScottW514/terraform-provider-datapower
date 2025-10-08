
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
  contacts                  = null
  ebms_profile_cpa_bindings = null
  ebms_message_properties = [{
    name  = "b2bmessagepropertyname"
    value = "value"
  }]
}