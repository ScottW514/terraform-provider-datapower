---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_wccservice Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  WebSphere Cell
---

# datapower_wccservice (Data Source)

WebSphere Cell

## Example Usage

```terraform
data "datapower_wccservice" "test" {
  depends_on = [datapower_wccservice.test]
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
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `odc_info_hostname` (String) Deployment Manager Host
- `odc_info_port` (Number) Deployment Manager Port number
- `ssl_client` (String) TLS client profile
- `ssl_client_config_type` (String) TLS client type
- `time_interval` (Number) Time Interval
- `update_type` (String) Update Method
- `user_summary` (String) Comments
