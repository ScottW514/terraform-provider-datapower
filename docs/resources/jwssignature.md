---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_jwssignature Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  JWS Signature
  CLI Alias: jws-signature
---

# datapower_jwssignature (Resource)

JWS Signature
  - CLI Alias: `jws-signature`

## Example Usage

```terraform
resource "datapower_jwssignature" "test" {
  id         = "JWSSignature_name"
  app_domain = "acc_test_domain"
  algorithm  = "RS256"
  key        = "TestAccCryptoKey"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `algorithm` (String) Algorithm
  - CLI Alias: `alg`
  - Choices: `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`, `ES512`, `PS256`, `PS384`, `PS512`
  - Default value: `RS256`
- `key` (String) Private Key
  - CLI Alias: `key`
  - Reference to: `datapower_cryptokey:id`
- `protected_header` (Attributes List) Protected Header
  - CLI Alias: `protected-header` (see [below for nested schema](#nestedatt--protected_header))
- `ss_key` (String) Shared Secret Key
  - CLI Alias: `sskey`
  - Reference to: `datapower_cryptosskey:id`
- `unprotected_header` (Attributes List) Unprotected Header
  - CLI Alias: `unprotected-header` (see [below for nested schema](#nestedatt--unprotected_header))
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--protected_header"></a>
### Nested Schema for `protected_header`

Required:

- `header_value` (String) Value

Optional:

- `header_name` (String) Name


<a id="nestedatt--unprotected_header"></a>
### Nested Schema for `unprotected_header`

Required:

- `header_value` (String) Value

Optional:

- `header_name` (String) Name
