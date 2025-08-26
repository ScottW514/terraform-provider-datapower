
resource "datapower_b2b_profile" "test" {
  id           = "ResTestB2BProfile"
  app_domain   = "acceptance_test"
  business_ids = ["businessid"]
  destinations = [{
    dest_name  = "b2bdestinationname"
    dest_url   = "https://localhost"
    ssl_client = "AccTest_SSLClientProfile"
  }]
}