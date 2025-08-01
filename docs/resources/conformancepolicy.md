---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_conformancepolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  Conformance Policy
  CLI Alias: conformancepolicy
---

# datapower_conformancepolicy (Resource)

Conformance Policy
  - CLI Alias: `conformancepolicy`

## Example Usage

```terraform
resource "datapower_conformancepolicy" "test" {
  id         = "ConformancePolicy_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `assert_bp10_conformance` (Boolean) BP1.0 Conformance Claim Assertion
  - CLI Alias: `assert-bp10-conformance`
  - Default value: `false`
- `fixup_stylesheets` (List of String) Corrective Stylesheets
  - CLI Alias: `fixup-stylesheets`
- `ignored_requirements` (List of String) Ignored Requirements
  - CLI Alias: `ignored-requirements`
- `log_target` (String) Destination
  - CLI Alias: `report-target`
- `profiles` (Attributes) Profiles
  - CLI Alias: `profiles` (see [below for nested schema](#nestedatt--profiles))
- `reject_include_summary` (Boolean) Include error summary
  - CLI Alias: `reject-include-summary`
  - Default value: `false`
- `reject_level` (String) Reject non-conforming messages
  - CLI Alias: `reject-level`
  - Choices: `never`, `failure`, `warning`
  - Default value: `never`
- `report_level` (String) Record Report
  - CLI Alias: `report-level`
  - Choices: `never`, `failure`, `warning`, `always`
  - Default value: `never`
- `response_log_target` (String) Destination
  - CLI Alias: `response-report-target`
- `response_properties_enabled` (Boolean) Distinct response behavior
  - CLI Alias: `response-properties-enabled`
  - Default value: `false`
- `response_reject_include_summary` (Boolean) Include response error summary
  - CLI Alias: `response-reject-include-summary`
  - Default value: `false`
- `response_reject_level` (String) Reject non-conforming response messages
  - CLI Alias: `response-reject-level`
  - Choices: `never`, `failure`, `warning`
  - Default value: `never`
- `response_report_level` (String) Record Report (response direction)
  - CLI Alias: `response-report-level`
  - Choices: `never`, `failure`, `warning`, `always`
  - Default value: `never`
- `result_is_conformance_report` (Boolean) Use analysis as result
  - CLI Alias: `result-is-conformance-report`
  - Default value: `false`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--profiles"></a>
### Nested Schema for `profiles`

Optional:

- `ap10` (Boolean) WS-I AP 1.0
  - Default value: `true`
- `bp10` (Boolean) WS-I BP 1.0
  - Default value: `false`
- `bp11` (Boolean) WS-I BP 1.1
  - Default value: `true`
- `bsp10` (Boolean) WS-I BSP 1.0
  - Default value: `true`
