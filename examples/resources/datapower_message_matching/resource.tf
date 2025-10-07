
resource "datapower_message_matching" "test" {
  id         = "ResTestMessageMatching"
  app_domain = "acceptance_test"
  http_header = [{
    name  = "HEADER"
    value = "VALUE"
  }]
  http_header_exclude = [{
    name  = "HEADER"
    value = "VALUE"
  }]
}