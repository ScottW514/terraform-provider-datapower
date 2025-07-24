
resource "datapower_b2bprofilegroup" "test" {
  id         = "B2BProfileGroup_name"
  app_domain = "acc_test_domain"
  b2b_profiles = [{
    partner_profile = datapower_b2bprofile.test.id
  }]
}