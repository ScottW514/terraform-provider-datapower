
resource "datapower_ftpfilepollersourceprotocolhandler" "test" {
  id                       = "FTPFilePollerSourceProtocolHandler_name"
  app_domain               = "acc_test_domain"
  target_directory         = "ftp://user:password@host:port/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  processing_seize_timeout = 0
  xml_manager              = "default"
}