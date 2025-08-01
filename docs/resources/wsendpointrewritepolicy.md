---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_wsendpointrewritepolicy Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  WS-Proxy Endpoint Rewrite
  CLI Alias: wsm-endpointrewrite
---

# datapower_wsendpointrewritepolicy (Resource)

WS-Proxy Endpoint Rewrite
  - CLI Alias: `wsm-endpointrewrite`

## Example Usage

```terraform
resource "datapower_wsendpointrewritepolicy" "test" {
  id         = "WSEndpointRewritePolicy_name"
  app_domain = "acc_test_domain"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `user_summary` (String) Comments
  - CLI Alias: `summary`
- `ws_endpoint_local_rewrite_rule` (Attributes List) Local Rewrite Rules
  - CLI Alias: `listener-rule` (see [below for nested schema](#nestedatt--ws_endpoint_local_rewrite_rule))
- `ws_endpoint_publish_rewrite_rule` (Attributes List) Publish Rewrite Rules
  - CLI Alias: `publisher-rule` (see [below for nested schema](#nestedatt--ws_endpoint_publish_rewrite_rule))
- `ws_endpoint_remote_rewrite_rule` (Attributes List) Remote Rewrite Rules
  - CLI Alias: `backend-rule` (see [below for nested schema](#nestedatt--ws_endpoint_remote_rewrite_rule))
- `ws_endpoint_subscription_local_rule` (Attributes List) Subscription Local Rewrite Rule
  - CLI Alias: `subscription-listener-rule` (see [below for nested schema](#nestedatt--ws_endpoint_subscription_local_rule))
- `ws_endpoint_subscription_publish_rule` (Attributes List) Subscription Publish Rewrite Rule
  - CLI Alias: `subscription-publisher-rule` (see [below for nested schema](#nestedatt--ws_endpoint_subscription_publish_rule))
- `ws_endpoint_subscription_remote_rule` (Attributes List) Subscription Remote Rewrite Rule
  - CLI Alias: `subscription-backend-rule` (see [below for nested schema](#nestedatt--ws_endpoint_subscription_remote_rule))

<a id="nestedatt--ws_endpoint_local_rewrite_rule"></a>
### Nested Schema for `ws_endpoint_local_rewrite_rule`

Optional:

- `front_protocol` (String) Local Endpoint Handler
- `frontside_port_suffix` (String) Local Endpoint WSDL Port Name Suffix
- `local_endpoint_hostname` (String) Local Endpoint Host
  - CLI Alias: `local-endpoint-hostname`
  - Default value: `0.0.0.0`
- `local_endpoint_port` (Number) Local Endpoint Port
  - CLI Alias: `local-endpoint-port`
- `local_endpoint_protocol` (String) Local Endpoint Protocol
  - CLI Alias: `local-endpoint-protocol`
  - Choices: `default`, `http`, `https`
  - Default value: `default`
- `local_endpoint_uri` (String) Local Endpoint URI
  - CLI Alias: `local-endpoint-uri`
- `service_port_match_regexp` (String) Service Port Match Expression
  - CLI Alias: `service-port-match`
  - Default value: `.*`
- `use_front_protocol` (Boolean) Use Local Endpoint Handler
  - Default value: `false`
- `wsdl_binding_protocol` (String) Local Endpoint WSDL Binding Protocol
  - Choices: `default`, `soap-11`, `soap-12`, `http-get`, `http-post`
  - Default value: `default`


<a id="nestedatt--ws_endpoint_publish_rewrite_rule"></a>
### Nested Schema for `ws_endpoint_publish_rewrite_rule`

Optional:

- `published_endpoint_hostname` (String) Published Endpoint Host
  - CLI Alias: `published-endpoint-hostname`
- `published_endpoint_port` (Number) Published Endpoint Port
  - CLI Alias: `published-endpoint-port`
- `published_endpoint_protocol` (String) Published Endpoint Protocol
  - CLI Alias: `published-endpoint-protocol`
  - Choices: `default`, `http`, `https`
  - Default value: `default`
- `published_endpoint_uri` (String) Published Endpoint URI
  - CLI Alias: `published-endpoint-path`
- `service_port_match_regexp` (String) Service Port Match Expression
  - CLI Alias: `service-port-match`
  - Default value: `.*`


<a id="nestedatt--ws_endpoint_remote_rewrite_rule"></a>
### Nested Schema for `ws_endpoint_remote_rewrite_rule`

Optional:

- `remote_endpoint_hostname` (String) Remote Endpoint Host
  - CLI Alias: `remote-endpoint-hostname`
- `remote_endpoint_port` (Number) Remote Endpoint Port
  - CLI Alias: `remote-endpoint-port`
- `remote_endpoint_protocol` (String) Remote Endpoint Protocol
  - CLI Alias: `remote-endpoint-protocol`
  - Choices: `default`, `http`, `https`, `dpmq`, `mq`, `idgmq`, `dptibems`, `tibems`, `dpwasjms`
  - Default value: `default`
- `remote_endpoint_uri` (String) Remote Endpoint URI
  - CLI Alias: `remote-endpoint-uri`
- `remote_mq_manager` (String) IBM MQ v9+
  - CLI Alias: `remote-idg-mq-qm`
- `remote_mqqm` (String) IBM MQ
  - CLI Alias: `remote-mq-qm`
- `remote_tibco_ems` (String) TIBCO EMS
  - CLI Alias: `remote-tibems-server`
  - Reference to: `datapower_tibcoemsserver:id`
- `remote_web_sphere_jms` (String) WebSphere JMS
  - CLI Alias: `remote-wasjms-server`
  - Reference to: `datapower_webspherejmsserver:id`
- `service_port_match_regexp` (String) Service Port Match Expression
  - CLI Alias: `service-port-match`
  - Default value: `.*`


<a id="nestedatt--ws_endpoint_subscription_local_rule"></a>
### Nested Schema for `ws_endpoint_subscription_local_rule`

Optional:

- `front_protocol` (String) Local Endpoint Handler
- `frontside_port_suffix` (String) Local Endpoint WSDL Port Name Suffix
- `local_endpoint_hostname` (String) Local Endpoint Host
  - CLI Alias: `local-endpoint-hostname`
  - Default value: `0.0.0.0`
- `local_endpoint_port` (Number) Local Endpoint Port
  - CLI Alias: `local-endpoint-port`
- `local_endpoint_protocol` (String) Local Endpoint Protocol
  - CLI Alias: `local-endpoint-protocol`
  - Choices: `default`, `http`, `https`
  - Default value: `default`
- `local_endpoint_uri` (String) Local Endpoint URI
  - CLI Alias: `local-endpoint-uri`
- `subscription` (String) Subscription
  - CLI Alias: `subscription`
- `use_front_protocol` (Boolean) Use Local Endpoint Handler
  - Default value: `false`
- `wsdl_binding_protocol` (String) Local Endpoint WSDL Binding Protocol
  - Choices: `default`, `soap-11`, `soap-12`, `http-get`, `http-post`
  - Default value: `default`


<a id="nestedatt--ws_endpoint_subscription_publish_rule"></a>
### Nested Schema for `ws_endpoint_subscription_publish_rule`

Optional:

- `published_endpoint_hostname` (String) Published Endpoint Host
  - CLI Alias: `published-endpoint-hostname`
- `published_endpoint_port` (Number) Published Endpoint Port
  - CLI Alias: `published-endpoint-port`
- `published_endpoint_protocol` (String) Published Endpoint Protocol
  - CLI Alias: `published-endpoint-protocol`
  - Choices: `default`, `http`, `https`
  - Default value: `default`
- `published_endpoint_uri` (String) Published Endpoint URI
  - CLI Alias: `published-endpoint-path`
- `subscription` (String) Subscription
  - CLI Alias: `subscription`


<a id="nestedatt--ws_endpoint_subscription_remote_rule"></a>
### Nested Schema for `ws_endpoint_subscription_remote_rule`

Optional:

- `remote_endpoint_hostname` (String) Remote Endpoint Host
  - CLI Alias: `remote-endpoint-hostname`
- `remote_endpoint_port` (Number) Remote Endpoint Port
  - CLI Alias: `remote-endpoint-port`
- `remote_endpoint_protocol` (String) Remote Endpoint Protocol
  - CLI Alias: `remote-endpoint-protocol`
  - Choices: `default`, `http`, `https`, `dpmq`, `mq`, `idgmq`, `dptibems`, `tibems`, `dpwasjms`
  - Default value: `default`
- `remote_endpoint_uri` (String) Remote Endpoint URI
  - CLI Alias: `remote-endpoint-uri`
- `remote_mq_manager` (String) Remote IBM MQ v9+ Queue Manager
  - CLI Alias: `remote-idg-mq-qm`
- `remote_mqqm` (String) Remote IBM MQ Queue Manager
  - CLI Alias: `remote-mq-qm`
- `remote_tibco_ems` (String) Remote TIBCO EMS
  - CLI Alias: `remote-tibems-server`
  - Reference to: `datapower_tibcoemsserver:id`
- `remote_web_sphere_jms` (String) Remote WebSphere JMS
  - CLI Alias: `remote-wasjms-server`
  - Reference to: `datapower_webspherejmsserver:id`
- `subscription` (String) Subscription
  - CLI Alias: `subscription`
