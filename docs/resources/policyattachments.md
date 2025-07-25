---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_policyattachments Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Policy Attachment
  CLI Alias: policy-attachments
---

# datapower_policyattachments (Resource)

Policy Attachment
  - CLI Alias: `policy-attachments`

## Example Usage

```terraform
resource "datapower_policyattachments" "test" {
  id                   = "PolicyAttachments_name"
  app_domain           = "acc_test_domain"
  enforcement_mode     = "enforce"
  policy_references    = false
  sla_enforcement_mode = "allow-if-no-sla"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `enforcement_mode` (String) Policy Enforcement Mode
  - CLI Alias: `enforcement-mode`
  - Choices: `filter`, `enforce`
  - Default value: `enforce`
- `external_policy` (Attributes List) External Policy
  - CLI Alias: `external-policy` (see [below for nested schema](#nestedatt--external_policy))
- `ignored_policy_attachment_points` (Attributes List) Ignore Embedded Policy
  - CLI Alias: `ignore-attachment-point` (see [below for nested schema](#nestedatt--ignored_policy_attachment_points))
- `policy_references` (Boolean) Policy References
  - CLI Alias: `policy-references`
  - Default value: `false`
- `sla_enforcement_mode` (String) SLA Enforcement Mode
  - CLI Alias: `sla-enforcement-mode`
  - Choices: `allow-if-no-sla`, `reject-if-no-sla`
  - Default value: `allow-if-no-sla`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--external_policy"></a>
### Nested Schema for `external_policy`

Required:

- `external_attach_policy_url` (String) Policy URL

Optional:

- `external_attach_message_content_filter` (String) Service Consumer MCF
  - Reference to: `datapower_messagecontentfilters:id`
- `external_attach_message_content_filter_service_provider` (String) Service Provider MCF
  - Reference to: `datapower_messagecontentfilters:id`
- `external_attach_policy_fragment_id` (String) Fragment Identifier
- `external_attach_wsdl_component_type` (String) Component Type
  - Choices: `service`, `port`, `fragmentid`, `rest`
- `external_attach_wsdl_component_value` (String) WSDL Component Value


<a id="nestedatt--ignored_policy_attachment_points"></a>
### Nested Schema for `ignored_policy_attachment_points`

Optional:

- `policy_attach_fragment_id` (String) Fragment Identifier
- `policy_attach_wsdl_component_type` (String) WSDL Component Type
  - Choices: `service`, `port`, `fragmentid`, `rest`
- `policy_attach_wsdl_component_value` (String) WSDL Component Value
