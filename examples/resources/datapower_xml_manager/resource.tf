
resource "datapower_xml_manager" "test" {
  id         = "ResTestXMLManager"
  app_domain = "acceptance_test"
  doc_cache_policy = [{
    type     = "protocol"
    priority = 128
  }]
  schema_validation = [{
    matching        = "__default-accept-service-providers__"
    validation_mode = "default"
  }]
  scheduled_rule = [{
    rule = "__dp-policy-begin__"
  }]
}