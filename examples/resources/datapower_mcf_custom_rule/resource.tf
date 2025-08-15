
resource "datapower_mcf_custom_rule" "test" {
  id                = "ResTestMCFCustomRule"
  app_domain        = "acceptance_test"
  custom_rule_name  = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}