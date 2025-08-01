---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_cryptokey Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Key
  CLI Alias: key
---

# datapower_cryptokey (Resource)

Key
  - CLI Alias: `key`

## Example Usage

```terraform
resource "datapower_cryptokey" "test" {
  id         = "CryptoKey_name"
  app_domain = "acc_test_domain"
  filename   = "cert:///acc-test-server.key"
  alias      = "TestAccPasswordAlias"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `filename` (String) File
  - CLI Alias: `file-name`
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `alias` (String) Password alias
  - CLI Alias: `password-alias`
  - Reference to: `datapower_passwordalias:id`
