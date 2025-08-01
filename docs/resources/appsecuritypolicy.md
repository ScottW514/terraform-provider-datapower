---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_appsecuritypolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Application Security Policy
  CLI Alias: application-security-policy
---

# datapower_appsecuritypolicy (Resource)

Application Security Policy
  - CLI Alias: `application-security-policy`

## Example Usage

```terraform
resource "datapower_appsecuritypolicy" "test" {
  id         = "AppSecurityPolicy_name"
  app_domain = "acc_test_domain"
  request_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "TestAccRequestProfile"
  }]
  response_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "TestAccResponseProfile"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `request_maps` (Attributes List) Request Maps
  - CLI Alias: `request-match` (see [below for nested schema](#nestedatt--request_maps))
- `response_maps` (Attributes List) Response Maps
  - CLI Alias: `response-match` (see [below for nested schema](#nestedatt--response_maps))

### Optional

- `error_maps` (Attributes List) Error Maps
  - CLI Alias: `error-match` (see [below for nested schema](#nestedatt--error_maps))
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--request_maps"></a>
### Nested Schema for `request_maps`

Required:

- `match` (String) Matching Rule
  - Reference to: `datapower_matching:id`
- `rule` (String) Request Profile
  - Reference to: `datapower_webapprequest:id`


<a id="nestedatt--response_maps"></a>
### Nested Schema for `response_maps`

Required:

- `match` (String) Matching Rule
  - Reference to: `datapower_matching:id`
- `rule` (String) Response Profile
  - Reference to: `datapower_webappresponse:id`


<a id="nestedatt--error_maps"></a>
### Nested Schema for `error_maps`

Required:

- `match` (String) Matching Rule
  - Reference to: `datapower_matching:id`
- `rule` (String) Processing Rule
  - Reference to: `datapower_stylepolicyrule:id`
