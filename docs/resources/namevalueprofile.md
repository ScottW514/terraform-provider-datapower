---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_namevalueprofile Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Name-Value Profile
  CLI Alias: webapp-gnvc
---

# datapower_namevalueprofile (Resource)

Name-Value Profile
  - CLI Alias: `webapp-gnvc`

## Example Usage

```terraform
resource "datapower_namevalueprofile" "test" {
  id            = "NameValueProfile_name"
  app_domain    = "acc_test_domain"
  default_fixup = "strip"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `default_fixup` (String) No Match Policy
  - CLI Alias: `unvalidated-fixup-policy`
  - Choices: `passthrough`, `strip`, `error`, `set`
  - Default value: `strip`
- `default_map_value` (String) No Match Map Value
  - CLI Alias: `unvalidated-fixup-map`
- `default_xss` (Boolean) No Match XSS Policy
  - CLI Alias: `unvalidated-xss-check`
  - Default value: `false`
- `max_aggregate_size` (Number) Total Size
  - CLI Alias: `max-aggregate-size`
  - Range: `1`-`4294967295`
  - Default value: `128000`
- `max_attributes` (Number) Maximum Count
  - CLI Alias: `max-attributes`
  - Range: `1`-`4294967295`
  - Default value: `256`
- `max_name_size` (Number) Maximum Name Length
  - CLI Alias: `max-name-size`
  - Range: `1`-`4294967295`
  - Default value: `512`
- `max_value_size` (Number) Maximum Value Length
  - CLI Alias: `max-value-size`
  - Range: `1`-`4294967295`
  - Default value: `1024`
- `no_match_xss_patterns_file` (String) XSS (Cross Site Scripting) Protection Patterns File
  - CLI Alias: `unvalidated-xss-patternsfile`
  - Default value: `store:///XSS-Patterns.xml`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `validation_list` (Attributes List) Validation List
  - CLI Alias: `validation` (see [below for nested schema](#nestedatt--validation_list))

<a id="nestedatt--validation_list"></a>
### Nested Schema for `validation_list`

Required:

- `name` (String) Name Expression
- `value` (String) Value Constraint

Optional:

- `fixup` (String) Failure Policy
  - Choices: `passthrough`, `strip`, `error`, `set`
  - Default value: `error`
- `map_value` (String) Map Value
- `xss` (Boolean) Check XSS
  - Default value: `false`
- `xss_patterns_file` (String) XSS (Cross Site Scripting) Protection Patterns File
  - Default value: `store:///XSS-Patterns.xml`
