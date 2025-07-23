
resource "datapower_apildapregistry" "test" {
  id                     = "APILDAPRegistry_test"
  app_domain             = "acc_test_domain"
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = datapower_ldapsearchparameters.test.id
}