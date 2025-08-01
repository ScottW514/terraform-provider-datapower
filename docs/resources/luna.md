---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_luna Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  SafeNet Luna HSM (default domain only)
  CLI Alias: luna
---

# datapower_luna (Resource)

SafeNet Luna HSM (`default` domain only)
  - CLI Alias: `luna`

## Example Usage

```terraform
resource "datapower_luna" "test" {
  id              = "Luna_name"
  remote_address  = "localhost"
  server_cert     = "cert:///cert.crt"
  security_option = "none"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of the object. Must be unique among object types in application domain.
- `remote_address` (String) Host
  - CLI Alias: `host`
- `server_cert` (String) Encryption certificate
  - CLI Alias: `server-cert`

### Optional

- `security_option` (String) Security option
  - CLI Alias: `option`
  - Choices: `none`, `htl`
  - Default value: `none`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
