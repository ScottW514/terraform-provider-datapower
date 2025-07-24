
resource "datapower_wccservice" "test" {
  id                = "WCCService_name"
  app_domain        = "acc_test_domain"
  odc_info_hostname = "example.com"
  odc_info_port     = 1024
  time_interval     = 10
}