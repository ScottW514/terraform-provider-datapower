---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_apirule Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  API rule
  CLI Alias: api-rule
---

# datapower_apirule (Resource)

API rule
  - CLI Alias: `api-rule`

## Example Usage

```terraform
resource "datapower_apirule" "test" {
  id         = "APIRule_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `actions` (List of String) API actions
  - CLI Alias: `action`
- `dynamic_actions` (Attributes List) API dynamic actions
  - CLI Alias: `dynamic-action` (see [below for nested schema](#nestedatt--dynamic_actions))
- `dynamic_actions_mode` (Boolean) Use dynamic actions
  - CLI Alias: `dynamic-actions-mode`
  - Default value: `false`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--dynamic_actions"></a>
### Nested Schema for `dynamic_actions`

Optional:

- `default` (String) Object reference
  - CLI Alias: `default`
- `literal` (String) Literal configuration
  - CLI Alias: `literal`
- `url` (String) URL reference
  - CLI Alias: `url`
