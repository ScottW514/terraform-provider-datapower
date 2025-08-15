
resource "datapower_smtp_server_connection" "test" {
  id               = "ResTest_SMTPServerConnection"
  app_domain       = "acceptance_test"
  mail_server_host = "localhost"
  mail_server_port = 25
}