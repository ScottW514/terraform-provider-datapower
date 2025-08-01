---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_b2bcpasendersetting Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  CPA sender setting
  CLI Alias: cpa-sender-setting
---

# datapower_b2bcpasendersetting (Resource)

CPA sender setting
  - CLI Alias: `cpa-sender-setting`

## Example Usage

```terraform
resource "datapower_b2bcpasendersetting" "test" {
  id         = "B2BCPASenderSetting_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `ack_requested` (String) Request acknowledgment
  - CLI Alias: `ack-requested`
  - Choices: `never`, `always`
  - Default value: `never`
- `ack_signature_requested` (String) Request signed acknowledgment
  - CLI Alias: `ack-signature-requested`
  - Choices: `never`, `always`
  - Default value: `never`
- `connection_timeout` (Number) Connection timeout
  - CLI Alias: `timeout`
  - Range: `3`-`7200`
  - Default value: `300`
- `dest_endpoint_url` (String) Destination URL
  - CLI Alias: `dest-url`
- `duplicate_elimination` (String) Duplicate elimination
  - CLI Alias: `duplicate-elimination`
  - Choices: `never`, `always`
  - Default value: `always`
- `enabled_doc_type` (Attributes) Enabled document types
  - CLI Alias: `enabled-doc-type` (see [below for nested schema](#nestedatt--enabled_doc_type))
- `encrypt_algorithm` (String) Encryption algorithm
  - CLI Alias: `encrypt-algorithm`
  - Choices: `http://www.w3.org/2001/04/xmlenc#tripledes-cbc`, `http://www.w3.org/2001/04/xmlenc#aes128-cbc`, `http://www.w3.org/2001/04/xmlenc#aes192-cbc`, `http://www.w3.org/2001/04/xmlenc#aes256-cbc`, `http://www.w3.org/2009/xmlenc11#aes128-gcm`, `http://www.w3.org/2009/xmlenc11#aes192-gcm`, `http://www.w3.org/2009/xmlenc11#aes256-gcm`
  - Default value: `http://www.w3.org/2001/04/xmlenc#tripledes-cbc`
- `encrypt_cert` (String) Encryption certificate
  - CLI Alias: `encrypt-cert`
  - Reference to: `datapower_cryptocertificate:id`
- `encryption_required` (Boolean) Require encryption
  - CLI Alias: `encrypt-required`
  - Default value: `false`
- `include_time_to_live` (Boolean) Include TimeToLive element
  - CLI Alias: `include-time-to-live`
  - Default value: `true`
- `max_retries` (Number) Retransmit attempts
  - CLI Alias: `max-retries`
  - Range: `1`-`30`
  - Default value: `3`
- `password_alias` (String) Password alias
  - CLI Alias: `password-alias`
  - Reference to: `datapower_passwordalias:id`
- `persist_duration` (Number) Persistence duration
  - CLI Alias: `persist-duration`
  - Range: `0`-`6000000`
- `retry` (Boolean) Retransmit unacknowledged messages
  - CLI Alias: `retry`
  - Default value: `false`
- `retry_interval` (Number) Retransmit interval
  - CLI Alias: `retry-interval`
  - Range: `1`-`86400`
  - Default value: `1800`
- `sign_digest_algorithm` (String) Signing digest algorithm
  - CLI Alias: `sign-digest-algorithm`
  - Choices: `sha1`, `sha256`, `sha512`, `ripemd160`, `sha224`, `sha384`, `md5`
  - Default value: `sha1`
- `sign_id_cred` (String) Signature identification credentials
  - CLI Alias: `sign-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `signature_algorithm` (String) Signature algorithm
  - CLI Alias: `sign-algorithm`
  - Choices: `dsa-sha1`, `rsa-sha1`, `rsa-sha256`, `rsa-sha384`, `rsa-sha512`, `rsa-ripemd160`, `rsa-ripemd160-2010`, `sha256-rsa-MGF1`, `rsa-md5`, `ecdsa-sha1`, `ecdsa-sha224`, `ecdsa-sha256`, `ecdsa-sha384`, `ecdsa-sha512`
  - Default value: `rsa-sha1`
- `signature_c14n_algorithm` (String) Signature canonicalization method
  - CLI Alias: `sign-c14n-algorithm`
  - Choices: `c14n`, `exc-c14n`, `c14n-comments`, `exc-c14n-comments`, `c14n11`, `c14n11-comments`
  - Default value: `exc-c14n`
- `signature_required` (Boolean) Require signature
  - CLI Alias: `sign-required`
  - Default value: `false`
- `ssl_client` (String) TLS client profile
  - CLI Alias: `ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ssl_client_config_type` (String) TLS client type
  - CLI Alias: `ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `sync_reply_mode` (String) Sync reply mode
  - CLI Alias: `syncreply-mode`
  - Choices: `mshSignalsOnly`, `none`
  - Default value: `none`
- `user_name` (String) Username
  - CLI Alias: `username`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--enabled_doc_type"></a>
### Nested Schema for `enabled_doc_type`

Optional:

- `enable_binary` (Boolean) Binary
  - Default value: `true`
- `enable_edifact` (Boolean) EDIFACT
  - Default value: `true`
- `enable_x12` (Boolean) X12
  - Default value: `true`
- `enable_xml` (Boolean) XML
  - Default value: `true`
