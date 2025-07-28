
resource "datapower_socialloginpolicy" "test" {
  id                      = "SocialLoginPolicy_name"
  app_domain              = "acc_test_domain"
  client_id               = "client_id"
  client_secret           = "client_secret"
  client_ssl_profile      = "TestAccSSLClientProfile"
  social_provider         = "google"
  provider_az_endpoint    = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token      = true
}