
resource "datapower_radius_settings" "test" {
  aaa_servers = [{
    number            = 100
    host              = "10.10.10.10"
    secret_wo         = "opensaysme"
    secret_wo_version = 1
  }]
}