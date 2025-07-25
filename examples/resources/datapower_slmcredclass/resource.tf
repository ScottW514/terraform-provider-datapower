
resource "datapower_slmcredclass" "test" {
  id         = "SLMCredClass_name"
  app_domain = "acc_test_domain"
  cred_type  = "aaa-mapped-credential"
}