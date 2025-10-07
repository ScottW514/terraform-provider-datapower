
resource "datapower_ssl_client_profile" "test" {
  id         = "ResTestSSLClientProfile"
  app_domain = "acceptance_test"
  protocols = {
    ssl_v3   = false
    tls_v1d0 = false
    tls_v1d1 = true
    tls_v1d2 = true
    tls_v1d3 = true
  }
  ciphers              = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
  validate_server_cert = false
  ssl_client_features = {
  }
  elliptic_curves = ["secp521r1", "secp384r1", "secp256k1", "secp256r1", ]
  hostname_validation_flags = {
  }
  sig_algs = ["ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "ecdsa_sha224", "ecdsa_sha1", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", ]
}