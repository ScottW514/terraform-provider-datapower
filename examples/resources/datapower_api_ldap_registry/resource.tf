
resource "datapower_api_ldap_registry" "test" {
  id                     = "ResTestAPILDAPRegistry"
  app_domain             = "acceptance_test"
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = "AccTest_LDAPSearchParameters"
}