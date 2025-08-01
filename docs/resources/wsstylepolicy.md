---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_wsstylepolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  WS-Proxy Processing Policy
  CLI Alias: wsm-stylepolicy
---

# datapower_wsstylepolicy (Resource)

WS-Proxy Processing Policy
  - CLI Alias: `wsm-stylepolicy`

## Example Usage

```terraform
resource "datapower_wsstylepolicy" "test" {
  id         = "0WSStylePolicy_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `def_stylesheet_for_soap` (String) Default Stylesheet for SOAP
  - CLI Alias: `filter`
  - Default value: `store:///filter-reject-all.xsl`
- `def_stylesheet_for_xsl` (String) Default Stylesheet for XSL Transforms
  - CLI Alias: `xsldefault`
  - Default value: `store:///identity.xsl`
- `policy_maps` (Attributes List) Policy Maps
  - CLI Alias: `match` (see [below for nested schema](#nestedatt--policy_maps))
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--policy_maps"></a>
### Nested Schema for `policy_maps`

Required:

- `match` (String) Match Rule
  - Reference to: `datapower_matching:id`
- `rule` (String) Stylepolicy Rule
  - Reference to: `datapower_wsstylepolicyrule:id`

Optional:

- `subscription` (String) Subscription
- `wsdl_component_type` (String) WSDL Component Type
  - Choices: `all`, `subscription`, `wsdl`, `service`, `port`, `operation`, `fragmentid`
  - Default value: `all`
- `wsdl_component_value` (String) WSDL Component Value
- `wsdl_fragment_id` (String) Fragment Identifier
