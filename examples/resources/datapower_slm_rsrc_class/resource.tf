
resource "datapower_slm_rsrc_class" "test" {
  id         = "ResTestSLMRsrcClass"
  app_domain = "acceptance_test"
  rsrc_type  = "aaa-mapped-resource"
}