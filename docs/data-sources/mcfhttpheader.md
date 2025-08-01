---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_mcfhttpheader Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  HTTP Header Message Content Filter
---

# datapower_mcfhttpheader (Data Source)

HTTP Header Message Content Filter

## Example Usage

```terraform
data "datapower_mcfhttpheader" "test" {
  depends_on = [datapower_mcfhttpheader.test]
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
- `http_name` (String) HTTP Header Name
- `http_value` (String) Header Value
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `user_summary` (String) Comments
