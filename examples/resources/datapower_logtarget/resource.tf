
resource "datapower_logtarget" "test" {
  id         = "___LogTarget_name"
  app_domain = "acc_test_domain"
  type       = "file"
}