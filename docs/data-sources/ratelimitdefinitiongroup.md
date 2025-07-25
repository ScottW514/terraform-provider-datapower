---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_ratelimitdefinitiongroup Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  Rate limit definition group
---

# datapower_ratelimitdefinitiongroup (Data Source)

Rate limit definition group

## Example Usage

```terraform
data "datapower_ratelimitdefinitiongroup" "test" {
  depends_on = [datapower_ratelimitdefinitiongroup.test]
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
- `rate_limit_definitions` (List of String) Rate limit definitions
- `update_on_exceed` (String) Update on exceed
- `user_summary` (String) Comments
