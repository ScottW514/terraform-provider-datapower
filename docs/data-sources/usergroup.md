---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_usergroup Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  User group
---

# datapower_usergroup (Data Source)

User group

## Example Usage

```terraform
data "datapower_usergroup" "test" {
  depends_on = [datapower_usergroup.test]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `result` (Attributes List) List of objects (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `access_policies` (List of String) Access policies
- `command_group` (List of String) Command group
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `user_summary` (String) Comments
