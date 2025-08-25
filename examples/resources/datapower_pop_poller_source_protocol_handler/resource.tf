
resource "datapower_pop_poller_source_protocol_handler" "test" {
  id                    = "ResTestPOPPollerSourceProtocolHandler"
  app_domain            = "acceptance_test"
  mail_server           = "localhost"
  port                  = 8888
  conn_security         = "none"
  account               = "account"
  password_alias        = "AccTest_PasswordAlias"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}