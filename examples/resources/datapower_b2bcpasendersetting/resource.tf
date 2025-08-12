
resource "datapower_b2bcpasendersetting" "test" {
  id                = "ResTestB2BCPASenderSetting"
  app_domain        = "acceptance_test"
  dest_endpoint_url = "ebms2://somehost/url"
}