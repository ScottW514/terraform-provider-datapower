---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_webb2bviewer Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  B2B viewer management service (default domain only)
  CLI Alias: b2b-viewer-mgmt
---

# datapower_webb2bviewer (Resource)

B2B viewer management service (`default` domain only)
  - CLI Alias: `b2b-viewer-mgmt`

## Example Usage

```terraform
resource "datapower_webb2bviewer" "test" {
  local_port    = 9091
  idle_timeout  = 600
  local_address = "0.0.0.0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `acl` (String) Access control list
  - CLI Alias: `acl`
  - Reference to: `datapower_accesscontrollist:id`
  - Default value: `web-b2b-viewer`
- `enabled` (Boolean) Administrative state
  - CLI Alias: `admin-state`
  - Default value: `false`
- `idle_timeout` (Number) Idle timeout
  - CLI Alias: `idle-timeout`
  - Range: `0`-`65535`
  - Default value: `600`
- `local_address` (String) Local address
  - CLI Alias: `ip-address`
  - Default value: `0.0.0.0`
- `local_port` (Number) Port Number
  - CLI Alias: `port`
  - Range: `1`-`65535`
  - Default value: `9091`
- `ssl_server` (String) Custom TLS server profile
  - CLI Alias: `ssl-server`
  - Reference to: `datapower_sslserverprofile:id`
- `ssl_server_config_type` (String) Custom TLS server type
  - CLI Alias: `ssl-config-type`
  - Choices: `server`, `sni`
  - Default value: `server`
- `sslsni_server` (String) Custom TLS SNI server profile
  - CLI Alias: `ssl-sni-server`
  - Reference to: `datapower_sslsniserverprofile:id`
- `user_agent` (String) Custom user agent
  - CLI Alias: `user-agent`
  - Reference to: `datapower_httpuseragent:id`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
