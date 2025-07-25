---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_nfsdynamicmounts Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  NFS dynamic mounts
  CLI Alias: nfs-dynamic-mounts
---

# datapower_nfsdynamicmounts (Resource)

NFS dynamic mounts
  - CLI Alias: `nfs-dynamic-mounts`

## Example Usage

```terraform
resource "datapower_nfsdynamicmounts" "test" {
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Optional

- `enabled` (Boolean) Administrative state
  - CLI Alias: `admin-state`
  - Default value: `false`
- `idle_unmount_seconds` (Number) Inactivity timeout
  - CLI Alias: `inactivity-timeout`
  - Default value: `900`
- `mount_timeout_seconds` (Number) Mount timeout
  - CLI Alias: `mount-timeout`
  - Range: `10`-`240`
  - Default value: `30`
- `mount_type` (String) Mount type
  - CLI Alias: `mount-type`
  - Choices: `hard`, `soft`
  - Default value: `hard`
- `read_only` (Boolean) Read-Only
  - CLI Alias: `read-only`
  - Default value: `false`
- `read_size` (Number) Read size
  - CLI Alias: `rsize`
  - Range: `1024`-`32768`
  - Default value: `4096`
- `retransmissions` (Number) Max retransmissions
  - CLI Alias: `retrans`
  - Range: `1`-`60`
  - Default value: `3`
- `timeout` (Number) Retransmission timeout
  - CLI Alias: `timeo`
  - Range: `1`-`600`
  - Default value: `7`
- `transport` (String) Transport protocol
  - CLI Alias: `transport`
  - Choices: `tcp`, `udp`
  - Default value: `tcp`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `version` (Number) NFS version
  - CLI Alias: `version`
  - Range: `2`-`4`
  - Default value: `3`
- `write_size` (Number) Write size
  - CLI Alias: `wsize`
  - Range: `1024`-`32768`
  - Default value: `4096`
