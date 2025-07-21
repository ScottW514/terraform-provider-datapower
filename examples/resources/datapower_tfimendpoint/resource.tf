
resource "datapower_tfimendpoint" "test" {
  id                = "TFIMEndpoint_name"
  app_domain        = "acc_test_domain"
  m_server_url      = "server.url"
  m_server_port     = 9080
  m_compatible_mode = "v6.0"
}