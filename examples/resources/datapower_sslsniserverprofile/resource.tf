
resource "datapower_sslsniserverprofile" "test" {
  id                 = "SSLSNIServerProfile_name"
  app_domain         = "acc_test_domain"
  sni_server_mapping = "TestAccSSLSNIMappingHostnameMap"
}