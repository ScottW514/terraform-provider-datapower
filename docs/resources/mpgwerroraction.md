---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_mpgwerroraction Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Multi-Protocol Gateway Error Action
  CLI Alias: mpgw-error-action
---

# datapower_mpgwerroraction (Resource)

Multi-Protocol Gateway Error Action
  - CLI Alias: `mpgw-error-action`

## Example Usage

```terraform
resource "datapower_mpgwerroraction" "test" {
  id         = "MPGWErrorAction_name"
  app_domain = "acc_test_domain"
  remote_url = "http://google.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `error_rule` (String) Error Rule
  - CLI Alias: `rule`
  - Reference to: `datapower_stylepolicyrule:id`
- `header_injection` (Attributes List) HTTP Header Injection
  - CLI Alias: `header-inject` (see [below for nested schema](#nestedatt--header_injection))
- `local_url` (String) Local page location
  - CLI Alias: `local-url`
- `reason_phrase` (String) Reason Phrase
  - CLI Alias: `reason-phrase`
- `remote_url` (String) Remote URL
  - CLI Alias: `remote-url`
- `status_code` (Number) Response Code
  - CLI Alias: `status-code`
  - Range: `100`-`999`
- `type` (String) Mode
  - CLI Alias: `type`
  - Choices: `error-rule`, `proxy`, `redirect`, `static`
  - Default value: `static`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--header_injection"></a>
### Nested Schema for `header_injection`

Required:

- `header_tag_value` (String) Header Value

Optional:

- `header_tag` (String) Header Name
