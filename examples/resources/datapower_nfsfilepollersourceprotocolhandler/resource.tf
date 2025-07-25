
resource "datapower_nfsfilepollersourceprotocolhandler" "test" {
  id                       = "NFSFilePollerSourceProtocolHandler_name"
  app_domain               = "acc_test_domain"
  target_directory         = "dpnfs://static-mount-name/path/"
  delay_between_polls      = 60000
  input_file_match_pattern = ".*"
  generate_result_file     = false
  processing_seize_timeout = 0
  xml_manager              = "default"
}