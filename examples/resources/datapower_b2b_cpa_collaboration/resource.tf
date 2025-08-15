
resource "datapower_b2b_cpa_collaboration" "test" {
  id            = "ResTestB2BCPACollaboration"
  app_domain    = "acceptance_test"
  internal_role = "internal"
  external_role = "external"
  service       = "service"
  actions = [{
    name           = "cpacollaborationactionname"
    value          = "value"
    capability     = "cansend"
    sender_setting = "AccTest_B2BCPASenderSetting"
  }]
}