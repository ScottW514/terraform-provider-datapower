---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_sshserverprofile Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  SSH server profile
  CLI Alias: sshserverprofile
---

# datapower_sshserverprofile (Resource)

SSH server profile
  - CLI Alias: `sshserverprofile`

## Example Usage

```terraform
resource "datapower_sshserverprofile" "test" {
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Optional

- `ciphers` (List of String) Ciphers
  - CLI Alias: `ciphers`
  - Choices: `CHACHA20-POLY1305_AT_OPENSSH.COM`, `AES128-CTR`, `AES192-CTR`, `AES256-CTR`, `AES128-GCM_AT_OPENSSH.COM`, `AES256-GCM_AT_OPENSSH.COM`
- `enabled` (Boolean) Administrative state
  - CLI Alias: `admin-state`
  - Default value: `true`
- `host_key_alg` (Attributes) Host key algorithms
  - CLI Alias: `host-key-alg` (see [below for nested schema](#nestedatt--host_key_alg))
- `kex_alg` (List of String) Key exchange algorithms
  - CLI Alias: `kex-alg`
  - Choices: `DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256`, `ECDH-SHA2-NISTP256`, `ECDH-SHA2-NISTP384`, `ECDH-SHA2-NISTP521`, `CURVE25519-SHA256_AT_LIBSSH.ORG`
- `mac_alg` (List of String) Message authentication codes
  - CLI Alias: `mac-alg`
  - Choices: `HMAC-SHA1`, `HMAC-SHA2-256`, `HMAC-SHA2-512`, `UMAC-64_AT_OPENSSH.COM`, `UMAC-128_AT_OPENSSH.COM`, `HMAC-SHA1-ETM_AT_OPENSSH.COM`, `HMAC-SHA2-256-ETM_AT_OPENSSH.COM`, `HMAC-SHA2-512-ETM_AT_OPENSSH.COM`, `UMAC-64-ETM_AT_OPENSSH.COM`, `UMAC-128-ETM_AT_OPENSSH.COM`
- `pre_auth_msg` (String) Preauthentication message
  - CLI Alias: `preauth-msg`
- `send_pre_auth_msg` (Boolean) Include SSH preauthentication message
  - CLI Alias: `send-preauth-msg`
  - Default value: `false`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--host_key_alg"></a>
### Nested Schema for `host_key_alg`

Optional:

- `ecdsa_sha2_nistp256` (Boolean) ecdsa-sha2-nistp256
  - Default value: `false`
- `ecdsa_sha2_nistp384` (Boolean) ecdsa-sha2-nistp384
  - Default value: `false`
- `ecdsa_sha2_nistp521` (Boolean) ecdsa-sha2-nistp521
  - Default value: `false`
- `rsa_sha2_256` (Boolean) rsa-sha2-256
  - Default value: `false`
- `rsa_sha2_512` (Boolean) rsa-sha2-512
  - Default value: `false`
- `ssh_ed25519` (Boolean) ssh-ed25519
  - Default value: `false`
