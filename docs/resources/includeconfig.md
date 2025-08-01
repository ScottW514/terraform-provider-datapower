---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_includeconfig Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Include configuration file
  CLI Alias: include-config
---

# datapower_includeconfig (Resource)

Include configuration file
  - CLI Alias: `include-config`

## Example Usage

```terraform
resource "datapower_includeconfig" "test" {
  id         = "IncludeConfig_name"
  app_domain = "acc_test_domain"
  url        = "http://localhost/config.zip"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `url` (String) URL
  - CLI Alias: `config-url`

### Optional

- `interface_detection` (Boolean) Interface detection
  - CLI Alias: `interface-detection`
  - Default value: `false`
- `on_startup` (Boolean) Import on startup
  - CLI Alias: `auto-execute`
  - Default value: `true`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
