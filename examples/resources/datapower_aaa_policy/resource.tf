
resource "datapower_aaa_policy" "test" {
  id                = "ResTestAAAPolicy"
  app_domain        = "acceptance_test"
  namespace_mapping = null
  extract_identity = {
    ei_bitmap = {
    }
  }
  authenticate = {
    au_method      = "ldap"
    au_host        = "10.1.10.1"
    au_cache_allow = "absolute"
    au_ltpa_token_versions_bitmap = {
    }
    au_sm_cookie_flow = {
    }
    au_sm_header_flow = {
    }
  }
  map_credentials = {
    mc_method = "none"
  }
  extract_resource = {
    er_bitmap = {
    }
  }
  map_resource = {
    mr_method = "none"
  }
  authorize = {
    az_method      = "anyauthenticated"
    az_cache_allow = "absolute"
    az_sm_cookie_flow = {
    }
    az_sm_header_flow = {
    }
  }
  post_process = {
    pp_saml_assertion_type = {
    }
  }
  saml_attribute = null
  ltpa_attributes = [{
    ltpa_user_attribute_name         = "name"
    ltpa_user_attribute_type         = "static"
    ltpa_user_attribute_static_value = "test"
  }]
  transaction_priority = [{
    credential = "cred"
    priority   = "unknown"
  }]
}