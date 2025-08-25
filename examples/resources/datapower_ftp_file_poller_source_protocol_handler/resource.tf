
resource "datapower_ftp_file_poller_source_protocol_handler" "test" {
  id                       = "ResTestFTPFilePollerSourceProtocolHandler"
  app_domain               = "acceptance_test"
  target_directory         = "ftp://user:password@host:port/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  result_name_pattern      = "../result/$1"
  processing_seize_timeout = 0
  xml_manager              = "default"
}