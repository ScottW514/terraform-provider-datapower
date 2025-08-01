---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_messagematching Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Message Matching
  CLI Alias: message-matching
---

# datapower_messagematching (Resource)

Message Matching
  - CLI Alias: `message-matching`

## Example Usage

```terraform
resource "datapower_messagematching" "test" {
  id         = "MessageMatching_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `http_header` (Attributes List) HTTP Headers
  - CLI Alias: `http-header` (see [below for nested schema](#nestedatt--http_header))
- `http_header_exclude` (Attributes List) Excluded HTTP Headers
  - CLI Alias: `http-header-exclude` (see [below for nested schema](#nestedatt--http_header_exclude))
- `http_method` (String) HTTP Method
  - CLI Alias: `method`
  - Choices: `any`, `OPTIONS`, `GET`, `HEAD`, `POST`, `PUT`, `PATCH`, `DELETE`, `TRACE`, `CONNECT`
  - Default value: `any`
- `ip_address` (String) IP Addresses
  - CLI Alias: `ip`
- `ip_exclude` (String) Excluded IP Addresses
  - CLI Alias: `ip-exclude`
- `request_url` (String) Request URL
  - CLI Alias: `request-url`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--http_header"></a>
### Nested Schema for `http_header`

Required:

- `name` (String) Name
- `value` (String) Value Match


<a id="nestedatt--http_header_exclude"></a>
### Nested Schema for `http_header_exclude`

Required:

- `name` (String) Name
- `value` (String) Value Match
