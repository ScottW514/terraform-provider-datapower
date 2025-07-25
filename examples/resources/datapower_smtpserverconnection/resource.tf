
resource "datapower_smtpserverconnection" "test" {
  id               = "___SMTPServerConnection_name"
  app_domain       = "acc_test_domain"
  mail_server_host = "localhost"
  mail_server_port = 25
}