---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_domainavailability Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Domain availability (updates restart domain)
  CLI Alias: domain-availability
---

# datapower_domainavailability (Resource)

Domain availability (updates restart domain)
  - CLI Alias: `domain-availability`

## Example Usage

```terraform
resource "datapower_domainavailability" "test" {
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
  - Default value: `false`
- `restart_domain_on_update` (Boolean) Set to true to restart the domain when changes are made to this resource.
- `user_summary` (String) Comments
  - CLI Alias: `summary`
