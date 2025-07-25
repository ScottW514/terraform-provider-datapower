---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_gitopsvariables Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  GitOps variables (default domain only)
  CLI Alias: gitops-variables
---

# datapower_gitopsvariables (Resource)

GitOps variables (`default` domain only)
  - CLI Alias: `gitops-variables`

## Example Usage

```terraform
resource "datapower_gitopsvariables" "test" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enabled` (Boolean) Administrative state
  - CLI Alias: `admin-state`
  - Default value: `true`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `variables` (Attributes List) Variables
  - CLI Alias: `variable` (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

Required:

- `variable_name` (String) Variable name
  - CLI Alias: `name`
- `variable_value` (String) Variable value
  - CLI Alias: `value`
