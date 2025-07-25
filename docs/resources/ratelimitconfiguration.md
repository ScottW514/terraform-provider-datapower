---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_ratelimitconfiguration Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Rate limit configuration
  CLI Alias: rate-limit-config
---

# datapower_ratelimitconfiguration (Resource)

Rate limit configuration
  - CLI Alias: `rate-limit-config`

## Example Usage

```terraform
resource "datapower_ratelimitconfiguration" "test" {
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Optional

- `enabled` (Boolean) Administrative state
  - CLI Alias: `admin-state`
  - Default value: `true`
- `parameters` (Attributes List) Parameters
  - CLI Alias: `parameter` (see [below for nested schema](#nestedatt--parameters))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Required:

- `name` (String) Parameter name
  - CLI Alias: `name`
- `value` (String) Parameter value
  - CLI Alias: `value`
