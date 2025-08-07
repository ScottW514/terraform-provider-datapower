
resource "datapower_b2bcpacollaboration" "test" {
  id         = "ResTestB2BCPACollaboration"
  app_domain = "acceptance_test"
  service    = "service"
  actions = [{
    name       = "cpacollaborationactionname"
    value      = "value"
    capability = "cansend"
  }]
}