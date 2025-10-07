
resource "datapower_domain" "test" {
  app_domain      = "domain_resource_test"
  neighbor_domain = ["default", ]
  file_map = {
    copy_from = true
    copy_to   = true
    delete    = true
    display   = true
    exec      = true
    subdir    = true
  }
  monitoring_map = {
  }
}