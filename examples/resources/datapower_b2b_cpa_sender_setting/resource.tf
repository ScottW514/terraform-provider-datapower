
resource "datapower_b2b_cpa_sender_setting" "test" {
  id         = "ResTestB2BCPASenderSetting"
  app_domain = "acceptance_test"
  enabled_doc_type = {
  }
  dest_endpoint_url = "ebms2://somehost/url"
}