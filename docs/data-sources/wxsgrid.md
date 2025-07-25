---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_wxsgrid Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  eXtreme Scale Grid
---

# datapower_wxsgrid (Data Source)

eXtreme Scale Grid

## Example Usage

```terraform
data "datapower_wxsgrid" "test" {
  depends_on = [datapower_wxsgrid.test]
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Read-Only

- `result` (Attributes List) List of objects (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `app_domain` (String) The name of the application domain the object belongs to
- `collective` (String) Collective
- `encrypt` (Boolean) Encrypted Data
- `encrypt_alg` (String) PKCS #7 algorithm for encryption and decryption
- `encrypt_ss_key` (String) Shared Secret Key for Encryption and Decryption
- `grid` (String) Grid Name
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `key_obfuscation` (Boolean) Hash Key Obfuscation
- `key_obfuscation_alg` (String) Hash Algorithm for Key Obfuscation
- `password_alias` (String) Password Alias
- `ssl_client` (String) TLS client profile
- `timeout` (Number) Timeout
- `user_name` (String) User
- `user_summary` (String) Comments
