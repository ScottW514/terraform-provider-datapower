
resource "datapower_sftpfilepollersourceprotocolhandler" "test" {
  id                       = "SFTPFilePollerSourceProtocolHandler_name"
  app_domain               = "acc_test_domain"
  ssh_client_connection    = "TestAccSSHClientProfile"
  target_directory         = "/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  processing_seize_timeout = 0
  xml_manager              = "default"
}