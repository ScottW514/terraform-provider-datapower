
resource "datapower_as1_poller_source_protocol_handler" "test" {
  id                    = "ResTestAS1PollerSourceProtocolHandler"
  app_domain            = "acceptance_test"
  mail_server           = "smtp.host"
  port                  = 25
  conn_security         = "none"
  account               = "account"
  password_alias        = "AccTest_PasswordAlias"
  delay_between_polls   = 300
  max_messages_per_poll = 5
}