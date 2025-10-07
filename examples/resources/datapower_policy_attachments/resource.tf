
resource "datapower_policy_attachments" "test" {
  id                               = "ResTestPolicyAttachments"
  app_domain                       = "acceptance_test"
  enforcement_mode                 = "enforce"
  policy_references                = false
  ignored_policy_attachment_points = null
  external_policy = [{
    external_attach_wsdl_component_value = "wsdlvalue"
    external_attach_policy_url           = "http://some.url"
  }]
  sla_enforcement_mode = "allow-if-no-sla"
}