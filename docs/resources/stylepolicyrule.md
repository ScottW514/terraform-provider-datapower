---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_stylepolicyrule Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Processing Rule
  CLI Alias: rule
---

# datapower_stylepolicyrule (Resource)

Processing Rule
  - CLI Alias: `rule`

## Example Usage

```terraform
resource "datapower_stylepolicyrule" "test" {
  id            = "___StylePolicyRule_name"
  app_domain    = "acc_test_domain"
  direction     = "rule"
  input_format  = "none"
  output_format = "none"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `actions` (List of String) Rule Action
  - CLI Alias: `action`
  - Reference to: `datapower_stylepolicyaction:id`
- `direction` (String) Rule Direction
  - CLI Alias: `type`
  - Choices: `rule`, `request-rule`, `response-rule`, `error-rule`
  - Default value: `rule`
- `input_format` (String) Input Filter
  - CLI Alias: `input-filter`
  - Choices: `none`, `gzip`, `pkzip`
  - Default value: `none`
- `non_xml_processing` (Boolean) Non-XML Processing
  - CLI Alias: `non-xml-processing`
  - Default value: `false`
- `output_format` (String) Output Filter
  - CLI Alias: `output-filter`
  - Choices: `none`, `gzip`, `pkzip`
  - Default value: `none`
- `unprocessed` (Boolean) Unprocessed
  - CLI Alias: `unprocessed`
  - Default value: `false`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
