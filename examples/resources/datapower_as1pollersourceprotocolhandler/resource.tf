
resource "datapower_as1pollersourceprotocolhandler" "test" {
  id                    = "AS1PollerSourceProtocolHandler_name"
  app_domain            = "acc_test_domain"
  mail_server           = "smtp.host"
  port                  = 25
  conn_security         = "none"
  account               = "account"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}