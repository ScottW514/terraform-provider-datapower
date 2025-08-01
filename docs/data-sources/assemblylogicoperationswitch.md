---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_assemblylogicoperationswitch Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  Operation switch assembly action
---

# datapower_assemblylogicoperationswitch (Data Source)

Operation switch assembly action

## Example Usage

```terraform
data "datapower_assemblylogicoperationswitch" "test" {
  depends_on = [datapower_assemblylogicoperationswitch.test]
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

- `action_debug` (Boolean) Enable debugging
- `app_domain` (String) The name of the application domain the object belongs to
- `case` (Attributes List) Case (see [below for nested schema](#nestedatt--result--case))
- `correlation_path` (String) Correlation path
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `otherwise` (String) Otherwise
- `title` (String) Title
- `user_summary` (String) Comments

<a id="nestedatt--result--case"></a>
### Nested Schema for `result.case`

Read-Only:

- `execute` (String) Execute
- `method` (String) HTTP method
  - Choices: `GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `PATCH`, `OPTIONS`, `TRACE`
  - Default value: `GET`
- `operation_id` (String) Operation ID
- `path` (String) API path
