
resource "datapower_b2bprofilegroup" "test" {
  id         = "ResTestB2BProfileGroup"
  app_domain = "acceptance_test"
  b2b_profiles = [{
    partner_profile = "AccTest_B2BProfile"
  }]
}