---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_sslclientprofile Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  TLS client profile
  CLI Alias: ssl-client
---

# datapower_sslclientprofile (Resource)

TLS client profile
  - CLI Alias: `ssl-client`

## Example Usage

```terraform
resource "datapower_sslclientprofile" "test" {
  id         = "SSLClientProfile_name"
  app_domain = "acc_test_domain"
  ciphers    = ["AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `cache_size` (Number) Session cache size
  - CLI Alias: `cache-size`
  - Range: `1`-`500000`
  - Default value: `100`
- `cache_timeout` (Number) Session cache timeout
  - CLI Alias: `cache-timeout`
  - Range: `1`-`86400`
  - Default value: `300`
- `caching` (Boolean) Enable session caching
  - CLI Alias: `caching`
  - Default value: `true`
- `ciphers` (List of String) Ciphers
  - CLI Alias: `ciphers`
  - Choices: `RSA_WITH_NULL_MD5`, `RSA_WITH_NULL_SHA`, `RSA_WITH_RC4_128_MD5`, `RSA_WITH_RC4_128_SHA`, `RSA_WITH_DES_CBC_SHA`, `RSA_WITH_3DES_EDE_CBC_SHA`, `DHE_DSS_WITH_DES_CBC_SHA`, `DHE_DSS_WITH_3DES_EDE_CBC_SHA`, `DHE_RSA_WITH_DES_CBC_SHA`, `DHE_RSA_WITH_3DES_EDE_CBC_SHA`, `RSA_WITH_AES_128_CBC_SHA`, `DHE_DSS_WITH_AES_128_CBC_SHA`, `DHE_RSA_WITH_AES_128_CBC_SHA`, `RSA_WITH_AES_256_CBC_SHA`, `DHE_DSS_WITH_AES_256_CBC_SHA`, `DHE_RSA_WITH_AES_256_CBC_SHA`, `RSA_WITH_NULL_SHA256`, `RSA_WITH_AES_128_CBC_SHA256`, `RSA_WITH_AES_256_CBC_SHA256`, `DHE_DSS_WITH_AES_128_CBC_SHA256`, `DHE_RSA_WITH_AES_128_CBC_SHA256`, `DHE_DSS_WITH_AES_256_CBC_SHA256`, `DHE_RSA_WITH_AES_256_CBC_SHA256`, `RSA_WITH_AES_128_GCM_SHA256`, `RSA_WITH_AES_256_GCM_SHA384`, `DHE_RSA_WITH_AES_128_GCM_SHA256`, `DHE_RSA_WITH_AES_256_GCM_SHA384`, `DHE_DSS_WITH_AES_128_GCM_SHA256`, `DHE_DSS_WITH_AES_256_GCM_SHA384`, `AES_128_GCM_SHA256`, `AES_256_GCM_SHA384`, `CHACHA20_POLY1305_SHA256`, `AES_128_CCM_SHA256`, `AES_128_CCM_8_SHA256`, `ECDHE_RSA_WITH_NULL_SHA`, `ECDHE_RSA_WITH_RC4_128_SHA`, `ECDHE_RSA_WITH_3DES_EDE_CBC_SHA`, `ECDHE_RSA_WITH_AES_128_CBC_SHA`, `ECDHE_RSA_WITH_AES_256_CBC_SHA`, `ECDHE_RSA_WITH_AES_128_CBC_SHA256`, `ECDHE_RSA_WITH_AES_256_CBC_SHA384`, `ECDHE_RSA_WITH_AES_128_GCM_SHA256`, `ECDHE_RSA_WITH_AES_256_GCM_SHA384`, `ECDHE_ECDSA_WITH_NULL_SHA`, `ECDHE_ECDSA_WITH_RC4_128_SHA`, `ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA`, `ECDHE_ECDSA_WITH_AES_128_CBC_SHA`, `ECDHE_ECDSA_WITH_AES_256_CBC_SHA`, `ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`, `ECDHE_ECDSA_WITH_AES_256_CBC_SHA384`, `ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`, `ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`
- `custom_sni_hostname` (String) Custom SNI hostname
  - CLI Alias: `custom-sni-hostname`
- `disable_renegotiation` (Boolean) Disable renegotiation
  - CLI Alias: `disable-renegotiation`
  - Default value: `true`
- `elliptic_curves` (List of String) Elliptic curves
  - CLI Alias: `curves`
  - Choices: `sect163k1`, `sect163r1`, `sect163r2`, `sect193r1`, `sect193r2`, `sect233k1`, `sect233r1`, `sect239k1`, `sect283k1`, `sect283r1`, `sect409k1`, `sect409r1`, `sect571k1`, `sect571r1`, `secp160k1`, `secp160r1`, `secp160r2`, `secp192k1`, `secp192r1`, `secp224k1`, `secp224r1`, `secp256k1`, `secp256r1`, `secp384r1`, `secp521r1`, `brainpoolP256r1`, `brainpoolP384r1`, `brainpoolP512r1`
- `enable_tls13_compat` (Boolean) Enable TLSv1.3 compatibility
  - CLI Alias: `enable-tls13-compat`
  - Default value: `true`
- `hostname_validation_fail_on_error` (Boolean) Hostname validation fail on error
  - CLI Alias: `hostname-validation-fail`
  - Default value: `false`
- `hostname_validation_flags` (Attributes) Hostname validation flags
  - CLI Alias: `hostname-validation-flags` (see [below for nested schema](#nestedatt--hostname_validation_flags))
- `idcred` (String) Identification credentials
  - CLI Alias: `idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `protocols` (Attributes) Protocols
  - CLI Alias: `protocols` (see [below for nested schema](#nestedatt--protocols))
- `require_closure_notification` (Boolean) Require closure notification
  - CLI Alias: `require-closure-notification`
  - Default value: `true`
- `sig_algs` (List of String) Signature algorithms
  - CLI Alias: `sign-alg`
  - Choices: `ecdsa_secp256r1_sha256`, `ecdsa_secp384r1_sha384`, `ecdsa_secp521r1_sha512`, `ed25519`, `ed448`, `ecdsa_sha224`, `ecdsa_sha1`, `rsa_pss_rsae_sha256`, `rsa_pss_rsae_sha384`, `rsa_pss_rsae_sha512`, `rsa_pss_pss_sha256`, `rsa_pss_pss_sha384`, `rsa_pss_pss_sha512`, `rsa_pkcs1_sha256`, `rsa_pkcs1_sha384`, `rsa_pkcs1_sha512`, `rsa_pkcs1_sha224`, `rsa_pkcs1_sha1`, `dsa_sha256`, `dsa_sha384`, `dsa_sha512`, `dsa_sha224`, `dsa_sha1`
- `ssl_client_features` (Attributes) Features
  - CLI Alias: `ssl-client-features` (see [below for nested schema](#nestedatt--ssl_client_features))
- `use_custom_sni_hostname` (Boolean) Use custom SNI host name
  - CLI Alias: `use-custom-sni-hostname`
  - Default value: `false`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `valcred` (String) Validation credentials
  - CLI Alias: `valcred`
  - Reference to: `datapower_cryptovalcred:id`
- `validate_hostname` (Boolean) Validate server hostname
  - CLI Alias: `validate-hostname`
  - Default value: `false`
- `validate_server_cert` (Boolean) Validate server certificate
  - CLI Alias: `validate-server-cert`
  - Default value: `true`

<a id="nestedatt--hostname_validation_flags"></a>
### Nested Schema for `hostname_validation_flags`

Optional:

- `x509_check_flag_always_check_subject` (Boolean) X509_CHECK_FLAG_ALWAYS_CHECK_SUBJECT
  - Default value: `false`
- `x509_check_flag_multi_label_wildcards` (Boolean) X509_CHECK_FLAG_MULTI_LABEL_WILDCARDS
  - Default value: `false`
- `x509_check_flag_no_partial_wildcards` (Boolean) X509_CHECK_FLAG_NO_PARTIAL_WILDCARDS
  - Default value: `false`
- `x509_check_flag_no_wildcards` (Boolean) X509_CHECK_FLAG_NO_WILDCARDS
  - Default value: `false`
- `x509_check_flag_single_label_subdomains` (Boolean) X509_CHECK_FLAG_SINGLE_LABEL_SUBDOMAINS
  - Default value: `false`


<a id="nestedatt--protocols"></a>
### Nested Schema for `protocols`

Optional:

- `ssl_v3` (Boolean) Enable SSL version 3
  - Default value: `false`
- `tls_v1d0` (Boolean) Enable TLS version 1.0
  - Default value: `false`
- `tls_v1d1` (Boolean) Enable TLS version 1.1
  - Default value: `true`
- `tls_v1d2` (Boolean) Enable TLS version 1.2
  - Default value: `true`
- `tls_v1d3` (Boolean) Enable TLS version 1.3
  - Default value: `true`


<a id="nestedatt--ssl_client_features"></a>
### Nested Schema for `ssl_client_features`

Optional:

- `compression` (Boolean) Enable compression
  - Default value: `false`
- `permit_insecure_servers` (Boolean) Permit connections without renegotiation
  - Default value: `false`
- `use_sni` (Boolean) Use SNI
  - Default value: `false`
