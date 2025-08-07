
resource "datapower_b2bprofile" "test" {
  id         = "ResTestB2BProfile"
  app_domain = "acceptance_test"
  destinations = [{
    dest_name = "b2bdestinationname"
    dest_url  = "https://localhost"
  }]
}