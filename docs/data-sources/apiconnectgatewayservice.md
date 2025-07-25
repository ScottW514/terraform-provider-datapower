---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_apiconnectgatewayservice Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  API Connect gateway service (default domain only)
---

# datapower_apiconnectgatewayservice (Data Source)

API Connect gateway service (`default` domain only)

## Example Usage

```terraform
data "datapower_apiconnectgatewayservice" "test" {
  depends_on = [datapower_apiconnectgatewayservice.test]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `api_gateway_address` (String) API gateway address
- `api_gateway_port` (Number) API gateway port
- `enabled` (Boolean) Administrative state
- `gateway_peering` (String) Gateway peering
- `gateway_peering_manager` (String) Gateway-peering manager
- `ip_multicast` (String) IP multicast
- `ip_unicast` (String) IP unicast
- `jwt_validation_mode` (String) JWT validation mode
- `jwturl` (String) JWT URL
- `local_address` (String) Local address
- `local_port` (Number) Local port
- `proxy_policy` (Attributes) API Manager proxy
  - CLI Alias: `proxy` (see [below for nested schema](#nestedatt--proxy_policy))
- `ssl_server` (String) TLS server profile
- `user_defined_policies` (List of String) User-defined policies
- `user_summary` (String) Comments
- `v5_compatibility_mode` (Boolean) V5 compatibility mode
- `v5c_slm_mode` (String) SLM peer mode

<a id="nestedatt--proxy_policy"></a>
### Nested Schema for `proxy_policy`

Read-Only:

- `proxy_policy_enable` (Boolean) Enable API Manager proxy
  - Default value: `false`
- `remote_address` (String) API Manager host
- `remote_port` (Number) API Manager port
