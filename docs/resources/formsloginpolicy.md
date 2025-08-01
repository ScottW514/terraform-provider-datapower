---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_formsloginpolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  HTML forms login policy
  CLI Alias: forms-login-policy
---

# datapower_formsloginpolicy (Resource)

HTML forms login policy
  - CLI Alias: `forms-login-policy`

## Example Usage

```terraform
resource "datapower_formsloginpolicy" "test" {
  id                    = "FormsLoginPolicy_name"
  app_domain            = "acc_test_domain"
  login_form            = "/LoginPage.htm"
  use_cookie_attributes = false
  enable_migration      = false
  error_page            = "/ErrorPage.htm"
  logout_page           = "/LogoutPage.htm"
  default_url           = "/"
  username_field        = "j_username"
  password_field        = "j_password"
  redirect_field        = "originalUrl"
  form_processing_url   = "/j_security_check"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `cookie_attributes` (String) Cookie attributes
  - CLI Alias: `cookie-attributes`
  - Reference to: `datapower_cookieattributepolicy:id`
- `default_url` (String) Default
  - CLI Alias: `default-url`
  - Default value: `/`
- `enable_migration` (Boolean) Enable session migration
  - CLI Alias: `enable-migration`
  - Default value: `false`
- `error_page` (String) Error
  - CLI Alias: `error-url`
  - Default value: `/ErrorPage.htm`
- `form_processing_url` (String) Form POST action URL
  - CLI Alias: `post-action-url`
  - Default value: `/j_security_check`
- `form_support_script` (String) Custom processing for form
  - CLI Alias: `script-location`
  - Default value: `store:///Form-Generate-HTML.xsl`
- `form_support_type` (String) Source for form processing
  - CLI Alias: `support-type`
  - Choices: `static`, `custom`
  - Default value: `static`
- `forms_location` (String) Forms storage location
  - CLI Alias: `forms-location`
  - Choices: `backend`, `appliance`
  - Default value: `backend`
- `inactivity_timeout` (Number) Inactivity timeout
  - CLI Alias: `inactivity-timeout`
  - Default value: `600`
- `local_error_page` (String) Local error page
  - CLI Alias: `local-error-page`
  - Default value: `store:///ErrorPage.htm`
- `local_login_form` (String) Local login form
  - CLI Alias: `local-login-page`
  - Default value: `store:///LoginPage.htm`
- `local_logout_page` (String) Local logout page
  - CLI Alias: `local-logout-page`
  - Default value: `store:///LogoutPage.htm`
- `login_form` (String) Login
  - CLI Alias: `login-url`
  - Default value: `/LoginPage.htm`
- `logout_page` (String) Logout
  - CLI Alias: `logout-url`
  - Default value: `/LogoutPage.htm`
- `password_field` (String) Password field name
  - CLI Alias: `password-field`
  - Default value: `j_password`
- `redirect_field` (String) Target URL field name
  - CLI Alias: `redirect-field`
  - Default value: `originalUrl`
- `redirect_url_type` (String) Redirect URL Type
  - CLI Alias: `redirect-url-type`
  - Choices: `urlin`, `host`
  - Default value: `urlin`
- `session_lifetime` (Number) Session lifetime
  - CLI Alias: `session-lifetime`
  - Default value: `10800`
- `shared_secret` (String) Shared secret
  - CLI Alias: `shared-secret`
  - Reference to: `datapower_cryptosskey:id`
- `ssl_port` (Number) TLS port
  - CLI Alias: `ssl-port`
  - Range: `1`-`65535`
  - Default value: `8080`
- `use_cookie_attributes` (Boolean) Attach cookie attribute policy
  - CLI Alias: `use-cookie-attribute`
  - Default value: `false`
- `use_ssl_for_login` (Boolean) Use TLS for Login
  - CLI Alias: `use-ssl`
  - Default value: `true`
- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `username_field` (String) Username field name
  - CLI Alias: `username-field`
  - Default value: `j_username`
