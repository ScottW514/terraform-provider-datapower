
resource "datapower_logtarget" "test" {
  id         = "ResTest_LogTarget"
  app_domain = "acceptance_test"
  type       = "file"
}