
resource "datapower_policy_attachments" "test" {
  id                   = "ResTestPolicyAttachments"
  app_domain           = "acceptance_test"
  enforcement_mode     = "enforce"
  policy_references    = false
  sla_enforcement_mode = "allow-if-no-sla"
}