---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_cryptofwcred Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Firewall credentials
  CLI Alias: fwcred
---

# datapower_cryptofwcred (Resource)

Firewall credentials
  - CLI Alias: `fwcred`

## Example Usage

```terraform
resource "datapower_cryptofwcred" "test" {
  id          = "CryptoFWCred_name"
  app_domain  = "acc_test_domain"
  private_key = ["TestAccCryptoKey"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `certificate` (List of String) Certificate
  - CLI Alias: `certificate`
  - Reference to: `datapower_cryptocertificate:id`
- `private_key` (List of String) Private key
  - CLI Alias: `key`
  - Reference to: `datapower_cryptokey:id`
- `shared_secret_key` (List of String) Shared secret key
  - CLI Alias: `sskey`
  - Reference to: `datapower_cryptosskey:id`
