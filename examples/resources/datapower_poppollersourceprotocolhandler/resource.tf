
resource "datapower_poppollersourceprotocolhandler" "test" {
  id                    = "ResTestPOPPollerSourceProtocolHandler"
  app_domain            = "acceptance_test"
  mail_server           = "localhost"
  port                  = 8888
  conn_security         = "none"
  account               = "account"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}