---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_apisecurityhttpscheme Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  API security HTTP scheme
---

# datapower_apisecurityhttpscheme (Data Source)

API security HTTP scheme

## Example Usage

```terraform
data "datapower_apisecurityhttpscheme" "test" {
  depends_on = [datapower_apisecurityhttpscheme.test]
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Read-Only

- `result` (Attributes List) List of objects (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `app_domain` (String) The name of the application domain the object belongs to
- `bearer_format` (String) Bearer format
- `bearer_validation_endpoint` (String) Bearer validation endpoint
- `bearer_validation_method` (String) Bearer validation method
- `bearer_validation_tls_profile` (String) Bearer validation TLS profile
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `scheme` (String) Scheme
- `user_summary` (String) Comments
