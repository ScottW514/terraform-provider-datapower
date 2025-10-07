
resource "datapower_git_ops_template" "test" {
  id         = "ResTestGitOpsTemplate"
  app_domain = "acceptance_test"
  templates = [{
    template_type = "change"
    class_name    = "someclass"
    name          = "templatename"
    field         = "fieldname"
    value         = "templatevalue"
  }]
}