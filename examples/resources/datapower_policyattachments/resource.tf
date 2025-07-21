
resource "datapower_policyattachments" "test" {
  id                   = "PolicyAttachments_name"
  app_domain           = "acc_test_domain"
  enforcement_mode     = "enforce"
  policy_references    = false
  sla_enforcement_mode = "allow-if-no-sla"
}