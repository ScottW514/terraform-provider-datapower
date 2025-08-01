---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_wccservice Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  WebSphere Cell
  CLI Alias: wcc-service
---

# datapower_wccservice (Resource)

WebSphere Cell
  - CLI Alias: `wcc-service`

## Example Usage

```terraform
resource "datapower_wccservice" "test" {
  id                = "WCCService_name"
  app_domain        = "acc_test_domain"
  odc_info_hostname = "example.com"
  odc_info_port     = 1024
  time_interval     = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `odc_info_hostname` (String) Deployment Manager Host
  - CLI Alias: `odc-info-hostname`
- `odc_info_port` (Number) Deployment Manager Port number
  - CLI Alias: `odc-info-port`

### Optional

- `ssl_client` (String) TLS client profile
  - CLI Alias: `ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ssl_client_config_type` (String) TLS client type
  - CLI Alias: `ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `time_interval` (Number) Time Interval
  - CLI Alias: `time-interval`
  - Range: `1`-`86400`
  - Default value: `10`
- `update_type` (String) Update Method
  - CLI Alias: `update-method`
  - Choices: `poll`, `subscribe`
  - Default value: `poll`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
