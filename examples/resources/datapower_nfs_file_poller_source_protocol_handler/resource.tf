
resource "datapower_nfs_file_poller_source_protocol_handler" "test" {
  id                       = "ResTestNFSFilePollerSourceProtocolHandler"
  app_domain               = "acceptance_test"
  target_directory         = "dpnfs://static-mount-name/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}