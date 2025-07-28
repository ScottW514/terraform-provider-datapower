
resource "datapower_mcfcustomrule" "test" {
  id                = "MCFCustomRule_name"
  app_domain        = "acc_test_domain"
  custom_rule_name  = "__dp-policy-begin__"
  custom_rule_value = "rulevalue"
}