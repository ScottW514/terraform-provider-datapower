
resource "datapower_apildapregistry" "test" {
  id                     = "APILDAPRegistry_name"
  app_domain             = "acc_test_domain"
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = "TestAccLDAPSearchParameters"
}