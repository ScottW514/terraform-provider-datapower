
resource "datapower_wccservice" "test" {
  id                = "ResTestWCCService"
  app_domain        = "acceptance_test"
  odc_info_hostname = "example.com"
  odc_info_port     = 1024
  time_interval     = 10
}