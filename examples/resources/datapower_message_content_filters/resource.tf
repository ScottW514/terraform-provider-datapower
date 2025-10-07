
resource "datapower_message_content_filters" "test" {
  id         = "ResTestMessageContentFilters"
  app_domain = "acceptance_test"
  filters = [{
    filter_name = "ResTestmc_filter"
    type        = "http"
    http_name   = "headername"
    http_value  = "headervalue"
  }]
}