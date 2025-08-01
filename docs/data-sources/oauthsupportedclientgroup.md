---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_oauthsupportedclientgroup Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  OAuth Client Group
---

# datapower_oauthsupportedclientgroup (Data Source)

OAuth Client Group

## Example Usage

```terraform
data "datapower_oauthsupportedclientgroup" "test" {
  depends_on = [datapower_oauthsupportedclientgroup.test]
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to

### Read-Only

- `result` (Attributes List) List of objects (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `app_domain` (String) The name of the application domain the object belongs to
- `client` (List of String) Client
- `client_template` (String) OAuth Client Template
- `customized` (Boolean) Customized OAuth
- `customized_type` (String) Customization Type
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `o_auth_role` (Attributes) OAuth Role
  - CLI Alias: `oauth-role` (see [below for nested schema](#nestedatt--result--o_auth_role))
- `template_process_url` (String) OAuth Template Process
- `user_summary` (String) Comments

<a id="nestedatt--result--o_auth_role"></a>
### Nested Schema for `result.o_auth_role`

Read-Only:

- `azsvr` (Boolean) Authorization and Token Endpoints
  - Default value: `false`
- `rssvr` (Boolean) Enforcement Point for Resource Server
  - Default value: `false`
