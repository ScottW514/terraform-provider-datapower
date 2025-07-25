---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_mtompolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  MTOM Policy
  CLI Alias: mtom
---

# datapower_mtompolicy (Resource)

MTOM Policy
  - CLI Alias: `mtom`

## Example Usage

```terraform
resource "datapower_mtompolicy" "test" {
  id         = "MTOMPolicy_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `include_content_type` (Boolean) Include Content Type
  - CLI Alias: `include-content-type`
  - Default value: `true`
- `mode` (String) MTOM Mode
  - CLI Alias: `mode`
  - Choices: `encode`, `decode`
  - Default value: `encode`
- `rule` (Attributes List) MTOM Rules
  - CLI Alias: `rule` (see [below for nested schema](#nestedatt--rule))
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--rule"></a>
### Nested Schema for `rule`

Required:

- `x_path` (String) XPath Expression
  - CLI Alias: `select`

Optional:

- `content_id` (String) Content ID
  - CLI Alias: `content-id`
- `content_type` (String) Content Type
  - CLI Alias: `content-type`
