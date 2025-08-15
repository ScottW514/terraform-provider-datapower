
resource "datapower_slm_cred_class" "test" {
  id         = "ResTestSLMCredClass"
  app_domain = "acceptance_test"
  cred_type  = "aaa-mapped-credential"
}