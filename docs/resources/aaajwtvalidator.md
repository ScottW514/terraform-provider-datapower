---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_aaajwtvalidator Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  JWT Validator
  CLI Alias: jwt-validator
---

# datapower_aaajwtvalidator (Resource)

JWT Validator
  - CLI Alias: `jwt-validator`

## Example Usage

```terraform
resource "datapower_aaajwtvalidator" "test" {
  id             = "AAAJWTValidator_name"
  app_domain     = "acc_test_domain"
  username_claim = "sub"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `aud` (String) Audience
  - CLI Alias: `aud`
- `claims` (Attributes List) Validate claims
  - CLI Alias: `claims` (see [below for nested schema](#nestedatt--claims))
- `customized_script` (String) Custom validation method processing
  - CLI Alias: `customized-script`
- `decrypt_credential_type` (String) Decrypt method
  - CLI Alias: `decrypt-type`
  - Choices: `pkix`, `ssecret`, `jwk`, `jwk-remote`, `custom`
- `decrypt_fetch_cred_ssl_profile` (String) Decrypt credential TLS client profile
  - CLI Alias: `decrypt-fetch-cred-sslprofile`
  - Reference to: `datapower_sslclientprofile:id`
- `decrypt_fetch_cred_url` (String) Decrypt credential URL
  - CLI Alias: `decrypt-fetch-cred-url`
  - Default value: `http://example.com/v3/key`
- `decrypt_jwk` (String) Decrypt JWK
  - CLI Alias: `decrypt-jwk`
- `decrypt_key` (String) Decrypt key
  - CLI Alias: `decrypt-key`
  - Reference to: `datapower_cryptokey:id`
- `decrypt_s_secret` (String) Decrypt shared secret
  - CLI Alias: `decrypt-ssecret`
  - Reference to: `datapower_cryptosskey:id`
- `issuer` (String) Issuer
  - CLI Alias: `iss`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `username_claim` (String) Claim used as username
  - CLI Alias: `username-claim`
  - Default value: `sub`
- `val_method` (Attributes) Validation method
  - CLI Alias: `validate-method` (see [below for nested schema](#nestedatt--val_method))
- `validate_custom` (String) Custom decrypt/verify processing
  - CLI Alias: `validate-custom`
- `verify_certificate` (String) Verify certificate
  - CLI Alias: `verify-certificate`
  - Reference to: `datapower_cryptocertificate:id`
- `verify_certificate_against_val_cred` (Boolean) Signature validation credentials
  - CLI Alias: `verify-certificate-against-valcred`
  - Default value: `false`
- `verify_credential_type` (String) Verify method
  - CLI Alias: `verify-type`
  - Choices: `pkix`, `ssecret`, `jwk`, `jwk-remote`, `custom`
- `verify_fetch_cred_ssl_profile` (String) Verify credential TLS client profile
  - CLI Alias: `verify-fetch-cred-sslprofile`
  - Reference to: `datapower_sslclientprofile:id`
- `verify_fetch_cred_url` (String) Verify credential URL
  - CLI Alias: `verify-fetch-cred-url`
  - Default value: `http://example.com/v3/certs`
- `verify_jwk` (String) Verify JWK
  - CLI Alias: `verify-jwk`
- `verify_s_secret` (String) Verify shared secret
  - CLI Alias: `verify-ssecret`
  - Reference to: `datapower_cryptosskey:id`
- `verify_val_cred` (String) Validation credentials
  - CLI Alias: `valcred`
  - Reference to: `datapower_cryptovalcred:id`

<a id="nestedatt--claims"></a>
### Nested Schema for `claims`

Required:

- `value` (String) Value

Optional:

- `name` (String) Name
- `type` (String) Type
  - Choices: `string`, `bool`, `number`


<a id="nestedatt--val_method"></a>
### Nested Schema for `val_method`

Optional:

- `customized` (Boolean) Custom processing
  - Default value: `false`
- `decrypt` (Boolean) Decrypt
  - Default value: `false`
- `verify` (Boolean) Verify
  - Default value: `false`
