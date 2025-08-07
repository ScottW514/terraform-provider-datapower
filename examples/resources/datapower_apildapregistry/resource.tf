
resource "datapower_apildapregistry" "test" {
  id                     = "ResTestAPILDAPRegistry"
  app_domain             = "acceptance_test"
  ldap_host              = "localhost"
  ldap_port              = 636
  ldap_search_parameters = "AccTest_LDAPSearchParameters"
}