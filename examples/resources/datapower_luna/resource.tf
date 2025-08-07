
resource "datapower_luna" "test" {
  id              = "ResTestLuna"
  remote_address  = "localhost"
  server_cert     = "cert:///cert.crt"
  security_option = "none"
}