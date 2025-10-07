
resource "datapower_ssh_server_profile" "test" {
  app_domain = "acceptance_test"
  ciphers    = ["AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM", ]
  kex_alg    = ["CURVE25519-SHA256_AT_LIBSSH.ORG", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", ]
  mac_alg    = ["UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256", "HMAC-SHA2-512", "HMAC-SHA1", ]
  host_key_alg = {
  }
}