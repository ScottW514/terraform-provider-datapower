---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_jweheader Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  JWE Header
---

# datapower_jweheader (Data Source)

JWE Header

## Example Usage

```terraform
data "datapower_jweheader" "test" {
  depends_on = [datapower_jweheader.test]
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
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `jwe_protected_header` (Attributes List) Protected Header (see [below for nested schema](#nestedatt--result--jwe_protected_header))
- `jwe_shared_unprotected_header` (Attributes List) Shared Unprotected Header (see [below for nested schema](#nestedatt--result--jwe_shared_unprotected_header))
- `recipient` (String) Recipient
- `user_summary` (String) Comments

<a id="nestedatt--result--jwe_protected_header"></a>
### Nested Schema for `result.jwe_protected_header`

Read-Only:

- `header_name` (String) Name
- `header_value` (String) Value


<a id="nestedatt--result--jwe_shared_unprotected_header"></a>
### Nested Schema for `result.jwe_shared_unprotected_header`

Read-Only:

- `header_name` (String) Name
- `header_value` (String) Value
