
resource "datapower_b2bprofile" "test" {
  id         = "B2BProfile_name"
  app_domain = "acc_test_domain"
  destinations = [{
    dest_name = "b2bdestinationname"
    dest_url  = "https://localhost"
  }]
}