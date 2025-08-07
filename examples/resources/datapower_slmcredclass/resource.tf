
resource "datapower_slmcredclass" "test" {
  id         = "ResTestSLMCredClass"
  app_domain = "acceptance_test"
  cred_type  = "aaa-mapped-credential"
}