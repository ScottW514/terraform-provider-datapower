
resource "datapower_slmrsrcclass" "test" {
  id         = "ResTestSLMRsrcClass"
  app_domain = "acceptance_test"
  rsrc_type  = "aaa-mapped-resource"
}