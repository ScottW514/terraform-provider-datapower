
resource "datapower_socialloginpolicy" "test" {
  id                      = "ResTestSocialLoginPolicy"
  app_domain              = "acceptance_test"
  client_id               = "client_id"
  client_secret           = "client_secret"
  client_ssl_profile      = "AccTest_SSLClientProfile"
  social_provider         = "google"
  provider_az_endpoint    = "https://example.com/auth"
  provider_token_endpoint = "https://example.com/token"
  validate_jwt_token      = true
}