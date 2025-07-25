
resource "datapower_mcfcustomrule" "test" {
  id                = "MCFCustomRule_name"
  app_domain        = "acc_test_domain"
  custom_rule_name  = datapower_stylepolicyrule.test.id
  custom_rule_value = "rulevalue"
}