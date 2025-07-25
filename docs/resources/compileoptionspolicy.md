---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_compileoptionspolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Compile Options Policy
  CLI Alias: compile-options
---

# datapower_compileoptionspolicy (Resource)

Compile Options Policy
  - CLI Alias: `compile-options`

## Example Usage

```terraform
resource "datapower_compileoptionspolicy" "test" {
  id         = "CompileOptionsPolicy_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `allow_soap_enc_array` (String) Specifically Allow xsi:type='SOAP-ENC:Array' Rule
  - CLI Alias: `allow-soap-enc-array`
  - Reference to: `datapower_urlmap:id`
- `allow_xop_include` (String) Accept MTOM/XOP Optimized Messages
  - CLI Alias: `allow-xop-include`
  - Reference to: `datapower_urlmap:id`
- `debug` (String) Debug Rule
  - CLI Alias: `debug`
  - Reference to: `datapower_urlmap:id`
- `disallow_xg4` (String) XML Hardware Acceleration Disallowed Rule
  - CLI Alias: `disallow-xg4`
  - Reference to: `datapower_urlmap:id`
- `minimum_escaping` (String) Minimum Output Escaping Rule
  - CLI Alias: `minesc`
  - Reference to: `datapower_urlmap:id`
- `prefer_xg4` (String) XML Hardware Acceleration Preferred Rule
  - CLI Alias: `prefer-xg4`
  - Reference to: `datapower_urlmap:id`
- `profile` (String) Profile Rule
  - CLI Alias: `profile`
  - Reference to: `datapower_urlmap:id`
- `stack_size` (Number) Maximum Stack Size
  - CLI Alias: `stack-size`
  - Range: `10240`-`104857600`
  - Default value: `524288`
- `stream` (String) Streaming Rule
  - CLI Alias: `stream`
  - Reference to: `datapower_urlmap:id`
- `strict` (Boolean) Strict
  - CLI Alias: `strict`
  - Default value: `false`
- `try_stream` (String) Attempt Streaming Rule
  - CLI Alias: `try-stream`
  - Reference to: `datapower_urlmap:id`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `validate_soap_enc_array` (String) Validate SOAP 1.1 Encoding Rule
  - CLI Alias: `validate-soap-enc-array`
  - Reference to: `datapower_urlmap:id`
- `wildcards_ignore_xsi_type` (String) Wildcards Ignore xsi:type Rule
  - CLI Alias: `wildcards-ignore-xsi-type`
  - Reference to: `datapower_urlmap:id`
- `wsdl_strict_soap_version` (Boolean) Strict SOAP Envelope Version
  - CLI Alias: `wsdl-strict-soap-version`
  - Default value: `false`
- `wsdl_validate_body` (String) Validate Message Body
  - CLI Alias: `wsdl-validate-body`
  - Choices: `strict`, `lax`, `skip`
  - Default value: `strict`
- `wsdl_validate_faults` (String) Validate Message Fault details
  - CLI Alias: `wsdl-validate-faults`
  - Choices: `strict`, `lax`, `skip`
  - Default value: `strict`
- `wsdl_validate_headers` (String) Validate Message Headers
  - CLI Alias: `wsdl-validate-headers`
  - Choices: `strict`, `lax`, `skip`
  - Default value: `lax`
- `wsdl_wrapped_faults` (Boolean) Require wrappers on fault-details specified by type
  - CLI Alias: `wsdl-wrapped-faults`
  - Default value: `false`
- `wsi_validation` (String) WS-I Basic Profile Validation
  - CLI Alias: `wsi-validate`
  - Choices: `ignore`, `warn`, `fail`
  - Default value: `ignore`
- `xacml_debug` (Boolean) Debug XACML Policy
  - CLI Alias: `xacml-debug`
  - Default value: `false`
- `xslt_version` (String) XSLT Version
  - CLI Alias: `xslt-version`
  - Choices: `XSLT10`, `XSLT10_IT23272`, `XSLT20`, `StylesheetSpecified`
  - Default value: `XSLT10`
