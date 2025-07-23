
resource "datapower_throttler" "test" {
  throttle_at          = 20
  terminate_at         = 5
  temp_fs_throttle_at  = 0
  temp_fs_terminate_at = 0
  qname_warn_at        = 10
  timeout              = 30
}