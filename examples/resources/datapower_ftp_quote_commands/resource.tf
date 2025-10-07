
resource "datapower_ftp_quote_commands" "test" {
  id         = "ResTestFTPQuoteCommands"
  app_domain = "acceptance_test"
  ftp_quoted_commands = [{
    quoted_command = "ls"
  }]
}