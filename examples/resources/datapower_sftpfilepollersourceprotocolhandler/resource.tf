
resource "datapower_sftpfilepollersourceprotocolhandler" "test" {
  id                       = "ResTestSFTPFilePollerSourceProtocolHandler"
  app_domain               = "acceptance_test"
  ssh_client_connection    = "AccTest_SSHClientProfile"
  target_directory         = "/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}