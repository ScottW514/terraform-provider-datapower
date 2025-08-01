---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_messagetype Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  Message Type
---

# datapower_messagetype (Data Source)

Message Type

## Example Usage

```terraform
data "datapower_messagetype" "test" {
  depends_on = [datapower_messagetype.test]
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
- `matching` (List of String) Message Matchings
- `user_summary` (String) Comments
