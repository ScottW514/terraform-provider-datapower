
resource "datapower_apisecurityrequirement" "test" {
  id         = "APISecurityRequirement_name"
  app_domain = "acc_test_domain"
  security   = [datapower_apisecurityapikey.test.id]
}