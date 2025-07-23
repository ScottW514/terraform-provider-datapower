
resource "datapower_apisecurityrequirement" "test" {
  id         = "APISecurityRequirement_test"
  app_domain = "acc_test_domain"
  security   = [datapower_apisecurityapikey.test.id]
}