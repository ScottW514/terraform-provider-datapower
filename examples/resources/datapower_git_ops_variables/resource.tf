
resource "datapower_git_ops_variables" "test" {
  variables = [{
    variable_name  = "variablename"
    variable_value = "variablevalue"
  }]
}