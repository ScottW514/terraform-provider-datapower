---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_urlmap Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  URL Map
  CLI Alias: urlmap
---

# datapower_urlmap (Resource)

URL Map
  - CLI Alias: `urlmap`

## Example Usage

```terraform
resource "datapower_urlmap" "test" {
  id         = "___URLMap_name"
  app_domain = "acc_test_domain"
  url_map_rule = [{
    pattern = "https://www.company.com/XML/stylesheets/*"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `url_map_rule` (Attributes List) URL Map Rule
  - CLI Alias: `match` (see [below for nested schema](#nestedatt--url_map_rule))

### Optional

- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--url_map_rule"></a>
### Nested Schema for `url_map_rule`

Required:

- `pattern` (String) Match Pattern
