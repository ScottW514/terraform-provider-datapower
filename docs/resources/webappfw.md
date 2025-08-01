---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_webappfw Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Web Application Firewall
  CLI Alias: web-application-firewall
---

# datapower_webappfw (Resource)

Web Application Firewall
  - CLI Alias: `web-application-firewall`

## Example Usage

```terraform
resource "datapower_webappfw" "test" {
  id                       = "WebAppFW_name"
  app_domain               = "acc_test_domain"
  remote_address           = "10.10.10.10"
  style_policy             = "TestAccAppSecurityPolicy"
  xml_manager              = "default"
  front_timeout            = 120
  back_timeout             = 120
  front_persistent_timeout = 180
  back_persistent_timeout  = 180
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `remote_address` (String) Remote Host
  - CLI Alias: `remote-address`
- `style_policy` (String) Security Policy
  - CLI Alias: `security-policy`
  - Reference to: `datapower_appsecuritypolicy:id`

### Optional

- `allow_cache_control_header` (Boolean) Allow Cache-Control Header
  - CLI Alias: `allow-cache-control`
  - Default value: `false`
- `back_http_version` (String) HTTP Version to Server
  - CLI Alias: `http-back-version`
  - Choices: `HTTP/1.0`, `HTTP/1.1`
  - Default value: `HTTP/1.1`
- `back_persistent_timeout` (Number) Back Persistent Timeout
  - CLI Alias: `back-persistent-timeout`
  - Range: `0`-`86400`
  - Default value: `180`
- `back_timeout` (Number) Back Side Timeout
  - CLI Alias: `back-timeout`
  - Range: `1`-`86400`
  - Default value: `120`
- `debug_history` (Number) Transaction History
  - CLI Alias: `debug-history`
  - Range: `10`-`250`
  - Default value: `25`
- `debug_mode` (String) Probe setting
  - CLI Alias: `debug-mode`
  - Choices: `on`, `off`, `unbounded`
  - Default value: `off`
- `debug_trigger` (Attributes List) Probe Triggers
  - CLI Alias: `debug-trigger` (see [below for nested schema](#nestedatt--debug_trigger))
- `delay_errors` (Boolean) Delay Error Messages
  - CLI Alias: `delay-errors`
  - Default value: `true`
- `delay_errors_duration` (Number) Duration to Delay Error Messages
  - CLI Alias: `delay-errors-duration`
  - Range: `250`-`300000`
  - Default value: `1000`
- `do_chunked_upload` (Boolean) Allow Chunked Uploads
  - CLI Alias: `chunked-uploads`
  - Default value: `false`
- `do_host_rewriting` (Boolean) Rewrite Host Names When Gatewaying
  - CLI Alias: `host-rewriting`
  - Default value: `true`
- `error_policy` (String) Error Policy
  - CLI Alias: `error-policy`
  - Reference to: `datapower_webapperrorhandlingpolicy:id`
- `follow_redirects` (Boolean) Follow Redirects
  - CLI Alias: `follow-redirects`
  - Default value: `true`
- `front_http_version` (String) HTTP Response Version
  - CLI Alias: `http-front-version`
  - Choices: `HTTP/1.0`, `HTTP/1.1`
  - Default value: `HTTP/1.1`
- `front_persistent_timeout` (Number) Front Persistent Timeout
  - CLI Alias: `front-persistent-timeout`
  - Range: `0`-`86400`
  - Default value: `180`
- `front_side` (Attributes List) Source Addresses
  - CLI Alias: `listen-on` (see [below for nested schema](#nestedatt--front_side))
- `front_timeout` (Number) Front Side Timeout
  - CLI Alias: `front-timeout`
  - Range: `1`-`86400`
  - Default value: `120`
- `http_client_ip_label` (String) HTTP Client IP Label
  - CLI Alias: `http-client-ip-label`
  - Default value: `X-Client-IP`
- `http_log_cor_id_label` (String) HTTP Global Transaction ID Label
  - CLI Alias: `http-global-tranID-label`
  - Default value: `X-Global-Transaction-ID`
- `priority` (String) Service Priority
  - CLI Alias: `priority`
  - Choices: `unknown`, `high-min`, `high`, `high-max`, `normal-min`, `normal`, `normal-max`, `low-min`, `low`, `low-max`
  - Default value: `normal`
- `remote_port` (Number) Remote Port
  - CLI Alias: `remote-port`
  - Range: `1`-`65535`
  - Default value: `80`
- `request_side_security` (Boolean) Request Security
  - CLI Alias: `request-security`
  - Default value: `true`
- `response_side_security` (Boolean) Response Security
  - CLI Alias: `response-security`
  - Default value: `true`
- `rewrite_errors` (Boolean) Rewrite Error Messages
  - CLI Alias: `rewrite-errors`
  - Default value: `true`
- `ssl_client` (String) TLS client profile
  - CLI Alias: `ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ssl_config_type` (String) TLS type
  - CLI Alias: `ssl-config-type`
  - Choices: `server`, `sni`
  - Default value: `server`
- `ssl_server` (String) TLS server profile
  - CLI Alias: `ssl-server`
  - Reference to: `datapower_sslserverprofile:id`
- `sslsni_server` (String) TLS SNI server profile
  - CLI Alias: `ssl-sni-server`
  - Reference to: `datapower_sslsniserverprofile:id`
- `stream_output_to_back` (String) Stream Output to Back
  - CLI Alias: `stream-output-to-back`
  - Choices: `buffer-until-verification`, `stream-until-infraction`
  - Default value: `buffer-until-verification`
- `stream_output_to_front` (String) Stream Output to Front
  - CLI Alias: `stream-output-to-front`
  - Choices: `buffer-until-verification`, `stream-until-infraction`
  - Default value: `buffer-until-verification`
- `uri_normalization` (Boolean) Normalize URI
  - CLI Alias: `uri-normalization`
  - Default value: `true`
- `url_rewrite_policy` (String) Header and URL Rewrite Policy
  - CLI Alias: `urlrewrite-policy`
  - Reference to: `datapower_urlrewritepolicy:id`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `xml_manager` (String) XML Manager
  - CLI Alias: `xml-manager`
  - Reference to: `datapower_xmlmanager:id`
  - Default value: `default`

<a id="nestedatt--debug_trigger"></a>
### Nested Schema for `debug_trigger`

Optional:

- `client_ip` (String) Client-IP Match
- `in_url_match` (String) Inbound URL Match
- `out_url_match` (String) Outbound URL Match
- `rule_match` (String) Processing Rule Match
- `rule_type` (String) Processing Type
  - Choices: `all`, `response`, `request`, `call`, `error`, `scheduled`, `lbhealth`
- `x_path` (String) XPath Expression Match


<a id="nestedatt--front_side"></a>
### Nested Schema for `front_side`

Optional:

- `credential_charset` (String) Credential Character Set
  - CLI Alias: `credential-charset`
  - Choices: `protocol`, `ascii`, `utf8`, `big5`, `cp1250`, `cp1251`, `cp1252`, `cp1253`, `cp1254`, `cp1255`, `cp1256`, `cp1257`, `cp1258`, `euc_jp`, `euc_kr`, `gb18030`, `gb2312`, `iso2022_jp`, `iso2022_kr`, `iso8859_1`, `iso8859_2`, `iso8859_4`, `iso8859_5`, `iso8859_6`, `iso8859_7`, `iso8859_8`, `iso8859_9`, `iso8859_15`, `sjis`, `tis620`, `unicode_le`
- `local_address` (String) Local IP Address
  - Default value: `0.0.0.0`
- `local_port` (Number) Port
  - Range: `1`-`65535`
  - Default value: `3000`
- `use_ssl` (Boolean) TLS
  - Default value: `false`
