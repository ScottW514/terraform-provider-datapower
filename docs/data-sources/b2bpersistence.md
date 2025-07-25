---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_b2bpersistence Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  B2B persistence (default domain only)
---

# datapower_b2bpersistence (Data Source)

B2B persistence (`default` domain only)

## Example Usage

```terraform
data "datapower_b2bpersistence" "test" {
  depends_on = [datapower_b2bpersistence.test]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `enabled` (Boolean) Administrative state
- `ha_enabled` (Boolean) Enable high availability
- `ha_local_ip` (String) Replication address
- `ha_local_port` (Number) Replication port
- `ha_other_hosts` (Attributes) Alternate host
  - CLI Alias: `ha-other-hosts` (see [below for nested schema](#nestedatt--ha_other_hosts))
- `ha_virtual_ip` (String) Virtual IP address
- `raid_volume` (String) RAID volume
- `storage_size` (Number) Storage size
- `user_summary` (String) Comments

<a id="nestedatt--ha_other_hosts"></a>
### Nested Schema for `ha_other_hosts`

Read-Only:

- `hostname` (String) Remote replication host
- `port` (Number) Remote replication port
  - Range: `1`-`65535`
  - Default value: `1320`
