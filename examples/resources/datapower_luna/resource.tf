
resource "datapower_luna" "test" {
  id              = "Luna_name"
  remote_address  = "localhost"
  server_cert     = "cert:///cert.crt"
  security_option = "none"
}