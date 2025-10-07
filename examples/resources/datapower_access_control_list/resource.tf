
resource "datapower_access_control_list" "test" {
  id         = "ResTestAccessControlList"
  app_domain = "acceptance_test"
  access_control_entry = [{
    access  = "allow"
    address = "10.0.0.0/8"
  }]
}