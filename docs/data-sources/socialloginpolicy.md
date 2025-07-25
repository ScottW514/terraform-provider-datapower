---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_socialloginpolicy Data Source - terraform-provider-datapower"
subcategory: ""
description: |-
  Social Login Policy
---

# datapower_socialloginpolicy (Data Source)

Social Login Policy

## Example Usage

```terraform
data "datapower_socialloginpolicy" "test" {
  depends_on = [datapower_socialloginpolicy.test]
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
- `client_grant` (String) Client grant type
- `client_id` (String) Client ID
- `client_optional_query_params` (String) Client Optional Query Parameters
- `client_redirect_uri` (String) Client redirection URI
- `client_scope` (String) Scope
- `client_secret` (String) Client secret
- `client_ssl_profile` (String) TLS client profile
- `id` (String) Name of the object. Must be unique among object types in application domain.
- `jwt_validator` (String) JWT Validator
- `provider_az_endpoint` (String) Authorization endpoint URL
- `provider_token_endpoint` (String) Token endpoint URL
- `social_provider` (String) Social login provider
- `user_summary` (String) Comments
- `validate_jwt_token` (Boolean) Enable JWT token validation
